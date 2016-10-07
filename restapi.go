package cardcastgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Error for when unmarshaling goes wrong
var ErrJSONUnmarshal = errors.New("json unmarshal")

// All requests will use this method
func (s *Session) Request(method, urlStr string, data interface{}) (response []byte, err error) {

	var body []byte
	if data != nil {
		body, err = json.Marshal(data)
		if err != nil {
			return
		}
	}
	return s.request(method, urlStr, "application/json", body)
}

// What goes behind the scenes
func (s *Session) request(method, urlStr, contentType string, b []byte) (response []byte, err error) {

	req, err := http.NewRequest(method, urlStr, bytes.NewBuffer(b))
	if err != nil {
		return
	}

	if s.Token != "" {
		req.Header.Set("X-Auth-Token", s.Token)
	}

	req.Header.Set("Content-Type", contentType)

	client := &http.Client{Timeout: (20 * time.Second)}

	resp, err := client.Do(req)
	defer func() {
		resp.Body.Close()
	}()

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	switch resp.StatusCode {

	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusNoContent:

	default: // Error condition
		err = fmt.Errorf("HTTP %s, %s", resp.Status, response)
	}

	return
}

func unmarshal(data []byte, v interface{}) error {

	err := json.Unmarshal(data, v)
	if err != nil {
		return ErrJSONUnmarshal
	}

	return nil
}

// ------------------------------------------------------------------------------------------------
// Functions specific to Decks
// ------------------------------------------------------------------------------------------------

// Returns a Carddeck type object
func (s *Session) GetDeck(deckID string) (cd *Carddeck, err error) {

	body, err := s.Request("GET", EndpointDeck(deckID), nil)
	if err != nil {
		return
	}

	err = unmarshal(body, &cd)
	return

}

// Returns a Card type object containing all calls (black cards)
func (s *Session) GetCalls(deckID string) (c *[]Card, err error) {

	body, err := s.Request("GET", EndpointCalls(deckID), nil)
	if err != nil {
		return
	}

	err = unmarshal(body, &c)
	return

}

// Returns a Card type object containing all responses (white cards)
func (s *Session) GetResponses(deckID string) (c *[]Card, err error) {

	body, err := s.Request("GET", EndpointResponses(deckID), nil)
	if err != nil {
		return
	}

	err = unmarshal(body, &c)
	return

}

// Posts a call to the deck (You must own the deck)
func (s *Session) PostCall(deckID string, callStr string) (c *[]Card, err error) {

	fcallStr := strings.Split(callStr, "_")

	type Call struct {
		Text   []string `json:"text"`
		String string   `json:"string"`
	}

	type Envelope struct {
		Calls []Call `json:"calls"`
	}

	data := Envelope{
		Calls: []Call{
			Call{Text: fcallStr, String: callStr},
		}}

	body, err := s.Request("POST", EndpointCalls(deckID), data)
	if err != nil {
		return
	}

	err = unmarshal(body, &c)
	return

}

// Posts a response to the deck (You must own the deck)
func (s *Session) PostResponse(deckID string, respStr string) (c *[]Card, err error) {

	type Response struct {
		Text   []string `json:"text"`
		String string   `json:"string"`
	}

	type Envelope struct {
		Responses []Response `json:"responses"`
	}

	data := Envelope{
		Responses: []Response{
			Response{Text: []string{respStr}, String: respStr},
		}}

	body, err := s.Request("POST", EndpointResponses(deckID), data)
	if err != nil {
		return
	}

	err = unmarshal(body, &c)
	return

}

// Deletes a call from the deck (You must own the deck)
func (s *Session) DeleteCall(deckID string, callID string) (c *Card, err error) {

	body, err := s.Request("DELETE", EndpointCall(deckID, callID), nil)
	if err != nil {
		return
	}

	err = unmarshal(body, &c)
	return

}

// Deletes a response from the deck (You must own the deck)
func (s *Session) DeleteResponse(deckID string, respID string) (c *Card, err error) {

	body, err := s.Request("DELETE", EndpointResponse(deckID, respID), nil)
	if err != nil {
		return
	}

	err = unmarshal(body, &c)
	return

}
