package main

import (
	"fmt"
	"reflect"
)

type Struct struct{
	Field0	[]int
	Field1	int
	Field2	[]int
}

func otherfunc(a Struct, size int) { //inter interface {},size int) {
	//arglist := make([]interface{},0,size)
	//a := reflect.ValueOf(inter)
	ps := reflect.ValueOf(&a)
	s := ps.Elem()
	fmt.Println(s)
	if s.Kind() == reflect.Struct {
		f := s.NumField();
		for i := 0; i<f ; i++ {
			if s.Field(i).Kind() == reflect.Slice {
				//j := s.Field(i).Interface()
				fmt.Println(s.Field(i).Index(0).Interface())
/*
				fmt.Println(s.Field(i).Kind())
				fmt.Println(s.Field(i).Type())
				fmt.Println(make([]int,2,3))
				fmt.Println(s.Field(i).Interface())
				fmt.Println(reflect.ValueOf(j))
				fmt.Println(s.Field(i).CanSet())
				*/
				//fmt.Println(s.Field(i).Interface().(slice))
				//fmt.Println(s.Field(i).Slice())
			}
		}

	}
	/*
	for i := 0; i<v.NumField(); i++ {
		fmt.Println(v)
	}*/
}
func main() {
	s := Struct{make([]int,3),3,make([]int,3)}
	s.Field0[0] = 3
	//fmt.Println(s)
	otherfunc(s,4)
	unroll_struct(s)

	/*
	v := reflect.ValueOf(s)
	i := v.Interface()
	n := reflect.ValueOf(i).NumField()
	for j := 0; j<n ; j++ {
		f := reflect.ValueOf(i).Field(j)
		if f.Kind() == reflect.Slice {
			fmt.Println(f.Len()) //Interface())
			fmt.Println(f.Index(0).Interface())
		}
	}
	*/
}

func unroll_struct(inter interface{}) {
	arg_count := 0
	n := reflect.ValueOf(inter).NumField()
	for j := 0; j<n ; j++ {
		f := reflect.ValueOf(inter).Field(j)
		if f.Kind() == reflect.Slice || f.Kind() == reflect.Array {
			for c := 0; c < f.Len(); c++ {
				fmt.Println(f.Index(c).Interface())
				arg_count += 1
			}
			arg_count -= 1
		}
		arg_count += 1
	}
}
/*
func unroll_query(strq []string, varargs ...interface{}) (string,[]interface{}) {
	stro := ""
	argo := make([]interface{},0,len(strq)-1)
	count := 1
	for i := range varargs {
		stro += strq[i]
		k := reflect.ValueOf(varargs[i])
		if k.Kind() == reflect.Slice || k.Kind() == reflect.Array {
			for c := 0; c < k.Len()-1; c++ {
				stro += '$'+strconv.Itoa(count)+','
				append(argo,k.Index(c).Interface())
			}
			stro += '$'+strconv.Itoa(count)
			append(argo,k.Index(k.Len()-1).Interface())
		} else {
			stro += varargs[i].Interface()
			append(argo,varargs[i].Interface())
		}
	}
	stro += strq[len(strq)-1]
	return stro,argo
}
*/
