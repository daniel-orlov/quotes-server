package pow_test

import (
	"testing"

	"github.com/daniel-orlov/quotes-server/pkg/pow"
)

func TestChallengeKey_ClientID(t *testing.T) {
	tests := []struct {
		name         string
		challengeKey *pow.ChallengeKey
		want         string
	}{
		{
			name:         "ChallengeKey is nil",
			challengeKey: nil,
			want:         "",
		},
		{
			name:         "ChallengeKey has empty ClientID",
			challengeKey: &pow.ChallengeKey{},
			want:         "",
		},
		{
			name:         "ChallengeKey has non-empty ClientID",
			challengeKey: pow.NewChallengeKey("client-id", ""),
			want:         "client-id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.challengeKey
			if got := c.ClientID(); got != tt.want {
				t.Errorf("ClientID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChallengeKey_ResourceID(t *testing.T) {
	tests := []struct {
		name         string
		challengeKey *pow.ChallengeKey
		want         string
	}{
		{
			name:         "ChallengeKey is nil",
			challengeKey: nil,
			want:         "",
		},
		{
			name:         "ChallengeKey has empty ResourceID",
			challengeKey: &pow.ChallengeKey{},
			want:         "",
		},
		{
			name:         "ChallengeKey has non-empty ResourceID",
			challengeKey: pow.NewChallengeKey("", "resource-id"),
			want:         "resource-id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.challengeKey
			if got := c.ResourceID(); got != tt.want {
				t.Errorf("ResourceID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChallengeKey_String(t *testing.T) {
	tests := []struct {
		name         string
		challengeKey *pow.ChallengeKey
		want         string
	}{
		{
			name:         "ChallengeKey is nil",
			challengeKey: nil,
			want:         "",
		},
		{
			name:         "ChallengeKey has empty ClientID and ResourceID",
			challengeKey: &pow.ChallengeKey{},
			want:         ":",
		},
		{
			name:         "ChallengeKey has non-empty ClientID and empty ResourceID",
			challengeKey: pow.NewChallengeKey("client-id", ""),
			want:         "client-id:",
		},
		{
			name:         "ChallengeKey has empty ClientID and non-empty ResourceID",
			challengeKey: pow.NewChallengeKey("", "resource-id"),
			want:         ":resource-id",
		},
		{
			name:         "ChallengeKey has non-empty ClientID and non-empty ResourceID",
			challengeKey: pow.NewChallengeKey("client-id", "resource-id"),
			want:         "client-id:resource-id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.challengeKey
			if got := c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
