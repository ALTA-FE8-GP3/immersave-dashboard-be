package delivery

import (
	"net/http"
	"project/immersive-dashboard/features/user"
	"project/immersive-dashboard/middlewares"
	"project/immersive-dashboard/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userUsecase user.UsecaseInterface
}

func New(e *echo.Echo, usecase user.UsecaseInterface) {
	handler := &UserDelivery{
		userUsecase: usecase,
	}

	e.GET("/users", handler.GetUser, middlewares.JWTMiddleware())
	e.POST("/login", handler.LoginUser)
	e.POST("/users", handler.PostData, middlewares.JWTMiddleware())
	e.PUT("/users/:id", handler.UpdateUser, middlewares.JWTMiddleware())
	e.DELETE("/users/:id", handler.DeleteDataUser, middlewares.JWTMiddleware())
}

func (delivery *UserDelivery) GetUser(c echo.Context) error {
	result, err := delivery.userUsecase.GetAllUser()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail get all product data"))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("get all products data", FromCoreList(result)))

}

func (delivery *UserDelivery) LoginUser(c echo.Context) error {

	var userRequest_Login UserRequest
	errBind := c.Bind(&userRequest_Login)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("Data Doesn't Exist"))
	}

	Token_JWT, err := delivery.userUsecase.PostLogin(ToCore(userRequest_Login))
	claim, _ := middlewares.ExtractClaims(Token_JWT)
	role := claim["userRole"].(string)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Data Doesn't Exist"))
	}

	return c.JSON(http.StatusOK, helper.Success_Login("Success Login", Token_JWT, role))

}

func (delivery *UserDelivery) PostData(c echo.Context) error {
	var userRequestData UserRequest
	errBind := c.Bind(&userRequestData)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("Fail Bind User Data"))
	}

	row, err := delivery.userUsecase.PostData(ToCore(userRequestData))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Fail Input User Data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Insert Row Affected Is Not 1"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("Success Insert"))

}

func (delivery *UserDelivery) UpdateUser(c echo.Context) error {

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}
	var userUpdate UserRequest
	errBind := c.Bind(&userUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("Fail Bind User Data"))
	}

	userUpdateCore := ToCore(userUpdate)
	userUpdateCore.ID = uint(id_conv)

	row, err := delivery.userUsecase.PutData(userUpdateCore)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Fail Update User Data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Update Row Affected Is Not 1"))
	}
	return c.JSON(http.StatusOK, helper.Success_Resp("Success Update Data"))
}

func (delivery *UserDelivery) DeleteDataUser(c echo.Context) error {

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}

	row, err := delivery.userUsecase.DeleteUser(id_conv)

	if err != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Fail Delete  Data"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("Success Delete  Data"))
}
