package common

import (
	"github.com/labstack/echo/v4"

	_admin "github.com/Rayato159/isekai-shop-api/domains/admin/exception"
)

func GetAdminID(pctx echo.Context) (string, error) {
	if adminID, ok := pctx.Get("adminID").(string); !ok || adminID == "" {
		return "", &_admin.AdminIDNotfound{}
	} else {
		return adminID, nil
	}
}
