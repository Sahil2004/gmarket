package utils

import (
	"encoding/json"
	"math"
	"math/rand"
	"strings"
	"time"

	"sync"

	"github.com/Sahil2004/gmarket/server/dtos"
	"github.com/gofiber/fiber/v3/client"
)

var priceStore = struct {
	sync.Mutex
	m map[string]float64
}{
	m: make(map[string]float64),
}

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

func GetMarketData(symbol string, exchange string) (float64, []dtos.LevelDTO, []dtos.LevelDTO, error) {
	key := GenerateDBStockSymbol(symbol, exchange)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	priceStore.Lock()
	defer priceStore.Unlock()

	ltp, exists := priceStore.m[key]

	if !exists {
		ltp = float64(1000 + r.Intn(2000))
	}

	delta := (r.Float64() - 0.5) * 2 * (ltp * 0.002) // +/-0.2% change
	ltp += delta
	ltp = math.Round(ltp*100) / 100

	priceStore.m[key] = ltp

	const levels = 5
	spreadBase := ltp * 0.0005

	bids := make([]dtos.LevelDTO, 0, levels)
	asks := make([]dtos.LevelDTO, 0, levels)

	for i := 1; i <= levels; i++ {
		spread := spreadBase * float64(i)

		bidPrice := math.Round((ltp-spread)*100) / 100
		askPrice := math.Round((ltp+spread)*100) / 100

		bidOrders := r.Intn(15) + 5
		askOrders := r.Intn(15) + 5

		bidQty := bidOrders * (r.Intn(40) + 10)
		askQty := askOrders * (r.Intn(40) + 10)

		bids = append(bids, dtos.LevelDTO{
			Price:  bidPrice,
			Orders: bidOrders,
			Qty:    bidQty,
		})

		asks = append(asks, dtos.LevelDTO{
			Price:  askPrice,
			Orders: askOrders,
			Qty:    askQty,
		})
	}

	return ltp, bids, asks, nil
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
