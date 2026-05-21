package utils

import (
	"database/sql"
	"math"

	"github.com/Sahil2004/gmarket/server/database"
	"github.com/Sahil2004/gmarket/server/dtos"
	"github.com/Sahil2004/gmarket/server/models"
	"github.com/google/uuid"
)

func DefaultAlgoConfig(symbol, exchange string) dtos.AlgoConfigDTO {
	return dtos.AlgoConfigDTO{
		Symbol:        symbol,
		Exchange:      exchange,
		Enabled:       false,
		RSIEnabled:    true,
		RSIPeriod:     14,
		RSIOverbought: 70,
		RSIOversold:   30,
		MAEnabled:     true,
		MAFastPeriod:  9,
		MASlowPeriod:  21,
	}
}

func AlgoConfigToDTO(cfg models.AlgoConfig) dtos.AlgoConfigDTO {
	return dtos.AlgoConfigDTO{
		Symbol:        cfg.Symbol,
		Exchange:      cfg.Exchange,
		Enabled:       cfg.Enabled,
		RSIEnabled:    cfg.RSIEnabled,
		RSIPeriod:     cfg.RSIPeriod,
		RSIOverbought: cfg.RSIOverbought,
		RSIOversold:   cfg.RSIOversold,
		MAEnabled:     cfg.MAEnabled,
		MAFastPeriod:  cfg.MAFastPeriod,
		MASlowPeriod:  cfg.MASlowPeriod,
	}
}

func GetAlgoConfig(queries *database.Queries, userID, symbol, exchange string) (dtos.AlgoConfigDTO, error) {
	cfg, err := queries.GetAlgoConfig(userID, symbol, exchange)
	if err == sql.ErrNoRows {
		return DefaultAlgoConfig(symbol, exchange), nil
	}
	if err != nil {
		return dtos.AlgoConfigDTO{}, err
	}
	return AlgoConfigToDTO(cfg), nil
}

func SaveAlgoConfig(queries *database.Queries, userID string, dto dtos.AlgoConfigDTO) (dtos.AlgoConfigDTO, error) {
	if dto.MAFastPeriod >= dto.MASlowPeriod {
		return dtos.AlgoConfigDTO{}, errInvalidMAPeriods()
	}
	uid, err := uuid.Parse(userID)
	if err != nil {
		return dtos.AlgoConfigDTO{}, err
	}
	cfg := models.AlgoConfig{
		UserID:         uid,
		Symbol:         dto.Symbol,
		Exchange:       dto.Exchange,
		Enabled:        dto.Enabled,
		RSIEnabled:     dto.RSIEnabled,
		RSIPeriod:      dto.RSIPeriod,
		RSIOverbought:  dto.RSIOverbought,
		RSIOversold:    dto.RSIOversold,
		MAEnabled:      dto.MAEnabled,
		MAFastPeriod:   dto.MAFastPeriod,
		MASlowPeriod:   dto.MASlowPeriod,
	}
	saved, err := queries.UpsertAlgoConfig(cfg)
	if err != nil {
		return dtos.AlgoConfigDTO{}, err
	}
	return AlgoConfigToDTO(saved), nil
}

type algoError string

func (e algoError) Error() string { return string(e) }

func errInvalidMAPeriods() error {
	return algoError("fast MA period must be less than slow MA period")
}

func ComputeSMA(closes []float64, period int) []float64 {
	if period <= 0 || len(closes) < period {
		return []float64{}
	}
	out := make([]float64, len(closes))
	for i := range closes {
		if i < period-1 {
			out[i] = math.NaN()
			continue
		}
		sum := 0.0
		for j := i - period + 1; j <= i; j++ {
			sum += closes[j]
		}
		out[i] = sum / float64(period)
	}
	return out
}

func ComputeRSI(closes []float64, period int) []float64 {
	if period <= 0 || len(closes) < period+1 {
		return []float64{}
	}
	out := make([]float64, len(closes))
	for i := 0; i < len(closes); i++ {
		out[i] = math.NaN()
	}
	avgGain := 0.0
	avgLoss := 0.0
	for i := 1; i <= period; i++ {
		change := closes[i] - closes[i-1]
		if change > 0 {
			avgGain += change
		} else {
			avgLoss -= change
		}
	}
	avgGain /= float64(period)
	avgLoss /= float64(period)

	if avgLoss == 0 {
		out[period] = 100
	} else {
		rs := avgGain / avgLoss
		out[period] = 100 - (100 / (1 + rs))
	}

	for i := period + 1; i < len(closes); i++ {
		change := closes[i] - closes[i-1]
		gain, loss := 0.0, 0.0
		if change > 0 {
			gain = change
		} else {
			loss = -change
		}
		avgGain = (avgGain*float64(period-1) + gain) / float64(period)
		avgLoss = (avgLoss*float64(period-1) + loss) / float64(period)
		if avgLoss == 0 {
			out[i] = 100
		} else {
			rs := avgGain / avgLoss
			out[i] = 100 - (100 / (1 + rs))
		}
	}
	return out
}

func lastValid(values []float64) float64 {
	for i := len(values) - 1; i >= 0; i-- {
		if !math.IsNaN(values[i]) {
			return values[i]
		}
	}
	return 0
}

func RSISignal(rsi float64, cfg dtos.AlgoConfigDTO) string {
	if !cfg.RSIEnabled || rsi == 0 {
		return "hold"
	}
	if rsi <= cfg.RSIOversold {
		return "buy"
	}
	if rsi >= cfg.RSIOverbought {
		return "sell"
	}
	return "hold"
}

func MASignal(fast, slow []float64) string {
	if len(fast) < 2 || len(slow) < 2 {
		return "hold"
	}
	prevFast, prevSlow := fast[len(fast)-2], slow[len(slow)-2]
	curFast, curSlow := fast[len(fast)-1], slow[len(slow)-1]
	if math.IsNaN(prevFast) || math.IsNaN(curFast) || math.IsNaN(prevSlow) || math.IsNaN(curSlow) {
		return "hold"
	}
	if prevFast <= prevSlow && curFast > curSlow {
		return "buy"
	}
	if prevFast >= prevSlow && curFast < curSlow {
		return "sell"
	}
	if curFast > curSlow {
		return "bullish"
	}
	if curFast < curSlow {
		return "bearish"
	}
	return "hold"
}

func CombineSignals(rsiSig, maSig string, cfg dtos.AlgoConfigDTO) string {
	if !cfg.Enabled {
		return "disabled"
	}
	buy, sell := 0, 0
	if cfg.RSIEnabled {
		switch rsiSig {
		case "buy":
			buy++
		case "sell":
			sell++
		}
	}
	if cfg.MAEnabled {
		switch maSig {
		case "buy":
			buy++
		case "sell":
			sell++
		case "bullish":
			buy++
		case "bearish":
			sell++
		}
	}
	if buy > sell && buy > 0 {
		return "buy"
	}
	if sell > buy && sell > 0 {
		return "sell"
	}
	return "hold"
}

func BuildAlgoIndicators(symbol, exchange, rangeKey string, cfg dtos.AlgoConfigDTO) (dtos.AlgoIndicatorsDTO, error) {
	candlesDTO, err := GetCandles(symbol, exchange, rangeKey)
	if err != nil {
		return dtos.AlgoIndicatorsDTO{}, err
	}

	times := make([]int64, len(candlesDTO.Candles))
	closes := make([]float64, len(candlesDTO.Candles))
	for i, c := range candlesDTO.Candles {
		times[i] = c.Time
		closes[i] = c.Close
	}

	rsiVals := ComputeRSI(closes, cfg.RSIPeriod)
	fastVals := ComputeSMA(closes, cfg.MAFastPeriod)
	slowVals := ComputeSMA(closes, cfg.MASlowPeriod)

	rsiSeries := make([]dtos.IndicatorPointDTO, 0)
	fastSeries := make([]dtos.IndicatorPointDTO, 0)
	slowSeries := make([]dtos.IndicatorPointDTO, 0)

	for i := range times {
		if !math.IsNaN(rsiVals[i]) {
			rsiSeries = append(rsiSeries, dtos.IndicatorPointDTO{Time: times[i], Value: RoundMoney(rsiVals[i])})
		}
		if !math.IsNaN(fastVals[i]) {
			fastSeries = append(fastSeries, dtos.IndicatorPointDTO{Time: times[i], Value: RoundMoney(fastVals[i])})
		}
		if !math.IsNaN(slowVals[i]) {
			slowSeries = append(slowSeries, dtos.IndicatorPointDTO{Time: times[i], Value: RoundMoney(slowVals[i])})
		}
	}

	curRSI := lastValid(rsiVals)
	rsiSig := RSISignal(curRSI, cfg)
	maSig := MASignal(fastVals, slowVals)

	return dtos.AlgoIndicatorsDTO{
		Symbol:         symbol,
		Exchange:       exchange,
		Range:          rangeKey,
		Config:         cfg,
		RSI:            rsiSeries,
		MAFast:         fastSeries,
		MASlow:         slowSeries,
		CurrentRSI:     RoundMoney(curRSI),
		CurrentMAFast:  RoundMoney(lastValid(fastVals)),
		CurrentMASlow:  RoundMoney(lastValid(slowVals)),
		RSISignal:      rsiSig,
		MASignal:       maSig,
		CombinedSignal: CombineSignals(rsiSig, maSig, cfg),
	}, nil
}

func GetAlgoIndicators(queries *database.Queries, userID, symbol, exchange, rangeKey string) (dtos.AlgoIndicatorsDTO, error) {
	cfg, err := GetAlgoConfig(queries, userID, symbol, exchange)
	if err != nil {
		return dtos.AlgoIndicatorsDTO{}, err
	}
	return BuildAlgoIndicators(symbol, exchange, rangeKey, cfg)
}
