package utils

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"sync"

	"github.com/Sahil2004/gmarket/server/dtos"
)

var priceStore = struct {
	sync.Mutex
	m map[string]struct {
		ltp       float64
		lcp       float64
		bids      []dtos.LevelDTO
		asks      []dtos.LevelDTO
		updatedAt time.Time
	}
}{
	m: make(map[string]struct {
		ltp       float64
		lcp       float64
		bids      []dtos.LevelDTO
		asks      []dtos.LevelDTO
		updatedAt time.Time
	}),
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

func GetMarketData(symbol string, exchange string) (float64, float64, []dtos.LevelDTO, []dtos.LevelDTO, error) {
	key := GenerateDBStockSymbol(symbol, exchange)
	now := time.Now()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	priceStore.Lock()
	defer priceStore.Unlock()

	// Return cached data if last update was within 500ms
	if data, exists := priceStore.m[key]; exists {
		if now.Sub(data.updatedAt) < 500*time.Millisecond {
			return data.ltp, data.lcp, data.bids, data.asks, nil
		}
	}

	// Generate new market data
	var ltp float64
	var lcp float64
	if data, exists := priceStore.m[key]; exists {
		ltp = data.ltp
		lcp = data.lcp
	} else {
		ltp = float64(1000 + r.Intn(2000))
		lcp = ltp - (float64(r.Intn(200)) - 100)
	}

	delta := (r.Float64() - 0.5) * 2 * (ltp * 0.002) // +/-0.2%
	ltp += delta
	ltp = math.Round(ltp*100) / 100

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

	priceStore.m[key] = struct {
		ltp       float64
		lcp       float64
		bids      []dtos.LevelDTO
		asks      []dtos.LevelDTO
		updatedAt time.Time
	}{
		ltp:       ltp,
		lcp:       lcp,
		bids:      bids,
		asks:      asks,
		updatedAt: now,
	}

	return ltp, lcp, bids, asks, nil
}

func GetSymbolsStatus(symbols []dtos.SymbolDTO) (dtos.SymbolStatusDTO, error) {
	statuses := make([]dtos.SymbolWithStatusDTO, 0, len(symbols))

	for _, sym := range symbols {
		ltp, lcp, _, _, err := GetMarketData(sym.Symbol, sym.Exchange)
		if err != nil {
			return dtos.SymbolStatusDTO{}, err
		}

		statuses = append(statuses, dtos.SymbolWithStatusDTO{
			Symbol:         sym.Symbol,
			Exchange:       sym.Exchange,
			LTP:            strconv.FormatFloat(ltp, 'f', 2, 64),
			LastClosePrice: strconv.FormatFloat(lcp, 'f', 2, 64),
		})
	}

	return dtos.SymbolStatusDTO{
		Symbols: statuses,
	}, nil
}

func GenerateDBStockSymbol(symbol string, exchange string) string {
	return symbol + "." + exchange
}

func ParseDBStockSymbol(dbSymbol string) (string, string) {
	parts := strings.SplitN(dbSymbol, ".", 2)
	return parts[0], parts[1]
}
