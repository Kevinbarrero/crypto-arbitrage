package main

import (
	"cryptoarbitrage/global"
	binance "cryptoarbitrage/market/binance"
	mex "cryptoarbitrage/market/mexc"
	"cryptoarbitrage/util"
	"fmt"
	"sync"
	"time"
)

func main() {
	var markets []global.Market
	binance := binance.Binance{}
	mexc := mex.Mexc{}
	markets = append(markets, &binance, &mexc)
	// symbols := make(map[string]global.Symbol)
	ws := sync.WaitGroup{}
	ws.Add(len(markets))
	start := time.Now()
	for _, i := range markets {
		go func(i global.Market) {
			start := time.Now()
			symbols, err := i.GetSymbols()
			if err != nil {
				fmt.Println("error", err)
			}
			// symbols = util.MergeSymbols(symbols, i_symbols)
			i.SetSymbolsStruct(symbols)
			ws.Done()
			fmt.Println(i.GetStructName(), "exec time", time.Since(start))
		}(i)
	}
	ws.Wait()
	fmt.Println("exec time", time.Since(start))
	symbols := util.MergeSymbols(mexc.GetSymbolsStruct(), binance.GetSymbolsStruct())
	fmt.Println("Binance", len(binance.GetSymbolsStruct()))
	fmt.Println("Mexc", len(mexc.GetSymbolsStruct()))
	fmt.Println("symbols", len(symbols))
}
