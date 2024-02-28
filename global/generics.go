package global

type Market interface {
	GetSymbols() (map[string]Symbol, error)
	GetSymbolsStruct() map[string]Symbol
	SetSymbolsStruct(map[string]Symbol)
	GetStructName() string
}

type Symbol struct {
	Price string
}
