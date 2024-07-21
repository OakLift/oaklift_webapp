package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"ai-tooling.com/internal/openai/project-creator"
	"github.com/pkg/errors"
)

const (
	host = "localhost"
	port = "8001"
	chatEndpoint = "/api/v1/code/review"
	projectCreationEndpoint = "/api/v1/project/create"
)

func main() {
	projectResponse, err := testProjectCreator()
	if err != nil {
		log.Fatalf("failed to hit project creator API: %v", err)
	}

	log.Printf("Poject Response:\n%s\n", projectResponse)
}

func setupClient() *http.Client {
	httpTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: httpTransport}
}


func testProjectCreator() (string, error ) {
	httpClient := setupClient()

	url := url.URL{
		Scheme: "http",
		Host: host + ":" + port,
		Path: projectCreationEndpoint,
	}

	createRequest := &project.CreateRequest{Language: project.HTMLCSS, Level: project.Beginner}
	body, err := json.Marshal(createRequest)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal create request")
	}

	resp, err := httpClient.Post(url.String(), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", errors.Wrapf(err, "POST request to %q failed", url.String())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed to read response body")
	}
	
	return string(respBody), nil
}

func testChatPost() (string, error) {
	httpClient := setupClient()

	url := url.URL{
		Scheme: "http",
		Host: host + ":" + port,
		Path: chatEndpoint,
	}

	bodyBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return "", errors.Wrap(err, "failed to read input file")
	}

	body := bytes.NewBuffer(bodyBytes)
	resp, err := httpClient.Post(url.String(), "application/json", body)
	if err != nil {
		return "", errors.Wrapf(err, "POST request to %q failed", url.String())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed to read response body")
	}
	
	return string(respBody), nil
}
