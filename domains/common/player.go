package common

import (
	_playerException "github.com/Rayato159/isekai-shop-api/domains/player/exception"
	"github.com/labstack/echo/v4"
)

func GetPlayerID(pctx echo.Context) (string, error) {
	if playerID, ok := pctx.Get("playerID").(string); !ok || playerID == "" {
		return "", &_playerException.PlayerIDNotFoundException{}
	} else {
		return playerID, nil
	}
}
