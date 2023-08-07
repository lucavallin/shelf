package herd

// Stock contains the amount of milk and wool produced
type Stock struct {
	Milk float64 `json:"milk"`
	Wool int `json:"wool"`
}

// GetStock retrieves milk and wool from the herd, aging the yaks accordingly
// Shepherd does this once a day, so every loop the yak is one day old
func (h Herd) GetStock(elapsedDays int) Stock {
	var totalMilk float64
	var totalWool int

	for day := 0; day < elapsedDays; day++ {
		for _, yak := range h.Yaks {
			totalMilk += yak.Milk()
			totalWool += yak.Shave()
			yak.IncreaseAge(1)
		}
	}

	return Stock{Milk: totalMilk, Wool: totalWool}
}