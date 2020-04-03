package main

import (
	"github.com/KoddiDev/koddi-framework-starter/config"
	"github.com/KoddiDev/koddi-framework-starter/endpoints"
	"github.com/KoddiDev/koddi-framework-starter/events"
	"github.com/KoddiDev/koddi-framework/endpoint"
	"github.com/KoddiDev/koddi-framework/event"
	"github.com/KoddiDev/koddi-framework/start"
)

func main() {
	//Load Config for everyone to use
	var configDir = "../koddi-framework-starter/config"
	_, err := config.LoadConfig()
	if err != nil {
		//Couldn't load Config File
		panic("Configs were unable to load....")
	}

	// create the events handlers and listeners....
	dmsEvent := event.NewEvent("records", "dms.new", "test", events.EvtRecord{}.ReadMessage)

	start.Start(start.StartData{
		Endpoints: []endpoint.Endpoint{
			endpoints.GetMyEndpoint(),
		},

		Events: []event.Event{
			*dmsEvent,
		},
	}, configDir)
}
