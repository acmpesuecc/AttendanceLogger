package main
import "regexp"

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func errorHandler(err error) {
	fmt.Println("Ops, something went wrong:", err)
}

func viewAttendance() {
	fmt.Println("This will show you attendance")
	
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
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			errorHandler(err)
			log.Fatal(err)
		}

	}
}	

func resetAttendance(){
	fmt.Println("This will reset attendance")
	
	if _, err := os.Stat("attendance.txt"); os.IsNotExist(err) {
		fmt.Println("Attendance already clear")
	} else {
		err := os.Remove("attendance.txt")
		if err != nil {
			errorHandler(err)
			log.Fatal(err)
		} else {
			fmt.Println("Attendance Cleared")
		}
	}
}



//Function to get student info with roll format validation
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

    //Regular expression for roll format validation
    rollFormat := regexp.MustCompile(`^[a-zA-Z]{3}\d{1}[a-zA-Z]{2}\d{2}[a-zA-Z]{2}\d{3}$`)
    
    //Loop until a valid roll format is entered
    for {
        fmt.Println("Enter the student SRN:")
        inputReader = bufio.NewReader(os.Stdin)
        roll, _ = inputReader.ReadString('\n')
        roll = strings.TrimSpace(roll)
        
        if rollFormat.MatchString(roll) {
            break
        } else {
            fmt.Println("Invalid SRN format. Please try again.")
        }
    }

    fmt.Println("Enter the course:")

    inputReader = bufio.NewReader(os.Stdin)
    course, _ = inputReader.ReadString('\n')

    return strings.TrimSpace(normtime), strings.TrimSpace(epochtime), strings.TrimSpace(name), strings.TrimSpace(roll), strings.TrimSpace(course)
}



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
		normtime, epochtime, name, roll, course := getStudentInfo()
		record := []string{normtime, epochtime, name, roll, course}

		file, err := os.OpenFile("attendance.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}
	
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()
		
		if err := writer.Write(record); err != nil {
			errorHandler(err)
			log.Fatalf("%s", err)
		}
	case 3:
		resetAttendance()
	}
}
