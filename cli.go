package main 

import "fmt"

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

func outputDataUserFormat(data *dataBase, i int) {
	fmt.Printf("%-40s\n\n", "GENERAL SPESIFICATIONS")
	fmt.Printf("%s %d%-20s%s %s\n", "User", i, "", ":", data[i].user)
	fmt.Printf("%-25s %s %s\n", "Serial Code", ":", data[i].serialCode)
	fmt.Printf("%-25s %s %s %s %s\n", "CPU", ":", data[i].cpuManufacturer, data[i].cpuModel, data[i].cpuSerial)

	if data[i].gpuManufacturer != "NONE" {
		fmt.Printf("%-25s %s %s %s %s\n", "GPU", ":", data[i].gpuManufacturer, data[i].gpuModel, data[i].gpuSerial)
	}

	if data[i].usingLaptop {
		fmt.Printf("%-25s %s %.2f%%\n", "Battery Health", ":", data[i].batteryHealth)
	}
	fmt.Printf("%-25s %s %s\n", "Operating System", ":", data[i].operatingSystem)

	footer()
	fmt.Printf("%-25s\n\n", "PERFORMANCE SPESIFICATIONS")
	fmt.Printf("CPU SPEC: \n")
	fmt.Printf("%-25s %s %.2f°C\n", "Average CPU Temperature", ":", data[i].rataCpuTemp)
	fmt.Printf("%-25s %s %.2f°C\n", "Median CPU Temperature", ":", data[i].medCpuTemp)
	fmt.Printf("%-25s %s %.2f°C\n", "Modus CPU Temperature", ":", data[i].modCpuTemp)
	fmt.Printf("%-25s %s %.2f°C\n", "Min CPU Temperature", ":", data[i].minCpuTemp)
	fmt.Printf("%-25s %s %.2f°C\n", "Max CPU Temperature", ":", data[i].maxCpuTemp)

	if data[i].gpuManufacturer != "NONE" {
		fmt.Printf("\nGPU SPEC: \n")
		fmt.Printf("%-25s %s %.2f°C\n", "Average GPU Temperature", ":", data[i].rataGpuTemp)
		fmt.Printf("%-25s %s %.2f°C\n", "Median GPU Temperature", ":", data[i].medGpuTemp)
		fmt.Printf("%-25s %s %.2f°C\n", "Modus GPU Temperature", ":", data[i].modGpuTemp)
		fmt.Printf("%-25s %s %.2f°C\n", "Min GPU Temperature", ":", data[i].minGpuTemp)
		fmt.Printf("%-25s %s %.2f°C\n", "Max GPU Temperature", ":", data[i].maxGpuTemp)
	}

	fmt.Printf("\nRAM SPEC: \n")
	fmt.Printf("%-25s %s %.2f GiB\n", "RAM Capacity", ":", data[i].ramCapacity)
	fmt.Printf("%-25s %s %.2f GiB\n", "RAM Used", ":", data[i].ramUsed)
	fmt.Printf("%-25s %s %.2f°C\n", "Average RAM Temperature", ":", data[i].rataRamTemp)
	fmt.Printf("%-25s %s %.2f°C\n", "Median RAM Temperature", ":", data[i].medRamTemp)
	fmt.Printf("%-25s %s %.2f°C\n", "Modus RAM Temperature", ":", data[i].modRamTemp)
	fmt.Printf("%-25s %s %.2f°C\n", "Min RAM Temperature", ":", data[i].minRamTemp)
	fmt.Printf("%-25s %s %.2f°C\n", "Max RAM Temperature", ":", data[i].maxRamTemp)

	fmt.Printf("\nDISK SPEC: \n")
	fmt.Printf("%-25s %s %.2f GiB\n", "Disk Capacity", ":", data[i].diskCapacity)
	fmt.Printf("%-25s %s %.2f GiB\n", "Disk Used", ":", data[i].diskUsed)

	footer()
	fmt.Printf("MAINTENANCE HISTORY\n\n")
	fmt.Printf("%-25s %s %v\n", "Last Maintenance Date", ":", data[i].lastMaintenanceDate)
	fmt.Printf("%-25s %s %v\n", "Next Maintenance Date", ":", data[i].nextMaintenanceDate)
	fmt.Printf("%-25s %s %v\n", "User Status", ":", data[i].status)
	fmt.Printf("%-25s %s %d\n", "Jumlah Komponen Rusak", ":", data[i].jumlahKomponenRusak)
	footer()

}


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

func showUserStatus(data *dataBase, totalUser *int) {
	var countGood, countWarning, countCritical, countVeriCritical, countNone int

	for i := 1; i < *totalUser; i++ {
		if !data[i].dataSudahDiisi {
			countNone++
		} else {
			status := upperCaseConverter(data[i].status)
			fmt.Printf("\nUser %d: %s\n", i, data[i].user)
			if data[i].dataSudahDiisi {
				fmt.Printf("%-2sStatus: %s\n", "", data[i].status)
			} else {
				fmt.Printf("%-2sStatus: Data unavailable\n", "")
			}

			if status == "GUD" {
				countGood++
			} else if status == "WARNING" {
				countWarning++
			} else if status == "CRITICAL" {
				countCritical++
			} else if status == "VERY_CRITICAL" {
				countVeriCritical++
			}
		}
	}

	fmt.Printf("\nTOTAL USER STATUS STATISTICS\n\n")
	fmt.Printf("%-20s: %d\n", "Good", countGood)
	fmt.Printf("%-20s: %d\n", "Warning", countWarning)
	fmt.Printf("%-20s: %d\n", "Critical", countCritical)
	fmt.Printf("%-20s: %d\n", "Very Critical", countVeriCritical)
	fmt.Printf("%-20s: %d\n", "None", countNone)
}


func showAverageTemp(data *dataBase, totalUser *int) {
	var countCpu, countGpu, countRam int
	var totalCpuTemp, totalGpuTemp, totalRamTemp float64
	for i := 1; i < *totalUser; i++ {
		fmt.Printf("\nUser %d: %s\n", i, data[i].user)
		fmt.Printf("%-2sAverage CPU Temperature: %.2f°C\n", "", data[i].rataCpuTemp)
		totalCpuTemp = totalCpuTemp + data[i].rataCpuTemp
		countCpu++
		if data[i].gpuManufacturer != "NONE" {
			totalGpuTemp = totalGpuTemp + data[i].rataGpuTemp
			countGpu++
		}
		totalRamTemp = totalRamTemp + data[i].rataRamTemp
		countRam++
	}
	fmt.Printf("\n%-36s%s\n\n", "", "AVERAGE TEMPERATURE STATISTICS")
	if countCpu > 0 {
		fmt.Printf("%-20s: %.2f°C\n", "Average CPU Temperature", totalCpuTemp/float64(countCpu))
	} else {
		fmt.Printf("%-20s: %s\n", "Average CPU Temperature", "Data Unavailable")
	}
	if countGpu > 0 {
		fmt.Printf("%-20s: %.2f°C\n", "Average GPU Temperature", totalGpuTemp/float64(countGpu))
	} else {
		fmt.Printf("%-20s: %s\n", "Average GPU Temperature", "Data Unavailable")
	}
	if countRam > 0 {
		fmt.Printf("%-20s: %.2f°C\n", "Average RAM Temperature", totalRamTemp/float64(countRam))
	} else {
		fmt.Printf("%-20s: %s\n", "Average RAM Temperature", "Data Unavailable")
	}
	footer()
}


func showSortedCpuTemp(data *dataBase, totalUser *int) {
	if *totalUser <= 1 {
		fmt.Println("No user data available")
	} else {
		var ascending bool = jenisSorting()
		var dataCopy dataBase
		cloningData(data, &dataCopy, *totalUser)
		if ascending {
			selectionSort(&dataCopy, *totalUser, 1)
		} else {
			selectionSortDesc(&dataCopy, *totalUser, 1)
		}
		for i := 1; i < *totalUser; i++ {
			fmt.Printf("User: %s with Average CPU Temperature: %.2f°C\n", dataCopy[i].user, dataCopy[i].rataCpuTemp)
		}
	}
	footer()
}

func showSortedGpuTemp(data *dataBase, totalUser *int) {
	if *totalUser <= 1 {
		fmt.Println("No user data available")
	} else {
		var ascending bool = jenisSorting()
		var dataCopy dataBase
		cloningData(data, &dataCopy, *totalUser)
		if ascending {
			selectionSort(&dataCopy, *totalUser, 2)
		} else {
			selectionSortDesc(&dataCopy, *totalUser, 2)
		}
		for i := 1; i < *totalUser; i++ {
			if dataCopy[i].gpuManufacturer != "NONE" {
				fmt.Printf("User: %s with Average GPU Temperature: %.2f°C\n", dataCopy[i].user, dataCopy[i].rataGpuTemp)
			}
		}
	}
	footer()
}

func showSortedRamTemp(data *dataBase, totalUser *int) {
	if *totalUser <= 1 {
		fmt.Println("No user data available")
	} else {
		var ascending bool = jenisSorting()
		var dataCopy dataBase
		cloningData(data, &dataCopy, *totalUser)
		if ascending {
			insertionSortData(&dataCopy, *totalUser, 3)
		} else {
			insertionSortDataDesc(&dataCopy, *totalUser, 3)
		}
		for i := 1; i < *totalUser; i++ {
			if dataCopy[i].dataSudahDiisi {
				fmt.Printf("User: %s with Average RAM Temperature: %.2f°C\n", dataCopy[i].user, dataCopy[i].rataRamTemp)
			}
		}
	}
	footer()
}

func showSortedSerialCode(data *dataBase, totalUser *int) {
	if *totalUser <= 1 {
		fmt.Println("No user data available")
	} else {
		var ascending bool = jenisSorting()
		var dataCopy dataBase
		cloningData(data, &dataCopy, *totalUser)
		if ascending {
			insertionSortDataString(&dataCopy, *totalUser, 9)
		} else {
			insertionSortDataStringDesc(&dataCopy, *totalUser, 9)
		}
		for i := 1; i < *totalUser; i++ {
			if dataCopy[i].dataSudahDiisi {
				fmt.Printf("User: %s with Serial Code: %s\n", dataCopy[i].user, dataCopy[i].serialCode)
			}
		}
	}
	footer()
}

func showUserTempStats(data *dataBase, loggedInUser int) {
	fmt.Printf("\n%-40s%s\n\n", " ", "YOUR TEMPERATURE STATISTICS")
	footer()
	fmt.Printf("CPU:\n")
	fmt.Printf("  Avg: %.2f°C  Median: %.2f°C  Mode: %.2f°C  Min: %.2f°C  Max: %.2f°C\n",
		data[loggedInUser].rataCpuTemp, data[loggedInUser].medCpuTemp,
		data[loggedInUser].modCpuTemp, data[loggedInUser].minCpuTemp, data[loggedInUser].maxCpuTemp)
	if data[loggedInUser].gpuManufacturer != "NONE" {
		fmt.Printf("GPU:\n")
		fmt.Printf("  Avg: %.2f°C  Median: %.2f°C  Mode: %.2f°C  Min: %.2f°C  Max: %.2f°C\n",
			data[loggedInUser].rataGpuTemp, data[loggedInUser].medGpuTemp,
			data[loggedInUser].modGpuTemp, data[loggedInUser].minGpuTemp, data[loggedInUser].maxGpuTemp)
	}
	fmt.Printf("RAM:\n")
	fmt.Printf("  Avg: %.2f°C  Median: %.2f°C  Mode: %.2f°C  Min: %.2f°C  Max: %.2f°C\n",
		data[loggedInUser].rataRamTemp, data[loggedInUser].medRamTemp,
		data[loggedInUser].modRamTemp, data[loggedInUser].minRamTemp, data[loggedInUser].maxRamTemp)
	footer()
}
