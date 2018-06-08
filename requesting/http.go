package requesting

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
)

func new(url, key string) (*bytes.Buffer, error) {
	// Create the request
	req, err := http.NewRequest("GET", "endpoint-url", nil)

	if err != nil {
		return nil, err
	}

	// Configure the  request
	req.Header.Set("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiJkNTFmY2Q5MC00YjBiLTAxMzYtNDRjNy0wZTc3YzRhYzg5MmEiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNTI4MjE2MjkzLCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6InB1Ymdfc3RhdF90cmFja2VyLTRhNzk1M2NkLTIzZDktNDJmNy05ZWNjLWJjNTY1YzQwZDcxNyJ9.i5b0XrU11RA8oq4bXRdr4NekwickWaqMt1GHJc71m40")
	req.Header.Set("Accept", "application/vnd.api+json")

	// Send the request
	client := &http.Client{}
	response, err := client.Do(req)

	if response.StatusCode != 200 {
		response.Body.Close()
		return nil, fmt.Errorf("HTTP request failed: %s", response.Status)
	}

	var reader io.ReadCloser

	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			return nil, err
		}
	default:
		reader = response.Body
	}

	var buffer bytes.Buffer
	buffer.ReadFrom(reader)

	return &buffer, nil
}
