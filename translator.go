package geolocation

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type yandexResponse struct {
	Code int      `json:"code"`
	Lang string   `json:"lang"`
	Text []string `json:"text"`
}

func translate(word string) (string, error) {
	rsp, err := http.Get(fmt.Sprintf("https://translate.yandex.net/api/v1.5/tr.json/translate?key=%s&text=%s&lang=ru", YandexKey, word))
	if err != nil {
		return "", err
	}

	var yr yandexResponse
	err = json.NewDecoder(rsp.Body).Decode(&yr)
	if err != nil {
		return "", err
	}

	if len(yr.Text) > 0 {
		return yr.Text[0], nil
	}

	return word, nil
}
