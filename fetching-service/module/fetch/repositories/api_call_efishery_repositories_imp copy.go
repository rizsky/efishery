package repositories

import (
	"encoding/json"
	"net/http"

	"github.com/efishery/fetchin-service/module/fetch/models"
)

//EfisheryCall :struct
type EfisheryCall struct {
	URL string
}

//NewEfisheryCallRepository init function
func NewEfisheryCallRepository(URL string) ApiCallEfisheryRepositories {
	return &EfisheryCall{
		URL: URL,
	}
}

//GetList get list from efishery api and bind with mnodel in return
func (api *EfisheryCall) GetList() ([]models.DataProducts, error) {
	var products []models.DataProducts
	resp, err := http.Get(api.URL)
	if err != nil {
		return []models.DataProducts{}, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&products)
	if err != nil {
		return []models.DataProducts{}, err
	}

	return products, nil
}
