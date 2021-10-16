package main

//Import packages
import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

//Funtion for logging errors
func errorHandler(err error) {
	println("Ops, something went wrong:", err)
}

//Function to get attendance
func getStudentInfo() (name, roll, course string) {
	fmt.Println("Enter the student name:")

	inputReader := bufio.NewReader(os.Stdin)
	name, _ = inputReader.ReadString('\n')

	fmt.Println("Enter the student roll number:")

	inputReader = bufio.NewReader(os.Stdin)
	roll, _ = inputReader.ReadString('\n')

	fmt.Println("Enter the course:")

	inputReader = bufio.NewReader(os.Stdin)
	course, _ = inputReader.ReadString('\n')

	return strings.TrimSpace(name), strings.TrimSpace(roll), strings.TrimSpace(course)
}

//Main
func main() {
	name, roll, course := getStudentInfo()
	record := []string{name, roll, course}
	
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
