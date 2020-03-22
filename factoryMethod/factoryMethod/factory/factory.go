package factory

import (
	"go_design_pattern/factoryMethod/factoryMethod/creator"
)

type Factory interface {
	Product() creator.Creator
}

type Factory1 struct {
}

func (f Factory1) Product() creator.Creator {
	return &creator.Creator1{}
}

type Factory2 struct {
}

func (f Factory2) Product() creator.Creator {
	return &creator.Creator2{}
}
