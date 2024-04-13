package api

import (
	"github.com/ingtk/MaruBatsuGame/api/model"
	"github.com/labstack/echo/v4"
)

type GameStartResponse struct {
	GameID string `json:"game_id"`
}

func (api *api) GameStart(c echo.Context) error {
	ctx := c.Request().Context()
	userID, err := auth(c)
	game, err := api.db.PopEmptyGame(ctx, userID)
	if err != nil {
		return err
	}

	// マッチング相手がいない
	if game == nil {
		gameID, err := api.idGenerator.Generate()
		if err != nil {
			return err
		}
		game = &model.Game{
			ID:         gameID,
			HostUserID: userID,
			CreatedAt:  api.clock.Now(),
		}
		err = api.db.PutGame(ctx, game)
		if err != nil {
			return err
		}
		return c.JSON(200, &GameStartResponse{gameID})
	}

	if game.HostUserID == userID {
		return c.JSON(200, &GameStartResponse{game.ID})
	}

	// TODO: 開始時は手番がHostUsr
	game.GuestUserID = userID
	game.Turn = game.HostUserID // TODO: turnを決める
	now := api.clock.Now()
	game.StartedAt = &now

	err = api.db.PutGame(ctx, game)
	if err != nil {
		return err
	}
	return c.JSON(200, &GameStartResponse{game.ID})
}
