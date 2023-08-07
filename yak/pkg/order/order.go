package order

import (
	"github.com/lucavallin/yak-webshop/pkg/herd"
	"sync"
)

// Book keeps a list of orders that were made
type Book struct {
	Day    int             `xml:"day"`
	Herd   *herd.Herd      `xml:"herd"`
	Stock  herd.Stock      `xml:"stock"`
	Orders map[int][]Order `xml:"orders"`
	mu     sync.Mutex
}

// Order represents a new order
type Order struct {
	Customer string `json:"customer"`
	Items    Items  `json:"order"`
}

// Items represents the item requested or provided in an order
type Items struct {
	Milk  float64 `json:"milk,omitempty"`
	Skins int     `json:"skins,omitempty"`
}

// BookRepository provides an interface for different book repositories
type BookRepository interface {
	Get() *Book
	Save(book *Book)
}

// NewBook creates a new empty orders book
func NewBook(herd *herd.Herd) *Book {
	return &Book{
		Herd:   herd,
		Orders: make(map[int][]Order),
	}
}

// AddOrder adds the order to the book and returns the available items
func (b *Book) AddOrder(order Order, day int) Items {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.updateStock(day)

	// Get available items and update stock
	items := b.getAvailableItems(order)
	if items.Milk == 0 && items.Skins == 0 {
		// Return right away so no new order is created
		return items
	}

	// Update stock and day
	b.Stock.Milk -= items.Milk
	b.Stock.Wool -= items.Skins
	b.Day = day

	// Replace the order items with the available, and add the order
	order.Items = items
	b.Orders[day] = append(b.Orders[day], order)

	return items
}

// Get daily stock for the total of days elapsed since last order
// and refill the stock with the produced milk and wool
func (b *Book) updateStock(day int) {
	dailyStock := b.Herd.GetStock(day - b.Day)
	b.Stock.Milk += dailyStock.Milk
	b.Stock.Wool += dailyStock.Wool
}

// Check what can be given to the customer and update the quantity in stock
func (b *Book) getAvailableItems(order Order) Items {
	items := Items{}

	if b.Stock.Milk >= order.Items.Milk {
		items.Milk = order.Items.Milk
	}
	if b.Stock.Wool >= order.Items.Skins {
		items.Skins = order.Items.Skins
	}

	return items
}
