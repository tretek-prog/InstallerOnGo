package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"log"
)

func main() {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\ZZZ`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
		fmt.Println("A")
	}
	if err := k.SetStringValue("xyz", "blahblah"); err != nil {
		log.Fatal(err)
		fmt.Println("B")
	}
	if err := k.Close(); err != nil {
		log.Fatal(err)
		fmt.Println("C")
	}
}
