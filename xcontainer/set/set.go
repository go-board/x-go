package set

type Set interface {
	Put(items ...interface{})
	Remove(items ...interface{})
	Values() []interface{}

	ContainAll(items ...interface{}) bool
	ContainAny(items ...interface{}) bool
	InterSection(o Set) Set
	Union(o Set) Set
	SubSet(o Set) bool
	SuperSet(o Set) bool
	Difference(o Set) Set
}
