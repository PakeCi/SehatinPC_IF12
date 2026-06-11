package main 

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
	jumlahKomponenRusak int
}

// Intel Data : https://www.intel.com/content/www/us/en/products/details/processors.html
// AMD Data : https://www.amd.com/en/products/specifications/processors.html
// Apple Data : https://everymac.com/systems/apple/index-apple-specs-applespec.html
//
// beberapa data yang bisa dipake : https://www.darkflash.com/article/safe-cpu-temperature-guide
// (no we i'm making datasheet for all of ts)

type dataBase [NMAX]dataComponent
