package undot

import (
	"testing"
)

var SIMPLE_DOT = `digraph {
	node [shape=record];
	rankdir=LR;
	subgraph cluster_0{
		label="CMP:1";
		A [label="A | 10"];
		B [label="B | 15"];
		C [label="C | 5"];
	}
	Client [label="Client | n/a"];
	A -> B;
	A -> C;
	Client -> A [label="(Ethernet,10M,3KM)"];
}`

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
	if len(u.Clusters["root"].Attributes) != 1 {
		t.Error("Should be one attribute")
	}
	if len(u.Clusters["root"].Nodes) != 2 {
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
	n, c := u.GetNodeByName("Client")
	if n == nil {
		t.Error("Should have found 'Client'")
	}
	if c == "" {
		t.Error("Should have found 'root' Cluster")
	}
	n2, c2 := u.GetNodeByName("bob")
	if n2 != nil {
		t.Error("should not have found 'bob'")
	}
	if c2 != "" {
		t.Error("should not have found cluster")
	}
}

var SIMPLE_DOT2 = `digraph{
	"A" [label="S0"];
	Switch [label="Switch | 1500"];
	"A" -> Switch [label="100m | 10n"];
}
`

func TestParseSimple2(t *testing.T) {
	u, err := Parse(SIMPLE_DOT2)
	if len(u.Clusters) != 1 {
		t.Error("Should be one clusters")
	}
	if len(u.Clusters["root"].Attributes) != 0 {
		t.Error("Should be no attribute")
	}
	if len(u.Clusters["root"].Nodes) != 2 {
		t.Error("Should be two nodes")
	}

	if len(u.Edges) != 1 {
		t.Error("Should be one edge")
	}
	if len(u.Edges["A"]) != 1 {
		t.Error("Should be one edge for 'A'")
	}
	if len(u.Edges["Switch"]) != 0 {
		t.Error("Should be no edges for 'Switch'")
	}
	if err != nil {
		t.Error("Should not have been an error")
	}
	n, c := u.GetNodeByName("A")
	if n == nil {
		t.Error("Should have found 'A'")
	}
	if c == "" {
		t.Error("Should have found 'root' Cluster")
	}
	n2, c2 := u.GetNodeByName("Bob")
	if n2 != nil {
		t.Error("should not have found 'Bob'")
	}
	if c2 != "" {
		t.Error("should not have found cluster")
	}
}
