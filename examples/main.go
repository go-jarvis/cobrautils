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

	/*
	   go run . --addr.home sichuan
	   Usage:
	     root [flags]

	   Flags:
	         --addr.home string    (default "chengdu")
	     -a, --age int            student age (default 20100)
	     -h, --help               help for root
	         --name string        student name (default "zhangsanfeng")

	   {Name:zhangsanfeng Age:20100 Gender:false Address:{Home:sichuan School:shuangliu}}
	*/
}
