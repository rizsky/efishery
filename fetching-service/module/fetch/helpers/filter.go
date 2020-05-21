package helpers

import (
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/efishery/fetchin-service/module/fetch/models"
)

//Filter :nodoc
type Filter struct {
	Province string `form:"area_provinsi" json:"area_provinsi"`
	TimeFrom string `form:"dari_tanggal" json:"dari_tanggal"`
	TimeTo   string `form:"sampai_tanggal" json:"sampai_tanggal"`
}

//Sort :nodoc
type Sort struct {
	Type string `form:"sort_by" json:"sort_by"`
}

func FilterProduct(filter Filter, products []models.DataProducts) []models.DataProducts {
	var dataFiltered []models.DataProducts
	for i, _ := range products {
		switch {
		case filter.Province != "" && filter.TimeFrom != "" && filter.TimeTo != "":
			if strings.Contains(products[i].AreaProvinsi, filter.Province) && TimeComparation(filter.TimeFrom, filter.TimeTo, products[i].TglParsed) {
				dataFiltered = append(dataFiltered, products[i])
			}

		case filter.Province != "":
			if strings.Contains(products[i].AreaProvinsi, filter.Province) {
				dataFiltered = append(dataFiltered, products[i])
			}
		case filter.TimeFrom != "" && filter.TimeTo != "":
			if TimeComparation(filter.TimeFrom, filter.TimeTo, products[i].TglParsed) {
				dataFiltered = append(dataFiltered, products[i])
			}
		}
	}

	return dataFiltered
}

//TimeComparation compare time from string
func TimeComparation(from, to, in string) bool {
	fromTime, _ := time.Parse(time.RFC3339, from)
	toTime, _ := time.Parse(time.RFC3339, to)
	inTime, err := time.Parse(time.RFC3339, in)

	if err != nil {
		return false
	}
	return inTime.After(fromTime) && inTime.Before(toTime)
}

func CalculateGraph(products []models.DataProducts) models.GraphData {
	graph := models.GraphData{}
	arrData := []int{}
	total := 0

	if len(products) == 0 {
		return graph
	}

	for _, value := range products {
		currentPrice, err := strconv.Atoi(value.Price)
		if err != nil {
			currentPrice = 0
		}
		arrData = append(arrData, currentPrice)
		total += currentPrice
	}
	sort.Ints(arrData)
	avg := total / len(products)
	median := arrData[len(products)/2]

	//assignin to struct
	graph.Avg = strconv.Itoa(avg)
	graph.Max = strconv.Itoa(arrData[len(arrData)-1])
	graph.Min = strconv.Itoa(arrData[0])
	graph.Median = strconv.Itoa(median)

	return graph

}
