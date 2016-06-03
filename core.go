package undot

type Undot struct {
	Clusters map[string]*Cluster
	Edges    map[string]*Edge
}

func Parse(dot string) (*Undot, error) {
	u := &Undot{make(map[string]*Cluster)}
	dot = ParseEdges(dot, u)
	dot, err := ParseClusters(dot, u)
	return u, err
}
