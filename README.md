## Introduction

Go wrapper for more easily interacting with the Postmark API.

## Usage

Import the package

```
go get -u github.com/brdsmth/gopostmark
```

Create a client and pass `X-Postmark-Server-Token`

```
client := gopostmark.Client("your-server-token")

email := gopostmark.EmailRequest{
	From:          "example@example.com",
	To:            "example@example.com",
	Subject:       "Postmark test",
	TextBody:      "Hello dear Postmark user.",
	HtmlBody:      "<html><body><strong>Hello</strong> dear Postmark user.</body></html>",
	MessageStream: "outbound",
}

response, err := client.SendEmail(email)
```

## Contribute

Contributions are welcome. Please open an issue or a pull request for any changes.
