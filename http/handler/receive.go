package handler

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"url-shortener/model"
)

type Receive struct {
	DB *gorm.DB
}

func (r *Receive) ReceiveHandler(c echo.Context) error {
	addr := c.Request().URL.Path

	var url model.Url
	if err := r.DB.Model(&model.Url{}).First(&url, "srt = ?", addr).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.Redirect(http.StatusSeeOther, url.Lng)
}
