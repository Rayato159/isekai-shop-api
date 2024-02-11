package controller

import "github.com/labstack/echo/v4"

type playerControllerImpl struct{}

func NewPlayerControllerImpl() PlayerController {
	return &playerControllerImpl{}
}

func (p *playerControllerImpl) EditProfile(pctx echo.Context) error {
	return nil
}
