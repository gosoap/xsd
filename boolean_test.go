package xsd_test

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/gosoap/xsd"
	"github.com/stretchr/testify/assert"
)

type xsdBoolean struct {
	XMLName xml.Name    `xml:"TS"`
	Flag    xsd.Boolean `xml:"flag,attr"`
}

func ExampleBoolean() {
	type Book struct {
		XMLName   xml.Name    `xml:"book"`
		Available xsd.Boolean `xml:"available,attr"`
		Title     string      `xml:",chardata"`
	}

	book := Book{Title: "1984", Available: xsd.Boolean(true)}
	b, _ := xml.Marshal(book)
	fmt.Println(string(b))
	// Output:
	// <book available="true">1984</book>
}

func TestBoolean_Marshal(t *testing.T) {
	ts := xsdBoolean{Flag: xsd.Boolean(true)}
	b, _ := xml.Marshal(ts)
	assert.Equal(t, `<TS flag="true"></TS>`, fmt.Sprintf("%s", b))

	ts.Flag = false
	b, _ = xml.Marshal(ts)
	assert.Equal(t, `<TS flag="false"></TS>`, fmt.Sprintf("%s", b))
}

func TestBoolean_Unmarshal(t *testing.T) {
	ts := new(xsdBoolean)

	err := xml.Unmarshal([]byte(`<TS flag="true"></TS>`), ts)
	assert.Nil(t, err)
	assert.Equal(t, true, ts.Flag.Bool())

	err = xml.Unmarshal([]byte(`<TS flag="false"></TS>`), ts)
	assert.Nil(t, err)
	assert.Equal(t, false, ts.Flag.Bool())

	err = xml.Unmarshal([]byte(`<TS flag="1"></TS>`), ts)
	assert.Nil(t, err)
	assert.Equal(t, true, ts.Flag.Bool())

	err = xml.Unmarshal([]byte(`<TS flag="0"></TS>`), ts)
	assert.Nil(t, err)
	assert.Equal(t, false, ts.Flag.Bool())

	err = xml.Unmarshal([]byte(`<TS flag="ok"></TS>`), ts)
	assert.NotNil(t, err)
}
