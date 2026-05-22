package utils

import (
	"errors"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"sync"

	"github.com/Sahil2004/gmarket/server/dtos"
)

type rangeConfig struct {
	barCount int
	interval time.Duration
}

var candleRangeConfigs = map[string]rangeConfig{
	"1D": {barCount: 78, interval: 5 * time.Minute},
	"1W": {barCount: 65, interval: 30 * time.Minute},
	"1M": {barCount: 22, interval: 24 * time.Hour},
	"1Y": {barCount: 252, interval: 24 * time.Hour},
}

type candleSeriesState struct {
	candles  []dtos.CandleDTO
	interval time.Duration
	barCount int
}

var candleStore = struct {
	sync.Mutex
	m map[string]*candleSeriesState
}{
	m: make(map[string]*candleSeriesState),
}

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

func roundPrice(value float64) float64 {
	return math.Round(value*100) / 100
}

func alignCandleTime(t time.Time, interval time.Duration) time.Time {
	if interval >= 24*time.Hour {
		return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	}
	intervalSec := int64(interval.Seconds())
	unix := t.Unix()
	return time.Unix((unix/intervalSec)*intervalSec, 0)
}

func generateCandleHistory(endPrice float64, count int, interval time.Duration, end time.Time) []dtos.CandleDTO {
	candles := make([]dtos.CandleDTO, count)
	endAligned := alignCandleTime(end, interval)
	price := endPrice
	r := rand.New(rand.NewSource(end.UnixNano()))

	for i := count - 1; i >= 0; i-- {
		barTime := endAligned.Add(-time.Duration(count-1-i) * interval)
		open := price
		change := (r.Float64() - 0.5) * open * 0.012
		close := open + change
		wick := open * 0.006 * r.Float64()
		high := math.Max(open, close) + wick
		low := math.Min(open, close) - wick*0.8

		candles[i] = dtos.CandleDTO{
			Time:  barTime.Unix(),
			Open:  roundPrice(open),
			High:  roundPrice(high),
			Low:   roundPrice(low),
			Close: roundPrice(close),
		}
		price = close
	}

	last := &candles[count-1]
	last.Close = roundPrice(endPrice)
	if endPrice > last.High {
		last.High = roundPrice(endPrice)
	}
	if endPrice < last.Low {
		last.Low = roundPrice(endPrice)
	}

	return candles
}

func GetCandles(symbol string, exchange string, rangeKey string) (dtos.CandlesDTO, error) {
	cfg, ok := candleRangeConfigs[rangeKey]
	if !ok {
		return dtos.CandlesDTO{}, errors.New("invalid range")
	}

	ltp, _, _, _, err := GetMarketData(symbol, exchange)
	if err != nil {
		return dtos.CandlesDTO{}, err
	}

	key := GenerateDBStockSymbol(symbol, exchange) + ":" + rangeKey
	now := time.Now()
	currentBarTime := alignCandleTime(now, cfg.interval).Unix()

	candleStore.Lock()
	defer candleStore.Unlock()

	series, exists := candleStore.m[key]
	if !exists {
		candles := generateCandleHistory(ltp, cfg.barCount, cfg.interval, now)
		candleStore.m[key] = &candleSeriesState{
			candles:  candles,
			interval: cfg.interval,
			barCount: cfg.barCount,
		}
		return dtos.CandlesDTO{
			Symbol:   symbol,
			Exchange: exchange,
			Range:    rangeKey,
			Candles:  candles,
		}, nil
	}

	lastBar := &series.candles[len(series.candles)-1]
	if currentBarTime > lastBar.Time {
		series.candles = append(series.candles, dtos.CandleDTO{
			Time:  currentBarTime,
			Open:  roundPrice(ltp),
			High:  roundPrice(ltp),
			Low:   roundPrice(ltp),
			Close: roundPrice(ltp),
		})
		if len(series.candles) > series.barCount {
			series.candles = series.candles[len(series.candles)-series.barCount:]
		}
	} else {
		lastBar.Close = roundPrice(ltp)
		if ltp > lastBar.High {
			lastBar.High = roundPrice(ltp)
		}
		if ltp < lastBar.Low {
			lastBar.Low = roundPrice(ltp)
		}
	}

	return dtos.CandlesDTO{
		Symbol:   symbol,
		Exchange: exchange,
		Range:    rangeKey,
		Candles:  series.candles,
	}, nil
}
