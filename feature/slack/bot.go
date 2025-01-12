package slack

import (
	"fmt"
	"github.com/slack-go/slack/slackevents"
	"log"
	"os"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func Bot() {

	token := os.Getenv("SLACK_TOKEN")
	appToken := os.Getenv("APP_TOKEN")

	print(token)
	print(appToken)
	client := slack.New(
		token,
		slack.OptionAppLevelToken(appToken),
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "remote", log.Lshortfile|log.LstdFlags)),
	)

	socketMode := socketmode.New(
		client,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)

	authTest, authTestError := client.AuthTest()
	if authTestError != nil {
		fmt.Print(authTestError)
		os.Exit(1)
	}
	selfUserId := authTest.UserID
	go func() {
		for envelope := range socketMode.Events {
			switch envelope.Type {
			case socketmode.EventTypeEventsAPI:
				// Events API:

				// Acknowledge the eventPayload first
				socketMode.Ack(*envelope.Request)

				eventPayload, _ := envelope.Data.(slackevents.EventsAPIEvent)
				switch eventPayload.Type {
				case slackevents.CallbackEvent:
					switch event := eventPayload.InnerEvent.Data.(type) {
					case *slackevents.MessageEvent:
						if event.User != selfUserId &&
							strings.Contains(strings.ToLower(event.Text), "hello") {
							_, _, err := client.PostMessage(
								event.Channel,
								slack.MsgOptionText(
									fmt.Sprintf(":wave: Hi there, <@%v>!", event.User),
									false,
								),
							)
							if err != nil {
								log.Printf("Failed to reply: %v", err)
							}
						}
					default:
						socketMode.Debugf("Skipped: %v", event)
					}
				default:
					socketMode.Debugf("unsupported Events API eventPayload received")
				}
			case socketmode.EventTypeInteractive:
				// Shortcuts:

				payload, _ := envelope.Data.(slack.InteractionCallback)
				switch payload.Type {
				case slack.InteractionTypeShortcut:
					socketMode.Debugf("Interaction callback received")
				case slack.InteractionTypeViewSubmission:
					// View Submission:
					if payload.CallbackID == "modal-id" {
						socketMode.Debugf("Submitted Data: %v", payload.View.State.Values)
						socketMode.Ack(*envelope.Request)
					}
				default:
					// Others
					socketMode.Debugf("Skipped: %v", payload)
				}

			default:
				socketMode.Debugf("Skipped: %v", envelope.Type)
			}
		}
	}()

	err := socketMode.Run()
	if err != nil {
		return
	}
}
