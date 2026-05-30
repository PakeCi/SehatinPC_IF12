package main 

func seedDummyData(data *dataBase, userIndex *int) {
	// --- Dummy User 1: Windows Gaming Laptop (Overheating CPU/GPU -> CRITICAL) ---
	data[*userIndex] = dataComponent{
		user: "gamer1", userPassword: "password123", dataSudahDiisi: true,
		serialCode: "SN-WIN-001", usingLaptop: true, batteryHealth: 88.0,
		cpuManufacturer: "AMD", cpuModel: "RYZEN", cpuSerial: "AMD-R7-5800H",
		gpuManufacturer: "NVIDIA", gpuModel: "RTX", gpuSerial: "NV-RTX3060",
		operatingSystem: "WINDOWS", dataLoad: true,
		ramCapacity: 16.0, ramUsed: 14.5, diskCapacity: 512.0, diskUsed: 400.0,
		rataCpuTemp: 98.0, medCpuTemp: 97.5, modCpuTemp: 98.0, minCpuTemp: 90.0, maxCpuTemp: 100.0,
		rataGpuTemp: 87.0, medGpuTemp: 86.5, modGpuTemp: 87.0, minGpuTemp: 80.0, maxGpuTemp: 89.0,
		rataRamTemp: 65.0, medRamTemp: 64.0, modRamTemp: 65.0, minRamTemp: 60.0, maxRamTemp: 68.0,
		lastMaintenanceDate: "15-10-2023",
	}
	setData(data, *userIndex) // Computes 'status' and 'nextMaintenanceDate'
	*userIndex++

	// --- Dummy User 2: Apple MacBook (Healthy / Idle -> GUD) ---
	data[*userIndex] = dataComponent{
		user: "macuser", userPassword: "password123", dataSudahDiisi: true,
		serialCode: "SN-MAC-002", usingLaptop: true, batteryHealth: 100.0,
		cpuManufacturer: "APPLE", cpuModel: "M", cpuSerial: "APP-M2-001",
		gpuManufacturer: "APPLE", gpuModel: "M", gpuSerial: "APP-M2-001",
		operatingSystem: "MACOS", dataLoad: false,
		ramCapacity: 8.0, ramUsed: 4.2, diskCapacity: 256.0, diskUsed: 100.0,
		rataCpuTemp: 45.0, medCpuTemp: 44.5, modCpuTemp: 45.0, minCpuTemp: 40.0, maxCpuTemp: 50.0,
		rataGpuTemp: 46.0, medGpuTemp: 45.5, modGpuTemp: 46.0, minGpuTemp: 41.0, maxGpuTemp: 51.0,
		rataRamTemp: 40.0, medRamTemp: 39.0, modRamTemp: 40.0, minRamTemp: 35.0, maxRamTemp: 42.0,
		lastMaintenanceDate: "01-01-2024",
	}
	setData(data, *userIndex)
	*userIndex++

	// --- Dummy User 3: Linux Server (No GPU, Low Disk Space -> WARNING) ---
	data[*userIndex] = dataComponent{
		user: "server_admin", userPassword: "password123", dataSudahDiisi: true,
		serialCode: "SN-LIN-003", usingLaptop: false, batteryHealth: -1,
		cpuManufacturer: "INTEL", cpuModel: "XEON", cpuSerial: "INT-XN-990",
		gpuManufacturer: "NONE", gpuModel: "NONE", gpuSerial: "NONE",
		operatingSystem: "LINUX", dataLoad: true,
		ramCapacity: 64.0, ramUsed: 32.0, diskCapacity: 2000.0, diskUsed: 1950.0, // Only 50GB left, triggers Warning
		rataCpuTemp: 75.0, medCpuTemp: 74.0, modCpuTemp: 75.0, minCpuTemp: 70.0, maxCpuTemp: 80.0,
		rataGpuTemp: 0.0, medGpuTemp: 0.0, modGpuTemp: 0.0, minGpuTemp: 0.0, maxGpuTemp: 0.0,
		rataRamTemp: 55.0, medRamTemp: 54.0, modRamTemp: 55.0, minRamTemp: 50.0, maxRamTemp: 60.0,
		lastMaintenanceDate: "20-02-2024",
	}
	setData(data, *userIndex)
	*userIndex++
}