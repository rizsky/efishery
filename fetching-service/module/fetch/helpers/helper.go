package helpers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

//EfisheryURL url constant
const EfisheryURL = "https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list"

//ConverterURL convert currency url
const ConverterURL = "https://free.currconv.com/api/v7/convert?q=IDR_USD&compact=ultra&apiKey=1f17a8dcbe8099f862b6"

//Convertion constant of usd to idr
const Convertion = "USD_IDR"

//SuccessResponse succes helper
func SuccessResponse(context *gin.Context, data interface{}, message string) {
	context.JSON(200, gin.H{
		"code":    200,
		"error":   nil,
		"data":    data,
		"message": message,
		"status":  "success",
	})
}

//HTTPResponseError help for error response
func HTTPResponseError(context *gin.Context, err error, errorCode int, message string) {
	context.JSON(errorCode, gin.H{
		"code":    errorCode,
		"error":   err,
		"data":    nil,
		"message": message,
		"status":  "failed",
	})
}

//RateConversion from ret string value
func RateConversion(fromAmount string, toRate float64) string {
	from, err := strconv.ParseFloat(fromAmount, 64)
	if err != nil {
		return ""
	}
	result := from * toRate
	return fmt.Sprintf("%f", result)
}
