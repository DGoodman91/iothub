package main

import (
	"context"
	"log"
	"os"

	"github.com/dgoodman91/iothub/iotdevice"
	iotmqtt "github.com/dgoodman91/iothub/iotdevice/transport/mqtt"
)

func main() {
	c, err := iotdevice.NewFromConnectionString(
		iotmqtt.New(), os.Getenv("IOTHUB_DEVICE_CONNECTION_STRING"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// connect to the iothub
	if err = c.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}

	// send a device-to-cloud message
	if err = c.SendEvent(context.Background(), []byte("hello")); err != nil {
		log.Fatal(err)
	}
}
