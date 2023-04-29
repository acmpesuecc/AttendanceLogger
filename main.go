package main

//Import packages
import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Function for logging errors
func errorHandler(err error) {
	println("Oops, something went wrong:", err)
}

//Function to get attendance

func viewAttendance() {
	println("This will show you attendance: \t")

	if _, err := os.Stat("attendance.txt"); os.IsNotExist(err) {
		println("No attendance to show\t")
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

// Function to reset attendance

func resetAttendance() {
	println("This will reset attendance\t")

	if _, err := os.Stat("attendance.txt"); os.IsNotExist(err) {
		println("Attendance already clear")
	} else {
		e := os.Remove("attendance.txt")
		if e != nil {
			errorHandler(e)
			log.Fatal(e)
		} else {
			println("\033[1m\033[93mAttendance Cleared\033[1m\033[93m")
		}
	}
}
func logout() (normtime, name, roll, course string) {
	now := time.Now()
	fmt.Println("Time: ", now.Local(), "\t")
	norm := now.Local()
	normtime = fmt.Sprint(norm)
	fmt.Println("\033[1m\033[93mEnter the student name:\033[1m\033[93m")

	inputReader := bufio.NewReader(os.Stdin)
	name, _ = inputReader.ReadString('\n')

	fmt.Println("\033[1m\033[93mEnter the student roll number:\033[1m\033[93m")

	inputReader = bufio.NewReader(os.Stdin)
	roll, _ = inputReader.ReadString('\n')

	time.Sleep(2 * time.Second)
	fmt.Println("\033[1m\033[93m-------\033[1m\033[93mYou have been logged out successfully\033[1m\033[93m-------")
	fmt.Println("\033[1m\033[93mThank You!!\033[1m\033[93m")
	return strings.TrimSpace(normtime), strings.TrimSpace(name), strings.TrimSpace(roll), strings.TrimSpace(course)
}

func getStudentInfo() (normtime, name, roll, course string) {
	now := time.Now()
	fmt.Println("Time: ", now.Local(), "\t")
	norm := now.Local()
	normtime = fmt.Sprint(norm)
	fmt.Println("\033[1m\033[93mEnter the student name:\033[1m\033[93m")

	inputReader := bufio.NewReader(os.Stdin)
	name, _ = inputReader.ReadString('\n')

	fmt.Println("\033[1m\033[93mEnter the student roll number:\033[1m\033[93m")

	inputReader = bufio.NewReader(os.Stdin)
	roll, _ = inputReader.ReadString('\n')

	fmt.Println("\033[1m\033[93mEnter the course:\033[1m\033[93m")

	inputReader = bufio.NewReader(os.Stdin)
	course, _ = inputReader.ReadString('\n')
	fmt.Println("\033[1m\033[93mPlease Wait!!\033[1m\033[93m")
	time.Sleep(1 * time.Second)
	fmt.Println("\033[1m\033[93mData added Successfully!\033[1m\033[93m\t")

	return strings.TrimSpace(normtime), strings.TrimSpace(name), strings.TrimSpace(roll), strings.TrimSpace(course)

}

// Main
func main() {

	fmt.Println("\033[1m\033[93mWhat would you like to do?\033[1m\033[93m\t")
	fmt.Println("\033[1m\033[93mPress 1 to view attendance\033[1m\033[93m\t")
	fmt.Println("\033[1m\033[93mPress 2 to log attendance\033[1m\033[93m\t")
	fmt.Println("\033[1m\033[93mPress 3 to reset attendance\033[1m\033[93m\t")
	fmt.Println("\033[1m\033[93mPress 4 to logout\033[1m\033[93m\t")

	var option int

	fmt.Scanln(&option)

	switch option {
	case 1:
		viewAttendance()
	case 2:
		normtime, name, roll, course := getStudentInfo()
		record := []string{normtime, name, roll, course}

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
	case 4:
		logout()
	}

}
