package xsd_test

import (
	"encoding/xml"
	"fmt"
	"testing"

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
