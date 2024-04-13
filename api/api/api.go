package api

import (
	"context"
	"time"

	"github.com/ingtk/MaruBatsuGame/api/model"
	"github.com/ingtk/MaruBatsuGame/api/pkg"
	"github.com/labstack/echo/v4"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Database interface {
	PopEmptyGame(ctx context.Context, userID string) (*model.Game, error)
	GetGameByID(ctx context.Context, gameID string) (*model.Game, error)
	PutGame(ctx context.Context, game *model.Game) error
}

type API interface {
	GameStart(c echo.Context) error
	GameStatus(c echo.Context) error
	PlayTurn(c echo.Context) error
}

type api struct {
	db          Database
	clock       pkg.Clock
	idGenerator pkg.IDGenerator
}

type clock struct{}

func (*clock) Now() time.Time {
	return time.Now()
}

type idGenerator struct{}

func (*idGenerator) Generate() (string, error) {
	return gonanoid.New()
}

func NewAPI(db Database) (API, error) {
	return &api{db: db, clock: &clock{}, idGenerator: &idGenerator{}}, nil
}
