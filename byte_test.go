package xsd_test

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/gosoap/xsd"
	"github.com/stretchr/testify/assert"
)

type xsdByte struct {
	XMLName xml.Name `xml:"TS"`
	Int8    xsd.Byte `xml:"int8,attr"`
}

func TestByte_Marshal(t *testing.T) {
	ts := xsdByte{Int8: xsd.Byte(-128)}
	b, _ := xml.Marshal(ts)
	assert.Equal(t, `<TS int8="-128"></TS>`, fmt.Sprintf("%s", b))

	ts.Int8 = 127
	b, _ = xml.Marshal(ts)
	assert.Equal(t, `<TS int8="127"></TS>`, fmt.Sprintf("%s", b))
}

func TestByte_Unmarshal(t *testing.T) {
	ts := new(xsdByte)

	err := xml.Unmarshal([]byte(`<TS int8="27"></TS>`), ts)
	assert.Nil(t, err)
	assert.Equal(t, int8(27), ts.Int8.Int8())

	err = xml.Unmarshal([]byte(`<TS int8="-34"></TS>`), ts)
	assert.Nil(t, err)
	assert.Equal(t, int8(-34), ts.Int8.Int8())

	err = xml.Unmarshal([]byte(`<TS int8="+105"></TS>`), ts)
	assert.Nil(t, err)
	assert.Equal(t, int8(105), ts.Int8.Int8())

	err = xml.Unmarshal([]byte(`<TS int8="0"></TS>`), ts)
	assert.Nil(t, err)
	assert.Equal(t, int8(0), ts.Int8.Int8())

	err = xml.Unmarshal([]byte(`<TS int8="0A"></TS>`), ts)
	assert.NotNil(t, err)

	err = xml.Unmarshal([]byte(`<TS int8="1524"></TS>`), ts)
	assert.NotNil(t, err)

	err = xml.Unmarshal([]byte(`<TS int8="INF"></TS>`), ts)
	assert.NotNil(t, err)
}
