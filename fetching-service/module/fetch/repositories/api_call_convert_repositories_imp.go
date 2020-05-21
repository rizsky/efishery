package repositories

import (
	"encoding/json"
	"net/http"

	"github.com/efishery/fetchin-service/module/fetch/models"
)

//ConverterCall :struct
type ConverterCall struct {
	URL string
}

//NewConverterRepository init function
func NewConverterRepository(URL string) ApiCallConverterRepository {
	return &ConverterCall{
		URL: URL,
	}
}

//GetCurrencyIDRtoUSD get list from efishery api and bind with mnodel in return
func (api *ConverterCall) GetCurrencyIDRtoUSD() (models.CurrencyIdrtoUsd, error) {
	var rate models.CurrencyIdrtoUsd
	resp, err := http.Get(api.URL)
	if err != nil {
		return models.CurrencyIdrtoUsd{}, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&rate)
	if err != nil {
		return models.CurrencyIdrtoUsd{}, err
	}
	return rate, nil
}
