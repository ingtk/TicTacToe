package db

import (
	"context"

	"github.com/ingtk/MaruBatsuGame/api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func NewDatabase() (*database, error) {
	dsn := "host=localhost user=user password=password dbname=db port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	return &database{db: db}, nil
}

type database struct {
	db *gorm.DB
}

func (d *database) PopEmptyGame(ctx context.Context, userID string) (*model.Game, error) {
	var game model.Game
	err := d.db.
		Where("host_user_id <> ?", userID).
		Where("guest_user_id = ''").
		First(&game).Error
	// v, err := client.LPop(ctx, "empty_rooms").Result()
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	err = game.UnmarshalBoard()
	if err != nil {
		return nil, err
	}
	return &game, nil
}

func (d *database) GetGameByID(ctx context.Context, gameID string) (*model.Game, error) {
	var game model.Game
	err := d.db.WithContext(ctx).
		Where("id = ?", gameID).
		First(&game).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	err = game.UnmarshalBoard()
	if err != nil {
		return nil, err
	}
	return &game, nil
}

func (d *database) PutGame(ctx context.Context, game *model.Game) error {
	game.MarshalBoard()
	err := d.db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.Assignments(
				map[string]any{
					"host_user_id":  game.HostUserID,
					"guest_user_id": game.GuestUserID,
					"turn":          game.Turn,
					"winner":        game.Winner,
					"board":         game.BoardData,
					"started_at":    game.StartedAt,
				},
			),
		}).Create(game).Error
	if err != nil {
		return err
	}
	err = game.MarshalBoard()
	if err != nil {
		return err
	}
	return nil
}
