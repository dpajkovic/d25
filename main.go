//	Copyright (c) Miloš Rackov 2020
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package main

import "fmt"

func main() {
	subject := 7
	modulo := 20201227
	cardPub, doorPub := 15335876, 15086442
	cardLoop, doorLoop, encryptionKey := part1(cardPub, doorPub, modulo, subject)
	fmt.Printf("cardLoop %v\ndoorLoop %v\nencryptionKey %v\n", cardLoop, doorLoop, encryptionKey)
}
