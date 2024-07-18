package irstats

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"fmt"
	"os"

	resty "github.com/go-resty/resty/v2"
)

const (
	defaultUserAgent = "go-irstats"
	defaultBaseURL   = "https://members-ng.iracing.com"
)

var (
	ErrAuthenticationFailed = errors.New("failed to authenticate with iRacing")
)

type Client struct {
	http *resty.Client // http client

	username      string
	password      string
	userAgent     string
	authenticated bool
}

func NewClient(username, password string, options ...ClientOptionFunc) (*Client, error) {
	http := resty.New()
	http.SetBaseURL(defaultBaseURL)
	http.SetRedirectPolicy(resty.NoRedirectPolicy())

	c := &Client{
		http:      http,
		username:  username,
		password:  password,
		userAgent: defaultUserAgent,
	}

	// Apply any given client options.
	for _, fn := range options {
		if fn == nil {
			continue
		}
		if err := fn(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) do(path urlPath, values *url.Values, v interface{}) (*http.Response, error) {
	if err := c.assertLoggedIn(); err != nil {
		return nil, err
	}

	var resp *resty.Response
	var err error

	r := c.http.R().SetHeader("User-Agent", c.username)
	if values == nil {
		resp, err = r.Get(path)
	} else {
		resp, err = r.SetFormDataFromValues(*values).Post(path)
	}
	rr := resp.RawResponse

	if err != nil {
		log.Printf("API request %s failed. %v\n", path, err)
		return rr, err
	}

	// If there is no target to unmarshal the json to, then we return out here
	if v == nil {
		return rr, nil
	}

	err = json.Unmarshal(resp.Body(), v)
	if err != nil {
		log.Println("Failed to unmarshal response body", err)
		return rr, err
	}

	return rr, nil
}

func (c *Client) Login() error {
	// Convert email to lowercase
	emailLower := strings.ToLower(c.username)

	// Encode the password with the lowercase email
	hash := sha256.New()
	hash.Write([]byte(c.password + emailLower))
	encodedPW := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	// Create the request body
	body := map[string]string{
		"email":    c.username,
		"password": encodedPW,
	}

	// Send the POST request
	resp, err := c.http.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("User-Agent", c.userAgent).
		SetBody(body).
		Post("/auth")

	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}

	// Save the cookies to a file
	cookieFile := "cookie-jar.txt"
	f, err := os.Create(cookieFile)
	if err != nil {
		return fmt.Errorf("error creating cookie file: %w", err)
	}
	defer f.Close()

	for _, cookie := range resp.Cookies() {
		_, err := f.WriteString(fmt.Sprintf("%s=%s\n", cookie.Name, cookie.Value))
		if err != nil {
			return fmt.Errorf("error writing to cookie file: %w", err)
		}
	}

	// Print the response body
	fmt.Println(string(resp.Body()))
	fmt.Println("Response Headers:")
	for key, values := range resp.Header() {
		fmt.Printf("%s: %s\n", key, values)
	}

	return nil
}

func (c *Client) assertLoggedIn() error {
	if c.authenticated {
		return nil
	}

	return c.Login()
}

func (c *Client) CheckLoggedIn() error {
	if c.authenticated {
		return nil
	}

	return c.Login()
}
