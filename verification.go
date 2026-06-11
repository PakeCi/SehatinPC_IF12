package main

func validateUserLogin(data *dataBase, username, password string, userIndex int) int {
	for i := 0; i < userIndex; i++ {
		if data[i].user == username && data[i].userPassword == password {
			return i
		}
	}
	return -1
}

func checkValidityUser(data *dataBase, username string, userIndex int) bool {
	for i := 0; i < userIndex; i++ {
		if data[i].user == username {
			return false
		}
	}
	return true
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
			if x[0] != 'M' || len(x) == 0 {
				return false
			}
		}
	} else if i == 6 {
		//formatnya DD-MM-YYYY
		if len(x) != 10 || ((x[0] < '0' || x[0] > '3') || (x[1] < '0' || x[1] > '9') || (x[3] < '0' || x[3] > '1') || (x[4] < '0' || x[4] > '9') || (x[6] < '0' || x[6] > '9') || (x[7] < '0' || x[7] > '9') || (x[8] < '0' || x[8] > '9') || (x[9] < '0' || x[9] > '9') || x[2] != '-' || x[5] != '-') {
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

