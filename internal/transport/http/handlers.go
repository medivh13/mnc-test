package http

import (
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/labstack/echo"
	"github.com/medivh13/mnc-test/internal/services"
	mncTestconst "github.com/medivh13/mnc-test/pkg/common/const"
	"github.com/medivh13/mnc-test/pkg/dto"
	btbErrors "github.com/medivh13/mnc-test/pkg/errors"
)

type HttpHandler struct {
	service services.Services
}

func NewHttpHandler(e *echo.Echo, srv services.Services) {
	handler := &HttpHandler{
		srv,
	}

	e.GET("api/v1/mnc-test/", handler.Ping)
	e.GET("api/v1/mnc-test/palindrome", handler.GetPalindrome)
	e.GET("api/v1/mnc-test/language", handler.GetLanguage)
}

func (h *HttpHandler) Ping(c echo.Context) error {

	version := os.Getenv("VERSION")
	if version == "" {
		version = "Hello Go developers"
	}

	data := version

	return c.JSON(http.StatusOK, data)

}

func (h *HttpHandler) GetPalindrome(c echo.Context) error {
	getDTO := dto.PalindromeReqDTO{}

	if err := c.Bind(&getDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	getDTO.Signature = c.Request().Header.Get("signature")
	getDTO.Text = c.Request().FormValue("text")

	data, err := h.service.Palindrome(&getDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), map[string]string{
			"error": err.Error(),
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mncTestconst.GetDataSuccess,
		Data:    data,
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HttpHandler) GetLanguage(c echo.Context) error {
	getDTO := dto.LanguageReqDTO{}

	if err := c.Bind(&getDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	getDTO.Signature = c.Request().Header.Get("signature")
	getDTO.Id = c.Request().FormValue("id")

	data, err := h.service.GetLangaugeById(&getDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), map[string]string{
			"error": err.Error(),
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mncTestconst.GetDataSuccess,
		Data:    data,
	}

	return c.JSON(http.StatusOK, resp)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case btbErrors.ErrInternalServerError:
		return http.StatusInternalServerError
	case btbErrors.ErrNotFound:
		return http.StatusNotFound
	case btbErrors.ErrConflict:
		return http.StatusConflict
	case btbErrors.ErrInvalidRequest:
		return http.StatusBadRequest
	case btbErrors.ErrFailAuth:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
