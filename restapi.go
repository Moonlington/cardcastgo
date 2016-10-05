package cardcastgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

var ErrJSONUnmarshal = errors.New("json unmarshal")

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

func (s *Session) request(method, urlStr, contentType string, b []byte) (response []byte, err error) {

	req, err := http.NewRequest(method, urlStr, bytes.NewBuffer(b))
	if err != nil {
		return
	}

	if s.Token != "" {
		req.Header.Set("x-auth-token", s.Token)
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

func (s *Session) Deck(deckID string) (cd *Carddeck, err error) {

	body, err := s.Request("GET", EndpointDeck(deckID), nil)
	if err != nil {
		return
	}

	err = unmarshal(body, &cd)
	return

}

func (s *Session) Calls(deckID string) (c *Card, err error) {

	body, err := s.Request("GET", EndpointCalls(deckID), nil)
	if err != nil {
		return
	}

	err = unmarshal(body, &c)
	return

}

func (s *Session) Responses(deckID string) (c *Card, err error) {

	body, err := s.Request("GET", EndpointResponses(deckID), nil)
	if err != nil {
		return
	}

	err = unmarshal(body, &c)
	return

}
