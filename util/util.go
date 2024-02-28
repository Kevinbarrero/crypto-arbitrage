package util

import "cryptoarbitrage/global"

func MergeSymbols(m1, m2 map[string]global.Symbol) map[string]global.Symbol {
	uniques := make(map[string]global.Symbol)
	if len(m1) < len(m2) {
		for k, v := range m1 {
			if _, ok := m2[k]; ok {
				uniques[k] = v
			}
		}
	} else {
		for k, v := range m2 {
			if _, ok := m1[k]; ok {
				uniques[k] = v
			}
		}
	}
	return uniques
}
