# FP is a lightweight functional programming library
## [Sample document](https://pkg.go.dev/github.com/stfujnkk/fp)
---
## Main interface
```txt
func Filter(fx interface{}, arr interface{}) []interface{}        
func Filter2(fx, arr, resPtr interface{}) int
func Flat(arr interface{}) []interface{}
func Flat2(arr interface{}, resPtr interface{}) int
func Fmap(fx interface{}, arr interface{}) []interface{}
func Fmap2(fx, arr, resPtr interface{}) int
func Fold(fx, arr interface{}) interface{}
func Fold2(fx, arr, resPtr interface{})
func Reduce(resPtr, fx, arr interface{})
func UnzipWith(fx, arr interface{}) ([]interface{}, []interface{})
func UnzipWith2(fx, arr, resPtr1, resPtr2 interface{}) int        
func ZipWith(fx, arr1, arr2 interface{}) []interface{}
func ZipWith2(fx, arr1, arr2, resPtr interface{}) int
type HalfFunc func(...interface{}) interface{}
    func Currying(p, fx interface{}) HalfFunc
    func Pipe(fx1, fx2 interface{}) HalfFunc
```
## Future plans
In the future, it will join the goruntine pool
Improve concurrency efficiency.

Add some functional functions, such as TakeWhile, Generator and Infinite list.
