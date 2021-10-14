package main

//Import packages
import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	return name, roll, course
}

//Main
func main() {

	fmt.Println("Welcome to Attendance Logger!")

	name, roll, course := getStudentInfo()

	file, err := os.Create("attendance.txt")
	if err != nil {
		errorHandler(err)
		log.Fatalf("%s", err)
	}

	defer file.Close()

	_, err2 := file.WriteString(name)
	_, err3 := file.WriteString(roll)
	_, err4 := file.WriteString(course)

	if err2 != nil {
		errorHandler(err2)
		log.Fatalf("%s", err2)
	}

	if err3 != nil {
		errorHandler(err3)
		log.Fatalf("%s", err3)
	}

	if err4 != nil {
		errorHandler(err4)
		log.Fatalf("%s", err4)
	}

}
