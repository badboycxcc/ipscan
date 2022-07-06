package tools

import (
	"fmt"
	"strconv"
	"strings"
)

type Cidr struct {
	CidrIpRange 	string
	Min 		string
	Max 		string
	Netmask 	string
	Count 		string
}

func NewCidr(ipRange string) *Cidr {
	tmp := Cidr{CidrIpRange:ipRange}
	return &tmp
}

// 获取ip最小值和最大值
// 192.168.6.1 192.168.6.255
func (this *Cidr) GetCidrIpRange() *Cidr {
	ip := strings.Split(this.CidrIpRange, "/")[0]
	ipSegs := strings.Split(ip, ".")
	maskLen := this.GetMaskLen()
	seg2MinIp, seg2MaxIp := this.GetIpSeg2Range(ipSegs, maskLen)
	seg3MinIp, seg3MaxIp := this.GetIpSeg3Range(ipSegs, maskLen)
	seg4MinIp, seg4MaxIp := this.GetIpSeg4Range(ipSegs, maskLen)
	//ipPrefix := ipSegs[0] + "." + ipSegs[1] + "."
	ipPrefix := ipSegs[0] + "."

	//this.Min = ipPrefix + strconv.Itoa(seg3MinIp) + "." + strconv.Itoa(seg4MinIp)
	this.Min = ipPrefix + strconv.Itoa(seg2MinIp) + "." +strconv.Itoa(seg3MinIp) + "." + strconv.Itoa(seg4MinIp)
	//this.Max = ipPrefix + strconv.Itoa(seg3MaxIp) + "." + strconv.Itoa(seg4MaxIp)
	this.Max = ipPrefix + strconv.Itoa(seg2MaxIp) + "." + strconv.Itoa(seg3MaxIp) + "." + strconv.Itoa(seg4MaxIp)
	return this
}

//计算得到CIDR地址范围内可拥有的主机数量
// 最大主机数
func (this *Cidr) GetCidrHostNum() *Cidr {
	cidrIpNum := uint(0)
	var i uint = uint(32 - this.GetMaskLen() - 1)
	for ; i >= 1; i-- {
		cidrIpNum += 1 << i
	}
	this.Count = fmt.Sprintf("%d",cidrIpNum)
	return this
}

func (this *Cidr) GetMaskLen() int {
	maskLen, _ := strconv.Atoi(strings.Split(this.CidrIpRange, "/")[1])
	return maskLen
}

//获取Cidr的子网掩码
func (this *Cidr) GetCidrIpMask() *Cidr {
	// ^uint32(0)二进制为32个比特1，通过向左位移，得到CIDR掩码的二进制
	cidrMask := ^uint32(0) << uint(32 - this.GetMaskLen())
	fmt.Println(fmt.Sprintf("%b \n", cidrMask))
	//计算CIDR掩码的四个片段，将想要得到的片段移动到内存最低8位后，将其强转为8位整型，从而得到
	cidrMaskSeg1 := uint8(cidrMask >> 24)
	cidrMaskSeg2 := uint8(cidrMask >> 16)
	cidrMaskSeg3 := uint8(cidrMask >> 8)
	cidrMaskSeg4 := uint8(cidrMask & uint32(255))

	this.Netmask = fmt.Sprint(cidrMaskSeg1) + "." + fmt.Sprint(cidrMaskSeg2) + "." + fmt.Sprint(cidrMaskSeg3) + "." + fmt.Sprint(cidrMaskSeg4)
	return this
}

//得到第二段IP的区间（第一片段.第二片段.第三片段.第四片段）
func (this *Cidr) GetIpSeg2Range(ipSegs []string, maskLen int) (int, int) {
	if maskLen > 8 {
		segIp, _ := strconv.Atoi(ipSegs[1])
		return segIp, segIp
	}
	ipSeg, _ := strconv.Atoi(ipSegs[1])
	return this.GetIpSegRange(uint8(ipSeg), uint8(24 - maskLen))
}


//得到第三段IP的区间（第一片段.第二片段.第三片段.第四片段）
func (this *Cidr) GetIpSeg3Range(ipSegs []string, maskLen int) (int, int) {
	if maskLen > 16 {
		segIp, _ := strconv.Atoi(ipSegs[2])
		return segIp, segIp
	}
	ipSeg, _ := strconv.Atoi(ipSegs[2])
	return this.GetIpSegRange(uint8(ipSeg), uint8(24 - maskLen))
}

//得到第四段IP的区间（第一片段.第二片段.第三片段.第四片段）
func (this *Cidr) GetIpSeg4Range(ipSegs []string, maskLen int) (int, int) {
	ipSeg, _ := strconv.Atoi(ipSegs[3])
	segMinIp, segMaxIp := this.GetIpSegRange(uint8(ipSeg), uint8(32 - maskLen))
	return segMinIp + 1, segMaxIp
}


//根据用户输入的基础IP地址和CIDR掩码计算一个IP片段的区间
func (this *Cidr) GetIpSegRange(userSegIp, offset uint8) (int, int) {
	var ipSegMax uint8 = 255
	netSegIp := ipSegMax << offset
	segMinIp := netSegIp & userSegIp
	segMaxIp := userSegIp & (255 << offset) | ^(255 << offset)
	return int(segMinIp), int(segMaxIp)
}




