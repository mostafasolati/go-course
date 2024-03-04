package main

import (
	"errors"
	"fmt"
	"strconv"
)

var err1 = errors.New("error1 happend")
var err2 = errors.New("error2 happend")

func TestFunc() error {
	return err2
}

func MultiErrors() error {
	return errors.Join(err1, err2, errors.New("new errrors"))
}

type MyError struct {
	Message string
	Code    int
}

func (s MyError) Error() string {
	return s.Message
}

func CheckPassword(p string) error {
	if p == "123" {
		return nil
	}
	return errors.New("invalid password")
}

func RegisterUser(username, password string) error {
	err := CheckPassword(password)
	if err != nil {
		// return err
		return fmt.Errorf("register user function failed : %w", err)
	}
	return nil
}

func CheckPassword2(p string) error {
	if p == "123" {
		return nil
	}
	err := MyError{
		Message: "Invalid Password",
		Code:    0xDEADC0DE, // 3735929054 hex speak
	}
	return err
}

func main() {
	s := "hello"
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error happened:", err)
	} else {
		fmt.Println("Value:", n)
	}

	if n, err := strconv.Atoi(s); err != nil {
		fmt.Println("Error happend:", err)
	} else {
		fmt.Println("Value:", n)
	}

	if err := CheckPassword("123"); err != nil {
		fmt.Println("UnAuthorized")
	} else {
		fmt.Println("Successfully logged in")
	}

	if err := CheckPassword2("xyz"); err != nil {
		myerr := new(MyError)
		err = errors.New("test error")
		if errors.As(err, myerr) {
			fmt.Println("CODE:", myerr.Code)
		} else {
			fmt.Println("Invalid Error Type")
		}
		fmt.Println("UnAuthorized")
	} else {
		fmt.Println("Successfully logged in")
	}

	if err := RegisterUser("myuser", "xyz"); err != nil {
		underlyingErr := errors.Unwrap(err)
		fmt.Println(err)
		fmt.Println(underlyingErr)
	}

	// if err := TestFunc(); err == err2 {
	// 	fmt.Println("Equal")
	// }

	if err := TestFunc(); errors.Is(err, err1) {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}

	err = MultiErrors()
	fmt.Println("Multi error:", err)

}
