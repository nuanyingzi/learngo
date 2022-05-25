package main

import (
	"fmt"
	"reflect"
)

// reflect demo

type Cat struct {
}
type Dog struct {
}

func reflectType(x interface{}) {
	// 我不知道别人调用我这个函数的时候会传进来什么类型的变量
	// 1. 方法1：通过类型断言
	// 2. 方法2：借助反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj, obj.Name(), obj.Kind())
	//fmt.Printf("%T\n", obj) // *reflect.rtype

}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v, %T\n", v, v)
	// 如何得到一个传入时候类型的变量
	k := v.Kind() // 拿到值对应的类型种类
	switch k {
	case reflect.Float32:
		// 把反射取到的值转换成一个float32类型的变量
		ret := float64(v.Float())
		fmt.Printf("%v, %T\n", ret, ret)
	case reflect.Int32:
		// 把反射取到的值转换成一个int32类型的变量
		ret := int8(v.Int())
		fmt.Printf("%v, %T\n", ret, ret)
	}
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	// Elem()用来根据指针取对应的值
	k := v.Elem().Kind()
	switch k {
	case reflect.Int32:
		v.Elem().SetInt(100)
	case reflect.Float32:
		v.Elem().SetFloat(3.21)
	}
}

func main() {
	//// 基本数据类型
	//var a float32 = 9.19
	//reflectType(a)
	//var b int8 = 9
	//reflectType(b)
	//// 结构体类型
	//var c Cat
	//reflectType(c)
	//var d Dog
	//reflectType(d)
	//// slice
	//var e []int
	//reflectType(e)
	//var f []string
	//reflectType(f)

	//var aa int32 = 99
	//reflectValue(aa)
	//var bb float32 = 1.234
	//reflectValue(bb)

	var aaa int32 = 10
	reflectSetValue(&aaa)
	fmt.Println(aaa)
}
