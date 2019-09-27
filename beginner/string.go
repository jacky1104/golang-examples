package main

// import required modules
import (
	"fmt"
)

// main function
func main() {

	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println(sample)

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")
	fmt.Printf("%x\n", sample)
	fmt.Printf("% x\n", sample)
	fmt.Printf("%q\n", sample)
	fmt.Printf("%+q\n", sample)

	const placeOfInterest = `⌘`

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")

	const country = "中国"
	// utf-8 byte
	fmt.Printf("%x\n", country)
	fmt.Printf("% x\n ", country)

	// unicode
	fmt.Printf("%+q\n ", country)
	fmt.Printf("%q\n ", country)

	for i := 0; i < len(country); i++ {
		fmt.Printf("%x ", country[i])
	}
	fmt.Println("chinese character to byte and unicode")
	for index, value := range country {
		fmt.Printf("%d, %x ", index, value)
		fmt.Printf("%d, %+q ", index, value)
		fmt.Printf("%d, %v ", index, value)
		fmt.Printf("%d, %+v ", index, value)
		fmt.Printf("%d, %s \n", index, string(value))
	}

	fmt.Println("chinese character to rune")
	for index, value := range []rune(country) {
		fmt.Printf("%d, %x ", index, value)
		fmt.Printf("%d, %+q ", index, value)
		fmt.Printf("%d, %v ", index, value)
		fmt.Printf("%d, %+v ", index, value)
		fmt.Printf("%d, %s \n", index, string(value))
	}
}
