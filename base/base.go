package main

import (
	"fmt"
	"log"

	"github.com/zabawaba99/firego"
)

type mailer interface {
	send() bool
}

type mailman struct {
	msg        string
	isPhysical bool
}

type postman struct {
	msg string
}

func (m mailman) send() bool {
	if m.isPhysical {
		return false
	}
	log.Println("Mailman Sending: ", m.msg)
	return true
}

func (p postman) send() bool {
	log.Println("Postman Sending: ", p.msg)
	return true
}

func processMailbox(m mailer) {
	m.send()
}

type deck []string

func (deck) String() string {
	return "My deck"
}

type vehicle struct {
	Plate  string `json:"plate"`
	Model  int    `json:"model"`
	Color  string `json:"color"`
	Type   string `json:"type"`
	IsMain bool   `json:"is_main"`
}

type user struct {
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	Vehicle  []vehicle `json:"vehicles"`
	Wfh      string    `json:"wfh"`
	IsActive bool      `json:"is_active"`
	Password string    `json:"password"`
	FreeDays []string  `json:"free_days"`
}

var fbEndPointBase string

func main() {
	fmt.Println("Hello Gopher!")

	fbEndPointBase = "https://vpparking-de51c.firebaseio.com/%s"

	getUsersFromFB()
	m := mailman{"Hi Mailman", false}
	processMailbox(m)
	p := postman{"Hi Postman"}
	processMailbox(p)

	d := deck{"a", "b"}
	fmt.Println(d)
	fmt.Println(d)
}

func getUsersFromFB() {
	fb := firego.New(fmt.Sprintf(fbEndPointBase, "users"), nil)

	var v []user
	if err := fb.Value(&v); err != nil {
		log.Fatal("Error: Cannot get users info", err)
		return
	}
	for _, user := range v {
		fmt.Printf("%s\n", user.Email)
		usrVehicle := user.Vehicle[0]
		vehModel := usrVehicle.Model
		vehIsMain := usrVehicle.IsMain
		fmt.Println(vehModel, vehIsMain)
	}
}
