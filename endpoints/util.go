package endpoints

import (
	"encoding/json"
	"errors"
	"fmt"
	contracts "go-dark-sky/datacontracts"
	"io/ioutil"
	"net/http"
	"strings"
)

func RespondWith(writer http.ResponseWriter, response contracts.CompletedJobResponse) {
	serialized, err := json.Marshal(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = writer.Write(serialized)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}

func getParameter(request *http.Request, field string, defaultValue string) string {
	var parameter interface{}
	var ok bool

	if request.Method == http.MethodGet {
		parameter, ok = request.URL.Query()[field]
		parameter = fmt.Sprintf("%v", parameter)
	} else {
		parameter = request.Form.Get(field)
		ok = true
	}

	if !ok || parameter == "" {
		return defaultValue
	}

	return fmt.Sprintf("%v", parameter)
}

func validateRequired(request *http.Request, required []string) error {
	var invalidParameters []string
	errorMessage := "Missing required parameters: "

	for i := 0; i < len(required); i++ {

		if getParameter(request, required[i], "") == "" {
			invalidParameters = append(invalidParameters, required[i])
		}
	}

	if len(invalidParameters) > 0 {
		return errors.New(errorMessage + strings.Join(invalidParameters, ", "))
	}

	return nil
}

func callApi(path string) (map[string]interface{}, error) {
	var result map[string]interface{}

	response, err := http.Get(path)
	if err != nil {
		return nil, err
	}

	if response.Body == nil {
		result = nil
		return nil, errors.New("unable to retrieve response")
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
