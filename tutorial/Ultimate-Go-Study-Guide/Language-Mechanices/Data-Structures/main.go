package main

import (
	"fmt"
	"unicode/utf8"
)

func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i := range slice {
		fmt.Printf("[%d] %p %s\n", i, &slice[i], slice[i])
	}
}

type user struct {
	name     string
	username string
}

func main() {
	var strings [5]string
	strings[0] = "Apple"
	strings[1] = "Orange"
	strings[2] = "Banana"
	strings[3] = "Grape"
	strings[4] = "Plum"
	fmt.Printf("\n=> Iterate over array\n")
	for i, fruit := range strings {
		fmt.Println(i, fruit)
	}

	six := [6]string{"Annie", "betty", "Charley", "Doug", "Edward", "Hoanh"}
	for i, v := range six {
		fmt.Printf("Value[%s]\tAddress[%p] IndexAddr[%p]\n", v, &v, &six[i])
	}

	slice2 := make([]string, 5, 8)
	slice2[0] = "Apple"
	slice2[1] = "Orange"
	slice2[2] = "Banana"
	slice2[3] = "Grape"
	slice2[4] = "Plum"

	fmt.Printf("\n=? Length vs Capacity\n")
	inspectSlice(slice2)

	var data []string
	lastCap := cap(data)
	for record := 1; record <= 102400; record++ {
		data = append(data, fmt.Sprintf("Rec: %d", record))
		if lastCap != cap(data) {
			capChg := float64(cap(data)-lastCap) / float64(lastCap) * 100
			lastCap = cap(data)
			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n", &data[0], record, cap(data), capChg)
		}
	}

	// Copy a slice
	slice4 := make([]string, len(slice2))
	copy(slice4, slice2)
	fmt.Printf("\n=> Copy a slice\n")
	inspectSlice(slice4)

	// Slice and reference
	x := make([]int, 7)
	for i := 0; i < 7; i++ {
		x[i] = i * 100
	}
	twohundred := &x[1]
	x = append(x, 800)
	x[1]++
	fmt.Printf("\n=> Slice and reference\n")
	fmt.Println("twohundred:", *twohundred, "x[1]:", x[1])

	// UTF-8
	s := "世界 means world"
	var buf [utf8.UTFMax]byte
	for i, r := range s {
		rl := utf8.RuneLen(r)
		si := i + rl
		copy(buf[:], s[i:si])
		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n", i, r, r, buf[:rl])
	}

	// Map
	users1 := make(map[string]user)
	users1["Roy"] = user{"Rob", "Roy"}
	users1["Ford"] = user{"Henry", "Ford"}
	users1["Mouse"] = user{"Mickey", "Mouse"}
	users1["Jackson"] = user{"Michael", "Jackson"}

	fmt.Printf("\n=> Iterate over map\n")
	for key, value := range users1 {
		fmt.Println(key, value)
	}

	users2 := map[string]user{
		"Roy":     {"Rob", "Roy"},
		"Ford":    {"Henry", "Ford"},
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Michael", "Jackson"},
	}

	fmt.Printf("\n=> Map literals\n")
	for key, value := range users2 {
		fmt.Println(key, value)
	}

	delete(users2, "Roy")

	u1, found1 := users2["Roy"]
	u2, found2 := users2["Ford"]

	fmt.Printf("\n=> Find key\n")
	fmt.Println("Roy", found1, u1)
	fmt.Println("Ford", found2, u2)

}
