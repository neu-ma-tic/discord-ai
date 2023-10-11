package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func RunAI(prompt string) string {
	jsonStr := fmt.Sprintf(
		`{"prompt":"{\"history\":[{\"role\":\"system\",\"content\":\"%s\"},{\"role\":\"user\",\"content\":\"%s\"},{\"role\":\"user\",\"content\":\"%s\"}]}","history":[]}`,
		os.Getenv("INJECT_SYSTEM_PROMPT"),
		os.Getenv("INJECT_USER_PROMPT"),
		prompt,
	)
	body := bytes.NewBuffer([]byte(jsonStr))

	req, err := http.NewRequest("POST", "https://cursor.sh/api/chat/stream", body)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	// maybe add more headers if+when they block
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/117.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading body: %s", err)
	}

	result := string(respBody)
	fmt.Println(result)

	return result
}
