package option

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func (suite *OptionTypeSuite) TestGet() {

	option := Empty()

	// invoke
	getResult := option.Get()

	// assert
	assert.EqualValues(suite.T(), &None{}, getResult)
}

func (suite *OptionTypeSuite) TestGetWithValue() {

	option := Of("test")

	// invoke
	getResult := option.Get()

	// assert
	assert.EqualValues(suite.T(), "test", getResult)
}

func (suite *OptionTypeSuite) TestGetWithNilValue() {

	option := Of(nil)

	// invoke
	getResult := option.Get()

	// assert
	assert.EqualValues(suite.T(), &None{}, getResult)
}

func (suite *OptionTypeSuite) TestOrElseWithValue() {

	option := Of("test")

	// invoke
	getResult := option.OrElse("elseTest")

	// assert
	assert.EqualValues(suite.T(), "test", getResult)
}

func (suite *OptionTypeSuite) TestOrElse() {

	option := Empty()

	// invoke
	getResult := option.OrElse("elseTest")

	// assert
	assert.EqualValues(suite.T(), "elseTest", getResult)
}

func (suite *OptionTypeSuite) TestOrElseError() {
	option := Empty()

	// invoke
	_, getResult := option.OrElseError(errors.New("My custom error"))

	// assert
	assert.EqualError(suite.T(), getResult, "My custom error", "unexpected Error msg")
}

func (suite *OptionTypeSuite) TestOrElseErrorWithValue() {
	option := Of("test")

	// invoke
	value, _ := option.OrElseError(errors.New("My custom error"))

	// assert
	assert.EqualValues(suite.T(), "test", value)
}

func (suite *OptionTypeSuite) TestIsPresent() {
	option := Empty()

	// invoke
	value := option.IsPresent()

	// assert
	assert.False(suite.T(), value, "must be false")
}

func (suite *OptionTypeSuite) TestIsPresentWithValue() {
	option := Of("test")

	// invoke
	value := option.IsPresent()

	// assert
	assert.True(suite.T(), value, "must be true")
}

func (suite *OptionTypeSuite) TestFilterEqualString() {
	option := Of("test")

	// invoke
	value := option.Filter(EqualString("test"))

	// assert
	assert.EqualValues(suite.T(), "test", value)
}

func (suite *OptionTypeSuite) TestFilterEqualInt() {
	option := Of(1)

	// invoke
	value := option.Filter(EqualInt(1))

	// assert
	assert.EqualValues(suite.T(), 1, value)
}

func (suite *OptionTypeSuite) TestFilterNotEqualInt() {
	option := Of(1)

	// invoke
	value := option.FilterNot(EqualInt(2))

	// assert
	assert.EqualValues(suite.T(), 1, value)
}

func (suite *OptionTypeSuite) TestFilterNotEqualIntNone() {
	option := Of(1)

	// invoke
	value := option.FilterNot(EqualInt(1))

	// assert
	assert.EqualValues(suite.T(), &None{}, value)
}

func (suite *OptionTypeSuite) TestFilterEqualIntWithNoneExistingValue() {
	option := Of(1)

	// invoke
	value := option.Filter(EqualInt(2))

	// assert
	assert.EqualValues(suite.T(), &None{}, value)
}

func (suite *OptionTypeSuite) TestFilterGreaterIntThanWithValue() {
	option := Of(1)

	// invoke
	value := option.Filter(GreaterIntThan(2))

	// assert
	assert.EqualValues(suite.T(), &None{}, value)
}

func (suite *OptionTypeSuite) TestFilterGreaterIntThanWithValueGreater() {
	option := Of(2)

	// invoke
	value := option.Filter(GreaterIntThan(1))

	// assert
	assert.EqualValues(suite.T(), 2, value)
}

func (suite *OptionTypeSuite) TestFilterlesserIntThanWithValue() {
	option := Of(1)

	// invoke
	value := option.Filter(LesserIntThan(2))

	// assert
	assert.EqualValues(suite.T(), 1, value)
}

func (suite *OptionTypeSuite) TestFilterlesserIntThanWithValueGreater() {
	option := Of(2)

	// invoke
	value := option.Filter(LesserIntThan(1))

	// assert
	assert.EqualValues(suite.T(), &None{}, value)
}

func (suite *OptionTypeSuite) TestArrayInt() {
	option := Of([]int{2, 3, 5, 7, 11, 13})

	// invoke
	value := option.Filter(ContainsElemIntArray(2))

	// assert
	assert.EqualValues(suite.T(), []int{2, 3, 5, 7, 11, 13}, value)
}

func (suite *OptionTypeSuite) TestArrayIntFilterNot() {
	option := Of([]int{2, 3, 5, 7, 11, 13})

	// invoke
	value := option.FilterNot(ContainsElemIntArray(2))

	// assert
	assert.EqualValues(suite.T(), &None{}, value)
}

func (suite *OptionTypeSuite) TestArrayIntNotElem() {
	option := Of([]int{2, 3, 5, 7, 11, 13})

	// invoke
	value := option.Filter(ContainsElemIntArray(1))

	// assert
	assert.EqualValues(suite.T(), &None{}, value)
}

func (suite *OptionTypeSuite) TestArrayFloat64NotElem() {
	option := Of([]float64{2.1, 3.1})

	// invoke
	value := option.Filter(ContainsElemFloat64Array(1.1))

	// assert
	assert.EqualValues(suite.T(), &None{}, value)
}

func (suite *OptionTypeSuite) TestArrayFloat64() {
	option := Of([]float64{2.1})

	// invoke
	value := option.Filter(ContainsElemFloat64Array(2.1))

	// assert
	assert.EqualValues(suite.T(), []float64{2.1}, value)
}

func (suite *OptionTypeSuite) TestArrayString() {
	testData := []string{"Ireland", "Spain", "Italy", "France", "Germany"}
	option := Of(testData)

	// invoke
	value := option.Filter(ContainsElemStringArray("France"))

	// assert
	assert.EqualValues(suite.T(), testData, value)
}

func (suite *OptionTypeSuite) TestArrayStringNone() {
	testData := []string{"Ireland", "Spain", "Italy", "France", "Germany"}
	option := Of(testData)

	// invoke
	value := option.Filter(ContainsElemStringArray("Russia"))

	// assert
	assert.EqualValues(suite.T(), &None{}, value)
}

func (suite *OptionTypeSuite) TestArrayBoolTrue() {
	testData := []bool{true, false, false}
	option := Of(testData)

	// invoke
	value := option.Filter(ContainsElemBoolArray(true))

	// assert
	assert.EqualValues(suite.T(), testData, value)
}

func (suite *OptionTypeSuite) TestArrayBooltrueNone() {
	testData := []bool{false, false, false}
	option := Of(testData)

	// invoke
	value := option.Filter(ContainsElemBoolArray(true))

	// assert
	assert.EqualValues(suite.T(), &None{}, value)
}

func (suite *OptionTypeSuite) TestArrayBoolfalseNone() {
	testData := []bool{true, true, true}
	option := Of(testData)

	// invoke
	value := option.Filter(ContainsElemBoolArray(false))

	// assert
	assert.EqualValues(suite.T(), &None{}, value)
}

func (suite *OptionTypeSuite) TestArrayBoolfalse() {
	testData := []bool{true, true, false}
	option := Of(testData)

	// invoke
	value := option.Filter(ContainsElemBoolArray(false))

	// assert
	assert.EqualValues(suite.T(), testData, value)
}

func (suite *OptionTypeSuite) TestBoolfalse() {
	option := Of(false)

	// invoke
	value := option.Filter(isFalse())

	// assert
	assert.EqualValues(suite.T(), false, value)
}

func (suite *OptionTypeSuite) TestBoolfalseNone() {
	option := Of(true)

	// invoke
	value := option.Filter(isFalse())

	// assert
	assert.EqualValues(suite.T(), &None{}, value)
}

func (suite *OptionTypeSuite) TestBooltrue() {
	option := Of(true)

	// invoke
	value := option.Filter(isTrue())

	// assert
	assert.EqualValues(suite.T(), true, value)
}

func (suite *OptionTypeSuite) TestBooltrueNone() {
	option := Of(false)

	// invoke
	value := option.Filter(isTrue())

	// assert
	assert.EqualValues(suite.T(), &None{}, value)
}

func (suite *OptionTypeSuite) TestFindInt() {
	option := Of([]int{2, 3, 5, 7, 11, 13})

	// invoke
	value := option.Find(FindElemIntArray(5))

	// assert
	assert.EqualValues(suite.T(), 5, value)
}

func (suite *OptionTypeSuite) TestFindIntNone() {
	option := Of([]int{2, 3, 5, 7, 11, 13})

	// invoke
	value := option.Find(FindElemIntArray(1))

	// assert
	assert.EqualValues(suite.T(), &None{}, value)
}

func (suite *OptionTypeSuite) TestFindFloat64() {
	option := Of([]float64{2.1, 3.1})

	// invoke
	value := option.Find(FindElemFloat64Array(2.1))

	// assert
	assert.EqualValues(suite.T(), 2.1, value)
}

func (suite *OptionTypeSuite) TestFindFloat64None() {
	option := Of([]float64{2.1, 3.1})

	// invoke
	value := option.Find(FindElemFloat64Array(2.2))

	// assert
	assert.EqualValues(suite.T(), &None{}, value)
}

func (suite *OptionTypeSuite) TestFindStringNone() {
	option := Of([]string{"Dublin", "Madrid"})

	// invoke
	value := option.Find(FindElemStringArray("Madrid"))

	// assert
	assert.EqualValues(suite.T(), "Madrid", value)
}

func (suite *OptionTypeSuite) TestFindIntGreaterThan() {
	option := Of([]int{2, 3, 5, 7, 11, 13})

	// invoke
	value := option.Find(FindElemIntArrayGreaterThan(5))

	// assert
	assert.EqualValues(suite.T(), []int{7, 11, 13}, value)
}

func (suite *OptionTypeSuite) TestFindIntLesserThan() {
	option := Of([]int{2, 3, 5, 7, 11, 13})

	// invoke
	value := option.Find(FindElemIntArrayLesserThan(5))

	// assert
	assert.EqualValues(suite.T(), []int{2, 3}, value)
}

func TestTypesOptionSuite(t *testing.T) {
	suite.Run(t, new(OptionTypeSuite))
}

type OptionTypeSuite struct {
	suite.Suite
}
