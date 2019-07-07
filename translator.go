package geolocation

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type yandexResponse struct {
	Code int      `json:"code"`
	Lang string   `json:"lang"`
	Text []string `json:"text"`
}

func translate(word string) (string, error) {
	v := make(url.Values)
	v.Set("key", YandexKey)
	v.Set("text", word)
	v.Set("lang", "en-ru")

	rsp, err := http.Get("https://translate.yandex.net/api/v1.5/tr.json/translate?" + v.Encode())
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
