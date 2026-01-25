package utils

import (
	"encoding/json"
	"strings"

	"github.com/gofiber/fiber/v3/client"
)

func GetStockSymbolWithExchange(symbol string, exchange string) string {
	var fullSymbol string
	switch exchange {
	case "NSE":
		fullSymbol = symbol + ".NS"
	default:
		fullSymbol = symbol
	}
	return fullSymbol
}

func GetExchangeData(symbol string, chartRange string, interval string) (map[string]interface{}, error) {

	url := "https://query1.finance.yahoo.com/v8/finance/chart/" + symbol + "?range=" + chartRange + "&interval=" + interval

	res, err := client.Get(url)

	if err != nil {
		return nil, err
	}

	var result map[string]interface{}

	if err := json.Unmarshal(res.Body(), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func GenerateDBStockSymbol(symbol string, exchange string) string {
	return symbol + "." + exchange
}

func ParseDBStockSymbol(dbSymbol string) (string, string) {
	parts := strings.SplitN(dbSymbol, ".", 2)
	return parts[0], parts[1]
}
