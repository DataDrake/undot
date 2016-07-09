package undot

import "testing"

func TestParseNodesEmpty(t *testing.T) {
	c := NewCluster()
	nodestring := ``
	dot := ParseNodes(nodestring, c)
	if len(c.Nodes) != 0 {
		t.Error("Should be no nodes")
	}
	if len(dot) != 0 {
		t.Error("dot strign should be empty")
	}
}

func TestParseNodesSingle(t *testing.T) {
	c := NewCluster()
	nodestring := ` A;`
	dot := ParseNodes(nodestring, c)
	if len(c.Nodes) != 1 {
		t.Error("Should be one node")
	}
	if c.Nodes["A"] == nil {
		t.Error("Could not find node 'A'")
	}
	if len(dot) != 0 {
		t.Error("dot strign should be empty")
	}
}

func TestParseNodesSingleAttributes(t *testing.T) {
	c := NewCluster()
	nodestring := ` A [label="1234 | bob"];`
	dot := ParseNodes(nodestring, c)
	if len(c.Nodes) != 1 {
		t.Error("Should be one node")
	}
	if c.Nodes["A"] == nil {
		t.Error("Could not find node 'A'")
	}
	if len(c.Nodes["A"].Attributes) != 1 {
		t.Error("Missign attributes for node 'A'")
	}
	if c.Nodes["A"].Attributes["label"] != "1234 | bob" {
		t.Errorf("Found wrong attributes for 'A', '%s'", c.Nodes["A"].Attributes["label"])
	}
	if len(dot) != 0 {
		t.Error("dot strign should be empty")
	}
}

func TestParseNodesMultiple(t *testing.T) {
	c := NewCluster()
	nodestring :=
		` A;
	 B;
	 C;`
	dot := ParseNodes(nodestring, c)
	if len(c.Nodes) != 3 {
		t.Error("Should be three nodes")
	}
	if c.Nodes["A"] == nil {
		t.Error("Could not find node 'A'")
	}
	if c.Nodes["B"] == nil {
		t.Error("Could not find node 'B'")
	}
	if c.Nodes["C"] == nil {
		t.Error("Could not find node 'C'")
	}
	if len(dot) != 0 {
		t.Error("dot strign should be empty")
	}
}

func TestParseNodesMultipleAttributes(t *testing.T) {
	c := NewCluster()
	nodestring :=
		` A;
	 B [label=1234];
	 C;`
	dot := ParseNodes(nodestring, c)
	if len(c.Nodes) != 3 {
		t.Error("Should be three nodes")
	}
	if c.Nodes["A"] == nil {
		t.Error("Could not find node 'A'")
	}
	if c.Nodes["B"] == nil {
		t.Error("Could not find node 'B'")
	}
	if c.Nodes["B"].Attributes["label"] != "1234" {
		t.Error("B's attributes did not match")
	}
	if c.Nodes["C"] == nil {
		t.Error("Could not find node 'C'")
	}
	if len(dot) != 0 {
		t.Error("dot strign should be empty")
	}
}
