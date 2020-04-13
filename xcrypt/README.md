# XCrypt

提供了快速加解密的方法。

## AES
```go
package xcrypt
// 参数分别对应 要加密的数据， 密钥， 初始化向量
// 加密
func AesEncryptRaw(data []byte, key []byte, iv []byte) ([]byte, error)
// 解密
func AesDecryptRaw(data []byte, key []byte, iv []byte) ([]byte, error)
```

## Bcrypt
bcrypt 提供了密码加密/验证相关的操作。

brcypt主要思想是拖慢破解的速度，在cost为10的情况下，加密要比md5慢上百万倍，对应的暴力破解也是如此。
在值域一定的情况下，破解越慢， 被破解的可能性越低，从而保护密码的安全。
```go
package xcrypt
func SetCost(cost int) {}

func BCryptHash(password []byte) ([]byte, error) {}

func BCryptValidate(password []byte, encryptedData []byte) bool {}
```
