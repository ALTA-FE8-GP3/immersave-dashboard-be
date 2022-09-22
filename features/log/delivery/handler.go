package delivery

import (
	"fmt"
	"net/http"
	"project/immersive-dashboard/features/log"
	"project/immersive-dashboard/middlewares"
	"project/immersive-dashboard/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type LogDelivery struct {
	logUsecase log.UsecaseInterface
}

func New(e *echo.Echo, usecase log.UsecaseInterface) {
	handler := &LogDelivery{
		logUsecase: usecase,
	}

	e.POST("/mentee", handler.PostLog, middlewares.JWTMiddleware())
	e.GET("/mentee/:id", handler.GetLogById, middlewares.JWTMiddleware())

}

func (delivery *LogDelivery) PostLog(c echo.Context) error {
	var log_RequestData LogRequest
	errBind := c.Bind(&log_RequestData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind data"))
	}
	fmt.Println(log_RequestData)
	fmt.Println(ToCore(log_RequestData))
	row, err := delivery.logUsecase.PostData(ToCore(log_RequestData))

	fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Fail Input User Data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Insert Row Affected Is Not 1"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("success insert data"))
}

func (delivery *LogDelivery) GetLogById(c echo.Context) error {

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}

	result, err := delivery.logUsecase.GetLogById(id_conv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail get data"))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("success get data", FromCore(result)))

}
