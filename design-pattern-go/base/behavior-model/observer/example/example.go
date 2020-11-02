package main

import "fmt"

type SubjectInterface interface {
	// 添加观察者
	Attach(observerInterface ...ObserverInterface)error
	// 删除观察者
	Detach(observerInterface ObserverInterface)error
	// 通知所有观察者
	Notify()error
}

type Subject struct {
	ObserverList []ObserverInterface
}
func (s *Subject) Attach(observer ...ObserverInterface) (err error){
	s.ObserverList = append(s.ObserverList, observer...)
	return
}
func (s *Subject) Detach(observer ObserverInterface) (err error) {
	for k,v := range s.ObserverList {
		if v == observer{
			s.ObserverList = append(s.ObserverList[:k], s.ObserverList[k+1:]...)
			return
		}
	}
	return
}
func (s *Subject) Notify() (err error) {
	for _,v := range s.ObserverList{
		_ = v.Do()
	}
	return
}

type ObserverInterface interface {
	Do()error
}

type Observer1 struct {}
func (o Observer1) Do() (err error) {
	fmt.Println("Observer1 do...")
	return 
}

type Observer2 struct {}
func (o Observer2) Do() (err error) {
	fmt.Println("Observer2 do...")
	return
}

type Observer3 struct {}
func (o Observer3) Do() (err error) {
	fmt.Println("Observer3 do...")
	return
}

func main(){
	sub := new(Subject)

	ob1 := new(Observer1)
	ob2 := new(Observer2)
	ob3 := new(Observer3)

	_ = sub.Attach(ob1, ob2, ob3)
	_ = sub.Detach(ob2)

	_ = sub.Notify()
}