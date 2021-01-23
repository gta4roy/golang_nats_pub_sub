package main

import (
	"log"
	"runtime"

	pb "nats_event/order"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"
)

func main() {

	// Create server connection
	natsConnection, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Println("Not able to Connect NATS Server " + nats.DefaultURL)
	} else {
		log.Println("Able to Connect NATS Server " + nats.DefaultURL)
	}

	natsConnection.Subscribe("Discovery.GetSystemInfo", func(msgSecondRequest *nats.Msg) {
		var systemInfo pb.GetSystemTime
		err := proto.Unmarshal(msgSecondRequest.Data, &systemInfo)
		if err != nil {
			log.Fatalf("Error on unmarshal: %v", err)
		}

		log.Println(" System Info Details :")
		log.Println("System Time ", systemInfo.Systemtime)
		log.Println("System Date ", systemInfo.Systemdate)
		log.Println("System Username ", systemInfo.Username)
		log.Println("System Server Ip ", systemInfo.Serverip)

	})

	// Keep the connection alive
	runtime.Goexit()
}
