package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/KoddiDev/koddi-framework/endpoint"
	"go-dark-sky/config"
	contracts "go-dark-sky/datacontracts"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GetWeather() endpoint.Endpoint {
	return endpoint.NewEndpoint("weather", http.MethodGet, "/weather", HandleGetWeatherRequest)
}

func HandleGetWeatherRequest(writer http.ResponseWriter, req *http.Request) {
	var errors []string

	latitudeQuery, ok := req.URL.Query()["latitude"]
	if !ok {
		errors = append(errors, "Missing parameter latitude")
	}

	longitudeQuery, ok := req.URL.Query()["longitude"]
	if !ok {
		errors = append(errors, "Missing parameter longitude")
	}

	latitude, _ := strconv.ParseFloat(strings.Join(latitudeQuery, ""), 64)
	longitude, _ := strconv.ParseFloat(strings.Join(longitudeQuery, ""), 64)

	path := fmt.Sprintf("/forecast/%s/%f,%f", config.ServiceConf.DarkSky.ApiKey, latitude, longitude)

	response, err := http.Get(config.ServiceConf.DarkSky.Url + path)
	if err != nil {
		errors = append(errors, err.Error())
	}

	var result map[string]interface{}

	if response.Body == nil {
		result = nil
		errors = append(errors, "Unable to retrieve response")
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			errors = append(errors, err.Error())
		}
		err = json.Unmarshal(body, &result)
		if err != nil {
			errors = append(errors, err.Error())
		}
	}

	RespondWith(writer, contracts.CompletedJobResponse{
		Message: "Test",
		Data:    result,
		Errors:  errors,
	})
}
