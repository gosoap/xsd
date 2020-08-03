package xsd_test

import (
	"encoding/xml"
	"fmt"
	"testing"
	"time"

	"github.com/gosoap/xsd"
	"github.com/stretchr/testify/assert"
)

type xsdDateTime struct {
	XMLName xml.Name     `xml:"TS"`
	D       xsd.DateTime `xml:"d,attr"`
}

func ExampleDateTime() {
	type person struct {
		XMLName xml.Name     `xml:"person"`
		Born    xsd.DateTime `xml:"born,attr"`
		Name    string       `xml:",chardata"`
	}
	ts := person{Name: "John Doe", Born: *xsd.NewDateTime(2020, time.June, 22, 12, 0, 0, time.UTC)}
	b, _ := xml.Marshal(ts)
	fmt.Println(string(b))
	// Output:
	// <person born="2020-06-22T12:00:00Z">John Doe</person>
}

func TestDateTime_Marshal(t *testing.T) {
	ts := xsdDateTime{D: *xsd.NewDateTime(2020, time.June, 22, 12, 0, 0, time.UTC)}
	b, err := xml.Marshal(ts)
	assert.Nil(t, err)
	assert.Equal(t, `<TS d="2020-06-22T12:00:00Z"></TS>`, fmt.Sprintf("%s", b))

	ts = xsdDateTime{D: *xsd.NewDateTime(-2020, time.June, 22, 12, 0, 0, time.UTC)}
	b, err = xml.Marshal(ts)
	assert.Nil(t, err)
	assert.Equal(t, `<TS d="-2020-06-22T12:00:00Z"></TS>`, fmt.Sprintf("%s", b))
}

func TestDateTime_Unmarshal(t *testing.T) {
	ts := new(xsdDateTime)

	err := xml.Unmarshal([]byte(`<TS d="2001-10-26T21:32:52"></TS>`), ts)
	assert.Nil(t, err)
	assert.Equal(t, 2001, ts.D.Time().Year())
	assert.Equal(t, time.October, ts.D.Time().Month())
	assert.Equal(t, 26, ts.D.Time().Day())
	assert.Equal(t, 21, ts.D.Time().Hour())
	assert.Equal(t, 32, ts.D.Time().Minute())
	assert.Equal(t, 52, ts.D.Time().Second())
	assert.Equal(t, "UTC", ts.D.Time().Location().String())

	err = xml.Unmarshal([]byte(`<TS d="2001-10-26T21:32:52+02:00"></TS>`), ts)
	assert.Nil(t, err)
	assert.Equal(t, 2001, ts.D.Time().Year())
	assert.Equal(t, time.October, ts.D.Time().Month())
	assert.Equal(t, 26, ts.D.Time().Day())
	assert.Equal(t, 21, ts.D.Time().Hour())
	assert.Equal(t, 32, ts.D.Time().Minute())
	assert.Equal(t, 52, ts.D.Time().Second())
	assert.NotEqual(t, "UTC", ts.D.Time().Location().String())

	err = xml.Unmarshal([]byte(`<TS d="2001-10-26T19:32:52Z"></TS>`), ts)
	assert.Nil(t, err)
	assert.Equal(t, 2001, ts.D.Time().Year())
	assert.Equal(t, time.October, ts.D.Time().Month())
	assert.Equal(t, 26, ts.D.Time().Day())
	assert.Equal(t, 19, ts.D.Time().Hour())
	assert.Equal(t, 32, ts.D.Time().Minute())
	assert.Equal(t, 52, ts.D.Time().Second())
	assert.Equal(t, "UTC", ts.D.Time().Location().String())

	err = xml.Unmarshal([]byte(`<TS d="2001-10-26T19:32:52+00:00"></TS>`), ts)
	assert.Nil(t, err)
	assert.Equal(t, 2001, ts.D.Time().Year())
	assert.Equal(t, time.October, ts.D.Time().Month())
	assert.Equal(t, 26, ts.D.Time().Day())
	assert.Equal(t, 19, ts.D.Time().Hour())
	assert.Equal(t, 32, ts.D.Time().Minute())
	assert.Equal(t, 52, ts.D.Time().Second())
	assert.NotEqual(t, "UTC", ts.D.Time().Location().String())

	err = xml.Unmarshal([]byte(`<TS d="2001-10-26T21:32:52.12679"></TS>`), ts)
	assert.Nil(t, err)
	assert.Equal(t, 2001, ts.D.Time().Year())
	assert.Equal(t, time.October, ts.D.Time().Month())
	assert.Equal(t, 26, ts.D.Time().Day())
	assert.Equal(t, 21, ts.D.Time().Hour())
	assert.Equal(t, 32, ts.D.Time().Minute())
	assert.Equal(t, 52, ts.D.Time().Second())
	assert.Equal(t, 126790000, ts.D.Time().Nanosecond())
	assert.Equal(t, "UTC", ts.D.Time().Location().String())

	err = xml.Unmarshal([]byte(`<TS d="2001-10-26"></TS>`), ts)
	assert.NotNil(t, err)

	err = xml.Unmarshal([]byte(`<TS d="2001-10-32T21:32"></TS>`), ts)
	assert.NotNil(t, err)

	err = xml.Unmarshal([]byte(`<TS d="2001-10-26T25:32:52+02:00"></TS>`), ts)
	assert.NotNil(t, err)

	err = xml.Unmarshal([]byte(`<TS d="01-10-26T21:32"></TS>`), ts)
	assert.NotNil(t, err)
}
