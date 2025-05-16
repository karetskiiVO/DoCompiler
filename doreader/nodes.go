package doreader

type commentDeletorNode interface {
	Next(r rune) (commentDeletorNode, string)
	OnEOF() string
}

type startNode struct{}

func (n *startNode) Next(r rune) (commentDeletorNode, string) {
	switch r {
	case '/':
		return &slashNode{}, ""
	default:
		return n, string(r)
	}
}

func (n *startNode) OnEOF() string {
	return ""
}

type slashNode struct{}

func (n *slashNode) Next(r rune) (commentDeletorNode, string) {
	switch r {
	case '/':
		return &singleLineCommentNode{}, ""
	case '*':
		return &multiLineCommentNode{}, ""
	default:
		return &startNode{}, "/" + string(r)
	}
}

func (n *slashNode) OnEOF() string {
	return "/"
}

type singleLineCommentNode struct{}

func (n *singleLineCommentNode) Next(r rune) (commentDeletorNode, string) {
	switch r {
	case '\n':
		return &startNode{}, "\n"
	default:
		return n, ""
	}
}

func (n *singleLineCommentNode) OnEOF() string {
	return ""
}

type multiLineCommentNode struct{}

func (n *multiLineCommentNode) Next(r rune) (commentDeletorNode, string) {
	switch r {
	case '*':
		return &endMultiLineCommentNode{}, ""
	default:
		return n, ""
	}
}

func (n *multiLineCommentNode) OnEOF() string {
	return ""
}

type endMultiLineCommentNode struct{}

func (n *endMultiLineCommentNode) Next(r rune) (commentDeletorNode, string) {
	switch r {
	case '/':
		return &startNode{}, ""
	default:
		return &multiLineCommentNode{}, ""
	}
}

func (n *endMultiLineCommentNode) OnEOF() string {
	return ""
}
