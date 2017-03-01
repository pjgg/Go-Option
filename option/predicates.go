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

func greaterIntThan(numA int) Predicate {
	return func(inputNum interface{}) bool {
		var result bool
		if numB, ok := inputNum.(int); ok == true {
			result = numA > numB
		} else {
			result = false
		}
		return result
	}
}

func lesserIntThan(numA int) Predicate {
	return func(inputNum interface{}) bool {
		var result bool
		if numB, ok := inputNum.(int); ok == true {
			result = numA < numB
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
