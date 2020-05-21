package models

//DataProducts :nodoc:
type DataProducts struct {
	UuID         string `json:"uuid"`
	Komoditas    string `json:"komoditas"`
	AreaProvinsi string `json:"area_provinsi"`
	AreaKota     string `json:"area_kota"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	PriceUsd     string `json:"price_usd"`
	TglParsed    string `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
}

//CurrencyIdrtoUsd :nodoc
type CurrencyIdrtoUsd struct {
	IDR_USD float64 `json:"IDR_USD"`
}
type DataProductsWithGraph struct {
	Products []DataProducts
	Graph    GraphData `json:"grafik_data_price"`
}
type GraphData struct {
	Min    string `json:"minimum"`
	Max    string `json:"maksimum"`
	Avg    string `json:"rata-rata"`
	Median string `json:"median"`
}
