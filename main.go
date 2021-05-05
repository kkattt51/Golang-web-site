package main

import (
	"fmt"; "html/template"; "net/http"
)

type User struct {
	Name string
	Age uint16
	Money int16
	AvgGrades, Happiness float64
	Hobbies []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money " +
		"equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func homePage(w http.ResponseWriter, request *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8, []string{"Football", "Skate", "Dance"}}
	//bob.setNewName("Alex")
	//fmt.Fprintf(w, bob.getAllInfo())
	//fmt.Fprintf(w, `<h1>Main Text</h1>
		//<b>Main Text</b>`)
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}

func contactsPage(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Contacts page!")
}

func handleRequest()  {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)
	http.ListenAndServe(":1234", nil)
}

func main() {
	//var bob User = ....
	//bob := User{name: "Bob", age: 25, money: -50, avgGrades: 4.2, happiness: 0.8}
	handleRequest()
}
