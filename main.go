package main

import (
	"TP/persist"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Player struct {
	Name   string
	Health int
	Attack int
}

type Data struct {
	Player1 Player
	Player2 Player
}

func handleGame(handleChannel chan Player) {
	for {
		select {
		case player := <-handleChannel:
			if player.Name == "Superman" {
				if err := persist.Save("./player1.tmp", player); err != nil {
					log.Fatalln(err)
				}
			} else {
				if err := persist.Save("./player2.tmp", player); err != nil {
					log.Fatalln(err)
				}
			}
		}
	}
}

func generateHandler(handleChan chan Player) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
		}

		var player1 Player
		if err := persist.Load("./player1.tmp", &player1); err != nil {
			log.Fatalln(err)
		}

		var player2 Player
		if err := persist.Load("./player2.tmp", &player2); err != nil {
			log.Fatalln(err)
		}

		switch r.Method {
		case "GET":

			tmpl, _ := template.ParseFiles("template.html")

			data := Data{Player1: player1, Player2: player2}

			tmpl.Execute(w, data)

		case "POST":
			// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}

			player := r.FormValue("player")
			action := r.FormValue("action")

			fmt.Println(player, action)

			switch player {
			case "superman":
				handle(&player1, &player2, action, handleChan)
			case "batman":
				handle(&player2, &player1, action, handleChan)
			}

			tmpl, _ := template.ParseFiles("template.html")

			data := Data{Player1: player1, Player2: player2}

			tmpl.Execute(w, data)

		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	}
}

func handle(p1 *Player, p2 *Player, action string, handleChan chan Player) {
	handleAction(p1, p2, action)
	handleChan <- *p2
}

func playerAttack(p1 *Player, p2 *Player) {
	p2.Health = p2.Health - p1.Attack
}

func playerHeal(p1 *Player, p2 *Player) {
	p2.Health = p2.Health + p1.Attack
}

func handleAction(p1 *Player, p2 *Player, action string) {
	switch action {
	case "attack":
		playerAttack(p1, p2)
	case "heal":
		playerHeal(p1, p2)
	}
}

func main() {
	var handleChan chan Player
	handleChan = make(chan Player)

	go handleGame(handleChan)

	http.HandleFunc("/", generateHandler(handleChan))

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
