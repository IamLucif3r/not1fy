# not1fy

A simple Go package that allows you to send messages to a Discord channel using webhooks. It supports sending plain text messages as well as more complex customized messages with embeds.

## Table of Contents

1. [Installation](#1-installation)
2. [Usage](#2-usage)
    2.1 Setting the Webhook URL
    2.2 Sending Simple Text Messages
    2.3 Sending Custom JSON Payloads
    2.4 Sending Custom Discord Messages
3. Examples
4. License

## 1. Installation
To install the not1fy package, you can use Go's module system:

```sh
go get github.com/iamlucif3r/not1fy
```

## 2. Usage

### 2.1 Setting the Webhook URL:
Before sending any messages, you need to set the Discord webhook URL. This can be done using the SetWebhookURL function:

```go

import "github.com/yourusername/discord_notifier"

webhookURL := "https://discord.com/api/webhooks/YOUR_WEBHOOK_ID/YOUR_WEBHOOK_TOKEN"
discord_notifier.SetWebhookURL(webhookURL)
```

### 2.2 Sending Simple Text Messages:
You can send simple text messages using the SendAlert function:

```go
err := not1fy.SendAlert("This is a test alert from not1fy!")
if err != nil {
    fmt.Printf("Error sending alert: %v\n", err)
}
```

### 2.3 Sending Custom JSON Payloads
For more complex use cases, you can send custom JSON payloads using the SendJSONPayload function:

```go
jsonPayload := map[string]interface{}{
    "content": "This is a custom JSON message!",
}
err := not1fy.SendJSONPayload(jsonPayload)
if err != nil {
    fmt.Printf("Error sending JSON payload: %v\n", err)
}
```

### 2.4 Sending Custom Discord Messages
You can also send customized Discord messages with embeds and other options using the SendCustomMessage function:

```go
customMsg := not1fy.DiscordMessage{
    Content: "This is a custom Discord message with an embed!",
    Embeds: []not1fy.Embed{
        {
            Title:       "Test Embed",
            Description: "Here's some embedded content.",
            Color:       0x00ff00, // Green color
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
}
err := not1fy.SendCustomMessage(customMsg)
if err != nil {
    fmt.Printf("Error sending custom message: %v\n", err)
}
```


## License
This package is licensed under the MIT License. See the LICENSE file for details.

## Contributions
Feel free to customize this template according to your specific needs and preferences!