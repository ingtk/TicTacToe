package model

import (
	"encoding/json"
	"time"
)

type Game struct {
	ID          string
	HostUserID  string
	GuestUserID string
	Turn        string
	Board       Board `gorm:"-"`
	Winner      string
	StartedAt   *time.Time
	CreatedAt   time.Time
	BoardData   string `gorm:"column:board" json:"-"`
}

func (g *Game) UpdateWinner() {
	g.Winner = g.Board.CheckWinner()
}

type Board [3][3]string

func (g *Game) UnmarshalBoard() error {
	if g.BoardData == "" {
		return nil
	}
	err := json.Unmarshal([]byte(g.BoardData), &g.Board)
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) MarshalBoard() error {
	board, err := json.Marshal(g.Board)
	if err != nil {
		return err
	}
	g.BoardData = string(board)

	return nil
}

// func (b *Board) Scan(value interface{}) error {
// 	bytes, ok := value.([]byte)
// 	if !ok {
// 		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
// 	}

// 	var result Board
// 	err := json.Unmarshal(bytes, &result)
// 	if err != nil {
// 		return err
// 	}
// 	b = &result
// 	return nil
// }

// func (b *Board) Value() (driver.Value, error) {
// 	if b == nil {
// 		return nil, nil
// 	}
// 	return json.Marshal(b)
// }

func (b *Board) CheckWinner() string {
	if b[0][0] == b[0][1] && b[0][1] == b[0][2] && b[0][0] != "" {
		return b[0][0]
	}
	if b[1][0] == b[1][1] && b[1][1] == b[1][2] && b[1][0] != "" {
		return b[1][0]
	}
	if b[2][0] == b[2][1] && b[2][1] == b[2][2] && b[2][0] != "" {
		return b[2][0]
	}
	if b[0][0] == b[1][0] && b[1][0] == b[2][0] && b[0][0] != "" {
		return b[0][0]
	}
	if b[0][1] == b[1][1] && b[1][1] == b[2][1] && b[0][1] != "" {
		return b[0][1]
	}
	if b[0][2] == b[1][2] && b[1][2] == b[2][2] && b[0][2] != "" {
		return b[0][2]
	}
	if b[0][0] == b[1][1] && b[1][1] == b[2][2] && b[0][0] != "" {
		return b[0][0]
	}
	if b[0][2] == b[1][1] && b[1][1] == b[2][0] && b[0][2] != "" {
		return b[0][2]
	}
	return ""
}
