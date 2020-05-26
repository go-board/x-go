package result

// Go a function with recover and do nothing.
func Go(fn func()) {
	SafeGo(fn, nil)
}

// SafeGo a function with recover and do callback.
func SafeGo(fn func(), handle func(x interface{})) {
	defer func() {
		if x := recover(); x != nil {
			if handle != nil {
				handle(x)
			}
		}
	}()
	fn()
}
