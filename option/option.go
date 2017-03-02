package option

type None struct {
}

type Option struct {
	value interface{}
}

type OptionBehavior interface {
	OrElse(defaultValue interface{}) interface{}
	OrElseError(err error) (interface{}, error)
	IsPresent() bool
	IfPresent(action PlayAction) interface{}
	Get() interface{}
	Filter(predicate Predicate) interface{}
	FilterNot(predicate Predicate) interface{}
}

// Returns the option's value.
func (o *Option) Get() interface{} {
	if o.value != nil {
		return o.value
	}

	return &None{}
}

//if value present then make an action
func (o *Option) IfPresent(action PlayAction) interface{} {
	if o.value != nil {
		return action()
	}

	return &None{}
}

// Returns the option's value if match predicate conditions
func (o *Option) Filter(f Predicate) interface{} {
	if f(o.value) {
		return o.value
	}

	return &None{}
}

// Returns the option's value if not match predicate conditions
func (o *Option) FilterNot(f Predicate) interface{} {
	if !f(o.value) {
		return o.value
	}

	return &None{}
}

func Empty() *Option {
	return &Option{}
}

//We can use the method orElse()
func (o *Option) OrElse(defaultValue interface{}) interface{} {
	if o.value == nil {
		return defaultValue
	}
	return o.value
}

//We can indicate the Optional to throw an error in case its value is null:
func (o *Option) OrElseError(err error) (interface{}, error) {
	if o.value == nil {
		return nil, err
	}
	return o.value, err
}

//There is the possiblity to check directly if the value is initialized and not null
func (o *Option) IsPresent() (result bool) {
	result = true
	if o.value == nil {
		result = false
	}

	return
}

func Of(value interface{}) *Option {
	var option = new(Option)
	if value == nil {
		option.value = &None{}
	} else {
		option.value = value
	}

	return option
}

type PlayAction func() error

type Predicate func(interface{}) bool
