package utils

import (
	_player "github.com/Rayato159/isekai-shop-api/pkg/player/exception"
	"github.com/labstack/echo/v4"
)

func GetPlayerID(pctx echo.Context) (string, error) {
	if playerID, ok := pctx.Get("playerID").(string); !ok || playerID == "" {
		return "", &_player.PlayerNotFound{PlayerID: "Unknown"}
	} else {
		return playerID, nil
	}
}
