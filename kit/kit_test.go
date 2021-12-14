package kit

import (
	"fmt"
	"time"
	// "unsafe"
)

func ExampleMask() {
	// 定义数据
	// Define type
	type data struct {
		Num int
		x   string
		t   time.Time
	}

	type data1 struct {
		Num int
		j   bool
		x   string
	}
	var a1, a2, a3 data
	// Mask 可以传入指针也可以传值
	// Mask can pass in a pointer or a value
	fmt.Println(Mask(1, &a1) == Mask(1, &a2))
	fmt.Println(Mask(7, a1) == Mask(7, a2))
	// 嵌套结构
	// Nested structure
	a3.t = time.Now()
	a2.t, _ = time.Parse("2006-01-02 15:04", "2020-12-14 10:12")
	// 不同mask
	// Different masks
	fmt.Println(Mask(4, &a2) == Mask(3, &a3))
	fmt.Println(Mask(3, &a2) == Mask(3, &a3))
	// 不同的 time.Time 结构体
	// Different time.Time structures
	a3.t, _ = time.Parse("2006-01-02 15:04", "2020-12-14 10:12")
	fmt.Println(Mask(6, &a2) == Mask(6, &a3))
	// 不同类型之间的比较
	// Comparison between different types
	a3.Num, a3.x = 0, "asx"
	aa := data1{
		0,
		false,
		"asx",
	}
	fmt.Println(Mask(5, aa) == Mask(3, a3))
	// Output:
	// true
	// true
	// false
	// true
	// true
	// true
}

func ExampleVisit() {
	type data struct {
		Public  int
		private string
	}
	a := data{
		334700,
		"test",
	}
	// 可以传入指针也可以传值
	// can pass in a pointer or a value

	// 访问公有字段
	// Access public fields
	ptr, size := Visit(0, a)
	fmt.Println(*(*int)(ptr), size)
	// 访问私有字段
	// Access private fields
	ptr, size = Visit(1, &a)
	fmt.Println(*(*string)(ptr), size)
	// Output:
	// 334700 8
	// test 16
}
