Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error

Программа выведет error, так как сравнивается не указатель, возвращенный функцией test(), а интерфейс error, который содержит в себе указатель на структуру customError. В нашем случае интерфейс error содержит в себе указатель на структуру customError, который не равен nil и условие if err != nil выполняется. Тем не менее указатель data в err равен nil.
```