package debug

import "fmt"

type Debug bool

func (d Debug) Printf(format string, a ...interface{}) (n int, err error) {
	if d {
		return fmt.Printf(format, a...)
	}

	return 0, nil
}

func (d Debug) Println(a ...interface{}) (n int, err error) {
	if d {
		return fmt.Println(a...)
	}

	return 0, nil
}
