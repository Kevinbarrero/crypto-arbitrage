package global

type Market interface {
	PullSymbolsFromAPI() error
	GetSymbols() map[string]Symbol
	GetMarketName() string
	SetSymbols(map[string]Symbol)
}

type Symbol struct {
	Price string
}
