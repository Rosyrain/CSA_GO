# GO常用函数库

### fmt包常用函数

使用时需要`import fmt`

fmt包实现了类似C语言printf和scanf的格式化I/O

#### 格式化的格式

通用范式

```perl
%v	打印值的默认格式表示。当输出结构体时，扩展标志（%+v）会添加字段名
%#v	打印值的Go语法表示
%T	打印值的类型的Go语法表示
%%	打印百分号,比如说在我们字符串中有一个%s要打印,就可以用%%s就行了
```

```go
p := Person{Name: "Alice", Age: 25}

	// 打印默认格式表示
	fmt.Printf("%v\n", p)

	// 打印带字段名的格式表示
	fmt.Printf("%+v\n", p)

	// 打印Go语法表示
	fmt.Printf("%#v\n", p)

	// 打印值的类型
	fmt.Printf("%T\n", p)

	// 打印百分号
	fmt.Printf("Printing a percentage symbol: %%\n")
```

```
{Alice 25}
{Name:Alice Age:25}
main.Person{Name:"Alice", Age:25}
main.Person
Printing a percentage symbol: %
```

布尔值:

```csharp
%t	单词true或false
```

整数:

```perl
%b	表示为二进制
%c	该值对应的unicode码值
%d	表示为十进制
%o	表示为八进制
%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
%x	表示为十六进制，使用a-f
%X	表示为十六进制，使用A-F
%U	表示为Unicode格式: U+1234，等价于"U+%04X"
```

浮点数、复数的两个组分:

```perl
%b	无小数部分、二进制指数的科学计数法，如-123456p-78；参见strconv.FormatFloat %e	科学计数法，如-1234.456e+78 %E	科学计数法，如-1234.456E+78 %f	有小数部分但无指数部分，如123.456 %F	等价于%f %g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
```

字符串和[]byte:

```less
%s	直接输出字符串或者[]byte %q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
%x	每个字节用两字符十六进制数表示（使用a-f）
%X	每个字节用两字符十六进制数表示（使用A-F）
```

指针:

```less
%p	表示为十六进制，并加上前导的0x
```

```go
// 布尔值
    fmt.Printf("%t\n", true) // true

    // 整数
    fmt.Printf("%d\n", 123)     // 123
    fmt.Printf("%b\n", 123)     // 1111011（二进制）
    fmt.Printf("%o\n", 123)     // 173（八进制）
    fmt.Printf("%x\n", 123)     // 7b（小写十六进制）
    fmt.Printf("%X\n", 123)     // 7B（大写十六进制）
    fmt.Printf("%c\n", 65)      // A（对应的Unicode字符）
    fmt.Printf("%q\n", 65)      // 'A'（带单引号的字符字面值）

    // 浮点数
    fmt.Printf("%f\n", 3.14)    // 3.140000
    fmt.Printf("%.2f\n", 3.14)  // 3.14（保留两位小数）
    fmt.Printf("%e\n", 3.14)    // 3.140000e+00（科学计数法）
    fmt.Printf("%E\n", 3.14)    // 3.140000E+00（科学计数法，大写E表示指数）
    fmt.Printf("%g\n", 3.14)    // 3.14（根据情况选择%f或%e格式）
    fmt.Printf("%G\n", 3.14)    // 3.14（根据情况选择%f或%E格式）

    // 字符串
    fmt.Printf("%s\n", "Hello") // Hello
    fmt.Printf("%q\n", "Hello") // "Hello"（带双引号的字符串字面值）

    // 指针
    a := 42
    fmt.Printf("%p\n", &a)     // 0xc0000140b8（指针的十六进制表示）

    // 宽度和精度
    fmt.Printf("%10d\n", 123)    //        123（最小宽度为10，右对齐）
    fmt.Printf("%-10d\n", 123)   // 123       （最小宽度为10，左对齐）
    fmt.Printf("%.2f\n", 3.14159) // 3.14（保留两位小数）

    // 百分号
    fmt.Printf("Printing a percentage symbol: %%\n") // Printing a percentage symbol: %
```



#### fmt.Println()函数 主要用来打印数据

采用默认格式将其参数格式化并写入标准输出。总是会在相邻参数的输出之间添加空格并在输出结束后添加换行符。

返回写入的字节数和遇到的任何错误。

```go
func Println(a ...interface{}) (n int, err error)
```

#### fmt.Errorf()函数 主要用来打印错误信息,但是一般我们都用error包里的打印

Errorf根据format参数生成格式化字符串并返回一个包含该字符串的错误。

```go
func Errorf(format string, a ...interface{}) error
```

#### fmt.Sprintf()函数 主要用来格式化字符串

根据format参数生成格式化的字符串并返回该字符串

```go
func Sprintf(format string, a ...interface{}) string
```

#### fmt.Fprint()函数 主要用来将数据写入到缓存中,F开头的都这样

采用默认格式将其参数格式化并写入w。如果两个相邻的参数都不是字符串，会在它们的输出之间添加空格。返回写入的字节数和遇到的任何错误。

```go
func Fprint(w io.Writer, a ...interface{}) (n int, err error)
```

代码案例

```go
package main
 
import (
	"fmt"
	"os"
	"bufio"
)
 
func main() {
 
	//os.Stdout,在命令行输出内容
	fmt.Fprintf(os.Stdout, "%s\n", "你好,直接在命令行打印数据")
	//使用bufio.NewWriter开启内存空间
	buf := bufio.NewWriter(os.Stdout)
	//想开辟的内存中写入数据
	fmt.Fprintf(buf, "%s\n", "你好,我将数据写入到了buffer中")
	//输出刚才开辟的内存中的所有数据
	buf.Flush()
}
 
```

#### fmt.Sprint()函数 可以将所有内容串联为一个字符串

采用默认格式将其参数格式化，串联所有输出生成并返回一个字符串。如果两个相邻的参数都不是字符串，会在它们的输出之间添加空格。

```go
func Sprint(a ...interface{}) string
```

代码案例

```go
package main
 
import "fmt"
 
func main() {
 
	sprint := fmt.Sprint("aaa", "bbb", "ccc")
	fmt.Println(sprint) //ssaaaa
}
```

#### fmt.Scanning()系列函数

- Scan、Scanf和Scanln从标准输入os.Stdin读取文本；
- Fscan、Fscanf、Fscanln从指定的io.Reader接口读取文本；
- Sscan、Sscanf、Sscanln从一个参数字符串读取文本。

canln、Fscanln、Sscanln会在读取到换行时停止，并要求一次提供一行所有条目；

Scanf、Fscanf、Sscanf只有在格式化文本末端有换行时会读取到换行为止；

其他函数会将换行视为空白。

Scanf、Fscanf、Sscanf会根据格式字符串解析参数，类似Printf。例如%x会读取一个十六进制的整数，%v会按对应值的默认格式读取。格式规则类似Printf，有如下区别:

```perl
%p 未实现
%T 未实现
%e %E %f %F %g %G 效果相同，用于读取浮点数或复数类型
%s %v 用在字符串时会读取空白分隔的一个片段
flag '#'和'+' 未实现  
```

### os包常用函数

#### 文件和进程相关

使用时需要`import os`

- os.Args 读取命令行参数,第一个参数是程序名
- os.Stdin 标准输入的文件描述符,不会阻塞命令行等待输入,需要配合ioutil.ReadAll(os.Stdin)读取内存,且需要用管道符号比如:nihao | go run ./nihao.go
- os.Stdout 标准输出的描述符
- os.Stderr 标准错误输出的描述符
- os.Hostname() 获取主机名
- os.Getwd() 获取当前目录
- os.Getpid() 获取当前进程ID
- os.Getpid() 获取当前进程的父进程ID
- os.Exit() 退出当前进程
- os.Getenv("GOROOT") 获取环境变量的值
- os.Setenv("GOPATH","/user/liuhao/go") 设置环境变量
- os.Mkdir() 创建目录,不能创建多级
- os.MkdirAll() 创建多级目录
- os.Remove() 只能删除一个空的目录或一个文件
- os.RemoveAll() 可以强制删除目录以及目录汇中的文件
- os.Rename() 重命名文件
- os.Chmod() 修改文件权限
- os.Chown() 修改文件所有者
- os.Open() 打开一个文件句柄用于读取文件
- os.Create() 创建文件
- os.OpenFile() 打开文件句柄,大多数调用者都应用Open或Create代替本函数
  - O_RDONLY 只读模式打开文件
  - O_WRONLY 只写模式打开文件
  - O_RDWR 读写模式打开文件
  - O_APPEND 写操作时将数据附加到文件尾部
  - O_CREATE 如果不存在将创建一个新文件
  - O_EXCL 和O_CREATE配合使用，文件必须不存在
  - O_SYNC 打开文件用于同步I/O
  - O_TRUNC 如果可能，打开时清空文件

代码案例

```go
package main
 
import (
	"os"
	"fmt"
	"io/ioutil"
)
 
func main() {
	//从标准输入获取用户输入的数据
	bytes, _ := ioutil.ReadAll(os.Stdin)
 
	if len(bytes) > 0 {
		fmt.Println("Something on STDIN: " + string(bytes))
	} else {
		fmt.Println("Nothing on STDIN")
	}
}
```

文件读取案例

```go
package main
 
import (
	"bufio"
	"io"
	"os"
	"fmt"
	"io/ioutil"
)
 
//逐行读取有的时候真的很方便，性能可能慢一些，但是仅占用极少的内存空间。
func ReadLine(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
 
	buffRead := bufio.NewReader(f)
	for {
		line, err := buffRead.ReadBytes('\n')
		os.Stdout.Write(line) 放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
 
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				fmt.Println("文件逐行读完")
			}
			panic(err)
		}
	}
}
 
func ReadBlock(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
 
	buf := make([]byte, 1024) //一次读取多少个字节
	bfRd := bufio.NewReader(f)
	for {
		n, err := bfRd.Read(buf)
		os.Stdout.Write(buf[:n]) // n 是成功读取字节数
 
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				fmt.Println("文件分块读取完毕")
			}
			panic(err)
		}
	}
}
 
func ReadAll(filePth string) {
	f, err := os.Open(filePth)
	if err != nil {
		panic(err)
	}
 
	content, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
 
	fmt.Println(content, "一次性读取文件完毕")
}
 
func main() {
 
	filePath := "nihao.txt"
	ReadLine(filePath)
	ReadBlock(filePath)
	ReadAll(filePath)
}
```

#### os/exec执行外部命令

使用时需要`import os/exec`

exec包执行外部命令，它将os.StartProcess进行包装使得它更容易映射到stdin和stdout，并且利用pipe连接i/o．

- exec.Cmd{} 表示一个正在准备或者正在运行的外部命令
- exec.LookPath()函数 在系统变量中查找命令保存在哪里
- cmd := exec.Command() 准备执行外部命令相当于句柄一样的感觉
- cmd.CombinedOutput() 在exec.Command之后调用该函数执行命令并返回标准输出和标准错误
- cmd.Output() 执行命令并返回标准输出(Output()和CombinedOutput()不能够同时使用)
- cmd.Run() 开始指定命令并且等待他执行结束，如果命令能够成功执行完毕，则返回nil，否则的话边会产生错误
- cmd.Start() 使某个命令开始执行，但是并不等到他执行结束，这点和Run命令有区别．然后使用Wait方法等待命令执行完毕并且释放响应的资源,注: 一个command只能使用Start()或者Run()中的一个启动命令，不能两个同时使用．
- cmd.StdoutPipe() 方法返回一个在命令Start后与命令标准输出关联的管道,Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
- cmd.StderrPipe() 方法返回一个在命令Start后与命令标准错误输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
- cmd.StdinPipe() 方法返回一个在命令Start后与命令标准输入关联的管道,Wait方法获知命令结束后会关闭这个管道。必要时调用者可以调用Close方法来强行关闭管道，例如命令在输入关闭后才会执行返回时需要显式关闭管道。
- cmd.Wait() 等待command退出，他必须和Start一起使用，如果命令能够顺利执行完并顺利退出则返回nil，否则的话便会返回error，其中Wait会是放掉所有与cmd命令相关的资源

代码案例

```go
 
package main
 
import (
	"os/exec"
	"fmt"
	"os"
	"io/ioutil"
)
 
func LookPath() {
	f, err := exec.LookPath("ps")
	if err != nil {
		panic(err)
	}
	fmt.Println(f) //  /bin/ps
}
 
func ExecPs() {
	cmd := exec.Command("ps")
	//执行上面的ps命令,并返回标准输出和错误输出
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out)) //很多,就不写了,输出结果和在系统中执行ps一样
}
 
func ExecRun() {
	cmd := exec.Command("ps")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println(cmd.Start()) //exec: already started,不过在此之前还是会正常执行ps命令
 
}
 
func ExecPipe() {
	cmd := exec.Command("cat")
	//方法返回一个在命令Start后与命令标准输入关联的管道
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	//写入管道内容
	_, err = stdin.Write([]byte("nihao.txt"))
	if err != nil {
		panic(err)
	}
	stdin.Close()
	cmd.Stdout = os.Stdout //终端标准输出tmp.txt
	cmd.Start()
}
 
func ExecPipe1() {
	cmd := exec.Command("ps")
	//方法返回一个在命令start后与命令标准输出关联的管道
	stdout, err := cmd.StdoutPipe()
	cmd.Start()
	//读取管道内容
	content, err := ioutil.ReadAll(stdout)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content)) //很多,就不写了,输出结果和在系统中执行ps一样
}
func main() {
 
	//LookPath()
	//ExecPs()
	//ExecRun()
	//下面都是管道的使用
	//ExecPipe()
	ExecPipe1()
}
```

### io包常用函数

`types.Buffer`类型也实现了Writer方法。

应该说其实很多包中,为了方便使用,都在不同的NewXXX()函数中实现了Writer方法,我们要从这些方法中获取数据要么用他们已经封装好的获取方式,也可以使用io的Reader方法读取

#### io/ioutil

使用时需要`import io/ioutil`

- ioutil.ReadAll() 从r读取数据直到EOF或遇到error，返回读取的数据和遇到的错误。
- ioutil.ReadFile() 从filename指定的整个文件中读取数据并返回文件的内容。成功的调用返回的err为nil而非EOF。
- ioutil.WriteFile() 函数向filename指定的文件中写入数据。如果文件不存在将按给出的权限创建文件，否则在写入数据之前清空文件。
- ioutil.ReadDir() 返回dirname指定的目录的目录信息的有序列表。
- ioutil.TempDir() 在dir目录里创建一个新的、使用prfix作为前缀的临时文件夹，并返回文件夹的路径。
- ioutil.TempFile() 在dir目录下创建一个新的、使用prefix为前缀的临时文件，以读写模式打开该文件并返回os.File指针。

### strconv包常用函数

使用时需要`import strconv`

- strconv.ParseBool() 将字符串转换为布尔值,真值: 1, t, T, TRUE, true, True; 假值: 0, f, F, FALSE, false, False.

```go
//str 是要转换的字符串
func ParseBool(str string) (value bool, err error) 
```

- strconv.ParseInt() 将字符串转换为 int 类型

```go
// s: 要转换的字符串
// base: 进位制（2 进制到 36 进制）
// bitSize: 指定整数类型（0:int、8:int8、16:int16、32:int32、64:int64）
// 返回转换后的结果和转换时遇到的错误
// 如果 base 为 0，则根据字符串的前缀判断进位制（0x:16，0:8，其它:10）
func ParseInt(s string, base int, bitSize int) (i int64, err error)
```

- strconv.ParseUint() 同 ParseInt 一样，只不过返回 uint 类型整数

```go
func ParseUint(s string, base int, bitSize int) (n uint64, err error)
```

- strconv.ParseFloat() 将字符串转换为浮点型

```go
//s 是要转换的字符串
//bitSize指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32），64是float64；
func ParseFloat(s string, bitSize int) (f float64, err error)
```

- strconv.FormatBool() 将布尔值转换为字符串 "true" 或 "false"

```go
func FormatBool(b bool) string
```

- strconv.FormatInt() 将字符串转换为整型

```go
// FormatUint 将 int 型整数 i 转换为字符串形式
// base: 进位制（2 进制到 36 进制）
// 大于 10 进制的数，返回值使用小写字母 'a' 到 'z'
func FormatInt(i int64, base int) string
```

- strconv.FormatUint() 和FormatInt一样,只不过返回的时uint

```go
// FormatUint 将 uint 型整数 i 转换为字符串形式
// base: 进位制（2 进制到 36 进制）
// 大于 10 进制的数，返回值使用小写字母 'a' 到 'z'
func FormatUint(i uint64, base int) string
```

- strconv.FormatFloat() 将浮点数转换为字符串值

```go
// f: 要转换的浮点数
// fmt: 格式标记（b、e、E、f、g、G）
// prec: 精度（数字部分的长度，不包括指数部分）
// bitSize: 指定浮点类型（32:float32、64:float64）
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
```

- strconv.Atoi() 相当于 ParseInt(s, 10, 0)将字符串转换为10进制的整型

```go
//Atoi是ParseInt(s, 10, 0)的简写。
func Atoi(s string) (i int, err error)
```

- strconv.Itoa() 相当于 FormatInt(i, 10)将整型转换为10进制的字符串

```go
//Itoa是FormatInt(i, 10) 的简写。
func Itoa(i int) string
```

- strconv.AppendBool()

```go
//等价于append(dst, FormatBool(b)...)
// AppendBool 将布尔值 b 转换为字符串 "true" 或 "false"
// 然后将结果追加到 dst 的尾部，返回追加后的 []byte
func AppendBool(dst []byte, b bool) []byte
```

- strconv.AppendInt()

```go
//等价于append(dst, FormatUint(I, base)...)
// AppendInt 将 int 型整数 i 转换为字符串形式，并追加到 dst 的尾部
// i: 要转换的字符串
// base: 进位制
// 返回追加后的 []byte
func AppendInt(dst []byte, i int64, base int) []byte
```

- strconv.AppendUint()

```go
//等价于append(dst, FormatUint(I, base)...)
// AppendUint 将 uint 型整数 i 转换为字符串形式，并追加到 dst 的尾部
// i: 要转换的字符串
// base: 进位制
// 返回追加后的 []byte
func AppendUint(dst []byte, i uint64, base int) []byte
```

- strconv.AppendFloat()

```go
//等价于append(dst, FormatFloat(f, fmt, prec, bitSize)...)
// AppendFloat 将浮点数 f 转换为字符串值，并将转换结果追加到 dst 的尾部
// 返回追加后的 []byte
func AppendFloat(dst []byte, f float64, fmt byte, prec int, bitSize int) []byte
```

- strconv.AppendQuote()

```go
//等价于append(dst, Quote(s)...)
// AppendQuote 将字符串 s 转换为“双引号”引起来的字符串，
// 并将结果追加到 dst 的尾部，返回追加后的 []byte
// 其中的特殊字符将被转换为“转义字符”
func AppendQuote(dst []byte, s string) []byte
```

- strconv.AppendQuoteToASCII()

```go
//等价于append(dst, QuoteToASCII(s)...)
// AppendQuoteToASCII 将字符串 s 转换为“双引号”引起来的 ASCII 字符串，
// 并将结果追加到 dst 的尾部，返回追加后的 []byte
// “非 ASCII 字符”和“特殊字符”将被转换为“转义字符”
func AppendQuoteToASCII(dst []byte, s string) []byte
```

- strconv.AppendQuoteRune()

```go
//等价于append(dst, QuoteRune(r)...)
// AppendQuoteRune 将 Unicode 字符转换为“单引号”引起来的字符串，
// 并将结果追加到 dst 的尾部，返回追加后的 []byte
// “特殊字符”将被转换为“转义字符”
func AppendQuoteRune(dst []byte, r rune) []byte
```

- strconv.AppendQuoteRuneToASCII()

```go
//等价于append(dst, QuoteRuneToASCII(r)...)
// AppendQuoteRune 将 Unicode 字符转换为“单引号”引起来的 ASCII 字符串，
// 并将结果追加到 dst 的尾部，返回追加后的 []byte
// “非 ASCII 字符”和“特殊字符”将被转换为“转义字符”
func AppendQuoteRuneToASCII(dst []byte, r rune) []byte
```

### strings包常用函数

- strings.HasPrefix(s string, prefix string) bool: 判断字符串s是否以prefix开头
- strings.HasSuffix(s string, suffix string) bool: 判断字符串s是否以suffix结尾。
- strings.Index(s string, str string) int: 判断str在s中首次出现的位置，如果没有出现，则返回-1
- strings.LastIndex(s string, str string) int: 判断str在s中最后出现的位置，如果没有出现，则返回-1
- strings.Replace(str string, old string, new string, n int): 字符串替换
- strings.Count(s, sep string) int: 返回字符串s中有几个不重复的sep子串。
- strings.Repeat(str string, count int)string: 重复count次str
- strings.ToLower(str string)string: 转为小写
- strings.ToUpper(str string)string: 转为大写
- strings.TrimSpace(str string): 去掉字符串首尾空白字符
- strings.Trim(str string, cut string): 去掉字符串首尾cut字符
- strings.TrimLeft(str string, cut string): 去掉字符串首cut字符
- strings.TrimRight(str string, cut string): 去掉字符串首cut字符
- strings.Field(str string): 返回str空格分隔的所有子串的slice
- strings.Split(str string, split string): 返回str split分隔的所有子串的slice
- strings.Join(s1 []string, sep string): 用sep把s1中的所有元素链接起来
- strings.Itoa(i int): 把一个整数i转成字符串
- strings.Atoi(str string)(int, error): 把一个字符串转成整数
- strings.Contains(s, substr string) bool: 判断字符串s是否包含子串substr。
- strings.ContainsRune(s string, r rune) bool: 判断字符串s是否包含utf-8码值r。
- r := strings.NewReader(s string) *Reader: 创建一个从s读取数据的Reader,同时该Reader有多个方法可用来访问NewReader中的数据。本函数类似bytes.NewBufferString，但是更有效率，且为只读的。
  - r.Len() 返回未读的字符串长度 `func (r *Reader) Len() int`
  - r.Size() 返回字符串的长度
  - r.Read() 读取字符串信息并保存到指定的byte类型的变量中 `func (r *Reader) Read(b []byte) (n int, err error)`
  - r.ReadAt()读取偏移off字节后的剩余信息到b中（需要注意的是，ReadAt函数不会影响Len的数值，和Read的数值，off不能为负数，不能大于size（）的长度）`func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)`
  - r.ReadByte() 从当前已读取位置继续读取一个字节
  - r.UnreadByte() 将当前已读取位置回退一位，当前位置的字节标记成未读取字节
  - 前面有ReadAt方法可以将字符串偏移多少位读取剩下的字符串内容，但是该方法不会影响正在用Read方法读取的内容，如果相对Read方法读取的内容做偏移就可以使用seek方法， offset是偏移的位置，whence是偏移起始位置，支持三种位置（io.SeekStart起始位，io.SeekCurrent当前位，io.SeekEnd末位）。需要注意的是offset可以未负数，当时偏移起始位 与offset相加得到的值不能小于0或者大于size()的长度

代码案例

```go
package main

import (
	"strings"
	"fmt"
)

func main() {
	read := strings.NewReader("大家好我是刘")
	fmt.Println(read.Len())             //返回未读取的长度
	fmt.Println(read.Size())            //返回字符串的长度
	buff := make([]byte, 18)            //开辟空间为后面的读取用,这里我用18,其实可以直接用rea.Len()或者read.Size()
	read.Read(buff)                     //通过指定的空间,读取指定长度的数据
	fmt.Println("中文字符串:", string(buff)) //打印空间中的内容
	buffAt := make([]byte, 18)
	read.ReadAt(buffAt, 3) //读取第三字节以后的数据, 注意啊,这里的如果不是3,6,9之类的,会出现方框.以这个中文3字节码!
	enRead := strings.NewReader("abcdefghijk")
	b, _ := enRead.ReadByte() //向后读取一个字节,这里是中文所有会输出方框,可以用英文
	b, _ = enRead.ReadByte()  //再向后读取一个字节,这里是中文所有会输出方框,可以用英文
	b, _ = enRead.ReadByte()  //再向后读取一个字节,这里是中文所有会输出方框,可以用英文
	fmt.Println("英文字符串:", string(b))
	fmt.Println(int(enRead.Size()) - enRead.Len()) //通过总字节数-还剩字节数,计算读了多少字节
	enRead.UnreadByte()                            //向前一个字节
	fmt.Println(int(enRead.Size()) - enRead.Len()) //通过总字节数-还剩字节数,计算读了多少字节
}
```

### bytes包常用函数

- func NewBuffer(buf []byte) *Buffer NewBuffer: 使用buf作为初始内容创建并初始化一个Buffer,大多数情况下，new(Buffer)（或只是声明一个Buffer类型变量）就足以初始化一个Buffer了。
- func NewBufferString(s string) *Buffer: 使用s作为初始内容创建并初始化一个Buffer。本函数用于创建一个用于读取已存在数据的buffer。
  - Reset(): 重设缓冲，因此会丢弃全部内容，等价于b.Truncate(0)
  - Len() int: 返回缓冲中未读取部分的字节长度；b.Len() == len(b.Bytes())。
  - Bytes(): 返回未读取部分字节数据的切片，len(b.Bytes()) == b.Len()。如果中间没有调用其他方法，修改返回的切片的内容会直接改变Buffer的内容。
  - String(): 将未读取部分的字节数据作为字符串返回，如果b是nil指针，会返回"<nil>"。
  - Truncate(n int): 丢弃缓冲中除前n字节数据外的其它数据，如果n小于零或者大于缓冲容量将panic。
  - Grow(n int): 必要时会增加缓冲的容量，以保证n字节的剩余空间。调用Grow(n)后至少可以向缓冲中写入n字节数据而无需申请内存。如果n小于零或者不能增加容量都会panic。
  - Read(p []byte): Read方法从缓冲中读取数据直到缓冲中没有数据或者读取了len(p)字节数据，将读取的数据写入p。返回值n是读取的字节数，除非缓冲中完全没有数据可以读取并写入p，此时返回值err为io.EOF；否则err总是nil。
  - Next(n int): 返回未读取部分前n字节数据的切片，并且移动读取位置，就像调用了Read方法一样。如果缓冲内数据不足，会返回整个数据的切片。切片只在下一次调用b的读/写方法前才合法。
  - ReadByte() : ReadByte读取并返回缓冲中的下一个字节。如果没有数据可用，返回值err为io.EOF。
  - ReadRune(): 读取并返回缓冲中的下一个utf-8码值。如果没有数据可用，返回值err为io.EOF。如果缓冲中的数据是错误的utf-8编码，本方法会吃掉一字节并返回(U+FFFD, 1, nil)。
  - Write(): 将p的内容写入缓冲中，如必要会增加缓冲容量。返回值n为len(p)，err总是nil。如果缓冲变得太大，Write会采用错误值ErrTooLarge引发panic。
  - WriteString(): 将字节c写入缓冲中，如必要会增加缓冲容量。返回值总是nil，但仍保留以匹配bufio.Writer的WriteByte方法。如果缓冲太大，WriteByte会采用错误值ErrTooLarge引发panic。
  - WriteByte(): 将字节c写入缓冲中，如必要会增加缓冲容量。返回值总是nil，但仍保留以匹配bufio.Writer的WriteByte方法。如果缓冲太大，WriteByte会采用错误值ErrTooLarge引发panic。
  - WriteRune(r rune) (n int, err error): WriteByte将unicode码值r的utf-8编码写入缓冲中，如必要会增加缓冲容量。返回值总是nil，但仍保留以匹配bufio.Writer的WriteRune方法。如果缓冲太大，WriteRune会采用错误值ErrTooLarge引发panic。
  - ReadFrom(r io.Reader) (n int64, err error): 从r中读取数据直到结束并将读取的数据写入缓冲中，如必要会增加缓冲容量。返回值n为从r读取并写入b的字节数；会返回读取时遇到的除了io.EOF之外的错误。如果缓冲太大，ReadFrom会采用错误值ErrTooLarge引发panic。
  - WriteTo(w io.Writer) (n int64, err error): 从缓冲中读取数据直到缓冲内没有数据或遇到错误，并将这些数据写入w。返回值n为从b读取并写入w的字节数；返回值总是可以无溢出的写入int类型，但为了匹配io.WriterTo接口设为int64类型。从b读取是遇到的非io.EOF错误及写入w时遇到的错误都会终止本方法并返回该错误。
- NewReader(b []byte) *Reader: 这个基本上和strings里面那个差不多,这里就不赘述了