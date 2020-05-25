package option

// Type Option represents an optional value:
// every Option is either Some and contains a value,
// or None, and does not.
// - Initial values
// - Return values for functions that are not defined over their entire input range (partial functions)
// - Return value for otherwise reporting simple errors, where None is returned on error
// - Optional struct fields
// - Struct fields that can be loaned or "taken"
// - Optional function arguments
// - Nullable pointers
// - Swapping things out of difficult situations
//
// Options are commonly paired with pattern matching to query the presence of a value and take action, always accounting for the None case.
type Option struct {
	data interface{}
	v    bool
}

// Some value T
func Some(s interface{}) Option { return Option{data: s, v: true} }

// No value
func None() Option { return Option{} }

func (o Option) IsSome() bool { return o.v }

func (o Option) IsNone() bool { return !o.IsSome() }

// Returns None if the option is None, otherwise returns optb.
func (o Option) And(optb Option) Option {
	if o.IsNone() {
		return None()
	}
	return optb
}

func (o Option) AndThen(fn func(interface{}) Option) Option {
	if o.IsNone() {
		return None()
	}
	return fn(o.data)
}

// Returns the option if it contains a value, otherwise returns optb.
func (o Option) Or(optb Option) Option {
	if o.IsNone() {
		return optb
	}
	return o
}

func (o Option) OrElse(fn func() Option) Option {
	if o.IsNone() {
		return fn()
	}
	return o
}

func (o Option) Map(fn func(interface{}) interface{}) Option {
	if o.IsNone() {
		return None()
	}
	return Some(fn(o.data))
}

// Applies a function to the contained value (if any), or returns the provided default (if not).
func (o Option) MapOr(defValue interface{}, fn func(interface{}) interface{}) Option {
	if o.IsSome() {
		return Some(fn(o.data))
	}
	return Some(defValue)
}

// Returns the contained Some value, consuming the self value.
//
// Because this function may panic, its use is generally discouraged.
// Instead, prefer to use pattern matching and handle the None case explicitly,
// or call `unwrap_or`, `unwrap_or_else`.
func (o Option) Unwrap() interface{} {
	if o.IsNone() {
		panic("unwrap Option None")
	}
	return o.data
}

func (o Option) UnwrapOr(v interface{}) interface{} {
	if o.IsNone() {
		return v
	}
	return o.data
}

func (o Option) UnwrapOrElse(fn func() interface{}) interface{} {
	if o.IsNone() {
		return fn()
	}
	return o.data
}
