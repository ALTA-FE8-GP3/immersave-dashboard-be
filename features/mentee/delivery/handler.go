package delivery

import (
	"net/http"
	"project/immersive-dashboard/features/mentee"
	"project/immersive-dashboard/middlewares"
	"project/immersive-dashboard/utils/helper"

	"github.com/labstack/echo/v4"
)

type MenteeDelivery struct {
	menteeUsecase mentee.UsecaseInterface
}

func New(e *echo.Echo, usecase mentee.UsecaseInterface) {
	handler := &MenteeDelivery{
		menteeUsecase: usecase,
	}

	e.POST("/mentee", handler.PostMentee, middlewares.JWTMiddleware())
	// e.GET("/carts", handler.GetAllCarts, middlewares.JWTMiddleware())
}

func (delivery *MenteeDelivery) PostMentee(c echo.Context) error {
	var mentee_RequestData MenteeRequest
	errBind := c.Bind(&mentee_RequestData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("Fail Bind Data"))
	}

	row, err := delivery.menteeUsecase.PostData(ToCore(mentee_RequestData))

	if err != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Fail Make New Cart"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("Success Make New Cart"))
}
