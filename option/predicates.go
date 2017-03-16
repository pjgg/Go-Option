package option

// Bool Predicates
// isTrue returns option value if option value is true
func IsTrue() Predicate {
	return func(value interface{}) (result bool) {
		if value, ok := value.(bool); ok == true {
			result = value == true
		} else {
			result = false
		}
		return
	}
}

// isFalse returns option value if option value is false
func IsFalse() Predicate {
	return func(value interface{}) (result bool) {
		if value, ok := value.(bool); ok == true {
			result = value == false
		} else {
			result = false
		}
		return

	}
}

// string Predicates
// EqualString returns option value if option value is equal to inputValue
func EqualString(inputValue string) Predicate {
	return func(value interface{}) (result bool) {
		if stringB, ok := value.(string); ok == true {
			result = inputValue == stringB
		} else {
			result = false
		}
		return

	}
}

// int Predicates
// EqualInt returns option value if option value is equal to inputValue
func EqualInt(inputValue int) Predicate {
	return func(value interface{}) (result bool) {
		if numB, ok := value.(int); ok == true {
			result = inputValue == numB
		} else {
			result = false
		}
		return
	}
}

// GreaterIntThan returns option value if option value is greater than inputNum
func GreaterIntThan(inputNum int) Predicate {
	return func(value interface{}) (result bool) {
		if valueInt, ok := value.(int); ok == true {
			result = valueInt > inputNum
		} else {
			result = false
		}
		return
	}
}

// LesserIntThan returns option value if option value is lesser than inputNum
func LesserIntThan(inputNum int) Predicate {
	return func(value interface{}) (result bool) {
		if valueInt, ok := value.(int); ok == true {
			result = valueInt < inputNum
		} else {
			result = false
		}
		return
	}
}

// array Predicates
// ContainsElemIntArray returns option value if inputNum is present
func ContainsElemIntArray(inputNum int) Predicate {
	return func(value interface{}) (result bool) {
		if arrayB, ok := value.([]int); ok == true {
			result, _ = findInt(arrayB, inputNum)
		}
		return
	}
}

// ContainsElemFloat64Array returns option value if inputFloat64 is present
func ContainsElemFloat64Array(inputFloat64 float64) Predicate {
	return func(value interface{}) (result bool) {
		if floatB, ok := value.([]float64); ok == true {
			result, _ = findFloat64(floatB, inputFloat64)
		}
		return
	}
}

// ContainsElemStringArray returns option value if inputString is present
func ContainsElemStringArray(inputString string) Predicate {
	return func(value interface{}) (result bool) {
		if arrayB, ok := value.([]string); ok == true {
			result, _ = findString(arrayB, inputString)
		}
		return
	}
}

// ContainsElemBoolArray returns option value if inputBool is present
func ContainsElemBoolArray(inputBool bool) Predicate {
	return func(value interface{}) bool {
		var result bool
		if arrayB, ok := value.([]bool); ok == true {
			for _, v := range arrayB {
				if v == inputBool {
					result = true
					break
				} else {
					result = false
				}
			}
		}
		return result
	}
}
