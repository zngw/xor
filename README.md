# xor
Go语言版简单异或加密处理  
异或加密是一种对称性加密，所以加密、解密可以使用同一方法
# Quick-start
```go
package main

import (
	"fmt"
	"github.com/zngw/xor"
)

func main() {
	// 原始数据
	raw := "zngw"
	fmt.Println("raw:",raw)

	// 加密
	encode := xor.Encode([]byte(raw),55)
	fmt.Println("encode:",encode)

	// 解密
	decode := xor.Encode(encode,55)
	fmt.Println("decode:",string(decode))
}
```