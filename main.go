package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bluele/slack"
)

const (
	token         = "YOUR TOKEN"
	channelName   = "general"
	startMessage  = "I'm start the work! :muscle:"
	finishMessage = "I'm finish the work! :beer:"
)

type IotButtonEvent struct {
	DeviceEvent struct {
		ButtonClicked struct {
			ClickType string `json:"clickType"`
		} `json:"buttonClicked"`
	} `json:"deviceEvent"`
}

func handler(event IotButtonEvent) error {
	// Check button click type
	message := startMessage
	if event.DeviceEvent.ButtonClicked.ClickType == "DOUBLE" {
		message = finishMessage
	}

	// Send message to slack
	msgOpts := &slack.ChatPostMessageOpt{AsUser: true}
	api := slack.New(token)
	err := api.ChatPostMessage(channelName, message, msgOpts)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
