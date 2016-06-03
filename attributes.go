package undot

import "regexp"

var ATTRIBUTE_MATCH = regexp.MustCompile("(\\S+)=(\\w+|\"[^\"]+\"),?")

type Attributable interface {
	SetAttribute(name, value string)
}

func ParseAttributes(dot string, a Attributable) string {
	for _, m := range ATTRIBUTE_MATCH.FindAllStringSubmatch(dot, -1) {
		a.SetAttribute(m[1], m[2])
	}
	dot = ATTRIBUTE_MATCH.ReplaceAllString(dot, "")
	return dot
}
