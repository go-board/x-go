package result

// Wrap a function with recover and do nothing.
func Wrap(fn func()) {
	defer recover()
	fn()
}

func WrapCallback(fn func(), handle func(x interface{})) {
	defer func() {
		if x := recover(); x != nil {
			handle(x)
		}
	}()
	fn()
}
