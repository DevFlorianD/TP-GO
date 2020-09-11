package main

import (
	"testing"
)

func TestHandlerAttack(t *testing.T) {
	channel := make(chan Player)
	exitChannel := make(chan bool)

	go handleGame(channel, exitChannel)

	playerTest1 := Player{"Superman", 2000, 7}
	playerTest2 := Player{"Batman", 1500, 5}

	handle(&playerTest1, &playerTest2, "attack", channel)

	if playerTest2.Health != 1493 {
		t.Errorf("L'attaque de superman sur batman n'a pas fonctionné")
	}
}

func TestHandlerHeal(t *testing.T) {
	channel := make(chan Player)
	exitChannel := make(chan bool)

	go handleGame(channel, exitChannel)

	playerTest1 := Player{"Superman", 2000, 7}
	playerTest2 := Player{"Batman", 1500, 5}

	handle(&playerTest1, &playerTest2, "heal", channel)

	if playerTest2.Health != 1507 {
		t.Errorf("Le soin de superman sur batman n'a pas fonctionné")
	}
}
