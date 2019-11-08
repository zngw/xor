/**
* Copyright: Copyright (c) 2019
* Created on 2019-11-09
* Author:zengwu
* Version 1.0
* Title: 简单的异或加密、解密算法
 */
package xor

// 异或加密是一种对称性加密，所以加密、解密可以使用同一方法
// 将data数据用seed种子进行异或加密/解密, 输出加密/解密后数据
func Encode(data []byte, seed int) (out []byte) {
	out = make([]byte, len(data))
	for i, v := range data {
		out[i] = v ^ (uint8(seed))
	}

	return
}
