package undot

import "testing"

func TestParseEdgesEmpty(t *testing.T) {
	u := NewUndot()
	edgestring := ``
	dot := ParseEdges(edgestring, u)
	if len(u.Edges) != 0 {
		t.Error("Should be no edges")
	}
	if len(dot) != 0 {
		t.Error("dot string should be empty")
	}
}

func TestParseEdgesSingle(t *testing.T) {
	u := NewUndot()
	edgestring := `A -> B;`
	dot := ParseEdges(edgestring, u)
	if len(u.Edges) != 1 {
		t.Error("Should be one edge")
	}
	if len(u.Edges["A"]) != 1 {
		t.Error("Should only be one Edge for 'A'")
	}
	if u.Edges["A"][0].NextNode != "B" {
		t.Error("Edge did not go from A to B")
	}
	if len(dot) != 0 {
		t.Error("dot string should be empty")
	}
}

func TestParseEdgesSingleAttributes(t *testing.T) {
	u := NewUndot()
	edgestring := `A -> B [label="1234 | bob"];`
	dot := ParseEdges(edgestring, u)
	if len(u.Edges) != 1 {
		t.Error("Should be one edge")
	}
	if len(u.Edges["A"]) != 1 {
		t.Error("Should only be one Edge for A")
	}
	if u.Edges["A"][0].NextNode != "B" {
		t.Error("Edge did not go from A to B")
	}
	if len(u.Edges["A"][0].Attributes) != 1 {
		t.Error("Missing attributes for edge 'A -> B'")
	}
	if u.Edges["A"][0].Attributes["label"] != "1234 | bob" {
		t.Errorf("Found wrong attributes for 'A', '%s'", u.Edges["A"][0].Attributes["label"])
	}
	if len(dot) != 0 {
		t.Error("dot string should be empty")
	}
}

func TestParseEdgesMultiple(t *testing.T) {
	u := NewUndot()
	edgestring :=
	`A -> B;
	B <- C;
	C -> D;`
	dot := ParseEdges(edgestring, u)
	if len(u.Edges) != 2 {
		t.Error("Should be two sets of edges")
	}
	if len(u.Edges["A"]) != 1 {
		t.Error("Should be a single EDGE for 'A'")
	}
	if u.Edges["A"][0].NextNode != "B" {
		t.Error("Edge did not go from A to B")
	}
	if len(u.Edges["B"]) != 0 {
		t.Error("Should not have found edges for 'B'")
	}
	if len(u.Edges["C"]) != 2 {
		t.Error("Should have found two edges for 'C'")
	}
	if u.Edges["C"][0].NextNode != "B" {
		t.Error("Edge did not go from C to B")
	}
	if u.Edges["C"][1].NextNode != "D" {
		t.Error("Edge did not go from C to D")
	}
	if len(dot) != 0 {
		t.Error("dot string should be empty")
	}
}

func TestParseEdgesMultipleAttributes(t *testing.T) {
	u := NewUndot()
	edgestring :=
	`A -> B;
	B <- C [label=1234];
	C -> D;`
	dot := ParseEdges(edgestring, u)
	if len(u.Edges) != 2 {
		t.Error("Should be two sets of edges")
	}
	if len(u.Edges["A"]) != 1 {
		t.Error("Should be a single EDGE for 'A'")
	}
	if u.Edges["A"][0].NextNode != "B" {
		t.Error("Edge did not go from A to B")
	}
	if len(u.Edges["B"]) != 0 {
		t.Error("Should not have found edges for 'B'")
	}
	if len(u.Edges["C"]) != 2 {
		t.Error("Should have found two edges for 'C'")
	}
	if u.Edges["C"][0].NextNode != "B" {
		t.Error("Edge did not go from C to B")
	}
	if u.Edges["C"][0].Attributes["label"] != "1234" {
		t.Error("B's attributes did not match")
	}
	if u.Edges["C"][1].NextNode != "D" {
		t.Error("Edge did not go from C to D")
	}
	if len(dot) != 0 {
		t.Error("dot string should be empty")
	}
}
