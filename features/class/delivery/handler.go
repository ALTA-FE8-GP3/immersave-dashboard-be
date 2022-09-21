package delivery

import (
	"net/http"
	"project/immersive-dashboard/features/class"
	"project/immersive-dashboard/middlewares"
	"project/immersive-dashboard/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ClassDelivery struct {
	classUsecase class.UsecaseInterface
}

func New(e *echo.Echo, usecase class.UsecaseInterface) {
	handler := &ClassDelivery{
		classUsecase: usecase,
	}

	e.GET("/class", handler.GetAllClass, middlewares.JWTMiddleware())
	e.POST("/class", handler.PostData, middlewares.JWTMiddleware())
	e.PUT("/class/:id", handler.UpdateClass, middlewares.JWTMiddleware())
}

func (delivery *ClassDelivery) GetAllClass(c echo.Context) error {
	result, err := delivery.classUsecase.GetAllClass()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail get data"))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("get data", FromCoreList(result)))

}

func (delivery *ClassDelivery) PostData(c echo.Context) error {
	var classRequestData ClassRequest
	errBind := c.Bind(&classRequestData)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("Fail Bind Data"))
	}

	row, err := delivery.classUsecase.PostData(ToCore(classRequestData))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Fail Input Data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Insert Row Affected Is Not 1"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("Success Insert"))

}

func (delivery *ClassDelivery) UpdateClass(c echo.Context) error {

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}
	var classUpdate ClassRequest
	errBind := c.Bind(&classUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("Fail Bind User Data"))
	}

	classUpdateCore := ToCore(classUpdate)
	classUpdateCore.ID = uint(id_conv)

	row, err := delivery.classUsecase.PutData(classUpdateCore)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Fail Update Data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Update Row Affected Is Not 1"))
	}
	return c.JSON(http.StatusOK, helper.Success_Resp("Success Update Data"))
}
