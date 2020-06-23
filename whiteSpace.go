package xsd

import "strings"

func WhiteSpacePreserve(s string) string {
	return s
}

func WhiteSpaceReplace(s string) string {
	s = strings.ReplaceAll(s, "\t", " ")
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\r", " ")
	return s
}

func WhiteSpaceCollapse(s string) string {
	s = WhiteSpaceReplace(s)
	s = strings.TrimLeft(s, " ")
	s = strings.TrimRight(s, " ")
	// TODO: replace 2+ spaces to one space
	return s
}
