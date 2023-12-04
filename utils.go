package aoc2023

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
