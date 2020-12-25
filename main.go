package main

import "fmt"

func main() {
	subject := 7
	modulo := 20201227
	cardPub, doorPub := 15335876, 15086442
	cardLoop, doorLoop, encryptionKey := part1(cardPub, doorPub, modulo, subject)
	fmt.Printf("cardLoop %v\ndoorLoop %v\nencryptionKey %v\n", cardLoop, doorLoop, encryptionKey)
}
