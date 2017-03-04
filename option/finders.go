package option

import "sort"

func FindElemIntArrayGreaterThan(inputNum int) Finder {
	return func(value interface{}) (result bool, elem interface{}) {
		if arrayB, ok := value.([]int); ok == true {
			result, elem = findIntGreaterThan(arrayB, inputNum)
		}
		return result, elem
	}
}

func FindElemIntArrayLesserThan(inputNum int) Finder {
	return func(value interface{}) (result bool, elem interface{}) {
		if arrayB, ok := value.([]int); ok == true {
			result, elem = findIntLesserThan(arrayB, inputNum)
		}
		return result, elem
	}
}

func FindElemIntArray(inputNum int) Finder {
	return func(value interface{}) (result bool, elem interface{}) {
		if arrayB, ok := value.([]int); ok == true {
			result, elem = findInt(arrayB, inputNum)
		}
		return result, elem
	}
}

func FindElemFloat64Array(inputNum float64) Finder {
	return func(value interface{}) (result bool, elem interface{}) {
		if arrayB, ok := value.([]float64); ok == true {
			result, elem = findFloat64(arrayB, inputNum)
		}
		return result, elem
	}
}

func FindElemStringArray(inputNum string) Finder {
	return func(value interface{}) (result bool, elem interface{}) {
		if arrayB, ok := value.([]string); ok == true {
			result, elem = findString(arrayB, inputNum)
		}
		return result, elem
	}
}

func findInt(value []int, inputNum int) (result bool, elem interface{}) {
	if !sort.IntsAreSorted(value) {
		sort.Ints(value)
	}

	i := sort.Search(len(value), func(i int) bool { return value[i] >= inputNum })
	if i >= len(value) || value[i] != inputNum {
		result = false
		elem = &None{}
	} else {
		result = true
		elem = value[i]
	}

	return result, elem
}

func findIntGreaterThan(value []int, inputNum int) (bool, interface{}) {
	if !sort.IntsAreSorted(value) {
		sort.Ints(value)
	}

	i := sort.Search(len(value), func(i int) bool { return value[i] > inputNum })
	if i >= len(value) {
		return false, &None{}
	} else {
		elem := make([]int, len(value)-i)
		copy(elem, value[i:len(value)])
		return true, elem
	}
}

func findIntLesserThan(value []int, inputNum int) (bool, interface{}) {
	if !sort.IntsAreSorted(value) {
		sort.Ints(value)
	}

	i := sort.Search(len(value), func(i int) bool { return value[i] >= inputNum })
	if i >= len(value) {
		return false, &None{}
	} else {
		elem := make([]int, i)
		copy(elem, value[0:i])
		return true, elem
	}
}

func findFloat64(value []float64, inputNum float64) (result bool, elem interface{}) {
	if !sort.Float64sAreSorted(value) {
		sort.Float64s(value)
	}

	i := sort.Search(len(value), func(i int) bool { return value[i] >= inputNum })
	if i >= len(value) || value[i] != inputNum {
		result = false
		elem = &None{}
	} else {
		result = true
		elem = value[i]
	}

	return result, elem
}

func findString(value []string, inputNum string) (result bool, elem interface{}) {
	if !sort.StringsAreSorted(value) {
		sort.Strings(value)
	}

	i := sort.Search(len(value), func(i int) bool { return value[i] >= inputNum })
	if i >= len(value) || value[i] != inputNum {
		result = false
		elem = &None{}
	} else {
		result = true
		elem = value[i]
	}

	return result, elem
}
