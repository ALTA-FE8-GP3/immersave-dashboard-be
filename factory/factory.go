package factory

import (
	userData "project/immersive-dashboard/features/user/data"
	userDelivery "project/immersive-dashboard/features/user/delivery"
	userUsecase "project/immersive-dashboard/features/user/usecase"

	classData "project/immersive-dashboard/features/class/data"
	classDelivery "project/immersive-dashboard/features/class/delivery"
	classUsecase "project/immersive-dashboard/features/class/usecase"

	menteeData "project/immersive-dashboard/features/mentee/data"
	menteeDelivery "project/immersive-dashboard/features/mentee/delivery"
	menteeUsecase "project/immersive-dashboard/features/mentee/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

	classDataFactory := classData.New(db)
	classUsecaseFactory := classUsecase.New(classDataFactory)
	classDelivery.New(e, classUsecaseFactory)

	menteeDataFactory := menteeData.New(db)
	menteeUsecaseFactory := menteeUsecase.New(menteeDataFactory, classDataFactory)
	menteeDelivery.New(e, menteeUsecaseFactory)

}
