package routes

import (
	"echolabstack/service"
	"echolabstack/ratelimitter"
	"github.com/labstack/echo/v4"
)


func Echoroutes(e *echo.Echo) {
	
	e.Static("/static", "static")
	e.GET("/users", service.NewAPI, ratelimitter.CombinedRateLimiter())

	e.GET("/pdfapi", service.PdfAPI,ratelimitter.CombinedRateLimiter())

}