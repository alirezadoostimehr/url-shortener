package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"url-shortener/model"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

type Submit struct {
	DB *gorm.DB
}

func (s *Submit) SubmitHandler(c echo.Context) error {
	inp := make(map[string]string)
	err := json.NewDecoder(c.Request().Body).Decode(&inp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	url := model.Url{
		Lng: inp["address"],
		Srt: "/receive/" + RandStringBytes(6),
	}

	c.Response().Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(c.Response()).Encode(url)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err = s.DB.Create(&url).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.Response().WriteHeader(http.StatusOK)
	return nil
}
