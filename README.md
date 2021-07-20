PhoneLoc-手机归属地查询
------------------
将46w+的手机号段归属地数据编码到一个不到1MB的bin文件中，并利用golang提供极致的查询性能。

### 数据文件结构

- 版本信息 100byte
- 号段映射区 100byte
- 数据记录区 46*10000*3byte
- 地址映射区 ≈3000byte

### 数据文件思路

```shell
# 以1891508为例
hexdump -Cv -n1 -s189 phone.dat # 读取号段索引
# 计算记录区偏移量 200+(49-1)*3*10000+1508*3 = 1444724
hexdump -Cv -n3 -s1444724 phone.dat # 读取记录区数据 即可得到归属地的adcode
# 00160b74  90 e3 84    得到的数据为0x84e390  最高2位表示运营商（0b00其他 0b01移动 0b10电信 0b11联通）
# 8最高两位0b10也就是电信 最高两位置0后得到归属城市的adcode 0x04e390（10进制320400，也就是常州）
```

### Thinks

- [xluohome/phonedata](https://github.com/xluohome/phonedata)
