/*
简易函数式编程库
Simple functional programming library
*/
package fp

import (
	"fmt"
	"reflect"

	"github.com/stfujnkk/fp/kit"
)

/*
将arr中每个元素包装为reflect.Value数组以便于反射调用。
*/
/*
Wrap each element in the ARR as reflect Value array for reflection calls.
*/
func prepare(arr []interface{}, num int) [][]reflect.Value {
	ps := make([][]reflect.Value, len(arr))
	for i := range arr {
		a, t := toArr(arr[i]), make([]reflect.Value, num)
		for j := range a {
			t[j] = reflect.ValueOf(a[j])
		}
		ps[i] = t
	}
	return ps
}

/*
fx 过滤函数, arr待处理列表。
用fx过滤arr中元素。
fx返回为真时保留,假时剔除。
*/
/*
FX filter function, arr pending list.
Filter elements in arr with FX.
If FX returns true, it will be retained and if false, it will be rejected.
*/
func Filter(fx interface{}, arr interface{}) []interface{} {
	f := reflect.ValueOf(fx)
	aa := toArr(arr)
	ps := prepare(aa, f.Type().NumIn())
	res := make([]interface{}, 0, len(ps)/2+1)
	for i := range aa {
		if f.Call(ps[i])[0].Interface().(bool) {
			res = append(res, aa[i])
		}
	}
	return res
}

/*
把 a 转换为数组。
若a不是数组,则用一个长度为1的数组包装。
*/
/*
Convert a to an array.
If a is not an array, wrap it with an array of length 1.
*/
func toArr(a interface{}) (res []interface{}) {
	v := reflect.ValueOf(a)
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		res = make([]interface{}, v.Len())
		for i := range res {
			res[i] = v.Index(i).Interface()
		}
	default:
		res = []interface{}{a}
	}
	return
}

/*
fx 过滤函数, arr待处理列表,resPtr 用于存储结果的地址,返回值为结果长度。
与Filter函数作用一样,只不过可以通过resPtr存储结果,并指定类型。
需要注意的是resPtr如果是切片或数组类型需要预留好足够空间。
*/
/*
FX filter function, arr pending list, resptr address used to store results,
The return value is the result length.
The function is the same as the filter function, except that the result can be stored through resptr and the type can be specified.
It should be noted that if resptr is a slice or array type, enough space should be reserved.
*/
func Filter2(fx, arr, resPtr interface{}) int {
	t := Filter(fx, arr)
	fill(resPtr, t...)
	return len(t)
}

/*
展开arr中各个元素
*/
/*
Expand the elements in the arr
*/
func Flat(arr interface{}) []interface{} {
	a := toArr(arr)
	res := make([]interface{}, 0, len(a))
	for i := range a {
		res = append(res, toArr(a[i])...)
	}
	return res
}

/*
展开arr中各个元素,并把结果存到resPtr。返回值为结果长度
*/
/*
Expand the elements in arr and save the results to resptr. The return value is the result length
*/
func Flat2(arr interface{}, resPtr interface{}) int {
	t := Flat(arr)
	fill(resPtr, t...)
	return len(t)
}

/*
对arr中每个元素应用一次fx
*/
/*
Apply FX once to each element in the arr
*/
func Fmap(fx interface{}, arr interface{}) []interface{} {
	f := reflect.ValueOf(fx)
	j := f.Type().NumOut()
	ps := prepare(toArr(arr), f.Type().NumIn())
	res := make([]interface{}, len(ps))
	if j == 1 {
		for i := range res {
			res[i] = f.Call(ps[i])[0].Interface()
		}
	} else {
		for i := range res {
			outs := f.Call(ps[i])
			t := make([]interface{}, j)
			for x := range t {
				t[x] = outs[x].Interface()
			}
			res[i] = t
		}
	}
	return res
}

/*
通过fx函数合并arr1和arr2
*/
/*
Merge Arr1 and arr2 through FX function
*/
func ZipWith(fx, arr1, arr2 interface{}) []interface{} {
	a1, a2 := toArr(arr1), toArr(arr2)
	n := len(a1)
	if n > len(a2) {
		n = len(a2)
	}
	res := make([]interface{}, n)
	f := reflect.ValueOf(fx)
	for i := range res {
		res[i] = f.Call([]reflect.Value{reflect.ValueOf(a1[i]), reflect.ValueOf(a2[i])})[0].Interface()
	}
	return res
}

/*
通过fx函数合并arr1和arr2,并把结果存到resPtr。返回值为结果长度
*/
/*
Merge Arr1 and arr2 through the FX function and save the results in resptr. The return value is the result length.
*/
func ZipWith2(fx, arr1, arr2, resPtr interface{}) int {
	t := ZipWith(fx, arr1, arr2)
	fill(resPtr, t...)
	return len(t)
}

/*
对arr中的元素两两应用fx函数。
若arr长度为0返回nil,长度为1返回原数组。
*/
/*
Apply the FX function to every two elements in the arr.
If the ARR length is 0, nil is returned, and if the ARR length is 1, the original array is returned.
*/
func Fold(fx, arr interface{}) interface{} {
	a := toArr(arr)
	n := len(a)
	if n == 0 {
		return nil
	}
	f, res := reflect.ValueOf(fx), a[0]
	for i := 1; i < n; i++ {
		res = f.Call([]reflect.Value{reflect.ValueOf(res), reflect.ValueOf(a[i])})[0].Interface()
	}
	return res
}

/*
对arr中的元素两两应用fx函数,结果存入resPtr。
若arr长度为0返回nil,长度为1返回原数组。
*/
/*
Apply the FX function to two elements in arr, and store the result in resptr.
If the ARR length is 0, nil is returned, and if the ARR length is 1, the original array is returned.
*/
func Fold2(fx, arr, resPtr interface{}) {
	t := Fold(fx, arr)
	fill(resPtr, t)
}

/*
把arr内容填充到ptr指向的对象。
*/
/*
Fill the ARR content into the object pointed to by PTR.
*/
func fill(ptr interface{}, arr ...interface{}) {
	v := reflect.ValueOf(ptr).Elem()
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		for i := range arr {
			v.Index(i).Set(reflect.ValueOf(arr[i]))
		}
	default:
		v.Set(reflect.ValueOf(arr[0]))
	}
}

/*
对arr中每个元素应用一次fx,并把结果存到resPtr。返回值为结果长度。
*/
/*
Apply FX to each element in arr once and save the result in resptr. The return value is the result length.
*/
func Fmap2(fx, arr, resPtr interface{}) int {
	t := Fmap(fx, arr)
	fill(resPtr, t...)
	return len(t)
}

/*
半函数。
*/
/*
Function with partial parameters fixed.
*/
type HalfFunc func(...interface{}) interface{}

/*
p 为参数列表,fx 为函数。
返回一个HalfFunc。
Currying可以固定一个函数的前几个参数。
*/
/*
P is the parameter list and FX is the function. Returns a HalfFunc.
Currying can fix the first few parameters of a function.
*/
func Currying(fx interface{}, p ...interface{}) HalfFunc {
	f := reflect.ValueOf(fx)
	return func(arr ...interface{}) interface{} {
		n, pl := len(arr)+len(p), len(p)
		if n != f.Type().NumIn() {
			msg := fmt.Sprintf("Takes %d positional argument but %d were given\n", f.Type().NumIn(), n)
			panic("Wrong number of parameters\n" + msg)
		}
		ps := make([]reflect.Value, n)
		for i := range p {
			ps[i] = reflect.ValueOf(p[i])
		}
		for i := pl; i < n; i++ {
			ps[i] = reflect.ValueOf(arr[i-pl])
		}
		t := Fmap(reflect.Value.Interface, f.Call(ps))
		if len(t) == 1 {
			return t[0]
		}
		return t
	}
}

/*
fx函数应接收两个参数,一个为resPtr,另一个为arr中的元素。
该函数遍历arr,并调用fx函数。
*/
/*
The FX function should take two arguments, one resptr and the other an element in arr.
This function traverses the ARR and calls the FX function.
*/
func Reduce(fx, resPtr, arr interface{}) {
	a, f := toArr(arr), reflect.ValueOf(fx)
	for i := range a {
		f.Call([]reflect.Value{reflect.ValueOf(resPtr), reflect.ValueOf(a[i])})
	}
}

/*
通过fx,分裂列表arr。
*/
/*
Split the list arr through FX.
*/
func UnzipWith(fx, arr interface{}) ([]interface{}, []interface{}) {
	a, f := toArr(arr), reflect.ValueOf(fx)
	r1, r2 := make([]interface{}, len(a)), make([]interface{}, len(a))
	for i := range a {
		outs := f.Call([]reflect.Value{reflect.ValueOf(a[i])})
		r1[i], r2[i] = outs[0].Interface(), outs[1].Interface()
	}
	return r1, r2
}

/*
通过fx,分裂列表arr,并把结果存入 resPtr1, resPtr2。
*/
/*
Split the list arr through FX and store the results in resptr1 and resptr2.
*/
func UnzipWith2(fx, arr, resPtr1, resPtr2 interface{}) int {
	a, b := UnzipWith(fx, arr)
	fill(resPtr1, a...)
	fill(resPtr2, b...)
	return len(a)
}

/*
组合两个函数,返回一个 HalfFunc。
返回的函数会把收到的参数尽量给fx1。
fx1的结果和余下的参数作为fx2的参数。
最终返回fx2的结果。
*/
/*
Combine two functions to return a HalfFunc.
The returned function will give the received parameters to FX1 as much as possible.
The result of FX1 and the remaining parameters are the parameters of FX2.
Finally, the result of FX2 is returned.
*/
func Pipe(fx1, fx2 interface{}) HalfFunc {
	f1, f2 := reflect.ValueOf(fx1), reflect.ValueOf(fx2)
	return func(p ...interface{}) interface{} {
		n := f1.Type().NumIn()
		p1, p2 := make([]reflect.Value, n), make([]reflect.Value, len(p)-n)
		Fmap2(reflect.ValueOf, p[:n], &p1)
		Fmap2(reflect.ValueOf, p[n:], &p2)
		rs := Fmap(reflect.Value.Interface, f2.Call(append(f1.Call(p1), p2...)))
		if len(rs) == 1 {
			return rs[0]
		}
		return rs
	}
}

/*
根据mask对arr分组，返回map。map键的类型为byte数组。注意不是切片,golang中切片无法比较。
*/
/*
Group arr according to mask and return map. The type of map key is byte array. Note that it is not a slice. Slices in golang cannot be compared.
*/
func Group(mask int64, arr interface{}) map[interface{}][]interface{} {
	a := toArr(arr)
	m := make(map[interface{}][]interface{}, len(a))
	for i := range a {
		k := kit.Mask(mask, a[i])
		v, ok := m[k]
		if !ok {
			v = make([]interface{}, 0, 3)
		}
		v = append(v, a[i])
		m[k] = v
	}
	return m
}

/*

根据mask对arr分组，对每组列表的元素都应用一次reducer函数,并把结果存在res里。
最后返回结果列表长度。
*/
/*
Group the ARR according to the mask, apply the reducer function to the elements of each list, and store the results in res.
Length of the last returned result list.
*/
func GroupReduce(mask int64, reducer, res, arr interface{}) int {
	m, i := Group(mask, arr), 0
	a, f := reflect.ValueOf(res).Elem(), reflect.ValueOf(reducer)
	for _, v := range m {
		for j := range v {
			f.Call(
				[]reflect.Value{
					a.Index(i).Addr(),
					reflect.ValueOf(v[j]),
				},
			)
		}
		i++
	}
	return i
}
