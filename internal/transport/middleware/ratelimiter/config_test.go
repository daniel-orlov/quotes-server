package ratelimiter_test

import (
	"testing"

	"github.com/daniel-orlov/quotes-server/internal/transport/middleware/ratelimiter"
)

func TestKey_String(t *testing.T) {
	tests := []struct {
		name string
		k    ratelimiter.Key
		want string
	}{
		{
			name: "ClientIP",
			k:    ratelimiter.ClientIP,
			want: "client_ip",
		},
		{
			name: "invalid",
			k:    ratelimiter.Key("invalid"),
			want: "invalid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRate_String(t *testing.T) {
	tests := []struct {
		name string
		r    ratelimiter.Rate
		want string
	}{
		{
			name: "Second",
			r:    ratelimiter.Second,
			want: "second",
		},
		{
			name: "Minute",
			r:    ratelimiter.Minute,
			want: "minute",
		},
		{
			name: "invalid",
			r:    ratelimiter.Rate("invalid"),
			want: "invalid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
