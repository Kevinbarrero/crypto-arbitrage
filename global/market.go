package global

import (
	"fmt"
	"sync"
	"time"
)

func PullSymbolsHandler(ws *sync.WaitGroup, market Market) {
	start := time.Now()
	err := market.PullSymbolsFromAPI()
	if err != nil {
		fmt.Println("error", err)
	}
	// symbols = util.MergeSymbols(symbols, i_symbols)
	fmt.Println(market.GetMarketName(), "exec time", time.Since(start))
	defer ws.Done()
}
