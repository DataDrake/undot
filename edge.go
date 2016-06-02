package undot

const (
	NONE = -1
	LEFT = 1
	RIGHT = 2
)

type Edge struct {
	Attributes map[string]string
	Direction int
	Left *Node
	Right *Node
}
