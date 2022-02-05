package fp

import (
	"fmt"

	"github.com/stfujnkk/fp/kit"
)

func ExampleFmap() {
	square := func(a int) int {
		return a * a
	}
	fmt.Printf("%#v\n", Fmap(square, []int{4, 7}))
	// Output:[]interface {}{16, 49}
}
func ExampleFmap2() {
	square := func(a int) int {
		return a * a
	}
	var res [10]int
	n := Fmap2(square, []int{4, 7}, &res)
	fmt.Printf("%#v\n", res[:n])
	// Output:[]int{16, 49}
}

func ExampleZipWith() {
	type Student struct {
		id   int
		name string
	}
	NewStudent := func(id int, name string) Student {
		return Student{
			id:   id,
			name: name,
		}
	}
	names, ids := []string{"Jack", "John"}, []int{7, 3, 11}
	fmt.Printf("%#v\n", ZipWith(NewStudent, ids, names))
	// Output:
	// []interface {}{fp.Student{id:7, name:"Jack"}, fp.Student{id:3, name:"John"}}
}

func ExampleZipWith2() {
	type Student struct {
		id   int
		name string
	}
	NewStudent := func(id int, name string) Student {
		return Student{
			id:   id,
			name: name,
		}
	}
	names, ids := []string{"Jack", "John"}, []int{7, 3, 11}
	stus := make([]Student, 10)
	n := ZipWith2(NewStudent, ids, names, &stus)
	fmt.Printf("%#v\n", stus[:n])
	// Output:
	// []fp.Student{fp.Student{id:7, name:"Jack"}, fp.Student{id:3, name:"John"}}
}

func ExampleFilter() {
	ids := []int{23, 90, 67, 6878, 90, 8}
	fx := func(x int) bool {
		return x >= 90
	}
	fmt.Printf("%#v\n", Filter(fx, ids))
	// Output:[]interface {}{90, 6878, 90}
}

func ExampleFilter2() {
	ids := []int{23, 90, 67, 6878, 90, 8}
	fx := func(x int) bool {
		return x >= 90
	}
	n := Filter2(fx, ids, &ids)
	fmt.Printf("%#v\n", ids[:n])
	// Output:
	// []int{90, 6878, 90}
}
func ExampleFold() {
	ids := []int{-23, 90, 67, 90, 8}
	sum := func(x, y int) int {
		return x + y
	}
	fmt.Printf("%T : %v\n", Fold(sum, ids), Fold(sum, ids))
	// Output:
	// int : 232
}

func ExampleFold2() {
	ids := []int{-23, 90, 67, 90, 8}
	sum := func(x, y int) int {
		return x + y
	}
	a := 0
	Fold2(sum, ids, &a)
	fmt.Printf("%T : %v\n", a, a)
	// Output:
	// int : 232
}

func ExampleCurrying() {
	add2 := func(a, b int) int {
		return a + b
	}
	add1 := Currying(add2, 7)
	fmt.Printf("%T : %v\n", add1(2), add1(2))
	// 连续柯里化
	// Compound Currying
	res := Currying(add1, 3)
	fmt.Printf("%T : %v\n", res(), res())
	// 固定多个参数
	// Fix multiple parameters
	res2 := Currying(add2, 7, 9)
	fmt.Printf("%T : %v\n", res2(), res2())
	// 和Fmap复合使用
	// Combined with fmap
	fmt.Printf("%#v\n", Fmap(add1, []int{4, -3}))
	// 多返回值
	// Multiple return values
	swap2 := func(a, b int) (int, int) {
		return b, a
	}
	swap1 := Currying(swap2, 7)
	swap0 := Currying(swap1, 3)
	fmt.Printf("%#v\n", swap0())
	// Output:
	// int : 9
	// int : 10
	// int : 16
	// []interface {}{11, 4}
	// []interface {}{3, 7}
}

func ExamplePipe() {
	var copyNum = func(a int) (int, int) {
		return a, a
	}
	add2 := func(a, b int) int {
		return a + b
	}
	// 部分参数作为第一个函数的参数
	// Some parameters are used as parameters of the first function
	square := func(a int) int {
		return a * a
	}
	f1 := Pipe(square, add2)
	fmt.Printf("%#v\n", f1(-2, 7))
	// 全部参数作为第一个函数的参数
	// All parameters are the parameters of the first function
	f2 := Pipe(copyNum, func(a, b int) int {
		return a * b
	})
	fmt.Printf("%#v\n", f2(8))
	// Output:
	// 11
	// 64
}

func ExampleReduce() {
	box := make([]float64, 0, 10)
	data := []float64{5, 76, 67, 69, 70, -7, 8}
	collect := func(r *[]float64, b float64) {
		*r = append(*r, b)
	}
	Reduce(collect, &box, data)
	fmt.Printf("%#v\n", box)
	// Output:
	// []float64{5, 76, 67, 69, 70, -7, 8}
}

func ExampleFlat() {
	var copyNum = func(a int) (int, int) {
		return a, a
	}
	arr := Fmap(copyNum, []int{7, 9})
	fmt.Printf("%#v\n", arr)
	fmt.Printf("%#v\n", Flat(arr))
	// Output:
	// []interface {}{[]interface {}{7, 7}, []interface {}{9, 9}}
	// []interface {}{7, 7, 9, 9}
}
func ExampleFlat2() {
	var copyNum = func(a int) (int, int) {
		return a, a
	}
	arr := Fmap(copyNum, []int{7, 9})
	fmt.Printf("%#v\n", arr)
	var res [6]int
	n := Flat2(arr, &res)
	fmt.Printf("%#v\n", res[:n])
	// Output:
	// []interface {}{[]interface {}{7, 7}, []interface {}{9, 9}}
	// []int{7, 7, 9, 9}
}

func ExampleUnzipWith() {
	var copyNum = func(a int) (int, int) {
		return a, a
	}
	a, b := UnzipWith(copyNum, []int{7, -2, 0})
	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
	// Output:
	// []interface {}{7, -2, 0}
	// []interface {}{7, -2, 0}
}

func ExampleUnzipWith2() {
	var copyNum = func(a int) (int, int) {
		return a, a
	}
	var a, b [8]int
	n := UnzipWith2(copyNum, []int{7, -2, 0}, &a, &b)
	fmt.Printf("%#v\n", a[:n])
	fmt.Printf("%#v\n", b[:n])
	// Output:
	// []int{7, -2, 0}
	// []int{7, -2, 0}
}

func ExampleGroupReduce() {
	// 药品信息结构体
	// Drug information structure
	type Drug struct {
		Name        string
		Producer    string
		Price       float32
		salesVolume int
	}
	data := []Drug{
		{"氟哌酸", "上海医药", 10.13, 476},
		{"氟哌酸", "智飞生物", 12.01, 312},
		{"洛贝林", "石药集团", 5.89, 621},
		{"甲硝唑", "哈药集团", 3.13, 781},
		{"洛贝林", "恒瑞", 6.54, 437},
	}
	// 药品销售额结构体
	// Drug sales structure
	type SaleOfDrug struct {
		Name        string
		TotalAmount float32
	}
	// 统计函数
	// Statistical function
	acc := func(s *SaleOfDrug, d Drug) {
		if s == nil {
			// 统计第一个时创建一个 SaleOfDrug
			// Create a SaleOfDrug when counting the first
			s = &SaleOfDrug{
				d.Name,
				d.Price * float32(d.salesVolume),
			}
		} else {
			(*s).Name = d.Name
			(*s).TotalAmount += d.Price * float32(d.salesVolume)
		}
	}
	res := make([]SaleOfDrug, 30)
	n :=GroupReduce(1, acc, &res, data)
	fmt.Println(res[:n])
	// Output:
	// [{氟哌酸 8569} {洛贝林 6515.67} {甲硝唑 2444.53}]
}

func ExampleGroup() {
	// 药品信息结构体
	// Drug information structure
	type Drug struct {
		Name        string
		Producer    string
		Price       float32
		salesVolume int
	}
	data := []Drug{
		{"氟哌酸", "上海医药", 10.13, 476},
		{"氟哌酸", "智飞生物", 12.01, 312},
		{"洛贝林", "石药集团", 5.89, 621},
		{"甲硝唑", "哈药集团", 3.13, 781},
		{"洛贝林", "恒瑞", 6.54, 437},
	}
	// 根据药品名称分组
	// Grouping by drug name
	m := Group(1, data)
	drugs := make([]Drug, 0, len(m))
	// 不保证顺序
	// Sequence is not guaranteed
	k := kit.Mask(1, data[0])
	for _, v := range m[k] {
		drugs = append(drugs, v.(Drug))
	}
	fmt.Println(drugs)
	// Output:
	// [{氟哌酸 上海医药 10.13 476} {氟哌酸 智飞生物 12.01 312}]
}