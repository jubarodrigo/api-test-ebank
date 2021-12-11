package domain

type Account struct {
	ID      int `json:"id"`
	Balance int `json:"balance"`
}

type Request struct {
	Type        string `json:"type"`
	Destination string `json:"destination"`
	Amount      int    `json:"amount"`
	Origin      string `json:"origin,omitempty"`
}

type Destination struct {
	Destination *Account `json:"destination"`
}

type Origin struct {
	Origin *Account `json:"origin"`
}

type Transfer struct {
	Origin      *Account `json:"origin"`
	Destination Account  `json:"destination"`
}
