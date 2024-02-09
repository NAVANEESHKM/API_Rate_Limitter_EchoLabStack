package service

import(
	"github.com/labstack/echo/v4"
	"net/http"

)

func NewAPI(c echo.Context) error {
	return c.String(http.StatusOK, "New API created")
}

func PdfAPI(c echo.Context) error{
	pdfPath := "static/Echo_static.pdf"
	return c.File(pdfPath)

}