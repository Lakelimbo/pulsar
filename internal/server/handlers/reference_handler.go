package handlers

import (
	"log"
	"net/http"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/labstack/echo/v4"
)

var swaggerFile string = "internal/server/docs/swagger.yaml"

func ServeSwagger(c echo.Context) error {
	return c.File(swaggerFile)
}

func InitScalar(c echo.Context) error {
	content, err := scalar.ApiReferenceHTML(&scalar.Options{
		SpecURL: swaggerFile,
		CustomOptions: scalar.CustomOptions{
			PageTitle: "Pulsar API",
		},
		DarkMode: true,
		Theme:    "alternate",
	})

	if err != nil {
		log.Fatalf("Failed to initialize Scalar: %v", err)
	}

	return c.HTML(http.StatusOK, content)
}
