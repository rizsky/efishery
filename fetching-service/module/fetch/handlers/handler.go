package handlers

import (
	"fmt"

	"github.com/efishery/fetchin-service/module/fetch/helpers"
	"github.com/efishery/fetchin-service/module/fetch/services"
	"github.com/gin-gonic/gin"
)

type handler struct {
	agregatorServices services.AgregatorServices
}

//NewHandler :nodoc:
func NewHandler(agregatorServices services.AgregatorServices) *handler {
	return &handler{
		agregatorServices: agregatorServices,
	}
}

// GetList godoc
// @Summary GET product data and convert price to idr
// @Description GET product data and convert price to from idr to usd
// @Tags Fetch
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authentication header" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiRXhjZWxlbmN5IiwicGhvbmUiOiIrNjI4MjI3MjQyMjMyMCIsInJvbGUiOiJ1c2VyIiwidGltZXN0YW1wIjoiMjAyMC0wNS0xOVQwNDo1MzowNS40NDNaIiwiaWF0IjoxNTg5ODY0MDE0fQ.pnQXvaa9vFlIM4AKnptrbzbfztEGPXWU-ISDH-K1Vsg)
// @Success 200 {object} models.DataProducts
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /v1/list [get]
func (handler *handler) GetList(c *gin.Context) {
	products, err := handler.agregatorServices.GetAllAgregator()
	if err != nil {
		helpers.HTTPResponseError(c, err, 500, "something went wrong")
		return
	}
	helpers.SuccessResponse(c, products, "success get products list")

}

// GetByParam godoc
// @Summary GET data by param and return it with filter and graphical data by Price
// @Description get data by param, data graph (mean, median, min, max) below "products" please notice it senpai
// @Tags Fetch
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authentication header" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQWRtaW4gRXhjZWxlbmN5IiwicGhvbmUiOiIrNjI4MjI3MjQyMjEyIiwicm9sZSI6ImFkbWluIiwidGltZXN0YW1wIjoiMjAyMC0wNS0yMVQxNDoxMzoyOS4wNDlaIiwiaWF0IjoxNTkwMDcwNDI4fQ.qOoaipBvBQdJOCWxcHeTLjBFhMHg1EEtRHsdzX7dHVk)
// @Param area_provinsi query string false "area_provinsi" default(SUMATERA)
// @Param dari_tanggal query string false "dari_tanggal" default(2019-11-10T17:00:00.00Z)
// @Param sampai_tanggal query string false "sampai_tanggal" default(2020-08-30T13:35:00Z)
// @Success 200 {object} models.DataProductsWithGraph
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /v1/admin/list [get]
func (handler *handler) GetByParam(c *gin.Context) {
	filters := helpers.Filter{}

	sort := helpers.Sort{
		Type: c.Query("sort_by"),
	}
	filters.Province = c.Query("area_provinsi")
	filters.TimeFrom = c.Query("dari_tanggal")
	filters.TimeTo = c.Query("sampai_tanggal")

	fmt.Println(filters)

	products, err := handler.agregatorServices.GetByParam(filters, sort)
	if err != nil {
		helpers.HTTPResponseError(c, err, 500, "something went wrong")
		return
	}
	helpers.SuccessResponse(c, products, "success get products data")
}

// GetPrivateClaim godoc
// @Summary GET Private claim from JWT
// @Description Get Private claim from JWT with data phone, role, name, timestamp
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authentication header" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiRXhjZWxlbmN5IiwicGhvbmUiOiIrNjI4MjI3MjQyMjMyMCIsInJvbGUiOiJ1c2VyIiwidGltZXN0YW1wIjoiMjAyMC0wNS0xOVQwNDo1MzowNS40NDNaIiwiaWF0IjoxNTg5ODY0MDE0fQ.pnQXvaa9vFlIM4AKnptrbzbfztEGPXWU-ISDH-K1Vsg)
// @Success 200 {string} string "oke"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /v1/private-claim [get]
func (handler *handler) GetPrivateClaim(c *gin.Context) {
	decodedJwt, ok := c.Get("decodedJwt")
	if !ok {
		helpers.SuccessResponse(c, nil, "jwt data null")
	}
	fmt.Println(decodedJwt)
	helpers.SuccessResponse(c, decodedJwt, "JWT Private claimed")
}
