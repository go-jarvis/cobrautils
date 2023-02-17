package main

import (
	"github.com/go-jarvis/cobrautils"
	"github.com/spf13/cobra"
)

func main() {
	root.Execute()
}

var root = &cobra.Command{
	Use:   "student",
	Short: "student info",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	cobrautils.BindFlags(root, stu)
}

var stu = &Strudent{
	Name: "zhangsan",
	Address: Address{
		Home: "Sichuan,China",
	},
	Contact: Contact{
		Phone: "138-8888-8888",
	},
}

type Strudent struct {
	Name    string  `flag:"name" usage:"student name" persistent:"true"`
	Age     int64   `flag:"age" usage:"student age" shorthand:"a"`
	Gender  bool    `flag:"gender" persistent:"true"`
	Address Address `flag:"addr"`
	Contact Contact `flag:"contcat"`
}

type Address struct {
	Home   string `flag:"home" usage:"home address"`
	School string `flag:"school" usage:"school address"`
}

type Contact struct {
	Phone string `flag:"phone" usage:"phone number"`
	Email string `flag:"email" usage:"email address"`
}
