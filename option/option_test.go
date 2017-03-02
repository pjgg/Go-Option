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

func (suite *OptionTypeSuite) TestOrElseThrow() {
	option := Empty()

	// invoke
	_, getResult := option.OrElseThrow(errors.New("My custom error"))

	// assert
	assert.EqualError(suite.T(), getResult, "My custom error", "unexpected Error msg")
}

func (suite *OptionTypeSuite) TestOrElseThrowWithValue() {
	option := Of("test")

	// invoke
	value, _ := option.OrElseThrow(errors.New("My custom error"))

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
	value := option.Filter(equalString("test"))

	// assert
	assert.EqualValues(suite.T(), "test", value)
}

func (suite *OptionTypeSuite) TestFilterEqualInt() {
	option := Of(1)

	// invoke
	value := option.Filter(equalInt(1))

	// assert
	assert.EqualValues(suite.T(), 1, value)
}

func (suite *OptionTypeSuite) TestFilterNotEqualInt() {
	option := Of(1)

	// invoke
	value := option.FilterNot(equalInt(2))

	// assert
	assert.EqualValues(suite.T(), 1, value)
}

func (suite *OptionTypeSuite) TestFilterNotEqualIntNone() {
	option := Of(1)

	// invoke
	value := option.FilterNot(equalInt(1))

	// assert
	assert.EqualValues(suite.T(), &None{}, value)
}

func (suite *OptionTypeSuite) TestFilterEqualIntWithNoneExistingValue() {
	option := Of(1)

	// invoke
	value := option.Filter(equalInt(2))

	// assert
	assert.EqualValues(suite.T(), &None{}, value)
}

func (suite *OptionTypeSuite) TestFilterGreaterIntThanWithValue() {
	option := Of(1)

	// invoke
	value := option.Filter(greaterIntThan(2))

	// assert
	assert.EqualValues(suite.T(), &None{}, value)
}

func (suite *OptionTypeSuite) TestFilterGreaterIntThanWithValueGreater() {
	option := Of(2)

	// invoke
	value := option.Filter(greaterIntThan(1))

	// assert
	assert.EqualValues(suite.T(), 2, value)
}

func (suite *OptionTypeSuite) TestFilterlesserIntThanWithValue() {
	option := Of(1)

	// invoke
	value := option.Filter(lesserIntThan(2))

	// assert
	assert.EqualValues(suite.T(), 1, value)
}

func (suite *OptionTypeSuite) TestFilterlesserIntThanWithValueGreater() {
	option := Of(2)

	// invoke
	value := option.Filter(lesserIntThan(1))

	// assert
	assert.EqualValues(suite.T(), &None{}, value)
}

func TestTypesOptionSuite(t *testing.T) {
	suite.Run(t, new(OptionTypeSuite))
}

type OptionTypeSuite struct {
	suite.Suite
}
