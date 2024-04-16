package market

import (
	"cryptoarbitrage/global"
	"encoding/json"
	"io"
	"net/http"
)

const BASEURL = "https://api.mexc.com/api/v3/"

type Mexc struct {
	symbols map[string]global.Symbol
}

func (m *Mexc) PullSymbolsFromAPI() error {
	type symbolsReq struct {
		Data []string `json:"data"`
	}
	requestURL := BASEURL + "defaultSymbols"
	res, err := http.Get(requestURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	// var coins map[string]interface{}
	var coins symbolsReq
	io, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(io), &coins); err == nil {
		// fmt.Println(coins)
		var gSymbols = make(map[string]global.Symbol)
		for _, i := range coins.Data {
			gSymbols[i] = global.Symbol{}
		}
		m.SetSymbols(gSymbols)
		return nil
	}
	return err
}

func (m *Mexc) GetSymbols() map[string]global.Symbol {
	return m.symbols
}
func (m *Mexc) SetSymbols(s map[string]global.Symbol) {
	m.symbols = s
}
func (m *Mexc) GetMarketName() string {
	return "Mexc"
}
