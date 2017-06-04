package models

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type CSV struct {
	header []string
	data   [][]string
}

func (c *CSV) GetHeader() []string {
	return c.header
}

func (c *CSV) GetData() [][]string {
	return c.data
}

func (c *CSV) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		msg := fmt.Sprintf("open file %s error, %s", path, err)
		log.Fatal(msg)
		return errors.New(msg)
	}
	defer f.Close()

	r := csv.NewReader(f)

	// read title
	record, err := r.Read()
	if err != nil {
		msg := fmt.Sprintf("read csv error, %s", err)
		log.Fatal(msg)
		return errors.New(msg)
	}
	c.header = record

	// format check
	if c.header[0] != "id" {
		msg := fmt.Sprint("csv format error, header 0 should be id field")
		log.Fatal(msg)
		return errors.New(msg)
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			msg := fmt.Sprintf("read csv error, %s", err)
			log.Fatal(msg)
			return errors.New(msg)
		}

		c.data = append(c.data, record)
	}

	return nil
}
