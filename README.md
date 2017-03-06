<img src="goLangLogo.jpeg"  width="380" height="150" border="0" /> 



[![Build Status](https://travis-ci.org/pjgg/Go-Option.svg?branch=master)](https://travis-ci.org/pjgg/Go-Option)
[![codecov](https://codecov.io/gh/pjgg/Go-Option/branch/master/graph/badge.svg)](https://codecov.io/gh/pjgg/Go-Option)


Option go lang implementation

A container struct which may not contain a non-null value. If a value is present, IsPresent() will return true and Get() will return the value.
Additional methods that depend on the presence or absence of a contained value are provided, such as OrElse() (return a default value if value not present).

Methods
=======

empty
-----
Returns an empty Option instance. No value is present for this Option.
**Returns**: 

an empty Option

usage:

`emptyOption := option.empty()`

Of
-----
Returns an Option with the specified present non-null value.

**Returns**: 

an Optional with the value present

usage:

`optionInstance := option.Of("test")`

Get
-----
If a value is present in this Option, returns the value, otherwise return None

**Returns**: 

the non-null value held by this Option

usage:

`optionInstance := option.Of("test")`
`optionValue := optionInstance.Get()`

IsPresent
----------
Return true if there is a value present, otherwise false.

**Returns**: 

true if there is a value present, otherwise false

usage:

`optionInstance := option.Of("test")`
`isPresent := optionInstance.IsPresent()`

OrElse
----------
Return the value if present, otherwise return other.

**Returns**: 

Return the value if present, otherwise return other.

usage:

`emptyOption := option.empty()`
`optionValue := optionInstance.OrElse("myDefaultValue")`

OrElseError
------------
Return the contained value, if present, otherwise returns error

**Returns**: 

Return the contained value, if present, otherwise returns error

usage:

`emptyOption := option.empty()`
`err := optionInstance.OrElseError(errors.New("My custom error"))`

or

`optionInstance := option.Of("test")`
`value := optionInstance.OrElseError(errors.New("My custom error"))`


Filter
-------
Returns this option if present and math with the given predicate, otherwise returns None

There are some predicates available. Have a look file [option/predicates.go](option/predicates.go) but you could develop your own ones. Just follow this signature

`type Predicate func(interface{}) bool
`

**Returns**: 

Returns this option if present and math with the given predicate, otherwise returns None

usage:

`option := Of(1)`
`optionValue := optionInstance.Filter(equalInt(1))`

**Note:** `equalInt` is a predefined predicate created as an example of how to use the type Predicate `type Predicate func(interface{}) bool`. You could have a look how is implemented and develop your own predicates for your custom struts.

FilterNot
----------
Returns this option if present and not math with the given predicate, otherwise returns None

There are some predicates available. Have a look file [option/predicates.go](option/predicates.go) but you could develop your own ones. Just follow this signature

`type Predicate func(interface{}) bool
`

* Input interface represents the option value.
* output bool represents if there are a value that match with the given inputValue

**Returns**: 

Returns this option if present and not math with the given predicate, otherwise returns None

usage:

`option := Of(2)`
`optionValue := optionInstance.FilterNot(equalInt(1))`

**Note:** `equalInt` is a predefined predicate created as an example of how to use the type Predicate `type Predicate func(interface{}) bool`. You could have a look how is implemented and develop your own predicates for your custom struts.

Find
-----
Finds elements of an array satisfying a finder function. If any element match with the given function then returns None.

There are some finders available. Have a look file [option/finders.go](option/finders.go) but you could develop your own ones. Just follow this signature

`type Finder func(interface{}) (bool, interface{})
`

* Input interface represents the option value.
* output bool represents if there are a value that match with the given inputValue
* output interface represents the value or values that match with the inputValue

**Returns**: 

Returns this option or a slice of this option if present and math with the given finder function, otherwise returns None

usage:

`option := Of([]int{2, 3, 5, 7, 11, 13})`
`result := option.Find(FindElemIntArrayGreaterThan(5))`

**Note:** `FindElemIntArrayGreaterThan ` is a predefined finder created as an example of how to use the type finder `type Finder func(interface{}) (bool, interface{})`. You could have a look how is implemented and develop your own finders for your custom struts.

Map
-----
apply a f action function to option value.

There are some actions function available. Have a look file [option/actions.go](option/actions.go) but you could develop your own ones. Just follow this signature

`type Action func(interface{}) (err error, result interface{})
`

* Input interface represents the option value.
* output error represents if there are any error in the mapping action
* output interface represents the option value mapped

**Returns**: 

Returns input interface (option.value) mapped by an action function

usage:

`option := Of([]string{"madrid", "toledo"})`
`option.Map(UpperCaseArrayString())`
`result := option.Get()`

**Note:** `UpperCaseString ` is a predefined action created as an example of how to use the type action `type Action func(interface{}) (err error, result interface{})`. You could have a look how is implemented and develop your own actions for your custom struts.