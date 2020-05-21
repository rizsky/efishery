package services

import (
	"github.com/efishery/fetchin-service/module/fetch/helpers"
	"github.com/efishery/fetchin-service/module/fetch/models"
	"github.com/efishery/fetchin-service/module/fetch/repositories"
)

//ShippingServices contract
type AgregatorServices interface {
	GetAllAgregator() ([]models.DataProducts, error)
	GetByParam(filter helpers.Filter, sort helpers.Sort) (models.DataProductsWithGraph, error)
}

//ShippingService godoc
type AgregatorService struct {
	efisheryCallRepo  repositories.ApiCallEfisheryRepositories
	converterCallRepo repositories.ApiCallConverterRepository
	cacheRepo         repositories.CacheRepository
}

//NewAgregatorServices instance
func NewAgregatorServices(
	efisheryCallRepo repositories.ApiCallEfisheryRepositories,
	converterCallRepo repositories.ApiCallConverterRepository,
	cacheRepo repositories.CacheRepository) AgregatorServices {
	return &AgregatorService{
		efisheryCallRepo:  efisheryCallRepo,
		converterCallRepo: converterCallRepo,
		cacheRepo:         cacheRepo,
	}
}

//GetAllAgregator services to return all data and with aggreated data fom currency
func (service *AgregatorService) GetAllAgregator() ([]models.DataProducts, error) {
	var (
		products []models.DataProducts
		rate     float64
	)
	//1. cache strategy first attemp every request
	rate, err := service.cacheRepo.Get(helpers.Convertion)
	if err != nil {
		return []models.DataProducts{}, err
	}

	//2. call the endpoint if no data from cache
	if rate == 0 {
		CurrencyFromAPI, err := service.converterCallRepo.GetCurrencyIDRtoUSD()
		if err != nil {
			return []models.DataProducts{}, err
		}
		rate = CurrencyFromAPI.IDR_USD

		//3. set cache after actually hit from endpoint
		err = service.cacheRepo.Set(helpers.Convertion, rate)
		if err != nil {
			return []models.DataProducts{}, err
		}
	}

	products, err = service.efisheryCallRepo.GetList()
	if err != nil {
		return []models.DataProducts{}, err
	}

	//4. assigning and conversion the price
	for i, value := range products {
		if value.Price == "" {
			continue
		}
		products[i].PriceUsd = helpers.RateConversion(value.Price, rate)
	}

	return products, nil
}

//GetAllAgregator services to return all data and with aggreated data fom currency
func (service *AgregatorService) GetByParam(filter helpers.Filter, sort helpers.Sort) (models.DataProductsWithGraph, error) {
	var (
		Response     models.DataProductsWithGraph
		dataFiltered []models.DataProducts
	)

	products, err := service.efisheryCallRepo.GetList()
	if err != nil {
		return models.DataProductsWithGraph{}, err
	}

	//filtering. assigning and conversion the price
	dataFiltered = helpers.FilterProduct(filter, products)

	//wildcard without filter
	if filter.Province == "" && filter.TimeFrom == "" && filter.TimeTo == "" && sort.Type == "" {
		dataFiltered = products
	}
	graph := helpers.CalculateGraph(dataFiltered)

	Response.Products = dataFiltered
	Response.Graph = graph

	return Response, nil
}
