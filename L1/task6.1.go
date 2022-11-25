package main

/*
	a 1000 and 1 ways to kill a goroutine (3 actually)

	just exit main...
	not good solution, but it gets the job done

	when main exits, program termitanes, and all running
	goroutines are terminated at random point of execution
	according to Golang reference docs and stackoverflow answers
*/

func main() {
	go func() {
		for {
			// do some stuff
		}
	}()
}
