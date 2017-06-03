package models

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type CSV struct {
	headers []string
	data    [][]string
}

func (c *CSV) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("open file %s error", path)
	}
	defer f.Close()

	r := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
	r.LazyQuotes = true

	// read title
	record, err := r.Read()
	if err != nil {
		log.Fatal("read csv error")
	}

	c.headers = record

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else {
			log.Fatal("read csv error")
		}

		c.data = append(c.data, record)
	}
}

func (c *CSV) toESJSON() ESJSON {
	return nil
}
