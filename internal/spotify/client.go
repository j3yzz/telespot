package spotify

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/j3yzz/telespot/internal/config"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

func NewClient(config *config.Config) (*Client, error) {
	authToken := "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", config.SpotifyClientId, config.SpotifyClientSecret)))
	authApi := config.SpotifyAuthApi
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	proxyUrl, err := url.Parse("http://127.0.0.1:2081")
	client := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
	}

	req, err := http.NewRequest("POST", authApi, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalln("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Authorization", authToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalln("Error in connecting to Spotify. Status Code:", resp.StatusCode)
		return nil, errors.New(fmt.Sprintf("Error in connecting to Spotify. Status Code:%s", resp.StatusCode))
	}

	var clientResponse Client
	err = json.NewDecoder(resp.Body).Decode(&clientResponse)
	if err != nil {
		log.Fatalln("Error decoding response:", err)
		return nil, err
	}

	return &clientResponse, nil
}
