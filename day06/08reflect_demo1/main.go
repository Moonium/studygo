package main

import (
	"fmt"
	"reflect"
)

// Cat 结构体
type Cat struct {
}

func reflectType(x interface{}) {
	value := reflect.TypeOf(x)
	fmt.Printf("type: %v\n", value)
	fmt.Printf("type name: %v type kind: %v\n", value.Name(), value.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}
func main() {
	var a float32 = 3.14
	reflectType(a)
	var b int64 = 100
	reflectType(b)

	reflectValue(a) // type is float32, value is 3.140000
	reflectValue(b)

	var c = Cat{}
	reflectType(c)

	reflectSetValue1(&b)
	fmt.Println(b)

	reflectSetValue2(&b)
	fmt.Println(b)
}
