package xsd

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// AnyURI represent a URI (Uniform Resource Identifier)
type AnyURI string

// MarshalText implements the encoding.TextMarshaler interface
func (a AnyURI) MarshalText() ([]byte, error) {
	s := WhiteSpaceCollapse(string(a))
	return []byte(s), nil
}

// Base64Binary represent a binary content coded as "base64"
type Base64Binary []byte

// MarshalText implements the encoding.TextMarshaler interface
func (b Base64Binary) MarshalText() ([]byte, error) {
	s := base64.StdEncoding.EncodeToString([]byte(b))
	return []byte(s), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface
func (b *Base64Binary) UnmarshalText(text []byte) error {
	bin, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return err
	}
	*b = Base64Binary(bin)
	return nil
}

func (b Base64Binary) Binary() []byte {
	return []byte(b)
}

// Boolean represent a boolean (true or false)
type Boolean bool

func (b Boolean) Bool() bool {
	return bool(b)
}

// Byte represent a signed value of 8 bits
type Byte int8

func (b Byte) Int8() int8 {
	return int8(b)
}

// Date represent a gregorian calendar date
type Date time.Time

// MarshalText implements the encoding.TextMarshaler interface
func (d Date) MarshalText() ([]byte, error) {
	s := time.Time(d).Format("2006-01-02Z07:00")
	return []byte(s), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface
func (d *Date) UnmarshalText(text []byte) error {
	t, err := time.Parse("2006-01-02Z07:00", string(text))
	if err != nil {
		t, err = time.Parse("2006-01-02", string(text))
		if err != nil {
			return err
		}
	}
	*d = Date(t)
	return nil
}

func (d Date) Time() time.Time {
	return time.Time(d)
}

func NewDate(year int, month time.Month, day int, loc *time.Location) *Date {
	d := Date(time.Date(year, month, day, 0, 0, 0, 0, loc))
	return &d
}

// DateTime represent a instant of time (Gregorian calendar)
type DateTime time.Time

// MarshalText implements the encoding.TextMarshaler interface. The time is formatted in RFC 3339 format.
func (o DateTime) MarshalText() ([]byte, error) {
	s := time.Time(o).Format(time.RFC3339)
	return []byte(s), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface
func (d *DateTime) UnmarshalText(text []byte) error {
	t, err := time.Parse(time.RFC3339, string(text))
	if err != nil {
		t, err = time.Parse("2006-01-02T15:04:05", string(text))
		if err != nil {
			return err
		}
	}
	*d = DateTime(t)
	return nil
}

func (d DateTime) Time() time.Time {
	return time.Time(d)
}

func NewDateTime(year int, month time.Month, day, hour, min, sec int, loc *time.Location) *DateTime {
	d := DateTime(time.Date(year, month, day, hour, min, sec, 0, loc))
	return &d
}

// Decimal represent a decimal numbers
type Decimal float32

var decimalRegExp = regexp.MustCompile("^[+|-]?[\\d\\.]+$")

// UnmarshalText implements the encoding.TextUnmarshaler interface
func (d *Decimal) UnmarshalText(text []byte) error {
	if !decimalRegExp.Match(text) {
		return fmt.Errorf("%s is not valid decimal", text)
	}
	f, err := strconv.ParseFloat(string(text), 32)
	if err != nil {
		return err
	}
	*d = Decimal(float32(f))
	return nil
}

func (d Decimal) Float32() float32 {
	return float32(d)
}

// Double represent a IEEE 64-bit floating-point
type Double float64

// Duration reptresent a time durations
type Duration time.Duration

// MarshalText implements the encoding.TextMarshaler interface
func (o Duration) MarshalText() ([]byte, error) {
	d := time.Duration(o)
	out := "PT"
	h := int(d.Hours())
	if h > 0 {
		out = fmt.Sprintf("%s%dH", out, h)
	}
	m := int(d.Minutes() - float64(h)*60)
	if m > 0 {
		out = fmt.Sprintf("%s%dM", out, m)
	}
	s := d.Seconds() - float64(m)*60 - float64(h)*3600
	if s > 0 {
		out = fmt.Sprintf("%s%.fS", out, s)
	}
	return []byte(out), nil
}

// Entities represent a whitespace-separated list of unparsed entity references
type Entities []Entity

// MarshalText implements the encoding.TextMarshaler interface
func (o Entities) MarshalText() ([]byte, error) {
	var l []string
	for idx := range o {
		l = append(l, string(o[idx]))
	}
	s := strings.Join(l, " ")
	return []byte(s), nil
}

// Entity represent a reference to an unparsed entity
type Entity NCName

// Float represent a IEEE 32-bit floating-point
type Float float64

// GDay represent a recurring period of time: monthly day
type GDay time.Time

// MarshalText implements the encoding.TextMarshaler interface
func (o GDay) MarshalText() ([]byte, error) {
	s := time.Time(o).Format("---02Z07:00")
	return []byte(s), nil
}

// GMonth represent a recurring period of time: yearly month
type GMonth time.Time

// MarshalText implements the encoding.TextMarshaler interface
func (o GMonth) MarshalText() ([]byte, error) {
	s := time.Time(o).Format("--01Z07:00")
	return []byte(s), nil
}

// GMonthDay represent a recurring period of time: yearly day
type GMonthDay time.Time

// MarshalText implements the encoding.TextMarshaler interface
func (o GMonthDay) MarshalText() ([]byte, error) {
	s := time.Time(o).Format("--01-02Z07:00")
	return []byte(s), nil
}

// GYear represents a period of one year
type GYear time.Time

// MarshalText implements the encoding.TextMarshaler interface
func (o GYear) MarshalText() ([]byte, error) {
	s := time.Time(o).Format("2006Z07:00")
	return []byte(s), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface
func (o *GYear) UnmarshalText(text []byte) error {
	t, err := time.Parse("2006Z07:00", string(text))
	if err != nil {
		t, err = time.Parse("2006", string(text))
		if err != nil {
			return err
		}
	}
	*o = GYear(t)
	return nil
}

func (o GYear) Time() time.Time {
	return time.Time(o)
}

// GYearMonth represents a period of one month
type GYearMonth time.Time

// MarshalText implements the encoding.TextMarshaler interface
func (o GYearMonth) MarshalText() ([]byte, error) {
	s := time.Time(o).Format("2006-01Z07:00")
	return []byte(s), nil
}

// ID represent a definition of unique identifiers
type ID NCName

// Name represent a XML 1.O name
type Name Token

// NCName represent a unqualified names
type NCName Name

// NormalizedString represent a whitespace-replaced strings
type NormalizedString String

// String represent a any string
type String string

// MarshalText implements the encoding.TextMarshaler interface
// func (o String) MarshalText() ([]byte, error) {
// 	s := string(o)
// 	return []byte(s), nil
// }

// Token represent a whitespace-replaced and collapsed strings
type Token NormalizedString
