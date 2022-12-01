package cobrautils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-jarvis/cobrautils/pflagvalue"
	"github.com/spf13/cobra"
)

func BindFlags(cmd *cobra.Command, opts interface{}, basename ...string) {

	rvPtr := reflect.ValueOf(opts)

	// 不是指针不能进行操作
	// Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装。如果v的Kind不是Interface或Ptr会panic；如果v持有的值为nil，会返回Value零值。
	if rvPtr.Kind() != reflect.Ptr && rvPtr.Elem().Kind() != reflect.Struct {
		fmt.Printf("want a Struct Ptr, got %#T \n", rvPtr.Type())
		return
	}

	// 获取 opts 结构体实例对象的反射
	// Indirect: 持有的指针指向的值的Value
	rv := reflect.Indirect(rvPtr)

	// fmt.Println(rv.Type()) // (ex) student :  具体的 结构体名字
	typ := rv.Type()
	for i := 0; i < typ.NumField(); i++ {
		/*
			var stu = student{
				Name: "zhangsan",
				Age:  20,
			}
		*/
		// typField : 结构体字段本身的属性， 与结构体实例化无关 (ex. stu.Name)
		typField := typ.Field(i)
		// valueField : 结构体实例化后字段对应的值的属性。 (ex. stu.Name -> zhangsan)
		valueField := rv.Field(i)

		// 2. 获取 name, shorthand。
		// 2.1. 获取字段名
		name := typField.Tag.Get("flag")

		// 2.1.0 如果 `name:"-"`
		if name == "-" {
			continue
		}

		// 2.1.1. 嵌套结构体, 继续循环
		if typField.Type.Kind() == reflect.Struct {
			if len(name) == 0 {
				name = strings.ToLower(typField.Name)
			}
			parts := append(basename, name)
			BindFlags(cmd, valueField.Addr().Interface(), parts...)
		}

		// 2.1.2 未设置 name 标签 或 name 为空 则跳过。
		if len(name) == 0 {
			continue
		}

		// 2.1.3 组合 flags 名字， 嵌套结构体以 . 合并
		parts := append(basename, name)
		name = strings.Join(parts, ".")

		// 2.3. 获取
		shorthand := typField.Tag.Get("shorthand")

		// 3. 获取 usage
		usage := typField.Tag.Get("usage")

		// 4. 初始化 flags 变量
		flags := cmd.Flags()

		// 4.1. 是否为 Persistent flags
		if val, ok := typField.Tag.Lookup("persistent"); ok && val == "true" {
			// fmt.Println("val=", val)
			flags = cmd.PersistentFlags()
		}

		// 6. get default value
		// value := typField.Tag.Get("value")

		// 5. 类型断言

		vIface := valueField.Interface()
		vAddrIface := valueField.Addr().Interface()
		switch v := vIface.(type) {
		case string:
			// 1.1 done : Addr() 获取值的内存地址， Interface() 并以 interface 类型返回， (*string) 并进行 类型指针类型 断言
			valuePtr := vAddrIface.(*string)
			// 1.2 done : 将 reflect.Type 值转换为对应的值
			// value := valueField.String()
			// 1.3 done: 设置 flag
			flags.StringVarP(valuePtr, name, shorthand, v, usage)

		case int:
			flags.IntVarP(vAddrIface.(*int), name, shorthand, v, usage)
		case int64:
			flags.Int64VarP(vAddrIface.(*int64), name, shorthand, v, usage)
		case uint:
			flags.UintVarP(vAddrIface.(*uint), name, shorthand, v, usage)
		case uint64:
			flags.Uint64VarP(vAddrIface.(*uint64), name, shorthand, v, usage)

		case bool:
			flags.BoolVarP(vAddrIface.(*bool), name, shorthand, v, usage)

		case []string:
			flags.StringSliceVarP(vAddrIface.(*[]string), name, shorthand, v, usage)
		case []int:
			flags.IntSliceVarP(vAddrIface.(*[]int), name, shorthand, v, usage)
		case []uint:
			flags.UintSliceVarP(vAddrIface.(*[]uint), name, shorthand, v, usage)

		case *string:
			vptr := vAddrIface.(**string)
			vv := pflagvalue.NewStringPtrValue(vptr, v)
			flags.VarP(vv, name, shorthand, usage)
		case *int:
			vv := pflagvalue.NewIntPtrValue(vAddrIface.(**int), v)
			flags.VarP(vv, name, shorthand, usage)
		case *int64:
			vv := pflagvalue.NewInt64PtrValue(vAddrIface.(**int64), v)
			flags.VarP(vv, name, shorthand, usage)
		case *bool:
			vv := pflagvalue.NewBoolPtrValue(vAddrIface.(**bool), v)
			flags.VarPF(vv, name, shorthand, usage).NoOptDefVal = "true"
		}
	}
}

func AppendCommand(child, parent *cobra.Command) {
	parent.AddCommand(child)
}
