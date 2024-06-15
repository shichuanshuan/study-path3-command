package main

import (
	"flag"
	"fmt"

	"github.com/spf13/pflag"
)

/*
支持 flag 类型
由于 pflag 对 flag 包兼容，所以可以在一个程序中混用二者
*/
func main() {
	var ip *int = pflag.Int("ip", 1234, "help message for ip")
	var port *int = flag.Int("port", 80, "help message for port")

	/*
		其中，ip 标志是使用 pflag.Int() 声明的，port 标志则是使用 flag.Int() 声明的。
		只需要通过 AddGoFlagSet 方法将 flag.CommandLine 注册到 pflag 中，那么 pflag 就可以使用 flag 中声明的标志集合了
	*/
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	fmt.Printf("ip: %d\n", *ip)
	fmt.Printf("port: %d\n", *port)
}
