package utils

import (
	"math/big"
)

const base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Encode(n uint64) string {
	if n == 0 {
		return string(base62Chars[0])
	}

	res := ""
	base := big.NewInt(62)
	num := big.NewInt(int64(n))

	for num.Cmp(big.NewInt(0)) > 0 {
		mod := new(big.Int)
		num.DivMod(num, base, mod)
		res = string(base62Chars[mod.Int64()]) + res
	}
	return res
}
