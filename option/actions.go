package option

import (
	"strings"
)

func UpperCaseString() Action {
	return func(value interface{}) (err error, result interface{}) {
		if v, ok := value.(string); ok == true {
			result = strings.ToUpper(v)
		}
		return nil, result
	}
}

func UpperCaseArrayString() Action {
	return func(value interface{}) (err error, result interface{}) {
		if arrayA, ok := value.([]string); ok == true {
			for k, v := range arrayA {
				arrayA[k] = strings.ToUpper(v)
			}
			result = arrayA
		}
		return nil, result
	}
}
