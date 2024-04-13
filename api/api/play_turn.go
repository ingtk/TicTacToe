package api

import (
	"errors"

	"github.com/labstack/echo/v4"
)

type playTurnRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (api *api) PlayTurn(c echo.Context) error {
	ctx := c.Request().Context()
	gameID := c.Param("game_id")
	userID, err := auth(c)
	if err != nil {
		return err
	}

	req := playTurnRequest{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}

	game, err := api.db.GetGameByID(ctx, gameID)
	if err != nil {
		return err
	}
	if game == nil {
		return errors.New("game not found")
	}
	// 自分が参加しているゲームではない
	if game.HostUserID != userID && game.GuestUserID != userID {
		return errors.New("invalid user")
	}

	// 自分のターンではない
	if game.Turn != userID {
		return errors.New("invalid user")
	}

	if game.HostUserID == userID {
		game.Turn = game.GuestUserID
	} else if game.GuestUserID == userID {
		game.Turn = game.HostUserID
	}

	if req.X < 0 || 2 < req.X || req.Y < 0 || 2 < req.Y {
		return errors.New("invalid point")
	}

	// すでに置かれている
	if game.Board[req.Y][req.X] != "" {
		return errors.New("invalid point")
	}

	game.Board[req.Y][req.X] = userID

	game.UpdateWinner()

	err = api.db.PutGame(ctx, game)
	if err != nil {
		return err
	}

	return c.JSON(200, nil)
}
