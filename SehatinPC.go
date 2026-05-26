package main

import "fmt"

const NMAX = 9999

type dataComponent struct {
	user, userPassword, serialCode               string
	batteryHealth                                float64 // dalam Persentase
	cpuManufacturer, gpuManufacturer             string  // CPU : Intel, AMD, Apple M-series GPU : Nvidia, AMD, Apple or NONE
	cpuModel, gpuModel, cpuSerial, gpuSerial     string  // CPU : Core, Pentium, Xeon, Atom, Ryzen, Athlon, Other, GPU : GeForce, Radeon, Apple or NONE
	rataCpuTemp, rataGpuTemp, rataRamTemp        float64 // dalam Celcius
	medCpuTemp, medGpuTemp, medRamTemp           float64 // dalam Celcius
	modCpuTemp, modGpuTemp, modRamTemp           float64 // dalam Celcius
	minCpuTemp, maxCpuTemp, minGpuTemp, maxGpuTemp, minRamTemp, maxRamTemp float64
	ramCapacity, ramUsed, diskCapacity, diskUsed float64 // dalam GiB
	dataLoad, dataSudahDiisi                     bool    // true = data dalam heavy load, false = data dalam idle load
	operatingSystem                              string  // Windows, Linux, MacOS
	lastMaintenanceDate, nextMaintenanceDate     string  // Format DD-MM-YYYY
	status                                       string  // Gud, Warning, Critical
	usingLaptop									 bool
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

	fmt.Printf("RULES : ")
	// nanti kutambahin rulesnya

	footer()
	fmt.Printf("\n%-45s%s\n\n", " ", "INPUT USER DATA (Page 1/2)")
	for !valid1 {
		fmt.Print("Are you using Laptop? (Yes/No): ")
		fmt.Scan(&usingLaptop)
		valid1 = checkValidityInput(usingLaptop, 8, "")
		if !valid1 {
			fmt.Println("Invalid Input")
		}else {
			usingLaptop = upperCaseConverter(usingLaptop)
			if usingLaptop == "YES" {
				isLaptop = true
			}else {
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
		}else if validGpuManuf && (x == "NONE") {
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
	fmt.Print("Input 10 of your current GPU Temperature in the last 20 seconds: ")
	for i = 0; i < 10; i++ {
		fmt.Scan(&gpuTemperature[i])
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
		}else {
			fmt.Println("Your RAM usage cannot exceed your RAM capacity.")
		}
	}
	valid2= false
	for !valid2 {
		fmt.Print("Current Disk Usage (GiB): ")
		fmt.Scan(&diskUsage)
		if diskUsage <= diskCapacity {
			valid2 = true
		}else {
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
	}else {
		batteryHealth = -1
	}
	for !valid2 {
		fmt.Print("Is it under Heavy Load? (Yes/No): ")
		fmt.Scan(&load)
		load = upperCaseConverter(load)
		if load == "YES" {
			heavyLoad = true
			valid2 = true
		}else if load == "NO" {
			heavyLoad = false
			valid2 = true
		}else {
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

func processData(min, max, med, mod, avg *float64, data [10]float64){
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
		}else if y == "AMD" {
			if x != "RYZEN" && x != "EPYC" && x != "ATHLON" {
				return false
			}
		}else if y == "APPLE" {
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
		}else if y == "AMD" {
			if x != "RADEON" {
				return false
			}
		}else if y == "APPLE" {
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
	fmt.Printf("GENERAL SPESIFICATIONS\n\n")
	fmt.Printf("User %d: %s\n", i, data[i].user)
	fmt.Println("Serial Code:", data[i].serialCode)
	fmt.Println("CPU: ", data[i].cpuManufacturer, data[i].cpuModel, data[i].cpuSerial)
	fmt.Println("GPU: ", data[i].gpuManufacturer, data[i].gpuModel, data[i].gpuSerial)
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

func binarySearch(data dataBase, searchData float64, totalUser, id int)int{
	//id yang dicari : 1. rataCpuTemp, 2. rataGpuTemp, 3. rataRamTemp, 4. medCpuTemp, 5. medGpuTemp, 6. medRamTemp
	// 7. modCpuTemp, 8. modGpuTemp, 9. modRamTemp, 10. minCpuTemp, 11. minGpuTemp, 12. minRamTemp,
	// 13. maxCpuTemp, 14. maxGpuTemp, 15. maxRamTemp, 16. ramCapacity, 17. ramUsed, 18. diskCapacity, 19.diskUsed 
	var right, left, middle int
	selectionSort(&data, totalUser, id)

	right = totalUser - 1
	left = 0 
	for left < right {
		middle = (right + left)/2
		if data[middle].rataCpuTemp > searchData {
			right = middle - 1
		}else if data[middle].rataCpuTemp < searchData {
			left = middle + 1 
		}else {
			return middle
		}
	}
	return -1
}

func selectionSort(data *dataBase, totalUser, id int) {
	//id yang dicari : 1. rataCpuTemp, 2. rataGpuTemp, 3. rataRamTemp, 4. medCpuTemp, 5. medGpuTemp, 6. medRamTemp
	// 7. modCpuTemp, 8. modGpuTemp, 9. modRamTemp, 10. minCpuTemp, 11. minGpuTemp, 12. minRamTemp,
	// 13. maxCpuTemp, 14. maxGpuTemp, 15. maxRamTemp, 16. ramCapacity, 17. ramUsed, 18. diskCapacity, 19.diskUsed 
	var idx int
	var temp dataComponent 	
	for i := 1; i < totalUser; i++ {
		idx = i
		temp = data[i]
		for y := i+1; y < totalUser; y++ {
			switch id {
			case 1:
				if data[y].rataCpuTemp < data[idx].rataCpuTemp {
					idx = y
				}
			case 2:
				if data[y].rataGpuTemp < data[idx].rataGpuTemp {
					idx = y
				}
			case 3:
				if data[y].rataRamTemp < data[idx].rataRamTemp {
					idx = y
				}
			case 4:
				if data[y].medCpuTemp < data[idx].medCpuTemp {
					idx = y
				}
			case 5:
				if data[y].medGpuTemp < data[idx].medGpuTemp {
					idx = y
				}
			case 6:
				if data[y].medRamTemp < data[idx].medRamTemp {
					idx = y
				}
			case 7:
				if data[y].modCpuTemp < data[idx].modCpuTemp {
					idx = y
				}
			case 8:
				if data[y].modGpuTemp < data[idx].modGpuTemp {
					idx = y
				}
			case 9:
				if data[y].modRamTemp < data[idx].modRamTemp {
					idx = y
				}
			case 10:
				if data[y].minCpuTemp < data[idx].minCpuTemp {
					idx = y
				}
			case 11:
				if data[y].minGpuTemp < data[idx].minGpuTemp {
					idx = y
				}
			case 12:
				if data[y].minRamTemp < data[idx].minRamTemp {
					idx = y
				}
			case 13:
				if data[y].maxCpuTemp < data[idx].maxCpuTemp {
					idx = y
				}
			case 14:
				if data[y].maxGpuTemp < data[idx].maxGpuTemp {
					idx = y
				}
			case 15:
				if data[y].maxRamTemp < data[idx].maxRamTemp {
					idx = y
				}
			case 16: 
				if data[y].ramCapacity < data[idx].ramCapacity {
					idx = y
				}
			case 17:
				if data[y].ramUsed < data[idx].ramUsed {
					idx = y
				}
			case 18:
				if data[y].diskCapacity < data[idx].diskCapacity {
					idx = y
				}
			case 19:
				if data[y].diskUsed < data[idx].diskUsed {
					idx = y
				}
			}
		}
		data[i] = data[idx]
		data[idx] = temp
	}
}

// func setIndex(x dataComponent, index int)dataComponent{
// 	switch index {
// 	case 1:return x.rataCpuTemp 
// 	case 2:return x.rataGpuTemp
// 	case 3:return x.rataRamTemp
// 	case 4:return x.medCpuTemp
// 	case 5:return x.medGpuTemp
// 	case 6:return x.medRamTemp
// 	case 7:return x.modCpuTemp
// 	case 8:return x.modGpuTemp
// 	case 9:return x.modRamTemp
// 	case 10:return x.minCpuTemp
// 	case 11:return x.minGpuTemp
// 	case 12:return x.minRamTemp
// 	case 13:return x.maxCpuTemp
// 	case 14:
// 	case 15:
// 	case 16:
// 	case 17:
// 	case 18:
// 	case 19:
// 	}
// }

func sequentialSearch(data dataBase, searchData string, totalUser int) {
	//buat cari status atau yang lainnya berdasarkan string
	
}

func deleteDataUser(data *dataBase, loggedInUser int, kill *bool, login *bool, totalUser *int) {

}

func statisticsMenu(data *dataBase, loggedInUser int, kill *bool, login *bool, totalUser *int) {

}

func changeDataUser(data *dataBase, loggedInUser int, kill *bool, login *bool, totalUser *int) {

}

func setData(OS, cpuManufacturer, cpuModel string, cpuOverheat, gpuOverheat, ramOverheat bool) {

}