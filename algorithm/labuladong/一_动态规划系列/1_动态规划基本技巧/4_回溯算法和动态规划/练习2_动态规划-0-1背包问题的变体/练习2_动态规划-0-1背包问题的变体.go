package main

import (
	"fmt"
	"sync"
)

type Set struct {
	sync.RWMutex
	m map[int]bool
}

// 创建集合对象
func NewSet(values ...int) *Set {
	s := &Set{
		m: make(map[int]bool, len(values)),
	}
	s.Add(values...)
	return s
}

// 添加元素
func (s *Set) Add(values ...int) {
	s.Lock()
	defer s.Unlock()
	for _, v := range values {
		s.m[v] = true
	}
}

// 删除元素
func (s *Set) Delete(values ...int) {
	s.Lock()
	defer s.Unlock()
	for _, v := range values {
		delete(s.m, v)
	}
}

// 判断元素是否存在
func (s *Set) IsExist(values ...int) bool {
	s.RLock()
	defer s.RUnlock()
	for _, v := range values {
		if _, ok := s.m[v]; !ok {
			return false
		}
	}
	return true
}

func main() {
	set := NewSet(1, 2, 3, 4, 5)
	set.Add(6)
	set.Delete(6)
	fmt.Println(set.IsExist(6))
	fmt.Println(set)
}
