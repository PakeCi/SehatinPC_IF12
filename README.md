# Tugas Besar Algoritma Pemrograman 2

**Anggota:**
- Yohanes Bimo Satrio Pinandito | IF-49-12 | 103012500049
- Maha Rayga Salim | IF-49-12 | 103012500183

## Program

Program **Sistem Monitoring Kesehatan Komponen PC**

**Fungsi**

Fungsi program yang diminta dalam spesifikasi adalah untuk mengawasi kondisi operasional perangkat keras secara real-time, dan data utama adalah data komponen, data suhu sensor, dan data penggunaan beban kerja. Pengguna aplikasi adalah teknisi komputer / pemilik sistem PC

**Spesifikasi Utama**

* Pengguna dapat menambahkan, mengubah, dan menghapus data komponen PC yang terpasang.
* Sistem dapat mencatat status kondisi perangkat terutama saat mengalami lag atau panas berlebih (overheat).
* Pengguna dapat mencari data komponen berdasarkan nama perangkat atau status kesehatan menggunakan Sequential dan Binary Search.
* Pengguna dapat mengurutkan data perangkat berdasarkan nomor seri komponen menggunakan Selection dan Insertion Sort.
* Sistem dapat menampilkan statistik jumlah komponen yang bermasalah dan rata-rata suhu kerja perangkat.

**Spesifikasi Sistem**

### 1. Klasifikasi Thermal Throttling 
| Manufacturer | Model | Limit Idle | Limit Heavy Load |
| :--- | :--- | :--- | :--- |
| **Intel** | Pentium | 70°C | 85°C |
| | Xeon    | 80°C | 95°C |
| | Atom    | 65°C | 80°C |
| | Core    | 90°C | 100°C |
| **AMD**   | Ryzen | 70°C | 95°C |
| | Epyc    | 85°C | 95°C |
| | Genereic| 65°C | 80°C |
| **Apple** | M-Series | 70°C | 95°C |

| Manufacturer | Model | Limit Idle | Limit Heavy Load |
| :--- | :--- | :--- | :--- |
| **NVIDIA** | RTX/GTX | 70°C | 85°C |
| **AMD**   | Radeon | 75°C | 90°C |
| **Apple** | M-Series | 75°C | 95°C |

### 2. Klasifikasi Storage / RAM Overload
* **Windows:** RAM < 15% || Disk < 15%  
* **Linux:** RAM < 5% || Disk < 15%
* **MacOS:** RAM < 15% || Disk < 10%

### 3. Next Maintenance Calculation
* **'GUD'** Next Maintenance di 6 bulan ke depan
* **'WARNING'** Next Maintenance di 3 bulan ke depan
* **'CRITICAL'** Next Maintenance di 7 hari ke depan
* **'VERY_CRITICAL'** Next Maintenance besok damn

**Module**
```
TB_SehatinPC
├─ README.md
├─ go.mod
├─ SehatinPC.go
├─ cli.go
├─ data_process.go
├─ searching.go
├─ sorting.go
├─ struct.go
├─ verification.go
├─ menu.go
└─ dummy.go
```

## How To Use
```
git clone https://github.com/PakeCi/SehatinPC_IF12
cd SehatinPC_IF12
go run . 
```