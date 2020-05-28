package Heap

type Heap interface {
length() int
isEmpty() bool
clear()
add(data interface{})
get() interface{}
remove() interface{}
replace(data interface{}) interface{}
}
