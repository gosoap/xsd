package xsd_test

import (
	"encoding/xml"
	"fmt"
	"testing"
	"time"

	"github.com/gosoap/xsd"
	"github.com/stretchr/testify/assert"
)

func TestAnyURI(t *testing.T) {
	type TS struct {
		XMLName xml.Name   `xml:"TS"`
		URI     xsd.AnyURI `xml:"uri,attr"`
	}
	ts := TS{URI: xsd.AnyURI(" 123 ")}
	b, _ := xml.Marshal(ts)
	assert.Equal(t, "<TS uri=\"123\"></TS>", fmt.Sprintf("%s", b))
}

func TestDate(t *testing.T) {
	type TS struct {
		XMLName xml.Name `xml:"TS"`
		D       xsd.Date `xml:"d,attr"`
	}
	ts := TS{D: xsd.Date(time.Date(2020, time.June, 22, 0, 0, 0, 0, time.Local))}
	b, _ := xml.Marshal(ts)
	assert.Equal(t, "<TS d=\"2020-06-22\"></TS>", fmt.Sprintf("%s", b))
}
