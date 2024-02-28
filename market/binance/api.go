package market

import (
	"cryptoarbitrage/global"
	"encoding/json"
	"io"
	"net/http"
)

const BASEURL = "https://api.binance.com/api/"

type Binance struct {
	symbols map[string]global.Symbol
}
type symbolReq struct {
	Symbol string `json:"symbol"`
}

func (b *Binance) GetSymbols() (map[string]global.Symbol, error) {
	requestURL := BASEURL + "v3/ticker/price"
	res, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	// var coins []map[string]interface{}
	var coins []symbolReq
	io, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(io), &coins); err == nil {
		var symbols = make(map[string]global.Symbol)
		for _, item := range coins {
			symbol := global.Symbol{}
			symbols[item.Symbol] = symbol
		}
		return symbols, nil
	}

	return nil, err

}

func (b *Binance) GetSymbolsStruct() map[string]global.Symbol {
	return b.symbols
}
func (b *Binance) SetSymbolsStruct(symbols map[string]global.Symbol) {
	b.symbols = symbols
}
func (b *Binance) GetStructName() string {
	return "Binance"
}
