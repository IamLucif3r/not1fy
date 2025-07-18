package main

import (
	"fmt"

	"github.com/iamlucif3r/not1fy/not1fy"
)

func main() {
	webhookURL := "https://discord.com/api/webhooks/<YOUR_WEBHOOK_ID>"
	not1fy.SetWebhookURL(webhookURL)

	err := not1fy.SendAlert("This is a test alert from the not1fy package!")
	if err != nil {
		fmt.Printf("Error sending alert: %v\n", err)
		return
	}
	fmt.Println("Simple alert sent successfully!")

	jsonPayload := map[string]interface{}{
		"content": "This is a custom JSON message!",
	}
	err = not1fy.SendJSONPayload(jsonPayload)
	if err != nil {
		fmt.Printf("Error sending JSON payload: %v\n", err)
		return
	}
	fmt.Println("Custom JSON payload sent successfully!")

	customMsg := not1fy.DiscordMessage{
		Content: "This is a custom Discord message with an embed!",
		Embeds: []not1fy.Embed{
			{
				Title:       "Test Embed",
				Description: "Here's some embedded content.",
				Color:       0x00ff00,
				Fields: []not1fy.Field{
					{
						Name:   "Field 1",
						Value:  "Value 1",
						Inline: true,
					},
					{
						Name:   "Field 2",
						Value:  "Value 2",
						Inline: true,
					},
				},
			},
		},
		AllowedMentions: &not1fy.AllowedMentions{
			Parse: []string{"users"},
		},
	}
	err = not1fy.SendCustomMessage(customMsg)
	if err != nil {
		fmt.Printf("Error sending custom message: %v\n", err)
		return
	}
	fmt.Println("Custom Discord message sent successfully!")
}
