package herd

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var filePath = "../../data/herd_test.xml"

func TestHerdXMLFileRepository_Get(t *testing.T) {
	expectedHerd := &Herd{Yaks: []*Yak{{
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
	repo := NewXMLFileRepository(filePath)

	assert.Equal(t, expectedHerd, repo.Get())
}

func TestHerdXMLFileRepository_Save(t *testing.T) {
	expectedHerd := &Herd{Yaks: []*Yak{{
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

	repo := NewXMLFileRepository(filePath)
	repo.Save(expectedHerd)

	assert.Equal(t, repo.Get(), expectedHerd)
}