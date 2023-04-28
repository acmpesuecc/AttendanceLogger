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

// Funtion for logging errors
func errorHandler(err error) {
	println("Ops, something went wrong:", err)
}

//Function to get attendance

func viewAttendance() {
	println("This will show you attendance")

	fmt.Println("Which class's attendance would you like to view?:")

	Reader := bufio.NewReader(os.Stdin)
	classname1, _ := Reader.ReadString('\n')

	if _, err := os.Stat(classname1); os.IsNotExist(err) {
		println("No attendance to show")
	} else {
		file, err := os.Open(classname1)
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
	println("This will reset attendance")

	fmt.Println("Which class's attendance do you wish to reset:")

	inputReader := bufio.NewReader(os.Stdin)
	classname, _ := inputReader.ReadString('\n')

	if _, err := os.Stat(classname); os.IsNotExist(err) {
		println("Attendance already clear")
	} else {
		e := os.Remove(classname)
		if e != nil {
			errorHandler(e)
			log.Fatal(e)
		} else {
			println("Attendance Cleared")
		}
	}
}

func getStudentInfo() (normtime, epochtime, name, class, roll, course string) {
	now := time.Now()
	fmt.Println("Time: ", now.Local(), "\n")
	epoch := now.Unix()
	norm := now.Local()
	epochtime = fmt.Sprint(epoch)
	normtime = fmt.Sprint(norm)
	fmt.Println("Enter the student name:")

	inputReader := bufio.NewReader(os.Stdin)
	name, _ = inputReader.ReadString('\n')

	fmt.Println("Enter the Student class/section:")

	inputReader = bufio.NewReader(os.Stdin)
	class, _ = inputReader.ReadString('\n')

	fmt.Println("Enter the student roll number:")

	inputReader = bufio.NewReader(os.Stdin)
	roll, _ = inputReader.ReadString('\n')

	fmt.Println("Enter the course:")

	inputReader = bufio.NewReader(os.Stdin)
	course, _ = inputReader.ReadString('\n')

	return strings.TrimSpace(normtime), strings.TrimSpace(epochtime), strings.TrimSpace(name), strings.TrimSpace(class), strings.TrimSpace(roll), strings.TrimSpace(course)

}

// Main
func main() {

	fmt.Println("Type \n")
	fmt.Println("1 to view attendance")
	fmt.Println("2 to log attendance")
	fmt.Println("3 to reset attendance")

	var option int

	fmt.Scanln(&option)

	switch option {
	case 1:
		viewAttendance()
	case 2:
		normtime, epochtime, name, class, roll, course := getStudentInfo()
		record := []string{normtime, epochtime, name, class, roll, course}

		filename := class + "attendance.txt"

		file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
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
	}

}
