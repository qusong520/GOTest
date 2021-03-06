/**
go language basic test
*/
package basic

import (
	"fmt"
	"reflect"
	"testing"
)

// https://juejin.im/post/5a75a4fb5188257a82110544

func TestAddPerson(t *testing.T) {
	//AddPerson()
}

// TestPersonReflect
func TestPersonReflect(t *testing.T) {

	per := Person{
		Name: "zhangsan",
		Age:  20,
	}

	perP := &per

	perV := reflect.ValueOf(perP)

	a := perV.Elem()

	fmt.Printf("poing value elem type is :%s\n", a.Type().Kind())

	for i := 0; i < a.Type().NumField(); i++ {
		fmt.Println(a.Type().Field(i).Name)
	}

	fmt.Println(a.FieldByName("Name"))

	fmt.Println(perV.Type())
	fmt.Println(perV.Type().Kind())

	fmt.Println(perV.Elem())

}

// TestPersonReflect2
func TestPersonReflect2(t *testing.T) {

	per := &Person{
		Name: "zhangsan",
		Age:  20,
	}

	perV := reflect.ValueOf(per)

	perV = perV.Elem()

	// operate filed value
	for i := 0; i < perV.Type().NumField(); i++ {
		vv := perV.Field(i)
		kk := perV.Type().Field(i)
		fmt.Printf("vv name: %s, type: %s, vv: %v, tag: %v\n", kk.Name, vv.Kind(), vv, kk.Tag)
	}

	perT := reflect.TypeOf(*per)
	fmt.Printf("the type of person is: %v\n", perT)
	//fmt.Printf("the element value of person is: %v\n", perT.NumField())  // 报错，不能计算一个指针的field
	//fmt.Printf("the element value of person is: %v\n", perT.Type().NumField()) // 报错，不能计算一个类型的field； struct field?
	fmt.Printf("the field length of person is: %d\n", perT.NumField())

	// operate field
	for i := 0; i < perT.NumField(); i++ {
		kk := perT.Field(i)
		vv := perV.Field(i)
		fmt.Printf("vv name: %s, type: %s, vv: %v, tag: %v,tagValue: %s\n", kk.Name, vv.Kind(), vv, kk.Tag, kk.Tag.Get("json"))

		// GetFieldTag
		value, ok := kk.Tag.Lookup("json")
		fmt.Printf("tag looup %s,isOk: %t\n", value, ok)
	}

	// operate field tag
	for i := 0; i < perT.NumField(); i++ {
		field := perT.Field(i)

		// GetFieldTag
		value, ok := field.Tag.Lookup("json")
		fmt.Printf("tag looup %s,isOk: %t\n", value, ok)

		fmt.Printf("tag value is: %s", field.Tag.Get("json"))

	}

	// To json
	fmt.Printf("person to jsonstr: %s\n", ToJsonString(per))

	// To map
	fmt.Printf("person to map: %v\n", per.toMap())
}

// 反射Slice 并进行反射操作
func TestPersonReflect3(t *testing.T) {
	var persons []Person
	for i := 0; i < 10; i++ {
		persons = append(persons, Person{
			Name: fmt.Sprintf("张三%d", i),
			Age:  20 + i,
		})
	}

	kk := reflect.TypeOf(persons).Kind()
	values := reflect.ValueOf(persons)

	fmt.Printf("type of interface, %v\n", kk)
	fmt.Printf("value of interface, %v\n", values)

	if values.Len() > 0 {
		types := reflect.TypeOf(values.Index(0))
		for i := 0; i < types.NumField(); i++ {
			fmt.Printf(types.Field(i).Name)
		}

	}

	//for i := 0; i < values.Len(); i++ {
	//	vv := values.Index(i)
	//	//fmt.Printf("value of interface, %v\n", vv)
	//	fmt.Printf("elem %v", vv.Type())
	//	types := reflect.TypeOf(vv)
	//	for i := 0; i < types.NumField(); i++ {
	//		fmt.Printf("type of interface, %v\n", types.Field(i).Name)
	//
	//	}
	//
	//}
}

func TestBaseReflect(t *testing.T) {
	aa := 1
	bb := 1.0
	cc := true
	dd := "123"
	ee := &Person{
		Name: "zhangsan",
		Age:  20,
	}
	ff := []string{"hello", "abc"}

	fmt.Printf("type of interface, %v\n", reflect.ValueOf(aa).Kind())
	fmt.Printf("type of interface, %v\n", reflect.ValueOf(bb).Kind())
	fmt.Printf("type of interface, %v\n", reflect.ValueOf(cc).Kind())
	fmt.Printf("type of interface, %v\n", reflect.ValueOf(dd).Kind())
	fmt.Printf("type of interface, %v\n", reflect.ValueOf(ee).Kind())
	fmt.Printf("type of interface, %v\n", reflect.ValueOf(ff).Kind())

}
