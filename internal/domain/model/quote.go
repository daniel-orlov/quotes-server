// Package model contains the domain models of the application.
package model

import "github.com/oklog/ulid/v2"

// Quote is a representation of a quote.
type Quote struct {
	// ID is the ID of the quote.
	// Read more about ULID here: https://github.com/oklog/ulid
	ID ulid.ULID `json:"id"`
	// Text is the text of the quote.
	Text string `json:"text"`
	// Author is the author of the quote.
	Author string `json:"author"`
}
