package undot

import (
	"testing"
)

func TestParseClustersEmpty(t *testing.T) {
	u := NewUndot()
	clusterstring := ""
	dot,err := ParseClusters(clusterstring, u)
	if len(u.Clusters) != 0 {
		t.Error("Should be no clusters")
	}
	if err == nil {t.Error("Should have been an error")}
	if len(dot) != 0 {
		t.Error("dot string should be empty")
	}
}

func TestParseClustersRoot(t *testing.T) {
	u := NewUndot()
	clusterstring := "graph { }"
	dot,err := ParseClusters(clusterstring, u)
	if len(u.Clusters) != 1 {
		t.Error("Should be a single cluster")
	}
	if err != nil {t.Error("Should not have been an error")}
	if u.Clusters["graph"] == nil {
		t.Error("Should have found a root cluster")
	}
	if len(dot) != 0 {
		t.Error("dot string should be empty")
	}
}

func TestParseClustersRootAttributes(t *testing.T) {
	u := NewUndot()
	clusterstring := "graph { rankdir=LR; }"
	dot,err := ParseClusters(clusterstring, u)
	if len(u.Clusters) != 1 {
		t.Error("Should be a single cluster")
	}
	if err != nil {t.Error("Should not have been an error")}
	if u.Clusters["graph"] == nil {
		t.Error("Should have found a root cluster")
	}
	if u.Clusters["graph"].Attributes["rankdir"] == "" {
		t.Error("Should have found an attribute")
	}
	if len(dot) != 0 {
		t.Error("dot string should be empty")
	}
}

func TestParseClustersSubcluster(t *testing.T) {
	u := NewUndot()
	clusterstring := "graph { subgraph cluster_1234 { } }"
	dot,err := ParseClusters(clusterstring, u)
	if len(u.Clusters) != 2 {
		t.Error("Should be a single cluster")
	}
	if err != nil {t.Error("Should not have been an error")}
	if u.Clusters["graph"] == nil {
		t.Error("Should have found a root cluster")
	}
	if u.Clusters["1234"] == nil {
		t.Error("Should have found cluster '1234'")
		t.Log(u.Clusters)
	}
	if len(dot) != 0 {
		t.Error("dot string should be empty")
	}
}