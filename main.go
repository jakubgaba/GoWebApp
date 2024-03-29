package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("this is athe about page and 2 + 2 is: %d", sum))
}

func Divide(w http.ResponseWriter, r *http.Request) {
	sum, err := divideValues(50.2, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Smthing happened wrong")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 50.2, 0.0, sum))
}

type Num interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func addValues[T ~int | float32](x T, y T) T {
	return x + y
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	_ = http.ListenAndServe(portNumber, nil)
}
