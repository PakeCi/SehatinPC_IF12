package main

func seedDummyData(data *dataBase, userIndex *int) {

	// ============================================================
	// Helper: hitung rata-rata, median, modus, min, max dari slice
	// ============================================================
	calcStats := func(temps [10]float64) (rata, med, mod, mn, mx float64) {
		rata = searchRataRata(temps)
		med = searchMedian(temps)
		mod = searchModus(temps)
		searchMinMax(temps, &mn, &mx)
		return
	}

	// ============================================================
	// User 1 — andi_gunawan | Windows | Intel Core | NVIDIA RTX
	// Status  : GUD
	// ============================================================
	i := *userIndex
	data[i].user = "andi_gunawan"
	data[i].userPassword = "andi1234"
	data[i].serialCode = "SN-001-ANDI"
	data[i].usingLaptop = true
	data[i].batteryHealth = 87.5
	data[i].cpuManufacturer = "INTEL"
	data[i].cpuModel = "CORE"
	data[i].cpuSerial = "i7-12700H"
	data[i].gpuManufacturer = "NVIDIA"
	data[i].gpuModel = "RTX"
	data[i].gpuSerial = "RTX 3060"
	data[i].ramCapacity = 16
	data[i].ramUsed = 7.2
	data[i].diskCapacity = 512
	data[i].diskUsed = 210.0
	data[i].operatingSystem = "WINDOWS"
	data[i].dataLoad = false
	data[i].lastMaintenanceDate = "10-01-2025"
	cpuT := [10]float64{55, 57, 56, 55, 58, 57, 56, 55, 57, 56}
	gpuT := [10]float64{52, 53, 54, 52, 53, 54, 52, 53, 54, 53}
	ramT := [10]float64{40, 41, 40, 41, 40, 41, 40, 41, 40, 41}
	data[i].rataCpuTemp, data[i].medCpuTemp, data[i].modCpuTemp, data[i].minCpuTemp, data[i].maxCpuTemp = calcStats(cpuT)
	data[i].rataGpuTemp, data[i].medGpuTemp, data[i].modGpuTemp, data[i].minGpuTemp, data[i].maxGpuTemp = calcStats(gpuT)
	data[i].rataRamTemp, data[i].medRamTemp, data[i].modRamTemp, data[i].minRamTemp, data[i].maxRamTemp = calcStats(ramT)
	data[i].dataSudahDiisi = true
	setData(data, i)
	*userIndex++

	// ============================================================
	// User 2 — siti_rahayu | MacOS | Apple M | Apple GPU
	// Status  : GUD
	// ============================================================
	i = *userIndex
	data[i].user = "siti_rahayu"
	data[i].userPassword = "siti5678"
	data[i].serialCode = "SN-002-SITI"
	data[i].usingLaptop = true
	data[i].batteryHealth = 92.0
	data[i].cpuManufacturer = "APPLE"
	data[i].cpuModel = "M"
	data[i].cpuSerial = "M2"
	data[i].gpuManufacturer = "APPLE"
	data[i].gpuModel = "M"
	data[i].gpuSerial = "M2 GPU"
	data[i].ramCapacity = 16
	data[i].ramUsed = 5.8
	data[i].diskCapacity = 256
	data[i].diskUsed = 80.0
	data[i].operatingSystem = "MACOS"
	data[i].dataLoad = false
	data[i].lastMaintenanceDate = "15-02-2025"
	cpuT = [10]float64{45, 46, 45, 47, 46, 45, 46, 45, 47, 46}
	gpuT = [10]float64{43, 44, 43, 44, 43, 44, 43, 44, 43, 44}
	ramT = [10]float64{38, 38, 39, 38, 38, 39, 38, 38, 39, 38}
	data[i].rataCpuTemp, data[i].medCpuTemp, data[i].modCpuTemp, data[i].minCpuTemp, data[i].maxCpuTemp = calcStats(cpuT)
	data[i].rataGpuTemp, data[i].medGpuTemp, data[i].modGpuTemp, data[i].minGpuTemp, data[i].maxGpuTemp = calcStats(gpuT)
	data[i].rataRamTemp, data[i].medRamTemp, data[i].modRamTemp, data[i].minRamTemp, data[i].maxRamTemp = calcStats(ramT)
	data[i].dataSudahDiisi = true
	setData(data, i)
	*userIndex++

	// ============================================================
	// User 3 — budi_santoso | Windows | AMD Ryzen | NVIDIA GTX
	// Status  : CRITICAL (CPU overheat saat heavy load)
	// ============================================================
	i = *userIndex
	data[i].user = "budi_santoso"
	data[i].userPassword = "budi9999"
	data[i].serialCode = "SN-003-BUDI"
	data[i].usingLaptop = false
	data[i].batteryHealth = 0
	data[i].cpuManufacturer = "AMD"
	data[i].cpuModel = "RYZEN"
	data[i].cpuSerial = "Ryzen 5 5600X"
	data[i].gpuManufacturer = "NVIDIA"
	data[i].gpuModel = "GTX"
	data[i].gpuSerial = "GTX 1660 Super"
	data[i].ramCapacity = 32
	data[i].ramUsed = 18.0
	data[i].diskCapacity = 1024
	data[i].diskUsed = 600.0
	data[i].operatingSystem = "WINDOWS"
	data[i].dataLoad = true
	data[i].lastMaintenanceDate = "20-03-2025"
	cpuT = [10]float64{96, 97, 96, 98, 97, 96, 97, 96, 98, 97}
	gpuT = [10]float64{70, 71, 70, 72, 71, 70, 71, 70, 72, 71}
	ramT = [10]float64{60, 61, 60, 62, 61, 60, 61, 60, 62, 61}
	data[i].rataCpuTemp, data[i].medCpuTemp, data[i].modCpuTemp, data[i].minCpuTemp, data[i].maxCpuTemp = calcStats(cpuT)
	data[i].rataGpuTemp, data[i].medGpuTemp, data[i].modGpuTemp, data[i].minGpuTemp, data[i].maxGpuTemp = calcStats(gpuT)
	data[i].rataRamTemp, data[i].medRamTemp, data[i].modRamTemp, data[i].minRamTemp, data[i].maxRamTemp = calcStats(ramT)
	data[i].dataSudahDiisi = true
	setData(data, i)
	*userIndex++

	// ============================================================
	// User 4 — dewi_lestari | Linux | Intel Xeon | NVIDIA RTX
	// Status  : GUD
	// ============================================================
	i = *userIndex
	data[i].user = "dewi_lestari"
	data[i].userPassword = "dewi2024"
	data[i].serialCode = "SN-004-DEWI"
	data[i].usingLaptop = false
	data[i].batteryHealth = 0
	data[i].cpuManufacturer = "INTEL"
	data[i].cpuModel = "XEON"
	data[i].cpuSerial = "Xeon E5-2680"
	data[i].gpuManufacturer = "NVIDIA"
	data[i].gpuModel = "RTX"
	data[i].gpuSerial = "RTX 4090"
	data[i].ramCapacity = 64
	data[i].ramUsed = 20.0
	data[i].diskCapacity = 2048
	data[i].diskUsed = 400.0
	data[i].operatingSystem = "LINUX"
	data[i].dataLoad = false
	data[i].lastMaintenanceDate = "05-04-2025"
	cpuT = [10]float64{60, 62, 61, 60, 63, 61, 60, 62, 61, 60}
	gpuT = [10]float64{55, 56, 55, 57, 56, 55, 56, 55, 57, 56}
	ramT = [10]float64{42, 43, 42, 43, 42, 43, 42, 43, 42, 43}
	data[i].rataCpuTemp, data[i].medCpuTemp, data[i].modCpuTemp, data[i].minCpuTemp, data[i].maxCpuTemp = calcStats(cpuT)
	data[i].rataGpuTemp, data[i].medGpuTemp, data[i].modGpuTemp, data[i].minGpuTemp, data[i].maxGpuTemp = calcStats(gpuT)
	data[i].rataRamTemp, data[i].medRamTemp, data[i].modRamTemp, data[i].minRamTemp, data[i].maxRamTemp = calcStats(ramT)
	data[i].dataSudahDiisi = true
	setData(data, i)
	*userIndex++

	// ============================================================
	// User 5 — rizky_firmansyah | Windows | Intel Core | AMD Radeon
	// Status  : WARNING (disk hampir penuh)
	// ============================================================
	i = *userIndex
	data[i].user = "rizky_firmansyah"
	data[i].userPassword = "rizky4321"
	data[i].serialCode = "SN-005-RIZKY"
	data[i].usingLaptop = true
	data[i].batteryHealth = 65.0
	data[i].cpuManufacturer = "INTEL"
	data[i].cpuModel = "CORE"
	data[i].cpuSerial = "i5-1135G7"
	data[i].gpuManufacturer = "AMD"
	data[i].gpuModel = "RADEON"
	data[i].gpuSerial = "Radeon RX 6600"
	data[i].ramCapacity = 8
	data[i].ramUsed = 4.0
	data[i].diskCapacity = 256
	data[i].diskUsed = 240.0 // > 85% kapasitas → WARNING
	data[i].operatingSystem = "WINDOWS"
	data[i].dataLoad = false
	data[i].lastMaintenanceDate = "01-05-2025"
	cpuT = [10]float64{65, 66, 65, 67, 66, 65, 66, 65, 67, 66}
	gpuT = [10]float64{60, 61, 60, 62, 61, 60, 61, 60, 62, 61}
	ramT = [10]float64{50, 51, 50, 52, 51, 50, 51, 50, 52, 51}
	data[i].rataCpuTemp, data[i].medCpuTemp, data[i].modCpuTemp, data[i].minCpuTemp, data[i].maxCpuTemp = calcStats(cpuT)
	data[i].rataGpuTemp, data[i].medGpuTemp, data[i].modGpuTemp, data[i].minGpuTemp, data[i].maxGpuTemp = calcStats(gpuT)
	data[i].rataRamTemp, data[i].medRamTemp, data[i].modRamTemp, data[i].minRamTemp, data[i].maxRamTemp = calcStats(ramT)
	data[i].dataSudahDiisi = true
	setData(data, i)
	*userIndex++

	// ============================================================
	// User 6 — nurul_hidayah | MacOS | Apple M | Apple GPU
	// Status  : GUD
	// ============================================================
	i = *userIndex
	data[i].user = "nurul_hidayah"
	data[i].userPassword = "nurul8888"
	data[i].serialCode = "SN-006-NURUL"
	data[i].usingLaptop = true
	data[i].batteryHealth = 95.0
	data[i].cpuManufacturer = "APPLE"
	data[i].cpuModel = "M"
	data[i].cpuSerial = "M3 Pro"
	data[i].gpuManufacturer = "APPLE"
	data[i].gpuModel = "M"
	data[i].gpuSerial = "M3 Pro GPU"
	data[i].ramCapacity = 18
	data[i].ramUsed = 6.0
	data[i].diskCapacity = 512
	data[i].diskUsed = 150.0
	data[i].operatingSystem = "MACOS"
	data[i].dataLoad = false
	data[i].lastMaintenanceDate = "12-06-2025"
	cpuT = [10]float64{40, 41, 40, 42, 41, 40, 41, 40, 42, 41}
	gpuT = [10]float64{38, 39, 38, 40, 39, 38, 39, 38, 40, 39}
	ramT = [10]float64{35, 36, 35, 37, 36, 35, 36, 35, 37, 36}
	data[i].rataCpuTemp, data[i].medCpuTemp, data[i].modCpuTemp, data[i].minCpuTemp, data[i].maxCpuTemp = calcStats(cpuT)
	data[i].rataGpuTemp, data[i].medGpuTemp, data[i].modGpuTemp, data[i].minGpuTemp, data[i].maxGpuTemp = calcStats(gpuT)
	data[i].rataRamTemp, data[i].medRamTemp, data[i].modRamTemp, data[i].minRamTemp, data[i].maxRamTemp = calcStats(ramT)
	data[i].dataSudahDiisi = true
	setData(data, i)
	*userIndex++

	// ============================================================
	// User 7 — fajar_pratama | Linux | AMD Ryzen | NVIDIA RTX
	// Status  : VERY_CRITICAL (CPU + GPU + RAM overheat)
	// ============================================================
	i = *userIndex
	data[i].user = "fajar_pratama"
	data[i].userPassword = "fajar7777"
	data[i].serialCode = "SN-007-FAJAR"
	data[i].usingLaptop = false
	data[i].batteryHealth = 0
	data[i].cpuManufacturer = "AMD"
	data[i].cpuModel = "RYZEN"
	data[i].cpuSerial = "Ryzen 9 7950X"
	data[i].gpuManufacturer = "NVIDIA"
	data[i].gpuModel = "RTX"
	data[i].gpuSerial = "RTX 4080"
	data[i].ramCapacity = 32
	data[i].ramUsed = 28.0
	data[i].diskCapacity = 1024
	data[i].diskUsed = 500.0
	data[i].operatingSystem = "LINUX"
	data[i].dataLoad = true
	data[i].lastMaintenanceDate = "18-07-2025"
	cpuT = [10]float64{96, 97, 96, 98, 96, 97, 96, 97, 96, 98}  // rata >95 → overheat
	gpuT = [10]float64{87, 88, 87, 89, 87, 88, 87, 88, 87, 89}  // rata >85 → overheat
	ramT = [10]float64{87, 88, 87, 89, 87, 88, 87, 88, 87, 89}  // rata >85 → overheat
	data[i].rataCpuTemp, data[i].medCpuTemp, data[i].modCpuTemp, data[i].minCpuTemp, data[i].maxCpuTemp = calcStats(cpuT)
	data[i].rataGpuTemp, data[i].medGpuTemp, data[i].modGpuTemp, data[i].minGpuTemp, data[i].maxGpuTemp = calcStats(gpuT)
	data[i].rataRamTemp, data[i].medRamTemp, data[i].modRamTemp, data[i].minRamTemp, data[i].maxRamTemp = calcStats(ramT)
	data[i].dataSudahDiisi = true
	setData(data, i)
	*userIndex++

	// ============================================================
	// User 8 — maya_setiawati | Windows | Intel Pentium | NONE
	// Status  : GUD
	// ============================================================
	i = *userIndex
	data[i].user = "maya_setiawati"
	data[i].userPassword = "maya2020"
	data[i].serialCode = "SN-008-MAYA"
	data[i].usingLaptop = false
	data[i].batteryHealth = 0
	data[i].cpuManufacturer = "INTEL"
	data[i].cpuModel = "PENTIUM"
	data[i].cpuSerial = "Pentium G6400"
	data[i].gpuManufacturer = "NONE"
	data[i].gpuModel = "NONE"
	data[i].gpuSerial = "NONE"
	data[i].ramCapacity = 8
	data[i].ramUsed = 2.5
	data[i].diskCapacity = 500
	data[i].diskUsed = 150.0
	data[i].operatingSystem = "WINDOWS"
	data[i].dataLoad = false
	data[i].lastMaintenanceDate = "22-08-2025"
	cpuT = [10]float64{55, 55, 56, 55, 56, 55, 56, 55, 55, 56}
	ramT = [10]float64{40, 41, 40, 41, 40, 41, 40, 41, 40, 41}
	data[i].rataCpuTemp, data[i].medCpuTemp, data[i].modCpuTemp, data[i].minCpuTemp, data[i].maxCpuTemp = calcStats(cpuT)
	data[i].rataGpuTemp, data[i].medGpuTemp, data[i].modGpuTemp, data[i].minGpuTemp, data[i].maxGpuTemp = 0, 0, 0, 0, 0
	data[i].rataRamTemp, data[i].medRamTemp, data[i].modRamTemp, data[i].minRamTemp, data[i].maxRamTemp = calcStats(ramT)
	data[i].dataSudahDiisi = true
	setData(data, i)
	*userIndex++

	// ============================================================
	// User 9 — hendra_kusuma | Linux | AMD Athlon | NONE
	// Status  : GUD
	// ============================================================
	i = *userIndex
	data[i].user = "hendra_kusuma"
	data[i].userPassword = "hendra3030"
	data[i].serialCode = "SN-009-HENDRA"
	data[i].usingLaptop = false
	data[i].batteryHealth = 0
	data[i].cpuManufacturer = "AMD"
	data[i].cpuModel = "ATHLON"
	data[i].cpuSerial = "Athlon 3000G"
	data[i].gpuManufacturer = "NONE"
	data[i].gpuModel = "NONE"
	data[i].gpuSerial = "NONE"
	data[i].ramCapacity = 16
	data[i].ramUsed = 4.0
	data[i].diskCapacity = 500
	data[i].diskUsed = 120.0
	data[i].operatingSystem = "LINUX"
	data[i].dataLoad = false
	data[i].lastMaintenanceDate = "03-09-2025"
	cpuT = [10]float64{50, 51, 50, 52, 51, 50, 51, 50, 52, 51}
	ramT = [10]float64{38, 39, 38, 40, 39, 38, 39, 38, 40, 39}
	data[i].rataCpuTemp, data[i].medCpuTemp, data[i].modCpuTemp, data[i].minCpuTemp, data[i].maxCpuTemp = calcStats(cpuT)
	data[i].rataGpuTemp, data[i].medGpuTemp, data[i].modGpuTemp, data[i].minGpuTemp, data[i].maxGpuTemp = 0, 0, 0, 0, 0
	data[i].rataRamTemp, data[i].medRamTemp, data[i].modRamTemp, data[i].minRamTemp, data[i].maxRamTemp = calcStats(ramT)
	data[i].dataSudahDiisi = true
	setData(data, i)
	*userIndex++

	// ============================================================
	// User 10 — putri_anggraini | MacOS | Apple M | Apple GPU
	// Status  : WARNING (RAM hampir penuh)
	// ============================================================
	i = *userIndex
	data[i].user = "putri_anggraini"
	data[i].userPassword = "putri1111"
	data[i].serialCode = "SN-010-PUTRI"
	data[i].usingLaptop = true
	data[i].batteryHealth = 78.0
	data[i].cpuManufacturer = "APPLE"
	data[i].cpuModel = "M"
	data[i].cpuSerial = "M1"
	data[i].gpuManufacturer = "APPLE"
	data[i].gpuModel = "M"
	data[i].gpuSerial = "M1 GPU"
	data[i].ramCapacity = 8
	data[i].ramUsed = 7.5 // > 85% kapasitas → WARNING
	data[i].diskCapacity = 256
	data[i].diskUsed = 100.0
	data[i].operatingSystem = "MACOS"
	data[i].dataLoad = false
	data[i].lastMaintenanceDate = "10-10-2025"
	cpuT = [10]float64{48, 49, 48, 50, 49, 48, 49, 48, 50, 49}
	gpuT = [10]float64{45, 46, 45, 47, 46, 45, 46, 45, 47, 46}
	ramT = [10]float64{42, 43, 42, 44, 43, 42, 43, 42, 44, 43}
	data[i].rataCpuTemp, data[i].medCpuTemp, data[i].modCpuTemp, data[i].minCpuTemp, data[i].maxCpuTemp = calcStats(cpuT)
	data[i].rataGpuTemp, data[i].medGpuTemp, data[i].modGpuTemp, data[i].minGpuTemp, data[i].maxGpuTemp = calcStats(gpuT)
	data[i].rataRamTemp, data[i].medRamTemp, data[i].modRamTemp, data[i].minRamTemp, data[i].maxRamTemp = calcStats(ramT)
	data[i].dataSudahDiisi = true
	setData(data, i)
	*userIndex++

	// ============================================================
	// User 11 — taufik_rahman | Windows | Intel Atom | NONE
	// Status  : CRITICAL (CPU overheat saat idle)
	// ============================================================
	i = *userIndex
	data[i].user = "taufik_rahman"
	data[i].userPassword = "taufik5555"
	data[i].serialCode = "SN-011-TAUFIK"
	data[i].usingLaptop = true
	data[i].batteryHealth = 50.0
	data[i].cpuManufacturer = "INTEL"
	data[i].cpuModel = "ATOM"
	data[i].cpuSerial = "Atom x7-E3950"
	data[i].gpuManufacturer = "NONE"
	data[i].gpuModel = "NONE"
	data[i].gpuSerial = "NONE"
	data[i].ramCapacity = 4
	data[i].ramUsed = 2.0
	data[i].diskCapacity = 128
	data[i].diskUsed = 60.0
	data[i].operatingSystem = "WINDOWS"
	data[i].dataLoad = false
	data[i].lastMaintenanceDate = "25-11-2024"
	cpuT = [10]float64{66, 67, 66, 68, 67, 66, 67, 66, 68, 67} // rata >65 idle → overheat
	ramT = [10]float64{50, 51, 50, 52, 51, 50, 51, 50, 52, 51}
	data[i].rataCpuTemp, data[i].medCpuTemp, data[i].modCpuTemp, data[i].minCpuTemp, data[i].maxCpuTemp = calcStats(cpuT)
	data[i].rataGpuTemp, data[i].medGpuTemp, data[i].modGpuTemp, data[i].minGpuTemp, data[i].maxGpuTemp = 0, 0, 0, 0, 0
	data[i].rataRamTemp, data[i].medRamTemp, data[i].modRamTemp, data[i].minRamTemp, data[i].maxRamTemp = calcStats(ramT)
	data[i].dataSudahDiisi = true
	setData(data, i)
	*userIndex++

	// ============================================================
	// User 12 — rini_wahyuni | Linux | Intel Core | AMD Radeon
	// Status  : GUD
	// ============================================================
	i = *userIndex
	data[i].user = "rini_wahyuni"
	data[i].userPassword = "rini6060"
	data[i].serialCode = "SN-012-RINI"
	data[i].usingLaptop = false
	data[i].batteryHealth = 0
	data[i].cpuManufacturer = "INTEL"
	data[i].cpuModel = "CORE"
	data[i].cpuSerial = "i9-13900K"
	data[i].gpuManufacturer = "AMD"
	data[i].gpuModel = "RADEON"
	data[i].gpuSerial = "Radeon RX 7900 XT"
	data[i].ramCapacity = 32
	data[i].ramUsed = 10.0
	data[i].diskCapacity = 2048
	data[i].diskUsed = 800.0
	data[i].operatingSystem = "LINUX"
	data[i].dataLoad = true
	data[i].lastMaintenanceDate = "30-12-2024"
	cpuT = [10]float64{75, 76, 75, 77, 76, 75, 76, 75, 77, 76}
	gpuT = [10]float64{68, 69, 68, 70, 69, 68, 69, 68, 70, 69}
	ramT = [10]float64{58, 59, 58, 60, 59, 58, 59, 58, 60, 59}
	data[i].rataCpuTemp, data[i].medCpuTemp, data[i].modCpuTemp, data[i].minCpuTemp, data[i].maxCpuTemp = calcStats(cpuT)
	data[i].rataGpuTemp, data[i].medGpuTemp, data[i].modGpuTemp, data[i].minGpuTemp, data[i].maxGpuTemp = calcStats(gpuT)
	data[i].rataRamTemp, data[i].medRamTemp, data[i].modRamTemp, data[i].minRamTemp, data[i].maxRamTemp = calcStats(ramT)
	data[i].dataSudahDiisi = true
	setData(data, i)
	*userIndex++

	// ============================================================
	// User 13 — agus_hermawan | Windows | AMD Ryzen | NVIDIA RTX
	// Status  : GUD
	// ============================================================
	i = *userIndex
	data[i].user = "agus_hermawan"
	data[i].userPassword = "agus1212"
	data[i].serialCode = "SN-013-AGUS"
	data[i].usingLaptop = true
	data[i].batteryHealth = 82.0
	data[i].cpuManufacturer = "AMD"
	data[i].cpuModel = "RYZEN"
	data[i].cpuSerial = "Ryzen 7 6800H"
	data[i].gpuManufacturer = "NVIDIA"
	data[i].gpuModel = "RTX"
	data[i].gpuSerial = "RTX 3070 Ti"
	data[i].ramCapacity = 16
	data[i].ramUsed = 8.0
	data[i].diskCapacity = 512
	data[i].diskUsed = 200.0
	data[i].operatingSystem = "WINDOWS"
	data[i].dataLoad = false
	data[i].lastMaintenanceDate = "05-01-2025"
	cpuT = [10]float64{62, 63, 62, 64, 63, 62, 63, 62, 64, 63}
	gpuT = [10]float64{58, 59, 58, 60, 59, 58, 59, 58, 60, 59}
	ramT = [10]float64{45, 46, 45, 47, 46, 45, 46, 45, 47, 46}
	data[i].rataCpuTemp, data[i].medCpuTemp, data[i].modCpuTemp, data[i].minCpuTemp, data[i].maxCpuTemp = calcStats(cpuT)
	data[i].rataGpuTemp, data[i].medGpuTemp, data[i].modGpuTemp, data[i].minGpuTemp, data[i].maxGpuTemp = calcStats(gpuT)
	data[i].rataRamTemp, data[i].medRamTemp, data[i].modRamTemp, data[i].minRamTemp, data[i].maxRamTemp = calcStats(ramT)
	data[i].dataSudahDiisi = true
	setData(data, i)
	*userIndex++

	// ============================================================
	// User 14 — laila_nurfitri | MacOS | Apple M | Apple GPU
	// Status  : GUD
	// ============================================================
	i = *userIndex
	data[i].user = "laila_nurfitri"
	data[i].userPassword = "laila3333"
	data[i].serialCode = "SN-014-LAILA"
	data[i].usingLaptop = true
	data[i].batteryHealth = 88.0
	data[i].cpuManufacturer = "APPLE"
	data[i].cpuModel = "M"
	data[i].cpuSerial = "M2 Pro"
	data[i].gpuManufacturer = "APPLE"
	data[i].gpuModel = "M"
	data[i].gpuSerial = "M2 Pro GPU"
	data[i].ramCapacity = 16
	data[i].ramUsed = 7.0
	data[i].diskCapacity = 512
	data[i].diskUsed = 200.0
	data[i].operatingSystem = "MACOS"
	data[i].dataLoad = false
	data[i].lastMaintenanceDate = "14-02-2025"
	cpuT = [10]float64{44, 45, 44, 46, 45, 44, 45, 44, 46, 45}
	gpuT = [10]float64{42, 43, 42, 44, 43, 42, 43, 42, 44, 43}
	ramT = [10]float64{36, 37, 36, 38, 37, 36, 37, 36, 38, 37}
	data[i].rataCpuTemp, data[i].medCpuTemp, data[i].modCpuTemp, data[i].minCpuTemp, data[i].maxCpuTemp = calcStats(cpuT)
	data[i].rataGpuTemp, data[i].medGpuTemp, data[i].modGpuTemp, data[i].minGpuTemp, data[i].maxGpuTemp = calcStats(gpuT)
	data[i].rataRamTemp, data[i].medRamTemp, data[i].modRamTemp, data[i].minRamTemp, data[i].maxRamTemp = calcStats(ramT)
	data[i].dataSudahDiisi = true
	setData(data, i)
	*userIndex++

	// ============================================================
	// User 15 — doni_prasetyo | Windows | Intel Core | NVIDIA MAX-Q
	// Status  : CRITICAL (GPU overheat saat heavy load)
	// ============================================================
	i = *userIndex
	data[i].user = "doni_prasetyo"
	data[i].userPassword = "doni7070"
	data[i].serialCode = "SN-015-DONI"
	data[i].usingLaptop = true
	data[i].batteryHealth = 70.0
	data[i].cpuManufacturer = "INTEL"
	data[i].cpuModel = "CORE"
	data[i].cpuSerial = "i7-11800H"
	data[i].gpuManufacturer = "NVIDIA"
	data[i].gpuModel = "MAX-Q"
	data[i].gpuSerial = "RTX 3060 Max-Q"
	data[i].ramCapacity = 16
	data[i].ramUsed = 9.0
	data[i].diskCapacity = 512
	data[i].diskUsed = 300.0
	data[i].operatingSystem = "WINDOWS"
	data[i].dataLoad = true
	data[i].lastMaintenanceDate = "20-03-2025"
	cpuT = [10]float64{88, 89, 88, 90, 89, 88, 89, 88, 90, 89}
	gpuT = [10]float64{86, 87, 86, 88, 87, 86, 87, 86, 88, 87} // rata >85 → overheat
	ramT = [10]float64{60, 61, 60, 62, 61, 60, 61, 60, 62, 61}
	data[i].rataCpuTemp, data[i].medCpuTemp, data[i].modCpuTemp, data[i].minCpuTemp, data[i].maxCpuTemp = calcStats(cpuT)
	data[i].rataGpuTemp, data[i].medGpuTemp, data[i].modGpuTemp, data[i].minGpuTemp, data[i].maxGpuTemp = calcStats(gpuT)
	data[i].rataRamTemp, data[i].medRamTemp, data[i].modRamTemp, data[i].minRamTemp, data[i].maxRamTemp = calcStats(ramT)
	data[i].dataSudahDiisi = true
	setData(data, i)
	*userIndex++

	// ============================================================
	// User 16 — yuni_astuti | Linux | AMD Epyc | NONE
	// Status  : GUD
	// ============================================================
	i = *userIndex
	data[i].user = "yuni_astuti"
	data[i].userPassword = "yuni4040"
	data[i].serialCode = "SN-016-YUNI"
	data[i].usingLaptop = false
	data[i].batteryHealth = 0
	data[i].cpuManufacturer = "AMD"
	data[i].cpuModel = "EPYC"
	data[i].cpuSerial = "EPYC 7302"
	data[i].gpuManufacturer = "NONE"
	data[i].gpuModel = "NONE"
	data[i].gpuSerial = "NONE"
	data[i].ramCapacity = 128
	data[i].ramUsed = 40.0
	data[i].diskCapacity = 4096
	data[i].diskUsed = 1000.0
	data[i].operatingSystem = "LINUX"
	data[i].dataLoad = false
	data[i].lastMaintenanceDate = "08-04-2025"
	cpuT = [10]float64{70, 71, 70, 72, 71, 70, 71, 70, 72, 71}
	ramT = [10]float64{55, 56, 55, 57, 56, 55, 56, 55, 57, 56}
	data[i].rataCpuTemp, data[i].medCpuTemp, data[i].modCpuTemp, data[i].minCpuTemp, data[i].maxCpuTemp = calcStats(cpuT)
	data[i].rataGpuTemp, data[i].medGpuTemp, data[i].modGpuTemp, data[i].minGpuTemp, data[i].maxGpuTemp = 0, 0, 0, 0, 0
	data[i].rataRamTemp, data[i].medRamTemp, data[i].modRamTemp, data[i].minRamTemp, data[i].maxRamTemp = calcStats(ramT)
	data[i].dataSudahDiisi = true
	setData(data, i)
	*userIndex++

	// ============================================================
	// User 17 — bagas_nugroho | Windows | Intel Core | NVIDIA GTX
	// Status  : WARNING (disk hampir penuh)
	// ============================================================
	i = *userIndex
	data[i].user = "bagas_nugroho"
	data[i].userPassword = "bagas9090"
	data[i].serialCode = "SN-017-BAGAS"
	data[i].usingLaptop = false
	data[i].batteryHealth = 0
	data[i].cpuManufacturer = "INTEL"
	data[i].cpuModel = "CORE"
	data[i].cpuSerial = "i5-9600K"
	data[i].gpuManufacturer = "NVIDIA"
	data[i].gpuModel = "GTX"
	data[i].gpuSerial = "GTX 1070"
	data[i].ramCapacity = 16
	data[i].ramUsed = 6.0
	data[i].diskCapacity = 500
	data[i].diskUsed = 465.0 // > 85% kapasitas → WARNING
	data[i].operatingSystem = "WINDOWS"
	data[i].dataLoad = false
	data[i].lastMaintenanceDate = "15-05-2025"
	cpuT = [10]float64{60, 61, 60, 62, 61, 60, 61, 60, 62, 61}
	gpuT = [10]float64{55, 56, 55, 57, 56, 55, 56, 55, 57, 56}
	ramT = [10]float64{45, 46, 45, 47, 46, 45, 46, 45, 47, 46}
	data[i].rataCpuTemp, data[i].medCpuTemp, data[i].modCpuTemp, data[i].minCpuTemp, data[i].maxCpuTemp = calcStats(cpuT)
	data[i].rataGpuTemp, data[i].medGpuTemp, data[i].modGpuTemp, data[i].minGpuTemp, data[i].maxGpuTemp = calcStats(gpuT)
	data[i].rataRamTemp, data[i].medRamTemp, data[i].modRamTemp, data[i].minRamTemp, data[i].maxRamTemp = calcStats(ramT)
	data[i].dataSudahDiisi = true
	setData(data, i)
	*userIndex++

	// ============================================================
	// User 18 — citra_dewanti | MacOS | Apple M | Apple GPU
	// Status  : GUD
	// ============================================================
	i = *userIndex
	data[i].user = "citra_dewanti"
	data[i].userPassword = "citra2222"
	data[i].serialCode = "SN-018-CITRA"
	data[i].usingLaptop = true
	data[i].batteryHealth = 97.0
	data[i].cpuManufacturer = "APPLE"
	data[i].cpuModel = "M"
	data[i].cpuSerial = "M3"
	data[i].gpuManufacturer = "APPLE"
	data[i].gpuModel = "M"
	data[i].gpuSerial = "M3 GPU"
	data[i].ramCapacity = 8
	data[i].ramUsed = 3.5
	data[i].diskCapacity = 256
	data[i].diskUsed = 90.0
	data[i].operatingSystem = "MACOS"
	data[i].dataLoad = false
	data[i].lastMaintenanceDate = "20-06-2025"
	cpuT = [10]float64{42, 43, 42, 44, 43, 42, 43, 42, 44, 43}
	gpuT = [10]float64{40, 41, 40, 42, 41, 40, 41, 40, 42, 41}
	ramT = [10]float64{34, 35, 34, 36, 35, 34, 35, 34, 36, 35}
	data[i].rataCpuTemp, data[i].medCpuTemp, data[i].modCpuTemp, data[i].minCpuTemp, data[i].maxCpuTemp = calcStats(cpuT)
	data[i].rataGpuTemp, data[i].medGpuTemp, data[i].modGpuTemp, data[i].minGpuTemp, data[i].maxGpuTemp = calcStats(gpuT)
	data[i].rataRamTemp, data[i].medRamTemp, data[i].modRamTemp, data[i].minRamTemp, data[i].maxRamTemp = calcStats(ramT)
	data[i].dataSudahDiisi = true
	setData(data, i)
	*userIndex++

	// ============================================================
	// User 19 — evan_suryana | Linux | AMD Ryzen | AMD Radeon
	// Status  : CRITICAL (GPU overheat saat heavy load)
	// ============================================================
	i = *userIndex
	data[i].user = "evan_suryana"
	data[i].userPassword = "evan1010"
	data[i].serialCode = "SN-019-EVAN"
	data[i].usingLaptop = false
	data[i].batteryHealth = 0
	data[i].cpuManufacturer = "AMD"
	data[i].cpuModel = "RYZEN"
	data[i].cpuSerial = "Ryzen 5 7600"
	data[i].gpuManufacturer = "AMD"
	data[i].gpuModel = "RADEON"
	data[i].gpuSerial = "Radeon RX 7800 XT"
	data[i].ramCapacity = 16
	data[i].ramUsed = 10.0
	data[i].diskCapacity = 1024
	data[i].diskUsed = 400.0
	data[i].operatingSystem = "LINUX"
	data[i].dataLoad = true
	data[i].lastMaintenanceDate = "28-07-2025"
	cpuT = [10]float64{80, 81, 80, 82, 81, 80, 81, 80, 82, 81}
	gpuT = [10]float64{91, 92, 91, 93, 92, 91, 92, 91, 93, 92} // rata >90 → overheat
	ramT = [10]float64{65, 66, 65, 67, 66, 65, 66, 65, 67, 66}
	data[i].rataCpuTemp, data[i].medCpuTemp, data[i].modCpuTemp, data[i].minCpuTemp, data[i].maxCpuTemp = calcStats(cpuT)
	data[i].rataGpuTemp, data[i].medGpuTemp, data[i].modGpuTemp, data[i].minGpuTemp, data[i].maxGpuTemp = calcStats(gpuT)
	data[i].rataRamTemp, data[i].medRamTemp, data[i].modRamTemp, data[i].minRamTemp, data[i].maxRamTemp = calcStats(ramT)
	data[i].dataSudahDiisi = true
	setData(data, i)
	*userIndex++

	// ============================================================
	// User 20 — fira_oktaviani | Windows | Intel Core | NVIDIA RTX
	// Status  : GUD
	// ============================================================
	i = *userIndex
	data[i].user = "fira_oktaviani"
	data[i].userPassword = "fira5050"
	data[i].serialCode = "SN-020-FIRA"
	data[i].usingLaptop = true
	data[i].batteryHealth = 91.0
	data[i].cpuManufacturer = "INTEL"
	data[i].cpuModel = "CORE"
	data[i].cpuSerial = "i7-1360P"
	data[i].gpuManufacturer = "NVIDIA"
	data[i].gpuModel = "RTX"
	data[i].gpuSerial = "RTX 4060"
	data[i].ramCapacity = 16
	data[i].ramUsed = 6.5
	data[i].diskCapacity = 512
	data[i].diskUsed = 180.0
	data[i].operatingSystem = "WINDOWS"
	data[i].dataLoad = false
	data[i].lastMaintenanceDate = "10-08-2025"
	cpuT = [10]float64{58, 59, 58, 60, 59, 58, 59, 58, 60, 59}
	gpuT = [10]float64{54, 55, 54, 56, 55, 54, 55, 54, 56, 55}
	ramT = [10]float64{44, 45, 44, 46, 45, 44, 45, 44, 46, 45}
	data[i].rataCpuTemp, data[i].medCpuTemp, data[i].modCpuTemp, data[i].minCpuTemp, data[i].maxCpuTemp = calcStats(cpuT)
	data[i].rataGpuTemp, data[i].medGpuTemp, data[i].modGpuTemp, data[i].minGpuTemp, data[i].maxGpuTemp = calcStats(gpuT)
	data[i].rataRamTemp, data[i].medRamTemp, data[i].modRamTemp, data[i].minRamTemp, data[i].maxRamTemp = calcStats(ramT)
	data[i].dataSudahDiisi = true
	setData(data, i)
	*userIndex++

	// ============================================================
	// User 21 — gilang_wibowo | Linux | Intel Xeon | NVIDIA RTX
	// Status  : GUD
	// ============================================================
	i = *userIndex
	data[i].user = "gilang_wibowo"
	data[i].userPassword = "gilang6969"
	data[i].serialCode = "SN-021-GILANG"
	data[i].usingLaptop = false
	data[i].batteryHealth = 0
	data[i].cpuManufacturer = "INTEL"
	data[i].cpuModel = "XEON"
	data[i].cpuSerial = "Xeon W-2295"
	data[i].gpuManufacturer = "NVIDIA"
	data[i].gpuModel = "RTX"
	data[i].gpuSerial = "RTX 4070"
	data[i].ramCapacity = 64
	data[i].ramUsed = 25.0
	data[i].diskCapacity = 2048
	data[i].diskUsed = 700.0
	data[i].operatingSystem = "LINUX"
	data[i].dataLoad = true
	data[i].lastMaintenanceDate = "05-09-2025"
	cpuT = [10]float64{78, 79, 78, 80, 79, 78, 79, 78, 80, 79}
	gpuT = [10]float64{72, 73, 72, 74, 73, 72, 73, 72, 74, 73}
	ramT = [10]float64{62, 63, 62, 64, 63, 62, 63, 62, 64, 63}
	data[i].rataCpuTemp, data[i].medCpuTemp, data[i].modCpuTemp, data[i].minCpuTemp, data[i].maxCpuTemp = calcStats(cpuT)
	data[i].rataGpuTemp, data[i].medGpuTemp, data[i].modGpuTemp, data[i].minGpuTemp, data[i].maxGpuTemp = calcStats(gpuT)
	data[i].rataRamTemp, data[i].medRamTemp, data[i].modRamTemp, data[i].minRamTemp, data[i].maxRamTemp = calcStats(ramT)
	data[i].dataSudahDiisi = true
	setData(data, i)
	*userIndex++

	// ============================================================
	// User 22 — hana_permata | Windows | AMD Ryzen | NVIDIA GTX
	// Status  : VERY_CRITICAL (CPU + GPU + RAM overheat, heavy load)
	// ============================================================
	i = *userIndex
	data[i].user = "hana_permata"
	data[i].userPassword = "hana2626"
	data[i].serialCode = "SN-022-HANA"
	data[i].usingLaptop = true
	data[i].batteryHealth = 45.0
	data[i].cpuManufacturer = "AMD"
	data[i].cpuModel = "RYZEN"
	data[i].cpuSerial = "Ryzen 5 4600H"
	data[i].gpuManufacturer = "NVIDIA"
	data[i].gpuModel = "GTX"
	data[i].gpuSerial = "GTX 1650"
	data[i].ramCapacity = 8
	data[i].ramUsed = 7.0
	data[i].diskCapacity = 256
	data[i].diskUsed = 120.0
	data[i].operatingSystem = "WINDOWS"
	data[i].dataLoad = true
	data[i].lastMaintenanceDate = "11-10-2024"
	cpuT = [10]float64{96, 97, 96, 98, 96, 97, 96, 97, 96, 98}  // rata >95 → overheat
	gpuT = [10]float64{86, 87, 86, 88, 86, 87, 86, 87, 86, 88}  // rata >85 → overheat
	ramT = [10]float64{86, 87, 86, 88, 86, 87, 86, 87, 86, 88}  // rata >85 → overheat
	data[i].rataCpuTemp, data[i].medCpuTemp, data[i].modCpuTemp, data[i].minCpuTemp, data[i].maxCpuTemp = calcStats(cpuT)
	data[i].rataGpuTemp, data[i].medGpuTemp, data[i].modGpuTemp, data[i].minGpuTemp, data[i].maxGpuTemp = calcStats(gpuT)
	data[i].rataRamTemp, data[i].medRamTemp, data[i].modRamTemp, data[i].minRamTemp, data[i].maxRamTemp = calcStats(ramT)
	data[i].dataSudahDiisi = true
	setData(data, i)
	*userIndex++
}
