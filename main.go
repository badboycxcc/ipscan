// 指定网段ip存活性探测

package main

import (
	"fmt"
	"go-ipscan/tools"
	"os"
)

func main() {
	fmt.Println("内网主机存活探测工具\n")
	fmt.Println("by  cxaqhq\n")
	var result string
	for k, v := range os.Args {
		if k == 1 {
			result = v
		}
	}
	if result != "" {
		ip := result
		fmt.Println("扫描IP段：", ip)

		iplist := tools.NewCidr(ip).GetCidrIpRange()
		fmt.Println("起始IP",iplist.Min)
		fmt.Println("结束IP",iplist.Max)
		allip := tools.ParseIP1(iplist.Min+ "-" + iplist.Max)
		tools.Task(allip)
	} else {
		fmt.Println("使用方法：scan 192.168.0.1/24")
	}
}