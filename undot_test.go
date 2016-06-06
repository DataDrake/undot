package undot

import (
	"testing"
)

var SIMPLE_DOT = string("digraph {\nnode [shape=record];\nrankdir=LR;\nsubgraph cluster_0{\nlabel=\"CMP:1\";\nA [label=\"A | 10\"];\nB [label=\"B | 15\"];\nC [label=\"C | 5\"];\n}\nClient [label=\"Client | n/a\"];\nA -> B;\nA -> C;\nClient -> A [label=\"(Ethernet,10M,3KM)\"];\n}")

func TestParseEmpty(t *testing.T) {
	dotstring := ""
	u, err := Parse(dotstring)
	if len(u.Clusters) != 0 {
		t.Error("Should be no clusters")
	}
	if err == nil {
		t.Error("Should have been an error")
	}
}

func TestParseSimple(t *testing.T) {
	u, err := Parse(SIMPLE_DOT)
	if len(u.Clusters) != 2 {
		t.Error("Should be two clusters")
	}
	if len(u.Clusters["0"].Nodes) != 3 {
		t.Error("Should be three nodes")
	}
	if len(u.Clusters["digraph"].Attributes) != 1 {
		t.Error("Should be one attribute")
	}
	if len(u.Clusters["digraph"].Nodes) != 2 {
		t.Error("Should be two nodes")
	}

	if len(u.Edges) != 2 {
		t.Error("Should be two sets of edges")
	}
	if len(u.Edges["A"]) != 2 {
		t.Error("Should be two edges for 'A'")
	}
	if len(u.Edges["Client"]) != 1 {
		t.Error("Should be one edge for 'Client'")
	}
	if err != nil {
		t.Error("Should not have been an error")
	}
}
