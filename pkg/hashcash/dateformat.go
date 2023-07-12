package hashcash

import "time"

// DateFormat is the date format used in hashcash.
type DateFormat string

const (
	/*
	* Extract from http://hashcash.org/docs/hashcash.txt:
	*
	* period >= 2 years then time format YY is used rounded down to the nearest year start;
	* 2 years < period <= 2 months then time format YYMM is used rounded down to the nearest month start;
	* 2 months < period <= 2 days then time format YYMMDD is used rounded down to the beginning of the nearest day;
	* 2 days < period <= 2 minutes then time format YYMMDDhhmm is used rounded down to the beginning of the nearest minute;
	* period < 2 minutes then time format YYMMDDhhmmss is used in seconds.
	*
	 */

	// DateFormatYY is the date format used in hashcash,
	// indicating a stamp validity period of 2 years or more.
	DateFormatYY DateFormat = "06" // 2006

	// DateFormatYYMM is the date format used in hashcash,
	// indicating a stamp validity period between 2 months and 2 years.
	DateFormatYYMM DateFormat = "0601"

	// DateFormatYYMMDD is the date format used in hashcash,
	// indicating a stamp validity period between 2 days and 2 months.
	DateFormatYYMMDD DateFormat = "060102"

	// DateFormatYYMMDDhhmm is the date format used in hashcash,
	// indicating a stamp validity period between 2 minutes and 2 days.
	DateFormatYYMMDDhhmm DateFormat = "0601021504"

	// DateFormatYYMMDDhhmmss is the date format used in hashcash,
	// indicating stamp validity period less than 2 minutes.
	DateFormatYYMMDDhhmmss DateFormat = "060102150455"

	// MaxDurationYYMM is the maximum duration for which the date format is YYMM.
	MaxDurationYYMM = 2 * 365 * 24 * time.Hour // 2 years

	// MaxDurationYYMMDD is the maximum duration for which the date format is YYMMDD.
	MaxDurationYYMMDD = 2 * 30 * 24 * time.Hour // 2 months

	// MaxDurationYYMMDDhhmm is the maximum duration for which the date format is YYMMDDhhmm.
	MaxDurationYYMMDDhhmm = 2 * 24 * time.Hour // 2 days

	// MaxDurationYYMMDDhhmmss is the maximum duration for which the date format is YYMMDDhhmmss.
	MaxDurationYYMMDDhhmmss = 2 * time.Minute // 2 minutes
)

// String returns the string representation of the DateFormat.
func (df DateFormat) String() string {
	return string(df)
}

// IsValid checks if the DateFormat is valid.
func (df DateFormat) IsValid() bool {
	switch df {
	case DateFormatYY, DateFormatYYMM, DateFormatYYMMDD, DateFormatYYMMDDhhmm, DateFormatYYMMDDhhmmss:
		return true
	default:
		return false
	}
}

// ParseDateFormat parses the date format from the given string.
// It is a rough implementation of the date format parser, but it is enough for the purpose of it's use.
func ParseDateFormat(date string) (DateFormat, error) {
	// Switch the length of the date string
	switch len(date) {
	case len(DateFormatYY.String()):
		return DateFormatYY, nil
	case len(DateFormatYYMM.String()):
		return DateFormatYYMM, nil
	case len(DateFormatYYMMDD.String()):
		return DateFormatYYMMDD, nil
	case len(DateFormatYYMMDDhhmm.String()):
		return DateFormatYYMMDDhhmm, nil
	case len(DateFormatYYMMDDhhmmss.String()):
		return DateFormatYYMMDDhhmmss, nil
	default:
		return "", ErrInvalidDateFormat
	}
}
