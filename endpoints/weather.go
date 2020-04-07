package endpoints

import (
	"fmt"
	"github.com/KoddiDev/koddi-framework/endpoint"
	"go-dark-sky/config"
	contracts "go-dark-sky/datacontracts"
	"net/http"
	"strconv"
)

func GetWeather() endpoint.Endpoint {
	return endpoint.NewEndpoint("weather", http.MethodGet, "/weather", HandleGetWeatherRequest)
}

func HandleGetWeatherRequest(writer http.ResponseWriter, request *http.Request) {
	var result map[string]interface{}
	var apiError error
	var errors []string

	validationErrors := validateRequired(request, []string{"latitude", "longitude"})

	if validationErrors != nil {
		errors = append(errors, validationErrors.Error())
	}

	if errors == nil {
		latitude, _ := strconv.ParseFloat(getParameter(request, "latitude", ""), 64)
		longitude, _ := strconv.ParseFloat(getParameter(request, "longitude", ""), 64)

		path := fmt.Sprintf("/forecast/%s/%f,%f", config.ServiceConf.DarkSky.ApiKey, latitude, longitude)

		result, apiError = callApi(config.ServiceConf.DarkSky.Url + path)
		if apiError != nil {
			errors = append(errors, apiError.Error())
		}
	}

	RespondWith(writer, contracts.CompletedJobResponse{
		Message: "Test",
		Data:    result,
		Errors:  errors,
	})
}
