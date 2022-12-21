Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false

Непустой интерйфейс под капотом имеет два указателя: на тиблицу методов и на значение. Тип error представляет из себя интерфейс, требующий от типа наличие метода Error() string. Тип *os.PathError реализует этот метод, поэтому *os.PathError может быть присвоен переменной типа error.
В нашем случае itable указывает на таблицу методов типа *os.PathError, а data равен nil. Тем не менее err не равен nil, так как интерфейс не пустой.

fmt.Println(err) выведет <nil>, так как в интерфейсе нет значения, а fmt.Println(err == nil) выведет false, так как интерфейс не пустой.
```