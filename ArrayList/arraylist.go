package ArrayList

import (
	"errors"
	"fmt"
)

type List interface {
	Size() int  // 数组的大小
	Get(index int) (interface{}, error)  // 获取第几个元素
	Set(index int, newval interface{}) error  // 修改指定索引位置的数据
	Insert(index int, newval interface{}) error  // 插入数据
	Append(newval interface{})   // 追加
	Clear()   // 清空
	Delete(index int) error  // 删除
	String() string  // 返回字符串
	Iterator() Iterator
}

type ArrayList struct {
	dataStore []interface{}   // 数组存储
	theSize int  // 数组的大小
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)  // 初始化结构体
	list.dataStore = make([]interface{}, 0, 10)  // 开辟空间10个
	list.theSize = 0
	return list
}

func (list *ArrayList) Size() int {
	return list.theSize   // 返回数组大小
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.theSize {
		return nil, errors.New("索引越界")
	}
	return list.dataStore[index], nil
}

func (list *ArrayList) Append(newval interface{}) {
	list.dataStore = append(list.dataStore, newval)  // 叠加数据
	list.theSize++
}

func (list *ArrayList) String() string {  // 返回数组的字符串
	return fmt.Sprint(list.dataStore)
}

func (list *ArrayList) Set(index int, newval interface{}) error {
	if index < 0 || index >= list.theSize {
		return errors.New("索引")
	}

	list.dataStore[index] = newval
	return nil
}

func (list *ArrayList) checkisFull() {
	if list.theSize == cap(list.dataStore) {
		// make 中间的参数为0，表示没有开辟空间
		newdataStore := make([]interface{}, 2*list.theSize, 2*list.theSize)
		//copy(newdataStore, list.dataStore)
		for i := 0; i < list.theSize; i++ {
			newdataStore[i] = list.dataStore[i]
		}
		list.dataStore = newdataStore
	}
}

func (list *ArrayList) Insert(index int, newval interface{}) error {
	if index < 0 || index >= list.theSize {
		return errors.New("索引越界")
	}
	list.checkisFull()  // 检测内存，如果满了，自动增加
	list.dataStore = list.dataStore[:list.theSize+1]  // 插入数据，内存移动一位
	for i := list.theSize; i > index; i-- {  // 从后往前移动
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = newval
	list.theSize++
	return nil
}

func (list *ArrayList) Clear() {
	/*
		go语言有自己的垃圾回收机制，所以只需覆盖掉之前开辟的空间，
		重新开辟一个空间
	*/
	list.dataStore = make([]interface{}, 0, 10)
	list.theSize = 0
}

func (list *ArrayList) Delete(index int) error {
	if index < 0 || index >= list.theSize {
		return errors.New("索引越界")
	}

	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...)  // 重新叠加，跳过index索引
	list.theSize--
	return nil
}

