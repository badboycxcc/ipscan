# 内网IP存活性探测工具 

## 使用方法
```
支持A、B、C段主机探测
scan -h 192.168.0.1/24
scan -h 192.168.0.1/24 -t 100
scan -h 192.168.0.1-192.168.0.255 -t 100
```
![image](https://user-images.githubusercontent.com/72059221/177450567-ab9175fa-cb82-4ce9-bf3c-5b5b13f65865.png)


## 参考项目

[phil-fly](https://github.com/phil-fly/go-ipscan)  
[cidr](https://github.com/lflxp/cidr/blob/master/cidr.go)  
[fscan](https://github.com/shadow1ng/fscan/blob/main/common/ParseIP.go)  

## 欢迎提BUG
