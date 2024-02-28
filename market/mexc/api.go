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
type symbolsReq struct {
	Data []string `json:"data"`
}

func (m *Mexc) GetSymbols() (map[string]global.Symbol, error) {
	requestURL := BASEURL + "defaultSymbols"
	res, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	// var coins map[string]interface{}
	var coins symbolsReq
	io, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(io), &coins); err == nil {
		// fmt.Println(coins)
		var gSymbols = make(map[string]global.Symbol)
		for _, i := range coins.Data {
			gSymbols[i] = global.Symbol{}
		}
		return gSymbols, nil
	}
	return nil, err
}

func (m *Mexc) GetSymbolsStruct() map[string]global.Symbol {
	return m.symbols
}
func (m *Mexc) SetSymbolsStruct(s map[string]global.Symbol) {
	m.symbols = s
}
func (m *Mexc) GetStructName() string {
	return "Mexc"
}
