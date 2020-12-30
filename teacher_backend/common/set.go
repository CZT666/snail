package common

type Set struct {
	value map[interface{}]void
}

type void struct{}

func NewSet() *Set {
	set := new(Set)
	set.value = make(map[interface{}]void)
	return set
}

func (set *Set) Add(key interface{}) {
	var value void
	set.value[key] = value
}

func (set *Set) Delete(key interface{}) {
	delete(set.value, key)
}

func (set *Set) Contains(key interface{}) bool {
	_, ok := set.value[key]
	return ok
}

func (set *Set) Length() int {
	return len(set.value)
}

func (set *Set) Value() map[interface{}]void {
	return set.value
}
