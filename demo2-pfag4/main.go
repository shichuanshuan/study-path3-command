package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
)

// 弃用/隐藏标志
// 使用 flags.MarkDeprecated 可以弃用一个标志，使用 flags.MarkShorthandDeprecated 可以弃用一个简短标志，
// 使用 flags.MarkHidden 可以隐藏一个标志。

func main() {
	flags := pflag.NewFlagSet("test", pflag.ExitOnError)

	var ip = flags.IntP("ip", "i", 1234, "help message for ip")

	var boolVar bool
	flags.BoolVarP(&boolVar, "boolVar", "b", true, "help message for boolVar")

	var h string
	flags.StringVarP(&h, "host", "H", "127.0.0.1", "help message for host")

	// 弃用标志 ip 时，其对应的简短标志 i 也会跟着被弃用；
	flags.MarkDeprecated("ip", "deprecated")

	// 弃用 boolVar 所对应的简短标志 b 时，boolVar 标志会被保留
	flags.MarkShorthandDeprecated("boolVar", "please use --boolVar only")

	// 隐藏标志
	flags.MarkHidden("host")

	flags.Parse(os.Args[1:])

	fmt.Printf("ip: %d\n", *ip)
	fmt.Printf("boolVar: %t\n", boolVar)
	fmt.Printf("host: %+v\n", h)
}
