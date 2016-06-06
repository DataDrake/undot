package undot

import (
	"errors"
	"regexp"
)

var ROOT_CLUSTER_MATCH = regexp.MustCompile("^((?:di)?graph)\\s*{((?:.|\\n)*)}")
var CLUSTER_MATCH = regexp.MustCompile("subgraph cluster_(\\S*)\\s*{((?:.|\\n)*?)}")

type Cluster struct {
	Attributes map[string]string
	Nodes      map[string]*Node
}

func NewCluster() *Cluster {
	return &Cluster{make(map[string]string), make(map[string]*Node)}
}

func (c *Cluster) SetAttribute(name, value string) {
	c.Attributes[name] = value
}

func ParseClusters(dot string, u *Undot) (string, error) {
	r := NewCluster()
	rm := ROOT_CLUSTER_MATCH.FindStringSubmatch(dot)
	if rm == nil || len(rm) == 0 {
		return dot, errors.New("Could not find dot graph")
	}
	for _, m := range CLUSTER_MATCH.FindAllStringSubmatch(rm[2], -1) {
		c := NewCluster()
		m[2] = ParseAttributes(m[2], c)
		m[2] = ParseNodes(m[2], c)
		u.Clusters[m[1]] = c
	}
	dot = CLUSTER_MATCH.ReplaceAllString(dot, "")
	rm[2] = CLUSTER_MATCH.ReplaceAllString(rm[2], "")
	rm[2] = ParseNodes(rm[2], r)
	rm[2] = ParseAttributes(rm[2], r)
	u.Clusters[rm[1]] = r
	dot = ROOT_CLUSTER_MATCH.ReplaceAllString(dot, "")
	return dot, nil
}
