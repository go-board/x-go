package result

type Result struct {
	ok  interface{}
	err error
}

func Ok(v interface{}) Result {
	return Result{ok: v}
}

func Err(err error) Result {
	return Result{err: err}
}

func (r Result) IsOk() bool {
	if r.err != nil {
		return false
	}
	return true
}

func (r Result) IsErr() bool { return !r.IsOk() }

func (r Result) Map(fn func(interface{}) interface{}) Result {
	if r.IsOk() {
		return Result{ok: fn(r.ok)}
	}
	return Result{err: r.err}
}

func (r Result) MapErr(fn func(err error) error) Result {
	if r.IsOk() {
		return Result{ok: r.ok}
	}
	return Result{err: fn(r.err)}
}

func (r Result) And(res Result) Result {
	if r.IsOk() {
		return res
	}
	return Result{err: r.err}
}

func (r Result) AndThen(fn func(interface{}) Result) Result {
	if r.IsOk() {
		return fn(r.ok)
	}
	return Result{err: r.err}
}

func (r Result) Or(res Result) Result {
	if r.IsOk() {
		return Result{ok: r.ok}
	}
	return res
}

func (r Result) OrElse(fn func(err error) Result) Result {
	if r.IsOk() {
		return Result{ok: r.ok}
	}
	return fn(r.err)
}

func (r Result) Unwrap() interface{} {
	if r.IsOk() {
		return r.ok
	}
	panic("err: try unwrap Result(err) into ok")
}

func (r Result) UnwrapOr(v interface{}) interface{} {
	if r.IsOk() {
		return r.ok
	}
	return v
}

func (r Result) UnwrapOrElse(fn func(err error) interface{}) interface{} {
	if r.IsOk() {
		return r.ok
	}
	return fn(r.err)
}
