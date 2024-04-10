package prototype

type INode interface {
	print(indentation string)
	clone() INode
}
