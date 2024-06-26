package redmine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Authenticator interface {

	// Request create and do a new request to url with method method.
	//
	// The body parameter is a pointer to any struct.
	// If body is not nil, then marshalled to JSON and use it in the new request's body.
	//
	// Returns the status code and the body bytes.
	Request(method string, url string, body any) (int, []byte, error)

	Impersonate(name string)
}

// Public is an instance of a public Redmine server (eg.: www.redmine.org).
type Public struct {
	server string // The server URL
}

// Request implements the Authenticator interface for Public.
func (p *Public) Request(method string, url string, body any) (int, []byte, error) {

	var bodyReader io.Reader = nil

	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return 0, nil, fmt.Errorf("failed to marshal body: %w", err)
		}

		bodyReader = bytes.NewBuffer(bodyBytes)
	}

	req, err := http.NewRequest(method, p.server+url, bodyReader)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to read body: %w", err)
	}

	return resp.StatusCode, respBody, nil
}

func (p *Public) Impersonate(name string) {}

// RegularLogin is authenticates to Redmine with the user's login and password via HTTP Basic Auth.
type RegularLogin struct {
	server   string // Server URL
	login    string // User login name
	password string // User login password
	become   string // The target user's name for user impoersonation with X-Redmine-Switch-User
}

func (rl *RegularLogin) Request(method string, url string, body any) (int, []byte, error) {

	var bodyReader io.Reader = nil

	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return 0, nil, fmt.Errorf("failed to marshal body: %w", err)
		}

		bodyReader = bytes.NewBuffer(bodyBytes)
	}

	req, err := http.NewRequest(method, rl.server+url, bodyReader)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.SetBasicAuth(rl.login, rl.password)

	req.Header.Add("Content-Type", "application/json")

	if rl.become != "" {
		req.Header.Add("X-Redmine-Switch-User", rl.become)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to read body: %w", err)
	}

	return resp.StatusCode, respBody, nil
}

func (rl *RegularLogin) Impersonate(name string) {
	rl.become = name
}

// AuthKey is authenticates to Redmine using the API key passed in as a username with a random password via HTTP Basic Auth
type AuthKey struct {
	server string // Server URL
	key    string // User API key
	become string // The target user's name for user impoersonation with X-Redmine-Switch-User
}

func (ak *AuthKey) Request(method string, url string, body any) (int, []byte, error) {

	var bodyReader io.Reader = nil

	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return 0, nil, fmt.Errorf("failed to marshal body: %w", err)
		}

		bodyReader = bytes.NewBuffer(bodyBytes)
	}

	req, err := http.NewRequest(method, ak.server+url, bodyReader)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.SetBasicAuth(ak.key, strconv.Itoa(int(time.Now().Unix())))

	req.Header.Add("Content-Type", "application/json")

	if ak.become != "" {
		req.Header.Add("X-Redmine-Switch-User", ak.become)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to read body: %w", err)
	}

	return resp.StatusCode, respBody, nil
}

func (ak *AuthKey) Impersonate(name string) {
	ak.become = name
}

// HeaderKey is authenticates to Redmine using the API key passed in "X-Redmine-API-Key" HTTP header.
type HeaderKey struct {
	server string // Server URL
	key    string // User API key
	become string // The target user's name for user impoersonation with X-Redmine-Switch-User
}

func (hk *HeaderKey) Request(method string, url string, body any) (int, []byte, error) {

	var bodyReader io.Reader = nil

	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return 0, nil, fmt.Errorf("failed to marshal body: %w", err)
		}

		bodyReader = bytes.NewBuffer(bodyBytes)
	}

	req, err := http.NewRequest(method, hk.server+url, bodyReader)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.SetBasicAuth(hk.key, strconv.Itoa(int(time.Now().Unix())))

	req.Header.Add("Content-Type", "application/json")

	if hk.become != "" {
		req.Header.Add("X-Redmine-Switch-User", hk.become)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to read body: %w", err)
	}

	return resp.StatusCode, respBody, nil
}

func (hk *HeaderKey) Impersonate(name string) {
	hk.become = name
}
