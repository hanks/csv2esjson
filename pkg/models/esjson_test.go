package models

import (
	"strings"
	"testing"
)

func TestItem_toJSON(t *testing.T) {
	var item Item

	item.index = map[string]map[string]string{}
	item.index["index"] = map[string]string{
		"_id": "1",
	}
	item.data = map[string]interface{}{
		"account_number": 1,
		"balance":        39225,
		"firstname":      "Amber",
		"lastname":       "Duke",
		"age":            32,
		"gender":         "M",
		"address":        "880 Holmes Lane",
		"employer":       "Pyrami",
		"email":          "amberduke@pyrami.com",
		"city":           "Brogan",
		"state":          "IL",
	}
	result, err := item.toJSON()
	if err != nil {
		t.Error("toJSON error")
	}

	t.Log(result)
	expect := `{"index":{"_id":"1"}}`
	if !strings.Contains(result, expect) {
		t.Errorf("result %s does contain %s", result, expect)
	}
}

func TestESJSON_AppendData(t *testing.T) {
	var e ESJSON

	header := []string{"id", "name", "kana"}
	record := []string{"1628", "グッデイ 玉名店", "グッデイ タマナテン"}

	e.AppendData(header, record)

	result := e.data[0].index["index"]["_id"]
	expect := "1628"
	if result != expect {
		t.Errorf("result %s != %s", result, expect)
	}

	result = e.data[0].data["name"].(string)
	expect = "グッデイ 玉名店"
	if result != expect {
		t.Errorf("result %s != %s", result, expect)
	}
}
