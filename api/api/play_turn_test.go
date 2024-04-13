package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ingtk/MaruBatsuGame/api/db"
	"github.com/ingtk/MaruBatsuGame/api/model"
	"github.com/labstack/echo/v4"
	"go.uber.org/mock/gomock"
)

func Test_api_PlayTurn(t *testing.T) {
	now := time.Now()
	type fields struct {
		db Database
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func(db *db.MockDatabase)
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			mock: func(mockdb *db.MockDatabase) {
				mockdb.EXPECT().GetGameByID(gomock.Any(), "aaa").Return(&model.Game{
					ID:          "aaa",
					HostUserID:  "user1",
					GuestUserID: "user2",
					Board:       model.Board{},
					Turn:        "user1",
					Winner:      "",
					StartedAt:   &now,
					CreatedAt:   now,
				}, nil)
				mockdb.EXPECT().PutGame(gomock.Any(), &model.Game{
					ID:          "aaa",
					HostUserID:  "user1",
					GuestUserID: "user2",
					Board:       model.Board{{"user1", "", ""}, {"", "", ""}, {"", "", ""}},
					Turn:        "user2",
					Winner:      "",
					StartedAt:   &now,
					CreatedAt:   now,
				}).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			ctrl := gomock.NewController(t)
			db := db.NewMockDatabase(ctrl)
			api := &api{
				db:          db,
				clock:       nil,
				idGenerator: nil,
			}
			if tt.mock != nil {
				tt.mock(db)
			}
			data, err := json.Marshal(playTurnRequest{Y: 0, X: 0})
			if err != nil {
				panic(err)
			}
			req := httptest.NewRequest(echo.POST, "/game/:game_id/play_turn", bytes.NewReader(data))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			cookie := &http.Cookie{
				Name:   cookieName,
				Value:  "user1",
				Path:   "/",
				Domain: "localhost",
			}
			req.AddCookie(cookie)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("game_id")
			c.SetParamValues("aaa")
			if err := api.PlayTurn(c); (err != nil) != tt.wantErr {
				t.Errorf("api.PlayTurn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
