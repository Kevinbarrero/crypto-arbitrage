package main

import (
	"cryptoarbitrage/global"
	binance "cryptoarbitrage/markets/binance"
	mex "cryptoarbitrage/markets/mexc"
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
	ws := sync.WaitGroup{}
	start := time.Now()
	for _, market := range markets {
		ws.Add(1)
		go global.PullSymbolsHandler(&ws, market)
	}
	ws.Wait()
	fmt.Println("exec time", time.Since(start))
	symbols := util.MergeSymbols(mexc.GetSymbols(), binance.GetSymbols())
	fmt.Println("Binance", len(binance.GetSymbols()))
	fmt.Println("Mexc", len(mexc.GetSymbols()))
	fmt.Println("symbols", len(symbols))
}
