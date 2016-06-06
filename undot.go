package undot

type Undot struct {
	Clusters map[string]*Cluster
	Edges    map[string][]*Edge
}

func NewUndot() *Undot {
	return &Undot{make(map[string]*Cluster), make(map[string][]*Edge)}
}

func Parse(dot string) (*Undot, error) {
	u := NewUndot()
	dot = ParseEdges(dot, u)
	dot, err := ParseClusters(dot, u)
	return u, err
}
