package main

import (
	"net/http"
)

func main() {

	//client := &http.Client{}
	req, _ := http.NewRequest("GET", "endpoint-url", nil)
	req.Header.Set("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiJkNTFmY2Q5MC00YjBiLTAxMzYtNDRjNy0wZTc3YzRhYzg5MmEiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNTI4MjE2MjkzLCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6InB1Ymdfc3RhdF90cmFja2VyLTRhNzk1M2NkLTIzZDktNDJmNy05ZWNjLWJjNTY1YzQwZDcxNyJ9.i5b0XrU11RA8oq4bXRdr4NekwickWaqMt1GHJc71m40")
	req.Header.Set("Accept", "application/vnd.api+json")
	//res, _ := client.Do(req)

}
