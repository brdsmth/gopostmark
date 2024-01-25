package goposhmark

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

type PoshmarkClient struct {
	ServerToken string
}

type EmailRequest struct {
	From          string `json:"From"`
	To            string `json:"To"`
	Subject       string `json:"Subject"`
	TextBody      string `json:"TextBody"`
	HtmlBody      string `json:"HtmlBody"`
	MessageStream string `json:"MessageStream"`
}

// Create Poshmark Email API client.
func Client(serverToken string) *PoshmarkClient {
	return &PoshmarkClient{
		ServerToken: serverToken,
	}
}

// Sends via the Poshmark API
func (pc *PoshmarkClient) SendEmail(emailRequest EmailRequest) ([]byte, error) {
	requestBody, err := json.Marshal(emailRequest)
	if err != nil {
		return nil, err
	}

	apiURL := "https://api.postmarkapp.com/email"

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Postmark-Server-Token", pc.ServerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http failure: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Format and print the response
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, fmt.Errorf("failed to format response: %s", err)
		}
		dumpStr := string(dump)
		return nil, errors.New(dumpStr)
	}

	// Read and return the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
