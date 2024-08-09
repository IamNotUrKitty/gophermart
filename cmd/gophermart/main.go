package main

import "github.com/IamNotUrKitty/gophermart/internal"

func main() {
	if err := internal.Run(); err != nil {
		panic(err)
	}
}
