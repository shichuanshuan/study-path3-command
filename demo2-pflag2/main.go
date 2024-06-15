package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

type host struct {
	value string
}

func (h *host) String() string {
	return h.value
}

func (h *host) Set(v string) error {
	h.value = v
	return nil
}

func (h *host) Type() string {
	return "host"
}

// 进阶用法
/*
从声明标志的方法名中我们能够总结出一些规律：

pflag.<Type> 类方法名会将标志参数值存储在指针中并返回。

pflag.<Type>Var 类方法名中包含 Var 关键字的，会将标志参数值绑定到第一个指针类型的参数。

pflag.<Type>P、pflag.<Type>VarP 类方法名以 P 结尾的，支持简短标志。
*/
func main() {
	flagset := pflag.NewFlagSet("test", pflag.ExitOnError)

	var ip = flagset.IntP("ip", "i", 1234, "help message for ip")

	var boolVar bool
	flagset.BoolVarP(&boolVar, "boolVar", "b", true, "help message for boolVar")

	var h host
	flagset.VarP(&h, "host", "H", "help message for host")

	flagset.SortFlags = false

	flagset.Parse(os.Args[1:])

	fmt.Printf("ip: %d\n", *ip)
	fmt.Printf("boolVar: %t\n", boolVar)
	fmt.Printf("host: %+v\n", h)

	i, err := flagset.GetInt("ip")
	fmt.Printf("i: %d, err: %v\n", i, err)
}

/*
go run main.go --ip 1 -H localhost --boolVar=false
一个完整标志在命令行传参时使用的分界符为 --，而一个简短标志的分界符则为 -。
flagset.SortFlags = false 作用是禁止打印帮助信息时对标志进行重排序。

每一个标志都会对应一个简短标志，如 -b 和 --boolVar 是等价的，可以更加方便的设置参数。
*/
