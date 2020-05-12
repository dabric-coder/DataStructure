package ArrayList

import "errors"

/*
	迭代器主要是解决访问数组方便的问题
*/
type Iterator interface {
	HasNext() bool  // 是否有下一个
	Next() (interface{}, error)  // 下一个
	Remove()  // 删除
	GetIndex() int  // 得到索引
}

type Iterative interface {
	Iterator() Iterator   // 构造初始化接口
}

type ArraylistIterator struct {
	list *ArrayList   // 数组指针
	currentIndex int  // 当前索引
}

func (list *ArrayList) Iterator() Iterator {
	it := new(ArraylistIterator)   // 构造迭代器
	it.currentIndex = 0
	it.list = list
	return it
}

func (it *ArraylistIterator) HasNext() bool {
	if it.currentIndex < it.list.theSize {
		return true
	}
	return false
}
func (it *ArraylistIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("没有下一个")
	}
	value, err := it.list.Get(it.currentIndex)  // 获取当前数据
	it.currentIndex++
	return value, err
}
func (it *ArraylistIterator) Remove() {
	it.currentIndex--
	it.list.Delete(it.currentIndex)   // 删除一个元素
}
func (it *ArraylistIterator) GetIndex() int {
	return it.currentIndex
}


