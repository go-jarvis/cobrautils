# cobrax 使用反射获取 flag 配置

## 安装


```bash
go get -u github.com/go-jarvis/cobrautils
```

## 使用 

> Attention: 由于 cobra 中对数据的处理方法很细致， 因此数据目前支持 `int, int64, uint, uint64`。 


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
    Gender  bool
    Address address `flag:"addr"`
}

type address struct {
    Home   string `flag:"home"`
    School string `flag:"-"`
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
        Age:    20100,
        Gender: false,
        Address: address{
            Home:   "chengdu",
            School: "shuangliu",
        },
    }

    cobrautils.BindFlags(rootCmd, &stu)
    _ = rootCmd.Execute()

    fmt.Printf("%+v", stu)
}
```

执行结果 

```bash
go run . --addr.home sichuan
Usage:
    root [flags]
Flags:
        --addr.home string    (default "chengdu")
    -a, --age int            student age (default 20100)
    -h, --help               help for root
        --name string        student name (default "zhangsanfeng")

{Name:zhangsanfeng Age:20100 Gender:false Address:{Home:sichuan School:shuangliu}}
```

`Demo`: [example](examples/main.go)

## QA

### `kind` and `type`

相较于 Type 而言，Kind 所表示的范畴更大。 类似于家用电器（Kind）和电视机（Type）之间的对应关系。或者电视机（Kind）和 42 寸彩色电视机（Type）

Type 是类型。Kind 是类别。Type 和 Kind 可能相同，也可能不同。

1. 通常基础数据类型的 Type 和 Kind 相同

2. 自定义数据类型则不同。


对于反射中的 kind 我们既可以通过 reflect.Type 来获取，也可以通过 reflect.Value 来获取。他们得到的值和类型均是相同的。




