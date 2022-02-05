## brief introduction
fp is a lightweight functional programming library

[** pkg.go.dev **](https://pkg.go.dev/github.com/stfujnkk/fp)

[**Github document**](https://stfujnkk.github.io/fp/)

## Version compatible

- **v0.0.1**

  Naming habits: adding 2 after the function name means to which means you can specify where the results are stored. The advantage is that you can add type information. You don't have to type assertions one by one.

- **v0.1.0**

  This version changes the position of some parameters because it is more convenient to Coriolis fixed parameters. Put the parameters that can control the behavior on the left and the parameters that transfer data on the right. For example, `reduce (resptr, fX, arr interface {})` is changed to `reduce (fX, resptr, arr interface {}`. Therefore, this also leads to and **v0.0.1** incompatible. In addition, three functions are added: ` group, visit, groupreduce and mask`





## Future plans
In the future, it will join the goruntine pool
Improve concurrency efficiency.

Add some functional functions, such as `TakeWhile`, `Generator` and `Infinite list`.

