package endpoints

import (
	"encoding/json"
	contracts "go-dark-sky/datacontracts"
	"net/http"
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
