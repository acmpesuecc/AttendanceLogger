package main

//Import packages
import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

//Funtion for logging errors
func errorHandler(err error) {
	println("Ops, something went wrong:", err)
}

//Function to get view attendance
//Function to get view attendance
func viewAttendance() {
	fmt.Println("This will show you attendance:")

	if _, err := os.Stat("attendance.txt"); os.IsNotExist(err) {
		fmt.Println("No attendance to show")
	} else {
		file, err := os.Open("attendance.txt")
		if err != nil {
			errorHandler(err)
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			fields := strings.Split(line, ",")
			dateTime := fields[0]
			name := fields[2]
			id := fields[3]
			branch := fields[4]
			fmt.Printf("Date: %s | Name: %-15s | ID: %s | Branch: %s\n" , dateTime[:10], name, id, branch)
			// dateTime[:10] extracts only the date part from the dateTime string
		}

		if err := scanner.Err(); err != nil {
			errorHandler(err)
			log.Fatal(err)
		}
	}
}

	
// Function to reset attendance
func resetAttendance(){
	println("This will reset attendance")
	
	if _, err := os.Stat("attendance.txt"); os.IsNotExist(err) {
		println("Attendance already clear")
	} else{
		e := os.Remove("attendance.txt")
		if e != nil {
			errorHandler(e)
			log.Fatal(e)
		} else{
			println("Attendance Cleared")
		}
	}
}
//Function to get attendance information
func getStudentInfo() (normtime, epochtime, name, roll, course string) {
	now := time.Now()
	epoch := now.Unix()
	norm := now.Local()
	epochtime = fmt.Sprint(epoch)
	normtime = fmt.Sprint(norm)

	pattern := `PES[12]UG[0129]{2}[CE][SC]\d{3}`

	// Loop until valid roll number is entered
	for {
		fmt.Println("Enter the student name:")
		inputReader := bufio.NewReader(os.Stdin)
		name, _ = inputReader.ReadString('\n')

		fmt.Println("Enter the student roll number:")
		inputReader = bufio.NewReader(os.Stdin)
		roll, _ = inputReader.ReadString('\n')

		// Validate roll number using regular expression
		matcher := regexp.MustCompile(pattern);
		matched := matcher.MatchString(roll)

		if matched {
			break // Exit loop if valid roll number is entered
		} else {
			fmt.Println("Invalid roll number. Roll number should be of the form", pattern)
		}
	}

	fmt.Println("Enter the course:")
	inputReader := bufio.NewReader(os.Stdin) // Declare inputReader again
	course, _ = inputReader.ReadString('\n')

	return strings.TrimSpace(normtime), strings.TrimSpace(epochtime), strings.TrimSpace(name), strings.TrimSpace(roll), strings.TrimSpace(course)
}


//Main
func main() {
	
	fmt.Println("Type: \n")
	fmt.Println("'1' to View Attendance")
	fmt.Println("'2' to Log Attendance")
	fmt.Println("'3' to Reset Attendance")
	
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
	}
	
	
	
}
