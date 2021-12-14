package kit

import (
	"reflect"
	"unsafe"
)

var byteType reflect.Type

// var anyType reflect.Type

func init() {
	// var in interface{}
	// anyType = reflect.TypeOf(&in).Elem()
	var b byte
	byteType = reflect.TypeOf(b)
}

/* 如果obj是指针返回他指向的类型。否者返回自己的类型 */
/* If obj is a pointer, return the type it points to. If not, return your own type */
func maybePtr(obj interface{}) reflect.Type {
	o := reflect.TypeOf(obj)
	switch o.Kind() {
	case reflect.Ptr:
		return o.Elem()
	default:
		return o
	}
}

/*访问obj结构体的第i个字段，返回字段的地址和长度*/
/*Access the ith field of obj structure and return the address and length of the field*/
func Visit(i int, obj interface{}) (unsafe.Pointer, uintptr) {
	f := maybePtr(obj).Field(i)
	return unsafe.Pointer((*[2]uintptr)(unsafe.Pointer(&obj))[1] + f.Offset), f.Type.Size()
}

/*
根据mask屏蔽obj结构体里的一些字段,返回一个byte数组。
该数组可来作为map键。
*/
/*
Mask some fields in the obj structure and return a byte array.
This array can be used as a map key.
*/
func Mask(mask int64, obj interface{}) interface{} {
	arr := reflect.New(
		reflect.ArrayOf(
			getSize(mask, obj),
			byteType,
		),
	).Elem()
	for i, l := 0, 0; mask != 0; i++ {
		if mask&1 == 1 {
			p, s := Visit(i, obj)
			ps := (uintptr)(p)
			pe := ps + s
			for ps != pe {
				arr.Index(l).Set(
					reflect.ValueOf(
						*((*byte)(unsafe.Pointer(ps))),
					),
				)
				l++
				ps++
			}
		}
		mask >>= 1
	}
	return arr.Interface()
}

/*
根据mask计算出所需byte数组的长度
*/
/*
Calculate the length of the required byte array according to the mask
*/
func getSize(mask int64, ptr interface{}) int {
	t, l := maybePtr(ptr), 0
	// fmt.Printf("\n%#v\n", t.Kind())
	for i := 0; mask != 0; i++ {
		if mask&1 == 1 {
			l += int(t.Field(i).Type.Size())
		}
		mask >>= 1
	}
	return l
}
