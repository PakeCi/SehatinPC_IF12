package main

import "fmt"

const NMAX = 9999

type dataComponent struct {
	user, userPassword, serialCode               string
	batteryHealth                                float64 // dalam Persentase
	cpuManufacturer, gpuManufacturer 			 string  // CPU : Intel, AMD, Apple M-series GPU : Nvidia, AMD, Apple or NONE
	cpuModel, gpuModel, cpuSerial, gpuSerial 	 string  // CPU : Core, Pentium, Xeon, Atom, Ryzen, Athlon, Other, GPU : GeForce, Radeon, Apple or NONE
	rataCpuTemp, rataGpuTemp, rataRamTemp        float64 // dalam Celcius
	medCpuTemp, medGpuTemp, medRamTemp           float64 // dalam Celcius
	modCpuTemp, modGpuTemp, modRamTemp           float64 // dalam Celcius
	ramCapacity, ramUsed, diskCapacity, diskUsed float64 // dalam GiB
	dataLoad, dataSudahDiisi                     bool    // true = data dalam heavy load, false = data dalam idle load
	operatingSystem                              string  // Windows, Linux, MacOS
	lastMaintenanceDate, nextMaintenanceDate     string  // Format DD-MM-YYYY
	status                                       string  // Gud, Warning, Critical
}

// Intel Data : https://www.intel.com/content/www/us/en/products/details/processors.html
// AMD Data : https://www.amd.com/en/products/specifications/processors.html
// Apple Data : https://everymac.com/systems/apple/index-apple-specs-applespec.html
//
// beberapa data yang bisa dipake : https://www.darkflash.com/article/safe-cpu-temperature-guide
// (no we i'm making datasheet for all of ts)

type dataBase [NMAX]dataComponent

func header() {
	Line := ("______________________________________________________________________________________________________") //102 karakter
	fmt.Printf("%s\n\n", Line)
	AsciiArt := (`      ::::::::  :::::::::: :::    :::     ::: ::::::::::: ::::::::::: ::::    ::: :::::::::   :::::::: 
    :+:    :+: :+:        :+:    :+:   :+: :+:   :+:         :+:     :+:+:   :+: :+:    :+: :+:    :+: 
   +:+        +:+        +:+    +:+  +:+   +:+  +:+         +:+     :+:+:+  +:+ +:+    +:+ +:+         
  +#++:++#++ +#++:++#   +#++:++#++ +#++:++#++: +#+         +#+     +#+ +:+ +#+ +#++:++#+  +#+          
        +#+ +#+        +#+    +#+ +#+     +#+ +#+         +#+     +#+  +#+#+# +#+        +#+           
#+#    #+# #+#        #+#    #+# #+#     #+# #+#         #+#     #+#   #+#+# #+#        #+#    #+#     
########  ########## ###    ### ###     ### ###     ########### ###    #### ###         ########       `)
	fmt.Println(AsciiArt)
	fmt.Printf("%s\n", Line)
	// source : https://stackoverflow.com/questions/77086187/how-do-i-work-with-ascii-art-strings-in-a-program soalnya gw gatau cara dapet back quotes
}

func footer() {
	Line := ("______________________________________________________________________________________________________")
	fmt.Printf("%s\n", Line)
}

func insertionSortAsc(data *[10]float64) {
	var idx int
	var temp float64
	for i := 1; i < 10; i++ {
		idx = i
		temp = data[i]
		for idx > 0 && data[idx-1] > temp {
			data[idx] = data[idx-1]
			idx = idx - 1
		}
		data[idx] = temp
	}
}

func searchRataRata(data [10]float64) float64 {
	var total float64
	for i := 0; i < 10; i++ {
		total = total + data[i]
	}
	return total / 10
}

func searchMedian(data [10]float64) float64 {
	insertionSortAsc(&data)
	return 0.5 * (data[4] + data[5]) //soalnya start dari 0
}

func searchMinMax(data [10]float64, min, max *float64) {
	insertionSortAsc(&data)
	*min = data[0]
	*max = data[9]
}

func searchModus(data [10]float64) float64 {
	var count, maxCount int
	var modus float64

	for i := 0; i < 10; i++ {
		count = 1
		for y := i + 1; y < 10; y++ {
			if data[i] == data[y] {
				count = count + 1
			}
		}
		if count > maxCount {
			maxCount = count
			modus = data[i]
		}
	}
	return modus
}

func loginPage(data *dataBase, userIndex *int, login *bool, loggedInUser *int) {
	var username, password string
	var valid bool = false
	// var firstOpenPage bool = true

	for !valid {
		// if firstOpenPage {
		// 	header()
		// 	firstOpenPage = false
		// }
		fmt.Printf("\n%-45s%s\n\n", " ", "LOGIN PAGE")
		fmt.Printf("%-44s%s\n%-44s%s\n%-41s%s\n", " ", "Input Username", " ", "Input Password", " ", "Type Exit to go back")
		footer()
		fmt.Print("Input Your Username : ")
		fmt.Scan(&username)
		if username == "Exit" {
			valid = true
		} else {
			fmt.Print("Input Your Password : ")
			fmt.Scan(&password)
			*loggedInUser = validateUserLogin(data, username, password, *userIndex)
			if *loggedInUser != -1 {
				valid = true
				*login = true
			} else {
				fmt.Println("Invalid username or password. Please try again.")
			}
		}
	}
}

func validateUserLogin(data *dataBase, username, password string, userIndex int) int {
	for i := 0; i < userIndex; i++ {
		if data[i].user == username && data[i].userPassword == password {
			return i
		}
	}
	return -1
}

func registerPage(data *dataBase, userIndex *int) {
	var username, password string
	var valid bool = false
	// var firstOpenPage bool = true
	var passwordValid bool = false

	for valid == false {
		// if firstOpenPage {
		// 	header()
		// 	firstOpenPage = false
		// }
		fmt.Printf("\n%-45s%s\n\n", " ", "REGISTER PAGE")
		fmt.Printf("%-47s%s\n%-36s%s\n%-41s%s\n", " ", "Username", " ", "Password at least 8 characters", " ", "Type Exit to go back")
		footer()
		fmt.Print("Username : ")
		fmt.Scan(&username)
		if username == "Exit" {
			valid = true
		} else {
			valid = checkValidityUser(data, username, *userIndex)
			if valid {
				data[*userIndex].user = username
				for !passwordValid {
					fmt.Print("Password : ")
					fmt.Scan(&password)
					if len(password) < 8 {
						fmt.Println("Password is not valid, Password must contain at least 8 characters.")
					} else {
						passwordValid = true
					}
				}
				data[*userIndex].userPassword = password
				*userIndex = *userIndex + 1
			} else {
				fmt.Println("Username already taken, Please try again with different username.")
			}
		}
	}
}

func checkValidityUser(data *dataBase, username string, userIndex int) bool {
	for i := 0; i < userIndex; i++ {
		if data[i].user == username {
			return false
		}
	}
	return true
}

func loginMenu(data *dataBase, kill *bool, login *bool, userIndex *int, firstOpenPage *bool, loggedInUser *int) {
	var input int
	if *firstOpenPage {
		header()
		*firstOpenPage = false
	}
	fmt.Printf("\n%-46s%s\n\n", " ", "LOGIN MENU") // 1/2*102 = 51
	fmt.Printf("%-45s%s\n%-45s%s\n%-45s%s\n", " ", "1. Register", " ", "2. Login", " ", "3. Kill Program")
	footer()
	fmt.Print("Input : ")
	fmt.Scan(&input)

	if input == 1 {
		registerPage(data, userIndex)
	} else if input == 2 {
		loginPage(data, userIndex, login, loggedInUser)
	} else if input == 3 {
		*kill = true
	} else {
		fmt.Println("Input Invalid")
	}
}

func main() {
	var data dataBase
	var exit bool = false
	var login bool = false
	var firstOpenPage bool = true
	var userIndex, loggedInUser int

	data[0].user = "admin"
	data[0].userPassword = "admin"
	userIndex = 1

	for !exit {
		if !login {
			loginMenu(&data, &exit, &login, &userIndex, &firstOpenPage, &loggedInUser)
			if login {
				firstOpenPage = true
			}
		} else {
			// if loggedInUser == 0 {
			// 	adminMenu(&data, &exit, &login, &userIndex, &firstOpenPage)
			// }else {
			// 	mainMenu(&data, &exit, &login, &userIndex, &firstOpenPage, loggedInUser)
			// }
			mainMenu(&data, &exit, &login, &userIndex, &firstOpenPage, loggedInUser)
			if !login {
				firstOpenPage = true
			}
		}
	}
}

func mainMenu(data *dataBase, kill *bool, login *bool, totalUser *int, firstOpenPage *bool, loggedInUser int) {
	var exit bool = false
	var input int

	if *firstOpenPage {
		header()
		*firstOpenPage = false
	}
	for !exit {
		fmt.Printf("\n%-46s%s\n\n", " ", "MAIN MENU")
		if loggedInUser == 0 {
			fmt.Printf("%-43s%s\n%-43s%s\n%-43s%s\n%-43s%s\n%-43s%s\n%-43s%s\n", " ", "1. Show All User", " ", "2. Change User Data", " ", "3. Delete User Data", " ", "4. Show All User Statistics", " ", "5. Logout", " ", "6. Kill Program")
			footer()
			fmt.Print("Input: ")
			fmt.Scan(&input)
			switch input {
			case 1:
				if *totalUser <= 1 {
					fmt.Println("No user data available")
				}else {
					fmt.Printf("\n%-45s%s\n", " ", "ALL USER DATA")
						for i := 1; i < *totalUser; i++ {
							fmt.Printf("\nUser %d: %s\n", i, data[i].user)
							fmt.Println()
						}
					footer()
				}	
			case 2:
				changeDataUser(data, loggedInUser, kill, login, totalUser)
			case 3:
				deleteDataUser(data, loggedInUser, kill, login, totalUser)
			case 4:
				statisticsMenu(data, loggedInUser, kill, login, totalUser)
			case 5:
				*login = false
				exit = true
			case 6:
				*kill = true
				exit = true
			default:
				fmt.Println("Invalid Input")
			}
		} else {
			if data[loggedInUser].dataSudahDiisi {
				fmt.Printf("%-43s%s\n%-43s%s\n%-43s%s\n%-43s%s\n%-43s%s\n%-43s%s\n", " ", "1. Show Your Data", " ", "2. Change Data", " ", "3. Delete Data", " ", "4. Statistics", " ", "5. Logout", " ", "6. Kill Program")
				footer()
				fmt.Print("Input: ")
				fmt.Scan(&input)

				switch input {
				case 1:
					showDataUser(data, loggedInUser, kill, login, totalUser)
				case 2:
					changeDataUser(data, loggedInUser, kill, login, totalUser)
				case 3:
					deleteDataUser(data, loggedInUser, kill, login, totalUser)
				case 4:
					statisticsMenu(data, loggedInUser, kill, login, totalUser)
				case 5:
					*login = false
					exit = true
				case 6:
					*kill = true
					exit = true
				default:
					fmt.Println("Invalid Input")
				}
			} else {
				fmt.Printf("%-44s%s\n%-46s%s\n%-43s%s\n%-23s%s\n", " ", "1. Input Data", " ", "2. Logout", " ", "3. Kill Program", " ", "To access other menu, you must input your PC data first!")
				footer()
				fmt.Print("Input : ")
				fmt.Scan(&input)

				switch input {
				case 1:
					inputDataUser(data, loggedInUser, kill, login)
				case 2:
					*login = false
					exit = true
				case 3:
					*kill = true
					exit = true
				default:
					fmt.Println("Invalid Input")
				}
			}
		}
	}
}

func inputDataUser(data *dataBase, loggedInUser int, kill *bool, login *bool) {
	var valid1 bool = false
	var valid2 bool = false
	var serialNumber, cpuManufacturer, gpuManufacturer, cpuModel, gpuModel, lastMaintenance string
	var cpuSerial, gpuSerial string
	var ramCapacity, diskCapacity float64

	fmt.Printf("RULES : ")
	// nanti kutambahin rulesnya

	footer()
	fmt.Printf("\n%-45s%s\n\n", " ", "INPUT USER DATA (Page 1/2)")
	for !valid1 {
		fmt.Print("Serial Number: ")
		fmt.Scan(&serialNumber)
		valid1 = checkValidityInput(serialNumber, 1)
		if !valid1 {
			fmt.Println("Serial Number cannot be null, or none")
		}
	}
	validManuf := false
	validModel := false

	for !validManuf || !validModel {
		fmt.Print("CPU: ")
		fmt.Scan(&cpuManufacturer, &cpuModel, &cpuSerial)
		validManuf = checkValidityInput(cpuManufacturer, 2)
		validModel = checkValidityInput(cpuModel, 3)
		if !validManuf || !validModel {
			fmt.Println("Invalid Manufacturer or Invalid Model Input")
		}
	}

	validGpuManuf := false
	validGpuModel := false

	for !validGpuManuf || !validGpuModel {
		fmt.Print("GPU: ")
		fmt.Scan(&gpuManufacturer, &gpuModel, &gpuSerial)
		validGpuManuf = checkValidityInput(gpuManufacturer, 4)
		validGpuModel = checkValidityInput(gpuModel, 5)
	}

	valid1 = false
	for !valid1 {
		fmt.Print("Ram Capacity(GiB): ")
		fmt.Scan(&ramCapacity)
		if ramCapacity > 0 {
			valid1 = true
		}else {
			fmt.Println("Insufficient Ram Space")
		}
	}
	valid1 = false
	for !valid1 {
		fmt.Print("Disk Capacity(GiB): ")
		fmt.Scan(&diskCapacity)
		if diskCapacity > 0 {
			valid1 = true
		}else {
			fmt.Println("Insufficient Disk Space")
		}
	}
	valid1 = false
	for !valid1 {
		fmt.Print("Last Maintenance (DD-MM-YYYY): ")
		fmt.Scan(&lastMaintenance)
		valid1 = checkValidityInput(lastMaintenance, 6)
	}
	footer()
	fmt.Printf("\n%-45s%s\n\n", " ", "INPUT USER DATA (Page 2/2)")
	for !valid2 {

	}
}

func checkValidityInput(x string, i int) bool {
	// SerialNumber = 1, cpuManufacture = 2, cpuModel = 3, gpuManufacture = 4, gpuModel = 5, lastMaintenance = 6
	x = upperCaseConverter(x)
	if i == 1 {
		if x == "" || x == "NONE" {
			return false
		}
	}else if i == 2 {
		if x != "INTEL" && x != "AMD" && x != "APPLE" {
			return false
		} 
	}else if i == 3 {
		if x != "CORE" && x != "PENTIUM" && x != "XEON" {
			return false
		}
	}else if i == 4 {
		
	}else if i == 5 {

	}else if i == 6 {

	}
	return true
}

func upperCaseConverter(x string) string {
	var char byte
	var upperCase string
	
	for i := 0; i < len(x); i++ {
		char = x[i]
		if char >= 'a' && char <= 'z' {
			upperCase = upperCase + string(char-32)
		}else {
			upperCase = upperCase + string(char)
		}
	}
	return upperCase
}

// func adminMenu(data *dataBase, kill *bool, login *bool, totalUser *int, firstOpenPage *bool) {
// 	var exit bool = false
// 	var input int

// 	if *firstOpenPage {
// 		header()
// 		*firstOpenPage = false
// 	}
// 	for !exit {
// 		fmt.Printf("\n%-46s%s\n\n", " ", "ADMIN MENU")
// 		fmt.Printf("%-43s%s\n%-43s%s\n%-43s%s\n%-43s%s\n%-43s%s\n", " ", "1. Show All User", " ", "2. Change User Data", " ", "3. Delete User Data", " ", "4. Show All User Statistics", " ", "5. Logout", " ", "6. Kill Program")
// 		footer()
// 		fmt.Print("Input : ")
// 		fmt.Scan(&input)
// 	}
// }

func showDataUser(data *dataBase, loggedInUser int, kill *bool, login *bool, totalUser *int) {
	if loggedInUser == 0 {
		if *totalUser <= 1 {
			fmt.Println("No user data available")
			footer()
		} else {
			for i := 1; i < *totalUser; i++ {
				outputDataUserFormat(data, i)
				fmt.Println()
			}
		}
	} else {
		fmt.Printf("\n%-40s%s\n\n", " ", "YOUR DATA INFORMATION")
		footer()
		outputDataUserFormat(data, loggedInUser)
	}
}

func outputDataUserFormat(data *dataBase, i int) {
	fmt.Printf("GENERAL SPESIFICATIONS\n\n")
	fmt.Printf("User %d: %s\n", i, data[i].user)
	fmt.Println("Serial Code:", data[i].serialCode)
	fmt.Println("CPU: ", data[i].cpuManufacturer, " ", data[i].cpuModel, " ", data[i].cpuSerial)
	fmt.Println("GPU: ", data[i].gpuManufacturer, " ", data[i].gpuModel, " ", data[i].gpuSerial)
	fmt.Println("Battery Health:", data[i].batteryHealth, "%")
	fmt.Println("Operating System:", data[i].operatingSystem)

	footer()
	fmt.Printf("PERFORMANCE SPESIFICATIONS\n\n")
	fmt.Printf("CPU SPEC: \n")
	fmt.Printf("Average CPU Temperature: %.2f°C\n", data[i].rataCpuTemp)
	fmt.Printf("Median CPU Temperature: %.2f°C\n", data[i].medCpuTemp)
	fmt.Printf("Modus CPU Temperature: %.2f°C\n", data[i].modCpuTemp)
	fmt.Printf("\nGPU SPEC: \n")
	fmt.Printf("Average GPU Temperature: %.2f°C\n", data[i].rataGpuTemp)
	fmt.Printf("Median GPU Temperature: %.2f°C\n", data[i].medGpuTemp)
	fmt.Printf("Modus GPU Temperature: %.2f°C\n", data[i].modGpuTemp)
	fmt.Printf("\nRAM SPEC: \n")
	fmt.Printf("RAM Capacity: %.2f GiB\n", data[i].ramCapacity)
	fmt.Printf("RAM Used: %.2f GiB\n", data[i].ramUsed)
	fmt.Printf("Average RAM Temperature: %.2f°C\n", data[i].rataRamTemp)
	fmt.Printf("Median RAM Temperature: %.2f°C\n", data[i].medRamTemp)
	fmt.Printf("Modus RAM Temperature: %.2f°C\n", data[i].modRamTemp)
	fmt.Printf("\nDISK SPEC: \n")
	fmt.Printf("Disk Capacity: %.2f GiB\n", data[i].diskCapacity)
	fmt.Printf("Disk Used: %.2f GiB\n", data[i].diskUsed)

	footer()
	fmt.Printf("MAINTENANCE HISTORY\n\n")
	fmt.Println("Last Maintenance Date:", data[i].lastMaintenanceDate)
	fmt.Println("Next Maintenance Date:", data[i].nextMaintenanceDate)
	fmt.Println("User Status:", data[i].status)
	footer()
}

func deleteDataUser(data *dataBase, loggedInUser int, kill *bool, login *bool, totalUser *int) {

}

func statisticsMenu(data *dataBase, loggedInUser int, kill *bool, login *bool, totalUser *int) {

}

func changeDataUser(data *dataBase, loggedInUser int, kill *bool, login *bool, totalUser *int) {

}

