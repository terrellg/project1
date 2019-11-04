package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	mainmenu()

}

//SSH into second computer
// view()
// create()
// sysinfo()

/*
	mainmenu
*/
func mainmenu() {
	var selection string
	fmt.Println(`
 ___  ___  _   _  _____  _    _    __  __  _  _    ____   ___ 
/ __)/ __)( )_( )(  _  )( \/\/ )  (  \/  )( \/ )  (  _ \ / __)
\__ \\__ \ ) _ (  )(_)(  )    (    )    (  \  /    )___/( (__ 
(___/(___/(_) (_)(_____)(__/\__)  (_/\/\_) (__)   (__)   \___)`)

	// fmt.Println("Welcome to the Remote PC Manager")
	fmt.Println("")
	fmt.Println("Main Menu")
	fmt.Println("Press 1 to view all files")
	// fmt.Println("Press 2 to change directories")
	fmt.Println("Press 3 to create a file")
	fmt.Println("Press 4 to delete a file")
	fmt.Println("Press 5 to view running processes")
	fmt.Println("Press 6 to kill running processes")
	fmt.Println("Press 7 to view system info")
	fmt.Println("Press 8 to search for an app")
	fmt.Println("Press 9 to install an app")
	fmt.Scan(&selection)

	switch selection {
	case "1":
		view()

	case "2":
		fmt.Println("Not Yet Implemented")

	case "3":
		create()

	case "4":
		delete()

	case "5":
		listprocesses()

	case "6":
		killprocess()

	case "7":
		sysinfo()

	case "8":
		search()
	case "9":
		install()
	default:
		fmt.Println("Invalid Option")
		mainmenu()
	}

}

/*
	ssh into device, run ls and print the results
*/
func view() {
	listed, _ := exec.Command("ssh", "user2@192.168.56.103", "ls", "-a").Output()
	fmt.Println("")
	fmt.Println("Current Files in directory:")
	fmt.Println(string(listed))
	goback()

}

/*
	ssh into device scan for user input. Create file with given input as the name
*/
func create() {
	fmt.Println("Please enter the name of the file you would like to create")
	var filename string
	var selection string
	fmt.Scan(&filename)
	//SSH and then create file
	exec.Command("ssh", "user2@192.168.56.103", "touch", filename).Run()
	fmt.Println("File Successfully Created!")

	fmt.Println("Press 1 to create another file or Press any other key to return")
	fmt.Println("")
	fmt.Scan(&selection)
	if selection == "1" {
		create()
	} else {
		mainmenu()

	}

}

func delete() {
	var filename string
	fmt.Println("Please enter the name of the file you would like to delete")
	fmt.Scan(&filename)
	exec.Command("ssh", "user2@192.168.56.103", "rm", "-r", "-f", filename)
	goback()

}
func listprocesses() {
	processes, _ := exec.Command("ssh", "user2@192.168.56.103", "top", "-b", "-n", "1").Output()
	fmt.Println("")
	fmt.Println(string(processes))
	var selection string
	fmt.Println("Press 1 to kill a process")
	fmt.Println("Press 2 to go back to the main menu")
	fmt.Scan(&selection)
	if selection == "1" {
		killprocess()
	}
	if selection == "2" {
		mainmenu()
	}

	// goback()

}
func killprocess() {
	var process string
	fmt.Println("Enter the pid of the process you would like to kill")
	fmt.Scan(&process)
	exec.Command("ssh", "user2@192.168.56.103", "kill", process).Run()
	goback()
}

func sysinfo() {
	sysinfo, _ := exec.Command("ssh", "user2@192.168.56.103", "lscpu").Output()
	fmt.Println("")
	fmt.Println(string(sysinfo))
	goback()
}

func search() {
	var app string
	fmt.Println("Type in the name of the program you would like to search for")
	fmt.Scan(&app)
	results, _ := exec.Command("ssh", "user2@192.168.56.103", "apt", "search", app).Output()

	fmt.Println("")
	fmt.Println(string(results))
	var selection string
	fmt.Println("Press 1 to search for another app")
	fmt.Println("Press 2 to install an app ")
	fmt.Println("Press 3 to return to the main menu")
	fmt.Scan(&selection)

	if selection == "1" {
		search()
	}
	if selection == "2" {
		install()
	}
	if selection == "3" {
		mainmenu()
	} else {
		os.Exit(0)
	}
}

func install() {
	var app string

	fmt.Println("Type in the name of the program you would like to install")
	fmt.Println("")
	fmt.Scan(&app)
	fmt.Println("")
	results, _ := exec.Command("ssh", "user2@192.168.56.103", "sudo", "apt", "install", app, "-y").Output()
	// exec.Command("ssh", "user2@192.168.56.103", "sudo", "apt", "install", app, "-y").Run()
	fmt.Println(string(results))
	goback()

}
func goback() {
	fmt.Println("Enter 1 to go back to the main menu")
	var input string
	fmt.Scan(&input)
	if input == "1" {
		mainmenu()
	} else {
		os.Exit(0)
	}
}
