package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/KoddiDev/koddi-framework/endpoint"
	"github.com/KoddiDev/koddi-framework/event"
	"github.com/KoddiDev/koddi-framework/server/log"
	"github.com/KoddiDev/koddi-framework/start"
	"math/rand"
	"net/http"
	"strconv"
)

var (
	eLog = log.GetStdLogger("test", "", "messageTest", "new", "v1", "v1.0")
)

type message struct {
	Message string `json:"message"`
}

// create all the endpoints handler...
//this can be done in a separate file
func GetMyEndpoint() endpoint.Endpoint {

	return endpoint.NewEndpoint("myEndpoint",
		http.MethodGet,
		"/new",
		EndpointHandleRequest)
}

/*
	Sample Endpoint handler
*/
func EndpointHandleRequest(writer http.ResponseWriter, req *http.Request) {
	//example to push message to Kafka stream.
	esbServicePtr := start.EsbInstance

	valid := `{
		"EventType": "records",
		"SubType": "dms.new",
		"CorrelationId": "12345",
		"Version": "1.0",
		"Payload": "just a string"
	}`

	var esbMessage event.ESBMessage
	json.Unmarshal([]byte(valid), &esbMessage)

	//set random payload
	esbMessage.Payload = strconv.Itoa(rand.Intn(500))
	payloadJSON, _ := json.Marshal(esbMessage)
	if esbServicePtr != nil {
		if err := esbServicePtr.SendMessage("records.dms.new", []byte("key"), payloadJSON); err != nil {
			fmt.Println("Error sending message: ", err)
		}
	} else {
		fmt.Println("esbServicePtr NOT Found")
	}
	//response to rest service
	writer.Write([]byte("sent")) //response message
}
