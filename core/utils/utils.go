package utils

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	models "github.com/hizbi-github/gost/new-project-core/models"
)

// func HttpPost(url string, body []byte, headers map[string][]string) (*models.HttpResponse, error) {
func HttpPost(httpRequest *models.HttpRequest) (*models.HttpResponse, error) {
	if httpRequest == nil {
		return nil, errors.New("request struct is nil")
	}

	req, err := http.NewRequest("POST", httpRequest.Url, bytes.NewReader(httpRequest.Body))
	if err != nil {
		return nil, err
	}

	if httpRequest.Headers != nil {
		req.Header = httpRequest.Headers
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")

	res, err := httpTlsClient().Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &models.HttpResponse{
		Headers: res.Header.Clone(),
		Body:    responseBody,
	}, nil
}

// func HttpGet(url string, headers map[string][]string) (*models.HttpResponse, error) {
func HttpGet(httpRequest *models.HttpRequest) (*models.HttpResponse, error) {
	if httpRequest == nil {
		return nil, errors.New("request struct is nil")
	}

	req, err := http.NewRequest("GET", httpRequest.Url, nil)
	if httpRequest.Body != nil {
		req, err = http.NewRequest("GET", httpRequest.Url, bytes.NewReader(httpRequest.Body))
	}
	if err != nil {
		return nil, err
	}

	if httpRequest.Headers != nil {
		req.Header = httpRequest.Headers
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")

	res, err := httpTlsClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &models.HttpResponse{
		Headers: res.Header.Clone(),
		Body:    responseBody,
	}, nil
}

func httpTlsClient() *http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS13,
		},
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   60 * time.Second,
	}

	return client
}

func Trim(untrimmed string) string {
	return (strings.TrimSpace(untrimmed))
}

func appendToFile(fileName string, fileContent string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	//f, err := os.Create(fileName)

	if err != nil {
		log.Fatalf("Unable to open file for appending: %v", err)
	}
	defer f.Close()

	_, err = f.WriteString(fileContent)
	if err != nil {
		log.Fatalf("Unable to write to the local backup file: %v", err)
	}
}

func whitespaceSplitAndJoin(unsplit string) string {
	//fmt.Println(strings.Fields(trim(unsplit)))
	return (strings.Join(strings.Fields(Trim(unsplit)), " "))
}
