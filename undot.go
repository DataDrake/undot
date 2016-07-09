package undot

type Undot struct {
	Clusters map[string]*Cluster
	Edges    map[string][]*Edge
}

func NewUndot() *Undot {
	return &Undot{make(map[string]*Cluster), make(map[string][]*Edge)}
}

func (u *Undot) GetNodeByName(name string) *Node {
	for _, c := range u.Clusters {
		n := c.Nodes[name]
		if n != nil {
			return n
		}
	}
	return nil
}

func Parse(dot string) (*Undot, error) {
	u := NewUndot()
	dot = ParseEdges(dot, u)
	_, err := ParseClusters(dot, u)
	return u, err
}
