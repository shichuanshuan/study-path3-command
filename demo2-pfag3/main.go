package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

func normalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	// alias
	switch name {
	case "old-flag-name":
		name = "new-flag-name"
		break
	}

	// --my-flag == --my_flag == --my.flag
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return pflag.NormalizedName(name)
}

// 借助 pflag.NormalizedName 我们能够给标志起一个或多个别名、规范化标志名等。
/*
要使用 pflag.NormalizedName，我们需要创建一个函数 normalizeFunc，然后将其通过 flagset.SetNormalizeFunc(normalizeFunc) 注入到 flagset 使其生效。

在 normalizeFunc 函数中，我们给 new-flag-name 标志起了一个别名 old-flag-name
*/
func main() {
	flagset := pflag.NewFlagSet("test", pflag.ExitOnError)

	var ip = flagset.IntP("new-flag-name", "i", 1234, "help message for new-flag-name")
	var myFlag = flagset.IntP("my-flag", "m", 1234, "help message for my-flag")

	flagset.SetNormalizeFunc(normalizeFunc)
	flagset.Parse(os.Args[1:])

	fmt.Printf("ip: %d\n", *ip)
	fmt.Printf("myFlag: %d\n", *myFlag)
}
