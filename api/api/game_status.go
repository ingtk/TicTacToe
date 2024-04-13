package api

import (
	"github.com/ingtk/MaruBatsuGame/api/model"
	"github.com/labstack/echo/v4"
)

type gameStatusResponse struct {
	HostUserID      string    `json:"hostUserId"`
	GuestUserID     string    `json:"guestUserId"`
	CurrentPlayerID string    `json:"currentPlayerId"`
	PlayerTurn      bool      `json:"playerTurn"`
	PlayerWin       *bool     `json:"playerWin"`
	Board           [3][3]int `json:"board"`
	GameStarted     bool      `json:"gameStarted"`
	GameEnded       bool      `json:"gameEnded"`
	Error           string    `json:"error"`
}

func (api *api) GameStatus(c echo.Context) error {
	ctx := c.Request().Context()
	gameID := c.Param("game_id")
	userID, err := auth(c)
	if err != nil {
		return err
	}
	game, err := api.db.GetGameByID(ctx, gameID)
	if err != nil {
		return err
	}
	// resp := gameStatusResponse{}
	if game == nil {
		return c.JSON(404, &gameStatusResponse{Error: "game not found"})
	}
	// 自分が参加しているゲームではない
	if game.HostUserID != userID && game.GuestUserID != userID {
		return c.JSON(400, &gameStatusResponse{Error: "invalid user"})
	}

	return c.JSON(200, toGameStatusResponse(userID, game))
}

func toGameStatusResponse(userID string, game *model.Game) gameStatusResponse {
	resp := gameStatusResponse{}
	resp.HostUserID = game.HostUserID
	resp.GuestUserID = game.GuestUserID
	resp.PlayerTurn = userID == game.Turn
	resp.CurrentPlayerID = game.Turn

	gameEnded := true
	resp.Board = [3][3]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if game.Board[i][j] == "" {
				gameEnded = false
			}
			square := 0
			if game.HostUserID != "" && game.Board[i][j] == game.HostUserID {
				square = 1
			}
			if game.GuestUserID != "" && game.Board[i][j] == game.GuestUserID {
				square = 2
			}
			resp.Board[i][j] = square
		}
	}

	if game.Winner != "" {
		win := userID == game.Winner
		resp.PlayerWin = &win
		gameEnded = true
	}
	if game.StartedAt != nil {
		resp.GameStarted = true
	}
	resp.GameEnded = gameEnded

	return resp
}
