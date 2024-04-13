package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ingtk/MaruBatsuGame/api/db"
	"github.com/ingtk/MaruBatsuGame/api/model"
	"github.com/labstack/echo/v4"
	"go.uber.org/mock/gomock"
)

// ほんとはmockを使ってテストした方がよい
type testClock struct{ now time.Time }

func (c *testClock) Now() time.Time {
	return c.now
}

type testIDGenerator struct{}

func (*testIDGenerator) Generate() (string, error) {
	return "aaa", nil
}

type testRandom struct{}

func (*testRandom) Int() int {
	return 100
}

func Test_api_GameStart(t *testing.T) {
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
			name: "待機中のゲームがない",
			mock: func(mockdb *db.MockDatabase) {
				mockdb.EXPECT().PopEmptyGame(gomock.Any(), "user1").Return(nil, nil)
				mockdb.EXPECT().PutGame(gomock.Any(), &model.Game{
					ID:          "aaa",
					HostUserID:  "user1",
					GuestUserID: "",
					Board:       model.Board{{"", "", ""}, {"", "", ""}, {"", "", ""}},
					Turn:        "",
					Winner:      "",
					StartedAt:   nil,
					CreatedAt:   now,
				}).Return(nil)
			},
		},
		{
			name: "待機中の相手とマッチング",
			mock: func(mockdb *db.MockDatabase) {
				mockdb.EXPECT().PopEmptyGame(gomock.Any(), "user1").Return(&model.Game{
					ID:          "aaa",
					HostUserID:  "user2",
					GuestUserID: "",
					Board:       model.Board{},
					Turn:        "",
					Winner:      "",
					StartedAt:   nil,
					CreatedAt:   now,
				}, nil)
				mockdb.EXPECT().PutGame(gomock.Any(), &model.Game{
					ID:          "aaa",
					HostUserID:  "user2",
					GuestUserID: "user1",
					Board:       model.Board{{"", "", ""}, {"", "", ""}, {"", "", ""}},
					Turn:        "user1",
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
				clock:       &testClock{now},
				idGenerator: &testIDGenerator{},
				random:      &testRandom{},
			}
			if tt.mock != nil {
				tt.mock(db)
			}
			req := httptest.NewRequest(echo.GET, "/game/start", nil)
			cookie := &http.Cookie{
				Name:   cookieName,
				Value:  "user1",
				Path:   "/",
				Domain: "localhost",
			}
			req.AddCookie(cookie)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			if err := api.GameStart(c); (err != nil) != tt.wantErr {
				t.Errorf("api.GameStart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
