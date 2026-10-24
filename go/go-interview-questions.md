# Go 面试题记录

## 1. map、slice 与 for 遍历

### 题目

```go
package main

import "fmt"

func main() {
	m := make(map[string][]int)

	for i := 0; i < 3; i++ {
		key := fmt.Sprintf("key%d", i%2)
		m[key] = append(m[key], i)
	}

	for key, values := range m {
		fmt.Println(key, values)
	}
}
```

### 输出

可能输出：

```text
key0 [0 2]
key1 [1]
```

也可能输出：

```text
key1 [1]
key0 [0 2]
```

### 题解

循环过程中，`i % 2` 的结果依次为 `0、1、0`：

```text
i = 0：key0 -> [0]
i = 1：key1 -> [1]
i = 2：key0 -> [0 2]
```

`map[string][]int` 的 value 是一个切片。访问不存在的 key 时，`m[key]` 会返回该切片类型的零值 `nil`，而 `append(nil, value)` 是合法的，因此可以直接追加。

需要注意，Go 的 map 遍历顺序是不固定的，所以 `key0` 和 `key1` 的输出顺序不能依赖。

## 2. range 遍历 slice 时修改元素

### 题目

```go
package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	for _, value := range s {
		value *= 2
	}

	fmt.Println(s)
}
```

### 输出

```text
[1 2 3]
```

### 题解

`range` 的第二个变量 `value` 是切片元素的副本：

```go
for _, value := range s {
```

因此：

```go
value *= 2
```

只修改了副本，不会修改原切片。

如果需要修改原切片，应使用下标：

```go
for i := range s {
	s[i] *= 2
}
```

此时输出为：

```text
[2 4 6]
```

## 4. 使用 append 删除 slice 中间元素

### 题目

```go
package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	s = append(s[:1], s[2:]...)

	fmt.Println(s)
}
```

### 输出

```text
[1 3]
```

### 题解

先看两个切片表达式：

```go
s[:1] // [1]
s[2:] // [3]
```

因此代码等价于：

```go
s = append([]int{1}, []int{3}...)
```

最终结果是：

```text
[1 3]
```

这里使用 `append(s[:1], s[2:]...)` 可以移除下标为 `1` 的元素 `2`。如果原切片容量足够，`append` 可能直接复用底层数组，并覆盖原来的数据；赋值给 `s` 后，切片长度变为 `2`。

## 3. 两个 Goroutine 交替打印字母和数字

### 题目

使用两个 Goroutine，通过 channel 交替打印字母和数字，输出：

```text
A1B2C3...Z26
```

### 参考实现

```go
package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	letters := make(chan int)
	numbers := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := range letters {
			fmt.Printf("%c", 'A'+i-1)
			if i == 26 {
				close(numbers)
				return
			}
			numbers <- i
		}
	}()

	go func() {
		defer wg.Done()

		for value := range numbers {
			fmt.Print(strconv.Itoa(value))
			letters <- value + 1
		}
	}()

	letters <- 1
	wg.Wait()
}
```

### 题解

两个 channel 分别负责两个 Goroutine 之间的交接：

```text
字母 Goroutine 打印 A
        |
        v
数字 Goroutine 打印 1
        |
        v
字母 Goroutine 打印 B
```

执行流程如下：

1. 主 Goroutine 向 `letters` 发送 `1`，启动第一次打印。
2. 字母 Goroutine 收到 `1`，打印 `A`，然后向 `numbers` 发送 `1`。
3. 数字 Goroutine 收到 `1`，打印 `1`，然后向 `letters` 发送 `2`。
4. 字母 Goroutine 收到 `2`，打印 `B`，继续交接给数字 Goroutine。
5. 当打印到 `Z26` 后，字母 Goroutine 关闭 `numbers`。
6. 数字 Goroutine 使用 `range numbers`，检测到 channel 关闭后自动退出。
7. `WaitGroup` 等待两个 Goroutine 结束，避免主 Goroutine 提前退出。

这里由生产数据的一方关闭 channel。字母 Goroutine 知道什么时候不会再发送数据，因此由它负责关闭 `numbers`。

## 5. slice、append 与 defer

### 题目

```go
package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	defer func() {
		s = append(s, 4)
		fmt.Println("defer:", s)
	}()

	s = append(s, 10)

	s2 := s[:2]
	s2 = append(s2, 20)

	fmt.Println("main:", s)
}
```

### 输出

```text
main: [1 2 20 10]
defer: [1 2 20 10 4]
```

### 题解

初始切片 `s` 的长度和容量都是 `3`。执行 `append(s, 10)` 时容量不足，因此会扩容并得到新的底层数组：

```text
s = [1 2 3 10]
```

接着执行：

```go
s2 := s[:2]
s2 = append(s2, 20)
```

`s2` 与 `s` 共享底层数组，并且 `s2` 的容量仍然足够，所以追加 `20` 时不会重新分配数组，而是直接覆盖底层数组下标 `2` 的元素：

```text
s  = [1 2 20 10]
s2 = [1 2 20]
```

`main` 函数中的普通语句先执行，因此先打印 `main`。函数返回前执行 `defer`，此时追加 `4`，最终打印 `defer`。
