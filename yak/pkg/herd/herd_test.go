package herd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// test when herd is empty
func TestHerd_GetMilkAndWoolWhenNoYaks(t *testing.T) {
	newHerd := &Herd{}
	expectedStock := Stock{Milk: 0, Wool: 0}
	stock := newHerd.GetStock(13)

	assert.Equal(t, expectedStock, stock)
}

func TestHerd_Age(t *testing.T) {
	newHerd := &Herd{Yaks: []*Yak{{
		Name: "Betty-1",
		Sex: "f",
		Age: 4,
	}}}

	newHerd.Age(13)
	assert.Equal(t, &Yak{
		Name: "Betty-1",
		Sex: "f",
		Age: 4.13,
		AgeLastShaved: 4,
	}, newHerd.Yaks[0])
}