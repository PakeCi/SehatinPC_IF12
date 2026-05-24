package main

import "fmt"

const NMAX = 9999

type dataComponent struct {
	user, userPassword, serialCode                 string
	batteryHealth                                  float64 // dalam Persentase
	cpuManufacturer, gpuManufacturer               string // CPU : Intel, AMD, Apple M-series GPU : Nvidia, AMD, Apple or NONE
	cpuModel, gpuModel                             string // CPU : Core, Ryzen, Series, GPU : GeForce, Radeon, Apple or NONE
	rataCpuTemp, rataGpuTemp, rataRamTemp float64 // dalam Celcius
	medCpuTemp, medGpuTemp, medRamTemp float64 // dalam Celcius
	modCpuTemp, modGpuTemp, modRamTemp float64 // dalam Celcius
	ramCapacity, ramUsed, diskCapacity, diskUsed   float64 // dalam GiB
	dataLoad, dataSudahDiisi bool // true = data dalam heavy load, false = data dalam idle load
	operatingSystem string // Windows, Linux, MacOS
	lastMaintenanceDate, nextMaintenanceDate       string //Format DD-MM-YYYY
	status                                         string // Gud, Warning, Critical
}

type dataBase [NMAX]dataComponent 

func header(){
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
	Line := ("______________________________________________________________________________________________________") //102 karakter
	fmt.Printf("%s\n", Line)
}

func checkAvailabilityUser(data dataBase, username string, userIndex int) bool {
	var valid bool

	valid = true
	for i := 0; i < userIndex; i++ {
		if data[i].user == username {
			valid = false
			i = userIndex 
		}
	}
	return valid
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

func loginPage(data *dataBase, userIndex *int, login *bool) {

}

func registerPage(data *dataBase, userIndex *int){

}

func loginMenu(data *dataBase, kill *bool, login *bool, userIndex *int, firstOpenPage *bool){
	var input int 
	if *firstOpenPage {
		header()
		*firstOpenPage = false
	}
	fmt.Printf("%-46s%s\n"," ","Login Menu") // 1/2*102 = 51 
	fmt.Printf("%-45s%s\n%-45s%s\n%-45s%s\n", " ", "1. Register"," ", "2. Login"," ", "3. Kill Program")
	footer()
	fmt.Print("Input : ")
	fmt.Scan(&input)

	if input == 1 {
		registerPage(data, userIndex)
	}else if input == 2 {
		loginPage(data, userIndex, login)
	}else if input == 3 {
		*kill = true
	}else {
		fmt.Println("Input Invalid")
	}
}

func main(){
	var data dataBase
	var exit bool = false
	var login bool = false
	var firstOpenPage bool = true
	var userIndex int

	for exit == false { 
		if login == false {
			loginMenu(&data, &exit, &login, &userIndex, &firstOpenPage)
			if login == true {
				firstOpenPage = true
			}
		} else {
			mainMenu(&data, &exit, &login, &userIndex, &firstOpenPage)
			if login == false {
				firstOpenPage = true
			}
		}
	}
}

func mainMenu(data *dataBase, kill *bool, login *bool, userIndex *int, firstOpenPage *bool) {
}