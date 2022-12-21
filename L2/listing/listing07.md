Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
1'''|
2   |
3   |
4   | - вперемешку, но элементы отдельного массива перемешиваться не могут
5   |
6   |
7   |
8...|
0
0
0
...

Функция слияния каналов реализована неверно, так как при закрытии одного из каналов-источников, стение из них по-прежнему будет выполняться, но возвращаемое значение всегда будет 0. Для исправления нужно при чтении из канал происводить проверку успешности чтения:
```

```go
case v, ok := <-a:
    if !ok {
        break loop
    }
    c <- v
```

```
Помимо прочего необходимо добавить вызов функции закрытия каналв по окончанию цикла в функции merge, иначе выбудет deadlock при чтении из канала c в main.

```