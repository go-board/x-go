# XCrypt

提供了快速加解密的方法。

## AES
```go
// 参数分别对应 要加密的数据， 密钥， 初始化向量
// 加密
func AesEncryptRaw(data []byte, key []byte, iv []byte) ([]byte, error)
// 解密
func AesDecryptRaw(data []byte, key []byte, iv []byte) ([]byte, error)
```
