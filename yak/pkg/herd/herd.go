package herd

// Herd represents an herd
type Herd struct {
	Yaks []*Yak `xml:"labyak" json:"herd"`
}

// Repository is an interface to allow different repositories (future move to mongo)
type Repository interface {
	Get() *Herd
	Save(herd *Herd)
}

// Age makes the herd older to provide an overview at X days
func (h Herd) Age(elapsedDays int) {
	for day := 0; day < elapsedDays; day++ {
		for _, yak := range h.Yaks {
			yak.Shave()
			yak.IncreaseAge(1)
		}
	}
}