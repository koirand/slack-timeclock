package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type IotButtonEvent struct {
	DeviceInfo struct {
		DeviceID      string  `json:"deviceId"`
		Type          string  `json:"type"`
		RemainingLife float64 `json:"remainingLife"`
		Attributes    struct {
			ProjectRegion      string `json:"projectRegion"`
			ProjectName        string `json:"projectName"`
			PlacementName      string `json:"placementName"`
			DeviceTemplateName string `json:"deviceTemplateName"`
		} `json:"attributes"`
	} `json:"deviceInfo"`
	DeviceEvent struct {
		ButtonClicked struct {
			ClickType    string    `json:"clickType"`
			ReportedTime time.Time `json:"reportedTime"`
		} `json:"buttonClicked"`
	} `json:"deviceEvent"`
	PlacementInfo struct {
		ProjectName   string `json:"projectName"`
		PlacementName string `json:"placementName"`
		Attributes    struct {
			Name string `json:"name"`
		} `json:"attributes"`
		Devices struct {
			SampleRequest string `json:"Sample-Request"`
		} `json:"devices"`
	} `json:"placementInfo"`
}

type EnvError struct{}

func (e *EnvError) Error() string {
	return "Env variable of TIMECLOCK_WEBHOOK_URL is not set."
}

const startMessage string = `{
	"text": ":muscle: %s started work!"
}
`

const finishMessage string = `{
	"text": ":beer: %s finished work!"
}
`

func handler(event IotButtonEvent) error {
	webhookUrl := os.Getenv("TIMECLOCK_WEBHOOK_URL")
	if webhookUrl == "" {
		return &EnvError{}
	}

	// check button click type
	message := startMessage
	if event.DeviceEvent.ButtonClicked.ClickType == "DOUBLE" {
		message = finishMessage
	}

	// send to slack
	_, err := http.Post(
		webhookUrl,
		"application/json",
		strings.NewReader(fmt.Sprintf(message, event.PlacementInfo.Attributes.Name)),
	)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
