package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Funtion for logging errors
func errorHandler(err error) {
	println("Ops, something went wrong:", err)
}

//Function to get attendance

func viewAttendance() {
	println()
	println("This will show you attendance")

	if _, err := os.Stat("attendance.txt"); os.IsNotExist(err) {
		println("No attendance to show")
		println()
	} else {
		file, err := os.Open("attendance.txt")
		if err != nil {
			errorHandler(err)
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			errorHandler(err)
			log.Fatal(err)
		}

	}

}

func resetAttendance() {
	println()
	println("This will reset attendance")

	if _, err := os.Stat("attendance.txt"); os.IsNotExist(err) {
		println("Attendance already clear")
		println()
	} else {
		e := os.Remove("attendance.txt")
		if e != nil {
			errorHandler(e)
			log.Fatal(e)
		} else {
			println("Attendance Cleared")
			println()
		}
	}
}

func getStudentInfo() (normtime, epochtime, name, roll, course string) {
	now := time.Now()
	fmt.Println("Time: ", now.Local(), "\n")
	epoch := now.Unix()
	norm := now.Local()
	epochtime = fmt.Sprint(epoch)
	normtime = fmt.Sprint(norm)
	fmt.Println("Enter the student name:")

	inputReader := bufio.NewReader(os.Stdin)
	name, _ = inputReader.ReadString('\n')

	fmt.Println("Enter the student roll number:")

	inputReader = bufio.NewReader(os.Stdin)
	roll, _ = inputReader.ReadString('\n')

	fmt.Println("Enter the course:")

	inputReader = bufio.NewReader(os.Stdin)
	course, _ = inputReader.ReadString('\n')

	return strings.TrimSpace(normtime), strings.TrimSpace(epochtime), strings.TrimSpace(name), strings.TrimSpace(roll), strings.TrimSpace(course)

}

// Main
func main() {

	// Looping until user enters -1
	for {
		fmt.Println()
		fmt.Println("===========================")
		fmt.Println("Type \n")
		fmt.Println("1 to view attendance")
		fmt.Println("2 to log attendance")
		fmt.Println("3 to reset attendance")
		fmt.Println("-1 to exit the program")
		fmt.Println("===========================")
		fmt.Println()
		var option int

		fmt.Scanln(&option)

		switch option {
		case 1:
			viewAttendance()
		case 2:
			normtime, epochtime, name, roll, course := getStudentInfo()
			record := []string{normtime, epochtime, name, roll, course}

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
		case 3:
			resetAttendance()
		case -1:
			fmt.Println("Exiting the program...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
