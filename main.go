// The sql database currently only hold the SRN, Name and Subject.
package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Function to connect to SQL
var db *sql.DB

func connectSQL() {
	var err error
	db, err = sql.Open("mysql", "root:TaNaY6969	@tcp(127.0.0.1:3306)/Student") //use local user pwd

	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MySQL!")
}

// Function for logging errors
func errorHandler(err error) {
	println("Ops, something went wrong:", err)
}

// Function to get student info
func getStudentInfo() (name, subject string) {
	fmt.Println("Enter the student name:")

	inputReader := bufio.NewReader(os.Stdin)
	name, _ = inputReader.ReadString('\n')

	fmt.Println("Enter the subject:")

	inputReader = bufio.NewReader(os.Stdin)
	subject, _ = inputReader.ReadString('\n')

	return strings.TrimSpace(name), strings.TrimSpace(subject)
}

type student struct {
	SRN     string
	Name    string
	Subject string
}

//Retrieve data from the Database

func retrieveStudent(db *sql.DB, name, subject string) (*student, error) {
	var s student
	query := "SELECT SRN, Names FROM student WHERE Names = ? AND Subject = ? LIMIT 1"
	err := db.QueryRow(query, name, subject).Scan(&s.SRN, &s.Name)
	if err != nil {
		return nil, err
	}
	s.Subject = subject
	return &s, nil
}

// Main function
func main() {
	connectSQL()

=======
	var opt string
	for true{
	fmt.Println("Type \n")
	fmt.Println("1 to log attendance")

	var option int

	fmt.Scanln(&option)

	switch option {
	case 1:
		name, subject := getStudentInfo()
		s, err := retrieveStudent(db, name, subject)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("SRN: ", s.SRN)
		fmt.Println("Name: ", s.Name)
		fmt.Println("Subject: ", s.Subject)

		normtime := time.Now().Format("2006-01-02 15:04:05")
		epochtime := fmt.Sprintf("%d", time.Now().Unix())
		record := []string{normtime, epochtime, name, s.SRN, subject}

		file, err := os.OpenFile("attendance.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		w := csv.NewWriter(file)

		defer file.Close()

		w.Write(record)
		w.Flush()
		err = w.Error()

		if err != nil {
			errorHandler(err)
			log.Fatalf("%s", err)
		}
	}
=======
	fmt.Println("Type next to log another student's attendance or exit to quit:")
	fmt.Scanln(&opt)
	if opt == "exit"{
		break
	}
}
  fmt.Println("Thank you")
  }