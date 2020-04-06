package main

import (
	"github.com/KoddiDev/koddi-framework/endpoint"
	"github.com/KoddiDev/koddi-framework/event"
	"github.com/KoddiDev/koddi-framework/start"
	"go-dark-sky/config"
	"go-dark-sky/endpoints"
	"go-dark-sky/events"
)

func main() {
	//Load Config for everyone to use
	var configDir = "../go-dark-sky/config"
	_, err := config.LoadConfig()
	if err != nil {
		//Couldn't load Config File
		panic("Configs were unable to load....")
	}

	// create the events handlers and listeners....
	dmsEvent := event.NewEvent("records", "dms.new", "test", events.EvtRecord{}.ReadMessage)

	start.Start(start.StartData{
		Endpoints: []endpoint.Endpoint{
			endpoints.GetWeather(),
		},

		Events: []event.Event{
			*dmsEvent,
		},
	}, configDir)
}
