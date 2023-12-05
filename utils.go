package aoc2023

import "strconv"

type addable interface {
	~int8 | ~uint8 | ~int16 | ~uint16 | ~int32 | ~uint32 | ~int64 | ~uint64 | ~int | ~uint
}

func SumSlice[S ~[]V, V addable](s S) (sum V) {
	for _, v := range s {
		sum += v
	}
	return sum
}

func SumMapVal[M ~map[K]V, K comparable, V addable](m M) (sum V) {
	for _, v := range m {
		sum += v
	}
	return sum
}

func ProdMapVal[M ~map[K]V, K comparable, V addable](m M) (prod V) {
	prod = 1
	for _, v := range m {
		prod *= v
	}
	return prod
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func IntFieldsIter(s string, fn func(i int)) {
	i := 0
	for {
		for ; i < len(s) && s[i] == ' '; i++ {
		}
		beg := i
		var v int
		for ; i < len(s) && s[i] != ' '; i++ {
			v = v*10 + int(s[i]-'0')
		}
		if beg == i {
			return
		}
		fn(v)
	}
}
