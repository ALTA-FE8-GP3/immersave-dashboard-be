package delivery

import (
	"net/http"
	"project/immersive-dashboard/features/user"
	"project/immersive-dashboard/middlewares"
	"project/immersive-dashboard/utils/helper"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userUsecase user.UsecaseInterface
}

func New(e *echo.Echo, usecase user.UsecaseInterface) {
	handler := &UserDelivery{
		userUsecase: usecase,
	}

	e.POST("/login", handler.LoginUser)
	e.POST("/users", handler.PostData, middlewares.JWTMiddleware())
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
