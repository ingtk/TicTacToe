package api

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

const cookieName = "myTicTacToe"

func auth(c echo.Context) (string, error) {
	cookie, err := c.Cookie(cookieName)
	if err != nil && err != http.ErrNoCookie {
		return "", err
	}

	// var cookie *http.Cookie
	if cookie == nil {
		cookie = &http.Cookie{}
	}
	if cookie.Value == "" {
		// TODO: 本来はローカルかどうかで切り替えるべき
		// cookie.HttpOnly = true
		// cookie.Secure = false
		// cookie.SameSite = http.SameSiteStrictMode
		cookie.Domain = "localhost"
		id, err := gonanoid.New()
		if err != nil {
			return "", err
		}
		cookie.Path = "/"
		cookie.Name = cookieName
		cookie.Value = id
		c.SetCookie(cookie)
		log.Printf("%+v", cookie)
	}

	return cookie.Value, nil
}
