package main

import (
	"fmt"

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

// 示例演示的 pflag 用法跟 flag 包用法一致，可以做到二者无缝替换。
// 与 flag 包不同的是，pflag 包参数定界符是两个 -，而不是一个 -，在 pflag 中 -- 和 - 具有不同含义
func main() {
	var ip1 *int = pflag.Int("ip", 1234, "help message for ip")

	var port int
	pflag.IntVar(&port, "port", 8080, "help message for port")

	var h host
	pflag.Var(&h, "host", "help message for host")

	// 解析命令行参数
	pflag.Parse()

	fmt.Printf("ip: %d\n", *ip1)
	fmt.Printf("port: %d\n", port)
	fmt.Printf("host: %+v\n", h)

	fmt.Printf("NFlag: %v\n", pflag.NFlag()) // 返回已设置的命令行标志个数
	fmt.Printf("NArg: %v\n", pflag.NArg())   // 返回处理完标志后剩余的参数个数
	fmt.Printf("Args: %v\n", pflag.Args())   // 返回处理完标志后剩余的参数列表
	fmt.Printf("Arg(1): %v\n", pflag.Arg(1)) // 返回处理完标志后剩余的参数列表中第 i 项
}

/*
go run main.go --ip 1 x y --host localhost a b
ip 标志的默认值已被命令行参数 1 所覆盖，由于没有传递 port 标志，所以打印结果为默认值 8080，host 标志的值也能够被正常打印。

还有 4 个非选项参数数 x、y、a、b 也都被 pflag 识别并记录了下来。这点比 flag 要强大，在 flag 包中，
非选项参数数只能写在所有命令行参数最后，x、y 出现在这里程序是会报错的。
*/
