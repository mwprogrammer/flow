# flow
A library for building interactive apps to engage with customers over the whatsapp business platform. Built ontop of the WhatsApp Cloud API provided by Meta technologies. [Learn more about it here](https://developers.facebook.com/documentation/business-messaging/whatsapp/overview).

## Project status
🛠️ Under construction

## Installation
```
go get github.com/mwprogrammer/flow
```

## Documentation

### Setup

Create a Meta developer account and register a new app with Meta. Follow the directions [here](https://developers.facebook.com/docs/development/create-an-app/) and [here](https://developers.facebook.com/docs/whatsapp/cloud-api/get-started).

#### Configure Flow Settings

Define your flow settings by specifying your Whatsapp business Account Id, The API version (24 is the default), your Access Token (Temporary or Permanent) and the Sender Phone Number ID. 
```
flow_settings := flow.FlowSettings{

    Id:      "XXXXXX", // Whatsapp Business Account ID
    Version: "24.0", // Whatsapp Cloud API version
    Token:   "XXXXXX", // Access Token
    Sender:  "XXXXXXX", // Sender Phone Number Id

}
```

#### Create your Flow app

Initialize a new Flow Object with the settings you created. You now have access to Flow methods which enable you to interact with users over the Whatsapp Business Platform.
```
flow_app := flow.New(new_flow_settings)
```

### Flow Methods

#### Reply with a Text Message

Send text messages to users by specifying the **receipient phone number**, the **message** and if the message includes a link, **whether or not to display the preview**.
```
err := flow_app.ReplyWithText("26588293345", "Hello", false)
```



