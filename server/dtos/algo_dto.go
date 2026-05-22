package dtos

type AlgoConfigDTO struct {
	Symbol        string  `json:"symbol"`
	Exchange      string  `json:"exchange"`
	Enabled       bool    `json:"enabled"`
	RSIEnabled    bool    `json:"rsi_enabled"`
	RSIPeriod     int     `json:"rsi_period"`
	RSIOverbought float64 `json:"rsi_overbought"`
	RSIOversold   float64 `json:"rsi_oversold"`
	MAEnabled     bool    `json:"ma_enabled"`
	MAFastPeriod  int     `json:"ma_fast_period"`
	MASlowPeriod  int     `json:"ma_slow_period"`
}

type IndicatorPointDTO struct {
	Time  int64   `json:"time"`
	Value float64 `json:"value"`
}

type AlgoIndicatorsDTO struct {
	Symbol     string              `json:"symbol"`
	Exchange   string              `json:"exchange"`
	Range      string              `json:"range"`
	Config     AlgoConfigDTO       `json:"config"`
	RSI        []IndicatorPointDTO `json:"rsi"`
	MAFast     []IndicatorPointDTO `json:"ma_fast"`
	MASlow     []IndicatorPointDTO `json:"ma_slow"`
	CurrentRSI float64             `json:"current_rsi"`
	CurrentMAFast float64        `json:"current_ma_fast"`
	CurrentMASlow float64          `json:"current_ma_slow"`
	RSISignal  string              `json:"rsi_signal"`
	MASignal   string              `json:"ma_signal"`
	CombinedSignal string          `json:"combined_signal"`
}
