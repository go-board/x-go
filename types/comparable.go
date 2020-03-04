package types

// Ordering is the result of a comparison between two values.
type Ordering int

const (
	OrderingLess Ordering = iota - 1
	OrderingEqual
	OrderingGreater
)

type Comparable interface{ Compare(o Comparable) Ordering }

func (i Int) Compare(o Comparable) Ordering {
	oo := o.(Int)
	if i < oo {
		return OrderingLess
	} else if i > oo {
		return OrderingGreater
	}
	return OrderingEqual
}

func (i Int64) Compare(o Comparable) Ordering {
	oo := o.(Int64)
	if i < oo {
		return OrderingLess
	} else if i > oo {
		return OrderingGreater
	}
	return OrderingEqual
}

func (i Int32) Compare(o Comparable) Ordering {
	oo := o.(Int32)
	if i < oo {
		return OrderingLess
	} else if i > oo {
		return OrderingGreater
	}
	return OrderingEqual
}

func (i Int16) Compare(o Comparable) Ordering {
	oo := o.(Int16)
	if i < oo {
		return OrderingLess
	} else if i > oo {
		return OrderingGreater
	}
	return OrderingEqual
}

func (i Int8) Compare(o Comparable) Ordering {
	oo := o.(Int8)
	if i < oo {
		return OrderingLess
	} else if i > oo {
		return OrderingGreater
	}
	return OrderingEqual
}

func (i Uint) Compare(o Comparable) Ordering {
	oo := o.(Uint)
	if i < oo {
		return OrderingLess
	} else if i > oo {
		return OrderingGreater
	}
	return OrderingEqual
}

func (i Uint64) Compare(o Comparable) Ordering {
	oo := o.(Uint64)
	if i < oo {
		return OrderingLess
	} else if i > oo {
		return OrderingGreater
	}
	return OrderingEqual
}

func (i Uint32) Compare(o Comparable) Ordering {
	oo := o.(Uint32)
	if i < oo {
		return OrderingLess
	} else if i > oo {
		return OrderingGreater
	}
	return OrderingEqual
}

func (i Uint16) Compare(o Comparable) Ordering {
	oo := o.(Uint16)
	if i < oo {
		return OrderingLess
	} else if i > oo {
		return OrderingGreater
	}
	return OrderingEqual
}

func (i Uint8) Compare(o Comparable) Ordering {
	oo := o.(Uint8)
	if i < oo {
		return OrderingLess
	} else if i > oo {
		return OrderingGreater
	}
	return OrderingEqual
}

func (i String) Compare(o Comparable) Ordering {
	oo := o.(String)
	if i < oo {
		return OrderingLess
	} else if i > oo {
		return OrderingGreater
	}
	return OrderingEqual
}
