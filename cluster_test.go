package undot

import (
	"testing"
)

func TestParseClustersEmpty(t *testing.T) {
	u := NewUndot()
	clusterstring := ""
	dot, err := ParseClusters(clusterstring, u)
	if len(u.Clusters) != 0 {
		t.Error("Should be no clusters")
	}
	if err == nil {
		t.Error("Should have been an error")
	}
	if len(dot) != 0 {
		t.Error("dot string should be empty")
	}
}

func TestParseClustersRoot(t *testing.T) {
	u := NewUndot()
	clusterstring :=
		`graph {
	}
	`
	_, err := ParseClusters(clusterstring, u)
	if len(u.Clusters) != 1 {
		t.Error("Should be a single cluster")
	}
	if err != nil {
		t.Error("Should not have been an error")
	}
	if u.Clusters["root"] == nil {
		t.Error("Should have found a root cluster")
	}
}

func TestParseClustersRootAttributes(t *testing.T) {
	u := NewUndot()
	clusterstring :=
		`graph {
		rankdir=LR;
	}
	`
	_, err := ParseClusters(clusterstring, u)
	if len(u.Clusters) != 1 {
		t.Error("Should be a single cluster")
	}
	if err != nil {
		t.Error("Should not have been an error")
	}
	if u.Clusters["root"] == nil {
		t.Error("Should have found a root cluster")
	}
	if u.Clusters["root"].Attributes["rankdir"] == "" {
		t.Error("Should have found an attribute")
		t.Log(u.Clusters["graph"].Nodes)
	}
}

func TestParseClustersSubcluster(t *testing.T) {
	u := NewUndot()
	clusterstring :=
		`graph {
		subgraph cluster_1234 {
		}
	}
	`
	_, err := ParseClusters(clusterstring, u)
	if len(u.Clusters) != 2 {
		t.Error("Should be a single cluster")
	}
	if err != nil {
		t.Error("Should not have been an error")
	}
	if u.Clusters["root"] == nil {
		t.Error("Should have found a root cluster")
	}
	if u.Clusters["1234"] == nil {
		t.Error("Should have found cluster '1234'")
		t.Log(u.Clusters)
	}
}

func TestParseClustersSubcluster2(t *testing.T) {
	u := NewUndot()
	clusterstring :=
		`graph {
	subgraph cluster_1234 {
		label="hello";
		T [label="bob"];
	}
}
`
	_, err := ParseClusters(clusterstring, u)
	if len(u.Clusters) != 2 {
		t.Error("Should be a single cluster")
	}
	if err != nil {
		t.Error("Should not have been an error")
	}
	if u.Clusters["root"] == nil {
		t.Error("Should have found a root cluster")
	}
	if u.Clusters["1234"] == nil {
		t.Error("Should have found cluster '1234'")
		t.Log(u.Clusters)
	}
	if len(u.Clusters["1234"].Nodes) != 1 {
		t.Error("Should have found exactly one node")
	}
	if u.Clusters["1234"].Nodes["T"] == nil {
		t.Error("Should have found node 'T'")
	}
	if len(u.Clusters["1234"].Nodes["T"].Attributes) != 1 {
		t.Error("Should have found one attribute")
	}
	if u.Clusters["1234"].Nodes["T"].Attributes["label"] == "" {
		t.Error("Should have found a label")
	}
}

var ASSIGNMENT_LOCAL_DEPENDENCY = `digraph{
	subgraph cluster_A{
		T [label="T | 100",type=TimedTask];
		S [label="S | 100",type=TimedTask];
	}
	Client;
	Client -> T;
	T -> S;
}
`

func TestParseClustersSubcluster3(t *testing.T) {
	u := NewUndot()
	_, err := ParseClusters(ASSIGNMENT_LOCAL_DEPENDENCY, u)
	if len(u.Clusters) != 2 {
		t.Error("Should be a single cluster")
	}
	if err != nil {
		t.Error("Should not have been an error")
	}
	if u.Clusters["root"] == nil {
		t.Error("Should have found a root cluster")
	}
	if u.Clusters["A"] == nil {
		t.Error("Should have found cluster 'A'")
		t.Log(u.Clusters)
	}
	if len(u.Clusters["A"].Nodes) != 2 {
		t.Error("Should have found exactly two nodes")
	}
	if u.Clusters["A"].Nodes["T"] == nil {
		t.Error("Should have found node 'T'")
	}
	if len(u.Clusters["A"].Nodes["T"].Attributes) != 2 {
		t.Error("Should have found one attribute")
		t.Log(u.Clusters["A"].Nodes["T"].Attributes)
	}
	if u.Clusters["A"].Nodes["T"].Attributes["label"] == "" {
		t.Error("Should have found a label")
	}
	if u.Clusters["A"].Nodes["S"] == nil {
		t.Error("Should have found node 'S'")
	}
	if len(u.Clusters["A"].Nodes["S"].Attributes) != 2 {
		t.Error("Should have found one attribute")
		t.Log(u.Clusters["A"].Nodes["S"].Attributes)
	}
	if u.Clusters["A"].Nodes["S"].Attributes["label"] == "" {
		t.Error("Should have found a label")
	}
}
