package main

import "container/list"

type MessageBuffer struct {
	buf *list.List
}

func NewMessageBuffer() *MessageBuffer  {
	return &MessageBuffer{
		buf:list.New(),
	}
}

func (mb *MessageBuffer)Size() int {
	return mb.buf.Len()
}
