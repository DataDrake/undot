package undot

import (
	"regexp"
	"strings"
)

var NODE_MATCH = regexp.MustCompile("\\s+([\\w|\"]+)\\s*(?:\\s+\\[(.*)\\])?;")

type Node struct {
	Attributes map[string]string
}

func NewNode() *Node {
	return &Node{make(map[string]string)}
}

func (n *Node) SetAttribute(name, value string) {
	n.Attributes[name] = value
}

func ParseNodes(dot string, c *Cluster) string {
	for _, m := range NODE_MATCH.FindAllStringSubmatch(dot, -1) {
		n := NewNode()
		ParseAttributes(m[2], n)
		if strings.HasPrefix(m[1], "\"") && strings.HasSuffix(m[1], "\"") {
			m[1] = m[1][1 : len(m[1])-1]
		}
		c.Nodes[m[1]] = n
	}
	return NODE_MATCH.ReplaceAllString(dot, "")
}
