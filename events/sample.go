package events

import (
	"context"
	"encoding/json"
	"github.com/KoddiDev/koddi-framework/event"
	"github.com/KoddiDev/koddi-framework/server/log"
)

type EVTEntry struct {
	Message string `json:"message"`
}
type EvtRecord struct{}

var eLog = log.GetStdLogger("leads", "", "myService", "event/myService", "v1", "v1.0")

/*
	Sample to ensure the data contract sent from sample Endpoint is same expected in ReadMessage
*/
func (mg EvtRecord) ReadMessage(ctx context.Context, esbMessage event.ESBMessage) int {

	var msg = EVTEntry{}
	payload, _ := json.Marshal(esbMessage.Payload)
	err := json.Unmarshal(payload, &msg.Message)
	eLog.Info("ReadMessages", esbMessage.CorrelationId, "Payload-"+msg.Message)
	if err != nil {
		//Failed to Unmarshal into the correct Data Contract
		eLog.Error("ReadMessage", esbMessage.CorrelationId, err, "Failed to unmarshal ReadMessage data")
		return event.ContractMismatch
	}
	return event.Success
}
