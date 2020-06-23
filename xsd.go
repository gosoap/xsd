package xsd

// AnyURI represent a URI (Uniform Resource Identifier)
//
// <xsd:simpleType name="anyURI" id="anyURI">
//  <xsd:restriction base="xsd:anySimpleType">
//   <xsd:whiteSpace value="collapse" fixed="true"/>
//  </xsd:restriction>
// </xsd:simpleType>
type AnyURI string

func (o AnyURI) MarshalText() ([]byte, error) {
	s := WhiteSpaceCollapse(string(o))
	return []byte(s), nil
}

// ID represent a definition of unique identifiers
//
// <xsd:simpleType name="ID" id="ID">
//  <xsd:restriction base="xsd:NCName"/>
// </xsd:simpleType>
type ID NCName

// Name represent a XML 1.O name
//
// <xsd:simpleType name="Name" id="Name">
//  <xsd:restriction base="xsd:token">
//   <xsd:pattern value="\i\c*"/>
//  </xsd:restriction>
// </xsd:simpleType>
type Name Token

// NCName represent a unqualified names
//
// <xsd:simpleType name="NCName" id="NCName">
//  <xsd:restriction base="xsd:Name">
//   <xsd:pattern value="[\i-[:]][\c-[:]]*"/>
//  </xsd:restriction>
// </xsd:simpleType>
type NCName Name

// normalizedstring represent a whitespace-replaced strings
//
// <xsd:simpleType name="normalizedString" id="normalizedString">
//  <xsd:restriction base="xsd:string">
//   <xsd:whiteSpace value="replace"/>
//  </xsd:restriction>
// </xsd:simpleType>
type NormalizedString String

// _string represent a any string
//
// <xsd:simpleType name="string" id="string">
//  <xsd:restriction base="xsd:anySimpleType">
//   <xsd:whiteSpace value="preserve"/>
//  </xsd:restriction>
// </xsd:simpleType>
type String string //anySimpleType

func (o String) MarshalText() ([]byte, error) {
	s := string(o)
	return []byte(s), nil
}

// token represent a whitespace-replaced and collapsed strings
//
// <xsd:simpleType name="token" id="token">
//  <xsd:restriction base="xsd:normalizedString">
//   <xsd:whiteSpace value="collapse"/>
//  </xsd:restriction>
// </xsd:simpleType>
type Token NormalizedString
