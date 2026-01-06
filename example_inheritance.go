package main

import "fmt"

type ParentI interface {
	echo(msg string)
	setName(s string)
}

type parent struct {
	name string
}

func (p *parent) echo(msg string) {
	fmt.Printf("%s: %s\n", p.name, msg)
}

func (p *parent) setName(s string) {
	p.name = s
}

func NewParent() ParentI {
	return &parent{}
}

type Child struct {
	parent
}

func (c *Child) echo(msg string) {
	fmt.Printf("%s: %s from child\n", c.name, msg)
}

func (c *Child) setName(s string) {
	c.name = s + " haha"
}

func NewChild() Child {
	return Child{}
}

func main() {
	parent := NewParent()
	parent.setName("test1")

	parent.echo("hello world")

	child := NewChild()
	child.setName("child 1")
	child.echo("hello world")
}
