package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/ingtk/MaruBatsuGame/api/db"
	"github.com/ingtk/MaruBatsuGame/api/model"
	"github.com/labstack/echo/v4"
	"go.uber.org/mock/gomock"
)

func ptr[T bool](v T) *T {
	return &v
}

func Test_api_GameStatus(t *testing.T) {
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
		want    gameStatusResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "マッチング中",
			mock: func(mockdb *db.MockDatabase) {
				mockdb.EXPECT().GetGameByID(gomock.Any(), "").Return(&model.Game{
					ID:          "aaa",
					HostUserID:  "user1",
					GuestUserID: "",
					Board:       model.Board{{"", "", ""}, {"", "", ""}, {"", "", ""}},
					Turn:        "",
					Winner:      "",
					StartedAt:   nil,
					CreatedAt:   now,
				}, nil)
			},
			want: gameStatusResponse{
				// ID:          "aaa",
				HostUserID:      "user1",
				GuestUserID:     "",
				CurrentPlayerID: "",
				Board:           [3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
				PlayerTurn:      false,
				PlayerWin:       nil,
				GameStarted:     false,
				GameEnded:       false,
			},
		},
		{
			name: "対戦中",
			mock: func(mockdb *db.MockDatabase) {
				mockdb.EXPECT().GetGameByID(gomock.Any(), "").Return(&model.Game{
					ID:          "aaa",
					HostUserID:  "user1",
					GuestUserID: "user2",
					Board:       model.Board{{"user1", "", ""}, {"", "user2", ""}, {"", "", ""}},
					Turn:        "user1",
					Winner:      "",
					StartedAt:   &now,
					CreatedAt:   now,
				}, nil)
			},
			want: gameStatusResponse{
				// ID:          "aaa",
				HostUserID:      "user1",
				GuestUserID:     "user2",
				CurrentPlayerID: "user1",
				Board:           [3][3]int{{1, 0, 0}, {0, 2, 0}, {0, 0, 0}},
				PlayerTurn:      true,
				PlayerWin:       nil,
				GameStarted:     true,
				GameEnded:       false,
			},
		},
		{
			name: "ゲーム終了、勝者あり",
			mock: func(mockdb *db.MockDatabase) {
				mockdb.EXPECT().GetGameByID(gomock.Any(), "").Return(&model.Game{
					ID:          "aaa",
					HostUserID:  "user1",
					GuestUserID: "user2",
					Board:       model.Board{{"user1", "", ""}, {"user1", "user2", ""}, {"user1", "user2", ""}},
					Turn:        "user1",
					Winner:      "user1",
					StartedAt:   &now,
					CreatedAt:   now,
				}, nil)
			},
			want: gameStatusResponse{
				// ID:          "aaa",
				HostUserID:      "user1",
				GuestUserID:     "user2",
				CurrentPlayerID: "user1",
				Board:           [3][3]int{{1, 0, 0}, {1, 2, 0}, {1, 2, 0}},
				PlayerTurn:      true,
				PlayerWin:       ptr(true),
				GameStarted:     true,
				GameEnded:       true,
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
				clock:       &testClock{now},
				idGenerator: &testIDGenerator{},
			}
			if tt.mock != nil {
				tt.mock(db)
			}
			req := httptest.NewRequest(echo.GET, "/game/aaa/status", nil)
			cookie := &http.Cookie{
				Name:   cookieName,
				Value:  "user1",
				Path:   "/",
				Domain: "localhost",
			}
			req.AddCookie(cookie)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			if err := api.GameStatus(c); (err != nil) != tt.wantErr {
				t.Errorf("api.GameStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
			resp := gameStatusResponse{}
			err := json.NewDecoder(rec.Body).Decode(&resp)
			if err != nil {
				t.Error(err)
			}
			if diff := cmp.Diff(tt.want, resp); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
