package models

import (
	"encoding/csv"
	"errors"
	"fmt"
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
		msg := fmt.Sprintf("open file %s error", path)
		log.Fatal(msg)
		return errors.New(msg)
	}
	defer f.Close()

	r := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
	r.LazyQuotes = true

	// read title
	record, err := r.Read()
	if err != nil {
		msg := "read csv error"
		log.Fatal(msg)
		return errors.New(msg)
	}
	c.headers = record

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else {
			msg := "read csv error"
			log.Fatal(msg)
			return errors.New(msg)
		}

		c.data = append(c.data, record)
	}

	return nil
}
