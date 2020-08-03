package xsd_test

import (
	"encoding/xml"
	"fmt"
	"testing"
	"time"

	"github.com/gosoap/xsd"
	"github.com/stretchr/testify/assert"
)

type xsdDate struct {
	XMLName xml.Name `xml:"TS"`
	D       xsd.Date `xml:"d,attr"`
}

func ExampleDate() {
	type person struct {
		XMLName xml.Name `xml:"person"`
		Born    xsd.Date `xml:"born,attr"`
		Name    string   `xml:",chardata"`
	}
	ts := person{Name: "John Doe", Born: *xsd.NewDate(2020, time.June, 22, time.UTC)}
	b, _ := xml.Marshal(ts)
	fmt.Println(string(b))
	// Output:
	// <person born="2020-06-22Z">John Doe</person>
}

func TestDate_Marshal(t *testing.T) {
	ts := xsdDate{D: *xsd.NewDate(2020, time.June, 22, time.UTC)}
	b, err := xml.Marshal(ts)
	assert.Nil(t, err)
	assert.Equal(t, `<TS d="2020-06-22Z"></TS>`, fmt.Sprintf("%s", b))

	ts = xsdDate{D: *xsd.NewDate(-2020, time.June, 22, time.UTC)}
	b, err = xml.Marshal(ts)
	assert.Nil(t, err)
	assert.Equal(t, `<TS d="-2020-06-22Z"></TS>`, fmt.Sprintf("%s", b))
}

func TestDate_Unmarshal(t *testing.T) {
	ts := new(xsdDate)

	err := xml.Unmarshal([]byte(`<TS d="2001-10-26"></TS>`), ts)
	assert.Nil(t, err)
	assert.Equal(t, 2001, ts.D.Time().Year())
	assert.Equal(t, time.October, ts.D.Time().Month())
	assert.Equal(t, 26, ts.D.Time().Day())
	assert.Equal(t, "UTC", ts.D.Time().Location().String())

	err = xml.Unmarshal([]byte(`<TS d="2001-10-26+02:00"></TS>`), ts)
	assert.Nil(t, err)
	assert.Equal(t, 2001, ts.D.Time().Year())
	assert.Equal(t, time.October, ts.D.Time().Month())
	assert.Equal(t, 26, ts.D.Time().Day())
	assert.NotEqual(t, "UTC", ts.D.Time().Location().String())

	err = xml.Unmarshal([]byte(`<TS d="2001-10-26Z"></TS>`), ts)
	assert.Nil(t, err)
	assert.Equal(t, 2001, ts.D.Time().Year())
	assert.Equal(t, time.October, ts.D.Time().Month())
	assert.Equal(t, 26, ts.D.Time().Day())
	assert.Equal(t, "UTC", ts.D.Time().Location().String())

	err = xml.Unmarshal([]byte(`<TS d="2001-10-26+00:00"></TS>`), ts)
	assert.Nil(t, err)
	assert.Equal(t, 2001, ts.D.Time().Year())
	assert.Equal(t, time.October, ts.D.Time().Month())
	assert.Equal(t, 26, ts.D.Time().Day())
	assert.NotEqual(t, "UTC", ts.D.Time().Location().String())

	err = xml.Unmarshal([]byte(`<TS d="2001-10"></TS>`), ts)
	assert.NotNil(t, err)

	err = xml.Unmarshal([]byte(`<TS d="2001-10-32"></TS>`), ts)
	assert.NotNil(t, err)

	err = xml.Unmarshal([]byte(`<TS d="2001-13-26+02:00"></TS>`), ts)
	assert.NotNil(t, err)

	err = xml.Unmarshal([]byte(`<TS d="01-10-26"></TS>`), ts)
	assert.NotNil(t, err)
}
