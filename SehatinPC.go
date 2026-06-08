package main


func main() {
	var data dataBase
	var exit bool = false
	var login bool = false
	var firstOpenPage bool = true
	var userIndex, loggedInUser int

	data[0].user = "admin"
	data[0].userPassword = "admin"
	userIndex = 1

	seedDummyData(&data, &userIndex)

	for !exit {
		if !login {
			loginMenu(&data, &exit, &login, &userIndex, &firstOpenPage, &loggedInUser)
			if login {
				firstOpenPage = true
			}
		} else {
			// if loggedInUser == 0 {
			// 	adminMenu(&data, &exit, &login, &userIndex, &firstOpenPage)
			// }else {
			// 	mainMenu(&data, &exit, &login, &userIndex, &firstOpenPage, loggedInUser)
			// }
			mainMenu(&data, &exit, &login, &userIndex, &firstOpenPage, loggedInUser)
			if !login {
				firstOpenPage = true
			}
		}
	}
}




// func adminMenu(data *dataBase, kill *bool, login *bool, totalUser *int, firstOpenPage *bool) {
// 	var exit bool = false
// 	var input int

// 	if *firstOpenPage {
// 		header()
// 		*firstOpenPage = false
// 	}
// 	for !exit {
// 		fmt.Printf("\n%-46s%s\n\n", " ", "ADMIN MENU")
// 		fmt.Printf("%-43s%s\n%-43s%s\n%-43s%s\n%-43s%s\n%-43s%s\n", " ", "1. Show All User", " ", "2. Change User Data", " ", "3. Delete User Data", " ", "4. Show All User Statistics", " ", "5. Logout", " ", "6. Kill Program")
// 		footer()
// 		fmt.Print("Input : ")
// 		fmt.Scan(&input)
// 	}
// }

//why the fuck did i make ts so fucking complicated AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA

// func copyData(data *dataBase, totalUser *int, arrayCoppied *dataBase) {
// 	for i := 1; i < *totalUser; i++ {
// 		arrayCoppied[i] = data[i]
// 	}
// }






