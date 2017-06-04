package models

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type Item struct {
	index map[string]map[string]string
	data  map[string]interface{}
}

func (item *Item) toJSON() (string, error) {
	var buf bytes.Buffer

	b, err := json.Marshal(item.index)
	if err != nil {
		msg := fmt.Sprintf("json parse error, %s", item.index)
		log.Fatal(msg)
		return "", errors.New(msg)
	}
	buf.Write(b)
	buf.WriteString("\n")

	b, err = json.Marshal(item.data)
	if err != nil {
		msg := fmt.Sprintf("json parse error, %s", item.index)
		log.Fatal(msg)
		return "", errors.New(msg)
	}
	buf.Write(b)
	buf.WriteString("\n")

	return buf.String(), nil
}

type ESJSON struct {
	data []Item
}

func (e *ESJSON) ExportJSON(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	defer f.Close()

	if err != nil {
		msg := fmt.Sprintf("open file %s error, %s", path, err)
		log.Fatal(msg)
		return errors.New(msg)
	}

	w := bufio.NewWriter(f)
	for _, item := range e.data {
		s, err := item.toJSON()
		if err != nil {
			msg := fmt.Sprintf("toJson error, %s", err)
			log.Fatal(msg)
			return errors.New(msg)
		}
		w.WriteString(s)
	}
	w.Flush()

	return nil
}

func (e *ESJSON) AppendData(header []string, record []string) error {
	var item Item

	item.index = map[string]map[string]string{}
	item.data = map[string]interface{}{}

	item.index["index"] = map[string]string{
		"_id": record[0],
	}

	for i, v := range header {
		item.data[v] = record[i]
	}

	e.data = append(e.data, item)

	return nil
}
