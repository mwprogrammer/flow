# 🌊 Flow

[![Pre-Release](https://img.shields.io/github/v/release/mwprogrammer/flow?include_prereleases&label=Beta&color=orange)](https://github.com/mwprogrammer/flow/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/mwprogrammer/flow/blob/develop/CONTRIBUTING.md)

---

A library for building interactive apps to engage with customers over the whatsapp business platform. Built ontop of the WhatsApp Cloud API provided by Meta technologies. [Learn more about it here](https://developers.facebook.com/documentation/business-messaging/whatsapp/overview).

## 🗺️ Project Roadmap

We are currently **🛠️ Under construction**. Our focus is on stabilizing core features, ensuring code quality through automated auditing, and gathering community feedback.

### 📍 Current Phase: Foundation & CI/CD
- [x] Complete project structure plus basic methods (read message, mark message as read, add text, etc)
- [x] Standardize contribution process and implement basic CI/CD.
- [x] Automated Go Linting (Static Analysis)
- [x] Publish beta on go packages.
- [ ] Add other methods for first release (In Progress)
- [ ] Add demo project utilizing library (In Progress)
- [ ] Publish tutorial.
- [ ] Accept contributions.

### 🔜 Upcoming Milestones
- **Publish first release v1.0.0: Release library for community feedback**
  
---
*Note: This roadmap is a living document and will be updated based on community feedback and project evolution.*

## Installation (Pre-release)
```bash
go get github.com/mwprogrammer/flow@v1.0.0-beta-0.1.1
```

## Documentation

### Setup

Create a Meta developer account and register a new app with Meta. Follow the directions [here](https://developers.facebook.com/docs/development/create-an-app/) and [here](https://developers.facebook.com/docs/whatsapp/cloud-api/get-started).

#### Configure Flow Settings

Define your flow settings by specifying your Whatsapp business Account Id, The API version (24 is the default), your Access Token (Temporary or Permanent) and the Sender Phone Number ID. 
```go
flow_settings := flow.FlowSettings{

    Id:      "XXXXXX", // Whatsapp Business Account ID
    Version: "24.0", // Whatsapp Cloud API version
    Token:   "XXXXXX", // Access Token
    Sender:  "XXXXXXX", // Sender Phone Number Id

}
```

#### Create your Flow app

Initialize a new Flow Object with the settings you created. You now have access to Flow methods which enable you to interact with users over the Whatsapp Business Platform.
```go
flow_app := flow.New(flow_settings)
```

### Flow Methods

#### Parse incoming messages

After setting up your webhook, read incoming messages from receipients.
```go
sample_json := `{"object":"whatsapp_business_account","entry":[{"id":"0","changes":[{"field":"messages","value":{"messaging_product":"whatsapp","metadata":{"display_phone_number":"16505551111","phone_number_id":"123456123"},"contacts":[{"profile":{"name":"test user name"},"wa_id":"16315551181"}],"messages":[{"from":"16315551181","id":"ABGGFlA5Fpa","timestamp":"1504902988","type":"text","text":{"body":"this is a text message"}}]}}]}]}`

message, err := new_flow.ParseMessage(sample_json)
```

#### Mark message as read

Mark a message as ready by specifying the **receipient phone** and **messageId**.
```go
err := flow_app.MarkAsRead("26588293345", "XXXXXXXXXX")
```

#### Display Typing Indicator

Display that the app is typing by specifying the **receipient phone** and **messageId**.
```go
err := flow_app.DisplayTypingIndicator("26588293345", "XXXXXXXXXX")
```

#### Reply with a Text Message

Send text messages to users by specifying the **receipient phone number**, the **message** and if the message includes a link, **whether or not to display the preview**.
```go
err := flow_app.ReplyWithText("26588293345", "Hello", false)
```