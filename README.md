Tugas Besar Algoritma Pemrograman 2

Yohanes Bimo Satrio Pinandito | IF-49-12 | 103012500049
Maha Rayga Salim | IF-49-12 | 103012500183

Program Sistem Monitoring Kesehatan Komponen PC
/_ Fungsi program yang diminta dalam spesifikasi adalah untuk mengawasi kondisi operasional perangkat keras secara real-time, dan data utama adalah data komponen, data suhu sensor, dan data penggunaan beban kerja. Pengguna aplikasi adalah teknisi komputer / pemilik sistem PC _/

Spesifikasi dalam PDF :
a. Pengguna dapat menambahkan, mengubah, dan menghapus data komponen PC yang terpasang.
b. Sistem dapat mencatat status kondisi perangkat terutama saat mengalami lag atau panas berlebih (overheat).
c. Pengguna dapat mencari data komponen berdasarkan nama perangkat atau status kesehatan menggunakan Sequential dan Binary Search.
d. Pengguna dapat mengurutkan data perangkat berdasarkan nomor seri komponen menggunakan Selection dan Insertion Sort.
e. Sistem dapat menampilkan statistik jumlah komponen yang bermasalah dan rata-rata suhu kerja perangkat.

Disini kami membuat sistem seperti berikut : 

> awal-awal akan masuk ke dalam login menu, yang di mana login menu ini akan menanyakan user apakah ingin login atau register, ketika user memilih register, maka akan pergi ke register page yang di register page, user akan membuat array baru yakni array nama user dan  login, dalam register page juga akan melakukan checking apakah di array sebelumnya sudah ada username yang sama, jika sudah ada akan diminta untuk input ulang nah kemudian akan terdapat opsi lanjutkan atau kembali, jika kembali akan kembali ke loginmenu, jika continue akan pergi ke main menu. Jika user memilih ke login page, di dalam login page user akan menginput username dan password yang telah dibuat di register page, dan akan melakukan checking apakah ada username dan password yang sama di dalam array. Jika ada berarti dapat masuk, jika tidak user akan diminta untuk login kembali. Dan terdapat base user yang berupa admin, di mana username dan passwordnya admin untuk masuk ke dalam. program. 

> di dalam program, akan diberikan beberapa menu, yakni untuk menginput data, dan jika user baru pertama kali masuk, yang hanya bisa diakses adalah page input data, namun jika sudah menginput data, dan datanya sudah ada di database user bisa melanjutkan ke menu yang lain. Dalam input data, pertama user akan menginput beberapa data  yakni serialCode laptop, cpuManufacturer, gpuManufacturer, cpuModel, dan gpuModel, ramCapacity, diskCapacity serta lastMaintenanceDate. yang kemudian akan ada continue dan back, jika back dia masih tidak bisa mengakses menu lain sebelum data diisi, namun ketika continue user akan di bawa ke page untuk menginput data lanjutan yang dimana user akan diminta menginput cpuTemperatur, gpuTemperatur, ramTemperatur yang masing-masing menginput 10 data dalam 20 detik terakhir (soalnya di btop pergantian data setiap 2 detik) untuk mencari rata-rata, median, min, max dan modus dari temperature. Ketika sudah mendapatkan rata-rata, median, dan modus program akan memasukkan datanya ke array dataComponent sebagai rataCpuTemp, medCpuTemp, modCpuTemp. yang user juga akan diminta mengisi data ramUsed, diskUsed, dataLoad, dan operatingSystem sistem akan menanyakan apakah data tersebut dalam heavy load atau normal load. dan beberapa variabel seperti cpuManufacturer, gpuManufacturer, cpuModel, gpuModel, operatingSystem, dan dataLoad akan mempengaruhi bagaimana pemrosesan data supaya dianggap lag atau overheat.

> Spesifikasi : 
    Untuk dinyatakan overheat : 
    1. Intel : 
        Pentium : HeavyLoad >= 85, Idle >= 70
        Xeon : HeavyLoad >= 95, Idle >= 80
        Atom : HeavyLoad >= 80, Idle >= 65
    2. Amd : 
