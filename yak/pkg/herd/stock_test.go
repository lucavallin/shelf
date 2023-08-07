package herd

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// test GetMilk
func TestHerd_GetStock(t *testing.T) {
	newHerd := &Herd{Yaks: []*Yak{{
		Name: "Betty-1",
		Age: 4,
		Sex: "f",
	}, {
		Name: "Betty-2",
		Age: 8,
		Sex: "f",
	}, {
		Name: "Betty-3",
		Age: 9.5,
		Sex: "f",
	}}}

	expectedStock := Stock{Milk: 1104.480, Wool: 3}
	stock := newHerd.GetStock(13)
	stock.Milk = math.Round(stock.Milk / 0.001) * 0.001

	assert.Equal(t, expectedStock, stock)
}

// test getting stock makes yak older (because it is done once a day)
func TestHerd_GetStockMakesYaksOlder(t *testing.T) {
	newHerd := &Herd{Yaks: []*Yak{{
		Name: "Betty-1",
		Age: 4,
		Sex: "f",
	}}}

	newHerd.GetStock(1)
	for _, yak := range newHerd.Yaks {
		assert.Equal(t, 4.01, yak.Age)
	}
}