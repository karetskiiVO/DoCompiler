package doreader

import (
	"bufio"
	"bytes"
	"io"
)

type CommentDeletor struct {
	reader *bufio.Reader

	node commentDeletorNode
}

var _ io.Reader = &CommentDeletor{}

func NewCommentDeletor(reader io.Reader) *CommentDeletor {
	return &CommentDeletor{
		reader: bufio.NewReader(reader),
		node:   &startNode{},
	}
}

func (cd *CommentDeletor) Read(p []byte) (int, error) {
	buffer := bytes.NewBuffer(p[:0])

	for buffer.Available() > 0 {
		r, _, err := cd.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				_, err := buffer.WriteString(cd.node.OnEOF())
				if err != nil {
					return buffer.Len(), err
				}
				return buffer.Len(), io.EOF
			}

			return buffer.Len(), err
		}

		bufnode, str := cd.node.Next(r)
		_, err = buffer.WriteString(str)
		if err != nil {
			cd.reader.UnreadRune()
			return buffer.Len(), err
		}

		cd.node = bufnode
	}

	return buffer.Len(), nil
}
