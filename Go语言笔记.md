### Go知识整理

​	努力的意义是经营好自己的每天，追求自己想要的目标(该阶段比较肤浅，买车买房)

#### 一、语言结构

Go语言的基础组成有以下几个部分：

- 包声明
- 引入包
- 函数
- 变量
- 语句&表达式
- 注释

接下来简单代码输出“Hello World！”：

```go
package main 

import "fmt"

func main(){
	//这是一个简单的程序
    fmt.printlm("Hello,World!")
}
```

**解析上程序的各个部分：**

1. 第一行代码package main定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。package main表示一个可执行的程序，每个Go应用程序都包含一个名为mian的包。

2. 下一行import “fmt”告诉Go编辑器这个程序需要使用fmt包（的函数或其他元素），fmt包实现了格式化IO（输入/输出）的函数。

3. 下一行func main（）是程序开始执行的函数。main函数是每一个可执行程序必须包含的，一般来说是在启动后第一个执行行的函数（如果有init（）函数则会先执行该函数）。

4. 下一行/\*...\*/是注释，在程序执行是被忽略。单行注释是最常见的注释形式，可以在任何地方使用以//开头的单行注释，均已/\*开头，并以\*/结尾，不可嵌套使用，多行注释一般用于包的文档描述或注释成块的代码片段。

5. 下一行fmt.Println（,,,）可以将字符串输出到控制台，并在最后自动增加到换行字符/n

   使用fmt.Print（"Hello,World!\n"）可以得到相同的结果。
   Print和println这两个函数也支持使用标量，如：fmt.Println(arr).如果没有指定，他们会默认的打印格式将变量arr输出到控制台。

6. 当标识符（包含常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是空间并可用的（像面向对象语言中的producted）。

#### 二、Go语言基础语法

Go标记

Go程序可以有多个标记组成，可以是关键字，标识符，常量，字符串，符号。以下Go语句由6个标记组成：

```go
fmt.Println("Hello,World!")
```

6个标记分别是

```go
1.fmt	2..	3.Println	4.(		5."Hello,World!"	6.)
```

**行分隔符**

在Go程序中，一行代表一个语句结束。每个语句不需要像C家族中的其他语言一样以分号" ; "结尾，因为这些工作都将由Go编译器自动完成。

如果你打算将多个语句写在同一行，它们则必须使用" ; "认为区分，但在实际开发中我们并鼓励这种做法。下面为两个语句：

```go
fmt.Println("Hello,World!")
fmt.Println("这是另一个语句!")
```

**标识符**

标识符用来命名变量、类型等程序实体。一个标识符实际上就是一个或是多个字母（A~Z和a~z）数字（0~9）、下划线_组成的序列，**但是第一个字符必须是字母或下划线，不能是数字。**

以下是**有效**标识符的例子：

```go
mahes	kumar	abc	move_name	a_123
myname50	_temp	j	a232	retVal
```

以下是**无效**的标识符例子：

```go
2ad(以数字开头)
case（Go语言的关键字）
a+b（运算符是不允许的）
```

**字符串连接**

Go语言的字符串连接可以通过`+`实现：

```go
实例
package main
import "fmt"
func main(){
    fmt.Println("Google" + "Runoob")
}
```

以上实例输出的结果为：

```
GoogleRunoob
```

**关键字**

下面列举Go代码中会使用到的25个关键字或保留字

| break    | default     | func   | interface | select |
| -------- | ----------- | ------ | --------- | ------ |
| case     | defer       | go     | map       | struct |
| chan     | else        | goto   | package   | switch |
| const    | fallthrough | if     | range     | type   |
| continue | for         | import | return    | var    |

除了以上介绍的这些关键字，Go语言还有36预定义标识符：

| append | bool    | byte    | cap     | close  | complex | complex64 | complex128 | uint16  |
| ------ | ------- | ------- | ------- | ------ | ------- | --------- | ---------- | ------- |
| copy   | false   | float32 | float64 | imag   | int     | int8      | int16      | uint32  |
| int32  | int64   | iota    | len     | make   | new     | nil       | panic      | uint64  |
| print  | println | real    | recover | string | true    | uint      | uint8      | uintptr |

程序一般由关键字、常量、变量、运算符、类型和函数组成。

程序中可能会使用到这些分隔符：括号（），中括号[]和大括号{}。

程序中可能会使用到这些标点符号：`.`、`,`、`;`、`:`和`...`。

**格式化字符串**

Go语言中使用`fmt.Sprintf`或`fmt.Printf`格式化字符串并赋值给新字符串：

- **Sprintf**根据格式化参数生成格式化的字符串**并返回该字符串**。
- **Printf**根据格式化参数生成格式化的字符串并写入标准输出

```go
Sprintf实例
package main

improt (	//表示可以导入多个包
	"fmt"
)

func main() {
    //%d  表示整型数字，%s  表示字符串
    var stockcode = 123
    var enddate = "2023-5-24"
    var url = "Code = %d&endData = %s"
    var target_url = fmt.Sprintf(url,stockcode,enddate)
    fmt.Println(target_url)
}
```

输出结果为：

```
Code = 123&endData = 2023-5-24
```

```
Printf实例
package main

import (
	"fmt"
)

func main() {
	// %d 表示整型数字， %s 表示字符串
	var stockcode = 123
	var enddate = "2023-5-23"
	var url = "Code = %d&endData = %s"
	fmt.Printf(url,stockcode,enddate)
	//或者这样输出
	fmt.Sprintf("Code = %d&endData = %s",stockcode,enddate)
}
```

输出结果为：

```
Code = 123&endData = 2023-5-24
Code = 123&endData = 2023-5-24
```

#### 三、结构类型

**派生类型:**
包括：

- (a) 指针类型（Pointer）
- (b) 数组类型
- (c) 结构化类型(struct)
- (d) Channel 类型
- (e) 函数类型
- (f) 切片类型
- (g) 接口类型（interface）
- (h) Map 类型

在 Go 中，**布尔值的类型为 bool，值是 true 或 false，默认为 false。**

```
//示例代码
var isActive bool  // 全局变量声明
var enabled, disabled = true, false  // 忽略类型的声明
func test() {
    var available bool  // 一般声明
    valid := false      // 简短声明
    available = true    // 赋值操作
}
```

字符串去除空格和换行符

```
package main  
  
import (  
    "fmt"  
    "strings"  
)  
  
func main() {  
    str := "这里是 www\n.runoob\n.com"  
    fmt.Println("-------- 原字符串 ----------")  
    fmt.Println(str)  
    // 去除空格  
    str = strings.Replace(str, " ", "", -1)  
    // 去除换行符  
    str = strings.Replace(str, "\n", "", -1)  
    fmt.Println("-------- 去除空格与换行后 ----------")  
    fmt.Println(str)  
}
```

输出结果为：

```
-------- 原字符串 ----------
这里是 www
.runoob
.com
-------- 去除空格与换行后 ----------
这里是www.runoob.com
```

#### 四、Go语言变量

变量来源于数学，是计算机语言中存储计算结果或能表示值抽象概念。

变量可以通过变量名访问。

Go语言变量名由字母、数字、下划线组成，其中首个字符不能为数字。

声明变量的一般形式是使用var关键字：

```go
var identifier1 type	//	关键字 变量名	变量类型
```

可以一次生命多个变量：

```go
var identifier1，identifier2 type

实例
package main
improt "fmt"
func main() {
    var a string = "Runoob"
    fmt.Println(a)
    
    var b,c int = 1, 2
    fmt.Println(b,c)
}
```

以上实例输出结果为：

```
Runoob
1 2
```

**变量声明**

**第一种，指定变量类型，如果没有初始化，则变量默认为零值**

```go
var	v_name v_type
v_name = value
```

零值就是变量没有做初始化时系统默认设置的值

```go
实例
package main
import "fmt"
func main() {
    
    //声明一个变量并初始化
    var a = "Runoob"
    fmt.Println(a)
    
    //没有初始化就为零值
    var b int
    fmt.Println(b)
    
    //bool 零值为 false
    var c bool
    fmt.Println(c)
}
```

以上实例执行结果为：

```
Runoob
0
false
```

- 数值类型 （包括complex64/128）为**0**
- 布尔类型为**false**
- 字符串为**""**(空字符串)
- 以下几种类型为**nil**：

```go
var a *int
var a []int
var a map[string]int
var a chan int
var a func(string) int
var a error	//	error是接口

实例
package main
import "fmt"
func main() {
    var i int
    var f float64
    var b bool
    var s string
    fmt.Printf("%v %v %v %q\n",i,f,b,s)
}
```

输出结果是：

```
0 0 false ""
```

**第二种，自行判定变量类型**

```go
var v_name = value

实例
package main
improt "fmt"
func main() {
    var d = ture
    fmt.Println(d)
}
```

输出结果是：

```
true
```

**第三种，如果变量已经被使用var声明过了，在使用`:=`声明变量，就产生编译错误（注 ：它只能被用在函数体内，而不可以用于全局变量的声明与赋值），格式：**

```go
v_name = value
```

例如：

```go
var intVal int
intVal := 1	// 这个时候会产生编译错误，因为intVal已经声明，不需要重新声明
```

一开始之间使用下面语句的话就不会编译错误：

```
intVal := 1	// 此时不会产生编译错误，因为相当于声明新的变量，:= 是一个声明语句
```

`intVal := 1`相等于：

```go
var intVal int
intVal = 1
```

可以将var f string = "Runoob"简写为 f := "Runoob"

```go
实例
package main
import "fmt"
func main() {
	f := "Runoob"	//相当于var f string = "Runoob"
    fmt.Println(f)
}
```

输出结果是：

```
Runoob
```

**多变量声明**

```go
//类型相同多个变量，非全局变量
var vname1,vname2,vname3 type
vname1,vname2,vname3 = v1, v2, v3

//知道变量类型,不显示声明类型，自动推断
var vname1,vname2,vname3 = v1, v2, v3
或者用 := 不显示声明类型，自动推断(注 ： 它只能被用在函数体内，而不可以用于全局变量的声明与赋值)
vname1,vname2,vname3 := v1, v2, v3

//这种因式分解关键字的写法一般用于声明全局变量
var (
	vname1	type
	vname2	type
)
```

实例

```go
package main
import "fmt"

var x,y int
var (	//这种因式分解关键字的写法一般用于声明全局变量
	a int
    b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main () {
    g, h := 123, "hello"
    Println(x, y, a, b, c, d, e, f, g, h)
}
```

以上实例执行结果为：

```
0 0 0 false 1 2 123 hello 123 hello
```

**注意事项**

如果在相同的代码块中，我们不可以再次对相同名称的变量使用初始化声明，例如：var a int = 20,再声明a :=20就是不被允许的，编译器会提示错误no new variables on left side of :=,但是a = 20是可以的，因为这是给相同的变量赋予一个新的值。

如果你在定义变量a之前使用它，则会得到编译错误undefined :a

如果你声明了一个局部变量却没有在相同的代码块中使用它，同样会得到编译错误，例如：

```go
实例
package main
import "fmt"

func main() {
    var a string = "abc"
    fmt.Println("hello world")
}
```

尝试编译这段代码将得到错误**a declared but not used**

此外，单纯地给a赋值也是不够的，这个值必须被使用，所以要使用这个a

```
fmt.Println("hello,world",a)	//	会移除错误
```

**空白标识符在函数返回值时的使用（有点像方法重载的意思，根据需要返回需要的参数数量）：**

空白标识符 _ 也被用于抛弃值，如值5在 `_, b = 5, 7`中被抛弃。

_ 实际上是一个只写变量，你不能得到它的值。这样做是因为Go语言中你必须使用所有被声明的变量，但是有时你并不需要使用从一个函数得到的所有返回值。

并行赋值也被用于当一个函数返回多个返回值时，比如这里的val和错误err是通过Func1函数同时得到：val, err = Func1(var1).

```go
package main

import "fmt"

func main() {
  _,numb,strs := numbers() //只获取函数返回值的后两个
  fmt.Println(numb,strs)
}

//一个可以返回多个值的函数
func numbers()(int,int,string){
  a , b , c := 1 , 2 , "str"
  return a,b,c
}
```

输出结果：

```
2 str
```



**局部变量和全局变量：**

局部变量：在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量。

全局变量：在函数体外声明的变量称之为全局变量，全局变量可以在整个包甚至外部包（被导出后）使用。

**备注：**

**1.局部变量不会一直存在，在函数被调用时存在，函数调用结束后变量就会被销毁，即生命周期。**

**2.Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。**



#### 五、Go语言常量

常量是一个简单值的标识符，在程序运行时，不会被修改的量

常量中的数据类型只可以是布尔类型、数字型（整数型、浮点型和复数）和字符串型

常量的定义格式：

```go
const identfier [type] = value
```

可以省略类型说明[type],因为编译器可以根据变量的值来推断类型

- 显式类型定义：`const b string = "abc"`
- 隐式类型定义：`const b = "abc"`

多个相同类型的声明可以简写为：

```go
const c_name1, c_name2 = value1, value2
```

以下实例演示了常量的应用：

```go
package main
import "fmt"

func main() {
	const LENGTH int = 10
    const WIDTH int = 5
    var area int
    const a, b, c = 1, false, "ste"		//多重赋值
    
    area = LENGTH * WIDTH
    fmt.Printf("面积为： %d",area)
    Println()
    Println(a, b, c)
}

```

以上实例运行结果为：

```
面积为： 50
1 false str
```

常量还可以用来作枚举：

```go
const (
	Unkown = 0
	Female = 1
	Male = 2
)
```

数字0、1和2分别代表未知性别、女性和男性。

常量可以用len(),cap(),unsafe.Sizeof()函数计算表达式的值。常量表达式中，**函数必须是内置函数，否则编译不过：**

```go
实例
package main
import (
	"unsafe"
    "fmt"
)
const(
	a = "abc"
    b = len(a)
    c = unsafe.Sizeof(a)
)
func main(){
    fmt.Println(a, b, c)
}
```

以上实例运行结果为：

```
abc 3 16
```

**iota特殊常量**

iota，特殊常量，可以认为是一个可以被编辑器修改的常量。

iota在const关键字出现时将被重置为0（const内部的第一行之前），const中美新增一行常量声明将使iota计数一次（iota可以理解为const语句块中的行索引）。

iota可以被用作枚举值：

```go
const(
	a = iota	//a = 0
	b = iota	//b = 1
	c = iota	//c = 2
)
```

第一个iota等于0，每当iota在新的一行被使用时，它的值就会被自动加1；所有a = 0，b = 1，c = 2可以简写为如下形式：

```go
const (
	a = iota	//a = 0
	b			//b = 1
	c			//c = 2
)
```

**iota用法**

```go
package main
import "fmt"

func main() {
	const (
		a = iota	//0
		b			//1
		c			//2
		d = "ha"	//独立值，iota += 1
		e			//"ha"  iota += 1
		f = 100		//iota += 1
		g			//100 iota += 1
		0 			//100 iota += 1
		h = iota	//8，恢复计数	
		i			//9
	)
	fmt.Println(a,b,c,d,e,f,g,o,h,i)
}
```

以上实例运行结果为：

```
0 1 2 ha ha 100 100 100 8 9
```

再看个有趣的iota实例：

```go
package main
import "fmt"
const(
	i = 1 << iota
	j = 3 << iota
	k
	l
)
func main() {
	fmt.Println("i=",i)
	fmt.Println("j=",j)
    fmt.Println("k=",k)
    fmt.Println("l=",l)
}
```

以上实例运行结果为：

```
i = 1
j = 6
k = 12
k = 24
```

iota表示从0开始自动加1，所以`i = 1 << 0`,`j = 3 << 1`(<< 表示左移的意思)，即：i = 1，j = 6，这没问题，关键在 k 和 l ，从输出结果看 `k = 3 << 2`  , `l = 3 << 3`

简单描述：

- **i = 1** : 左移0位，不变仍为1
- **j = 3** : 左移1位，变为二进制**110**，即6。
- **k = 3** : 左移1位，变为二进制**1100**，即12。
- **l = 3** : 左移1位，变为二进制**110**，即24。

注：`<<n==*(2^n)`。

#### **六、指针类型**

Go中，指针是一种类型，指向变量所在的内存单元(不是内存地址)。

这里展开做一个说明，在Go中，指针是一种特殊的数据类型，它存储了变量所在内存单元的地址。这里的内存地址指的是一个唯一的标识符，用来指示变量在计算机内存中的位置。而内存单元则是计算机内存中的一个存储单元，它可以存储某种特顶定类型的数据。

申明：在变量名前加上星号字符，比如`*age`,指向变量`age`所在的内存单元

具体来说：

- **内存地址：**内存地址是计算机内存中某个特定存储单元的唯一标识符。在go中，使用指针可以获取变量在内存中的地址。每个变量在内存中都有一个唯一的地址，可以通过指针来引用这个地址。Go使用十六进制表示内存地址，通常以`0x`开头.
- **内存单元：**内存单元是计算机内存中的一个存储单元，它可以存储某种特定类型的数据。内存单元通常由连续的字节组成，每个字节都有一个唯一的地址。在程序中，变量被存储在内存单元中，并且可以同各国变量名或者指针来访问这些内存单元中存储的数据。

因此，在Go中，指针存储的是变量在内存中的地址，而这个地址对应的内存单元存储了实际的数据。指针通过引用地址来访问内存的单元中的数据，从而允许对变量进行间接访问和修改。

**&(取址符)：对变量取址**

& ：获取变量在计算机内存中的地址，`&age`，取出变量`age`所在内存地址，一般地址是十六进制。

**\*(声明指针)：对指针取值**

申明指针`*age`，打印指针内存单元的值`**age`。如下面`x  *int`,`*x`就是指针对应的值。

```go
package main

import (
   "fmt"
   "reflect"
)

func main(){
   k:=40
   fmt.Println(k)
   fmt.Println(say("hello,world","lf"))
   fmt.Println(reflect.TypeOf(k))  //检查变量类型
   
   // 获取变量在计算机内存中的地址，可在变量名前面加上&字符。
   fmt.Println(&k) 
   
   // &k 引用的是变量k的值，值所在的内存地址
   showMemoryAddress(&k)  //返回的地址是相同的
}

// *int参数类型位指向整数的指针，而不是整数
func showMemoryAddress(x *int){ 
	// 本身就是指针，打印地址不需要  &这个符号，如果想使用指针指向的变量的值，而不是其内存地址，可在指针变量前面加上星号
   fmt.Println(*x)  
   return
}

func say(param,tt string) string{
   return param+"--"+tt
}
```

结果如下

```go
40
hello,world--lf
int
0xc00007e020
40
```

#### 七、Go语言条件语句

**if语法**

```go
if 布尔表达式 {
   /* 在布尔表达式为 true 时执行 */
}
```

实例

```go
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 10
 
   /* 使用 if 语句判断布尔表达式 */
   if a < 20 {
       /* 如果条件为 true 则执行以下语句 */
       fmt.Printf("a 小于 20\n" )
   }
   fmt.Printf("a 的值为 : %d\n", a)
}
```

以上代码执行结果为：

```go
a 小于 20
a 的值为 : 10
```

**switch语法**

```go
switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
```

变量val1可以是任何类型，而val1和val2则可以是同类型的任意值。类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。

![image-20230524203734265](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20230524203734265.png)

```go
实例
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var grade string = "B"
   var marks int = 90

   switch marks {
      case 90: grade = "A"
      case 80: grade = "B"
      case 50,60,70 : grade = "C"
      default: grade = "D"  
   }

   switch {
      case grade == "A" :
         fmt.Printf("优秀!\n" )    
      case grade == "B", grade == "C" :
         fmt.Printf("良好\n" )      
      case grade == "D" :
         fmt.Printf("及格\n" )      
      case grade == "F":
         fmt.Printf("不及格\n" )
      default:
         fmt.Printf("差\n" );
   }
   fmt.Printf("你的等级是 %s\n", grade );      
}
```

以上代码执行结果为：

```go
优秀!
你的等级是 A
```

**select语法**

```go
select {
	case <- channel1:
		//执行的代码
	case value := <- channel2:
		//执行的代码
	case channel3 <- value:
		//执行的代码
		
	...
		
	default:
	//所有通道都没有准备好，执行的代码
}
```

以下描述了select语句的语法：

- 每个case都必须是一个通道
- 所有的channel表达式都会被求值
- 所有发送的表达式都会被求值
- 如果任意某个通道可以进行，它就执行，其他被忽略
- 如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行。
  否则：
  1. 如果有defualt子句，则执行该语句。
  2. 如果没有default子句，select将阻塞，直到某个通道可以运行；

```go
实例
package main

import (
	"fmt"
    "time"
)
func main() {
    c1 := make(chan string)		//跟Java创建开启一个线程，线程名c1
    c2 := make(chan string)		//跟Java创建开启一个线程，线程名c2
    go func(){		//运行线程
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()
    
    for i := 0;i < 2;i++ {
        select {
            case msg1 := <- c1:         
            fmt.Println("received",msg1)
            case msg2 := <- c2:
            fmt.Println("received",msg2)
        }
    }
}
```

以上实例中，我们创建了两个通道c1和c2。
select语句等待两个通道的数据。如果接收到c1的数据，就会打印"received one";如果接收到c2的数据，就会打印"received two"。
以下实例中，我们定义两个通道，并启动了两个协程(Goroutine)从两个通道中获取数据。在main函数中，我们使用select语句在这两个通道中进行非阻塞的选择，如果两个通道都没有可用数据，就执行default中的语句。
以下实例执行后会不断地从两个通道获取到的数据，当两个通道都没有可用的数据时，会输出"no message received"。

```go
实例
package main
import "fmt"

func main() {
    //定义两个通道
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    //启动两个goroutine，分别从两个通道中获取数据
    go func() {
        for {	//相当于while(true)；用for true{}也可以
            ch1 <- "from 1"
        }
    }()
    go func() {
        for {
            ch1 <- "from 2"
        }
    }()
    
    //使用select语句非阻塞地从两个通道中获取数据
    for {
        select {
            case msg1 := <- ch1:
            fmt.Println(msg1)
            case msg2 := <- ch2:
            fmt.Println(msg2)
            default:
            //如果两个通道都没有可用的数据，则执行这里的语句
            fmt.Println("no massage recived")
        }
    }
}
```

以上代码执行结果为：

```
from 1
from 2
from 1
from 2
...		//	死循环执行线程
```

#### 八、Go语言函数

函数是基本的代码块，用于执行一个任务。
Go语言最少有个main()函数。
你可以通过函数来划分不同的功能，逻辑上每个函数的执行的是指定的任务。
函数声明告诉了编辑器函数的名称，返回类型、和参数。
Go语言标准库提供了多种可动用的内置的函数。例如，len()函数可以接受不同类型的长度。如果我们传入的是字符串则返回字符串的长度，如果传入的是数据，则返回数组中包含的元素个数。

**函数定义**

Go语言函数定义格式如下：

```go
func function_name([parameter list](参数列表) [return_types(返回的类型)]) {
	函数体
}
```

函数定义解析：

- func：函数有func开始声明
- function_name：函数名称，参数列表和返回值类型构成了函数签名。
- parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数(实参：被传过来的参数)。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
- return_type：返回类型，函数返回一列值。return_types是该值的数据类型。有些功能不需要返回值，这种情况下return_types不是必须的。
- 函数体：函数定义的代码集合。

```go
实例
package main
import "fmt"

func main() {
    //定义局部变量
    var a int = 10
    var b int = 20
    var ret int
    
    //调用函数并返回最大值
    ret = max(a, b)
    fmt.Printf("最大值是：%d\n",ret)
}

//函数返回两个数的最大值
func max(num1, num2 int) int {
    //定义局部变量
    var result int
    
    if num1 > num2 {
        result = num1
    }else {
        result = num2
    }
    return result
}
```

以上实例在main()函数中调用max()函数，执行结果为：

```
最大值是：20
```

**函数返回多个值**

Go函数可以返回多个值，例如：

```go
实例
package main
import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Google", "Runoob")
	fmt.Println(a, b)
}
```

以上实例执行结果为：

```
Runoob Google
```

**函数参数**

函数如果使用参数，该变量可称为函数的形参。
形参就像定义在函数体内的局部变量。
调用函数，可以通过两种方式传递参数：

**注意：Go中的函数参数是传值传递，即在传递参数时会拷贝实参的值**

| 传递类型 | 描述                                                         |
| -------- | ------------------------------------------------------------ |
| 值传递   | 值传递是指在调用函数时将实际参数赋值一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响实际参数。 |
| 引用传递 | 引用传递时指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数进行的修改，将影响到实际参数。 |

值传递实例(调用函数后，实际的参数值不变)

```go
package main
import "fmt"

func main() {
	//定义局部变量
    var a int = 100
    var b int = 200
    
    fmt.Printf("交换前 a 的值：%d\n",a)
    fmt.Printf("交换前 b 的值：%d\n",b)
    
    //通过调用函数来交换值
    swap(a, b)	//有个细节，不定义变量去接收，就可以避免使用不了
    
    fmt.Printf("交换后 a 的值：%d\n",a)
    fmt.Printf("交换后 b 的值：%d\n",b)
}

//定义相互交换值的函数
func swap(x, y int) int {
    var temp int
    temp = x	//保存x的值
    x = y		//保存y的值
    y = temp	//将temp值给y
    
    return temp
}
```

以上实例执行结果为：

```
交换前 a 的值： 100	//交换前
交换前 b 的值： 200
交换后 a 的值： 100	//交换后，在函数改变，但实际的地址值没变，所有还是100
交换后 b 的值： 200
```

**引用传递值**

引用传递是指在调用函数时将实际参数的地址传到函数中，那么在函数中对参数进行的修改，将影响实际参数，引用传递指针参数传递到函数内，以下是交换函数swap()使用了引用传递：

```go
//定义交互值函数
func swap(x *int, y *int){
	var temp int
	temp = *x	//保持x地址上的值
	*x = *y		//将y值赋给x
	*y = temp	//将temp值赋给y
}
```

以下是使用引用传递来调用swap()函数：

```go
package main
import "fmt"

func main() {
    //定义局部变量
    var a int = 100
    var b int = 200
    
    fmt.Printf("交换前，a的值：%d\n",a)
    fmt.Printf("交换前，b的值：%d\n",b)
    
    /*调用swap()函数
    *&a指向a指针，a变量的地址
    *&b指向b指针，a变量的地址
    */
    swap(&a, &b)
    fmt.Printf("交换前，a的值：%d\n",a)
    fmt.Printf("交换前，b的值：%d\n",b)
}

func swap(x *int, y *int) {
    var temp int
    temp = *x	//保存x地址上的值
    *x = *y		//将y值赋给x
    *y = temp	//将temp值赋给y
}
```

以上代码执行结果为：

```
交换前 a 的值为 : 100		//引用传递前的值
交换前 b 的值为 : 200
交换后 a 的值 : 200		//引用传递后的值，原地址值被修改，所以实际参数被修改为200
交换后 b 的值 : 100
```

#### 九、Go语言变量作用域

作用域为已声明标识符所表示的常量、类型、变量、函数或包在源代码中的作用范围。

Go 语言中变量可以在三个地方声明：

- 函数内定义的变量称为局部变量
- 函数外定义的变量称为全局变量
- 函数定义中的变量称为形式参数

**局部变量**

在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量。
以下实例中 main() 函数使用了局部变量 a, b, c：

```go
package main

import "fmt"

func main() {
   /* 声明局部变量 */
   var a, b, c int

   /* 初始化参数 */
   a = 10
   b = 20
   c = a + b

   fmt.Printf ("结果： a = %d, b = %d and c = %d\n", a, b, c)
}
```

以上实例执行输出结果为：

```
结果： a = 10, b = 20 and c = 30
```

**全局变量**

在函数体外声明的变量称之为全局变量，全局变量可以在整个包甚至外部包（被导出后）使用。
全局变量可以在任何函数中使用，以下实例演示了如何使用全局变量：

```go
实例
package main

import "fmt"

/* 声明全局变量 */
var g int

func main() {

   /* 声明局部变量 */
   var a, b int

   /* 初始化参数 */
   a = 10
   b = 20
   g = a + b

   fmt.Printf("结果： a = %d, b = %d and g = %d\n", a, b, g)
}
```

以上实例执行输出结果为：

```
结果： a = 10, b = 20 and g = 30
```

**Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。实例：**

```go
package main

import "fmt"

/* 声明全局变量 */
var a int = 20;

func main() {
   /* main 函数中声明局部变量 */
   var a int = 10
   var b int = 20
   var c int = 0

   fmt.Printf("main()函数中 a = %d\n",  a);
   c = sum( a, b);
   fmt.Printf("main()函数中 c = %d\n",  c);
}

/* 函数定义-两数相加 */
func sum(a, b int) int {
   fmt.Printf("sum() 函数中 a = %d\n",  a);
   fmt.Printf("sum() 函数中 b = %d\n",  b);

   return a + b;
}
```

以上实例执行输出结果为：

```
main()函数中 a = 10
sum() 函数中 a = 10
sum() 函数中 b = 20
main()函数中 c = 30
```

**初始化局部和全局变量**

不同类型的局部和全局变量的默认值为：

| 数据类型 | 初始化默认值 |
| :------- | :----------- |
| int      | 0            |
| float32  | 0            |
| pointer  | nil          |

#### 十、Go语言数组

Go语言提供了数组类型的数据结构

数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型，例如整型、字符串或者自定义类型。
相对于去声明number0，number1，...，number99的变量，使用数组形式number[0]，number[1]...，number[99]更加方便且易于拓展。
数组元素可以通过索引(位置)来读取(或者修改)，索引从0开始，第一个元素的索引为0，第二个索引为1，以此类推。

![image-20230525112644344](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20230525112644344.png)

**声明数组**

Go语言数组声明需要指定元素类型及元素个数，语法格式如下：

```go
var variable_name [SIZE] variable_type
```

以上为一维数组的定义方式，例如定义了数组balance长度为10，类型为float32：

```go
var balance [10] float
```

**初始化数组**

```go
var balance = [5]float32{100.1,12.5,1.3,20.3}
```

我们也可以通过字面量在声明数组的同时快速初始化数组：

```
balance := [5]g=float{100.0,45.5,3.4}
```

如果数组长度不确定，可以使用`...`代替数组的长度，编译器会根据元素个数自行推断数组的长度：

```go
var balance = [...]float{100.0,45.5,3.4}
或者
balance := [...]float{100.0,45.5,3.4}
```

如果设置了数组的长度，我们还可以通过指定下标来初始化元素：

```go
//将索引为1和3的元素初始化
balance :=[5]float{1:2.0,3:7.0}
```

初始化数组中`{}`的元素个数不能大于`[]`中的数字。

如果忽略`[]`中的数字不设置数组大小，Go语言会根据元素的个数来设置数组的大小：

```go
balance[4] = 50.0
```

![image-20230525114140640](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20230525114140640.png)

**访问数组元素**

数组元素可以通过索引(位置)来读取。格式为数组名后加中括号(`[])`，中括号中为索引的值，如：

```go
var salary float32 = balance[6]
```

以上实例读取了数组balance第7个元素的值。

以下演示了数组完整操作（声明、赋值、访问）的实例：

```go
实例
package main
import "fmt"

func main() {
    var n [10]int	//n是一个长度为10的数组
    var i, j int
    
    //为数组n初始化元素
    for i = 0;i < 10; i++ {
        n[i] = i + 100		//设置元素为i + 100
    }
    
    //输出每个数组元素的值
    for j = 0; j < len(n); j++ {
        fmt.Printf("Element[%d] = %d\n", j, n[j] )
    }
}
```

以上实例执行结果如下：

```
Element[0] = 100
Element[1] = 101
Element[2] = 102
Element[3] = 103
Element[4] = 104
Element[5] = 105
Element[6] = 106
Element[7] = 107
Element[8] = 108
Element[9] = 109
```

```go
实例2
package main
import "fmt"

func main() {
    var i, j, k int
    //声明数组的同时快速初始化数组
    balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
    
    //输出数组元素
    for i = 0; i < len(balance); i++ {
        fmt.Printf("balance[%d] = %f\n",i,balance[i])
    }
    
    balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
    for j = 0; j < len(balance2); j++ {
        fmt.Printf("balance2[%d] = %f\n",j,balance2[j])
    }
    
    //将索引为1和3的元素初始化
    balance3 := [5]float32{1:2.0,3:7.0}
    for k = 0; k < 5; k++ {
        fmt.Printf("balance3[%d] = %f\n",k,balance3[k])
    }
}
```

以上实例执行结果如下：

```
balance[0] = 1000.000000
balance[1] = 2.000000
balance[2] = 3.400000
balance[3] = 7.000000
balance[4] = 50.000000
balance2[0] = 1000.000000
balance2[1] = 2.000000
balance2[2] = 3.400000
balance2[3] = 7.000000
balance2[4] = 50.000000
balance3[0] = 0.000000
balance3[1] = 2.000000
balance3[2] = 0.000000
balance3[3] = 7.000000
balance3[4] = 0.000000
```

**Go语言多维数组**

Go语言支持多维数组，以下为常用多维数组声明方式：

```go
var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type
```

以下实例声明 了三维数组：

```go
var threedim [5][10][4] int
```

**二维数组**

二维数组是最简单的多维数组，二维数组本质上是由一维数组组成的。二维数组定义方式如下：

```go
var arrayNamr [x][y] variable_type
```

variable_type为Go语言的数据类型，arrayName为数组名，二维数组可认为是一个表格，x为行，y为列，下图演示了一个二维数组a为三行四列：

![image-20230525134826694](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20230525134826694.png)

二维数组中的元素可以通过`a[i][j]`来访问。

```go
实例
package main 
import "fmt"

func main() {
    //Step 1 : 创建数组
    values := [][]int{}
    
    //Step 2 : 使用append()函数向空的二维数组加两行一维数组
    row1 := []int{1,2,3}
    row2 := []int{4,5,6}
    values = append(values,row1)
    values = append(values,row2)
    
    // Step 3 : 显示两行数据
    fmt.Println("Row 1")
    fmt.Println(values[0])
    fmt.Println("Row 2")
    fmt.Println(values[1])
    
    //Step 4 : 访问第一个元素
    fmt.Printf("第一个元素为： %d",values[1][0])
}
```

以上实例运行输出结果为：

```
Row 1
[1 2 3]
Row 2
[4 5 6]
第一个元素为： 4
```

**初始化二维数组**

多维数组可通过大括号`{}`来初始值。以下实例为一个3行4列的二维数组：

```go
a := [3][4]int {
	{0,1,2,3},		//第一行索引为0
	{4,5,6,7},		//第二行索引为1
    {8,9,10,11}	,	//注意：最后一行这样写必须要有 , 号
}
```

**注意：**以上代码中倒数第二行的`}`必须有逗号，因为最后一行的 `}`不能单独一行，也可以写成这样：

```go
a := [3][4]int {
	{0,1,2,3},		//第一行索引为0
	{4,5,6,7},		//第二行索引为1
    {8,9,10,11}}	//第二行索引为2
}
```

以下实例化一个2行2列的二维数组：

```go
实例
package main 
import "fmt"

func main() {
	//创建二维数组
	sites := [2][2]string{}
	
	//向二维数组添加元素
	sites[0][0] = "Google"
	sites[0][1] = "Runoob"
	sites[1][0] = "Taobao"
	sites[1][1] = "Weibo"
	
	//显示结果
	fmt.Println(sites)
}
```

以上实例运行输出结果为：

```
[[Google Runoob] [Taobao Weibo]]
```

**访问二维数组**

二维数组通过指定坐标来访问。如数组中的行索引与列索引，例如：

```go
val := a[2][3]
或
var val int = a[2][3]
```

以上实例访问了二维数组val第三行的第四个元素。

二维数组可以使用循环嵌套来输出元素：

```
实例
package main
import "fmt"

func main() {
	//初始化5行2列数组
	a := [5][2]int{{0,0},{1,2},{2,4},{3,6},{4,8}}
	var i, j int
	
	//输出数组元素
	for i = 0;i < 5; i++ {
		for j = 0;j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n",i,j,a[i][j])
		}
	}
}
```

以上实例运行输出结果为：

```
a[0][0] = 0
a[0][1] = 0
a[1][0] = 1
a[1][1] = 2
a[2][0] = 2
a[2][1] = 4
a[3][0] = 3
a[3][1] = 6
a[4][0] = 4
a[4][1] = 8
```

创建各个维度元素个数不一致的多维数组：

```go
实例
package main
import "fmt"

func main() {
    //创建空的二位数组
    animals := [][]string{}
    
    //创建三个一维数组，各数组长度不同
    row1 := []string{"fish","shark","eel"}
    row2 := []string{"bird"}
    row3 := []string{"lizard","salamander"}
    
    //使用append()函数将一维数组添加到二维数组中
    animals = append(animals,row1)
    animals = append(animals,row2)
    animals = append(animals,row3)
    
    //循环输出
    for i := range animals {
        fmt.Printf("Row: %v\n",i)	//增强for遍历，得到每一个一维数组
        fmt.Println(animals[i])
    }
}
```

以上实例运行输出结果为：

```
Row: 0
[fish shark eel]
Row: 1
[bird]
Row: 2
[lizard salamander]
```

#### 十一、Go语言指针

Go语言中指针是很容易学习的，Go语言中使用指针可以更简单的执行一些任务。

变量是一种使用方便的占位符，用于引用计算机的内存地址。

Go语言的取地址符是`&`，放到一个变量前使用就会返回相应的内存地址。

以下实例演示了变量在内存中地址：

```go
实例
package main
import "fmt"

func main() {
    var a int = 10
    fmt.Printf("变量的计算机内存地址是： %x\n",&a )
}
```

执行以上代码输出结果为：

```
变量的计算机内存地址是： c00000e0a8
```

现在我们已经了解了什么是内存地址和如何去访问它。下面将具体介绍指针。

**什么是指针**

一个指针变量指向了一个值的内存地址

类似于变量和常量，在使用指针前你需要声明指针，指针声明格式如下：

```go
var var_name *var_type
```

var_type为指针类型，var_name为指针变量名，*号用于指定变量是作为一个指针。指针声明：

```go
var ip *int		//指向整型
var fp *float32 //指向浮点型
```

上例是一个指向int和float32的指针

**如何使用指针**

指针使用流程：

- 定义指针变量
- 为指针变量赋值
- 访问指针变量中指向地址的值

在指针类型前面加上`*`号(前缀)来获取指针所指向的内容。

```go
package main
import "fmt"

func main() {
	var a int = 20		//声明一个实际变量
	var ip *int			//声明指针变量
	
	ip = &a				//指针变量的内存地址
	
	fmt.Printf("a变量的内存地址是： %x\n",&a )
	
	//指针变量的存储地址
	fmt.Printf("ip变量存储的指针地址是： %x\n",ip )
	
	//使用指针访问值
	fmt.Printf("ip变量的值是： %d\n",*ip )
}
```

以上实例执行输出结果为：

```
a变量的内存地址是： c00000e0a8
ip变量存储的指针地址是： c00000e0a8
ip变量的值是： 20
```

**Go空指针**

当一个指针被定义后没有分配任何变量时，它的值为`nil`。
`nil`指针也称为空指针
nil在概念上和其他语言的null、None、nil、NULL一样，都指代零值或空值。
一个指针变量通常缩写为ptr(pointer)。

```go
实例
package main
import "fmt"

func main() {
    var ptr *int
    
    fmt.Printf("ptr的值为： %x\n",ptr)
}
```

以上实例输出结果为：

```
ptr的值为： nil
```

空指针判断：

```go
if(ptr != nil)		//	prt不是空指针
if(ptr == nil)		//ptr时空指针
```

**Go语言指针数组**

了解指针数组前，看个实例，定义了长度为3的整型数组：

```go
package main
import "fmt"

const MAX int = 3	//声明一个全局变量，变量名用大写

func main() {
    
    a := []int{10,100,200}	//	不指定长度，初始化值，编辑器会根据元素个数推断长度
    for i := 0;i < MAX; i++ {
        fmt.Printf("a[%d] = %d\n",i, a[i] )
    }
}
```

以上代码执行输出结果为：

```
a[0] = 10
a[1] = 100
a[2] = 200
```

有一种情况，我们可能需要保存数组，这样我们就需要使用指针。
以下声明了整型指针数组：

```go
var ptr [MAX]*int;
```

ptr为整型指针数组。因此每个元素都指向了一个值。以下实例的三个整数将存储在指针数组中：

```go
实例
package main
import "fmt"

const MAX int = 3

func main() {
    a := []int{10,100,200}
    var ptr [MAX]*int;
    
    for i := 0; i < MAX;i++ {
        ptr[i] = &a[i]		//整型地址赋值给指针数组
    }
    
    for i := 0;i < MAX;i++ {
        fmt.Printf("a[%d] = %d\n",i,*ptr[i] )
    }
}
```

以上代码指向输出结果为：

```
a[0] = 10
a[1] = 100
a[2] = 200
```

**Go语言指向指针的指针**

如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。
当定义一个指向指针的变量时，第一个指针存放第二个指针的地址，第二给指针存放变量的地址：

![image-20230525170903727](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20230525170903727.png)

指向指针的指针变量声明格式如下：

```go
var ptr **int;
```

以上指向指针的指针变量为整型。

访问指向指针的指针变量需要使用两个`**`号，如下所示：

```go
package main
import "fmt"

func main() {
    
    var a int
    var ptr *int
    var pptr **int
    
    a = 3000
    
    //指针ptr地址
    ptr = &a
    
    //指向ptr地址
    pptr = &ptr
    
    //获取pptr的值
    fmt.Printf("变量 a = %d\n",a)
    fmt.Printf("指针变量 ptr = %d\n",*ptr)
    fmt.Printf("指向指针的指针变量 pptr = %d\n",**pptr)
}
```

以上实例执行输出结果为：

```
变量 a = 3000
指针变量 *ptr = 3000
指向指针的指针变量 **pptr = 3000
```

**Go语言指针作为函数参数**

Go语言允许向函数传递指针，只需要在函数定义的参数上设置为指针类型即可。

以下实例演示了如何向函数传递指针，并在函数调用后修改函数内的值：

```go
package main
import "fmt"

func main() {
    //定义局部变量
    a := 100
    b := 200
    
    fmt.Printf("交换前 a 的值 ： %d\n",a)
    fmt.Printf("交换前 b 的值 ： %d\n",b)
    
    //调用函数用于交换值
    // &a指向 a 变量的地址
    // &b指向 b 变量的地址
    swap(&a,&b);
    
    fmt.Printf("交换后 a 的值 ： %d\n",a)
    fmt.Printf("交换后 b 的值 ： %d\n",b)
}

func swap(x *int, y *int) {
    var temp int
    temp = *x		//保存x地址的值
    *x = *y			//将y地址的值赋给x
    *y = temp		//将temp赋值给y
}

//交换函数这样写更加简洁，也是go语言的特性，c++和c#是不能这么干的
func swap(x *int,y *int) {
    *x,*y = *y,*x
}
```

#### 十二、Go语言结构体

Go语言中数组可以存储同一类型的数据，但在结构体中我们为不同项定义不同的数据类型。
结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
结构体表示一项纪录，比如保存图书馆的书籍记录，每本书有以下属性：

- Title : 标题
- Author : 作者
- Subject : 学科
- ID : 书籍ID

**定义结构体**

结构体定义需要使用type和struct语句。struct语句定义一个新的数据类型，结构体中有一个或多个成员(属性)。type语句设定了结构体的名称，结构体的格式如下：

```go
type struct_variable_type struct {
    member definition
    member definition
    ...
    member definition
}
```

一旦定义了结构体类型，它就能用于变量的声明，语法格式下：

```go
variable_name := struct_variable_type{value1,value2...valuen}
或
variable_name := struct_variable_type{key1: value1,key2: value2...,keyn:valuen}
```

实例如下：

```go
package main
import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	
    //创建一个新的结构体
    fmt.Println(Books{"Go语言","www.runoob.com","Go语言教程",6396407})
    
    //也可以使用key => value 格式
    fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})
    
    //忽略的字段为 0 或 空(nil)
    fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}
```

输出结果为：

```
{Go语言 www.runoob.com Go语言教程 6396407}
{Go 语言 www.runoob.com Go 语言教程 6495407}
{Go 语言 www.runoob.com  0}
```

**访问结构体成员**

如果要访问结构体成员，需要使用点号 `.`操作符，实例如下：

```
结构体.成员名
```

结构体类型变量使用struct关键字定义，实例如下：

```go
实例
package main 
import "fmt"

type Books struct {
    title string
    author string 
    subject string
    book_id int
}

func main() {
	var Book1 Books		//声明Book1 为 Books类型
    var Book2 Books		//声明Book12 为 Books类型
    
    //Book1描述
    Book1.title = "Go 语言"
    Book1.author = "www.runoob.com"
    Book1.subject = "Go 语言教程"
    Book1.book_id = 6343452
    
    //Book2描述
    Book2.title = "Python 语言"
    Book2.author = "www.runoob.com"
    Book2.subject = "Python 语言教程"
    Book2.book_id = 6343452
    
    //打印Book1信息
    fmt.Printf("Book 1 title : %s\n",Book1.title)
    fmt.Printf("Book 1 author : %s\n",Book1.author)
    fmt.Printf("Book 1 subject : %s\n",Book1.subject)
    fmt.Printf("Book 1 book_id : %d\n",Book1.book_id)
    
    //打印Book2信息
    fmt.Printf("Book 2 title : %s\n",Book2.title)
    fmt.Printf("Book 2 author : %s\n",Book2.author)
    fmt.Printf("Book 2 subject : %s\n",Book2.subject)
    fmt.Printf("Book 2 book_id : %d\n",Book2.book_id)
}
```

以上实例执行运行结果为：

```
Book 1 title : Go 语言
Book 1 author : www.runoob.com
Book 1 subject : Go 语言教程
Book 1 book_id : 6343452
Book 2 title : Python 语言
Book 2 author : www.runoob.com 
Book 2 subject : Python 语言教程
Book 2 book_id : 6343452
```

**结构体作为函数参数**

你可以像其它数据类型一样将结构体作为参数传递给函数，并加以实例方访问结构体变量：

```go
实例
package main 
import "fmt"

type Books struct {
    title string
    author string 
    subject string
    book_id int
}

func main() {
	var Book1 Books		//声明Book1 为 Books类型
    var Book2 Books		//声明Book12 为 Books类型
    
    //Book1描述
    Book1.title = "Go 语言"
    Book1.author = "www.runoob.com"
    Book1.subject = "Go 语言教程"
    Book1.book_id = 6343452
    
    //Book2描述
    Book2.title = "Python 语言"
    Book2.author = "www.runoob.com"
    Book2.subject = "Python 语言教程"
    Book2.book_id = 6343452
    
    //打印Book1信息
    printBook(Book1)
    
    //打印Book2信息
    printBook(Book2)
}

func printBook(book Books) {
    fmt.Printf("Book  title : %s\n",book.title)
    fmt.Printf("Book  author : %s\n",book.author)
    fmt.Printf("Book  subject : %s\n",book.subject)
    fmt.Printf("Book  book_id : %d\n",book.book_id)

}
```

以上执行运行结果为：

```
Book title : Go 语言
Book author : www.runoob.com
Book subject : Go 语言教程
Book book_id : 6495407
Book title : Python 教程
Book author : www.runoob.com
Book subject : Python 语言教程
Book book_id : 6495700         
```

**结构体指针**

你可以定义结构体的指针类似于其它指针变量，格式如下 ：

```go
var struct_pointer *Books
```

以上定义的指针变量可以存储结构体变量的地址。查看结构体变量的地址，可以将`&`符号至于结构体变量前：

```
struct_pointer = &Book1
```

使用结构体指针访问结构体成员，使用 `.`操作符：

```
struct_pointer.title
```

 接下来让我们使用结构体指针重写以上实例，代码如下：

```go
package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}

func main() {
    
    var Book1 Books        /* 声明 Book1 为 Books 类型 */
    var Book2 Books        /* 声明 Book2 为 Books 类型 */

    /* book 1 描述 */
    Book1.title = "Go 语言"
    Book1.author = "www.runoob.com"
    Book1.subject = "Go 语言教程"
    Book1.book_id = 6495407

    /* book 2 描述 */
    Book2.title = "Python 教程"
    Book2.author = "www.runoob.com"
    Book2.subject = "Python 语言教程"
    Book2.book_id = 6495700
    
   //打印Book1信息
    printBook(&Book1)
    
    //打印Book2信息
    printBook(&Book2)
}

func printBook(book *Books) {
    fmt.Printf("Book title : %s\n",book.title)
    fmt.Printf("Book author : %s\n",book.author )
    fmt.Printf("Book subject : %s\n",book.subject )
    fmt.Printf("Book book_id  : %d\n",book.book_id )
}
```

以上实例执行运行结果为：

```
Book title : Go 语言
Book author : www.runoob.com
Book subject : Go 语言教程
Book book_id  : 6495407
Book title : Python 教程
Book author : www.runoob.com
Book subject : Python 语言教程
Book book_id  : 6495700
```

**结构体中属性的首字母大小写问题**

- 首字母大写相当于public
- 首字母小写相当于private

**注意：**这个public和private是相对于包（go文件首行的package后面跟的包名）来说的。

**敲黑板，划重点**

当要将结构体对象转换为JSON时，对象中的属性首字母必须是大写，才能正常转为JSON。

实例一：

```go
package main
import (
	"fmt"
	"encoding/json"
)

type Person struct {
    Name string		//Name字段首字母大写
    age int 			//age字段首字母小写
}

func main() {
    person := Person{"小明",18}
    if result,err := json.Marshal(&person);err == nil {	//json.Marshal将对象转换为json字符串
        fmt.Println(string(result))
    }
}
```

控制台输出：

```
{"Name":"小明"}	//只有Name，没有age
```

实例二：

```
type Person struct {
	Name string		//都是大写
	Age int
}
```

控制台输出：

```
{"Name":"小明","Age":18}   //两个字段都有
```

那这样JSON字符串以后就只能都是大写了么？当然不是，可以使用tag标记要返回的字段名。
实例三：

```go
package main
import (
	"fmt"
	"encoding/json"
    "time"
)

type Person struct {
    Name string	 `josn:"name"`//标记json名字name
    Age int      `josn:"age"`
    Time int64   `json:"-"`	  //标记忽略该字段
}

func main() {
    person := Person{"小明",18,time.Now().Unix()}
    if result,err := json.Marshal(&person);err == nil {	//json.Marshal将对象转换为json字符串
        fmt.Println(string(result))
    }
}
```

控制台输出：

```
{"Name":"小明","Age":18}
```

#### 十三、Go语言切片（Slice）

Go语言切片是对数组的抽象。
Go数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，概念强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

**定义切片**

你可以声明一个为指定大小的数组来定义切片：

```
var identifier []type
```

切片不需要说明长度。
或使用make()函数来创建切片：

```go
var slice1 []type = make([type],len)
也可简写为
slice1 := make([]type,len)
```

也可以指定容量，其中**capacity**为可选参数。

```
make([]T,length,capacity)
```

这里len是数组的长度并且也是切片的初始长度。

**切片初始化**

直接初始化切片，`[]`表示是切片类型，{1,2,3}初始化值依次是`1,2,3`,其中cap=len=3.

```
s := []int{1,2,3}
```

初始化切片s，是数组arr的引用

```
s := arr[:]
```

将arr中下标**startIndex**到 **endIndex - 1**下的元素创建一个新的切片。

```
s := arr[startIndex:endIndex]
```

默认endIndex时将表示一直到arr的最后一个元素。

```
s := arr[startIndex:]
```

默认startIndex时，将表示从arr的第一个元素开始。

```
s := arr[:endIndex]
```

通过切片s初始化切片s1

```
s1 := s[startIndex:endIndex]
```

通过内置函数make()初始化切片s，[]int标识为其元素类型为int的切片。

```
s := make([]int,len,cap)
```

**len()和cap()函数**

切片时可索引的，并且可以有len()方法获取长度。
切牌你提供计算容量的方法cap(),可以测量切片最长可以达到多少。
以下为具体实例：

```go
package main
import "fmt"

func main() {
    var numbers = make([]int,3,5)
    
    printSilce(numbers)
}

func printSilce(x []int){
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```

控制台输出结果为：

```
len=3 cap=5 slice=[0 0 0]
```

**空(nil)切片**

一个切片在未初始化之前默认为nil，长度为0，实例如下：

```go
实例
package main 
import "fmt"

func main() {
	var numbers [] int
    
    printSlice(numbers)
    
    if numbers == nil {
        fmt.Printf("切片是空的")
    }
}

func printSlice(x []int) {
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```

以上实例运行输出结果为：

```
len=0 cap=0 slice=[]
切片是空的
```

**切片截取**

可以通过设置下限及上限来设置截取切片[lower-bound:upper-bound]，实例如下：

```go
实例
package main
import "fmt"

func main() {
	//创建切片
    numbers := []int{0,1,2,3,4,5,6,7,8}
    printSlice(numbers)
    
    //打印原始切片
    fmt.Println("numbers ==",numbers)
    
    //打印子切片从索引1(包含)到索引4(不包含)
    fmt.Println("numbers[1:4] ==",numbers[1:4])
    
    //默认下限为0
    fmt.Println("numbers[:3] == ",numbers[:3])
    
    //默认上限为len(s)
    fmt.Println("numbers[4:] == ",numbers[4:])
    
    numbers1 := make([]int,0,5)
    printSlice(numbers1)	//cap =5
    
    //打印子切片从索引0(包含)到索引2(不包含)
    numbers2 := numbers[:2]
    printSlice(numbers2)	//cap=9,因为是从0索引开始截取，numbers的cap为9
    
    //打印子切片从索引2(包含)到索引5(不包含)
    numbers3 := numbers[2:5]//cap=7,因为2(包含2)开始截取，少了0和1索引的值，cap=9-2=7
    printSlice(numbers3)
}

func printSlice(x []int) {
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```

执行以上代码输出结果为：

```
len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8]
numbers == [0 1 2 3 4 5 6 7 8]
numbers[1:4] == [1 2 3]
numbers[:3] ==  [0 1 2]
numbers[4:] ==  [4 5 6 7 8]
len=0 cap=5 slice=[]
len=2 cap=9 slice=[0 1]
len=3 cap=7 slice=[2 3 4]
```

**append()和copy()函数**

如果向增加切片的容量，我们必须创建一个新的更大的切片并把原分切片的内容都拷贝过来。

下面的代码描述了从拷贝切片到copy方法和向切片追加新元素的方法。

```go
实例
package main
import "fmt"

func main() {
	var numbers []int
    printSlice(numbers)
    
    //允许追加空切片
    numbers = append(numbers,0)
    printSlice(numbers)
    
    //向切片添加一个元素
    numbers = append(numbers,1)
    printSlice(numbers)
    
    //同时添加多个元素
    numbers = append(numbers,2,3,4)
    printSlice(numbers)
    
    //创建切片numbers1是之前切片容量的两倍
    numbers1 := make([]int,len(numbers),(cap(numbers))*2)
    
    //拷贝Numbers的内容到numbers1
    copy(numbers1,numbers)
    printSlice(numbers1)
}

func printSlice(x []int) {
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```

以上代码执行输出结果为：

```
len=0 cap=0 slice=[]
len=1 cap=1 slice=[0]
len=2 cap=2 slice=[0 1]
len=5 cap=6 slice=[0 1 2 3 4]
len=5 cap=12 slice=[0 1 2 3 4]
```

#### 十四、Go语言范围(Range)

Go语言中range关键字用于for循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中返回元素的索引和元素对应的值，在集合中返回key-value对。
for循环的range格式可以对slice、map、数组、字符串等进行迭代循环。格式如下：

```go
for key,value := range oldMap {
    newMap[key] = value
}
```

以上代码中的key和value是可以省略。

```go
//如果只想读取key，格式如下：
for key := range oldMap

//或者这样：
for key:_ :=range oldMap

//如果只想读取value，格式如下：
for _:value := range oldMap
```

实例：遍历简单的数组，`2**%d`的结果为2对应的次方数：

```go
package main
import "fmt"

var pow = []int{1,2,4,8,16,32,64,128}

func main() {
    for i,v := range pow {
        fmt.Printf("2**%d = %d\n",i,v)
    }
}	
```

控制台输出结果为：

```
2**0 = 1
2**1 = 2
2**2 = 4
2**3 = 8
2**4 = 16
2**5 = 32
2**6 = 64
2**7 = 128
```

实例：for循环的range格式可以省略key和value，如下：

```go
package main
import "fmt"

func main() {
    map1 := make(map[int]float32)
    map1[1] = 1.2
    map1[2] = 2.2
    map1[3] = 3.2
    map1[4] = 4.2
    
    //读取key和value
    for key,value :=range map1 {
        fmt.Printf("key is : %d -- value is : %f\n",key,value )
    }
    
    //读取key
    for key := range map1 {
        fmt.Printf("key is : %d\n",key)
    }
    
    //读取value
    for _,value := range map1 {
        fmt.Printf("value is : %f\n",value)
    }
}
```

控制台输出结果为：

```
key is : 1 -- value is : 1.200000
key is : 2 -- value is : 2.200000
key is : 3 -- value is : 3.200000
key is : 4 -- value is : 4.200000
key is : 1
key is : 2
key is : 3
key is : 4
value is : 1.200000
value is : 2.200000
value is : 3.200000
value is : 4.200000
```

实例：range遍历其他数据结构：

```go
package main
import "fmt"

func main() {
    //这是我们使用range去求一个slice的和。使用数组跟这个类似
    nums := []int{2,3,4}
    sum := 0
    for _,num := range nums {
        sum += num
    }
    fmt.Println("sum : ",sum)
    //在数组上使用range将传入索引和值两个变量。上面那个例子我们不需要使用该元素的序号，索引我们使用空白符 "_" 省略了，有时候我们确实需要知道它的索引
    for i,num := range nums {
        if num == 3 {
            fmt.Println("index : ",i)
        }
    }
    
    //range也可以用在map的键值对上
    kvs := map[string]string{"a" : "apple","b" : "banana"}
    for k,v := range kvs {
        fmt.Printf("%s -> %s\n",k,v)
    }
 //range也可以用枚举Unicode字符串，第一个参数是字符串的索引，第二个是字符串(Unicode的值)本身
    for i,c := range "go" {
        fmt.Println(i,c)
    }
}
```

控制台输出结果为：

```
sum :  9
index :  1
a -> apple
b -> banana
0 103
1 111
```

**注意：涉及到指针是需要注意，v是个单独的地址。**

```go
func main() {
    nums :=[3]int{5,6,7}
    for k,v :=range nums {
        fmt.Println("源值地址： ",&nums[k],"\t value的地址：",&v)
    }
}
```

输出结果为：

```
源值地址： 0xc000012138          value的地址： 0xc00000e0a8
源值地址： 0xc000012140          value的地址： 0xc00000e0a8
源值地址： 0xc000012148          value的地址： 0xc00000e0a8
```

#### 十五、Go语言Map（集合）

Map是一种无序的键值对的集合。
Map最重要的一点是通过key来快速检索数据，key类似于索引，指向数据的值。
Map是一种集合，所以我们可以先迭代数据和切片那样迭代它。不过，Map是无序的，遍历Map时返回的顺序是不确定的。
在获取Map的值时，如果键值对不存在，返回该类型的零值，例如int类型的零值是0，string类型的零值是`" "`。
Map是引用类型，如果将一个Map传递给一个函数或赋值给另一个变量，他们都指向一个底层数据结构，因此对Map的修改会影响到所有引用它的变量。

**定义Map**

可以使用内建函数make或是使用map关键字来定义Map：（发现一个细节，使用make函数创建可以不用初始值）

```
//使用make()来定义
map_variable := make(map[KeyType]ValueType,initialCapacity)
```

其中KeyType是键的类型，ValueType是值的类型，initialCapacity是可选参数，用于指定Map的初始容。Map的容量是指Map中可以保存的键值对的数量，当Map中的键值对数量达到容量是，Map会自动扩容。如果不指定initialCapacity，Go语言会根据实际情况选择一个合适的值。

```go
实例
//创建一个空的Map
m := make(map[string]int)

//创建一个初始容量为10的Map
m := make(map[string]int,10)
```

也可以使用字面量创建Map：

```go
//使用字面量创建Map
m := map[string]int{
    "apple":1,
    "banana":2,
    "orange":3,
}
```

```go
//获取键值对
v1 := m["apple"]

//判断键是否存在
v2,ok := m["pear"]	//如果键不存在，ok的值为false，v2的值为该类型的零值

//修改元素
m["apple"] = 5

//获取Map的长度
len := len(m)

//遍历Map
for k,v :=range m {
    fmt.Printf("key=%s,value=%d\n",k,v)
}

//删除元素：
delete(m,"banana")
```

实例：下面演示创建和使用map

```go
package main
import "fmt"

func main() {
    siteMap := map[string]string {
        "Google":"谷歌",
        "Runoob":"菜鸟教程",
        "Baidu":"百度",
        "Wiki":"维基百科",
    }
    
    //使用键输出对应值
    for site := range siteMap {
        fmt.Println(site,"对应的中文值是->",siteMap[site])
    }
    
    //查看元素在集合中是否存在
    name,ok := siteMap["Facebook"]//如果确定存在，则ok为true，否则不存在
    if ok {
        fmt.Println("Facebook 的对应中文值是->",name)
    }else {
        fmt.Println("不存在中文值")
    }
}
```

控制台输出结果为：

```
Google 对应的中文值是-> 谷歌
Runoob 对应的中文值是-> 菜鸟教程
Baidu 对应的中文值是-> 百度
Wiki 对应的中文值是-> 维基百科
不存在中文值
```

**delete()函数**

delete()函数用于山删除集合的元素，参数为map和对应的key。实例如下：

```go
package main
import "fmt"

func main() {
    countryCapitalMap := map[string]string {
        "France":"Paris",
        "Italy":"Rome",
        "China":"Beijing",
        "India":"New delhi",
    }
    fmt.Println("国家对应的首都")
    
    //根据国家打印对应的首都
    for country := range countryCapitalMap {
        fmt.Println(country,"的首都是",countryCapitalMap[country])
    }
    
    //删除元素
    delete(countryCapitalMap,"India")
    fmt.Println("India被删除")
    
    //删除后的国家首都
    for country := range countryCapitalMap {
        fmt.Println(country,"的首都是",countryCapitalMap[country])
    }
}
```

控制台输出结果为：

```
国家对应的首都
France 的首都是 Paris
Italy 的首都是 Rome
China 的首都是 Beijing
India 的首都是 New delhi
India被删除
France 的首都是 Paris
Italy 的首都是 Rome
China 的首都是 Beijing
```

#### 十六、Go语言递归函数

递归，就是在运行的过程中调用自己。语法格式如下：

```go
func recursion() {
	recursion()	//函数调用自身
}

func main() {
    recursion()
}
```

go语言支持递归。但我们在使用递归时，开发者需要设置退出条件，否则递归将陷入无限循环中。

递归函数对于解决数学上的问题时非常有用的，就像计算阶乘，生成斐波那契数列等。

**阶乘**

以下实例通过Go语言的递归函数实现阶乘：

```go
package main
import "fmt"

func Factoral(n uint64)(result uint64){
    if (n > 0) {
        result = n * Factoral(n-1)
        return result
    }
    return 1
}

func main() {
    i := 15
    fmt.Printf("%d的阶乘是 %d\n",i,Factoral(uint64(i)))
}
```

以上实例输出结果为：

```
15 的阶乘是 1307674368000
```

**斐波那契数列**

以下实例通过Go语言的递归函数实现斐波那契数列：

```go
package main
import "fmt"

func fibonacci(n int) int {
    if n < 2 {
        return n
    }
    return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
    for i := 0;i <10; i++ {
        fmt.Printf("%d\t",fibonacci(i))
    }
}
```

控制台输出结果为：

```
0       1       1       2       3       5       8       13      21      34
```

#### 十七、Go语言类型转换

类型转换用于将一种类型的变量转换为另一种类型的变量。
Go语言类型转换基本格斯如下：

```
type_name(expression)
```

type_name为类型，expression为表达式。

**数值类型转换**

将整型转换为浮点型：

```go
a := 10
var b float64 = float64(a)
```

实例：将整型转化为浮点型，并计算结果，将结构域赋值给浮点型变量：

```go
package main
import "fmt"

func main() {
    sum := 17
    count := 5
    var mean float32
    
    mean = float32(sum)/float32(count)
    fmt.Printf("mean的值为： %f\n",mean)
}
```

控制台输出结果为：

```
mean的值为： 3.400000
```

**字符串类型转换**

将一个字符串转换成另一个类型，可以使用以下语法：

```go
str := "10"
var num int
num,_ = strconv.Atoi(str)
```

以上代码将字符串变量str转型为整型变量num。
**注意：**`strconv.Atoi`函数返回两个值第一个是转换后的整型值，第二个是可能发生的错误，我们可以使用空白标识符 `_` 来忽略这个错误。
以下实例将字符串转换为整数：

```go
package main 
import (
	"fmt"
    "strconv"
)

func main() {
    str := "123"
    num,err := strconv.Atoi(str)
    if err != nil {
        fmt.Println("转换错误:",err)
    } else {
        fmt.Printf("字符串 '%s' 转换为整型 ：%d\n",str,num)  
    }
}
```

控制台输出结果为：

```
字符串 '123' 转换为整型 ：123
```

以下实例将字符串转换为浮点数：

```go
package main
import (
	"fmt"
	"strconv"
)

func main() {
    str := "3.14"
    num,err := strconv.ParseFloat(str,64)
    if err != nil {
        fmt.Println("转换错误：",err)
    } else {
        fmt.Printf("字符串 '%s' 转为浮点型 ：%f\n",str,num)
    }
}
```

控制台输出结果为：

```
字符串 '3.14' 转为浮点型 ：3.140000
```

**接口类型转换**

接口类型转换有两种情况：**类型断言**和**类型转换
类型断言用于将接口类

型转换为指定类型，其语法为：

```go
value.(type)
或者
value.(T)
```

其中value是接口类型的变量，type或T是要转换成的类型。
如果类型断言成功，它将返回转换后的值和一个布尔值，表示转换是否成功。

```go
package main
import "fmt"

func main() {
    var i interface{} = "Hello,World"
    str,ok := i.(string)
    if ok {
        fmt.Printf("'%s' is a string\n",str)
    } else {
         fmt.Println("conversion failed")
    }   
}
```

以上实例中，我们定于了一个接口类型变量 `i`，并将它赋值为字符串 "Hello,world!"。然后。我们使用类型断言将 `i`转换为字符串类型，并将转换后的值赋值给变量str。最后我们使用ok变量检查类型转换是否成功，如果成功，我们打印转换后的字符串；否则，我们打印转换失败的消息。
类型转换用于将一个接口类型的值转换为另一个接口类型，其语法：

```
T(value)
```

T是目标接口类型，value要转换的值。

在类型转换中，我们必须保证要转换的值和目标接口类型之间是兼容的，否则编译器会报错。

```go
实例
package main
import "fmt"

type Writer interface {
    Write([]byte)(int,error)
}

type StringWriter struct {
    str string
}

func (sw *StringWriter) Write(data []byte) (int,error){
    sw.str += string(data)
    return len(data),nil
}

func main() {
    var w Writer = &StringWriter{} 	//将StringWriter类型的指针赋值给Writer接口类型的w
    sw := w.(*StringWriter)		//将w转换为StringWriter类型
    sw.str = "Hello,World"
    fmt.Println(sw.str)
}
```

控制台输出结果为：

```
Hello,World
```

以上实例中，我们定义了一个Writer接口和一个事项该接口的结构体StringWriter。然后，我们将StringWriter类型的指针赋值给Writer接口类的变量w，接着，我们使用类型转换将StringWriter类型，并将转换后的值赋值给变量sw。最后，我们使用sw访问StringWriter结构体中的字段str，并打印它的值。

#### 十八、Go语言接口

Go语言提供了另外一种类型，即接口，它把所有的具有共性方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
接口可以让我们将不同的类型绑定到一组公共的方法上，从而实现多态和灵活的设计。
Go语言中的接口时隐式实现的，也就是说，如果一个类型实现了一个接口定义的所有方法，那么它就会自动的实现该接口。因此，我们可以通过将接口作为参数实现对不同类型的调用，从而实现多态。

```go
实例
//定义接口
type interface_name interface {
    mothod_name [retrun_type]
    mothod_name [retrun_type]
    mothod_name [retrun_type]
    ...
    mothod_name [retrun_type]
}

//定义结构体
type struct_name struct {
    //variable
}
//实现接口方法
func (struct_name_variable struct_name) method_name [return_type]{
    //方法实现
}
...
func (struct_name_variable struct_name) method_name [return_type]{
    //方法实现
}
```

实例：以下两个方法演示接口的使用：

```go
package main
import "fmt"

type Phone interface {
    call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia,I can call you")
}

type IPhone struct {
}
func (iPhone IPhone) call() {
    fmt.Println("I am iPhone,I can call you")
}

func main() {
    var phone Phone
    phone = new(NokiaPhone)
    phone.call()
    
    phone = new(IPhone)
    phone.call()
}
```

上面的例子中，我们定义了一个接口Phone，接口里面有一个方法`call()`。然后我们在main函数里面定义了一个Phone类型变量，并分别为值赋值**NokiaPhone**和**IPhone**然后调用`call()`方法，输出结果如下：

```
I am Nokia,I can call you
I am iPhone,I can call you
```

第二个口实例：

```go
package main
import "fmt"

type Shape interface {
    area()	float64	
}

type Rectangle struct {
    width float64
    heigth float64
}

func (r Rectangle) area() float64 {
    return r.width * r.heigth
}

type Circle struct {
    radius float64
}

func (c Circle) area() float64 {
    return 3.14 * c.radius * c.radius
}

func main() {
    var s Shape
    
    s = Rectangle{width:10,heigth:5}
    fmt.Printf("矩形面积：%f\n",s.area())
    
    s = Circle{3}
    fmt.Printf("圆形面积：%f\n",s.area())
}
```

控制台输出结果为：

```
矩形面积：50.000000
圆形面积：28.260000
```

#### 十九、Go错误处理

Go语言通过内置的错误接口提供了非常简单的错误处理机制。
error类型时一个接口类，这是它的定义：

```go
type error interface {
	Error() string
}
```

我们可以在编码中通过实现error接口类型来生成错误信息。
函数通常在最后的返回值中返回错误信息。使用errors.New可返回一个错误信息：

```go
func Sqrt(f float64) (float64,error) {
    if f < 0 {
        return 0,error.New("math : square root of negative number")
    }	
    //实现
}
```

在下面的例子中，我们在调用Sqrt的时候传递的一个负数，然后就得到了non-nil的error对象，将此对象与nil比较，结果为true，所以fmt.Println(fmt包在处理error时会调用Erroe方法)被调用，以输出错误，请看下面调用的示例代码：

```go
result,err := Sqrt(-1)

if err != nil {
	fmt.Print(err)
}
```

实例

```go
package main
import "fmt"

//定义一个DivideError结构
type DivideError struct {
    dividee int
    divider int
}

//实现'error'接口
func (de *DivideError) Error() string {
    strFormat := `
    	Cannot proceed ,the divider is zero.
    	dividee: %d
    	divider: 0
    `
    return fmt.Sprintf(strFormat,de.dividee)
}

//定义'int'类型除法运算函数
func Divide(varDividee int,varDivider int) (result int,errorMsg string) {
    if varDivider == 0 {
        dData := DivideError{
            dividee : varDividee,
            divider : varDivider,
        }
        errorMsg = dData.Error()
        return
    } else {
        return varDividee / varDivider,""
    }
}

func main() {
	//正常情况
    if result,errorMsg := Divide(100,10);errorMsg == "" {
        fmt.Println("100 / 10 = ",result)
    }
    
    //但除数为零的时候会返回错误信息
    if _,errorMsg := Divide(100,0);errorMsg != "" {
        fmt.Println("errorMsg is: ", errorMsg)
    }
}
```

控制台输出结果为：

```
100 / 10 =  10
errorMsg is:
        Cannot proceed ,the divider is zero.
        dividee: 100
        divider: 0
```

介绍一下**panic**与**recover**，一个用于主动抛出错误，一个用于捕获panic抛出的错误。
**概念**

panic与recover是Go的两个内置函数，这两个内置函数用于处理Go运行时的错误，panic用于主动抛出错误，recover用来捕获panic抛出的错误。

#### 二十、Go并发

Go语言支持并发，我们只需要通过go关键字来开启goroutine即可。
goroutine是轻量级线程，goroutine的调度是由Golang运行时进行管理的。
goroutine语法格式：

```
go 函数名(参数列表)
```

例如

```
go f(x,y)
```

开启一个新的goroutine

```
f(x,y,x)
```

GO允许使用go语句开启一个新的运行期线程，即goroutine，以一个不同的、新建的goroutine来执行一个函数，同一个程序中的所有的goroutine共享一个地址空间。

```go
实例
package main 
import (
	"fmt"
    "time"
)

func say(s string) {
    for i := 0;i < 5; i++ {
        time.Sleep(100 *time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("world")
    say("hello")
}
```

执行以上代码，你会看到输出的hello和world是没有固定先后顺序。因为它们是两个goroutine在执行

```
hello
world
world
hello
hello
world
world
hello
hello
```

**通道（channel）**

通道(channel)是用来传递数据的一个数据结构。
通道可用于量goroutine之间通过传递一个指定类型的值来同步运行和通讯。操作符 `<-`用于指定通道的方向，发送或界首市。如果未指定方向，则为双向通道。

```
ch <- v  //把v发送到通道ch
v := <- ch //从ch接收数据，并把值赋给v
```

声明一个通道很简单，我们使用chan关键字即可，通道在使用前必须先创建：

```
ch := make(chan int)
```

**注意：**默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。
以下实例通过两个goroutine来计算数字之和，goroutine完成计算后，它会计算两个结果的和：

```go
package main
import "fmt"

func sum(s []int,c chan int) {
    sum := 0
    for _,v :=range s {
        sum += v
    }
    c <- sum	//把sum发送到通道c
}

func main() {
    s := []int{7,2,8,-9,4,0}
    
    c := make(chan int)
    go sum(s[:len(s)/2],c)
    go sum(s[len(s)/2:],c)
    x,y := <-c, <-c	//从通道C中接收
    
    fmt.Println(x,y,x+y)
}
```

输出结果为：

```
-5 17 12
```

**通道缓冲区**

通道可以设置缓冲区，通过make的第二个参数指定缓冲区大小：

```
ch := make(chan int,100)
```

带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于异步状态，就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。
不过由于缓冲区的大小是有限的，所以还是必须有接收端来接收数据的，否则缓冲区一满，数据发送端就无法再发送数据了。
**注意：**然后通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。如果通道带缓冲，发送方则胡阻塞直到发送的值被拷贝到缓冲区内；如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。接收方在有值可以接收之前会一直阻塞。

```go
实例
package main 
import "fmt"

func main() {
    //这里我们定义了一个可以存储整数类型的带缓冲通道
    //缓冲区大小为2
    ch := make(chan int,2)
    
    //因为ch是带缓冲的通道，我们可以同时发送两个数据
    //而不用立刻需要去同步读取数据
    ch <- 1
    ch <- 2
    
    //获取这两个数据
    fmt.Println(<- ch)
    fmt.Println(<- ch)
}
```

执行输出结果为：

```
1
2
```

**Go遍历通道与关闭通道**

Go通过range关键字来实现遍历读取到的数据，类似于与数组或切片。格式如下：

```
v,ok := <-ch
```

如果通道接收不到数据后就ok为false，这是通道就可以使用`close()`函数来关闭

```go
实例
package main 
import "fmt"

func fibonacci(n int,c chan int) {
    x,y := 0,1
    for i := 0;i <n;i++ {
        c <- x
        x,y = y,x+y
    }
    close(c)
}

func main() {
    c := make(chan int,10)
    go fibonacci(cap(c),c)
    //range函数遍历每个从通道接收的数据，因为c在发送完10个
    //数据之后就关闭了通道，所以这里我们range函数在接收到10个数据
    //之后就接收了，如果上面的c通道不关闭，那么range函数就不会结束
    //从而在接收第11个数据的时候就阻塞了。
    for i := range c{
        fmt.Println(i)
    }
}
```

执行输出结果为：

```
0
1
1
2
3
5
8
13
21
34
```

#### 二十一、标签校验和常见标签

在Goland中，`binding`和`validate`都是结构体标签，用于指定结构体字段的验证规则。

`binding`标签通常用于Gin框架中的请求参数绑定，它定义了该字段在绑定请求参数时的验证规则。例如，`binding:"required"`表示该字段是必填的，如果请求参数中缺少该字段，则会返回一个HTTP 400错误响应。

示例代码：

```go
type User struct {
    Name string `json:"name" binding:"required"`
    Age  int    `json:"age" binding:"gte=18"`
}
go复制代码
```

在这个例子中，`Name`字段是必填的，`Age`字段的值必须大于等于18。

`validate`标签通常用于对结构体字段进行更加复杂的验证，例如验证邮箱、手机号等。`validate`标签需要使用第三方库如`go-playground/validator`等来实现。

示例代码：

```go
import "github.com/go-playground/validator/v10"

type User struct {
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
}

func main() {
    validate := validator.New()

    user := User{Name: "Alice", Email: "alice@example.com"}
    if err := validate.Struct(user); err != nil {
        panic(err)
    }
}
go复制代码
```

在这个例子中，`Name`字段是必填的，`Email`字段必须是一个合法的邮箱地址。

总的来说，`binding`和`validate`都是用于结构体字段验证的标签，但是`binding`更加简单易用，适用于简单的验证场景，而`validate`则更加灵活，适用于复杂的验证场景。

**常见标签**

在 Go 语言中，标签（Tag）通常用于为结构体字段提供元数据信息，这些信息可以在运行时通过反射来访问。标签是一种以键值对形式附加到结构体字段上的字符串，放在字段的后面，例如：`json:"name"`。以下是一些常见的 Go 标签以及它们的用途：

1. **`json` 标签**：用于指定结构体字段在 JSON 序列化和反序列化时的行为。例如，`json:"name"` 可以指定 JSON 键名为 "name"。
2. **`xml` 标签**：用于指定结构体字段在 XML 序列化和反序列化时的行为。
3. **`bson` 标签**：用于指定结构体字段在 BSON 格式（MongoDB 中常用）序列化和反序列化时的行为。
4. **`db` 标签**：用于指定结构体字段与数据库表字段的映射关系。在 ORM 框架中经常使用。
5. **`gorm` 标签**：类似于 `db` 标签，用于 GORM ORM 框架中，指定字段与数据库表字段的映射关系以及其他配置选项。
6. **`form` 标签**：用于表单数据的绑定，通常与 Web 框架中的`post`请求处理有关。
7. **`validate` 标签**：通常与数据验证库一起使用，用于指定字段的验证规则。
8. **`protobuf` 标签**：用于 Protocol Buffers（protobuf）消息的字段映射。
9. **`yaml` 标签**：用于 YAML 格式的序列化和反序列化。
10. **`toml` 标签**：用于 TOML 格式的序列化和反序列化。
11. **`asn1` 标签**：用于 ASN.1 格式的编码和解码。
12. **`swaggertype` 标签**：与 Swagger 文档生成相关，用于指定字段的类型信息。
13. **`mapstructure` 标签**：通常与 `mapstructure` 库一起使用，用于将数据映射到结构体字段。

#### 二十二、gin框架请求头

在Gin框架中，`c *gin.Context`表示一个HTTP请求的上下文，它包含了HTTP请求的所有信息，例如请求方法、请求头、请求参数、请求体等等。在处理HTTP请求时，我们通常需要使用`c *gin.Context`来获取这些信息，并且根据请求的内容来生成响应。

`c *gin.Context`提供了一系列方法来获取HTTP请求的信息，例如：

- `c.Request.Method`：获取HTTP请求的方法（GET、POST、PUT、DELETE等）
- `c.Request.Header`：获取HTTP请求的头部信息
- `c.Query("key")`：获取HTTP请求的URL参数
- `c.Param("id")`：获取HTTP请求的参数
- `c.PostForm("key")`：获取HTTP请求的表单参数
- `c.BindJSON(&obj)`：将HTTP请求的JSON数据绑定到一个结构体中
- `c.ShouldBind(&obj)`：将HTTP请求的参数绑定到一个结构体中（自动根据Content-Type来判断参数类型）
- `c.String(code int, format string, values ...interface{})`：生成一个字符串响应
- `c.JSON(code int, obj interface{})`：生成一个JSON响应

除了获取HTTP请求的信息之外，`c *gin.Context`还提供了一些方法来设置HTTP响应的信息，例如：

- `c.Header("key", "value")`：设置HTTP响应的头部信息
- `c.SetCookie("name", "value", maxAge, "/", "domain", secure, httpOnly)`：设置HTTP响应的Cookie信息
- `c.Redirect(code int, location string)`：生成一个HTTP重定向响应
- `c.AbortWithStatus(code int)`：终止HTTP请求并返回指定的状态码

总之，`c *gin.Context`是Gin框架中非常重要的一个组件，它提供了处理HTTP请求和生成HTTP响应的一系列方法，是我们开发Gin应用程序的重要工具。

#### 二十三、接口实现和多态

在Golang中，接口是一种类型，它定义了一组方法的签名。任何实现了接口中定义的所有方法的类型，都可以被认为是实现了该接口。接口可以被用来实现多态，使得我们可以在不知道具体类型的情况下，调用实现了接口的对象的方法。

下面是一个简单的接口定义的例子：

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
go复制代码
```

上面的代码定义了一个`Shape`接口，它包含了两个方法：`Area()`和`Perimeter()`，这两个方法都返回一个`float64`类型的值。任何实现了`Shape`接口中定义的这两个方法的类型，都可以被认为是实现了`Shape`接口。

下面是一个实现了`Shape`接口的`Rectangle`类型的例子：

```go
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}
go复制代码
```

上面的代码定义了一个`Rectangle`类型，它包含了两个字段`Width`和`Height`，以及实现了`Shape`接口中定义的`Area()`和`Perimeter()`方法。因此，`Rectangle`类型可以被认为是实现了`Shape`接口。

在Golang中，我们可以使用类型断言来判断一个对象是否实现了某个接口。例如：

```go
func printShapeInfo(s Shape) {
    fmt.Printf("Area: %f, Perimeter: %f\n", s.Area(), s.Perimeter())
}

func main() {
    r := Rectangle{Width: 10, Height: 5}
    printShapeInfo(r)
}
go复制代码
```

上面的代码中，我们定义了一个`printShapeInfo()`函数，它接受一个`Shape`类型的参数，并打印出该对象的面积和周长。在`main()`函数中，我们创建了一个`Rectangle`类型的对象`r`，并将其传递给`printShapeInfo()`函数。由于`Rectangle`类型实现了`Shape`接口，因此`printShapeInfo()`函数可以正常地打印出`r`的面积和周长。

总之，接口是Golang中非常重要的一个特性，它可以帮助我们实现多态，使得我们可以在不知道具体类型的情况下，调用实现了接口的对象的方法。



#### 二十四、时间格式化规则

在 Go 语言中，时间格式化是基于特定的格式化字符串来定义的。以下是常用的时间格式化规则：

| 符号        | 说明                                     | 示例    |
| ----------- | ---------------------------------------- | ------- |
| 2006        | 年份（四位数）                           | 2023    |
| 01          | 年份（两位数）                           | 23      |
| 02          | 月份（带前导零）                         | 01      |
| 1           | 月份（不带前导零）                       | 1       |
| Jan         | 月份（英文缩写）                         | Jan     |
| January     | 月份（英文全称）                         | January |
| _2          | 月份的日期（带前导零）                   | 02      |
| 2           | 月份的日期（不带前导零）                 | 2       |
| 15          | 小时（24小时制，带前导零）               | 15      |
| 3           | 小时（12小时制，不带前导零）             | 3       |
| 04          | 分钟（带前导零）                         | 04      |
| 5           | 分钟（不带前导零）                       | 5       |
| 05          | 秒钟（带前导零）                         | 05      |
| 5           | 秒钟（不带前导零）                       | 5       |
| .000        | 纳秒（固定三位小数点后的数字）           | .000    |
| -0700       | 时区偏移量（数字表示）                   | -0700   |
| -07:00      | 时区偏移量（冒号分隔）                   | -07:00  |
| MST         | 时区缩写                                 | MST     |
| Mon         | 星期几（英文缩写）                       | Mon     |
| Monday      | 星期几（英文全称）                       | Monday  |
| PM          | 上午/下午标记                            | PM      |
| pm          | 上午/下午标记（小写）                    | pm      |
| .999999     | 微秒（固定六位小数点后的数字）           | .999999 |
| -07         | 时区偏移量（数字表示，不包含冒号）       | -7      |
| -7          | 时区偏移量（数字表示，不带前导零和冒号） | -7      |
| -07:00      | 时区偏移量（冒号分隔，不带前导零）       | -7:00   |
| -7:00       | 时区偏移量（冒号分隔，不包含前导零）     | -7:00   |
| -07Z        | 时区偏移量（数字表示，带 'Z' 后缀）      | -7Z     |
| -0700Z07:00 | 时区偏移量（数字表示，带 'Z' 后缀和冒号  |         |

**指定时区：**`location,err := time.LoadLocation("Asia/Shanghai")`   **东八区固定写法**

**这段代码声明了一些全局变量：**

- `GVA_DB` 是一个 `*gorm.DB` 类型的变量，用于表示数据库连接对象。
- `GVA_DBList` 是一个 `map[string]*gorm.DB` 类型的变量，用于存储多个数据库连接对象。
- `GVA_REDIS` 是一个 `*redis.Client` 类型的变量，用于表示 Redis 客户端连接对象。
- `GVA_CONFIG` 是一个 `config.Server` 类型的变量，用于存储服务器配置信息。
- `GVA_VP` 是一个 `*viper.Viper` 类型的变量，用于处理配置文件。
- `GVA_LOG` 是一个 `*zap.Logger` 类型的变量，用于记录日志。
- `GVA_Timer` 是一个 `timer.Timer` 类型的变量，用于处理定时任务。
- `GVA_Concurrency_Control` 是一个 `singleflight.Group` 类型的变量，用于并发控制。
- `BlackCache` 是一个 `local_cache.Cache` 类型的变量，用于表示缓存对象。
- `lock` 是一个 `sync.RWMutex` 类型的变量，用于控制并发访问。

这些全局变量在程序中被用于存储和管理数据库连接、缓存、配置、日志等重要的共享资源。它们可以在程序的不同模块中被访问和使用，以提供全局范围的功能和数据共享

**字符串转换成time.Time类型**

定义一个`layout  := "2006-01-02`"固定格式，相当于"yyyy-mm-dd"
定义一个时区，东八区`:location,_  := time.LoadLocation("Asia/Shanghai")`,固定写法
然后解析字符串转化为时间类型:
`parse,_ := time.ParseIntLocation(layout,string,location)`

#### 二十五、go单元测试

文件名要以`_test.go`结尾，如(practive_test.go),测试运行，方法名如`TestAdd()`，参数是testing下的，(t *testing.T).

#### 二十六、零值

官方文档中零值称为`zero value`，零值并不仅仅只是字面上的数字零，而是一个类型的空值或者说默认值更为准确。

| 类型                                 | 零值                         |
| ------------------------------------ | ---------------------------- |
| 数字类型                             | `0`                          |
| 布尔类型                             | `false`                      |
| 字符串类型                           | `""`                         |
| 数组                                 | 固定长度的对应类型的零值集合 |
| 结构体                               | 内部字段都是零值的结构体     |
| 切片，映射表，函数，接口，通道，指针 | `nil`                        |

Golang 更明确的数字类型命名，支持 Unicode，支持常用数据结构。

| 类型          | 长度(字节) | 默认值 | 说明                                      |
| ------------- | ---------- | ------ | ----------------------------------------- |
| bool          | 1          | false  |                                           |
| byte          | 1          | 0      | uint8                                     |
| rune          | 4          | 0      | Unicode Code Point, int32                 |
| int, uint     | 4或8       | 0      | 32 或 64 位                               |
| int8, uint8   | 1          | 0      | -128 ~ 127, 0 ~ 255，byte是uint8 的别名   |
| int16, uint16 | 2          | 0      | -32768 ~ 32767, 0 ~ 65535                 |
| int32, uint32 | 4          | 0      | -21亿~ 21亿, 0 ~ 42亿，rune是int32 的别名 |
| int64, uint64 | 8          | 0      |                                           |
| float32       | 4          | 0.0    |                                           |
| float64       | 8          | 0.0    |                                           |
| complex64     | 8          |        |                                           |
| complex128    | 16         |        |                                           |
| uintptr       | 4或8       |        | 以存储指针的 uint32 或 uint64 整数        |
| array         |            |        | 值类型                                    |
| struct        |            |        | 值类型                                    |
| string        |            | ""     | UTF-8 字符串                              |
| slice         |            | nil    | 引用类型                                  |
| map           |            | nil    | 引用类型                                  |
| channel       |            | nil    | 引用类型                                  |
| interface     |            | nil    | 接口                                      |
| function      |            | nil    | 函数                                      |

#### 二十七、输出

输出一句`Hello 世界!`，比较常用的有三种方法，第一种是调用`os.Stdout`

```go
os.Stdout.WriteString("Hello 世界!")
```



第二种是使用内置函数`println`

```go
println("Hello 世界!")
```



第三种也是最推荐的一种就是调用`fmt`包下的`Println`函数

```go
fmt.Println("Hello 世界!")
```

`fmt.Println`会用到反射，因此输出的内容通常更容易使人阅读，不过性能很差强人意。

#### 二十八、格式化

| 0    | 格式化    | 描述                                           | 接收类型           |
| ---- | --------- | ---------------------------------------------- | ------------------ |
| 1    | **`%%`**  | 输出百分号`%`                                  | `任意类型`         |
| 2    | **`%s`**  | 输出`string`/`[] byte`值                       | `string`,`[] byte` |
| 3    | **`%q`**  | 格式化字符串，输出的字符串两端有双引号`""`     | `string`,`[] byte` |
| 4    | **`%d`**  | 输出十进制整型值                               | `整型类型`         |
| 5    | **`%f`**  | 输出浮点数                                     | `浮点类型`         |
| 6    | **`%e`**  | 输出科学计数法形式 ,也可以用于复数             | `浮点类型`         |
| 7    | **`%E`**  | 与`%e`相同                                     | `浮点类型`         |
| 8    | **`%g`**  | 根据实际情况判断输出`%f`或者`%e`,会去掉多余的0 | `浮点类型`         |
| 9    | **`%b`**  | 输出整型的二进制表现形式                       | `数字类型`         |
| 10   | **`%#b`** | 输出二进制完整的表现形式                       | `数字类型`         |
| 11   | **`%o`**  | 输出整型的八进制表示                           | `整型`             |
| 12   | **`%#o`** | 输出整型的完整八进制表示                       | `整型`             |
| 13   | **`%x`**  | 输出整型的小写十六进制表示                     | `数字类型`         |
| 14   | **`%#x`** | 输出整型的完整小写十六进制表示                 | `数字类型`         |
| 15   | **`%X`**  | 输出整型的大写十六进制表示                     | `数字类型`         |
| 16   | **`%#X`** | 输出整型的完整大写十六进制表示                 | `数字类型`         |
| 17   | **`%v`**  | 输出值原本的形式，多用于数据结构的输出         | `任意类型`         |
| 18   | **`%+v`** | 输出结构体时将加上字段名                       | `任意类型`         |
| 19   | **`%#v`** | 输出完整Go语法格式的值                         | `任意类型`         |
| 20   | **`%t`**  | 输出布尔值                                     | `布尔类型`         |
| 21   | **`%T`**  | 输出值对应的Go语言类型值                       | `任意类型`         |
| 22   | **`%c`**  | 输出Unicode码对应的字符                        | `int32`            |
| 23   | **`%U`**  | 输出字符对应的Unicode码                        | `rune`,`byte`      |
| 24   | **`%p`**  | 输出指针所指向的地址                           | `指针类型`         |

#### 二十九、闭包

闭包(Closure)这一概念，在一些语言中又被称为Lamda表达式，经常与匿名函数一起使用，`函数+环境引用 = 闭包`。看一个文件名_test.go结尾的单元测试例子：

```go
packege main

func TestSum(t *testing.T){
    sum :=Sum(1,2)
    fmt.Println(sum(3,4))
    fmt.Println(sum(5,6))
}

func Sum(a,b int) func(int, int) int {
    return func(int, int) int{
        return a + b
    }
}
```

```
3
3
```

在上述代码中，无论传入什么数字，输出结果都是3，返回的是一个sum函数，传3和4进去，函数做的操作只是将a和b相加，传进去的3和4并不是a和b的值，如果想将传过去的形参被使用，你可以给形参起名字，修改后代码如下所示：

```go
packege main

func TestSum(t *testing.T){
    sum :=Sum(1,2)
    fmt.Println(sum(3,4))
    fmt.Println(sum(5,6))
}

func Sum(a,b int) func(c int, d int) int {
    return func(c int, d int) int{
        return a + b + c + d
    }
}
```

```
10
14
```

匿名函数引用参数a,b，即便Sum函数已经执行完毕，虽然已经超出了它的生命周期，但是对其返回的函数传入参数，依然可以成功的引用a，b的值，这一个过程就是闭包。事实上a，b已经逃逸到了堆上，只要其返回函数的生命周期没有结束，就不会被回收掉。

**延时调用**

`defer`关键字描述的一个匿名函数会在函数返回之前(即return之前)执行，下面是一个案例:

```go
packege main
import "fmt"

func TestDo(t *testing.T){
    Do()
}

func Do(){
    defer func(){
        fmt.Println("1")
    }()
    fmt.Println("2")
}
```

输出结果为：

```
2
1
```

当有多个defer语句时，会按照后进先出的顺序执行。如数据结构中的栈，下面是一个案例：

```go
packege main
import "fmt"

func TestDo(t *testing.T){
    Do()
}

func Do(){
    defer func(){
        fmt.Println("1")
    }()
    defer func(){
        fmt.Println("2")
    }()
    defer func(){
        fmt.Println("3")
    }()
    defer func(){
        fmt.Println("4")
    }()
    fmt.Println("2")
    defer func(){
        fmt.Println("5")
    }()
}
```

控制台输出结果为：

```
2
5
4
3
2
1
```

**延迟调用通常用于释放文件资源，关闭连接等操作，另一个常用的写法是用于捕获异常`panic`。**

#### 三十、基础知识回顾

在Go中，我们可以使用`unsafe.Sizeof(x)`来查看变量所占的内存大小。以下是Go内置的数据类型占用的内存大小：

| 类型                        | 内存大小（字节数） |
| :-------------------------- | :----------------- |
| bool                        | 1                  |
| int8/uint8                  | 1                  |
| int/uint                    | 8                  |
| int32/uint32                | 4                  |
| int64/uint64                | 8                  |
| float32                     | 4                  |
| float64                     | 8                  |
| complex64                   | 8                  |
| complex128                  | 16                 |
| 指针类型：*T, map,func,chan | 8                  |
| string                      | 16                 |
| interface                   | 16                 |
| []T                         | 24                 |

协程：用户态，轻量级线程，栈KB级别，创建和调度由go语言直接调度
线程：内核态，线程跑多个协程，栈MB级别

- **Go协程goroutine: 是一种轻量线程，它不是操作系统的线程，而是将一个操作系统线程分段使用，通过调度器实现协作式调度。是一种绿色线程，微线程，它与Coroutine协程也有区别，能够在发现堵塞后启动新的微线程。**
- **通道channel: 类似Unix的Pipe，用于协程之间通讯和同步。协程之间虽然解耦，但是它们和Channel有着耦合。**

#### 三十一、Go数据类型|slice进阶

**1.1源码实现**

slice是基于array实现的，它的底层是array，可以理解为对底层array的抽象。源码包中src/runtime、silice.go定义了slice的数据结构：

```go
//$GOROOT/src/runtime/slice.go
type slice struct {
	array unsafe.Pointer
    len int
    cap int
}
```

可以看到`slice`占用24个字节，其中：

- `array`：指向底层数据的指针，占用8个字节；
- `len`：切片的长度，占用8个字节
- `cap`：切片的容量，`cap`总是大于等于`len`的，占用8个字节。

**1.2初始化时的底层实现**

slice共有4种初始化方式：

```
//1、直接声明
var slice1 []int

//2、使用字面量
slice2 := []int{1,2,3,4}

//3、使用make创建slice
slice3 := make([]int,3,5)

//4、从切片或数组截取
slice4 := arr[1:3]
```

我们可以通过 go tool compile -S test.go | grep CALL 得到汇编代码：

```
$ go tool compile -S main.go| grep CALL
0x002c 00044 (main.go:6)        CALL    runtime.makeslice(SB)
0x0050 00080 (main.go:7)        CALL    runtime.growslice(SB)
0x0080 00128 (main.go:8)        CALL    runtime.convTslice(SB)
0x0094 00148 (main.go:8)        CALL    runtime.convT64(SB)
0x00a8 00168 (main.go:8)        CALL    runtime.convT64(SB)
0x011c 00284 ($GOROOT/src/fmt/print.go:274)     CALL    fmt.Fprintln(SB)
0x0130 00304 (main.go:5)        CALL    runtime.morestack_noctxt(SB)
```

从汇编代码中我们可以知道初始化 slice 底层调用的是 runtime.makeslice，makeslice 函数的工作主要就是计算 slice 所需内存大小，然后调用 mallocgc 进行内存的分配。

所需的内存大小 = 切片中元素大小 * 切片的容量。

以下是 makeslice 源码：

```go
fun makeslice(et *_type,len,cap int) unsafe.Pointer {
    mem,overflow := math.MulUintptr(et.size,unitptr(cap))
    if overflow || mem > maxAlloc || len < 0 || len > cap {
        mem,overflow := math.MulUintptr(et.size,unitptr(len))
        if ovaeflow || mem > maxAlloc || len > 0 {
            panicmakeslicelen()
        }
        paincamkeslicecap()
    }
    
    return mallocgc(mem,et,true)
}
```

`makeslice`输入参数包括元素整型指针`et`、长度`len`、容量`cap`。这个函数的返回值就是一个指向新slice内存空间的指针。

在函数内部，首先使用`math.MulUintptr`ha念书计算出要分配的内存大小(即`et.size * cap`):

```go
fun MulUintptr(a,b uinptr) (uintpr,bool){
    if a|b < 1 <<(4*goarch.PtrSize) || a == 0 {
        return a*b,false
    }
    overflow := b > MaxUintptr/a
    return a*b,overflow
}
```

其中`a|b < 1 <<  (4*goarch.PtrSize)`这段代码是一个Go语言中的条件语句，用于比较`a`和`b`的值是否都小于`(4*goarch.PtrSize)`，如果都小于，那么返回`true`。其中`goarch.PtrSize`是一个常量，表示指针类型在当前系统架构下的字节大小。

具体来说，右侧表达式`1 <<  (4*goarch.PtrSize)`表示将二进制数1左移`(4*goarch.PtrSize)`位，相当于将其乘以`2^(4*goarch.PtrSize)`。这样就得到了一个大于等于1的整数，在与a|b进行位运算后，可以得到它们的一部分二进制位，进而比较它们的大小关系。

回到`makeslice`，通过`et.size*cap`与`maxAlloc`常量的比较和其他条件的检查，确保分配的内存不会超出系统限制或者非法操作。如果检查失败，则会抛出一个异常，否则调用`mallocgc`ha函数分配内存，并将分配到的内存指针返回调用方。

**2.slice 是线程安全的吗？**
不是。

我们可以简单写个函数验证一下：

```go
func main() {
    slice := make([]int, 0)
	for i := 0; i < 10; i++ {     
    	go func() {
        	slice = append(slice, 1)
    	}()
	}

	for i := 0; i < 10; i++ {
   		go func() {
        	if len(slice) > 0 {
            	slice = slice[:len(slice)-1]
        	}
    	}()
}

	time.Sleep(time.Second * 2)

	fmt.Println("length of slice:", len(slice))
}
```
**可以发现运行后每次结果都不同，说明 slice 并不支持并发读写。**

**从 slice 的底层结构也可以看出，因为它没有使用锁等方式。**

**slice 在并发执行中不会报错，但是数据会丢失。**

**3.slice的扩容机制**

**3.1Go 1.17及更早版本**

1. 如果期望容量大于当前容量的2倍，那么扩容的容量大小为期望容量；

2. 否则：

   - 如果原始容量小于1024，那么扩容后的容量为原始容量的2倍；
   - 如果原始容量大于等于1024，那么会进入一个循环，每次循环扩容1.25倍，知道扩容后的容量大于期望容量，如果循环过程中发生整数溢出，则将扩容后的容量为期望容量。

   相关源码：

   ```go
   // $GOROOT/src/runtime/slice.go
   //old.cap	原始容量
   //cap 期望容量
   //doublecap  当前容量的两倍
   newcap := old.cap
   doublecap := newcap + newcap
   //如果期望的容量大于当前容量的两倍，那么扩容的容量大小为期望容量
   if cap > doublecap {
       newcap = cap
   } else {
       //如果原始容量小于1024，则扩容的容量为原始容量的2倍
       if old.cap < 1024 {
           newcap = doublecap
       } else {
       //不小于1024的话，进入一个for循环，结束for循环的条件是0小于当前容量和小于期望容量，每次做1.25倍扩容直到大于期望容量
           for 0 < newcap && newcap < cap {
               newcap += newcap/4
           }
           if newcap <= 0 {
               newcap = cap
           }
           
       }
   }
   ```

**3.2Go 1.18至Go 1.20版本**

1. 如果期望容量大于当前容量的2倍，那么扩容后的容量大小为期望容量；
2. 否则：
   - 如果原始容量大于等于256，那么扩容后的熔炼为原始容量的2倍；
   - 如果原始容量大于256，那么会进入一个循环，每次循环都会增加`（当前容量 + 3 * 256）/ 4`的容量，这条公式意味这从2倍扩容到1.25倍扩容的平滑过渡，如果循环过程发生整数溢出，则将扩容后的容量置为期望容量。

​	相关源码

```go
// $GOROOT/src/runtime/slice.go
//old.cap	原始容量
//cap 期望容量
//doublecap  当前容量的两倍
newcap := old.cap
doublecap := newcap + newcap
//如果期望的容量大于当前容量的两倍，那么扩容的容量大小为期望容量
if cap > doublecap {
    newcap = cap
} else {
	const threshod = 256
    //如果原始容量小于256，则扩容的容量为原始容量的2倍
    if old.cap < 256 {
        newcap = doublecap
    } else {
    //不小于256的话，进入一个for循环，结束for循环的条件是0小于当前容量和小于期望容量，每次做1.25倍扩容直到大于期望容量
        for 0 < newcap && newcap < cap {
            newcap += (newcap + 3 * threshod)/ 4
        }
        if newcap <= 0 {
            newcap = cap
        }
        
    }
}
```

**4.array和slice的区别**

1. array是值类型，slice是引用类型。所以在作为函数参数传递时，array会拷贝整个数组，slice传递的时指针；
2. array是定长的，slice是变长的，并且总是指向底层的array。

**5.slice的深拷贝和浅拷贝**

``浅拷贝是指在复制对象时，只复制对象本身和其中包含的基本类型数据，而不会复制对象所引用的其他对象，也就是说，在浅拷贝中，复制出的新对象与元对象共用一些引用对象。因此，如果修改了其中的一个对象中的引用对象，令一个对象也会随之改变。``
`深拷贝则意味着在复制对象时，除了复制对象本身和其中包含的基本类型数据外，还会递归地复制对象所引用的其他对象。也就是说，在深拷贝中，复制出的新对象和原对象时是完全独立的，它们之间没有任何引用关系。因此，即使修改其中一个对象的引用对象，另一个对象也不会受到影响。`

实现深拷贝的方式：1.copy函数；2.遍历slice再赋值

实现浅拷贝的方式：1.默认的赋值操作。

****

**切片总结**

1.什么是切片？

slice不是数组或数组指针。它通过内部指针和相关属性引用数组片段，以实现变长方案。

- 切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值考本传递。
- 切片的长度是可以该改变的，所有又称动态数组
- 如果slice == nil，那么len()、cap()结果都等于0
- cap可以求出slice最大扩张容量，不能超出数组限制



2.什么时候slice == nil ？为什么会出现这种情况？

- 首先这个涉及到变量声明和内存的分配问题，如果使用var关键字声明有两种情况，如果是值类型，系统会默认分配内存空间，并赋值为零值，如果是指针类型或切片类型，系统不会分配默认内存的，底层是一个对数组地址引用的指针，零值默认为nil，表示并未引用任何有效的底层数组，会出现slice == nil，使用make函数可以进行底层数组的内存分配并进行初始化赋值零值。

3.切片的扩容在什么时候发生？

在Go中，切片的扩容是再Go的运行时系统控制的。切片的容量策略是在达到其容量上限时开始扩张的。当切片的长度(len)达到其容量(cap)时，再添加新元素到切片中时，Go会自动进行容量扩张。扩容到一定倍数时，这个倍数时根据切片的大小动态调整的。

#### 三十二、空接口类型的支持

在Go语言中，空接口是一个特殊的接口类型，也被称为任意类型。空接口不包含任何方法，因此可以表示任意类型的值。空接口的定义非常简单，它没有任何方法声明：

```go
interface{}
```

由于空接口不包含任何方法，因此它可以接收任何类型的值。这使得空接口在需要处理不同类型的值的情况下非常有用，因为我们无需提前指定具体的类型。

以下是一个简单的示例来展示空接口的用法：

```go
package main

import "fmt"

func printValue(v interface{}) {
        fmt.Println(v)
}

func main() {
        printValue(42)           // 输出 42
        printValue("Hello")      // 输出 Hello
        printValue(3.14)         // 输出 3.14
        printValue([]int{1, 2, 3}) // 输出 [1 2 3]
}
```

在这个示例中，我们定义了一个函数 `printValue`，它接收一个空接口类型的参数 `v`。在函数内部，我们直接通过 `fmt.Println` 打印了接收到的值 `v`。通过将不同类型的值传递给 `printValue` 函数，我们可以看到它可以接收任意类型的值，并打印出对应的结果。

使用空接口时需要注意的是，由于空接口可以接收任意类型的值，因此在使用其内部的值时，我们需要进行类型断言或类型判断，以确定其具体类型并进行相应的操作。

```go
package main

import "fmt"

func processValue(v interface{}) {
        if str, ok := v.(string); ok {
                fmt.Println("Received a string:", str)
        } else if num, ok := v.(int); ok {
                fmt.Println("Received an integer:", num)
        } else {
                fmt.Println("Received an unknown type")
        }
}

func main() {
        processValue("Hello")  // 输出 "Received a string: Hello"
        processValue(42)       // 输出 "Received an integer: 42"
        processValue(true)     // 输出 "Received an unknown type"
        processValue(3.14)     // 输出 "Received an unknown type"
        processValue([]int{1, 2, 3}) // 输出 "Received an unknown type"
}
```

在这个示例中，我们定义了一个函数 `processValue`，它接收一个空接口类型的参数 `v`。在函数内部，我们使用类型断言来判断 `v` 的具体类型，并根据类型执行相应的操作。

在 `if` 语句中，我们使用 `t, ok := v.(type)` 来进行类型断言，将 `v` 转换为 目标 `type` 类型，并将转换后的值存储在`t` 中。如果转换成功，`ok` 的值为 `true`，我们就可以执行对应的操作。如果转换失败，那么 `ok` 的值为 `false`，表示 `v` 不是目标类型。

总结而言，`Go`语言中的空接口是一种特殊的接口类型，它不包含任何方法，可以表示任意类型的值。空接口在需要处理不同类型的值的情况下非常有用，但在使用时需要注意类型断言或类型判断。

#### 三十三、控制并发数量

在go语言中启动多个goroutine实现并发(这里使用sync.WaitGroup来实现goroutine的同步)。

```go
var wg sync.WaitGroup

func hello(i int){
    defer wg.Done() //goroutine结束就等级 -1
    fmt.PrintIn("Hello Gorouter!",1)
}

func mian(){
    
    for i := 0;i < 10;i++ {
        wg.add(1) //启动一个goroutine
        go hello(i)
    }
    wg.Wait()	//等待所有等级的goroutine都结束
}
```

**问题**

```go
func main(){
    userCount := math.MaxInt64
    for i:=0;i<userCount;i++ {
        go func(i int){
            //做一些各种各样的业务逻辑
            fmt.Printf("go func: %d\n",i)
            time.Sleep(time.Second)
        }
    }
}
```

在这里，假设`userCount`是一个外部传入的参数(不可预测，有可能非常大)，有人会全部丢进去循环。想着全部并发goroutine去同时做某一件事。觉得这样子会效率更高。

**噩梦般的开始**

当然，在特定的场景下，问题可大了。因为在本文被丢进去同时并发的可是一个极端值。

**输出结果**

```
...
go func: 5839
go func: 5840
go func: 5841
go func: 5842
go func: 5915
go func: 5524
go func: 5916
go func: 8209
go func: 8264
signal: killed
```

- 系统资源占用率不断上涨
- 输出一定数量后：控制台就不在输出最新的值了
- 信号量：signal：killed
- 短时间内系统负载暴增
- 短时间内占用的虚拟内存暴增

**小结**

就是在不控制并发的goroutine数量会发生什么问题？大致如下：

- CPU使用率浮动上涨
- Memory（内存）占用不断上涨。
- 主进程崩溃（被杀掉了）

简单来说，”崩溃“的原因就是对系统资源的占用过大。常见的比如：打开文件数（too many files open）、内存占用等等

**危害**

对该台服务器产生非常大的影响，影响自身及相关的应用。很有可能导致不可用或者响应缓慢，另外启动了复数”失控“的goroutine，导致程序流转混乱

**解决方案**

思考解决方案。如下：

1. 控制/限制goroutine同时并发运行的数量
2. 改变应用程序的逻辑写法(避免大规模的使用系统资源和等待)

**控制goroutine并发数量**

**尝试chan**

```go
func mian(){
    userCount := 10
    ch := make(chan bool,2)
    for i := 0;i<userCount;i++{
        ch <- true
        go Read(ch,i)
    }
    
    //time.Sleep(time.Second)
}

func Read(ch chan bool,i int){
    fmt.Printf("go func : %d\n",i)
    <- ch
}
```

输出结果：

```
go func : 1
go func : 2
go func : 3
go func : 4
go func : 0
go func : 5
go func : 7
go func : 8
go func : 6
go func : 9
```

我这个好像输出对应的值，好像有些会出现输出9个值，这种情况在当主协程结束时，子协程也是会被终止掉的。因此剩余的goroutine还没来得及输出，就被终止了(可以打开time.Sleep看看输出数量)

**尝试sync**

```go
var wg = sync.WaitGroup{}

func main(){
    userCount := 10
    for i:=0;i<userCount;i++{
        wg.Add(1)
        go Read(i)
    }
    wg.Wait()
}

func Read(i int){
    defer wg.Done()
    fmt.Printf("fo func : %d\n",i)
}
```

单纯的使用`sync.WaitGroup`也不行，没有控制到同时并发的goroutine数量(代指达不到本文所要求的目标)

**小结**

单纯简单使用channel或sync都有明显缺席，不行。我们再看看组件配合能不能实现

**尝试chan+sync**

```go
var wg = sync.WaitGroup{}

func mian(){
    userCount := 10
    ch := make(chan bool,2)
    for i := 0;i<userCount;i++{
        wg.Add(1)
        go Read(ch,i)
    }
    wg.Wait()
}

func Read(ch chan bool,i int){
    defer wg.Done()
    
    ch <- true
    fmt.Printf("go func: %d,time:%d\n",i,time.Now().Unix())
    time.Sleep(time.Second)
    <- ch
}
```

输出结果：

```
go func: 9, time: 1547911938
go func: 1, time: 1547911938
go func: 6, time: 1547911939
go func: 7, time: 1547911939
go func: 8, time: 1547911940
go func: 0, time: 1547911940
go func: 3, time: 1547911941
go func: 2, time: 1547911941
go func: 4, time: 1547911942
go func: 5, time: 1547911942
```

从输出结果来看，确实实现控制goroutine以2个2个的数量去执行我们的"业务逻辑"，当然结果集也理所当然的时乱序输出

**方案一：简单Semapore**

在确立了简单使用chan+sync的方案可行后，我们重新将流转逻辑封装为gsema。主程序变成如下：

```go
import(
	"fmt"
    "time"
    "github.com/EDDYCJY/gsema"
)

var sema = gsema.NewSemaphore(3)

func mian(){
    userCount := 10
    for i:=0;i<userCount;i++{
        go Read(i)
    }
    sema.Wait()
}

func Read(i int){
    defer sema.Done()
    sema.Add(1)
    
    fmt.Printf("go func :%d, time: %d\n",i,time.Now().unix())
    time.Sleep(time.Second)
}
```

**分析方案**

在上述代码中，程序流程如下：

- 设置允许的并发数目为3个
- 循环10次，每次启动一个goroutine来执行任务
- 每一个goroutine在内部利用sema进行调控是否阻塞
- 按允许并发数逐渐释出goroutine，最后结束任务

看上去没什么严重问题。单却有一个”大“坑，认真看到第二点”每次启动一个goroutine”这句话。这里有点问题，提前产生那么多的goroutine会不会有什么问题，接下来一起分析下利弊，如下：

**利**

- 适合量不大、复杂度低的使用场景
  - 几百个千个、几十万个也是可以接收的（看具体业务场景）
  - 实际业务逻辑在运行前就已经被阻塞等待了（因为并发数受限），基本实际业务逻辑损耗的性能比goroutine本身大
  - gotoutine本身很轻便，仅损耗少许的内存空间和调度。这种等待响应的情况都躺好了，等待任务唤醒
- Semaphore操作复杂度低且流转简单，容易控制

**弊**

- 不适合量很大。复杂度高的使用场景
  - 有几百万、几千万个goroutine的话，就浪费了大量调度goroutine和内存空间。恰好你的服务器也接受不了的话
- Semapore操作复杂度提高，要管理更多的状态

**小结**

- 基于什么业务场景，就用什么方案去做事
- 有足够的时间，允许你去追求跟优秀、极致的方案(用第三方库也行)

用哪种方案，我认为主要基于以上两点去思考，都是OK的。没有对错，只有当前业务场景能不能接受，这个预先启动的goroutine数量你的系统是否能够接受

**灵活控制goroutine并发数量**

在输出输入一体的请况下，在常见的业务场景中确实可以。但，如这次新的业务场景比较特殊，要控制输入的数量，以达到改变允许并发允许goroutine的数量。我们仔细想想，要做出如下改变：

- 输入/输出要抽离，才可以分别控制
- 输入/输出要可变，理所应当在for-loop中(可设置数值的地方)
- 允许改变goroutine并发数量，但它也必须有一个最大值(因为允许改变是相对)

**方案二：灵活chan+sync**

```go
package main

import (
	"fmt"
    "sync"
    "time"
)
var wg = sync.WaitGroup()
func main(){
    userCount := 10
    ch := make(chan int,5)
    for i := 0;i<userCount;i++{
        wg.Add(1)
        go func(){
            defer wg.Done()
            for d := range ch {
                fmt.Printf("go func :%d,time: %d\n",d,time.Now().Unix())
                time.Sleep(time.Second * time.Duration(d))
            }
        }()
    }
    for i :=0;i<10;i++{
        ch <- 1
        ch <- 2
        //time.Sleep(time.Second)
    }
    close(ch)
    wg.Wait()
}
```

输出结果：

```
go func :3,time: 1692690351
go func :1,time: 1692690351
go func :1,time: 1692690351
go func :2,time: 1692690351
go func :3,time: 1692690351
go func :2,time: 1692690351
go func :1,time: 1692690351
go func :2,time: 1692690351
go func :3,time: 1692690351
go func :1,time: 1692690351
go func :2,time: 1692690352
go func :3,time: 1692690352
go func :1,time: 1692690352
go func :2,time: 1692690352
go func :3,time: 1692690353
go func :1,time: 1692690353
go func :2,time: 1692690353
go func :3,time: 1692690353
go func :1,time: 1692690354
go func :2,time: 1692690354
go func :3,time: 1692690354
go func :1,time: 1692690354
go func :2,time: 1692690354
go func :3,time: 1692690354
go func :1,time: 1692690355
go func :2,time: 1692690355
go func :3,time: 1692690355
go func :1,time: 1692690355
go func :2,time: 1692690356
go func :3,time: 1692690356
```

在方案二中，我们可以随时地根据新的业务需求，做如下事情：

- 变更channel的输入数量
- 能够根据特殊情况，变更channel的循环值
- 变更最大允许并发的goroutine数量

总的来说，就是可控空间都尽量放开了，是不是更加灵活。

**方案三：第三方库**

[go-playground/pool](https://github.com/go-playground/pool)

[nozzle/throttler](https://github.com/nozzle/throttler)

[Jeffail/tunny](https://github.com/Jeffail/tunny)

[panjf2000/ants](https://github.com/panjf2000/ants)

比较成熟的第三方库也不少，基本都是以生成和管理 goroutine 为目标的池工具。我简单列了几个，具体建议大家阅读下源码或者多找找，原理相似

#### 三十四、变量在栈上还是堆上

在写代码的时候，有时候会想这个变量到底分配到哪里，是堆上还是在栈上。深挖一下这块的奥妙

**问题**

```go
type User struct {
    ID int64
    Name string
}

func GetUserInfo() *User {
    return &User{ID:123,Name:"EDDYCJY"}
}
func main(){
    _ = GetUserInfo()
}
```

带着问题学习。请问main调用`GetUserInfo`返回的`&User{...}`。这个变量是分配到栈上还是分配到堆上？
**什么是堆/栈**

在这里仅简单介绍所需的基础知识。如下：

- 堆(Heap)：一般来讲是人为手动进行管理，手动申请、分配、释放。一般所涉及的内存大小并不定，一般会存放比较大的对象。另外其分配相对慢，涉及到的指令动作也相对多
- 栈(Stack)：有编译器进行管理，自动申请、分配、释放。一般不会太大，我们常见的**函数参数**(不同平台允许存放的数量不同)，**局部变量**等等都会存放在栈上

介绍Go语言，它的堆栈分配时通过Compiler进行分析，GC去管理的，而对其的分析选择动作就是探讨的重点

**什么时逃逸分析**

在编译程序优化理论中，逃逸分析是一种确定指针动态范围的方法，简单来说就是分析在程序的哪些地方可以访问到指针，通俗的讲就是确定一个变量要放在堆上还是栈上，规则如下：

1. 是否有在其他地方(非局部)被引用。只有有可能被引用了，那么它一定分配到堆上。否则分配到栈上
2. 即便没有被外部引用，但对象过大，无法存放在栈区上，依然有可能分配到堆上

可以理解为，逃逸分析是编译器用于决定变量分配到堆上还是栈上的一种行为

**在上面阶段确立逃逸**

在编译阶段确立逃逸，注意并不是在运行时

**为什么需要逃逸**

这个问题可以反过来想，如果变量都分配到堆上会出现什么情况，例如：

- 垃圾回收(GC)的压力不断增大
- 申请、分配、回收内存的系统开销增大(相对于栈)
- 动态分配产生一定量的内存碎片

其实总的来说，就是频繁申请、分配堆内存是有一定"代价"的，会影响程序运行的效率，简介影响到整个系统。因此"按需分配"最大限度的灵活利用资源，才是正确的治理之道。这就是为什么需要逃逸分析的原因。

**怎么确定是否逃逸**

第一，通过编译器命令，就可以看到详细的逃逸分析过程。而指令集`-gcflags`用于将标识参数传入递给Go编译器，涉及如下：

- `-m`会打印逃逸分析的优化策略，实际上最多总共可以用4个`-m`，但是信息量较大，一般用1个就可以了
- `-l`会禁用函数内联，在这里禁掉inline能更好的观察逃逸情况，减少干扰

```
go build -gcflags '-m -l' main.go
```

第二，通过反编译命令查看

```
go tool compile -S main.go
```

注：可以通过`go tool compile -help`查看所有允许传递给编译器的标识参数 

```
PS C:\Users\FWR\Desktop\awesomeProject1> go build -gcflags '-m -l' main.go
# command-line-arguments                  
.\main.go:10:9: &User{...} escapes to heap
```

通过查看分析结果，可得知`&User`逃到了堆里，也就是分配到堆上了。这是不是有问题，再看看汇编代码确定一下，如下：

```
PS C:\Users\FWR\Desktop\awesomeProject1> go tool compile -S main.go
main.GetUserInfo STEXT size=103 args=0x0 locals=0x18 funcid=0x0 align=0x0                                               
        0x0000 00000 (C:\Users\FWR\Desktop\awesomeProject1\main.go:9)   TEXT    main.GetUserInfo(SB), ABIInternal, $24-0
        0x0000 00000 (C:\Users\FWR\Desktop\awesomeProject1\main.go:9)   CMPQ    SP, 16(R14)                             
        0x0004 00004 (C:\Users\FWR\Desktop\awesomeProject1\main.go:9)   PCDATA  $0, $-2
        0x0004 00004 (C:\Users\FWR\Desktop\awesomeProject1\main.go:9)   JLS     92
        0x0006 00006 (C:\Users\FWR\Desktop\awesomeProject1\main.go:9)   PCDATA  $0, $-1
        0x0006 00006 (C:\Users\FWR\Desktop\awesomeProject1\main.go:9)   SUBQ    $24, SP
        0x000a 00010 (C:\Users\FWR\Desktop\awesomeProject1\main.go:9)   MOVQ    BP, 16(SP)
        0x000f 00015 (C:\Users\FWR\Desktop\awesomeProject1\main.go:9)   LEAQ    16(SP), BP
        0x0014 00020 (C:\Users\FWR\Desktop\awesomeProject1\main.go:9)   FUNCDATA        $0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
        0x0014 00020 (C:\Users\FWR\Desktop\awesomeProject1\main.go:9)   FUNCDATA        $1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
        0x0014 00020 (C:\Users\FWR\Desktop\awesomeProject1\main.go:10)  LEAQ    type:main.User(SB), AX
        0x001b 00027 (C:\Users\FWR\Desktop\awesomeProject1\main.go:10)  PCDATA  $1, $0
        0x001b 00027 (C:\Users\FWR\Desktop\awesomeProject1\main.go:10)  NOP
        0x0020 00032 (C:\Users\FWR\Desktop\awesomeProject1\main.go:10)  CALL    runtime.newobject(SB)

```

目光集中到CALL指令，发现其执行了 `runtime.newobject(SB)`方法，也就是确实分配到了堆上，这是为什么？

**分析结果**

这是因为`GetUserInfo()`返回的是指针对象，引用被返回了方法之外了，因此编译器会把改对象分配到堆上，而不是栈上。否则方法结束之后，局部变量就被回收了，岂不是翻车。所以最终分配到堆上是理所当然的

**在想想**

难道所以的指针对象，都应该在堆上？

```
func main(){
	str := new(string)
	*str = "EDDYCJY"
}
```

```
$ go build -gcflags '-m -l' main.go
# command-line-arguments
./main.go:4:12: main new(string) does not escape
```

在go中，指针对象通常指向分配在堆上的数据。但是，并不是所有指针对象都必须在堆上分配。

在go中，在大部分情况下，如果一个对象被分配在堆上，那么它的地址会被存储在指针中。但是，go的编译器会对一些情况进行逃逸分析，如果编译器可以确定对象的生命周期，那么它就可以在栈上分配对象，而不需要使用指针来引用。这种情况会在函数调用结束时自动释放，无需手动管理内存。以下一些情况下，对象可能不会在堆上分配，而是在栈上分配：

1. 对于较小的对象，编译器可能会将其分配在栈上，而不是堆上。这样可以减少内存分配和垃圾回收的开销。
2. 对于局部变量或函数参数，如果编译器可以确定对象的生命周期仅限于函数调用内部，那么它可能会在栈上分配对象。
3. 在闭包中，如果编译器可以确定闭包不会逃逸到函数外部，那么它可能会在栈上分配的开销，因此并不是所有的指针对象都需要在堆上分配。

#### 项目swagger配置

```
cd ./server
go install github.com/swaggo/swag/cmd/swag@lastest
swag init --paseDependency -g ./entry/main.go -o ./docs
```

#### 三十五、git代码提交

> 在本地修改与远程代码无冲突的情况下，git先pull在commit，因为这样减少git没有必要的merge;
>
> 在本地修改与远程代码有冲突的情况下，git先commit在pull，这是为了应对多人合并开发的情况，避免覆盖源代码情况出现。

**一、git先pull再commit**

在本地修改与远程代码无冲突情况下，优先使用：pull ->commit ->push。在协商好的团队合作或者个人开发中，优先pull，因为这样会减少git没有必要的merge。

**二、git先commit再pull**

先commit再pull的情况就是为了应对多人合并开发的情况：

- commit是为了告诉git我这次提交改了哪些东西，不然你只是改了但是git不知道你改了，也就无从判断比较；
- pull是为了本地commit和远程commit的对比记录，git是按照文件的行数操作进行比较的，如果同时操作了某个文件的同一行那么就会产生冲突，git也会把这个冲突标记出来，这个时候就需要先把你的冲突和那个人拉过来问问保留谁的代码，然后再git add && git commit && git pull 这三连，再次pull一次是为了防止再你们协商的时候另一个人给又提交了一版东西，如果真发生了那流程重复一遍，通常没有冲突的时候就直接给你合并了，不会把你的代码覆盖掉；
- 出现代码覆盖或者丢失的情况：比如A B两人的代码版本都是1，A在本地提交了2、3并且推送到远程了，B进行修改了时候没有commit操作，他先自己写了东西，然后git pull这个时候B本地版本已经到3了，B在本地版本3的时候改了A写过的代码，再进行了git commit && git push 那么在远程版本中就是4，而且A的代码被覆盖了，所以所有人都要先commit再pull，不然真的会覆盖代码的。

#### 三十六、**Go语言的并发编程**

**并发介绍**

- 进程与线程

```
A.进程是程序操作系统中的一次执行过程，系统进行资源分配和调度的一个独立单位。
B.线程是进程的一个执行实体，是CPU调度和分配的基本单位，它是比进程更小的能独立运行的基本单位。
C.一个进程可以创建和撤销多个线程；同一个进程中的多线程之间可以并发执行。
```

- 并发和并行

```
A.多线程程序在一个核的CPU上运行，就是并发
B.多线程程序在多个核的CPU上运行，就是并行
```

![image-20240312135321481](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240312135321481.png)

- 协程和线程

```
协程：独立的栈空间，共享堆空间，调度又用户自己控制，本质有点类似于用户级线程，这些用户级线程的调度也是自己实现的
线程：一个线程上可以跑多个协程，协程是轻量级的线程
```

**goroutine只是由官方实现的超级“线程池”**

每个实例`4~5KB`的栈内存占用和由实现机制而大幅减少的创建和销毁开销是go高并发的根本原因。

**并发不是并行：**

并发只要是由切换时间片来实现“同时”运行，并行则是直接利用多核实现多线程的运行，go可以设置使用核数，以发挥多核计算机的能力。

goroutine奉行通过通信来共享内存，而不是共享内存来通信。



**GMP原理与调度**

**golang“调度器”的由来？**

**(1)单进程时代不需要调度器**

我们知道，一切的软件都是跑在操作系统上，真正用来干活（计算）的是CPU。早期的操作系统每个程序就是一个进程，知道一个程序运行完，才能进行下一个进程，就是“单进程时代”

一切的程序只能串行发生

![image-20240312141942590](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240312141942590.png)

早期的单进程操作系统，面临2个问题：

1. 单一的执行流程，计算机只能一个任务一个任务处理
2. 进程阻塞所带来的CPU时间浪费

后来操作系统就具有了最早的并发能力：多进程并发，当一个进程阻塞的时候，切换到另一个等待执行的进程，这样就能尽量把CPU利用起来，CPU就不浪费了。

**(2)多进程/线程时代有了调度器需求**

![image-20240312143636602](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240312143636602.png)

在多进程/线程的操作系统中，就解决了阻塞的问题，因为一个进程阻塞CPU可以立刻切换到其他进程中去执行，而且调度CPU的算法可以保证在运行的进程都可以被分配CPU的运行时间片。这样从宏观来看，似乎多个进程是在同时被运行。

但新的问题就又出现了，进程拥有太多的资源，进程的创建、切换、销毁，都会占用很长的时间，CUP虽然利用起来了，但如果进程过多，CPU有很大的一部分被用来进行进程调度了。

怎么才能提高CPU的利用率呢？

但是对于Linux操作系统来讲，cpu对进程的态度和线程的态度是一样的。

![image-20240312144528963](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240312144528963.png)

很明显，CPU调度切换的是进程和线程。尽管线程看起来很美好，但实际上多线程开发设计会变得更加复杂，要考虑很多同步竞争等问题，如锁、竞争冲突等。

**(3)协程来提高CPU的利用率**

多进程、多协程已经提高了系统的并发能力，到那时在当今互联网高并发场景下，为每个任务都创建一个线程是不现实的，因为会消耗大量的内存(进程虚拟内存会占用4GB[32位操作系统])，而线程也要大约4MB。

大量的进程/线程出现了新的问题

- 高内存占用
- 调度的高消耗CPU

然后工程师们就发现，其实一个线程分为“内核态”线程和“用户态”线程。

一个“用户态线程”必须绑定一个“内核态线程”，但是CPU并不知道有“用户态线程”的存在，它只知道它允许的是一个“内核态线程”（Linux的PCB进程控制块）

![image-20240312151251656](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240312151251656.png)

这样，我们再去细化去分类一下，，内核线程依然教“线程(thread)”,用户线程叫“协程（co-routine）”

![image-20240312151537824](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240312151537824.png)

看到这里，我们就要开脑洞了，既然一个协程(co-routine)可以绑定一个线程（thread），那么能不能多个线程(co-routine)绑定一个或者多个线程（thread）上呢

|| N : 1关系

N个线程绑定1个线程，优点就是协程在用户态线程即完成切换，不会陷入到内核态，这种切换非常的轻量快速。但也有很大的缺点，1个进程的所有协程都绑定在1个线程上

缺点：

- 某个程序用不了硬件的多核加速能力
- 一旦某协程阻塞，造成线程阻塞，本进程的其他协程就无法执行了，根本就没有并发能力了。

![image-20240312153049776](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240312153049776.png)

|| M ： N 关系

M个协程绑定在1个线程，是N：1和1：1类型的结合，克服了以上2种模型的缺点，但是先起来最为复杂。

![image-20240312153536100](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240312153536100.png)

协程跟协程是区别的，线程由CPU调度是抢占式的，协程由用户态调度式协作式的，一个协程让出CPU后，才执行下一个协程。

前面有了计算机进程、协程的基础知识后，接下来学习Golang调度器的实现和调度

**(4)Go语言的协程goroutine**

Go为了提供更容易使用的并发方法，使用了goroutine和channel。goroutine来自协协程的概念，让一组可复用的函数运行在一组线程上，即使有协程阻塞，该线程的其他协程也可以被runtime调度，转译到其他可运行的线程上。最关键的是，程序员看不到这些底层的细节，这就降低了编程的难度，提供了更容易的并发。

Go中，协程被成为goroutine，它非常轻量，一个goroutine只占几KB，并且这几KB就足够goroutine运行完，这就能在有限的内存空间内支持大量的goroutine，支持了更多的并发。虽然一个goroutine的栈只占几KB，但实际是可伸缩的，如果需要更多内容，runtime会自动为routine分配。

goroutine特点：

- 占用内存更小（几KB）
- 调度更灵活(runtime调度)

**(5)被废弃的goroutine调度器**

Go目前使用的调度器是2012年重新设计的，因为之前的调度器性能存在问题，所以使用了4年就被废弃了，那么我们先来分析一下被废弃的调度器是如何运作的？

> 大部分文章都是用G来表示Goroutine，用M来表示线程，那么我们也会用这种表达的对应关系

![image-20240312161554688](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240312161554688.png)

下main我们来看看被废弃的golang调度器是如何实现的？

![image-20240312162201682](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240312162201682.png)

M想要执行、放回G都必须访问全局G队列，并且M有多个，即多个线程访问同一资源需要加锁进行保证互斥/同步，所以全局G队列是有互斥锁进行保护的。

老调度器有几个缺点：

- 创建、销毁、调度G都需要每个M获取锁，这就形成了激烈的锁竞争。
- M转移G会造成低延迟和额外的系统负载。比如当G中包含创建新协程的时候，M创建了G‘，为了继续执行G，需要把G’交给M‘执行，也造成了很差的局部性，因为G’和G是相关的，最好放在M上执行，而不是其他M‘
- 系统调用（CPU在M之间的切换）导致频繁的线程阻塞和取消阻塞操作增加了系统开销

**（6）Goroutine调度器的GMP模型的设计思想**

面对之前调度器的问题，Go设计了新的调度器。

在新的调度器中，除了列M(thread)和G(goroutine)外,又引进了P(processor)

![image-20240312170804626](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240312170804626.png)

Processor，它包含了运行gorotine的资源，如果线程想运行goroutine，必需先获取P，P中还包含了可运行的G队列。

**（7）GMP模型**

在Go中，线程是运行goroutine的实体，调度器的功能是把可运行的goroutine分配到工作线程上。

- 全局队列(Global Queue):存放等待运行的G。
- P的本地队列：同全局队列类似 ，存放的也是等待运行的G，存的数量有限，**不超过256个**。新建G’时，G‘优先加入到P的本地队列，如果队列满了，则会把本地队列中的一半的G移动到全局队列。
- P列表：所有的P都在程序启动时创建，并保存在数组中，最多有GOMAXPROCS（可配置）个。
- M：线程想运行任务就得获取P，或从其他P的本地队列偷一半放到自己P的本地队列。M运行G，G执行之后，M会从P获取下一个G，不断的重复下去。

Goroutine调度器和OS调度器是通过M结合起来的，每个M都代表一个内核线程，OS调度器负责把内核线程分配到CPU的核上执行。

****

我们之前说过，Go语言是一种面向过程的语言。默认程序会很`main`函数开始，逐行执行代码，而且只能使用CPU的一个核。举个例子：

```
files := []string{"a.txt", "b.txt", "c.txt"}
for _, f := range files {
  Compress(f)
}
```

假设`Compress`会根据传入的路径来压缩文件。上面的循环会自动压缩`files`中保存的所有文件，但是一次只能压缩一个。后面的文件必须等前面的压完了才能开始。考虑到压缩操作比较耗时，而且当今的CPU都是多核，所以如果能使用多核并发压缩，一定会降低整个过程的耗时。在Go语言中可以使用协程实现这种效果。

```
files := []string{"a.txt", "b.txt", "c.txt"}
for _, f := range files {
  go Compress(f) // 启动协程
}
```

Go语言起协程非常简单，只需要在调用函数的前面加上`go`关键字。上面的代码会在循环中启动三个协程，每个协程独自运行`Compress`函数，互不影响。一般来说不同的协程会使用不同的CPU核，所以整个压缩过程可以并发执行。

- **并发的问题**

以上是一个虚拟的示例，没法直接运行。下面我用一个更简单的例子来说明使用协程需要注意的问题。

```
package main

func main() {
  for i := 0; i < 10; i++ {
    go func() {
      fmt.Println("hello")
    }()
  }
}
```

这里的`func(){}()`是一种函数的惯用法，意思是声明一个匿名函数并且立即执行。再加上前面的`go`关键字，表示声明匿名函数并且在新协程中立即执行。上面的例子是批量创建十个协程，并在每个协程里输出`hello`字符串。

如果你自己运行一把就会发现，程序什么内容也没输出，直接退出了😂造成这个问题的根本原因是`main`函数退出了，整个程序结束，之前创建的协程没有机会执行，也就不可能输出什么内容了。

要想解决这个问题，`main`函数在创建协程后，需要等待所有协程都结束后才能退出。不能自己先退出。这种协程之间的等待又称为**协程同步**。Go语言为协程同步提供了特殊的数据类型，叫通道(channel)。

- **通道(channel)**

通道跟切片有点类似，我们可以把它想象成是一消息队列，我们需要给消息指定类型，我们可以把消息放到队列，也可以从队列里获取消息。通道的声明语法如下：

```go
var ch1 chan int	//队列中没能放int数据
ch1 = make(chan int)
ch2 := make(chan string,4)
```

声明通道使用`chan`关键字，后面跟消息的类型。通道跟字典类似，声明之后还需初始化，不然无法使用。所有`ch1`还需要配合`make(chan int)`来完成初始化。

我们在使用`make`创建通道的时候还可以额外指定第二个参数，用来表示队列的缓冲区长度。如果不指定，默认长度就是零。缓冲区的用途后面再细说。

有了通道变量之后，我们就可以往里面投递数据或者读取数据：

```
ch1 <- 1	//写入数据
v := <- ch1 //读取
```

这里用到了箭头操作符`<-`,箭头表示数据的流向。`ch1 <- 1`表示数据流向通道，`<- ch1`表示数据流出通道。 `v := <- ch1`表示从`ch1`提出一条数据并保存到变量`v`中。

如果运行上面的两行代码写道`main`并运行，程序会直接报错：

```
fatal error:all goroutines are asleep - deadkock!
```

解决办法也很简单，把`ch1`改为`make(chan int,1)`就可以了，给它设一个缓冲区。

**对于没有缓冲区的通道，如果某协程想投递数据，则该协程会暂停执行。一直等到另一个协程想从该通道读取消息才能恢复。这种暂停也叫挂起。这种挂起是双向的。如果协程尝试`ch1`读取消息，但此时又没有其他协程尝试写入，则该协程也会被挂起。通道再这里是两个协程纽带，两个协程必须同时读写才行。**简单的说就是必须是由两个协程来完成读写数据，一个协程里面读写会报错，上面再一个`main`进行协程的读写就会出现错误。

回到上面的例子，我们先尝试给`ch1`写入消息，这个时候运行时就会挂起`main`函数所在的协程。因为后续的读取操作也在`main`函数中，所以不可能有协程从`ch1`读取内容了，这样`main`函数就会一直处理挂起状态。这就是所谓的死锁，自己跟自己死锁了😂

一般发生死锁后程序不会退出。因为运行时很难判断到底有没有产生死锁。但我们的示例中只有一个协程，也就是运行`main`函数的协程，它都挂起了说明一定产生了死锁，所以就直接报错了。但在实际生产系统中，死锁问题很难发现。所以大家在写并发代码的时候一定要小心。

如果把`ch1`改成`make(chan int,1)`,它就有了缓冲区。这个时候`ch1 <- 1`就不会挂起当前协程，然后面的`v := <- ch1`也就能正常执行了。

- **协程同步**

有了通道这个工具，我们就可以解决前面说的问题。改造代码如下:

```go
package main 

func main(){
    ch := make(chan int,10)
    for i := 0;i < 10;i++ {
        go func(){
            fmt.Println("hello")
            ch <- 1
        }()
    }
    for i:= 0;i < 10 ;i++ {
        <— ch
    }
}
```

我们定义了缓冲长度为10的`int`型通道。每个协程完成后会往`ch`写入消息`1`。因为有缓冲，所有不论有没有其他协程读取`ch`，刚的写入肯定会成功，对应的协程执行完成后会退出。

我们在`main`函数的最后尝试从`ch`读取消息。如果协程没有完成，读取操作会挂起`main`所在的协程。

运行改造后的代码就会看到程序输出十行`hello`。上例中也可以使用无缓冲通道。无缓冲通道需要读写双方同时操作，所以`main`函数读取一条消息才会有一个协程退出。而有缓冲的版本则没有此限制，所以打开协程完成工作后就会直接退出，不受`main`函数所在的协程影响。但无论如何，并发协程的执行顺序是不确定的。

- **并发顺序**

我们稍微修改一下上述代码：

```go
package main 

func main(){
    ch := make(chan int,10)
    for i:= 0;i < 10;I++ {
        go func(id int){
            fmt.Println("hello",id)
            ch <- 1
        }(i)
    }
    for i := 0;i < 10;i++ {
        <- ch
    }
}
```

我们给每一个协程加一个序号，通过匿名函数的参数传进去，然后在运行的时候打开出来。多执行几次，会发现每次的执行顺序都不一样，这就印证了刚才的判断。所以，协程在并发运行的时候是不可控的！

- **range语法**

在实践中经常会循环读取通道，所以Go语言也支持使用`range`关键字实现[遍历] 通道的效果。刚才`main`函数最后的for循环可以改为：

```go
for v := range ch {
	fmt.Println(v)
}
```

使用`range`的时候不需要指定箭头操作符！它实际的运行过程如下：

```go
for{
    v := <- ch
    fmt,Println(v)
}
```

所以这个for循环不会自动结束，除非有其他协程使用了`close`函数关闭`ch`通道。

最实战的例子是定时器。如果我们想每隔一段时间输出一条消息，可以使用`time`标准库提供的通道：

```go
ch := time.Tick(3 * time.Second)
go fun(){
    for t := range ch {
        fmt.Println(t)
    }
}()
```

`time.Tick(3 * time.Second)`会返回一个`chan time.time`通道，每隔三秒就可以从里面读取一个时间对象。上面的代码创建了一个单独的协程，不停从定时器通道读取时间。只要`main`不退出，该协程就会不停打印定时器触发的时间。

但问题来了，上面的循环什么时候结束呢？答案是永远不会结束。如果通道里面没有内容了，协程在读取的时候就会被挂起。但有时候我们希望告知对应的协程已经没有消息了，工作干完了，可以退出了。

通知退出有几种方法。最简单的就是关闭对应的通道：

```go
ch := make(chan int)
go func(){
    time.Sleep(1.*time.Second)
    close(ch)
}()
for v := range ch {
    fmt.Println(v)
}
```

通道关闭之后`for`循环会收到信号，然后结束循环。但这种方法有个副作用：

```
v := <- ch
```

如果是这就用箭头操作符读取消息，当通道关闭代码后还是会收到一个值，只不过这次是零值。本例中`v`最后的取值是`0`。未来区分正常的消息零值和关闭零值，Go语言还支持另一种读取语法：

```
v，ok := <- ch
```

通道关闭的时候上述读取也会返回，而且`ok`会被设置为`false`。程序可以通过检查`ok`变量来确定是否已经关闭。

- **select语法**

除了关闭通道之外，我们还可以使用一个单独的控制通道来通知协程退出。比如：

```go
ch1 := make(chan int)
ch2 := make(chan int)

go func(){
    for{
        v := <- ch1
        fmt.Println(v)
        if v := ch2;v > 0 {
            return
        }
    }  
}()
```

这段程序在协程里先从`ch1`读取数据，然后打印出来。它在继续循环之前会尝试从`ch2`读取消息。如果有消息则表明活干完了，可以退出。这段程序有一个问题，如果`ch1`里面迟迟没有消息，那么协程就会卡在`v:= <- ch1`这一行。此时外界通过`ch2`发信号没法关闭协程，因为协程被挂起了。即使有协程给`ch1`发消息，程序会输出对应的值。但如果外界没有给`ch2`发消息，那么`v := <- ch2`也会卡住。所以说上述代码根本无法实现我们的设计意图。

这个问题的核心是在同一个协程中从通道读取消息必须是按顺序执行。而我们希望同时从多个通道等到消息。无论是`ch1`还是ch2，只要有消息就会唤醒协程处理。这个功能需要`select`关键字：

```go
go func(){
    for{
        select{
            case v := <- ch1:
            fmt.Println(v)
            case <- ch2:
            	return
        }
    }
}()
```

因为用上了`select`，只要`ch1`和`ch2`任何通道有消息，处理协程就会被唤醒。

`select`还支持`default`分支，对应的是所有`case`分支都没有新消息的情景，业务代码使用较少，初学者记住有这么回事就行。后面边用边学。

除了用通道来同步协程外，Go官方话封装了`WaitGroup`对象，之前的实例可以改写为：

```go
import "sync"

func main(){
	var wg sync.WaitGroup
    for i := 0;i < 10 ;i++ {
        wg.Add(1)
        go func(){
            defer wg,Done()
            fmt.Println("hello")
        }()
    }
    wg.Wait()
}
```

声明`wg`之后，每创建一个协程就调用一个`wg.Add(1)`，表示要多等待一个协程。`main`函数最终调用`wg.Wait()`等待所有协程结束。每一个协程结束后需要调用`wg.Done()`。所有协程都结束后`wg.Wait()`函数就会返回，整个程序才能退出。

新代码中使用`defer wg.Done()`来调用函数，它能确保每个协程结束时一定能执行`wg.Done()`函数。如果不用`defer`而且协程在执行的过程产生了`panic`，那就可能没法执行对应的`wg.Done()`函数，所以`main`函数会一直等待，从而产生死锁。

- **协程争用**

所谓争用，就是多个协程争相读写同一个变量。前文所讲的通道就是最典型的例子，不同的协程可能会同时读写通道。因为有争用，所以也就有了并发安全的概念。如果类型的变量支持多协程同时读写，我们就称之为并发安全。通道类型就是并发安全的。但是除非特殊说明，几乎所有的类型都不是并发安全的。如果多协程同时读写某变量，轻则会产生panic进而程序退出(是的，这是简单情况)；重则会破环数据，但程序以一种错误的方式持续运行，等发现的时候已经很难收场 。下面是一个案例：

```go
func main(){
    a := 0
    g := sync.WaitGtoup{}
    for i := 0;i < 100;i++ {
        g.Add(1)
        go func(){
            defer g.Done()
            a++
        }()
    }
    g.Wait()
    fmt.Println(a)
}
```

程序声明一个变量`a`，然后启动一百个协程对该变量做加一操作。运行几次就会发现，结果并非总是一万。

出现这种现象的根本原因是所有协程同时操作一个变量`a`。因为协程的执行顺序和时间并不固定，当某个协程把`a`从`0`改为`1`的时候，有可能另外一个协程没有拿到`a`是最新值`1`，依然是在`0`的基础上加`1`再写会`a`对应的内存，这样就可能产生错误的结果。

- **并发锁**

最简单的办法就是加锁。协程在更新`a`之前先获取一把锁，这个时候其他协程因为不可能同时获取锁，所以只能等当前协程更新之后再更新。当前协程更新完成后需要主打释放锁，这样其他协程才能继续尝试获取锁并更新`a`的值。

改良后的代码如下：

```go
func main(){
    a := 0
    m := sync.Mutex{}
    g := sync.WaitGroup{}
    for i := 0; i < 10000;i++ {
        g.Add(1)
        go func(){
            defer g.Done()
            m.Lock()
            defer m.UnLock()
            a++
        }()
    }
    g.Wait()
    fmt.Pritln(a)
}
```

先声明`m := sync.Mutex{}`。然后各协程在更行之前尝试获取锁`m.Luck()`。一次只能有一个协程锁定成功，其他协程都会被挂起。等成功的协程更新完成后会通过`defer m.UnLock()`释放锁定。这时候被挂起的协程又会别唤醒，开始新一轮的争抢过程。

加锁的本质是排队，所有的协程按照获取锁的顺序依次更新`a`的值，这样就不会产生并发总题。但是，加锁是有代价的。没有抢到锁的协程会被挂起，而且协程多了相互争锁也会给操作系统带来一定的负担。所以说还是少写可能产生争用的代码。

减少急用最简单的办法是不要共享内存(变量)。不共享就不会有争用，大家各干各的，互不影响。所以我们要尽量避免使用共享变量、指针、切片和字节，这些变量都是引用传递，不同协程可能会操作同一段内存，从而产生征用问题。Go语言本身也是鼓励传值，鼓励拷贝内存，虽然会有一些性能上的损耗，但过跟并发引起的BUG相比，这种损耗不值一提。

- **并发安全类型**

再一个办法就是使用并发安全的数据类型。比如字典可以使用`sync.Map`代替内置的`map`。但是无脑使用并发版本类型也会导致性能问题，因为多数情况下并不会碰到多协程争用问题。

```go
m := sync.Map{}
for i := 0;i < 10;i++ {
    go func(){
        //无需加锁
        m.Store(i,fmt.Sprint(i))//i => "1"
    }()
}
v := m.Load(1)
v.(string)
```

因为不是内置类型，所以`sync.Map`不支持`range`关键字，只能通过`Range()`来遍历。

- **原子类型**

有一类并发安全的类型叫原子类型，它们由底层硬件实现并发安全，**几乎没有性能损耗**。可以优先考虑使用。相关的类型都封装在`sync/atomic`这个包。

前面的例子可以改写为：

```go
import "sync/atomic"
func main(){
    a := &atomic.Int32{}
    g := sync.WaitGroup{}
    for i := 0;i < 1000;i++ {
        g.Add(1)
        go func(){
            defer g.Done()
            a.Add(1)
        }()
    }
    g.Wait()
    fmt.Println(a.Load())
}
```

- **读写锁**

为了降低加锁成本，人们发明了读写锁。简单来说多数变量都是读多写少。多个协程同时读一个变量不会有问题，只要读的过程中没有人修改这个变量就行了。

首先创建读写锁：

```
rwm := sync.RWMutex{}
```

读协程需要获取读锁，写协程需要获取写锁：

```
rwm.RLock()	//获取读锁
rwm.Lock()	//获取写锁
```

因为读不会修改数据，所以允许多个协程获取读锁，也就是并发读取变量内容。此时如果有少量协程想修改内容，它们需要获取写锁。写锁是排他锁，需要等所有读锁和其他锁释放才能获取。一旦写协程成功获取写锁，在它解锁之前，所有其他尝试获取读锁或写锁的协程都会被挂起。**读锁也叫共享锁，写锁也叫排他锁。**

- **goroutine调度**

GPM是Go语言运行时(runtime)层面的实现，是go语言自己实现的一套调度系统。区别于操作系统调度OS线程。

1. G很好理解，就是个goroutine的，里面除了存放本goroutine信息外，还有与所在P的绑定等信息。
2. P管理着一组goroutine队列，P里面会存储当前goroutine运行的上下文环境（函数指针，堆栈地址及地址边界），P会对自己管理的gotourine队列做一些调度（比如把占用CPU时间较长的goroutine暂停、运行后续的goroutine等等）当自己的队列消费完了就回去全局队列里取，如果全局队列也消费完了回去其他P的队列里抢任务。
3. M（machine）是Go运行时（runtime）对操作系统内核线程的虚拟，M与内核线程一般是一一映射的关系，一个goroutine最终是要放到M上执行的；

P与M一般也是一一对应关系。他们关系是：P管理者着一组G挂载在M上运行。当一个G长久阻塞在一个M上时，runtime会新建一个M上。当旧的G阻塞完成或者认为其已经死掉时，回收旧的M。

P的个数是通过runtime.GOMAXPROCS设定（最大256），Go1.5版本之后默认为物理线程数。在并发量大的时候会增加一些P和M，但不会太多，切换太频繁的话的得不偿失。

当从线程角度讲，Go语言相比起其他语言的哟是在于OS线程是由OS内核来调度的，goroutine则是由Go运行时（runtime）自己的调度器调度的，这个调度器使用过一个成为m：n调度的技术（复用/调度m个goroutine到n个OS线程。）其一大特点是goroutine的调度是在用户态下完成的，不涉及内核与用户态之间的频繁切换，包括内存的分配与释放，都是用户维护者一块大的内存池，不直接调用系统的malloc函数（除非内存池需要该百年），成本比调度OS线程低很多。另一方面充分利用了多核的硬件资源，近似的把若干goroutine均分在物理线程上，在加上本身goroutine的超轻量，以上种种保证了go调度方面的性能。

- **channel**

单纯地将函数并发执行是没有意义的。函数与函数之间需要交换数据才能体现并执行函数的意义。

虽然可以使用共享内存进行数据交换，但是共享内存在不同的goroutine中容易发生竞态问题。未来保证数据交换的正确性，必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。

Go语言的并发模型是CSP（Communicating Sequential Processess），提供通信共享内存二不是通过共享内存而实现通信。

通过说goroutine是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。

Go语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

- **总结**

到这讲完了并发相关的主要内容。不论是哪一种语言，并发编程都是非常困难的领域。Go语言内置协程只是降低了并发编程的门槛，但绝对没有降低并发编程的难度。

#### 三十七、Golang进阶

**参数传递**

****

go的参数传递分两种：

- 值传递: 值传递实际就是一份拷贝，函数内部对该值的修改，不会影响外部的值

```go
func main() {
    x := 16
    fmt.Println("修改前的值%d:",x)	//16
    change(x)
    fmt.Println("修改后的值%d",x)	//16，没有改变
}

func change(x int) {
     x = 100
    fmt.Println("修改时%d:",x) //修改拷贝的值变成100
}
```

- 引用传递: **引用传递本质也是值传递，只不过这份值是一个指针（地址）。**所以我们在函数内对这份值的修改，其实不是改这个值，而是去修改这个值所指向的数据，所以是会影响到函数外部的值。

```go
func main() {
    x := 10
    fmt.Println("修改前的值:%d",x)	//16
    change(&x)
    fmt.Println("修改后的值:%d",x)	//100，被修改了
}

func change(x *int) {
    *x = 100
    fmt.Println("修改时的值:%d",x)	//100
}
```

传指针使得多个函数能操作同一个对象

传指针比较轻量级（8bytes），只是传内存地址，我们可以用指针传递体积大的结构体。如果用参数传递的话，在每次copy上面会花费相对多的系统开销(内存和时间)。所以当你要传递一个大的结构体的时候，用指针时一个明智的选择。

Go语言中**slice,map和channel**这三种类型的实现机制类似指针，所以可以直接传递，而不是取地址后传递指针。

**返回值**

****

返回值可以直接起别名返回，避免初始化多个返回值

```go
func hello(x,y string) (res string,err error){
    return
}
```

**匿名函数**

****

Go语言支持函数式编程:

- 将匿名函数作为另一个函数的参数：回调函数
- 将匿名函数作为另一个函数的返回值：闭包

```go
func main(){
    //匿名函数
    fun(){
        fmt.Println("这是一个匿名函数")
    }()
    
    //用一个变量来接收一个匿名函数，就可以在他的作用域内多次调用该函数
    fun1 :- func(){
        fmt.Printlin("这也是一个匿名函数")
    }
    fun1()
    fun1()
    
    //定义一个带参数的匿名函数
    fun(a,b int) {
        fmt.Println(a,b)
    }(1,2)
    
    //定义一个带返回值的匿名函数
    res1 := func(a,b int) int {
        return a + b
    }(10,20)
    fmt.Println(res1)
}
```

**回调函数**

回调函数：callback，就是将一个函数func2作为函数fun1的一个参数，那么func2叫做回调函数，fun1叫做高阶函数。

```go
type CallBack func(x,y int) int

//根据传入的回调函数进行算数运算
func oper(a,b int,callbackFunc CallBack) int {
    res := callbackFunc (a,b)
    return res
}

//加法运算回调函数
func add(a,b int) int {
    return a + b
}

//减法运算回调函数
funa sub(a,b int) int {
    return a - b
}

func main() {
    //将add作为回调函数传入oper
    res1 := oper(10,20,add)
    fmt.Println(res1)
    
    //将sub作为回调函数传入oper
    res2 := oper(5,2,sub)
    fmt.Println(res2)
}
```

许多官方库就是使用了回调函数的思想，把灵活的处理逻辑交给用户自身来实现，这样代码的可定制化大大增强；比如上面的排序函数，由于每个业务场景都有不同的排序诉求，所以把具体的排序实现交给用户，让用户去实现回调函数逻辑。

**闭包**

闭包：一个外层函数中，有内层函数，该内层函数中，会有操作外层函数的局部变量，并且该外层函数的返回值就是这个内层函数。那么这个内层函数和外层函数的局部变量，统称为闭包结构

闭包 = 函数 + 引用环境

```go
func increment() func() int {
    //1.定义一个局部变量
    i := 0 
    //2.定义一个匿名函数，给变量自增返回
    fun := func() int {
        i++
        return i
    }
   	//3.返回该匿名函数
    return fun
}

func main() {
    res1 := increment()  //res1 = fun
    fmt.Printf("%T\n",res1) //func() int
    
    //带括号表示自执行函数res1，得到返回结果v1
    v1 := = res1()
    fmt.Println(v1) //1
    //再次执行res1，得到返回结果v2
    v2 := res1()
    fmt.Println(v2) //2
    fmt.Println(res1()) //3    
    fmt.Println(res1()) //4
    fmt.Println(res1()) //5
    fmt.Println(res1()) //6
	
    //用一个新的变量来接收increment()的返回结果
    //这个时候increment函数又重新执行了一遍
    res2 := increment()
    fmt.Printf("%T\n"res2)
    //执行res2
    v3 := res2()
    fmt.Println(v3) //1
    fmt.Println(res2()) //2
    
    //res1 和 res2 没什么关系
    fmt.Println(res1()) //7
}
```

被捕获到闭包中的变量让闭包本身拥有了记忆效应，闭包中的逻辑可以修改闭包捕获的变量，变量会跟随闭包生命其一直存在。

隔离数据：假设你想创建一个函数，该函数可以访问即使退出后仍然存在的数据。举个例子，如果你想统计函数被调用的次数，单不希望任何人访问该数据(这样他们就不会意外更改它)，你就可以用闭包来实现它

```go
//函数计数器
func counter(f func()) func() int {
    n := 0
    return func() int {
        f()
        n += 1
        return n
    }
}

//测试的调用函数
func foo(){
    fmt.Println("call foo")
}

func main() {
    cnt := counter(foo)
    cnt()
    cnt()
    cnt()
    cnt()
    fmt.Println(cnt())
}
```

还可以通过闭包的记忆效应来实现设计模式中工厂模式的生成器

```go
//定义一个bytedancer生成器，输入名称，返回新的用户数据
func getBytedancer(name string) func() (string,int){
    //定义字节范分数
    style := 100
    //返回闭包
    return func()(string,int){
        //引用外部的style变量，形成了闭包
        return name,style
    }
}

func main() {
    //创建一个bytedancer生成器
    genertor := getBytedancer("bytedancer001")
    
    //返回新创建bytedancer的姓名，字节范分数
    name,style := genertor ()
    fmt.Println(name,style)
}
```

闭包具有面向对象语言的特性——封装性，变量style无法从外部直接访问和修改。

**并发编程基础讲解**

**背景知识**

- 协程：独立的栈空间，共享堆空间，调度由用户自己控制，本质上有点类似于用户级线程，这些用户线程的调度也是自己实现的。
- 线程：一个线程上可以跑多个协程，协程也是轻量级的线程。

Goroutine是官方实现的超级“线程池”，每个实例占用4~5kb的栈空间且极少的创建销毁开销是go高并发的根本原因。

并发是通过切换时间片来实现“并行”运行，go可以设置使用核心数`runtime.GOMAXPROCS`,发挥多核主机能力。

**初探并发**

****

Go语言中使用共routine非常简单，只需要调用函数的时候在前面加上go关键字，就可以为函数创建一个goroutine。

```go
func hello() {
    fmt.Println("Hello Goroutine")
}

func main(){
    hello()
    go hello()
    time.Sleep(3 * time.Second)
    fmt.Println("main goroutine done")
}
```

在程序启动时，Go程序就会为main()函数创建一个默认的goroutine。当main()函数返回的时候该goroutine就结束了，所有在main()函数中启动的goroutine会一同结束，所以在实际使用goroutine时需要特别注意其调度。

**多个goroutine**

```go
func hello(i int) {
    fmt.Println("hello goroutine!",i)
}

func main(){
    for i := 0; i < 10 ;i++ {
        go hello(i)
    }
    time.Sleep(3 * time.Second)
}
```

多次执行上面的代码，会发现每次打印的数字的顺序都不一致。这是因为10个goroutine是并发执行的，而goroutine的调度是随机的。

**当goroutine遇上loop**

```go
func hello(i int) {
    fmt.Println("hello goroutine!",i)
}

func main() {
    for i := 0;i < 10; i++ {
        go func(){
            hello(i)
        }()
    }
    time.Sleep(time.Second)
}
```

由于**闭包**只是绑定到i变量上，并没有保存到goroutine栈中，这样写会导致for循环结束后才执行goroutine多线程操作，这时候value值指向了最后一个元素，所以上面的代码极大可能都是输出最后一个元素。

1. 通过参数传递数据到协程

   ```go
   func hello(i int) {
       fmt.Println("hello goroutine!",i)
   }
   
   func main() {
       for i := 0;i < 10;i++ {
           go func(idx int) {
               hello(idx)
           }(i)
       }
       time.Sleep(time.Second)
   }
   ```

   这里将idx作为一个参数传入goroutine中，每个idx都会独立计算并保存到goroutine的栈中

​	2.定义临时变量

```go
func hello(i int){
    fmt.Println("hello Goroutine!",i)
}

func main(){
    for i := 0;i < 10;i++ {
        va1 := i
        go func(){
            hello(va1)
        }()
    }
    time.Sleep(time.Second)
}
```

另一种方法是在循环内定义新的变量，由于在循环内定义的变量在循环遍历的过程中是不共享的，因此也可以达到同样的效果

**并发同步**

****

在代码中生硬的使用time.Sleep肯定是不合适的，我们推荐使用sync.WaitGroup来实现并发任务的同步。

sync.WaitGroup有以下几个方法：

| 方法名                     | 功能                |
| -------------------------- | ------------------- |
| (wg *WaitGroup) Add(n int) | 计数器+n            |
| (wg *WaitGroup) Done()     | 计数器-1            |
| (wg *WaitGroup) Wait()     | 阻塞直到计数器变为0 |

sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。例如当我们启动了N个并发任务时，就将计数器值增加N。每个任务完成时通过调用Done()方法将计数器减1.通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成。

我们利用sync.WaitGroup将上面的代码优化一下：

```go
var wg sync.WaitGroup
func hello(i int) {
    defer wg.Done()	//goroutine结束就登记-1
    fmt.Println("Hello Goroutine!",i)
}

func main() {
    for i := 0;i < 10;i++ {
        wg.Add(1)	//启动一个goroutine就登记+1
        go hello(i)
    }
    fmt.Println("main goroutine done!")
    wg.Wait()	//等待所有登记的goroutine都结束
}
```

**并发安全**

****

有时候会存在多个goroutine同时操作一个资源（临界区），这种情况就会发生数据竞态问题。类比卫生间被整层同性别的同学竞争使用场景。

```go
var x int64
var wg sync.WaitGroup

func add() {
    defer wg.Done()
    for i := 0;i < 10000; i++ {
        x = x + 1
    }
}

func main() {
    wg.Add(2)
    go add()
    go add()
    wg.Wait()
    fmt.Println(x)
}
```

上面的代码中我们开启了两个goroutine去累加变量x的值，这两个goroutine再访问和修改x变量的时候就会存在数据竞争，导致最后的结果与期待的不符。

**互斥锁**

****

互斥锁是一种常用的控制共享资源访问的办法，它能够保证同时只有一个goroutine可以共享资源。Go语言中使用sync.Mutex类型来实现互斥锁。使用互斥锁来修复上面代码的问题：

```go
var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
    defer wg,Done()
    for i := 0;i < 10000;i++ {
        lock.Lock() 	//加锁
        x = x +1
        lock.Unlock()	//解锁
    }
}

func main() {
    wg.Add(2)
    go add()
    go add()
    wg.Wait()
    fmt.Println(x)
}
```

**读写互斥锁**

****

互斥锁是完全互斥的，实际上我们更多的场景是读多写少的，当我们并发的读取不涉及修改的资源是没必要加锁的，这时我们使用读写锁sync.RWMutex时一种更好的选择。

读写锁分为读锁和写锁。

当一个goroutine获取读锁之后，其他的goroutine如果获取读锁就会继续获得锁，如果获取写锁就会等待；

当一个goroutine获取写锁之后，其他协程无论获取读锁还是写锁均会等待。

互斥锁

```go
var (
	x int64
    wg sync.WaitGroup
    lock sync.Mutex
)

func write(){
    lock.Lock()	//加上互斥锁
    x = x + 1
    time.Sleep(10 * time.Millisecond)//假设写操作耗时10毫秒
    lock Unlock()	//解互斥锁
    wg.Done()
}

func read(){
    lock.Lock()	//加互斥锁
    time.Sleep(time.Millisecond)//假设读操作耗时1毫秒
    lock.Unlock()	//解互斥锁
    wg.Done()
}

func main(){
    start := time.Now()
    for i := 0;i < 10;i++ {
        wg.Add(1)
        go write()
    }
    
    for i := 0;i < 1000;i++ {
        wg.Add(1)
        go read()
    }
    
    wg.Wait()
    end := time.Now()
    fmt.Println(end.Sub(start))
}
```

> 互斥锁的输出值是1.1s

读写锁

```go
var (
	x int64
    wg sync.WaitGroup
    rwlock sync.RWMutex
)

func write(){
    rwlock.Lock()	//加写锁
    x = x +1
    time.Sleep(10 * time.Millisecond)	//假设写操作耗时10毫秒
    rwlock.Unlock()	//解写锁
    wg.Dome()
}

func read() {
    rwlock.RLock()	//加读锁
    time.Sleep(time.Millisecond)	//假设读操作耗时1毫秒
    rwlock.RUnlock()	//解读锁
    wg.Done()
}

func main(){
    start := time.Now()
    for i := 0;i < 10;i++ {
        wg.Add(1)
        go write()
    }
    
    for i := 0;i < 1000;i++ {
        wg.Add(1)
        go read()
    }
    
    wg.Wait()
    end := time.Now()
    fmt.Println(end.Sub(start))
}
```

> 读写锁的输出值是101ms，约等于0.1s，性能提升了10倍

我们可以看到读写锁非常适合读多写少的场景，如果读和写操作差别不大，读写锁的优势就发挥不出来。

**并发通信**

****

单纯的将函数并发执行是没有使用场景的，函数与函数之间需要交换数据才能真正体现并发执行的意义。

虽然可以使用共享内存进行数据交互，但是共享内存在不同的goroutine中容易发生竞态问题。为了保证数据交换的正确性，需要使用互斥量对内存进行加锁，而不是共享内存来通信。

channel是go并发的通信桥梁，可以让一个goroutine发送特定值到另一个goroutine进行通信。channel遵循先进先出(FIFO)机制，保证收发数据的顺序，每个channel都是一个具体类型的通道，声明时需要指定其元素类型。

**channel操作**

通道有发送(send)、接收(receive)和关闭(close)三种操作

发送和接收都使用 <- 操作符，关闭使用内置的close函数

```go
//初始化一个channel
ch := make(chan int)

//发送操作
ch <- 10 //把10发送到ch中

//接收操作
x := <- ch //从ch中接收值并复制给变量x

//关闭操作
close(ch)
```

对于channel关闭操作，需要注意的是，只有在通知接收方goroutine把所有的数据都发送完毕的时候才需要关闭通道。且channel是可以被GC机制回收掉的，所以关闭通道不是必须操作的。

**channel常见操作整理**

| 操作/状态 | nil   | 空               | 非空                               | 未满                             | 满                               | 已关闭                 |
| --------- | ----- | ---------------- | ---------------------------------- | -------------------------------- | -------------------------------- | ---------------------- |
| 发送      | 阻塞  | 发送值           | 发送值                             | 发送值                           | 阻塞                             | panic                  |
| 接收      | 阻塞  | 阻塞             | 接受值                             | 接收值                           | 接收值                           | 读取完毕数据后返回零值 |
| 关闭      | panic | 关闭成功返回零值 | 关闭成功后，读取完毕数据后返回零值 | 关闭成功，读取完毕数据后返回零值 | 关闭成功，读取完毕数据后返回零值 | panic                  |

**无缓冲channel**

无缓冲的通道又称为阻塞的通道

```go
func mian() {
    ch := make(chan string)
    ch <- "chan test"
    fmt.Println("发送成功")
}
```

上面这段代码能够编译通过，但是执行的时候会出现deadlock(死锁)错误；为啥出现这种错误呢？

我们发现这个通道没有接收值，我们定义一个变量去接收通道的值，这样就解决没缓冲需要接受值，用代码去实验一下：

```go
func main() {
	ch := make(chan string)
	ch <- "chan test"
	v := <-ch
	fmt.Println("发送成功", v)
}
```

发现还是会出现执行时异常：

```
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	/tmp/sandbox19711939/prog.go:11 +0x36

Program exited.
```

这说明在一个协程里面进行同一个通道的发送和接收操作，会出现死锁，因为我们使用ch :=make(chan int)创建了无缓冲的通道，无缓冲的通道只有在又接受值的时候才能发送值，简单的来说就是无缓冲的通道必须有接收才能发送。

左边的代码会阻塞在ch <- 10这一行代码形成死锁，如何解决这种问题？

一种方法是启用一个goroutine去接收值

```go
func recv(c chan int) {
    ret := <- c
    fmt.Println("接收成功",ret)
}

func main() {
    ch := make(chan int)
    go recv(ch)	//这行代码要在发送前面，不然也会出现执行时错误，产生死锁
    ch <- 10
    fmt.Println("发送成功")
}
```

无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成功，两个goroutine将继续执行。相反，如果接收操作先执行，接收放的goroutine将阻塞，直达另一个goroutine在该通道上发送一个值。

使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为同步通道。

**有缓冲channel**

解决上面问题的方法还有一种就是使用有缓冲区的通道。

我们可以在使用make函数初始化通道的时候为其指定通道的容量

```go
func main(){
    ch := make(chan int,1)	//创建一个容量为1的有缓冲区通道
    ch <- 10
    fmt.Println("发送成功")
}
```

只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量。

**Context**

****

我们能发现调用大部分外部仓库方法时，第一个参数都是`ctx context.Context`,包括公共库的大部分接口设计也是遵循该规范。

标准要求：每个方法的第一个参数都将context作为第一个参数，并使用ctx变量名惯用语。

****

**Context接口**

```go
type Context interface{
    Deadline() (deadline time.Time,ok bool)
    Done() <- chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

`context`接口包含四个方法:

- `Deadline()`:返回绑定当前`context`的任务被取消的截止时间，如果没有设定期限，将返回ok ==false
- `Done()`:当当前的`context`被取消时，将返回一个关闭的`channel`，如果当前`context`不会被取消，将返回nil
- `Err()`:
  - 如果`Done()`返回的`channel`没有关闭，将返回`nil`；
  - 如果`Done()`返回的channel已经关闭，将返回非空的值表示任务结束的原因；
  - 如果`context`被取消，`Err()`将返回`canceled`；
  - 如果是`context`超时，Err()将返回`DeadlineExceeded`
- `Value()`:返回`context`存储的键值对中当前`key`对于的值，如果没有对于的`key`，则返回`nil`

可以看到Done()方法返回的channel正是用来传递结束尾信号以抢占并中断当前任务；Deadline()方法表示一段时间后当前goroutine是否会被取消；以及一个Err()方法，来解释goroutine被取消的原因；而Value()则用于获取特定于当前任务树的额外信息。

**依赖管理**

****

为什么需要依赖管理？

- 最早的时候，Go所依赖的所有的第三方库都放在GOPATH这个目录下面。这就导致了同一个库只能保存一个版本的代码。如果不同的项目依赖同一个第三放库的不同版本，就会出现冲突。

go midule 是Go1.11版本之后官方退出的版本版本管理工具，并且从1.13本版本开始，go mudule将是Go语言默认的依赖管理工具。

**Go Module常用命令**

```
go mod download 下载依赖的module到本地cache（默认为$GOPATH/pkg.mod目录）
go mod edit	编辑go.mod文件
go mod graph 打印模块依赖图
go mod init 初始化当前文件夹，创建go.mod文件
go mod tidy 增加缺少的module，删除无用的module
go mod vendor 将依赖复制到vendor下
go mod verify 校验依赖
go mod why 解释为什么需要依赖
```



#### 三十八、工具通用方法记录

1.根据两个日期时间，获取相差的月数

```go
func MonthBetween(t1,t2 time.Time) (int,error) {
    if !t1.Before(t2) {
        return 0,errors.New("开始时间必须在结束时间之前")
    }
    yearDiff := t2.Year() - t1.Year()
    monthDiff := int(t2.Month()) - int(t1.Month())
    
    totalMonths := yearDiff * 12 + monthDiff
    if t2.Day() < t1.Day() {
        totalMonths--
    }
    return totalMonths,nil
}
```

#### 三十九、Go泛型

**1.泛型介绍**

泛型编程是计算机科学中一个相当重要的概念，广泛应用于各种编程语言和框架中。泛型允许在强类型编程语言设计代码时，**能够在实例化时指定类型**作为参来指明使用哪些数据类型，有助于提高代码复用性，增加类型的安全，在某种情况下还能达到性能优化的效果。

Go语言的泛型在`Go 1.18`版本中引入的一个新特性，它允许开发者编写可以处理不同数据类型的代码，而无需为每种类型都编写重复的代码。泛型的引入使得Go语言在编写一些通用功能时更加方便，在开发时能够通过止指定数据类型来处理不同数据类型的代码，而无需为每种数据类型都编写重复的代码，从而提高了代码的复用性和可读性。

举个例子，如下是一个实现反转切片的函数如下：

```go
func Reverse(s []int) []int {
    n := len(s)
    for i,j := 0,n-1;i < j;i,j = i+1,j -1 {
        s[i],s[j] = s[j],s[i]
    }
    return s
}

func main() {
    intSlice := []int{1,2,3,4,5}
    fmt.Println("Revesed int slice:",Reverse(intSlice)) //Revesed int slice: [5 4 3 2 1]
}
```

上述`Reverse`函数接收一个`[]int`类型的参数，但如果想要将一个`[]float64`类型的切片实现反转则无法调用`Reverse`函数，可能导致需要重新定义一个接收参数类型`[]float64`的参数，这样就导致相同逻辑的函数却因为类型不同而重复定义。

```go
func Reverse(s []float) []float {
    n := len(s)
    for i,j := 0,n-1;i < j;i,j = i+1,j -1 {
        s[i],s[j] = s[j],s[i]
    }
    return s
}
```

**2.泛型的意义**

**类型安全：弱类型的利弊**

在没有提供泛型语法时，Go语言中的由于interface{}的特性，经常被用作通用类型，通过使用interface{}可以接收任何类型的参数，但带来的坏处则是失去了编译器类型检查的好处。

```go
func Add(x,y interface{}) interface{} {
    return x.(int) + y.(int)	//需要使用类型断言，且不安全
}
```

上述Add()函数使用类型断言将x于y转换成int类型。这种类型断言的不完全主要体系那在两个方面：

- **类型断言失败**

​		如果传入的`x`与`y`并非`int`类型，则在进行类型断言时会发生运行时错误，导致程序崩溃，产生一个可控的错误响应，例如传入`float64`类型，则会报`panic: interface conversion: interface {} is float64,not int`错误

为了处理这种情况，需要在`x.(int)`之前，先检查x是否实现了`int`类型的方法。如果类型不正确，可以返回一个错误或者提示信息。

- **类型转换错误**

​		即使类型断言成功，但如果传入的参数值在转换成目标类型后发生溢出或下溢，也可能导致程序崩溃。例如当两个int类型的数相加导致整数溢出时。

为了避免这种情况，需要在进行类型转换之前检查参数值是否在目标类型的范围内。

**强类型的优势**

通过使用泛型，可以在编辑器进行类型检查，从而解决上面类型断言失败和类型转换错误等问题。

```go
type Addable interface {
    int | float64
}

func Add[T Addable](x,y T) {
    return x + y
}
```

上述代码定义了一个名为`Addable`的接口,`Addable`接口使用了类型断言 `（int | float64）`，指定了实现该接口的类型必须是`int`或`float64`。`Addable`实际上是一个类型约束，只允许那些满足条件的类型（比如，可以进行加法操作的类型）作为泛型参数。

在`Add`函数中，通用使用泛型，让其接收两个类型为`Addable`的参数`x`和`y`，并返回它的和。

通过使用 泛型：

- 可以确保只有满足接口要求的类型（即`int`和`float64`）才能被用于Add函数。这有助于在编译时检查类型错误，提高代码的安全性。
- 泛型函数`Add`可以处理满足`Addable`接口要求的任何类型，无需为每种类型编写重复的代码，使得代码更加简洁，易于维护。
- 若后续需要支持其他类型进行加法操作的类型，只需确保这些类型实现了`Abbable`接口，即可在`Add`函数中使用。这使得代码具有很好的可拓展性。

```go
func main() {
    add1 := Add[int](1,3)
    add2 := Add[float64](3.14,4.75)
    fmt.Println(add1)	//3
    fmt.Println(add2)	//7.890000000000001
}
```

**代码复用**

在没有泛型的情况下，如果说想为不同的类型实现相同的逻辑，通常需要多写几个几乎相同的函数。例如上述的`Add`函数，若没有泛型，则需要为int类型与`float64`类型分别定义相加函数。

```go
func Add(x,y int) int {
    return x + y
}

func Add(x,y float64) float64 {
    return x + y
}
```

有了泛型后，通过使用泛型，可以写出更加普适且通用的函数，无需担心类型安全问题。

```go
func Add[T int|float64](x,y T) T {
    return x + y
}
```

**性能优化**

一般而言，泛型代码由于其高度抽象，可能会让使用者担心性能损失。但在Go语言中，泛型的实现方式是在编译期间生成特定类型的代码，因此性能损失通常是可控的。

由于Go编译器在编译期会为每个泛型参数生成具体的实现，因此，运行时不需要额外的类型检查或转换，这有助于优化性能。

```go
//上面使用泛型经过编译期生成以下代码
func Add_int(x,y int) int {
    return x + y
}

func Add_float64(x,y float64) float64 {
    return x + y
}
```

**3.泛型语法**

Go语言泛型包括**类型参数、类型约束、泛型函数、泛型结构体**等应用。

**类型参数**

在Go中，泛型的类型参数通常使用`[]`方括号来声明，方括号紧跟随函数名称或结构体名称。例如：

```go
func Max[T int|float64](x,y T) T {
    if x > y {
        return x
    }
    return y
}
```

上述代码中，定义了一个泛型函数，接受两个类型为`T`的参数`x`和`y`，T是一个类型参数，表示可以是`int`或`float64`类型。`int|float64`表示类型约束，允许`T`可以是int或`float64`类型。

```go
func main(){
    maxInt := Max[int](1,2)
    maxFloat64 := Max[float64](1.1,2.2)
    fmt.Println(maxInt)	//2
    fmt.Println(maxFloat64)	//2.2
}
```

上述`main`函数中，在调用`Max`函数时，可以指定传入`int`类型来调用函数，也可以传入`float64`类型的参数。向`Max`函数提供类型参数(在本例中为`int`和`float64`)称为实例化`（instantiation）`。

类型实例化分为两步进行：

- 首先，编译器在整个泛型函数或类型中将所有类型形参(type paramenters)替换为它们各自的类型实参(type arguments)。
- 其次，编译器验证每个类型参数是否满足相应的约束。

在成功实例化之后，即可得到一个非泛型函数，它可以像任何其他正常函数一样被调用。例如：

```go
func main() {
    max := Max[int]
    fmt.Println(max(1,2)) //2
}
```

**类型约束**

- **内置约束**

Go泛型中，内置了几种类型约束。例如`any`，表示任何类型都可以作为参数。

```go
func Add[T any](x,y T) T {
    return x + y
}
```

上述代码中，`T`表示一个类型参数，该类型参数使用了`any`约束，`any`代表该类型参数可以是任意类型。

Go泛型**不仅支持单一的类型参数，同样也支持定义多个类型参数。**

```go
func ReturnData[T,V any](x T,y T)(T,V){
    return x,y
}
```

- **自定义约束**

类型约束定义了一根数据类型集，只有在这个数据类型集中的数据类型，才能用作泛型的类型参数。

除了内置约束外，也同样支持自定义约束。自定义约束通常是定义`interface`接口来实现的。

```go
func Add[T interface{ int | float64}](x,y T) T {
    return x + y
}
```

上述Add函数中，类型约束接口可以直接在类型参数列表中使用。规定了支持的数据类型为int与float64类型。Go泛型允许使用类型结合，在一个约束中指定多种允许的类型。

通常情况下，外层`interface{}`是可以省略的。

```go
func Add[T int|float64](x,y T) T {
    return x + y
}
```

同样，类型约束接口可以实现定义并加以复用。

```go
type Addable interface {
	int | float64
}

func Add[T Addable](x,y T) T {
    return x + y
}
```

上述通过定义Addable接口来定义一个自定义类型约束接口。在不同的函数中，可以复用该类型约束接口。

**泛型函数**

上述的例子中属于泛型函数，泛型函数可以允许在不同类型上执行相同的函数逻辑。

```go
func Max[T int|float64](x,y T) T {
    if x >= y{
        return x
    }
    return y
}
```

**泛型结构体**

Go泛型不仅支持泛型函数，同样也支持泛型结构体，例如：

```go
type Number[T int|float64] struct {
	num1 T
    num2 int
    num3 float64
}
```

上述定义了一个`Number`的泛型结构体，结构体内的成员变量声明为`T`类型。

**泛型方法**

在使用泛型定义了泛型结构体后，同样也可以使用为泛型结构体定义方法。

```go
func (n Number[T])returnNUm1() T {
    return n.num1
}
```

**泛型应用**

- **泛型接口**

在Go中，同样可以在接口中定义包含泛型的方法，例如：

```go
package main

import "fmt"

type Number[T int|float64] struct {
    num T
}

type OperatorContainer[T int|float64] interface {
    Add(x T) T
    Sub(x T) T
}

func (n Number[T]) Add(element T) T {
    return n.num + element
}

func (n Number[T]) Sub(element T) T {
    return n.num - element
}

func main() {
    var operator1 OperatorContainer[int]
    operator1 = Number[int]{num:10}
    fmt.Println(operator1.Add(20))	//30
    fmt.Println(operator1.Sub(5))	//5
    
    var operator2 OperatorContainer[float64]
    operator2 = Number[float64]{num:10.5}
    fmt.Println(operator2 .Add(20.6))	//31.1
    fmt.Println(operator2 .Sub(4.56))	//5.94
}
```

上述代码中：

- 定义了一个泛型类型`Number[T]`，其中`T`可以是整数`(int)`或浮点数`(float64)`;
- 定义一个泛型接口`OperatorContainer[T],`包含两方法：`Add(x T) T` 和 `Sub(x T) T`,类型约束规定其能使用`int`与`float64`类型；
- 然后，为`Number`结构体实现了`OperatorContainer[T]`接口的方法。`Add`方法将两个数值相加，`Sub`方法将两个数值相减。
- 在`main`函数中，创建了两个泛型变量`operator1`和`operator2`,分别用于存储`int`和`float64`的运算器。同时为两个变量分别创建一个`Number`实例，并调用`Add`和`Sub`方法进行加法和减法运算。

通过泛型类型和接口可以编写更通用、更灵活的代码。在这个例子中，使用泛型实现了一个简单的数值计算器，可以处理`int`和`float64`的加法和减法操作。

- **泛型数据结构**

在实际应用场景中，可以通过泛型来为不同的数据类型定义通用的数据结构，例如栈、队列和堆等。

例如，**使用泛型实现一个栈**：

```go
package stack

type Stack[T any] struct {
    stack []T
}

//Push入栈操作
func (s *Stack[T]) Push(element T){
    if s == nil {
        panic("stack is nil")
    }
    s.stack = append(s.stack,element)
}

//Pop出栈操作
func (s *Stack[T]) Pop() T {
    if len(s.stack) == 0 {
        panic("stack is empty")
    }
    element := s.stack[len(s.stack) - 1]
    s.stack = s.stack[:len(s.stack) - 1]
    return element
}

func main() {
	stack := Stack[int]{}
	stack.Push(5)
	stack.Push(4)
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)
	fmt.Println(stack) //{[5 4 3 2 1]}
	
	stack.Pop()
	fmt.Println(stack)	//{[5 4 3 2]}
}
```

**泛型构建简单缓存系统**

缓存系统通常需要能够存储任意的数据类型，并且能够在给定的时间内根据特定的条件去检索到想获取到的值，可以通过泛型的特性以及Go内置`map`类型来实现一个简单的缓存系统。

首先，可以通过`any`类型约束来定义泛型结构体，设置其内的map类型的成员变量可以存放任意的类型。

```go
type Cachep[T any] struct {
    cache map[string]T
}
```

可以为`Cache`结构体，实现简单的`set`方法与`get`方法来为操作缓存。

```go
func (c *Cache[T]) Set(key string,value T) {
    c.cache[key] = value
}

func (c *Cache[T]) Get(key string) (T,bool) {
    value,exists := c.cache[key]
    return value,exists
}
```

某些场景下，可能需要缓存字符串

```go
func main() {
    c :- &Cache[string] {
        cache : make(make[string]string),
    }
    c.Set("name","serena")
    value,exists := c.Get("name")
    if exists {
        fmt.Println(value)
    }else {
        fmt.Println("存在")
    }
}
```

**总结**

Go泛型是Go中一个极其强大以及灵活的特性，不仅解决了类型安全的问题，同时还带了代码复用以及可维护的强大能力。

泛型不仅仅是一种编程语言的功能或者一个语法糖，它更是一种编程范式的体现。

适当而精妙地应用泛型可以极大地提高代码质量，减少错误，并加速开发过程。

特别是构建大型、复杂的系统时，泛型能够帮助更好地组织代码结构，降低模块之间的耦合度，提高系统的可维护性和可拓展性。

Go泛型没有引入过多复杂的规则和特性，而是集中解决最广泛和最实际的问题。在大多数场景下，Go泛型都能提供清晰、直观和高效的解决方案。

通过深入理解和应用Go的泛型特性，不仅能成为高效的Go开发者，也能更好地理解泛型编程这一通用的编程范式，从而在更广泛的编程任务和问题解决中收益。

#### 四十、反射

**1、什么是反射**

反射是一种程序在运行时可以检查自身变量类型、自身值、有哪些方法等的一种能力。

Go作为一种静态语言，一般而言，我们的变量是什么类型都是一开始就定好的；大多数时候，也能满足我们的使用要求；但是在某些时候，我们无法一开始就确定下来它的具体类型，需要在程序运行的时候，才知道它的具体类型、以及一些其他信息。

那这个"**某些时候**"指的是哪些时候呢？比如：我们编写一个根据任意结构体实例，生成sql语句的功能，我们的结构体实例在一开始可能是Person类型，后续变成了Order类型，它们的类型是不确定的，没办法一开始定好；

另外，我们需要知道这个结构体实例上有哪些字段，才能生成sql语句。

go作为一门静态语言，大多数都是老老实实的，一是一，二是二；但是有了反射，为它提供了一种超越静态的一些能力——它可以在运行时候做很多操作（比如：生成函数、动态调用方法、改变结构体字段值），这么看它已具备了一部分编程的能力。

这种能力在go的许多包、以及框架中都有许多使用，所以才为我们提供了许多好用的功能，所以学习它非常有必要。

在我们为go有这样一种能力欢呼的同时，也需要像大多数go反射相关文章的那样，提醒你反射它的性能不是很好~

最后，由于反射它处理的就是事先不确定具体类型的情况，因此在实际的代码种，我们常常可以看到反射的使用一般和`interface`勾兑在一起，这也见怪不怪。

**2、反射初探**

**初探1**

```go
package main

import (
	"fmt"
    "reflect"
)

type Person struct {
    Name string `param:"name" max:"10"`	//添加两个tag param 和 max
    Age int `param:"age" max:"20"`
}

func main(){
    //先定义一个结构体实例
    p := Person{Name : "张三",Age : 12}
    
    //返回一个反射的类型 reflect.Type
    t := reflect.TypeOf(p)
    //返回反射的值reflect.Value
    v := reflect.ValueOf(p)
    
    //reflect.Type这里的返回是带包名的
    fmt.Println("reflect.Type",t)
    //Kind返回的底层的具体类型  比如这里结构体返回struct
    fmt.Println("kind",t.Kind())
    //Name返回无包名的  名称对于有些类型没有Name比如切片 这里Person
    fmt.Println("Name",t.Name())
    
    //NumField返回结构体有多少个字段
    for i := 0;i < t.NumField();i++ {
        field  := t.Field(i)	//返回字段reflect.StructField是一个结构体
        
        //字段值 注意这里从reflect.Value中取
        fieldVal := v.Field(i)
        fieldType := field.Type //字段类型
        
        //name是结构体中的一个字段  代表字段名称
        //使用Interface()将反射值转换成interface{}类型
        fmt.Printf("第%d个字段：%s type: %s,值：%v\n",i+1,field.Name,fieldType,fieldVal.Interface())
        
        //取tag在内部也是一个结构体reflect.StructTag
        fmt.Printf("字段：%s,param tag is :%s\n",field.Name,field.Tag.Get("param"))
        fmt.Printf("字段：%s,max tag is :%s\n",field.Name,field.Tag.Get("max"))
    }
}
```

```
reflect.Type main.Person
kind struct
Name Person
第1个字段：Name type: string,值：张三
字段：Name,param tag is :name
字段：Name,max tag is :10
第2个字段：Age type: int,值：12
字段：Age,param tag is :age
字段：Age,max tag is :20
```

这里我们需要注意的是`Type`、`Kind`、`Name`三者的区别:

- **Type**：带包名的类型（可理解为全称）
- **Kind**：底层类型比如`int`、`string`、`struct`等等（忽略类型别名）
- **Name**：不带包的名称，结构体返回结构体名称（有些类型无Name，比如切片）

**初探2**

我们除了可以解析结构体中的字段、类型、tag外，我们还可以修改结构体中的字段值、动态调用结构体方法，我们一起来看下。

```go
package main

import (
	"fmt"
    "reflect"
)

type Person struct {
    Name string `param:"name" max:"10"`	//添加两个tag param 和 max
    Age int `param:"age" max:"20"`
}

func (u Person) PrintInfo() {
    fmt.Printf("姓名：%s,年龄：%d\n",u.Name,u.Age)
}

func main() {
    //先定义一个结构体实例
    p := Person{Name : "张三",Age : 12}
    
    //返回一个反射类型reflect.Type
    t := reflect.TypeOf(p)
    //返回反射的值reflect.Value
    v := reflect.ValueOf(p)
    
    //列出结构体有哪些方法
    //NumMethod返回方法个数
    for i := 0; i < t.NumMethod(); i++ {
        //reflect.Method返回一个结构体
        m := t.Method(i)
        fmt.Printf("方法名：%s，方法类型：%s\n",m.Name,m.Type)
    }
    
    //根据名称找调用方法  从value中取（凡是和执行调用相关的都从value中取）
    m1 := v.MethodByName("PrintInfo")
    //调用方法 参数这里应该用reflect.Value切片 如果无参数可以用nil
    m1.Call(nil)
}
```

我还可以修改变量的值，有三点需要注意：

1. 传递指针地址作为参数
2. 使用Elem()
3. `reflect.Value`转换回具体的类型采用`interface().(xx类型)`

**3、反射与正常使用之间的桥梁**

经过前面初探部分代码实践，我们已经对go的反射使用有些了解；但是还不够清晰；比较零散，这里我们一起归类总结下。

go的反射和我们正常（普通）使用相比，就像两个不同的世界，在它们的世界都有各自的规律 —— 不同的使用方式;但是它们并非完全隔绝,有两个东西是他们之间的桥梁：

1. `reflect.TypeOf()` 类型相关
2. `reflect.ValueOf()` 值相关

透过这两个东西，我们们可以穿梭于反射、正常（普通）使用两个世界。

相信您在看了更多反射代码后会有更深刻的体会，好的！我们继续前行。

前面我们对go结构体相关反射就行了初探，但实际上反射能做的事远远不止于此，我们可以先有这么一个认知：

> 通过反射能做到许多go正常（普通）代码能做到的许多操作，比如：构建切片、构建映射、构建函数、结构体等等。

在开始之前先记住一个规律，操作反射世界就需要使用反射世界的元素去操作，比如在反射赋值就需要给一个`reflect.Value`类型的值。

**4、构建切片**

我们可以通过反射创建切片，主要通过`reflect.MakeSlice`实现

```go
package main

import (
	"fmt"
    "reflect"
)

func main() {
    //准备一个切片
    s:= make([]int,0)
    
    //获取切片的类型
    sliceType := reflect.TypeOf(s)
    //构建切片和普通make写法类似
    sliceValue := reflect.MakeSlice(sliceType,0,0)
    //准备填充reflect Slice的值 类型必为reflect.Value类型
    svalue := reflect.ValueOf(1)
    //往切片中放元素
    //PS : 不能这么写sliceValue[0] = svalue
    sliceValue = reflect.Append(sliceValue,svalue)
    
    //转成普通世界的值
    slice := sliceValue.Interface().([]int)
    fmt.Println("当前slice:",slice)
}

// 当前slice:  [1]
```

**5、构建map**

主要通过`reflect.MakeMap`实现

```go
package main 

import (
	"fmt"
    "reflect"
)

func main() {
    //准备一个map类型
    m := make(map[string]int)
    
    //获取反射map类型
    mapType := reflect.TypeOf(m)
    //构建反射map值
    mapValue := reflect.MakeMap(mapType)
    //准备反射key
    key := reflect.ValueOf("one")
    val := reflect.ValueOf(1)
    
    //map赋值
    mapValue.SetMapIndex(key,val)
    
    //转换称普通世界值
    convertMap := mapValue.Interface().(map[string]int)
    
    //打印看看
    fmt.Printf("%#v\n",convertMap)
}

// map[string]int{"one":1}
```

**6、构建函数**

主要通过reflect.MakeFunc实现

```go
package main

import (
	"fmt"
    "reflect"
    "runtime"
    "time"
)

func main() {
    //末尾转换成一个正常普通函数
    tHello := timeSpend(hello).(func(string))
    //执行新函数
    tHello("dmy")
}

//一个hello函数 做什么无所谓主要用于测试
func hello(name string) {
    fmt.Printf("hello %s\n",name)
    
    //模拟耗时操作
    time.Sleep(time.Second)
}

//生成一个计算函数执行时间的函数
//入参数 ： 一个函数 interface{}
//返回值 ： 一个函数 interface{}类型
func timeSpend(f interface{}) interface{} {
    //反射类型
    t := reflect.TypeOf(f)
    //如果传入的不是函数类型就报错
    if t.Kind() != reflect.Func {
        panic("need a function")
    }
    
    //反射值
    v := reflect.ValueOf(f)
    //构建反射函数
    wrapperF := reflect.MakeFunc(t,func(in []reflect.Value) []reflect.Value {
        //获取当前时间
        start := time.Now()
        //调用函数
        out := v.Call(in)
        //计算执行时间
        //FuncForPC 用于计算函数名
        fmt.Printf("call %s spend %v\n",runtime.FuncForPC(v.Pointer()).Name(),time.Since(start))
        //返回
        return out
    })
	//返回函数的反射值的Interface
    return wrapperF.Interface()
}

//hello dmy
//call main.hello spend 1s
```

**7、构建结构体**

构建结构体和前面的稍微复杂些，主要有三步：

1. 构建结构体字段切片
2. 通过reflect.StructOf()构建结构体类型
3. 通过reflect.New()构建出结构体reflect.value

```go
package main

import (
	"fmt"
    "reflect"
)

func main() {
    //这会构建出一个三个字段
    //Field1 int
    //Field2 string
    //Field3 bool
    //需要注意的是这里构建的结构体
    //一般我们是无法实现写一个结构体来转换的 因为类型随意了
    s := MakeStruct(10,"abc",true).(*struct {
        Field1 int
    	Field2 string
    	Field3 bool
    })
    s.Field1 = 12
    fmt.Printf("%#v\n",*s)
    
    //大多数时候我们通过反射去修改生成结构体中的值
    //返回的是一个指针
    s1 := MakeStruct("a",10)
    //取valueOf 转Elem来改结构体的值
    s1Value := reflect.ValueOf(s1).Elem()
    s1Value.Field(0).SetString("abc")
    s1Value.Field(1).SetInt(12)
    
    //打印值
    fmt.Printf("s1: %#v\n",s1Value.Interface())
}

//构建结构体
//入参数：任意个任意类型参数
//每个参数代码 字段的类型 顺序
func MakeStruct(vals ...interface{}) interface{} {
    //1.准备结构体切片
    //准备一个切片 用于存储结构体字段数据
    //这里类型用的是structField
    structSlice := make([]reflect.StructField,len(vals))
    
    //遍历参数
    for i,val :=range vals {
        //StructField本身就是一个结构体
        sf := reflect.StructField {
            Name : fmt.Sprintf("Field%d",i+1),	//字段名
            Type : reflect.TypeOf(val),			//字段类型这里要用reflect.Type
        }
        
        structSlice[i] = sf //存入切片
    }
    
    //2.构建结构体类型
    sType := reflect.StructOf(structSlice)
    //3.构建结构体反射值
    sValue := reflect.New(sType)
    //4.转为interface类型返回
    return sValue.Interface()
}

//输出结果：
//struct { Field1 int; Field2 string; Field3 bool }{Field1:12, Field2:"", Field3:false}
//s1: struct { Field1 string; Field2 int }{Field1:"abc", Field2:12}
```

**8、应用的例子**

最后我们写一个常用的例子收尾，写一个在数据映射中经常要做的操作，通过结构体实例生成相关的插入sql语句。

```go
package main

import (
	"fmt"
    "reflect"
    "strings"
)

//user结构体
type user struct {
    name string
    age int
}

func main() {
    u := user{name : "zhangs",age : 17}
    fmt.Println(GenerateSql(u))
}

//生成任意结构体实例的insert sql语句
//为简单起见 这里只考虑字段有两种类型 string 和 int
//为简单起见 这里没有考虑复数 大小写的情况
//比如一个user{Name:"zhangs",age:17}生成的sql语句为 insert into user(name,age) values("张三",18)
func GenerateSql(s interface{}) (sql string) {
    //反射类型
    t := reflect.TypeOf(s)
    //检查传入的类型
    if t.Kind() != reflect.Struct {
        fmt.Println("传入的不是结构体")
        return ""
    }
    
    //反射的值备用
    v := reflect.ValueOf(s)
    //字段名切片
    fieldNames := make([]string,0,t.NumField())
    //字段值切片 任意类型 统统先转换成字符串 方便后续join后拼接字符串
    fieldValues := make([]string,0,t.NumField())
    
    //循环取值
    for i := 0;i < t.NumField(); i++ {
        //取字段名
        fieldNames = append(fieldNames,t.Field(i).Name)
        //字段
        field := v.Field(i)
        //根据字段类型做不同操作
        switch field.Kind() {
        case reflect.String:
            //如果字段类型不是string用string回报错
            fieldValues = append(fieldValues,field.String())
        case reflect.Int:
            //int转字符
            fieldValues = append(fieldValues,fmt.Sprintf("%v",field.Int()))
        }
    }
    
    //表明
    tableName := t.Name()
    //拼接sql
    sql = fmt.Sprintf("insert into %s(%s) values(%s)",tableName,strings.Join(fieldNames,","),strings.Join(fieldValues,","))
    //将数组数据转换成字符串格式
    fmt.Println(fieldValues)
	fmt.Println(strings.Join(fieldValues, ","))
    return
}
//[zhangs 17]
//zhangs,17
//输出结果：insert into user(name,age) values(zhangs,17)
```

****

#### 四十一、网络编程

go这门语言由于它有goroutine和channel非常适合web编程，能够实现很高的并发。

网络编程本质是利用socket，底层借助的是操作系统提供的一系列能力；当然go自身也做了很多优化；

go提供的网络编程总体而言，和其他语言差别不大，也比较好懂；go屏蔽了底层的复杂性，将复杂留给自己，将简单留给使用者，真的很不错。

**Websocket是什么？**

- Websocket是一种在单个TCP连接上进行全双工通信的协议
- Websocket使得客户端和服务器之间的数据交换变得更加简单，允许服务端主动向客户端推送数据
- 在Websocket API中，浏览器和服务器只需完成一次握手，两者之间就直接可以创建持久性的连接，并且进行双向数据传输
- 需要安装第三方包：
  - cmd中：go get -u -v github.com/gorilla/websocket

**1.TCP实现**

**1.1简单TCP服务端-客户端**

- **服务端**

```go
package main

import (
	"fmt"
    "net"
)

func main() {
    //1.直接listen
    listen,err := net.Listen("tcp",":8080")
    if err != nil {
        fmt.Println("listen error:",err)
        return
    }
    
    defer listen.Close()
    
    //打印出监听的端口信息
    fmt.Println("listening on ",listen.Addr().String())
    
    for{
        //2.Accept接受客户端请求
        conn,err := listen.Accept()
        if err != nil {
            fmt.Println("accept error:",err)
            continue //跳过本次 继续下一次
        }
        
        //3.开一个协程 单独处理连接
        go handelConn(conn)
    }
}

//处理接收到的连接
//net.Conn是一个i接口，拥有Read和Write方法
func handelConn(conn net.Conn) {
    //准备从conn中读取数据 存储的切片
    data := make([]byte,512)
    //但会读取的字节数，错误
    n,err := conn.Read(data)
    if err != nil {
        fmt.Println("Read error:",err)
        return
    }
    fmt.Printf("读取到客户端ip:%s,长度为:%v的信息:%s\n",conn.RemoteAddr(),n,string(data))
    
    //回写客户端消息
    _,err = conn.Write([]byte("我们已经收到你的请求"))
    if err != nil {
        fmt.Println("write error:",err)
        return
    }
}
```

- 客户端

```go
package main

import (
	"fmt"
    "net"
)

func main() {
    //1.客户端拨号 tcp ip:port
    conn,err := net.Dial("tcp","localhost:8080")
    if err != nil {
        fmt.Println("Dial error:",err)
        return
    }
    
    defer conn.Close()
    
    //2.发送数据
    _,err = conn.Write([]byte("hello world"))
    if err != nil {
        fmt.Println("Write error:",err)
        return
    }
    
    //3.读取数据
    data := make([]byte,512)
    n,err := conn.Read(data)
    if err != nil {
        fmt.Println("Read error:",err)
        return
    }
    
    fmt.Printf("Read data length : %d,conten : %s", n, data)
}


// Read data length : 30,conten : 我们已经收到你的请求
```

**1.2粘包问题**

上述的代码虽然能正常工作，但是如果客户端一次性发送很多信息的适合，会有粘包问题。

1. 把客户端改成循环发送很多次（比如60次）
2. 服务端读取客户端信息改成循环读取

- 客户端

```go
package mian

import (
	"fmt"
    "net"
    "strings"
)

func main() {
    //1.客户端拨号tcp ip:port
    conn,err := net.Dial("tcp","localhost:8080")
    if err != nil {
        fmt.Println("Dial error:",err)
        return
    }
    
    defer conn.Close()
    
    //构建一个很长的字符串 由60个 hello world组成
    str := strings.Repeat("hello world",60)
    
    //2.发送数据
    _,err = conn.Write([]byte(str))
    if err != nil {
        fmt.Println("Write error:",err)
        return
    }
}
```

- 服务端

```go
package main

import (
	"fmt"
    "net"
)

func main() {
    //1.直接listen
    listen,err := net.Listen("tcp",":8080")
    if err != nil {
        fmt.Println("listen error:",err)
        return
    }
    
    defer listen.Close()
    
    //打印监听的端口信息
    fmt.Println("listening on ",listen.Addr().String())
    
    for {
        //2.accept接收客户端请求
        conn,err := listen.Accept()
        if err != nil {
            fmt.Println("accept error:",err)
            continue	//跳过本次 继续下一次
        }
        
        go handelConn(conn)
    }
}

//处理接收到的连接
//net.Conn是一个接口，拥有Read和Write方法
func handelConn(conn net.Conn) {
    //准备从conn中读取是估计 存储的切片
    data := make([]byte,512)
    
    for {
        //返回读取的字节数，错误
        n,err := conn.Read(data)
        if err != nil {
            fmt.Println("Read error:",err)
            return
        }
        fmt.Printf("读取到客户端ip:%s,长度为：%v的信息:%s\n",conn.RemoteAddr(),n,string(data))
    }
}
```

