package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var dbm *sql.DB

//connecting database

func connectDB() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/userinfo?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Connected Sucessfully......")
	dbm = db
}

//creating a table in a database init only runs for the one time

func createTable() {
	query := `Create table logininfo (
		id int auto_increment,
		username text not null,
		email text not null,
		password text not null,
		primary key(id)
	);`
	_, err := dbm.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

// storing the sign up information

func storeSignupInfo() {
	fmt.Printf("\n")
	fmt.Println("User Authentication System ")
	fmt.Printf("\n\n")
	//id := 0
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Username: ")
	username, _ := reader.ReadString('\n')
	reader1 := bufio.NewReader(os.Stdin)
	fmt.Printf("Email: ")
	email, _ := reader1.ReadString('\n')

	reader2 := bufio.NewReader(os.Stdin)
	fmt.Printf("Password: ")
	password1, _ := reader2.ReadString('\n')

	//Encrypting a password

	var letterStorage []string
	var e string
	for _, letter := range password1 {
		a := int(letter)
		if a >= 97 && a <= 122 {
			b := int(letter) + 3
			if b > 122 {
				c := b - 122
				b = 96 + c
				e = string(b)
			}
			e = string(b)
		} else {
			b := int(letter)
			e = string(b)
		}
		letterStorage = append(letterStorage, string(e))
	}
	password := strings.Join(letterStorage, "")

	_, err := dbm.Exec(`insert into logininfo (username,email,password) value (?,?,?)`,
		username, email, password)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Sign Up Sucessfull")
		time.Sleep(2 * time.Second)
		clearScreen()
		menu()
	}

}

// user verification here

func signin() {
A:
	type logininfo struct {
		id       int
		username string
		email    string
		password string
	}
	var s logininfo
	row, err := dbm.Query(`select * from logininfo`)
	if err != nil {
		log.Fatal(err)
	} else {
		for row.Next() {
			row.Scan(&s.id, &s.username, &s.email, &s.password)
			fmt.Printf("\n")
			fmt.Println("User Authentication System")
			fmt.Printf("\n\n")
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("Username / Email: ")
			username, _ := reader.ReadString('\n')
			reader1 := bufio.NewReader(os.Stdin)
			fmt.Printf("Password: ")
			password1, _ := reader1.ReadString('\n')

			// Encrypting the password
			var letterStorage []string
			var e string
			for _, letter := range password1 {
				a := int(letter)
				if a >= 97 && a <= 122 {
					b := int(letter) + 3
					if b > 122 {
						c := b - 122
						b = 96 + c
						e = string(b)
					}
					e = string(b)
				} else {
					b := int(letter)
					e = string(b)
				}
				letterStorage = append(letterStorage, string(e))
			}
			password := strings.Join(letterStorage, "")

			if username == s.username && password == s.password || username == s.email && password == s.password {
				fmt.Println("Sign in Successful!!! ")
				time.Sleep(5 * time.Second)
				menu()

			} else if username != s.email && username != s.username {
				fmt.Println("The Username is not registered!!")
				time.Sleep(1 * time.Second)
				clearScreen()
				menu()
			} else if username == s.email && password != s.password || username == s.username && password != s.password {
				fmt.Println("Invalid Password!! ")
				time.Sleep(1 * time.Second)
				clearScreen()
				goto A
			}

		}
	}

}

// clearing the screen
func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// main menu
func menu() {
	clearScreen()
A:
	fmt.Printf("\n")
	fmt.Println("User Authentication System.")

	fmt.Println("Press 1 to sign up! ")
	fmt.Println("Press 2 to sign in! ")
	fmt.Println("Press 3 to exit! ")
	fmt.Printf("\n\n")

	fmt.Printf("Enter your choice: ")
	var inp int
	fmt.Scanln(&inp)

	if inp == 1 {
		clearScreen()
		storeSignupInfo()
	} else if inp == 2 {
		clearScreen()
		signin()
	} else if inp == 3 {
		os.Exit(1)
	} else {
		fmt.Println("Invalid Input !!!")
		time.Sleep(1 * time.Second)
		goto A
	}
}
func main() {
	clearScreen()
	connectDB()
	//createTable()
	menu()
}
