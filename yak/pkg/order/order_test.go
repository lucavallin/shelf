package order

import (
	"github.com/lucavallin/yak-webshop/pkg/herd"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBook(t *testing.T) {
	newHerd := &herd.Herd{Yaks: []*herd.Yak{{
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
	newBook := NewBook(newHerd)

	assert.Equal(t, newHerd, newBook.Herd)
	assert.Equal(t, 0, newBook.Day)
}

// Here it would be good to have more sad-path tests
func TestBook_AddOrder(t *testing.T) {
	newHerd := &herd.Herd{Yaks: []*herd.Yak{{
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
	newBook := NewBook(newHerd)

	newOrder := Order{Customer: "Medvedev", Items: Items{Milk: 1100, Skins: 3}}
	newItems := newBook.AddOrder(newOrder, 13)

	assert.Equal(t, newOrder.Items, newItems)
	// Fast way to round, sorry
	assert.Equal(t, 4, int(newBook.Stock.Milk))
	assert.Equal(t, 0, newBook.Stock.Wool)
	assert.Equal(t, 13, newBook.Day)
	assert.Equal(t, 1, len(newBook.Orders))
}