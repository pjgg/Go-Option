# Go-Option
[![Build Status](https://travis-ci.org/pjgg/Go-Option.svg?branch=master)](https://travis-ci.org/pjgg/Go-Option)
[![Coverage Status](https://coveralls.io/repos/github/pjgg/Go-Option/badge.svg?branch=master)](https://coveralls.io/github/pjgg/Go-Option?branch=master)

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

`optionValue := optionInstance.Get()`

IsPresent
----------
Return true if there is a value present, otherwise false.

**Returns**: 

true if there is a value present, otherwise false

usage:

`isPresent := optionInstance.IsPresent()`

OrElse
----------
Return the value if present, otherwise return other.

**Returns**: 

Return the value if present, otherwise return other.

usage:

`optionValue := optionInstance.OrElse("myDefaultValue")`

OrElseThrow
------------
Return the contained value, if present, otherwise returns error

**Returns**: 

Return the contained value, if present, otherwise returns error

usage:

`optionValue := optionInstance.OrElseThrow(errors.New("My custom error"))`

Filter
-------
Returns this option if present and math with the given predicate, otherwise returns None

There are some predicates available, you could find them under file [option/predicates.go](option/predicates.go) but you could develop your own ones. Just follow this signature

`type Predicate func(interface{}) bool
`

**Returns**: 

Returns this option if present and math with the given predicate, otherwise returns None

usage:

`optionValue := optionInstance.Filter(equalInt(1))`