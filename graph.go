package undot

type Graph struct{
	Name string
	Attributes map[string]string
	Subgraphs map[string]Graph
	Nodes []Node
}