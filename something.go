package main

import (
  "fmt"
  "log"
  "golang.org/x/sys/windows/registry"
) 


func main() {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Intel`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()
	fmt.Println("dd")

	s, _, err := k.GetStringValue("Default")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Windows system root is %q\n", s)
}
