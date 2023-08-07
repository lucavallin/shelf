package herd

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

var xmlYak = `
<labyak name="Betty-1" age="4" sex="f"/>
`

// Test unmarshalling from XML
func TestYak_Unmarshalling(t *testing.T) {
	expectedFemaleYak := Yak{Name: "Betty-1", Sex: "f", Age: 4}
	newFemaleYak := Yak{}
	xml.Unmarshal([]byte(xmlYak), &newFemaleYak)

	assert.Equal(t, expectedFemaleYak, newFemaleYak)
}

func TestYak_Milk(t *testing.T) {
	milkAmount := Yak{Name: "Betty-1", Sex: "f", Age: 4}.Milk()

	assert.Equal(t, float64(38), milkAmount)
}

// test canBeMilked when yak is male
func TestYak_OnlyFemalesCanBeMilked(t *testing.T) {
	maleYak := Yak{Name: "Betto-1", Sex: "m", Age: 4}
	milkAmount := maleYak.Milk()

	assert.Equal(t, float64(0), milkAmount)
}

// test Milk when yak is dead
func TestYak_CannotMilkDeadYak(t *testing.T) {
	deadYak := Yak{Name: "Betty-3", Sex: "f", Age: 11}
	milkAmount := deadYak.Milk()

	assert.Equal(t, float64(0), milkAmount)
}

func TestYak_Shave(t *testing.T) {
	yak := &Yak{Name: "Betty-1", Sex: "f", Age: 4}
	woolAmount := yak.Shave()

	assert.Equal(t, 1, woolAmount)
}

// test Shave updates AgeLastShaved
func TestYak_ShaveUpdatesAgeLastShaved(t *testing.T) {
	yak := Yak{Name: "Betty-1", Sex: "f", Age: 4}
	// 0.13 = 13 days in year format
	expectedAgeLastShaved := yak.Age + 0.13
	// Minimum shaving age, day 12 is elapsed
	yak.IncreaseAge(13)
	yak.Shave()

	assert.Equal(t, expectedAgeLastShaved, yak.AgeLastShaved)
}

func TestYak_canBeShaved(t *testing.T) {
	femaleYak := Yak{Name: "Betty-1", Sex: "f", Age: 4}
	femaleYak.IncreaseAge(13)

	assert.Equal(t, true, femaleYak.canBeShaved())
}

// test Shave when yak is dead
func TestYak_CannotShaveDeadYak(t *testing.T) {
	deadYak := Yak{Name: "Betty-3", Sex: "f", Age: 11}
	woolAmount := deadYak.Shave()

	assert.Equal(t, 0, woolAmount)
}

// test canBeShaved when yak is younger than a year
func TestYak_CannotShaveYoungerThanOneYear(t *testing.T) {
	youngYak := Yak{Name: "Betty-2", Sex: "f", Age: 0}
	woolAmount := youngYak.Shave()

	assert.Equal(t, 0, woolAmount)
}

// test cannot shave when yak have been shaved less than
// shavingConstant + D * shavingCoefficient days (D = age in days) days ago
func TestYak_CannotShaveRecentlyShaved(t *testing.T) {
	youngYak := Yak{Name: "Betty-2", Sex: "f", Age: 0}
	woolAmount := youngYak.Shave()

	assert.Equal(t, 0, woolAmount)
}

func TestYak_isAlive(t *testing.T) {
	yak := Yak{Name: "Betty-1", Sex: "f", Age: 4}
	yak.IncreaseAge(601)

	assert.Equal(t, false, yak.isAlive())
}

func TestYak_IncreaseAge(t *testing.T) {
	yak := Yak{Name: "Betty-1", Sex: "f", Age: 4}
	yak.IncreaseAge(13)

	assert.Equal(t, 4.13, yak.Age)
}