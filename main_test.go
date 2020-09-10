package main

import "testing"

func TestHandler(t *testing.T) {
	channel := make(chan Player)

	p1 := Player{"Superman", 2000, 7}
	p2 := Player{"Batman", 1500, 5}

	handle(&p1, &p2, "attack", channel)

	if p2.Health != 1473 {
		t.Errorf("L'attaque de superman sur batman n'a pas fonctionné")
	}
}
