package main 

import "fmt"

func binarySearch(data *dataBase, searchData float64, totalUser *int, id int, batasKanan, batasKiri *int) {
	var right, left, middle int
	// var dataCopy dataBase
	// ini buat float64
	selectionSort(data, *totalUser, id)
	*batasKiri = -1
	*batasKanan = -1
	right = *totalUser - 1
	left = 1
	for left <= right {
		middle = (right + left) / 2
		if data[middle].indexFloat(id) > searchData {
			right = middle - 1
		} else if data[middle].indexFloat(id) < searchData {
			left = middle + 1
		} else {
			*batasKiri = middle
			*batasKanan = middle
			left = right + 1
		}
	}

	if *batasKiri != -1 && *batasKanan != -1 {

		for data[*batasKiri-1].indexFloat(id) == searchData && *batasKiri > 1 {
			*batasKiri = *batasKiri - 1
		}

		for data[*batasKanan+1].indexFloat(id) == searchData && *batasKanan < *totalUser-1 {
			*batasKanan = *batasKanan + 1
		}
	}
}


func sequentialSearch(data *dataBase, searchData string, totalUser *int, id int) {
	//buat cari status atau yang lainnya berdasarkan string
	for i := 1; i < *totalUser; i++ {
		if data[i].indexString(id) == searchData {
			fmt.Printf("User: %s with data: %s \n", data[i].user, data[i].indexString(id))
		}
	}
}

func sequentialSearchIndex(data *dataBase, searchData string, totalUser *int) int {
	for i := 1; i < *totalUser; i++ {
		if data[i].user == searchData {
			return i
		}
	}
	return -1
}

// gah damn dikit lg co
func searchUserStatus(data *dataBase, totalUser *int) {
	var searchStatus string
	// var input string
	if *totalUser <= 1 {
		fmt.Println("No user data available")
	} else {
		fmt.Print("Input Status that you want to search (GUD / WARNING / CRITICAL / VERY_CRITICAL):")
		fmt.Scan(&searchStatus)
		searchStatus = upperCaseConverter(searchStatus)
		sequentialSearch(data, searchStatus, totalUser, 8)
	}
}

func searchUserCpuTemp(data *dataBase, totalUser *int) {
	if *totalUser <= 1 {
		fmt.Println("No user data available")
	} else {
		var batasKiri, batasKanan int
		var input float64
		fmt.Print("Input CPU that you want to search: ")
		fmt.Scan(&input)
		var dataCopy dataBase
		cloningData(data, &dataCopy, *totalUser)
		binarySearch(&dataCopy, input, totalUser, 1, &batasKanan, &batasKiri)
		if batasKiri == -1 || batasKanan == -1 {
			fmt.Println("No user data found with that CPU Temperature")
		} else {
			fmt.Printf("Users with Average CPU Temperature of %.2f°C:\n", input)
			for i := batasKiri + 1; i < batasKanan; i++ {
				fmt.Printf("User: %s\n", dataCopy[i].user)
				fmt.Printf("%-2sAverage CPU Temperature: %.2f°C\n", "", dataCopy[i].rataCpuTemp)
			}
		}
	}
	footer()
}