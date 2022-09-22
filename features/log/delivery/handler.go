package delivery

import (
	"fmt"
	"net/http"
	"project/immersive-dashboard/config"
	"project/immersive-dashboard/features/log"
	"project/immersive-dashboard/middlewares"
	"project/immersive-dashboard/utils/helper"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type LogDelivery struct {
	logUsecase log.UsecaseInterface
}

func New(e *echo.Echo, usecase log.UsecaseInterface) {
	handler := &LogDelivery{
		logUsecase: usecase,
	}

	e.POST("/log", handler.PostLog, middlewares.JWTMiddleware())
	e.GET("/log/:id", handler.GetLogById, middlewares.JWTMiddleware())
	e.GET("/log", handler.GetAlllog, middlewares.JWTMiddleware())

}

func (delivery *LogDelivery) PostLog(c echo.Context) error {
	var log_RequestData LogRequest
	errBind := c.Bind(&log_RequestData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind data"))
	}

	dataFoto, infoFoto, fotoerr := c.Request().FormFile("url_file")
	if fotoerr != http.ErrMissingFile || fotoerr == nil {
		format, errf := helper.CheckFile(infoFoto.Filename)
		if errf != nil {
			return c.JSON(http.StatusBadRequest, helper.Fail_Resp("Format Error"))
		}
		//checksize
		err_image_size := helper.CheckSize(infoFoto.Size)
		if err_image_size != nil {
			return c.JSON(http.StatusBadRequest, err_image_size)
		}
		//rename
		waktu := fmt.Sprintf("%v", time.Now())
		imageName := strconv.Itoa(int(log_RequestData.MenteeID)) + "_" + strconv.Itoa(int(log_RequestData.UserID)) + waktu + "." + format

		imageaddress, errupload := helper.UploadFileToS3(config.FolderName, imageName, config.FileType, dataFoto)
		if errupload != nil {
			return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail to upload file"))
		}

		log_RequestData.Url_File = imageaddress
	}

	row, err := delivery.logUsecase.PostData(ToCore(log_RequestData))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail input data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("insert row affected is not 1"))
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

func (delivery *LogDelivery) GetAlllog(c echo.Context) error {
	result, err := delivery.logUsecase.GetAlllog()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail get data"))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("get data", FromCoreList(result)))

}
