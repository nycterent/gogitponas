package main

import (
	"gogitponas/gitlab"
	"gogitponas/rocketchat"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var RocketChatClient *rocketchat.RocketChat

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gitlab_url := os.Getenv("GITLAB_URL")
	gitlab_token := os.Getenv("GITLAB_TOKEN")
	gitlab_projects := strings.Split(os.Getenv("GITLAB_PROJECTS"), ",")

	// slack_hook := os.Getenv("SLACK_HOOK")
	// slack_channel := os.Getenv("SLACK_CHANNEL")

	g := gitlab.New(gitlab_token, gitlab_url)
	RocketChatClient = rocketchat.New(os.Getenv("ROCKET_HOOK"))

	for _, project := range gitlab_projects {
		SendNotifications(g.GetOldMergeRequests(project))
	}

}

func SendNotifications(MRRequests []gitlab.GitlabMergeInformation) {
	notification_targets := strings.Split(os.Getenv("NOTIFICATION_TARGETS"), ",")
	for _, mr := range MRRequests {
		for _, target := range notification_targets {
			switch target {
			case "rocket":
				RocketChatClient.Send(rocketchat.RocketChatMessage{
					Text: mr.Title,
					Attachments: []rocketchat.RocketChatMessageAttachment{
						rocketchat.RocketChatMessageAttachment{
							Title:     mr.Reference,
							TitleLink: mr.MRUrl,
							Text:      mr.Author,
						},
					},
				})
				break
			}
		}
	}
}
