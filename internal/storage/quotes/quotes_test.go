package quotes_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/daniel-orlov/quotes-server/internal/storage/quotes"
)

func TestGetQuotes(t *testing.T) {
	tests := []struct {
		name    string
		wantLen int
	}{
		{
			name:    "GetQuotes() should return 24 quotes",
			wantLen: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantLen, len(quotes.GetQuotes()), "GetQuotes()")
		})
	}
}
