package models

import "testing"

func TestCSV_Load(t *testing.T) {
	var c CSV
	result := c.Load("stores_detail_test.csv")

	if result != nil {
		t.Error("load error")
	}

	if c.header[0] != "id" {
		t.Errorf(`header[0] %s != "id"`, c.header[0])
	}

}
