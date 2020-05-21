package repositories

import (
	"github.com/efishery/fetchin-service/module/fetch/models"
)

//ApiCallEfisheryRepositories :nodoc:
type ApiCallEfisheryRepositories interface {
	GetList() ([]models.DataProducts, error)
}

//ApiCallConverterRepository :nodoc
type ApiCallConverterRepository interface {
	GetCurrencyIDRtoUSD() (models.CurrencyIdrtoUsd, error)
}

//CacheRepository :nodoc:
type CacheRepository interface {
	Get(key string) (float64, error)
	Set(key string, value float64) error
}
