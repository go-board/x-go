# XDebug

一些调试工具。

## Dump 美化输出

## Gc 垃圾收集 http endpoint handler
```go
type gc struct{}

func (gc) ServeHTTP(http.ResponseWriter, *http.Request) { ... }
```

## PProf 性能调优
```go
// 注册pprof handler 到 *http.ServeMux上面
func PProfHandlers(h *http.ServeMux)
```
