package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type exa struct {
	flag    bool
	counter int16
	pi      float32
}

type example struct {
	counter int64
	pi      float32
	flag    bool
}

func increment1(inc int) {
	inc++
	println("inc1:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}

func increment2(inc *int) {
	*inc++
	println("inc2:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tvalue Points To[", *inc, "]")
}

type user struct {
	ID   int
	Name string
}
type updateStats struct {
	Modified int
	Duration float64
	Success  bool
	Message  string
}

func retrieveUser(name string) (*user, error) {
	r, err := getUser(name)
	if err != nil {
		return nil, err
	}

	var u user
	err = json.Unmarshal([]byte(r), &u)

	return &u, err
}

func getUser(name string) (string, error) {
	response := `{"ID":101, "Name":"Kwangh"}`
	return response, nil
}

func updateUser(u *user) (*updateStats, error) {
	response := `{"Modified":1, "Duration":0.005, "Success" : true, "Message": "updated"}`

	var us updateStats
	if err := json.Unmarshal([]byte(response), &us); err != nil {
		return nil, err
	}

	if us.Success != true {
		return nil, errors.New(us.Message)
	}

	return &us, nil
}

func main() {
	var a int
	var b string
	var c float64
	var d bool
	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)

	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)

	aaa := int32(10)
	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)

	var e1 exa
	fmt.Printf("%+v\n\n", e1)

	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141692,
	}

	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)

	e3 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	fmt.Println()
	fmt.Println("Flag", e3.flag)
	fmt.Println("Counter", e3.counter)
	fmt.Println("Pi", e3.pi)

	var e4 exa
	e4 = e3
	fmt.Printf("%+v\n", e4)

	count := 10
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
	increment1(count)
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
	increment2(&count)
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	u, err := retrieveUser("Kwangh")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", *u)

	if _, err := updateUser(u); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Updated user record for ID", u.ID)

	const (
		A1 = iota
		B1 = iota
		C1 = iota
	)

	fmt.Println("1:", A1, B1, C1)

	const (
		A2 = iota
		B2
		C2
	)

	fmt.Println("2:", A2, B2, C2)

	const (
		A3 = iota + 1
		B3
		C3
	)

	fmt.Println("3:", A3, B3, C3)

	const (
		Ldate = 1 << iota
		Ltime
		Lmicroseconds
		Llongfile
		Lshortfile
		LUTC
	)

	fmt.Println("Log:", Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC)
}
