package pushover

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const apiurl = "https://api.pushover.net/1/messages.json"
const MaxMsgLength = 1024
const MaxTitleLength = 250

const (
	PriorityLowest = -2 + iota
	PriorityLow
	PriorityNormal
	PriorityHigh
)

// Message represents a message in the Pushover Message API.
type Message struct {
	User string
	Token string
	Title string
	Message string
	Priority int
}

type response struct {
	Status int
	Request string
	Errors errors
}

type errors []string

func (e errors) Error() string {
	return strings.Join(e, ", ")
}

func (m *Message) validate() error {
	nchar := strings.Count(m.Message, "")
	if nchar > MaxMsgLength {
		return fmt.Errorf("%d character message too long, allowed %d characters", nchar, MaxMsgLength)
	}
	nchar = strings.Count(m.Title, "")
	if nchar > MaxTitleLength {
		return fmt.Errorf("%d-character title too long, allowed %d characters", nchar, MaxTitleLength)
	}
	return nil
}

// Push sends the Message m to Pushover.
func Push(m Message) error {
	if err := m.validate(); err != nil {
		return err
	}
	req := url.Values{}
	req.Add("token", m.Token)
	req.Add("user", m.User)
	req.Add("title", m.Title)
	req.Add("message", m.Message)
	if m.Priority != 0 {
		req.Add("priority", strconv.Itoa(m.Priority))
	}
	resp, err := http.PostForm(apiurl, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	var presp response
	if err := json.NewDecoder(resp.Body).Decode(&presp); err != nil {
		return fmt.Errorf("decode error response: %v", err)
	}
	return presp.Errors
}
