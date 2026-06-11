package main

import "fmt"

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
		if upperCaseConverter(username) == "EXIT" {
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
		if upperCaseConverter(username) == "EXIT" {
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
						if data[i].dataSudahDiisi {
							fmt.Printf("%-2sStatus: %s\n", "", data[i].status)
						} else {
							fmt.Printf("%-2sStatus: Data unavailable\n", "")
						}
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

	fmt.Printf("RULES : \n  ")
	fmt.Printf("%-20s: Intel | AMD | Apple\n  ", "CPU Manufacturer")
	fmt.Printf("%-20s: Core | Pentium | Xeon | Atom\n  ", "CPU Model for Intel")
	fmt.Printf("%-20s: Ryzen | Epyc | Athlon\n  ", "CPU Model for AMD")
	fmt.Printf("%-20s: M (contoh: M1, M2...)\n  ", "CPU Model for Apple")
	fmt.Printf("%-20s: NVIDIA | AMD | Apple | NONE\n  ", "GPU Manufacturer")
	fmt.Printf("%-20s: RTX | GTX | MAX-Q\n  ", "GPU Model for NVIDIA")
	fmt.Printf("%-20s: Radeon\n  ", "GPU Model for AMD")
	fmt.Printf("%-20s: M (contoh: M1, M2...)\n  ", "GPU Model for Apple")
	fmt.Printf("%-20s: Windows | Linux | MacOS\n\n", "Operating System")

	fmt.Printf("FORMAT : \n  ")
	fmt.Printf("%-20s: CPU Manufacturer Cpu Model Cpu Serial (contoh : Intel Core i7)\n  ", "CPU")
	fmt.Printf("%-20s: GPU Manufacturer GPU Model GPU Serial (contoh : NVIDIA GTX 1650)\n", "GPU")
	fmt.Printf("%-20s: DD-MM-YYYY\n", "Date Format")

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
			if !validGpuModel {
				fmt.Println("Invalid GPU Model Input")
			}
		} else if validGpuManuf && (x == "NONE") {
			validGpuModel = true
		}
		if !validGpuManuf {
			fmt.Println("Invalid Manufacturer Input")
		}
	}
	valid1 = false
	for !valid1 {
		fmt.Print("Ram Capacity(GiB): ")
		fmt.Scan(&ramCapacity)
		if ramCapacity > 0 {
			valid1 = true
		} else {
			fmt.Println("Insufficient RAM Space")
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
	// fmt.Printf("RULES : \n")
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
			} else {
				fmt.Println("Invalid input")
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
	fmt.Println("Type (1) to Save, Type (2) to Exit Menu")
	footer()
	var input int
	for !valid2 {
		fmt.Print("Input: ")
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
			if upperCaseConverter(gpuManufacturer) != "NONE" {
				processData(&min, &max, &med, &mod, &avg, gpuTemperature)
				data[loggedInUser].rataGpuTemp = avg
				data[loggedInUser].medGpuTemp = med
				data[loggedInUser].modGpuTemp = mod
				data[loggedInUser].minGpuTemp = min
				data[loggedInUser].maxGpuTemp = max
			} else {
				data[loggedInUser].rataGpuTemp = -1
				data[loggedInUser].medGpuTemp = -1
				data[loggedInUser].modGpuTemp = -1
				data[loggedInUser].minGpuTemp = -1
				data[loggedInUser].maxGpuTemp = -1
			}
			processData(&min, &max, &med, &mod, &avg, ramTemperature)
			data[loggedInUser].rataRamTemp = avg
			data[loggedInUser].medRamTemp = med
			data[loggedInUser].modRamTemp = mod
			data[loggedInUser].minRamTemp = min
			data[loggedInUser].maxRamTemp = max
			data[loggedInUser].ramUsed = ramUsage
			data[loggedInUser].diskUsed = diskUsage
			setData(data, loggedInUser)
			fmt.Println("Data saved")
			valid2 = true
		case 2:
			valid2 = true
		// case 3:
		// 	*login = false
		// 	valid2 = true
		// case 4:
		// 	*kill = true
		// 	valid2 = true
		default:
			fmt.Println("Invalid Input")
		}
	}
}

func jenisSorting() bool {
	var input int
	var valid bool = false
	var ascending bool
	fmt.Printf("%-45s%s\n%-45s%s\n", "", "1. Ascending", "", "2. Descending")
	footer()
	for !valid {
		fmt.Print("Input: ")
		fmt.Scan(&input)
		if input == 1 || input == 2 {
			valid = true
		} else {
			fmt.Println("Invalid Input")
		}
	}
	ascending = input == 1
	return ascending
}

func deleteDataUser(data *dataBase, loggedInUser int, kill *bool, login *bool, totalUser *int) {
	var exit bool = false
	var input int

	if loggedInUser == 0 {
		deleteDataMenuAdministrator(data, kill, login, totalUser)
	} else {
		fmt.Printf("%45s", "DELETE PAGE\n")
		fmt.Printf("%-42s%s\n%-42s%s\n", " ", "1. Delete My Data", " ", "2. Exit")
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
				} else if upperCaseConverter(confirmation) == "NO" {
					exit = false
				} else {
					fmt.Println("Invalid Input")
				}
			case 2: //exit
				exit = true
			// case 3: //logout
			// 	*login = false
			// 	exit = true
			// case 4: //kill
			// 	*kill = true
			// 	exit = true
			default:
				fmt.Println("Invalid Input")
				footer()
			}
		}
	}
}

// nanti dibenerin
func deleteDataMenuAdministrator(data *dataBase, kill *bool, login *bool, totalUser *int) {
	var input int
	var exit bool = false
	var valid bool = false
	var searchType int
	var searchDataF float64
	var searchDataS string

	fmt.Printf("\n%-36s%s\n\n", " ", "DELETE MENU FOR ENDMINISTRATOR")
	fmt.Printf("%-33s%s\n", " ", "1. Show All User With Specific Data")
	fmt.Printf("%-33s%s\n", " ", "2. Delete User PC Data")
	fmt.Printf("%-33s%s\n", " ", "3. Delete User Account")
	fmt.Printf("%-33s%s\n", " ", "4. Exit")
	// fmt.Printf("%-33s%s\n", " ", "5. Kill Program")

	for !exit {
		footer()
		fmt.Print("Input: ")
		fmt.Scan(&input)
		switch input {
		case 1: //show all data with specific thing
			if *totalUser <= 1 {
				fmt.Println("No user data available")
				footer()
			} else {
				fmt.Printf("%-33s%s\n", " ", "1. Average CPU Temperature")
				fmt.Printf("%-33s%s\n", " ", "2. Average GPU Temperature")
				fmt.Printf("%-33s%s\n", " ", "3. Average RAM Temperature")
				fmt.Printf("%-33s%s\n", " ", "4. Median CPU Temperature")
				fmt.Printf("%-33s%s\n", " ", "5. Median GPU Temperature")
				fmt.Printf("%-33s%s\n", " ", "6. Median RAM Temperature")
				fmt.Printf("%-33s%s\n", " ", "7. Modus CPU Temperature")
				fmt.Printf("%-33s%s\n", " ", "8. Modus GPU Temperature")
				fmt.Printf("%-33s%s\n", " ", "9. Modus RAM Temperature")
				fmt.Printf("%-33s%s\n", " ", "10. Minimum CPU Temperature")
				fmt.Printf("%-33s%s\n", " ", "11. Minimum GPU Temperature")
				fmt.Printf("%-33s%s\n", " ", "12. Minimum RAM Temperature")
				fmt.Printf("%-33s%s\n", " ", "13. Maximum CPU Temperature")
				fmt.Printf("%-33s%s\n", " ", "14. Maximum GPU Temperature")
				fmt.Printf("%-33s%s\n", " ", "15. Maximum RAM Temperature")
				fmt.Printf("%-33s%s\n", " ", "16. RAM Capacity")
				fmt.Printf("%-33s%s\n", " ", "17. RAM Used")
				fmt.Printf("%-33s%s\n", " ", "18. Disk Capacity")
				fmt.Printf("%-33s%s\n", " ", "19. Disk Used")
				fmt.Printf("%-33s%s\n", " ", "20. CPU Manufacturer")
				fmt.Printf("%-33s%s\n", " ", "21. GPU Manufacturer")
				fmt.Printf("%-33s%s\n", " ", "22. CPU Model")
				fmt.Printf("%-33s%s\n", " ", "23. GPU Model")
				fmt.Printf("%-33s%s\n", " ", "24. CPU Serial")
				fmt.Printf("%-33s%s\n", " ", "25. GPU Serial")
				fmt.Printf("%-33s%s\n", " ", "26. Operating System")
				fmt.Printf("%-33s%s\n", " ", "27. User Status")
				fmt.Printf("%-33s%s\n", " ", "28. Serial Code")
				fmt.Printf("%-33s%s\n", " ", "29. Exit")
				// fmt.Printf("%-33s%s\n", " ", "30. Logout")
				// fmt.Printf("%-33s%s\n", " ", "31. Kill Program")

				valid = false
				for !valid {
					fmt.Print("Input: ")
					fmt.Scan(&searchType)
					if searchType >= 0 && searchType <= 31 {
						valid = true
					} else {
						fmt.Println("Invalid Input")
					}
				}
				if searchType == 29 {
					exit = true
					// } else if searchType == 30 {
					// 	*login = false
					// 	exit = true
					// } else if searchType == 31 {
					// 	*kill = true
					// 	exit = true
				} else if searchType >= 1 && searchType <= 19 {
					fmt.Print("Data: ")
					fmt.Scan(&searchDataF)
					var copyData dataBase
					cloningData(data, &copyData, *totalUser)
					var batasKanan, batasKiri int
					binarySearch(&copyData, searchDataF, totalUser, searchType, &batasKanan, &batasKiri)
					if batasKiri == -1 {
						fmt.Println("No data found on any user")
					} else {
						for i := batasKiri; i <= batasKanan; i++ {
							fmt.Printf("User: %s with data: %.2f\n", copyData[i].user, copyData[i].indexFloat(searchType))
						}
					}
				} else if searchType >= 20 && searchType <= 28 {
					fmt.Print("Data:")
					fmt.Scan(&searchDataS)
					sequentialSearch(data, searchDataS, totalUser, searchType-19)
				}
			}
			exit = true
		case 2: //delete user data
			var deleteUser string
			valid = false
			for !valid {
				fmt.Print("Which user's data do you want to delete: ")
				fmt.Scan(&deleteUser)
				deleteIndex := sequentialSearchIndex(data, deleteUser, totalUser)
				if deleteIndex != -1 {
					fmt.Println("This User's Data: ")
					outputDataUserFormat(data, deleteIndex)
					fmt.Print("Are you sure about that? (YES/NO)")
					var confirm string
					fmt.Scan(&confirm)
					if upperCaseConverter(confirm) == "YES" {
						deletion(data, deleteIndex, totalUser, 1)
						valid = true
					} else {
						valid = true
					}
				} else {
					fmt.Println("User not found")
				}
			}
			exit = true
		case 3: //delete user
			if *totalUser <= 1 {
				fmt.Println("No user data available")
				footer()
				exit = true
			} else {
				var deleteIndex int
				var deleteUser string

				valid = false
				for !valid {
					for !valid {
						fmt.Print("Which user's data do you want to delete: ")
						fmt.Scan(&deleteUser)
						deleteIndex = sequentialSearchIndex(data, deleteUser, totalUser)
						if deleteIndex != -1 {
							fmt.Println("This User's Data: ")
							outputDataUserFormat(data, deleteIndex)
							fmt.Print("Are you sure about that this is permanent? (YES/NO)")
							var confirm string
							fmt.Scan(&confirm)
							if upperCaseConverter(confirm) == "YES" {
								deletion(data, deleteIndex, totalUser, 2)
								valid = true
							} else {
								valid = true
							}
						} else {
							fmt.Println("User not found")
						}
					}
				}
			}
			exit = true
		case 4: //exit
			exit = true
		// case 5: //kill
		// 	exit = true
		// 	*kill = true
		default:
			fmt.Println("Invalid Input")
			footer()
		}
	}
}

func statisticsMenu(data *dataBase, loggedInUser int, kill *bool, login *bool, totalUser *int) {
	var exit bool = false
	var input int

	fmt.Printf("\n%-44s%s\n\n", "", "STATISTIC MENU")
	if loggedInUser == 0 {
		for !exit {
			fmt.Printf("%-33s%s\n", " ", "1. Show component status (all users)")
			fmt.Printf("%-33s%s\n", " ", "2. Show temperature statistics (all users)")
			fmt.Printf("%-33s%s\n", " ", "3. Sort users by serial code") //selection Sort
			fmt.Printf("%-33s%s\n", " ", "4. Sort users by CPU temp")    //selection Sort
			fmt.Printf("%-33s%s\n", " ", "5. Sort users by GPU temp")    //selection Sort
			fmt.Printf("%-33s%s\n", " ", "6. Sort users by RAM temp")    //Insertion sort
			fmt.Printf("%-33s%s\n", " ", "7. Search user by status")     //sequential search
			fmt.Printf("%-33s%s\n", " ", "8. Search user by CPU temp")   //binary search
			fmt.Printf("%-33s%s\n", " ", "9. Search user by SerialCode") //sequential search
			fmt.Printf("%-33s%s\n", " ", "10. Exit")
			// fmt.Printf("%-33s%s\n", " ", "11. Logout")
			// fmt.Printf("%-33s%s\n", " ", "12. Kill Program")
			footer()
			fmt.Print("Input: ")
			fmt.Scan(&input)
			switch input {
			case 1:
				showUserStatus(data, totalUser)
			case 2:
				showAverageTemp(data, totalUser)
			case 3:
				showSortedSerialCode(data, totalUser)
			case 4:
				showSortedCpuTemp(data, totalUser)
			case 5:
				showSortedGpuTemp(data, totalUser)
			case 6:
				showSortedRamTemp(data, totalUser)
			case 7:
				searchUserStatus(data, totalUser)
			case 8:
				searchUserCpuTemp(data, totalUser)
			case 9:
				searchUserSerialCode(data, totalUser)
			case 10:
				exit = true
			// case 11:
			// 	*login = false
			// 	exit = true
			// case 12:
			// 	*kill = true
			// 	exit = true
			default:
				fmt.Println("Invalid Input")
			}
		}
	} else {
		fmt.Printf("%-33s%s\n", " ", "1. Show my status")
		fmt.Printf("%-33s%s\n", " ", "2. Show my temperature statistics")
		fmt.Printf("%-33s%s\n", " ", "3. Exit")
		// fmt.Printf("%-33s%s\n", " ", "4. Logout")
		// fmt.Printf("%-33s%s\n", " ", "5. Kill Program")
		for !exit {
			footer()
			fmt.Print("Input: ")
			fmt.Scan(&input)
			switch input {
			case 1:
				fmt.Printf("\nUser %s\n", data[loggedInUser].user)
				fmt.Printf("%-2sStatus: %s\n", " ", data[loggedInUser].status)
			case 2:
				showUserTempStats(data, loggedInUser)
			case 3:
				exit = true
			// case 4:
			// 	*login = false
			// 	exit = true
			// case 5:
			// 	*kill = true
			// 	exit = true
			default:
				fmt.Println("Invalid Input")
			}
		}
	}
}

func changeDataUser(data *dataBase, loggedInUser int, kill *bool, login *bool, totalUser *int) {
	var exit bool = false
	var input string
	// buat admin
	if loggedInUser == 0 {
		if *totalUser <= 1 {
			fmt.Println("No user data, so you kanot chenj aowkaokwaoakow")
			footer()
		} else {
			fmt.Printf("\n%-41s%s\n\n", " ", "CHANGE DATA FOR USER")
			footer()
			fmt.Printf("%-34sInput User index you want to change\n", " ")
			fmt.Printf("%-44sTotal User: %d\n", " ", *totalUser-1)
			fmt.Printf("Type Exit to go back\n")
			footer()
			fmt.Println("USERNAME: ")
			for i := 1; i < *totalUser; i++ {
				fmt.Printf("User: %s\n", data[i].user)
			}
			footer()
			for !exit {
				fmt.Print("Input: ")
				fmt.Scan(&input)
				if upperCaseConverter(input) == "EXIT" {
					exit = true
				} else {
					userIndex := sequentialSearchIndex(data, input, totalUser)
					if userIndex != -1 {
						if data[userIndex].dataSudahDiisi {
							changeDataUserLogic(data, userIndex, kill, login)
						} else {
							fmt.Println("No data available to change for this user")
						}
						exit = true
					} else {
						fmt.Println("User not found")
					}
				}
			}
		}
		//buat user (bisa ubah data sendiri)
	} else {
		if data[loggedInUser].dataSudahDiisi {
			changeDataUserLogic(data, loggedInUser, kill, login)
		} else {
			fmt.Println("No data available to change")
		}
	}
}
