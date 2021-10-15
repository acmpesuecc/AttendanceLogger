package main

//Import packages
import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
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

// function to view attendance
func viewStudentInfo() {
	fmt.Println("Viewing student info")
	recordFile, err := os.Open("attendance.csv")
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return
	}

	// Setup the reader
	reader := csv.NewReader(recordFile)

	// Read the records
	header, err := reader.Read()
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return
	}
	fmt.Printf("%v ", header)
	fmt.Println()

	for i := 0; ; i = i + 1 {
		record, err := reader.Read()
		if err == io.EOF {
			break // reached end of the file
		} else if err != nil {
			fmt.Println("An error encountered ::", err)
			return
		}
		fmt.Printf("%v ", record)
		fmt.Println()
	}

}

//function to reset all information on the given csv file
func resetStudentInfo() {
	fmt.Println("Resetting student info")
	if err := os.Truncate("/home/johnjj/Documents/hacknight_forks/AttendanceLogger/attendance.csv", 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}

}

//Main
func main() {
	fmt.Println("Welcome to the Attendance Logger")
	fmt.Println("Enter 1 to log attendance")
	fmt.Println("Enter 2 to view attendance")
	fmt.Println("Enter 3 to reset attendance")

	var input int
	fmt.Scanf("%d", &input)

	if input == 1 {
		name, roll, course := getStudentInfo()
		recordFile, err := os.OpenFile("/home/johnjj/Documents/hacknight_forks/AttendanceLogger/attendance.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			errorHandler(err)
			log.Fatalf("%s", err)
		}
		defer recordFile.Close()

		writer := csv.NewWriter(recordFile)
		var csvData = [][]string{
			{name, roll, course},
		}
		err = writer.WriteAll(csvData)
		if err != nil {
			fmt.Println("Error while writing to the file ::", err)
			return
		}
		fmt.Println("Succesfully added")
		err = recordFile.Close()
		if err != nil {
			fmt.Println("Error while closing the file ::", err)
			return
		}

	} else if input == 2 {
		viewStudentInfo()

	} else if input == 3 {
		resetStudentInfo()
	} else {
		fmt.Println("Invalid input, try again.")
	}

}
