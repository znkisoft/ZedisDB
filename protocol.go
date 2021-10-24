package main

type Messager interface {
	String() []byte
}

type Message []byte

func (m Message) String() []byte {
	return m
}
