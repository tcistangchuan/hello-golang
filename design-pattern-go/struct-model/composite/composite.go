package main

import "fmt"

/*组合模式
组合模式，其实说的是对象包含对象的问题，通过组合的方式（在对象内部引用对象）来进行布局，
我认为这种组合是区别于继承的，而另一层含义是指树形结构子节点的抽象（将叶子节点与数枝节
点抽象为子节点），区别于普通的分别定义叶子节点与数枝节点的方式。

应用场景：
文件目录显示，多及目录呈现等树形结构数据的操作。
*/

type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(string)
	AddChild(Component)
	Print(string)
}

const (
	LeafNode = iota
	CompositeNode
)

func NewComponent(kind int, name string) Component {
	var c Component
	switch kind {
	case LeafNode:
		c = NewLeaf()
	case CompositeNode:
		c = NewComposite()
	}

	c.SetName(name)
	return c
}

type component struct {
	parent Component
	name   string
}

func (c *component) Parent() Component {
	return c.parent
}

func (c *component) SetParent(parent Component) {
	c.parent = parent
}

func (c *component) Name() string {
	return c.name
}

func (c *component) SetName(name string) {
	c.name = name
}

func (c *component) AddChild(Component) {}

func (c *component) Print(string) {}

type Leaf struct {
	component
}

func NewLeaf() *Leaf {
	return &Leaf{}
}

func (c *Leaf) Print(pre string) {
	fmt.Printf("%s-%s\n", pre, c.Name())
}

type Composite struct {
	component
	childes []Component
}

func NewComposite() *Composite {
	return &Composite{
		childes: make([]Component, 0),
	}
}

func (c *Composite) AddChild(child Component) {
	child.SetParent(c)
	c.childes = append(c.childes, child)
}

func (c *Composite) Print(pre string) {
	fmt.Printf("%s+%s\n", pre, c.Name())
	pre += " "
	for _, comp := range c.childes {
		comp.Print(pre)
	}
}

func main() {
	root := NewComponent(CompositeNode, "root")
	c1 := NewComponent(CompositeNode, "c1")
	c2 := NewComponent(CompositeNode, "c2")
	c3 := NewComponent(CompositeNode, "c3")

	l1 := NewComponent(LeafNode, "l1")
	l2 := NewComponent(LeafNode, "l2")
	l3 := NewComponent(LeafNode, "l3")

	root.AddChild(c1)
	root.AddChild(c2)
	c1.AddChild(c3)
	c1.AddChild(l1)
	c2.AddChild(l2)
	c2.AddChild(l3)

	root.Print("")
}
