package cardcastgo

import (
	"time"
)

type Session struct {
	Token string
}

type Carddeck struct {
	Name              string    `json:"name"`
	Code              string    `json:"code"`
	Description       string    `json:"description"`
	Unlisted          bool      `json:"unlisted"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	ExternalCopyright bool      `json:"external_copyright"`
	Category          string    `json:"category"`
	CallCount         string    `json:"call_count"`
	ResponseCount     string    `json:"response_count"`
	Author            struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	} `json:"author"`
	Rating string `json:"rating"`
}

type Card struct {
	Text      []string  `json:"text"`
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type Callpost struct {
	Calls []struct {
		Text       []string `json:"text"`
		String     string   `json:"string"`
		Validation struct {
			State   string `json:"state"`
			Message string `json:"message"`
		} `json:"validation"`
		EventChain interface{} `json:"eventChain"`
	} `json:"calls"`
}

type Responsepost struct {
	Responses []struct {
		Text       []string `json:"text"`
		String     string   `json:"string"`
		Validation struct {
			State   string `json:"state"`
			Message string `json:"message"`
		} `json:"validation"`
		EventChain interface{} `json:"eventChain"`
	} `json:"responses"`
}

// type Searchresult struct {
// 	Total   int `json:"total"`
// 	Results struct {
// 		Count  int `json:"count"`
// 		Offset int `json:"offset"`
// 		Data   []struct {
// 			Code              string    `json:"code"`
// 			Name              string    `json:"name"`
// 			Category          string    `json:"category"`
// 			ExternalCopyright bool      `json:"external_copyright"`
// 			CreatedAt         time.Time `json:"created_at"`
// 			UpdatedAt         time.Time `json:"updated_at"`
// 			CallCount         string    `json:"call_count"`
// 			ResponseCount     string    `json:"response_count"`
// 			Author            struct {
// 				ID       string `json:"id"`
// 				Username string `json:"username"`
// 			} `json:"author"`
// 			Rating string `json:"rating"`
// 		}
// 	}
// }
