package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//
// type error interface {
// 	Error() string
// }
//

var MyError2 = errors.New("I am Error 2")

type MyError1 struct {
	Message string
	Number  int
}

func (me MyError1) Error() string {
	return fmt.Sprintf("%s: %d", me.Message, me.Number)
}

func toNumber(s string) int {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		fmt.Println(err.Error())
	}

	return int(i)
}

func iReturnError() (int, error) {

	n := rand.Int()

	if n%2 == 0 {
		return 0, MyError1{
			Message: "Hello, I an Error 1",
			Number:  n,
		}
	}

	return 1, MyError2
}

func iReturnErrorError() error {
	_, err := iReturnError()
	if err != nil {
		return fmt.Errorf("from iReturnErrorError %w", err)
	}

	return nil
}

func main() {
	rand.Seed(time.Now().UnixMilli())
	fmt.Println(toNumber("123"))
	fmt.Println(toNumber("abc"))
	fmt.Println(toNumber("123abc"))

	fmt.Println(errors.Is(iReturnErrorError(), MyError2))
	fmt.Println(errors.Is(iReturnErrorError(), MyError2))
	fmt.Println(errors.Is(iReturnErrorError(), MyError2))
	fmt.Println(errors.Is(iReturnErrorError(), MyError2))
	fmt.Println(errors.Is(iReturnErrorError(), MyError2))
	fmt.Println(errors.Is(iReturnErrorError(), MyError2))
	fmt.Println(errors.Is(iReturnErrorError(), MyError2))
	fmt.Println(errors.Is(iReturnErrorError(), MyError1{}))

	myError1 := new(MyError1)

	if ok := errors.As(iReturnErrorError(), myError1); ok {
		fmt.Println(myError1.Message)
	}


}