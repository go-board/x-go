# XTime

time 工具集

## parse 支持多种类型的parse

go 默认的 `time.Parse` 需要指定对应的layout，导致使用起来不够便利，因此催生了 `xtime.ParseAny` 

``` go
package main

import (
    "time"

    "github.com/go-board/x-go/xtime"
)

func main() {
    // 标准库
    t, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z07:00")
    println(t, err)
    // vs xtime
    t, err = xtime.ParseAny("2018/11/12")
    println(t, err)
}
```

## duration 时间度量

``` go
package main

import (
    "github.com/go-board/x-go/xtime"
)

func main() {
    seconds := xtime.MeasureDuration(func () {
        // do lots of work
    })
    println(seconds)
}
```

## relative time 相对时间计算

``` go
package main

import (
    "time"

    "github.com/go-board/x-go/xtime"
)

func main()  {
    xtime.StartOfDay(time.Now()) // 给定时间对应日期的开始时间
    xtime.EndOfDay(time.Now()) // 给定时间对应日期的结束时间
    xtime.StartOfToday() // 今天的开始时间
    xtime.EndOfToday() // 今天的结束时间
    xtime.UTC() // 当前UTC时间
    xtime.Local() // 当前的本地时间
    xtime.Yesterday() // 昨天
    xtime.Tomorrow() // 明天
    xtime.NowMillis() // 当前时间的毫秒单位
}
```
