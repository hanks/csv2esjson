package transfer

import (
	"csv2esjson/pkg/models"
	"errors"
	"fmt"
	"log"
)

func Execute(in string, out string) error {
	c, _ := readCSV(in)
	e, _ := parse(c)
	writeESJSON(e, out)

	return nil
}

func readCSV(in string) (models.CSV, error) {
	var c models.CSV
	err := c.Load(in)

	if err != nil {
		msg := fmt.Sprintf("open csv file %s error, %s", in, err)
		log.Fatal(msg)
		return c, errors.New(msg)
	}

	return c, nil
}

func parse(c models.CSV) (models.ESJSON, error) {
	var e models.ESJSON

	header := c.GetHeader()

	for _, record := range c.GetData() {
		e.AppendData(header, record)
	}

	return e, nil
}

func writeESJSON(e models.ESJSON, out string) error {
	err := e.ExportJSON(out)

	if err != nil {
		msg := fmt.Sprintf("write esjson to %s error, %s", out, err)
		log.Fatal(msg)
		return errors.New(msg)
	}

	return nil
}
