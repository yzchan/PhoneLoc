PhoneLoc-手机归属地查询
------------------
将46w+的手机号段归属地数据编码到一个不到1MB的bin文件中，并利用golang提供极致的查询性能。

### 数据文件结构

- 版本信息 100byte
- 号段映射区 100byte
- 数据记录区 46*10000*2byte
- 地址映射区 ≈3000byte

### 数据文件思路

```shell
# 1342021归属地查询
hexdump -Cv -n1 -s134 phone.dat # 读取号段索引
# 计算记录区偏移量 200+(5-1)*10000*2+2021*2 = 84242 
hexdump -Cv -n2 -s84242 phone.dat # 读取记录区数据 即可得到归属地的adcode
```

### Thinks
 - [xluohome/phonedata](https://github.com/xluohome/phonedata)
