package factory

import (
	userData "project/immersive-dashboard/features/user/data"
	userDelivery "project/immersive-dashboard/features/user/delivery"
	userUsecase "project/immersive-dashboard/features/user/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)
}
