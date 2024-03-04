# GO常用函数库2

### time 包

- 场景：处理时间和日期相关的操作，如时间格式化、解析、比较、计算等。

常用函数：

- `time.Now()`：获取当前时间。

- `time.Parse(layout, value)`：根据给定的格式解析时间字符串。

  - 参数：
    - `layout`：时间字符串的格式，例如 "2006-01-02 15:04:05"
    - `value`：要解析的时间字符串

- `time.Format(layout)`：将时间格式化为指定的布局字符串。

  - 参数：
    - `layout`：时间格式化的布局字符串，例如 "2006-01-02 15:04:05"

- `time.Duration`：表示时间间隔，可用于计算时间差。

- `time.Sleep(duration)`：让程序暂停一段时间。

  - 参数：

    - `duration`：暂停的时间间隔，例如 `time.Second` 表示 1 秒

    - ```
      const (
              Nanosecond  Duration = 1
              Microsecond          = 1000 * Nanosecond
              Millisecond          = 1000 * Microsecond
              Second               = 1000 * Millisecond
              Minute               = 60 * Second
              Hour                 = 60 * Minute
      )
      可以看到，time.Second 的值是 1000 * time.Millisecond，而 time.Millisecond 是 1000 * time.Microsecond，time.Microsecond 是 1000 * time.Nanosecond。
      
      这些常量用于表示不同时间单位的持续时间，例如纳秒（time.Nanosecond）、微秒（time.Microsecond）、毫秒（time.Millisecond）、秒（time.Second）、分钟（time.Minute）和小时（time.Hour）。
      
      使用 time.Second 可以方便地表示一秒钟的时间间隔，例如在调用 time.Sleep(time.Second) 时，程序会暂停一秒钟。
      ```

      

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Println("Current time:", now)

	// 解析时间字符串
	layout := "2006-01-02"
	dateStr := "2022-12-31"
	date, _ := time.Parse(layout, dateStr)
	fmt.Println("Parsed date:", date)

	// 格式化时间
	fmt.Println("Formatted date:", date.Format(layout))

	// 计算时间差
	duration := now.Sub(date)
	fmt.Println("Time difference:", duration)

	// 程序暂停一秒
	time.Sleep(time.Second)
	fmt.Println("Resumed after sleep")
}
```

```
Current time: 2024-03-04 10:30:15.123456789 +0000 UTC m=+0.000000001
Parsed date: 2022-12-31 00:00:00 +0000 UTC
Formatted date: 2022-12-31
Time difference: 7988h30m15.123456788s
Resumed after sleep
```



### strings 包

- 场景：处理字符串相关的操作，如拼接、分割、替换、大小写转换等。

常用函数：

- `strings.Join(strs, sep)`：将字符串切片按指定分隔符连接成一个字符串。
  - 参数：
    - `strs`：要连接的字符串切片
    - `sep`：连接字符串之间的分隔符
- `strings.Split(str, sep)`：根据指定分隔符将字符串拆分成字符串切片。
  - 参数：
    - `str`：要拆分的字符串
    - `sep`：拆分字符串的分隔符
- `strings.Replace(str, old, new, n)`：将字符串中的指定子串替换为新的子串。
  - 参数：
    - `str`：要替换的字符串
    - `old`：要被替换的子串
    - `new`：替换后的新子串
    - `n`：指定替换的次数，-1 表示全部替换
- `strings.ToLower(str)`：将字符串转换为小写。
  - 参数：
    - `str`：要转换为小写的字符串
- `strings.ToUpper(str)`：将字符串转换为大写。
  - 参数：
    - `str`：要转换为大写的字符串
- `strings.Contains(str, substr)`：判断字符串是否包含指定子串。
  - 参数：
    - `str`：要搜索的字符串
    - `substr`：要查找的子串

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串拼接
	strs := []string{"Hello", "world"}
	joined := strings.Join(strs, " ")
	fmt.Println("Joined string:", joined)

	// 字符串拆分
	str := "one,two,three,four"
	split := strings.Split(str, ",")
	fmt.Println("Split string:", split)

	// 字符串替换
	original := "Hello, World!"
	replaced := strings.Replace(original, "World", "Golang", 1)
	fmt.Println("Replaced string:", replaced)

	// 大小写转换
	lower := strings.ToLower(original)
	upper := strings.ToUpper(original)
	fmt.Println("Lowercase string:", lower)
	fmt.Println("Uppercase string:", upper)

	// 判断字符串是否包含子串
	contains := strings.Contains(original, "Hello")
	fmt.Println("Contains 'Hello':", contains)
}
```

```
Joined string: Hello world
Split string: [one two three four]
Replaced string: Hello, Golang!
Lowercase string: hello, world!
Uppercase string: HELLO, WORLD!
Contains 'Hello': true
```



### strconv 包

- 场景：处理基本类型和字符串之间的转换，如整数转字符串、字符串转整数、浮点数转字符串等。

常用函数：

- `strconv.Itoa(i)`：将整数转换为字符串。

  - 参数：
    - `i`：要转换为字符串的整数

- `strconv.Atoi(s)`：将字符串转换为整数。

  - 参数：
    - `s`：要转换为整数的字符串

- `strconv.ParseFloat(s, bitSize)`：将字符串转换为浮点数。

  - 参数：
    - `s`：要转换为浮点数的字符串
    - `bitSize`：转换后的浮点数类型的位数，可以是 32 或 64

- `strconv.FormatFloat(f, 'f', prec, bitSize)`：将浮点数转换为字符串。

  - 参数：

    - `f`：要转换为字符串的浮点数

    - `fmt`：转换格式，可以是 `'f'`、`'b'`、`'e'`、`'E'`、`'g'` 或 `'G'`

      - ```
        'f':定将浮点数转换为十进制表示法的格式
        'b'：科学计数法，例如 -123456p-78
        'e'：科学计数法，例如 -1.234560e+02
        'E'：科学计数法（大写），例如 -1.234560E+02
        'g'：根据数值的大小，自动选择 'f' 或 'e' 格式
        'G'：根据数值的大小，自动选择 'f' 或 'E' 格式
        ```

        

    - `prec`：保留的小数位数（对 `'f'`、`'e'`、`'E'` 和 `'g'` 类型有效）

    - `bitSize`：转换的浮点数类型的位数，可以是 32（`float32`）或 64（`float64`）

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 整数和字符串之间的转换
	num := 42
	str := strconv.Itoa(num)
	fmt.Println("Integer to string:", str)

	num, _ = strconv.Atoi(str)
	fmt.Println("String to integer:", num)

	// 浮点数和字符串之间的转换
	floatNum := 3.14
	str = strconv.FormatFloat(floatNum, 'f', 2, 64)
	fmt.Println("Float to string:", str)

	floatNum, _ = strconv.ParseFloat(str, 64)
	fmt.Println("String to float:", floatNum)
}
```



```
Integer to string: 42
String to integer: 42
Float to string: 3.14
String to float: 3.14
```



### encoding/json 包

$$
此处json包可以使用字节跳动的sonic（效率更高）
$$



- 场景：处理 JSON 数据的编码和解码，将结构体、切片等数据类型转换为 JSON 字符串，以及从 JSON 字符串解码为 Go 数据类型。

常用函数：

- `json.Marshal(v)`：将 Go 数据类型编码为 JSON 字符串。
  - 参数：
    - `v`：要转换为 JSON 字符串的数据结构
- `json.Unmarshal(data, v)`：将 JSON 字符串解码为 Go 数据类型。
  - 参数：
    - `data`：要解码的 JSON 字符串
    - `v`：解码后存储数据的结构体变量的指针

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 结构体到 JSON 字符串
	p := Person{Name: "Alice", Age: 25}
	jsonData, _ := json.Marshal(p)
	fmt.Println("JSON data:", string(jsonData))

	// JSON 字符串到结构体
	jsonStr := `{"name":"Bob","age":30}`
	var person Person
	json.Unmarshal([]byte(jsonStr), &person)
	fmt.Println("Decoded person:", person)
}
```

```
JSON data: {"name":"Alice","age":25}
Decoded person: {Bob 30}
```

