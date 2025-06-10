package ds

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestDoubleLLTestSuite(t *testing.T) {
	suite.Run(t, new(DoubleLLTestSuite))
}

type DoubleLLTestSuite struct {
	suite.Suite
}

func (s *DoubleLLTestSuite) TestDoubleLL_AddNodeAtTail() {
	doubleLL := NewDoubleLL[int]()
	doubleLL.AddNodeAtTail(1)
	s.Equal(doubleLL.Tail().Value, 1)

	doubleLL.AddNodeAtTail(2)
	s.Equal(doubleLL.Tail().Value, 2)
}

func (s *DoubleLLTestSuite) TestDoubleLL_ReturnHead() {
	doubleLL := NewDoubleLL[int]()
	head := doubleLL.Head()
	s.Nil(head)

	doubleLL.AddNodeAtTail(1)
	doubleLL.AddNodeAtTail(2)
	s.Equal(doubleLL.Head().Value, 1)
}

func (s *DoubleLLTestSuite) TestDoubleLL_AddNodeAfter() {
	doubleLL := NewDoubleLL[int]()
	doubleLL.AddNodeAtTail(1)
	doubleLL.AddNodeAtTail(2)
	head := doubleLL.Head()
	floatVal := 1.2
	doubleLL.AddNodeAfter(head, int(floatVal))
	s.Equal(doubleLL.Head().next.Value, 1)
}
