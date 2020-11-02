package main

import "fmt"

// 汽车产品类接口
type Vehicle interface {
	setColor() Vehicle
	setSize() Vehicle
	getVehicle() Vehicle
}

// 产品主管（生成器）
type director struct {
	vehicle Vehicle
}
func (d *director) SetVehicle(v Vehicle){
	d.vehicle = v
}

func (d *director) Construct(){
	d.vehicle.setColor().setSize()
}

// 具体的产品
type car struct {
	color string
	size  int
}

func (c *car) setColor() Vehicle {
	c.color = "red"
	return c
}

func (c *car) setSize() Vehicle {
	c.size = 100
	return c
}

func (c *car) getVehicle() Vehicle {
	return c
}

func main(){
	// 产品
	c := &car{}

	// 生成器
	d := director{}
	d.SetVehicle(c)
	d.Construct()

	// 最终产品
	fmt.Println(d.vehicle.getVehicle())
	fmt.Println(c)
}


