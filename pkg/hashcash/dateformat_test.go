package hashcash_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/daniel-orlov/quotes-server/pkg/hashcash"
)

func TestDateFormat_String(t *testing.T) {
	tests := []struct {
		name string
		df   hashcash.DateFormat
		want string
	}{
		{
			name: "DateFormatYY",
			df:   hashcash.DateFormatYY,
			want: "06",
		},
		{
			name: "DateFormatYYMM",
			df:   hashcash.DateFormatYYMM,
			want: "0601",
		},
		{
			name: "DateFormatYYMMDD",
			df:   hashcash.DateFormatYYMMDD,
			want: "060102",
		},
		{
			name: "DateFormatYYMMDDhhmm",
			df:   hashcash.DateFormatYYMMDDhhmm,
			want: "0601021504",
		},
		{
			name: "DateFormatYYMMDDhhmmss",
			df:   hashcash.DateFormatYYMMDDhhmmss,
			want: "060102150455",
		},
		{
			name: "InvalidDateFormat",
			df:   hashcash.DateFormat("invalid"),
			want: "invalid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.df.String(), "String()")
		})
	}
}

func TestDateFormat_IsValid(t *testing.T) {
	tests := []struct {
		name string
		df   hashcash.DateFormat
		want bool
	}{
		{
			name: "DateFormatYY",
			df:   hashcash.DateFormatYY,
			want: true,
		},
		{
			name: "DateFormatYYMM",
			df:   hashcash.DateFormatYYMM,
			want: true,
		},
		{
			name: "DateFormatYYMMDD",
			df:   hashcash.DateFormatYYMMDD,
			want: true,
		},
		{
			name: "DateFormatYYMMDDhhmm",
			df:   hashcash.DateFormatYYMMDDhhmm,
			want: true,
		},
		{
			name: "DateFormatYYMMDDhhmmss",
			df:   hashcash.DateFormatYYMMDDhhmmss,
			want: true,
		},
		{
			name: "InvalidDateFormat",
			df:   hashcash.DateFormat("invalid"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.df.IsValid(), "IsValid()")
		})
	}
}

func TestParseDateFormat(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		want    hashcash.DateFormat
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "DateFormatYY",
			args: args{
				date: "06",
			},
			want:    hashcash.DateFormatYY,
			wantErr: assert.NoError,
		},
		{
			name: "DateFormatYYMM",
			args: args{
				date: "0601",
			},
			want:    hashcash.DateFormatYYMM,
			wantErr: assert.NoError,
		},
		{
			name: "DateFormatYYMMDD",
			args: args{
				date: "060102",
			},
			want:    hashcash.DateFormatYYMMDD,
			wantErr: assert.NoError,
		},
		{
			name: "DateFormatYYMMDDhhmm",
			args: args{
				date: "0601021504",
			},
			want:    hashcash.DateFormatYYMMDDhhmm,
			wantErr: assert.NoError,
		},
		{
			name: "DateFormatYYMMDDhhmmss",
			args: args{
				date: "060102150405",
			},
			want:    hashcash.DateFormatYYMMDDhhmmss,
			wantErr: assert.NoError,
		},
		{
			name: "InvalidDateFormat",
			args: args{
				date: "invalid",
			},
			want:    "",
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hashcash.ParseDateFormat(tt.args.date)
			if !tt.wantErr(t, err, fmt.Sprintf("ParseDateFormat(%v)", tt.args.date)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ParseDateFormat(%v)", tt.args.date)
		})
	}
}
