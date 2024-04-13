package main

import (
	// "fmt"

	"net/http"

	"github.com/ingtk/MaruBatsuGame/api/api"
	"github.com/ingtk/MaruBatsuGame/api/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// TOOD: 本番環境ではCORSを設定する
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
	db, err := db.NewDatabase()
	if err != nil {
		panic(err)
	}

	api, err := api.NewAPI(db)
	if err != nil {
		panic(err)
	}


	// Routes
	// e.GET("/ws", hello)
	// e.POST("/auth", authorize)
	e.POST("/game/start", api.GameStart)
	e.GET("/game/:game_id/status", api.GameStatus)
	e.POST("/game/:game_id/play_turn", api.PlayTurn)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// type Param struct {
// 	Matrix [3][3]int `json:"matrix"`
// 	Turn   int       `json:"turn"`
// }
