# GO常用函数库2

### time 包

- 场景：处理时间和日期相关的操作，如时间格式化、解析、比较、计算等。

常用函数：

- `time.Now()`：获取当前时间。
- `time.Parse(layout, value)`：根据给定的格式解析时间字符串。
- `time.Format(layout)`：将时间格式化为指定的布局字符串。
- `time.Duration`：表示时间间隔，可用于计算时间差。
- `time.Sleep(duration)`：让程序暂停一段时间。

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

### strings 包

- 场景：处理字符串相关的操作，如拼接、分割、替换、大小写转换等。

常用函数：

- `strings.Join(strs, sep)`：将字符串切片按指定分隔符连接成一个字符串。
- `strings.Split(str, sep)`：根据指定分隔符将字符串拆分成字符串切片。
- `strings.Replace(str, old, new, n)`：将字符串中的指定子串替换为新的子串。
- `strings.ToLower(str)`：将字符串转换为小写。
- `strings.ToUpper(str)`：将字符串转换为大写。
- `strings.Contains(str, substr)`：判断字符串是否包含指定子串。

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

### strconv 包

- 场景：处理基本类型和字符串之间的转换，如整数转字符串、字符串转整数、浮点数转字符串等。

常用函数：

- `strconv.Itoa(i)`：将整数转换为字符串。
- `strconv.Atoi(s)`：将字符串转换为整数。
- `strconv.ParseFloat(s, bitSize)`：将字符串转换为浮点数。
- `strconv.FormatFloat(f, 'f', prec, bitSize)`：将浮点数转换为字符串。

```
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

### encoding/json 包

- 场景：处理 JSON 数据的编码和解码，将结构体、切片等数据类型转换为 JSON 字符串，以及从 JSON 字符串解码为 Go 数据类型。

常用函数：

- `json.Marshal(v)`：将 Go 数据类型编码为 JSON 字符串。
- `json.Unmarshal(data, v)`：将 JSON 字符串解码为 Go 数据类型。

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