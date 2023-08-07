package herd

import (
	"math"
)

// Yak represents a yak in the herd
type Yak struct {
	Name string `xml:"name,attr" json:"name"`
	Sex string `xml:"sex,attr" json:"-"`
	Age float64 `xml:"age,attr" json:"age"`
	AgeLastShaved float64 `xml:"-" json:"age-last-shaved"`
}

const (
	daysInYear = 100
	firstShaveInDays = daysInYear
	deathInDays = 10 * daysInYear
	// Milk yield is given by the formula milkingConstant - D * milkingCoefficient liters (D = age in days)
	milkingConstant = 50
	milkingCoefficient = 0.03
	// Shaving period is given by the formula shavingConstant + D * shavingCoefficient days (D = age in days)
	shavingConstant = 8.00
	shavingCoefficient = 0.01
	sexFemale = "f"
)

// Milk returns how many liters of milk can be taken from the Yak
func (y Yak) Milk() float64 {
	if y.isAlive() && y.canBeMilked() {
		return milkingConstant - (getAgeInDays(y.Age) * milkingCoefficient)
	}

	return 0
}

// Shave returns how many units of wool can be taken from the Yak
func (y *Yak) Shave() int {
	if y.isAlive() && y.canBeShaved() {
		// y.AgeLastShaved has to be updated for checks on future shaving
		y.AgeLastShaved = y.Age
		return 1
	}

	return 0
}

// IncreaseAge adds days to the Yak's age and returns the new age in years
func (y *Yak) IncreaseAge(daysOlder int) {
	if daysOlder < 0 {
		return
	}

	newAge := getAgeInDays(y.Age) + float64(daysOlder)
	y.Age = math.Round((newAge / daysInYear) / 0.01) * 0.01
}

// canBeShaved checks whether or not the Yak can be shaved
// the yak must be at least 1 year old and haven't been shaved in the last shaving period
// these checks fulfill also Assumption n.2 (see assignment paper)
func (y Yak) canBeShaved() bool {
	ageInDays := getAgeInDays(y.Age)
	ageLastShavedInDays := getAgeInDays(y.AgeLastShaved)

	// These lines add the check so that a yak can only be shaved after {formula-value} days after the webshop
	// starts, but doing so doesn't give the results specified in the assignment.
	//if ageLastShavedInDays == 0 {
	//	y.AgeLastShaved = y.Age
	//	ageLastShavedInDays = getAgeInDays(y.AgeLastShaved)
	//}

	shavingPeriod := math.Ceil(shavingConstant + (ageInDays * shavingCoefficient))

	return ageInDays > firstShaveInDays &&
		ageInDays - ageLastShavedInDays > shavingPeriod
}

func (y Yak) canBeMilked() bool {
	return y.Sex == sexFemale
}

func (y Yak) isAlive() bool {
	return getAgeInDays(y.Age) < deathInDays
}

// ageInDays converts the yak's age (in years) to days
func getAgeInDays(age float64) float64 {
	return age * daysInYear
}