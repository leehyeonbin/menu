package slack

import (
	"log"
	"os"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func Bot() {

	token := os.Getenv("SLACK_TOKEN")
	appToken := os.Getenv("APP_TOKEN")

	client := slack.New(token, slack.OptionDebug(true), slack.OptionAppLevelToken(appToken))

	app := socketmode.New(
		client,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)

	err := app.Run()
	if err != nil {
		log.Fatalf("Error starting socketmode listener: %v", err)
		return
	}
}
