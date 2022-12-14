# cobrax 使用反射获取 flag 配置， 支持指针字段

cobrax 通过反射方法， 解析 struct 字段中的 **特定** 参数， 绑定 flag 。

## 安装

```bash
go get -u github.com/go-jarvis/cobrautils
```

## 使用方式

> Attention: 由于 cobra 中对数据的处理方法很细致， 因此数据目前支持 

1. 字符串:
    + `string, *string`
2. 数字:
    + `int, int64, uint, uint64`。 
    + `*int, *int8, *int16, *int32, *int64`
    + `*uint, *uint8, *uint16, *uint32, *uint64`
3. 布尔类型:
    + `bool, *bool`
4. 切片:
    + `[]string, []int, []int64`
5. 时间:
    + `timeDuration, *time.Duration`

flag 与 `cobra` 定义一致

```go
func (f *FlagSet) Uint64VarP(p *uint64, name, shorthand string, value uint64, usage string) {
	f.VarP(newUint64Value(value, p), name, shorthand, usage)
}
```

### flag 设置

```go
type student struct {
    Name    string `flag:"name" usage:"student name" persistent:"true"`
    Age     int64  `flag:"age" usage:"student age" shorthand:"a"`

    NamePtr    *string `flag:"nameptr" usage:"student name" persistent:"true"`
    AgePtr     *int64  `flag:"ageptr" usage:"student age" shorthand:"a"`
}
```

1. `flag:"config"` : flag 的名字, `--config`， 嵌套 struct 之间使用 `.` 连接, `--config.password`
2. `shorthand:"c"` : 参数简写 `-c`, 简写没有潜逃
3. `usage:"comment balalal"`: 参数说明
4. `persistent` : 全局

### flag 设置指针

从 v1.3.0 开始， flag 支持有限指针集， 以解决常用类型的 0 值问题。

包括 `*string, *int, *int64, *bool`

> https://runsisi.com/2019/04/29/go-pointer-flag/


### 默认值设置

由于所有参数的值最终都需要一个接收者， 保存之后才能够背调用。
因此， 默认值的设置就放在 `struct` 实例化一个对象中。

```go
stu := student{
    Name:   "zhangsanfeng",
    Age:    20100
}
```

### 键值绑定

```go
// 绑定
cobrautils.BindFlags(rootCmd, &stu)
_ = rootCmd.Execute()

// 打印结果
fmt.Printf("%+v", stu)
```
## 完整 Demo

```go
package main

import (
    "fmt"

    "github.com/go-jarvis/cobrautils"
    "github.com/spf13/cobra"
)

type student struct {
    Name    string `flag:"name" usage:"student name" persistent:"true"`
    Age     int64  `flag:"age" usage:"student age" shorthand:"a"`
}

var rootCmd = &cobra.Command{
    Use: "root",
    Run: func(cmd *cobra.Command, args []string) {
        _ = cmd.Help()
    },
}

func main() {
    stu := student{
        Name:   "zhangsanfeng",
        Age:    20100
    }

    cobrautils.BindFlags(rootCmd, &stu)
    _ = rootCmd.Execute()

    fmt.Printf("%+v", stu)
}
```

执行结果 

```bash
go run . --name wenzhaolun
Usage:
    root [flags]
Flags:
    -a, --age int            student age (default 20100)
    -h, --help               help for root
        --name string        student name (default "zhangsanfeng")

{Name:wenzhaolun Age:20100}
```

`Demo`: [example](examples/main.go)

## QA

### `kind` and `type`

相较于 Type 而言，Kind 所表示的范畴更大。 类似于家用电器（Kind）和电视机（Type）之间的对应关系。或者电视机（Kind）和 42 寸彩色电视机（Type）

Type 是类型。Kind 是类别。Type 和 Kind 可能相同，也可能不同。

1. 通常基础数据类型的 Type 和 Kind 相同

2. 自定义数据类型则不同。


对于反射中的 kind 我们既可以通过 reflect.Type 来获取，也可以通过 reflect.Value 来获取。他们得到的值和类型均是相同的。




