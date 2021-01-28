package main

import (
	"fmt"
)

type reader interface {
	read(b []byte) (int, error)
}

type notifier interface {
	notify()
}

type printer interface {
	print()
}

type file struct {
	name string
}

type pipe struct {
	name string
}

type user struct {
	name  string
	email string
}

type data struct {
	name string
	age  int
}

func (u user) print() {
	fmt.Printf("My name is %s and my email is %s\n", u.name, u.email)
}

func (u *user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func (a *admin) notify() {
	fmt.Printf("Sending admin email To %s<%s>\n", a.name, a.email)
}

func (file) read(b []byte) (int, error) {
	s := "<rss><channel><title>Going Go Programming</title></channel></rss>"
	copy(b, s)
	return len(s), nil
}

func (pipe) read(b []byte) (int, error) {
	s := `{name: "hoanh", title: "developer"}`
	copy(b, s)
	return len(s), nil
}

func retrieve(r reader) error {
	data := make([]byte, 100)
	len, err := r.read(data)
	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))
	return nil
}

func (d data) displayName() {
	fmt.Println("My Name Is", d.name)
}

func (d *data) setAge(age int) {
	d.age = age
	fmt.Println(d.name, "Is Age", d.age)
}

func (u *user) changeEmail(email string) {
	u.email = email
	fmt.Printf("Changed User Email To %s\n", email)
}

func (u *user) String() string {
	return fmt.Sprintf("My name is %q and my email is %q", u.name, u.email)
}

func sendNotification(n notifier) {
	n.notify()
}

type duration int

func (d *duration) notify() {
	fmt.Println("Sending Notification in", *d)
}

type admin struct {
	user
	level string
}

func main() {
	// Method
	bill := user{"Bill", "bil@email.com"}
	bill.notify()
	bill.changeEmail("bill@hotmail.com")

	hoanh := &user{"Hoanh", "hoanhan@email.com"}
	hoanh.notify()
	hoanh.changeEmail("hoanhan101@gmail.com")

	users := []user{
		{"bill", "bil@email.com"},
		{"hoanh", "hoanhan@email.com"},
	}
	for _, u := range users {
		u.notify()
	}

	d := data{
		name: "Hoanh",
	}
	fmt.Println("Proper Calls to Methods:")

	d.displayName()
	d.setAge(21)

	fmt.Println("\nWhat the Compiler is Doing:")
	data.displayName(d)
	(*data).setAge(&d, 21)

	fmt.Println("\nCall Value Receiver Methods with Variable:")
	f1 := d.displayName
	d.name = "Hoanh An"
	f1()

	fmt.Println("\nCall Pointer Receiver Method with Variable:")
	f2 := d.setAge
	d.name = "Hoanh An Dinh"
	f2(21)

	f := file{"data.json"}
	p := pipe{"cfg_service"}

	retrieve(f)
	retrieve(p)
	u := user{"Hoanh", "hoanhan@email.com"}
	sendNotification(&u)
	//duration(42).notify()

	fmt.Println(u)
	fmt.Println(&u)

	entities := []printer{
		u,
		&u,
	}

	u.name = "Hoanh An"
	u.email = "hoanhan101@gmail.com"

	for _, e := range entities {
		e.print()
	}

	ad := admin{
		user: user{
			name:  "Hoanh An",
			email: "hoanhan101@gmail.com",
		},
		level: "superuser",
	}

	ad.user.notify()
	ad.notify()

	sendNotification(&ad)
}
