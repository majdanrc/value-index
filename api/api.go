package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"value-index/internal/search"
)

type RestAPI struct {
	searchService *search.SearchService
	echo          *echo.Echo
	port          string
}

func NewRest(searchService *search.SearchService, port string) *RestAPI {
	apiService := &RestAPI{
		searchService: searchService,
		echo:          echo.New(),
		port:          port,
	}

	apiService.echo.GET("/endpoint/:value", apiService.getIndex)

	return apiService
}

func (a *RestAPI) Start() error {
	return a.echo.Start(":" + a.port)
}

// getIndex is the single endpoint providing the functionality of the search service
func (a *RestAPI) getIndex(c echo.Context) error {
	valueStr := c.Param("value")
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid value"})
	}

	index, found := a.searchService.FindIndex(value)

	if index == -1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "index for value not found",
			"value": valueStr,
		})
	}

	response := map[string]interface{}{
		"index": index,
		"value": found,
	}

	if found != value {
		response["message"] = "closest match found"
	}

	return c.JSON(http.StatusOK, response)
}
