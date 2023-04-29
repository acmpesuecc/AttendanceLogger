package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Function for logging errors
func errorHandler(err error) {
	println("Ops, something went wrong:", err)
}

// Function to get attendance
func viewAttendance(collection *mongo.Collection) {
	fmt.Println()
	fmt.Println("This will show you attendance")

	cur, err := collection.Find(context.Background(), nil)
	if err != nil {
		errorHandler(err)
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	if cur.Err() != nil {
		errorHandler(err)
		log.Fatal(err)
	}

	if cur.RemainingBatchLength() == 0 {
		fmt.Println("No attendance to show")
		fmt.Println()
		return
	}

	for cur.Next(context.Background()) {
		var result map[string]string
		err := cur.Decode(&result)
		if err != nil {
			errorHandler(err)
			log.Fatal(err)
		}
		fmt.Println(result)
	}

	if err := cur.Err(); err != nil {
		errorHandler(err)
		log.Fatal(err)
	}
}

func resetAttendance(collection *mongo.Collection) {
	fmt.Println("This will reset attendance")
	err := collection.Drop(context.Background())
	if err != nil {
		errorHandler(err)
		log.Fatal(err)
	}
	fmt.Println("Attendance Collection Dropped")
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
func displayCourseInfo(collection *mongo.Collection) {
	fmt.Println("Enter the course name:")
	inputReader := bufio.NewReader(os.Stdin)
	courseName, _ := inputReader.ReadString('\n')
	courseName = strings.TrimSpace(courseName)

	filter := bson.M{"course": courseName}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		errorHandler(err)
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	var students []bson.M
	if err := cur.All(context.Background(), &students); err != nil {
		errorHandler(err)
		log.Fatal(err)
	}

	fmt.Println("Total students in the course:", len(students))
	for _, student := range students {
		fmt.Println(student["name"])
	}
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
			// View attendance from MongoDB
			cur, err := collection.Find(context.Background(), bson.M{})
			if err != nil {
				errorHandler(err)
				log.Fatal(err)
			}
			defer cur.Close(context.Background())

			for cur.Next(context.Background()) {
				var result bson.M
				err := cur.Decode(&result)
				if err != nil {
					errorHandler(err)
					log.Fatal(err)
				}
				fmt.Println(result["normtime"], result["name"], result["roll"], result["course"])
			}
			if err := cur.Err(); err != nil {
				errorHandler(err)
				log.Fatal(err)
			}
		case 2:
			// Log attendance to MongoDB
			normtime, epochtime, name, roll, course := getStudentInfo()

			res, err := collection.InsertOne(context.TODO(), bson.M{
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
			fmt.Println("Attendance added successfully with ID:", res.InsertedID)
		case 3:
			// Reset attendance collection
			err := collection.Drop(context.Background())
			if err != nil {
				errorHandler(err)
				log.Fatal(err)
			}
			fmt.Println("Attendance Collection Dropped")
		case 4:
			displayCourseInfo(collection)

		case 5:
			fmt.Println("Exited the program...")
			return
		case -1:
			fmt.Println("Exited the program...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
