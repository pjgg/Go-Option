package option

func equalInt(numA int) Predicate {
	return func(inputNum interface{}) bool {
		var result bool
		if numB, ok := inputNum.(int); ok == true {
			result = numA == numB
		} else {
			result = false
		}
		return result

	}
}

func greaterIntThan(inputNum int) Predicate {
	return func(value interface{}) bool {
		var result bool
		if valueInt, ok := value.(int); ok == true {
			result = valueInt > inputNum
		} else {
			result = false
		}
		return result
	}
}

func lesserIntThan(inputNum int) Predicate {
	return func(value interface{}) bool {
		var result bool
		if valueInt, ok := value.(int); ok == true {
			result = valueInt < inputNum
		} else {
			result = false
		}
		return result
	}
}

func equalString(stringA string) Predicate {
	return func(inputString interface{}) bool {
		var result bool
		if stringB, ok := inputString.(string); ok == true {
			result = stringA == stringB
		} else {
			result = false
		}
		return result

	}
}
