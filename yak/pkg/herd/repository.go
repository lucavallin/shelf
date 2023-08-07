package herd

import (
	"encoding/xml"
	"github.com/prometheus/common/log"
	"io/ioutil"
)

// XMLFileRepository is a XML-file-based repo
type XMLFileRepository struct {
	filePath string
}

// NewXMLFileRepository creates a new HerdXMLFileRepository
func NewXMLFileRepository(filePath string) *XMLFileRepository {
	return &XMLFileRepository{filePath}
}

// Get reads a new Herd
func (h *XMLFileRepository) Get() *Herd {
	newHerd := &Herd{}
	content, _ := ioutil.ReadFile(h.filePath)
	if err := xml.Unmarshal(content, &newHerd); err != nil {
		log.Fatal(err)
	}

	return newHerd
}

// Save stores the herd
func (h *XMLFileRepository) Save(herd *Herd) {
	content, _ := xml.Marshal(herd)
	err := ioutil.WriteFile(h.filePath, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}