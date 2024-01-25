## Introduction

Go wrapper for more easily interacting with the Poshmark API.

## Usage

```
go get -u github.com/brdsmth/goposhmark
```

```
package main

import (
	"fmt"

	goposhmark "github.com/brdsmth/goposhmark/client"
)

func main() {

	client := goposhmark.Client("your-server-token")

	email := goposhmark.EmailRequest{
		From:          "example@example.com",
		To:            "example@example.com",
		Subject:       "Postmark test",
		TextBody:      "Hello dear Postmark user.",
		HtmlBody:      "<html><body><strong>Hello</strong> dear Postmark user.</body></html>",
		MessageStream: "outbound",
	}

	response, err := client.SendEmail(email)
	if err != nil {
		// Handle the error
	}
}

```

## Contribute

Contributions are welcome. Please open an issue or a pull request for any changes.
