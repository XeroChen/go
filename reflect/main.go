package main

import (
	"fmt"
	"reflect"
)

type BaseClassItf interface {
	Print()
	Id() int
	SetId(int)
}

type BaseClass struct {
	BaseClassItf
	ID   int
	Info string
}

type ClassA struct {
	BaseClass
	A string
}

func (a ClassA) Print() {
	fmt.Printf("This is ClassA with Id=0x%08X\n", a.ID)
}

func (a ClassA) Id() int {
	return a.ID
}

func (a *ClassA) SetId(Id int) {
	a.ID = Id
}

type ClassB struct {
	BaseClass
	B string
}

func (a ClassB) Print() {
	fmt.Printf("This is ClassB with Id=0x%08X\n", a.ID)
}

func (a ClassB) Id() int {
	return a.ID
}

func (a *ClassB) SetId(Id int) {
	a.ID = Id
}

func PrintMe(obj BaseClassItf) {
	obj.Print()
}

func SetId(obj BaseClassItf, Id int) {
	obj.SetId(Id)
}

func main() {
	var a ClassA
	var b ClassB

	SetId(&a, 1)
	SetId(&b, 2)

	valueA := reflect.ValueOf(a) // <main.NotknownType Value>
	typA := reflect.TypeOf(a)    // main.NotknownType
	kndA := valueA.Kind()
	//fmt.Println(valueA, typA, kndA)
	fmt.Printf("%+v %+v %+v\n", valueA, typA, kndA)
	var x interface{}
	x = a
	valueX := reflect.ValueOf(x) // <main.NotknownType Value>
	typX := reflect.TypeOf(x)    // main.NotknownType
	kndX := valueX.Kind()
	fmt.Printf("%+v %+v %+v\n", valueX, typX, kndX)

	// 通过反射遍历处结构体中的字段
	for i := 0; i < valueA.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, valueA.Field(i))
	}

	// call the first method, which is String():
	results := valueA.Method(0).Call(nil)
	fmt.Println(results)

	PrintMe(&a)
	PrintMe(&b)

}
