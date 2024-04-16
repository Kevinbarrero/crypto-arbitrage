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

func (b *Binance) PullSymbolsFromAPI() error {
	type symbolReq struct {
		Symbol string `json:"symbol"`
	}
	requestURL := BASEURL + "v3/ticker/price"
	res, err := http.Get(requestURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	// var coins []map[string]interface{}
	var coins []symbolReq
	io, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(io), &coins); err == nil {
		var symbols = make(map[string]global.Symbol)
		for _, item := range coins {
			symbol := global.Symbol{}
			symbols[item.Symbol] = symbol
		}
		b.SetSymbols(symbols)
		return nil
	}

	return err

}

func (b *Binance) GetSymbols() map[string]global.Symbol {
	return b.symbols
}
func (b *Binance) SetSymbols(symbols map[string]global.Symbol) {
	b.symbols = symbols
}
func (b *Binance) GetMarketName() string {
	return "Binance"
}
