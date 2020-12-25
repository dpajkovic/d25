package main

func part1(cardPub, doorPub, modulo, subject int) (cardLoop, doorLoop, encryptionKey int) {
	foundCard, foundDoor := false, false
	cardLoop, doorLoop = 0, 0
	loop, value := 0, 1
	for !(foundCard && foundDoor) {
		loop++
		value = (value * subject) % modulo
		foundCard = foundCard || (value == cardPub)
		foundDoor = foundDoor || (value == doorPub)
		if value == cardPub {
			cardLoop = loop
		}
		if value == doorPub {
			doorLoop = loop
		}
	}
	if cardLoop < doorLoop {
		subject = doorPub
		loop = cardLoop
	} else {
		subject = cardPub
		loop = doorLoop
	}
	value = 1
	for loop > 0 {
		loop--
		value = (value * subject) % modulo
	}
	encryptionKey = value
	return
}
