// 指定网段ip存活性探测

package main

import (
	"flag"
	"fmt"
	"go-ipscan/tools"
	"strings"
)

func main() {
	fmt.Println("内网主机存活探测工具\n")
	fmt.Println("by  cxaqhq\n")
	// ip地址
	var result string
	// 线程数
	var tasknum int

	flag.StringVar(&result, "h", "", "主机名,例如：192.168.0.1/24  或者 192.168.1.1-192.168.1.254")
	flag.IntVar(&tasknum, "t", 10, "线程数")

	flag.Parse()

	if result != "" {
		ip := result
		fmt.Println("扫描IP段：", ip)
		if strings.Contains(ip, "-") == true {
			allip := tools.ParseIP1(ip)
			tools.Task(allip, tasknum)
		} else {
			iplist := tools.NewCidr(ip).GetCidrIpRange()
			allip := tools.ParseIP1(iplist.Min+ "-" + iplist.Max)
			// 输出全部ip
			tools.Task(allip, tasknum)
		}

	} else {
		fmt.Println("使用方法：\n")
		fmt.Println("scan -h 192.168.0.1/24 -t 10")
		fmt.Println("scan -h 192.168.0.1-192.168.1.1-192.168.1.254 -t 10")
		fmt.Println("默认线程数：10，最大线程数：1000")
	}
}