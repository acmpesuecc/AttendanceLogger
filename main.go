package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func resetAttendance(collection *mongo.Collection) {
	println("This will reset attendance")

	if _, err := os.Stat("attendance.txt"); os.IsNotExist(err) {
		println("Attendance already clear")
	} else {
		e := os.Remove("attendance.txt")
		if e != nil {
			errorHandler(e)
			log.Fatal(e)
		} else {
			err := collection.Drop(context.TODO())
			if err != nil {
				errorHandler(err)
				log.Fatal(err)
			}
			println("Attendance Collection Dropped")
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
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB", err)

	}
	defer client.Disconnect(context.TODO())
	collection := client.Database("Attendance").Collection("attendance")
	for {
		fmt.Println()
		fmt.Println("===========================")
		fmt.Println("Type \n")
		fmt.Println("1 to view attendance")
		fmt.Println("2 to log attendance")
		fmt.Println("3 to reset attendance")
		fmt.Println("4 to exit")
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
			_, err := collection.InsertOne(context.TODO(), map[string]string{
				"normtime":  normtime,
				"epochtime": epochtime,
				"name":      name,
				"roll":      roll,
				"course":    course,
			})
			if err != nil {
				errorHandler(err)
				log.Fatal(err)
			}
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
			resetAttendance(collection)
		case 4:
			fmt.Println("Exiting the program...")
			return
		case -1:
			fmt.Println("Exiting the program...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
