package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	models "github.com/hizbi-github/gost/new-project-core/models"
)

func SomeHandler(ctx echo.Context) error {
	ctx.Response().Header().Add("Content-Type", "application/json")

	var someRequestStruct models.ApiGenericResponse
	err := ctx.Bind(&someRequestStruct)
	if err != nil {
		logrus.Errorln(err)
		ctx.JSON(http.StatusBadRequest, models.ApiGenericResponse{
			Data:    nil,
			Message: "fix your reqeust body",
		})
		return err
	}

	logrus.Infoln("some log")
	return nil
}
