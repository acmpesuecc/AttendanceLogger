package main

//Import packages
import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Funtion for logging errors
func errorHandler(err error) {
	println("Ops, something went wrong:", err)
}

// Main
func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Attendance Logger")
	myWindow.Resize(fyne.NewSize(350, 300))

	//View Attendance Button
	content := widget.NewButton("View attendance", func() {

		//Creates a new window for section entry

		w2 := myApp.NewWindow("View Attendance")
		w2.Resize(fyne.NewSize(310, 90))
		class_sec := widget.NewEntry()
		class_sec.SetPlaceHolder("Enter section-course (format)")
		class_sec.Resize(fyne.NewSize(100, 200))

		content := container.NewVBox(class_sec, widget.NewButton("Check", func() {

			file := class_sec.Text + "attendance.txt"
			if _, err := os.Stat(file); os.IsNotExist(err) {
				println("No attendance to show")
			} else {
				file, err := os.Open(file)
				if err != nil {
					errorHandler(err)
					log.Fatal(err)
				}
				defer file.Close()

				scanner := bufio.NewScanner(file)
				myList := []string{}
				for scanner.Scan() {
					myList = append(myList, scanner.Text())
				}

				//Shows the log in the window as a list
				myL := widget.NewList(func() int { return len(myList) },

					func() fyne.CanvasObject { return widget.NewLabel("My new list") },
					func(lii widget.ListItemID, co fyne.CanvasObject) {
						co.(*widget.Label).SetText(myList[lii])
					})

				w2.Resize(fyne.NewSize(400, 300))
				w2.SetContent(myL)
				if err := scanner.Err(); err != nil {
					errorHandler(err)
					log.Fatal(err)
				}

				class_sec.SetText("")

			}
		}))

		content.Resize(fyne.NewSize(300, 200))

		w2.SetContent(

			container.NewWithoutLayout(content),
		)

		w2.Show()

	})
	content.Resize(fyne.NewSize(300, 50))
	content.Move(fyne.NewPos(20, 50))

	//Log Attendance Button

	content2 := widget.NewButton("Log attendance", func() {

		//Creates a new window to log attendance

		w3 := myApp.NewWindow("Log Attendance")
		w3.Resize(fyne.NewSize(310, 90))

		name := widget.NewEntry()
		name.SetPlaceHolder("Enter the name of the student")
		name.Resize(fyne.NewSize(400, 200))

		section := widget.NewEntry()
		section.SetPlaceHolder("Enter the section of the student")
		section.Resize(fyne.NewSize(400, 200))

		roll := widget.NewEntry()
		roll.SetPlaceHolder("Enter the roll no of the student")
		roll.Resize(fyne.NewSize(400, 200))

		course := widget.NewEntry()
		course.SetPlaceHolder("Enter the course which was attended")
		course.Resize(fyne.NewSize(400, 200))

		now := time.Now()
		epoch := now.Unix()
		norm := now.Local()
		epochtime := fmt.Sprint(epoch)
		normtime := fmt.Sprint(norm)

		content_input2 := container.NewVBox(name, section, roll, course, widget.NewButton("Save", func() {

			record := []string{normtime, epochtime, name.Text, section.Text, roll.Text, course.Text}
			filename := section.Text + "-" + course.Text + "attendance.txt"

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

			name.SetText("")
			section.SetText("")
			roll.SetText("")
			course.SetText("")

		}))

		content_input2.Resize(fyne.NewSize(300, 200))

		w3.SetContent(

			container.NewWithoutLayout(content_input2),
		)

		w3.Show()

	})

	content2.Resize(fyne.NewSize(300, 50))
	content2.Move(fyne.NewPos(20, 110))

	//Reset Attendance Button

	content3 := widget.NewButton("Reset attendance", func() {

		w4 := myApp.NewWindow("View Attendance")
		w4.Resize(fyne.NewSize(310, 90))
		Class_Section := widget.NewEntry()
		Class_Section.SetPlaceHolder("Enter section-course (format)")
		Class_Section.Resize(fyne.NewSize(100, 200))

		content4 := container.NewVBox(Class_Section, widget.NewButton("Delete", func() {

			filename1 := Class_Section.Text + "attendance.txt"
			if _, err := os.Stat(filename1); os.IsNotExist(err) {
				println("Attendance already clear")
			} else {
				e := os.Remove(filename1)
				if e != nil {
					errorHandler(e)
					log.Fatal(e)
				} else {
					println("Attendance Cleared")
				}
			}

			Class_Section.SetText("")

		}))

		content4.Resize(fyne.NewSize(300, 200))

		w4.SetContent(container.NewWithoutLayout(content4))

		w4.Show()

	})

	content3.Resize(fyne.NewSize(300, 50))
	content3.Move(fyne.NewPos(20, 170))

	myWindow.SetContent(

		container.NewWithoutLayout(
			content,
			content2,
			content3,
		),
	)
	myWindow.ShowAndRun()

}
