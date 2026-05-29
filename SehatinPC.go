package main

import "fmt"

const NMAX = 10000

type dataComponent struct {
	user, userPassword, serialCode                                         string
	batteryHealth                                                          float64 // dalam Persentase
	cpuManufacturer, gpuManufacturer                                       string  // CPU : Intel, AMD, Apple M-series GPU : Nvidia, AMD, Apple or NONE
	cpuModel, gpuModel, cpuSerial, gpuSerial                               string  // CPU : Core, Pentium, Xeon, Atom, Ryzen, Athlon, Other, GPU : GeForce, Radeon, Apple or NONE
	rataCpuTemp, rataGpuTemp, rataRamTemp                                  float64 // dalam Celcius
	medCpuTemp, medGpuTemp, medRamTemp                                     float64 // dalam Celcius
	modCpuTemp, modGpuTemp, modRamTemp                                     float64 // dalam Celcius
	minCpuTemp, maxCpuTemp, minGpuTemp, maxGpuTemp, minRamTemp, maxRamTemp float64
	ramCapacity, ramUsed, diskCapacity, diskUsed                           float64 // dalam GiB
	dataLoad, dataSudahDiisi                                               bool    // true = data dalam heavy load, false = data dalam idle load
	operatingSystem                                                        string  // Windows, Linux, MacOS
	lastMaintenanceDate, nextMaintenanceDate                               string  // Format DD-MM-YYYY
	status                                                                 string  // Gud, Warning, Critical
	usingLaptop                                                            bool
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
		if count >= maxCount {
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
		fmt.Printf("\n%-46s%s\n\n", " ", "LOGIN PAGE")
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
				} else {
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
	var ramCapacity, diskCapacity, batteryHealth float64
	var cpuTemperature [10]float64
	var gpuTemperature [10]float64
	var ramTemperature [10]float64
	var isLaptop bool
	var usingLaptop string

	fmt.Printf("RULES : \n")
	// nanti kutambahin rulesnya

	footer()
	fmt.Printf("\n%-45s%s\n\n", " ", "INPUT USER DATA (Page 1/2)")
	for !valid1 {
		fmt.Print("Are you using Laptop? (Yes/No): ")
		fmt.Scan(&usingLaptop)
		valid1 = checkValidityInput(usingLaptop, 8, "")
		if !valid1 {
			fmt.Println("Invalid Input")
		} else {
			usingLaptop = upperCaseConverter(usingLaptop)
			if usingLaptop == "YES" {
				isLaptop = true
			} else {
				isLaptop = false
			}
		}
	}
	valid1 = false
	for !valid1 {
		fmt.Print("Serial Number: ")
		fmt.Scan(&serialNumber)
		valid1 = checkValidityInput(serialNumber, 1, "")
		if !valid1 {
			fmt.Println("Serial Number cannot be null, or none")
		}
	}
	validManuf := false
	validModel := false
	for !validManuf || !validModel {
		fmt.Print("CPU: ")
		fmt.Scan(&cpuManufacturer, &cpuModel, &cpuSerial)
		validManuf = checkValidityInput(cpuManufacturer, 2, "")
		validModel = checkValidityInput(cpuModel, 3, cpuManufacturer)
		if !validManuf || !validModel {
			fmt.Println("Invalid Manufacturer or Invalid Model Input")
		}
	}

	validGpuManuf := false
	validGpuModel := false
	for !validGpuManuf || !validGpuModel {
		fmt.Print("GPU: ")
		fmt.Scan(&gpuManufacturer)
		validGpuManuf = checkValidityInput(gpuManufacturer, 4, "")
		x := upperCaseConverter(gpuManufacturer)
		if validGpuManuf && (x != "NONE") {
			fmt.Scan(&gpuModel, &gpuSerial)
			validGpuModel = checkValidityInput(gpuModel, 5, gpuManufacturer)
		} else if validGpuManuf && (x == "NONE") {
			validGpuModel = true
		}
	}
	valid1 = false
	for !valid1 {
		fmt.Print("Ram Capacity(GiB): ")
		fmt.Scan(&ramCapacity)
		if ramCapacity > 0 {
			valid1 = true
		} else {
			fmt.Println("Insufficient Ram Space")
		}
	}
	valid1 = false
	for !valid1 {
		fmt.Print("Disk Capacity(GiB): ")
		fmt.Scan(&diskCapacity)
		if diskCapacity > 0 {
			valid1 = true
		} else {
			fmt.Println("Insufficient Disk Space")
		}
	}
	valid1 = false
	for !valid1 {
		fmt.Print("Last Maintenance (DD-MM-YYYY): ")
		fmt.Scan(&lastMaintenance)
		valid1 = checkValidityInput(lastMaintenance, 6, "")
	}
	footer()
	fmt.Printf("RULES : \n")
	fmt.Printf("\n%-38s%s\n\n", " ", "INPUT USER DATA (Page 2/2)")
	var i int
	fmt.Print("Input 10 of your current CPU Temperature in the last 20 seconds: ")
	for i = 0; i < 10; i++ {
		fmt.Scan(&cpuTemperature[i])
	}
	if gpuManufacturer != "NONE" {
		fmt.Print("Input 10 of your current GPU Temperature in the last 20 seconds: ")
		for i = 0; i < 10; i++ {
			fmt.Scan(&gpuTemperature[i])
		}
	}

	fmt.Print("Input 10 of your current RAM Temperature in the last 20 seconds: ")
	for i = 0; i < 10; i++ {
		fmt.Scan(&ramTemperature[i])
	}
	var load, OS string
	var heavyLoad bool
	var ramUsage, diskUsage float64
	for !valid2 {
		fmt.Print("Current RAM Usage (GiB): ")
		fmt.Scan(&ramUsage)
		if ramUsage <= ramCapacity {
			valid2 = true
		} else {
			fmt.Println("Your RAM usage cannot exceed your RAM capacity.")
		}
	}
	valid2 = false
	for !valid2 {
		fmt.Print("Current Disk Usage (GiB): ")
		fmt.Scan(&diskUsage)
		if diskUsage <= diskCapacity {
			valid2 = true
		} else {
			fmt.Println("Your Disk usage cannot exceed your Disk capacity.")
		}
	}
	valid2 = false
	for !valid2 {
		fmt.Print("Current Operating System: ")
		fmt.Scan(&OS)
		valid2 = checkValidityInput(OS, 7, " ")
		if !valid2 {
			fmt.Println("Invalid Operating System")
		}
	}
	valid2 = false
	if isLaptop {
		for !valid2 {
			fmt.Print("Battery Health: ")
			fmt.Scan(&batteryHealth)
			if batteryHealth >= 0 && batteryHealth <= 100 {
				valid2 = true
			}
		}
		valid2 = false
	} else {
		batteryHealth = -1
	}
	for !valid2 {
		fmt.Print("Is it under Heavy Load? (Yes/No): ")
		fmt.Scan(&load)
		load = upperCaseConverter(load)
		if load == "YES" {
			heavyLoad = true
			valid2 = true
		} else if load == "NO" {
			heavyLoad = false
			valid2 = true
		} else {
			fmt.Println("Input must be (Yes / No)")
		}
	}
	valid2 = false
	footer()
	var input int
	for !valid2 {
		fmt.Print("Type 1 to Save, Type 2 to Exit Menu, Type 3 to Logout, Type 4 to Kill Program")
		fmt.Scan(&input)
		switch input {
		case 1:
			var avg, med, mod, min, max float64
			data[loggedInUser].serialCode = serialNumber
			data[loggedInUser].batteryHealth = batteryHealth
			data[loggedInUser].cpuManufacturer = cpuManufacturer
			data[loggedInUser].gpuManufacturer = gpuManufacturer
			data[loggedInUser].cpuModel = cpuModel
			data[loggedInUser].gpuModel = gpuModel
			data[loggedInUser].cpuSerial = cpuSerial
			data[loggedInUser].gpuSerial = gpuSerial
			data[loggedInUser].ramCapacity = ramCapacity
			data[loggedInUser].diskCapacity = diskCapacity
			data[loggedInUser].lastMaintenanceDate = lastMaintenance
			data[loggedInUser].usingLaptop = isLaptop
			data[loggedInUser].operatingSystem = OS
			data[loggedInUser].dataLoad = heavyLoad
			data[loggedInUser].dataSudahDiisi = true
			processData(&min, &max, &med, &mod, &avg, cpuTemperature)
			data[loggedInUser].rataCpuTemp = avg
			data[loggedInUser].medCpuTemp = med
			data[loggedInUser].modCpuTemp = mod
			data[loggedInUser].minCpuTemp = min
			data[loggedInUser].maxCpuTemp = max
			processData(&min, &max, &med, &mod, &avg, gpuTemperature)
			data[loggedInUser].rataGpuTemp = avg
			data[loggedInUser].medGpuTemp = med
			data[loggedInUser].modGpuTemp = mod
			data[loggedInUser].minGpuTemp = min
			data[loggedInUser].maxGpuTemp = max
			processData(&min, &max, &med, &mod, &avg, ramTemperature)
			data[loggedInUser].rataRamTemp = avg
			data[loggedInUser].medRamTemp = med
			data[loggedInUser].modRamTemp = mod
			data[loggedInUser].minRamTemp = min
			data[loggedInUser].maxRamTemp = max
			data[loggedInUser].ramUsed = ramUsage
			data[loggedInUser].diskUsed = diskUsage
			setData(data, loggedInUser)

			valid2 = true
		case 2:
			valid2 = true
		case 3:
			*login = false
			valid2 = true
		case 4:
			*kill = true
			valid2 = true
		}
	}
}

func processData(min, max, med, mod, avg *float64, data [10]float64) {
	searchMinMax(data, min, max)
	*avg = searchRataRata(data)
	*med = searchMedian(data)
	*mod = searchModus(data)
}

func checkValidityInput(x string, i int, y string) bool {
	// SerialNumber = 1, cpuManufacture = 2, cpuModel = 3, gpuManufacture = 4, gpuModel = 5, lastMaintenance = 6, operatingSystem = 7
	x = upperCaseConverter(x)
	if i == 1 {
		if x == "" || x == "NONE" {
			return false
		}
	} else if i == 2 {
		if x != "INTEL" && x != "AMD" && x != "APPLE" {
			return false
		}
	} else if i == 3 {
		y = upperCaseConverter(y)
		if y == "INTEL" {
			if x != "CORE" && x != "PENTIUM" && x != "XEON" && x != "ATOM" {
				return false
			}
		} else if y == "AMD" {
			if x != "RYZEN" && x != "EPYC" && x != "ATHLON" {
				return false
			}
		} else if y == "APPLE" {
			if x[0] != 'M' {
				return false
			}
		}
	} else if i == 4 {
		if x != "NVIDIA" && x != "AMD" && x != "APPLE" && x != "NONE" {
			return false
		}
	} else if i == 5 {
		y = upperCaseConverter(y)
		if y == "NVIDIA" {
			if x != "RTX" && x != "GTX" && x != "MAX-Q" {
				return false
			}
		} else if y == "AMD" {
			if x != "RADEON" {
				return false
			}
		} else if y == "APPLE" {
			if x[0] != 'M' {
				return false
			}
		}
	} else if i == 6 {
		//formatnya DD-MM-YYYY
		if len(x) != 10 && (x[0] < '0' || x[0] > '3') && (x[1] < '0' || x[1] > '9') && (x[3] < '0' || x[3] > '1') && (x[4] < '0' || x[4] > '9') && (x[6] < '0' || x[6] > '9') && (x[7] < '0' || x[7] > '9') && (x[8] < '0' || x[8] > '9') && (x[9] < '0' || x[9] > '9') && x[2] != '-' && x[5] != '-' {
			return false
		}
	} else if i == 7 {
		if x != "WINDOWS" && x != "LINUX" && x != "MACOS" {
			return false
		}
	} else if i == 8 {
		if x != "YES" && x != "NO" {
			return false
		}
	}
	return true
}

func upperCaseConverter(x string) string {
	var char byte
	var upperCase string

	for i := 0; i < len(x); i++ {
		char = x[i]
		if char >= 'a' && char <= 'z' {
			upperCase = upperCase + string(char-('a'-'A'))
		} else {
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
	fmt.Printf("%-40s\n\n", "GENERAL SPESIFICATIONS")
	fmt.Printf("%s %d%-19s%s %s\n", "User", i, "", ":", data[i].user)
	fmt.Printf("%-25s %s %s\n", "Serial Code", ":", data[i].serialCode)
	fmt.Printf("%-25s %s %s %s %s\n", "CPU", ":", data[i].cpuManufacturer, data[i].cpuModel, data[i].cpuSerial)

	if data[i].gpuManufacturer != "NONE" {
		fmt.Printf("%-25s %s %s %s %s\n", "GPU", ":", data[i].gpuManufacturer, data[i].gpuModel, data[i].gpuSerial)
	}

	fmt.Printf("%-25s %s %.2f%%\n", "Battery Health", ":", data[i].batteryHealth)
	fmt.Printf("%-25s %s %s\n", "Operating System", ":", data[i].operatingSystem)

	footer()
	fmt.Printf("%-25s\n\n", "PERFORMANCE SPESIFICATIONS")
	fmt.Printf("CPU SPEC: \n")
	fmt.Printf("%-25s %s %.2f°C\n", "Average CPU Temperature", ":", data[i].rataCpuTemp)
	fmt.Printf("%-25s %s %.2f°C\n", "Median CPU Temperature", ":", data[i].medCpuTemp)
	fmt.Printf("%-25s %s %.2f°C\n", "Modus CPU Temperature", ":", data[i].modCpuTemp)

	if data[i].gpuManufacturer != "NONE" {
		fmt.Printf("\nGPU SPEC: \n")
		fmt.Printf("%-25s %s %.2f°C\n", "Average GPU Temperature", ":", data[i].rataGpuTemp)
		fmt.Printf("%-25s %s %.2f°C\n", "Median GPU Temperature", ":", data[i].medGpuTemp)
		fmt.Printf("%-25s %s %.2f°C\n", "Modus GPU Temperature", ":", data[i].modGpuTemp)
	}

	fmt.Printf("\nRAM SPEC: \n")
	fmt.Printf("%-25s %s %.2f GiB\n", "RAM Capacity", ":", data[i].ramCapacity)
	fmt.Printf("%-25s %s %.2f GiB\n", "RAM Used", ":", data[i].ramUsed)
	fmt.Printf("%-25s %s %.2f°C\n", "Average RAM Temperature", ":", data[i].rataRamTemp)
	fmt.Printf("%-25s %s %.2f°C\n", "Median RAM Temperature", ":", data[i].medRamTemp)
	fmt.Printf("%-25s %s %.2f°C\n", "Modus RAM Temperature", ":", data[i].modRamTemp)

	fmt.Printf("\nDISK SPEC: \n")
	fmt.Printf("%-25s %s %.2f GiB\n", "Disk Capacity", ":", data[i].diskCapacity)
	fmt.Printf("%-25s %s %.2f GiB\n", "Disk Used", ":", data[i].diskUsed)

	footer()
	fmt.Printf("MAINTENANCE HISTORY\n\n")
	fmt.Printf("%-25s %s %v\n", "Last Maintenance Date", ":", data[i].lastMaintenanceDate)
	fmt.Printf("%-25s %s %v\n", "Next Maintenance Date", ":", data[i].nextMaintenanceDate)
	fmt.Printf("%-25s %s %v\n", "User Status", ":", data[i].status)
	footer()
}

func binarySearch(data *dataBase, searchData float64, totalUser *int, id int) int {
	var right, left, middle int
	// var dataCopy dataBase

	selectionSort(data, *totalUser, id)

	right = *totalUser - 1
	left = 0
	for left < right {
		middle = (right + left) / 2
		if data[middle].indexFloat(id) > searchData {
			right = middle - 1
		} else if data[middle].indexFloat(id) < searchData {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}

func selectionSort(data *dataBase, totalUser, id int) {
	var idx int
	var temp dataComponent

	for i := 1; i < totalUser; i++ {
		idx = i
		temp = data[i]
		for y := i + 1; y < totalUser; y++ {
			if data[y].indexFloat(id) < data[idx].indexFloat(id) {
				idx = y
			}
		}
		data[i] = data[idx]
		data[idx] = temp
	}
}

// https://go.dev/tour/methods/4 pake receiver biar ga banyak if else
func (x dataComponent) indexFloat(idx int) float64 {
	//id yang dicari : 1. rataCpuTemp, 2. rataGpuTemp, 3. rataRamTemp, 4. medCpuTemp, 5. medGpuTemp, 6. medRamTemp
	// 7. modCpuTemp, 8. modGpuTemp, 9. modRamTemp, 10. minCpuTemp, 11. minGpuTemp, 12. minRamTemp,
	// 13. maxCpuTemp, 14. maxGpuTemp, 15. maxRamTemp, 16. ramCapacity, 17. ramUsed, 18. diskCapacity, 19.diskUsed

	switch idx {
	case 1:
		return x.rataCpuTemp
	case 2:
		return x.rataGpuTemp
	case 3:
		return x.rataRamTemp
	case 4:
		return x.medCpuTemp
	case 5:
		return x.medGpuTemp
	case 6:
		return x.medRamTemp
	case 7:
		return x.modCpuTemp
	case 8:
		return x.modGpuTemp
	case 9:
		return x.modRamTemp
	case 10:
		return x.minCpuTemp
	case 11:
		return x.minGpuTemp
	case 12:
		return x.minRamTemp
	case 13:
		return x.maxCpuTemp
	case 14:
		return x.maxGpuTemp
	case 15:
		return x.maxRamTemp
	case 16:
		return x.ramCapacity
	case 17:
		return x.ramUsed
	case 18:
		return x.diskCapacity
	case 19:
		return x.diskUsed
	default:
		return 0
	}
}

func (x dataComponent) indexString(idx int) string {
	// id yang dicari : 1. cpuManufacturer, 2. gpuManufacturer, 3. cpuModel, 4.gpuModel, 5. cpuSerial, 6. gpuSerial
	// 7. OperatingSystem, 8. Status
	switch idx {
	case 1:
		return x.cpuManufacturer
	case 2:
		return x.gpuManufacturer
	case 3:
		return x.cpuModel
	case 4:
		return x.gpuModel
	case 5:
		return x.cpuSerial
	case 6:
		return x.gpuSerial
	case 7:
		return x.operatingSystem
	case 8:
		return x.status
	default:
		return ""
	}
}

func sequentialSearch(data *dataBase, searchData string, totalUser *int, id int) {
	//buat cari status atau yang lainnya berdasarkan string
	for i := 1; i < *totalUser; i++ {
		if data[i].indexString(id) == searchData {
			fmt.Printf("User %d: %s with data: %s \n", i, data[i].user, data[i].indexString(id))
		}
	}
}

func deleteDataUser(data *dataBase, loggedInUser int, kill *bool, login *bool, totalUser *int) {
	var exit bool = false
	var input int

	fmt.Printf("%45s", "DELETE PAGE")
	if loggedInUser == 0 {
		deleteDataMenuAdministrator(data, kill, totalUser)
	} else {
		for !exit {
			footer()
			fmt.Print("Input: ")
			fmt.Scan(&input)
			switch input {
			case 1:
				var confirmation string
				fmt.Println("Are you sure about that? note : ALL of your data will be gone in a blink of an eye.")
				footer()
				fmt.Print("Input (YES/NO): ")
				fmt.Scan(&confirmation)
				if upperCaseConverter(confirmation) == "YES" {
					deletion(data, loggedInUser, totalUser, 1)
					exit = true
				}
			case 2: //exit
				exit = true
			case 3: //kill
				exit = true
				*kill = true
			default:
				fmt.Println("Invalid Input")
				footer()
			}
		}
	}
}

func deleteDataMenuAdministrator(data *dataBase, kill *bool, totalUser *int) {
	var input int
	var exit bool = false
	var valid bool = false
	var searchType, index int
	var searchDataF float64
	var searchDataS string
	for !exit {
		footer()
		fmt.Print("Input: ")
		fmt.Scan(&input)
		switch input {
		case 1: //show all data with specific thing
			for !valid {
				fmt.Scan(&searchType)
				if searchType >= 1 && searchType <= 27 {
					valid = true
				}
			}
			if searchType >= 1 && searchType <= 19 {
				fmt.Print("Data: ")
				fmt.Scan(&searchDataF)
				index = binarySearch(data, searchDataF, totalUser, searchType)

				for data[index-1].indexFloat(searchType) == searchDataF {
					fmt.Printf("User %d: %s with data: %.2f\n ", index, data[index].user, data[index].indexFloat(searchType))
					index = index - 1
				}
				for data[index+1].indexFloat(searchType) == searchDataF {
					fmt.Printf("User %d: %s with data: %.2f\n", index, data[index].user, data[index].indexFloat(searchType))
				}
			} else {
				fmt.Print("Data: ")
				fmt.Scan(&searchDataS)
				sequentialSearch(data, searchDataS, totalUser, searchType-19)
			}
		case 2: //delete user data
			var deleteIndex int
			for !valid {
				fmt.Print("Which user's data do you want to delete: ")
				fmt.Scan(&deleteIndex)
				if deleteIndex < totalUser && deleteIndex > 0 {
					deletion(data, deleteIndex, totalUser, 1)
					valid = true
				}
			}
		case 3: //delete user

		case 4: //exit
			exit = true
		case 5: //kill
			exit = true
			*kill = true
		default:
			fmt.Println("Invalid Input")
			footer()
		}
	}
}

func deletion(data *dataBase, loggedInUser int, totalUser *int, id int) {
	tempUser := data[loggedInUser].user
	tempPassword := data[loggedInUser].userPassword

	switch id {
	case 1: //only delete data User
		data[loggedInUser] = dataComponent{}
		data[loggedInUser].user = tempUser
		data[loggedInUser].userPassword = tempPassword
		data[loggedInUser].dataSudahDiisi = false
	case 2: //delete User (admin only)
		for i := loggedInUser; i < *totalUser-1; i++ {
			data[i] = data[i+1]
		}
		data[*totalUser-1] = dataComponent{}
		*totalUser--
	}
	fmt.Println("Data Sucessfuly Deleted")
}

//why the fuck did i make ts so fucking complicated AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA

// func copyData(data *dataBase, totalUser *int, arrayCoppied *dataBase) {
// 	for i := 1; i < *totalUser; i++ {
// 		arrayCoppied[i] = data[i]
// 	}
// }

func statisticsMenu(data *dataBase, loggedInUser int, kill *bool, login *bool, totalUser *int) {

}

func changeDataUser(data *dataBase, loggedInUser int, kill *bool, login *bool, totalUser *int) {

}

func setData(data *dataBase, loggedInUser int) {
	var cpuOverheat bool = false
	var gpuOverheat bool = false
	var ramOverheat bool = false
	var sisaAvailableRam, sisaAvailableDisk float64

	nextMaintenance = data[loggedInUser].lastMaintenanceDate
	if data[loggedInUser].cpuManufacturer == "INTEL" {
		if data[loggedInUser].cpuModel == "PENTIUM" {
			if (data[loggedInUser].dataLoad && data[loggedInUser].rataCpuTemp >= 85) || (!data[loggedInUser].dataLoad && data[loggedInUser].rataCpuTemp >= 70) {
				cpuOverheat = true
			}
		} else if data[loggedInUser].cpuModel == "XEON" {
			if (data[loggedInUser].dataLoad && data[loggedInUser].rataCpuTemp >= 95) || (!data[loggedInUser].dataLoad && data[loggedInUser].rataCpuTemp >= 80) {
				cpuOverheat = true
			}
		} else if data[loggedInUser].cpuModel == "ATOM" {
			if (data[loggedInUser].dataLoad && data[loggedInUser].rataCpuTemp >= 80) || (!data[loggedInUser].dataLoad && data[loggedInUser].rataCpuTemp >= 65) {
				cpuOverheat = true
			}
		} else {
			if (data[loggedInUser].dataLoad && data[loggedInUser].rataCpuTemp >= 100) || (!data[loggedInUser].dataLoad && data[loggedInUser].rataCpuTemp >= 90) {
				cpuOverheat = true
			}
		}
	} else if data[loggedInUser].cpuManufacturer == "AMD" {
		if data[loggedInUser].cpuModel == "RYZEN" {
			if (data[loggedInUser].dataLoad && data[loggedInUser].rataCpuTemp >= 95) || (!data[loggedInUser].dataLoad && data[loggedInUser].rataCpuTemp >= 70) {
				cpuOverheat = true
			}
		} else if data[loggedInUser].cpuModel == "EPYC" {
			if (data[loggedInUser].dataLoad && data[loggedInUser].rataCpuTemp >= 95) || (!data[loggedInUser].dataLoad && data[loggedInUser].rataCpuTemp >= 85) {
				cpuOverheat = true
			}
		} else {
			if (data[loggedInUser].dataLoad && data[loggedInUser].rataCpuTemp >= 80) || (!data[loggedInUser].dataLoad && data[loggedInUser].rataCpuTemp >= 65) {
				cpuOverheat = true
			}
		}
	} else if data[loggedInUser].cpuManufacturer == "APPLE" {
		if (data[loggedInUser].dataLoad && data[loggedInUser].rataCpuTemp >= 95) || (!data[loggedInUser].dataLoad && data[loggedInUser].rataCpuTemp >= 70) {
			cpuOverheat = true
		}
	}

	if data[loggedInUser].gpuManufacturer != "NONE" {
		if data[loggedInUser].gpuManufacturer == "NVIDIA" {
			if (data[loggedInUser].dataLoad && data[loggedInUser].rataGpuTemp >= 85) || (!data[loggedInUser].dataLoad && data[loggedInUser].rataGpuTemp >= 70) {
				gpuOverheat = true
			}
		} else if data[loggedInUser].gpuManufacturer == "AMD" {
			if (data[loggedInUser].dataLoad && data[loggedInUser].rataGpuTemp >= 90) || (!data[loggedInUser].dataLoad && data[loggedInUser].rataGpuTemp >= 75) {
				gpuOverheat = true
			}
		} else if data[loggedInUser].gpuManufacturer == "APPLE" {
			if (data[loggedInUser].dataLoad && data[loggedInUser].rataGpuTemp >= 95) || (!data[loggedInUser].dataLoad && data[loggedInUser].rataGpuTemp >= 75) {
				gpuOverheat = true
			}
		}
	}

	if (data[loggedInUser].dataLoad && data[loggedInUser].rataRamTemp >= 85) || (!data[loggedInUser].dataLoad && data[loggedInUser].rataRamTemp >= 70) {
		ramOverheat = true
	}

	// var diskUsageByOS, ramUsageByOS float64
	var minimumDiskAvailable, minimumRamAvailable float64

	if data[loggedInUser].operatingSystem == "WINDOWS" {
		// ramUsageByOS = 0.35 *data[loggedInUser].ramCapacity
		// diskUsageByOS = 64
		minimumRamAvailable = 0.15 * data[loggedInUser].ramCapacity
		minimumDiskAvailable = 0.15 * data[loggedInUser].diskCapacity //%
	} else if data[loggedInUser].operatingSystem == "MACOS" {
		// ramUsageByOS = 0.2 * data[loggedInUser].ramCapacity //%
		// diskUsageByOS = 25
		minimumRamAvailable = 0.15 * data[loggedInUser].ramCapacity
		minimumDiskAvailable = 0.1 * data[loggedInUser].diskCapacity //%
	} else {
		// ramUsageByOS = 0.1 * data[loggedInUser].ramCapacity //%
		// diskUsageByOS = 20
		minimumRamAvailable = 0.05 * data[loggedInUser].ramCapacity
		minimumDiskAvailable = 0.15 * data[loggedInUser].diskCapacity //%
	}

	sisaAvailableRam = data[loggedInUser].ramCapacity - data[loggedInUser].ramUsed
	sisaAvailableDisk = data[loggedInUser].diskCapacity - data[loggedInUser].diskUsed

	var lowOnDisk, lowOnRam bool

	lowOnRam = sisaAvailableRam < minimumRamAvailable
	lowOnDisk = sisaAvailableDisk < minimumDiskAvailable

	if cpuOverheat && gpuOverheat && ramOverheat {
		data[loggedInUser].status = "VERY_CRITICAL"
	} else if cpuOverheat || gpuOverheat || ramOverheat {
		data[loggedInUser].status = "CRITICAL"
	} else if lowOnDisk || lowOnRam {
		data[loggedInUser].status = "WARNING"
	} else {
		data[loggedInUser].status = "GUD"
	}
}
