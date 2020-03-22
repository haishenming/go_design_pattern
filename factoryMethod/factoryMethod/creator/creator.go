package creator

import "fmt"

type Creator interface {
	SetMaterial(material string) // 准备原料
	Material() string            // 原料
	Result()                     // 获取工厂返回

}

// golang中没有继承，需要通过组合实现
type CreatorBase struct {
	material string
}

func (c *CreatorBase) Material() string {
	return c.material
}

func (c *CreatorBase) SetMaterial(material string) {
	c.material = material
}

type Creator1 struct {
	CreatorBase
}

func (c Creator1) Result() {
	fmt.Printf("creator 1 使用 %s\n", c.Material())
}

type Creator2 struct {
	CreatorBase
}

func (c Creator2) Result() {
	fmt.Printf("creator 2 使用 %s\n", c.Material())
}
