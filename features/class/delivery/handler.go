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
	e.GET("/class/:id", handler.GetClassById, middlewares.JWTMiddleware())
	e.POST("/class", handler.PostData, middlewares.JWTMiddleware())
	e.PUT("/class/:id", handler.UpdateClass, middlewares.JWTMiddleware())
	e.DELETE("/class/:id", handler.DeleteClass, middlewares.JWTMiddleware())

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
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind data"))
	}

	row, err := delivery.classUsecase.PostData(ToCore(classRequestData))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail input data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("insert row affected is not 1"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("success insert"))

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
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind user data"))
	}

	classUpdateCore := ToCore(classUpdate)
	classUpdateCore.ID = uint(id_conv)

	row, err := delivery.classUsecase.PutData(classUpdateCore)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail update data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("update row affected is not 1"))
	}
	return c.JSON(http.StatusOK, helper.Success_Resp("success update data"))
}

func (delivery *ClassDelivery) GetClassById(c echo.Context) error {

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}

	result, err := delivery.classUsecase.GetClassById(id_conv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail get data"))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("success get data", FromCore(result)))

}

func (delivery *ClassDelivery) DeleteClass(c echo.Context) error {

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}

	row, err := delivery.classUsecase.DeleteClass(id_conv)

	if err != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail delete data"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("success delete data"))
}
