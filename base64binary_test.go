package xsd_test

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/gosoap/xsd"
	"github.com/stretchr/testify/assert"
)

type xsdBase64Binary struct {
	XMLName xml.Name         `xml:"TS"`
	Value   xsd.Base64Binary `xml:",chardata"`
}

func ExampleBase64Binary() {
	type BinData struct {
		XMLName xml.Name         `xml:"binData"`
		Value   xsd.Base64Binary `xml:",chardata"`
	}
	content := "qwertyuiop"
	binData := BinData{Value: xsd.Base64Binary(content)}
	b, _ := xml.Marshal(binData)
	fmt.Println(string(b))
	// Output:
	// <binData>cXdlcnR5dWlvcA==</binData>
}
func TestBase64Binary_Marshal(t *testing.T) {
	ts := xsdBase64Binary{Value: xsd.Base64Binary("qwertyuiop")}
	b, _ := xml.Marshal(ts)
	assert.Equal(t, "<TS>cXdlcnR5dWlvcA==</TS>", fmt.Sprintf("%s", b))
}

func TestBase64Binary_Unmarshal(t *testing.T) {
	ts := new(xsdBase64Binary)

	err := xml.Unmarshal([]byte(`<TS>cXdlcnR5dWlvcA==</TS>`), ts)
	assert.Nil(t, err)
	assert.Equal(t, "qwertyuiop", string(ts.Value.Binary()))

	err = xml.Unmarshal([]byte(`<TS>qaz123</TS>`), ts)
	assert.NotNil(t, err)
}
