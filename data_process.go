package main 

import "fmt"

func processData(min, max, med, mod, avg *float64, data [10]float64) {
	searchMinMax(data, min, max)
	*avg = searchRataRata(data)
	*med = searchMedian(data)
	*mod = searchModus(data)
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

func cloningData(dataAsli *dataBase, dataCopy *dataBase, totalUser int) {
	for i := 0; i < totalUser; i++ {
		dataCopy[i] = dataAsli[i]
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
	case 9:
		return x.serialCode
	case 10:
		return x.user
	default:
		return ""
	}
}



func changeDataUserLogic(data *dataBase, loggedInUser int, kill *bool, login *bool) {
	var exit bool = false
	var input int

	fmt.Printf("\n%-43s%s\n\n", " ", "CHANGE DATA MENU USER")
	fmt.Printf("%-40s%s\n", " ", "1. Serial Code")
	fmt.Printf("%-40s%s\n", " ", "2. CPU")
	fmt.Printf("%-40s%s\n", " ", "3. GPU")
	fmt.Printf("%-40s%s\n", " ", "4. RAM Capacity")
	fmt.Printf("%-40s%s\n", " ", "5. Disk Capacity")
	fmt.Printf("%-40s%s\n", " ", "6. CPU Temperature Data")
	fmt.Printf("%-40s%s\n", " ", "7. GPU Temperature Data")
	fmt.Printf("%-40s%s\n", " ", "8. RAM Temperature Data")
	fmt.Printf("%-40s%s\n", " ", "9. Ram Usage")
	fmt.Printf("%-40s%s\n", " ", "10. Disk Usage")
	fmt.Printf("%-40s%s\n", " ", "11. Operating System")
	fmt.Printf("%-40s%s\n", " ", "12. Load Status")
	fmt.Printf("%-40s%s\n", " ", "13. Last Maintenance Date")
	if data[loggedInUser].usingLaptop {
		fmt.Printf("%-40s%s\n", " ", "14. Battery Health")
	}
	fmt.Printf("%-40s%s\n", " ", "15. Exit")
	// fmt.Printf("%-40s%s\n", " ", "16. Logout")
	// fmt.Printf("%-40s%s\n", " ", "17. Kill Program")
	footer()
	for !exit {
		fmt.Print("Input: ")
		fmt.Scan(&input)

		switch input {
		case 1:
			var newSerial string
			var valid bool = false
			for !valid {
				fmt.Print("New Serial Code: ")
				fmt.Scan(&newSerial)
				valid = checkValidityInput(newSerial, 1, "")
				if !valid {
					fmt.Println("Serial Code cannot be null or none")
				}
			}
			data[loggedInUser].serialCode = newSerial
			fmt.Println("Serial Code Changed")
			exit = true

		case 2:
			var newManuf, newModel, newSerial string
			validManuf := false
			validModel := false
			for !validManuf || !validModel {
				fmt.Print("CPU: ")
				fmt.Scan(&newManuf, &newModel, &newSerial)
				validManuf = checkValidityInput(newManuf, 2, "")
				validModel = checkValidityInput(newModel, 3, newManuf)
				if !validManuf || !validModel {
					fmt.Println("Invalid Manufacturer or Invalid Model Input")
				}
			}
			data[loggedInUser].cpuManufacturer = upperCaseConverter(newManuf)
			data[loggedInUser].cpuModel = upperCaseConverter(newModel)
			data[loggedInUser].cpuSerial = newSerial
			setData(data, loggedInUser)
			fmt.Println("CPU Info updated.")
			exit = true

		case 3:
			var newManuf, newModel, newSerial string
			validGpuManuf := false
			validGpuModel := false
			for !validGpuManuf || !validGpuModel {
				fmt.Print("GPU: ")
				fmt.Scan(&newManuf)
				validGpuManuf = checkValidityInput(newManuf, 4, "")
				x := upperCaseConverter(newManuf)
				if validGpuManuf && (x != "NONE") {
					fmt.Scan(&newModel, &newSerial)
					validGpuModel = checkValidityInput(newModel, 5, newManuf)
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
			data[loggedInUser].gpuManufacturer = upperCaseConverter(newManuf)
			data[loggedInUser].gpuModel = upperCaseConverter(newModel)
			data[loggedInUser].gpuSerial = newSerial
			setData(data, loggedInUser)
			fmt.Println("GPU Info updated.")
			exit = true

		case 4:
			var newCap float64
			var valid bool = false
			for !valid {
				fmt.Print("New RAM Capacity (GiB): ")
				fmt.Scan(&newCap)
				if newCap > 0 {
					valid = true
				} else {
					fmt.Println("Capacity must be greater than 0")
				}
			}
			data[loggedInUser].ramCapacity = newCap
			fmt.Println("RAM Capacity updated.")
			exit = true

		case 5:
			var newCap float64
			var valid bool = false
			for !valid {
				fmt.Print("New Disk Capacity (GiB): ")
				fmt.Scan(&newCap)
				if newCap > 0 {
					valid = true
				} else {
					fmt.Println("Capacity must be greater than 0")
				}
			}
			data[loggedInUser].diskCapacity = newCap
			fmt.Println("Disk Capacity updated.")
			exit = true

		case 6:
			var newTemps [10]float64
			fmt.Print("Input 10 CPU Temperatures: ")
			for i := 0; i < 10; i++ {
				fmt.Scan(&newTemps[i])
			}
			var avg, med, mod, min, max float64
			processData(&min, &max, &med, &mod, &avg, newTemps)
			data[loggedInUser].rataCpuTemp = avg
			data[loggedInUser].medCpuTemp = med
			data[loggedInUser].modCpuTemp = mod
			data[loggedInUser].minCpuTemp = min
			data[loggedInUser].maxCpuTemp = max
			setData(data, loggedInUser)
			fmt.Println("CPU Temperatures updated.")
			exit = true

		case 7:
			if data[loggedInUser].gpuManufacturer == "NONE" {
				fmt.Println("No GPU installed, skipping.")
			} else {
				var newTemps [10]float64
				fmt.Print("Input 10 GPU Temperatures: ")
				for i := 0; i < 10; i++ {
					fmt.Scan(&newTemps[i])
				}
				var avg, med, mod, min, max float64
				processData(&min, &max, &med, &mod, &avg, newTemps)
				data[loggedInUser].rataGpuTemp = avg
				data[loggedInUser].medGpuTemp = med
				data[loggedInUser].modGpuTemp = mod
				data[loggedInUser].minGpuTemp = min
				data[loggedInUser].maxGpuTemp = max
				setData(data, loggedInUser)
				fmt.Println("GPU Temperatures updated.")
			}
			exit = true

		case 8:
			var newTemps [10]float64
			fmt.Print("Input 10 RAM Temperatures: ")
			for i := 0; i < 10; i++ {
				fmt.Scan(&newTemps[i])
			}
			var avg, med, mod, min, max float64
			processData(&min, &max, &med, &mod, &avg, newTemps)
			data[loggedInUser].rataRamTemp = avg
			data[loggedInUser].medRamTemp = med
			data[loggedInUser].modRamTemp = mod
			data[loggedInUser].minRamTemp = min
			data[loggedInUser].maxRamTemp = max
			setData(data, loggedInUser)
			fmt.Println("RAM Temperatures updated.")
			exit = true

		case 9:
			var newUsage float64
			var valid bool = false
			for !valid {
				fmt.Print("New RAM Usage (GiB): ")
				fmt.Scan(&newUsage)
				if newUsage >= 0 && newUsage <= data[loggedInUser].ramCapacity {
					valid = true
				} else {
					fmt.Println("RAM usage cannot exceed RAM capacity or be negative")
				}
			}
			data[loggedInUser].ramUsed = newUsage
			setData(data, loggedInUser)
			fmt.Println("RAM Usage updated.")
			exit = true

		case 10:
			var newUsage float64
			var valid bool = false
			for !valid {
				fmt.Print("New Disk Usage (GiB): ")
				fmt.Scan(&newUsage)
				if newUsage >= 0 && newUsage <= data[loggedInUser].diskCapacity {
					valid = true
				} else {
					fmt.Println("Disk usage cannot exceed Disk capacity or be negative")
				}
			}
			data[loggedInUser].diskUsed = newUsage
			setData(data, loggedInUser)
			fmt.Println("Disk Usage updated.")
			exit = true

		case 11:
			var newOS string
			var valid bool = false
			for !valid {
				fmt.Print("New Operating System (Windows/Linux/MacOS): ")
				fmt.Scan(&newOS)
				valid = checkValidityInput(newOS, 7, "")
				if !valid {
					fmt.Println("Invalid OS, must be Windows / Linux / MacOS")
				}
			}
			data[loggedInUser].operatingSystem = upperCaseConverter(newOS)
			setData(data, loggedInUser)
			fmt.Println("Operating System updated.")
			exit = true

		case 12:
			var newLoad string
			var valid bool = false
			for !valid {
				fmt.Print("Heavy Load? (Yes/No): ")
				fmt.Scan(&newLoad)
				newLoad = upperCaseConverter(newLoad)
				if newLoad == "YES" {
					data[loggedInUser].dataLoad = true
					valid = true
				} else if newLoad == "NO" {
					data[loggedInUser].dataLoad = false
					valid = true
				} else {
					fmt.Println("Input must be Yes or No")
				}
			}
			setData(data, loggedInUser)
			fmt.Println("Load Status updated.")
			exit = true

		case 13:
			var newDate string
			var valid bool = false
			for !valid {
				fmt.Print("New Last Maintenance Date (DD-MM-YYYY): ")
				fmt.Scan(&newDate)
				valid = checkValidityInput(newDate, 6, "")
				if !valid {
					fmt.Println("Invalid date format, use DD-MM-YYYY")
				}
			}
			data[loggedInUser].lastMaintenanceDate = newDate
			data[loggedInUser].nextMaintenanceDate = nextMaintenance(newDate, data[loggedInUser].status)
			fmt.Println("Maintenance Date updated.")
			exit = true

		case 14:
			if data[loggedInUser].usingLaptop {
				var newBatt float64
				var valid bool = false
				for !valid {
					fmt.Print("New Battery Health (%): ")
					fmt.Scan(&newBatt)
					if newBatt >= 0 && newBatt <= 100 {
						valid = true
					} else {
						fmt.Println("Battery health must be between 0 and 100")
					}
				}
				data[loggedInUser].batteryHealth = newBatt
				fmt.Println("Battery Health updated.")
			} else {
				fmt.Println("Invalid Input")
			}
			exit = true
		case 15:
			exit = true
		// case 16:
		// 	exit = true
		// 	*login = false
		// case 17:
		// 	exit = true
		// 	*kill = true
		default:
			fmt.Println("Invalid Input")
		}
	}
}

func setData(data *dataBase, loggedInUser int) {
	var cpuOverheat bool = false
	var gpuOverheat bool = false
	var ramOverheat bool = false
	var sisaAvailableRam, sisaAvailableDisk float64
	var jumlahKomponen int = 0  

	// nextMaintenance := data[loggedInUser].lastMaintenanceDate
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

	if cpuOverheat {
		jumlahKomponen++
	}
	if gpuOverheat {
		jumlahKomponen++
	}
	if ramOverheat {
		jumlahKomponen++
	}
	if lowOnDisk {
		jumlahKomponen++
	}
	if lowOnRam {
		jumlahKomponen++
	}
	data[loggedInUser].jumlahKomponenRusak = jumlahKomponen

	if cpuOverheat && gpuOverheat && ramOverheat {
		data[loggedInUser].status = "VERY_CRITICAL"
	} else if cpuOverheat || gpuOverheat || ramOverheat {
		data[loggedInUser].status = "CRITICAL"
	} else if lowOnDisk || lowOnRam {
		data[loggedInUser].status = "WARNING"
	} else {
		data[loggedInUser].status = "GUD"
	}
	data[loggedInUser].nextMaintenanceDate = nextMaintenance(data[loggedInUser].lastMaintenanceDate, data[loggedInUser].status)
}

func nextMaintenance(lastMaintenanceDate string, status string) string {
	if len(lastMaintenanceDate) != 10 {
		return "Invalid Date"
	}

	tanggal := int(lastMaintenanceDate[0]-'0')*10 + int(lastMaintenanceDate[1]-'0')
	bulan := int(lastMaintenanceDate[3]-'0')*10 + int(lastMaintenanceDate[4]-'0')
	tahun := int(lastMaintenanceDate[6]-'0')*1000 + int(lastMaintenanceDate[7]-'0')*100 + int(lastMaintenanceDate[8]-'0')*10 + int(lastMaintenanceDate[9]-'0')

	var tanggalDalamBulan int
	if bulan == 1 || bulan == 3 || bulan == 5 || bulan == 7 || bulan == 8 || bulan == 10 || bulan == 12 {
		tanggalDalamBulan = 31
	} else if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
		tanggalDalamBulan = 30
	} else if (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0) {
		tanggalDalamBulan = 29
	} else {
		tanggalDalamBulan = 28
	}

	if status == "VERY_CRITICAL" {
		tanggal = tanggal + 1
	} else if status == "CRITICAL" {
		tanggal = tanggal + 7
	} else if status == "WARNING" {
		bulan = bulan + 3
	} else if status == "GUD" {
		bulan = bulan + 6
	}

	if tanggal > tanggalDalamBulan {
		tanggal = tanggal - tanggalDalamBulan
		bulan = bulan + 1
	}
	if bulan > 12 {
		bulan = bulan - 12
		tahun = tahun + 1
	}

	hasil := ""
	if tanggal < 10 {
		hasil = hasil + "0"
	}
	hasil = hasil + stringConverter(tanggal) + "-"
	if bulan < 10 {
		hasil = hasil + "0"
	}
	hasil = hasil + stringConverter(bulan) + "-"
	hasil = hasil + stringConverter(tahun)
	return hasil
}

func stringConverter(x int) string {
	if x == 0 {
		return "0"
	}
	var hasil string
	hasil = ""
	for x > 0 {
		hasil = string(rune('0'+x%10)) + hasil
		x = x / 10
	}
	return hasil
}