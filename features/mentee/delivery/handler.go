package delivery

import (
	"fmt"
	"net/http"
	"project/immersive-dashboard/features/mentee"
	"project/immersive-dashboard/middlewares"
	"project/immersive-dashboard/utils/helper"
	"strconv"

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
	e.GET("/mentee", handler.GetMentee, middlewares.JWTMiddleware())
	e.GET("/mentee/:id", handler.GetMenteeById, middlewares.JWTMiddleware())
	e.DELETE("/mentee/:id", handler.DeleteDataMentee, middlewares.JWTMiddleware())
	e.PUT("/mentee/:id", handler.UpdateMenteeId, middlewares.JWTMiddleware())

}

func (delivery *MenteeDelivery) PostMentee(c echo.Context) error {
	var mentee_RequestData MenteeRequest
	errBind := c.Bind(&mentee_RequestData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind data"))
	}
	fmt.Println(mentee_RequestData)
	fmt.Println(ToCore(mentee_RequestData))
	row, err := delivery.menteeUsecase.PostData(ToCore(mentee_RequestData))

	fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Fail Input User Data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Insert Row Affected Is Not 1"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("success insert data"))
}

func (delivery *MenteeDelivery) GetMentee(c echo.Context) error {
	category := c.QueryParam("category")
	status := c.QueryParam("status")
	class_id, err := strconv.Atoi(c.QueryParam("class_id"))
	if err != nil && class_id != 0 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail converse class_id param"))
	}

	dataMentee, err := delivery.menteeUsecase.GetAllMentee(class_id, category, status)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("get all data mentees success", FromCoreList(dataMentee)))
}

func (delivery *MenteeDelivery) DeleteDataMentee(c echo.Context) error {

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}

	row, err := delivery.menteeUsecase.Delete(id_conv)

	if err != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail delete data"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("success delete data"))
}

func (delivery *MenteeDelivery) GetMenteeById(c echo.Context) error {

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}

	result, err := delivery.menteeUsecase.GetMenteeById(id_conv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail get data"))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("success get data", FromCore(result)))

}

func (delivery *MenteeDelivery) UpdateMenteeId(c echo.Context) error {

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}
	var menteeUpdate MenteeRequest
	errBind := c.Bind(&menteeUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind user data"))
	}

	menteeUpdateCore := ToCore(menteeUpdate)
	menteeUpdateCore.ID = uint(id_conv)

	row, err := delivery.menteeUsecase.UpdateMenteeId(menteeUpdateCore)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail update data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("update row affected is not 1"))
	}
	return c.JSON(http.StatusOK, helper.Success_Resp("success update data"))
}
