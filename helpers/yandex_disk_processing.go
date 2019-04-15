package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ImagesFromYandexDisk struct {
	Embedded EmbeddedStruct `json:"_embedded"`
}

type EmbeddedStruct struct {
	Items []ItemsStruct `json:"items"`
}

type ItemsStruct struct {
	ResourceId string `json:"resource_id" db:"resource_id"`
	ImageUrl   string `json:"file" db:"image_url"`
}

func GetImagesFromUserAppFolder(userToken string, client *http.Client) ([]ItemsStruct, error) {
	// todo: limit

	url := "https://cloud-api.yandex.net/v1/disk/resources?path=app:/&limit=100"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Printf("Error when try http.NewRequest, err: %s", err)
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("OAuth %s", userToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Printf("Error when try client.Do, err: %s", err)
		return nil, err
	}

	imagesFromYandexDisk := &ImagesFromYandexDisk{}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal([]byte(bodyBytes), imagesFromYandexDisk)

	if err != nil {
		log.Printf("Error when try json.Unmarshal, err: %s", err)
		return nil, err
	}

	return imagesFromYandexDisk.Embedded.Items, nil
}
