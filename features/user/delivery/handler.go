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
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("data doesn't exist"))
	}

	Token_JWT, err := delivery.userUsecase.PostLogin(ToCore(userRequest_Login))
	claim, _ := middlewares.ExtractClaims(Token_JWT)
	role := claim["userRole"].(string)
	user := claim["userName"].(string)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("data doesn't exist"))
	}

	return c.JSON(http.StatusOK, helper.Success_Login("success login", Token_JWT, role, user))

}

func (delivery *UserDelivery) PostData(c echo.Context) error {
	var userRequestData UserRequest
	errBind := c.Bind(&userRequestData)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind user data"))
	}

	row, err := delivery.userUsecase.PostData(ToCore(userRequestData))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail input user data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("insert row affected is not 1"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("success insert"))

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
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind user data"))
	}

	userUpdateCore := ToCore(userUpdate)
	userUpdateCore.ID = uint(id_conv)

	row, err := delivery.userUsecase.PutData(userUpdateCore)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail update user data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("update row affected is not 1"))
	}
	return c.JSON(http.StatusOK, helper.Success_Resp("success update data"))
}

func (delivery *UserDelivery) DeleteDataUser(c echo.Context) error {

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}

	row, err := delivery.userUsecase.DeleteUser(id_conv)

	if err != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail delete data"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("success delete data"))
}
