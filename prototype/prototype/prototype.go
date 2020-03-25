package prototype

// 原型对象需要实现此接口
type Cloneable interface {
	Clone() Cloneable
}

type Prototype struct{}

func (p *Prototype) Clone() Cloneable {
	pc := *p
	return &pc
}
