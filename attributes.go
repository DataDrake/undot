package undot

import (
	"regexp"
	"strings"
)

var ATTRIBUTE_MATCH = regexp.MustCompile("(\\w+)=(\\w+|\".+?\"),?")

type Attributable interface {
	SetAttribute(name, value string)
}

func ParseAttributes(dot string, a Attributable) string {
	for _, m := range ATTRIBUTE_MATCH.FindAllStringSubmatch(dot, -1) {
		m[2] = strings.Replace(m[2], "\"", "", -1)
		a.SetAttribute(m[1], m[2])
	}
	return ATTRIBUTE_MATCH.ReplaceAllString(dot, "")
}
