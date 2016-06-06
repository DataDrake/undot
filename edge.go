package undot

import "regexp"

var EDGE_MATCH = regexp.MustCompile("\\s*(\\S+)\\s+(<?->?)\\s+(\\S+)(?:\\s*\\[(.*)\\])?;")

type Edge struct {
	Attributes map[string]string
	NextNode   string
}

func NewEdge() *Edge {
	return &Edge{make(map[string]string), ""}
}

func (e *Edge) SetAttribute(name, value string) {
	e.Attributes[name] = value
}

func ParseEdges(dot string, u *Undot) string {
	for _, m := range EDGE_MATCH.FindAllStringSubmatch(dot, -1) {
		e := NewEdge()
		ParseAttributes(m[4], e)
		switch m[2] {
		case "->":
			e.NextNode = m[3]
			u.Edges[m[1]] = append(u.Edges[m[1]], e)
		case "<-":
			e.NextNode = m[1]
			u.Edges[m[3]] = append(u.Edges[m[3]], e)
		default:
		}

	}
	dot = EDGE_MATCH.ReplaceAllString(dot, "")
	return dot
}
