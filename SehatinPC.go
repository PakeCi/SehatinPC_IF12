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
			mainMenu(&data, &exit, &login, &userIndex, &firstOpenPage, loggedInUser)
			if !login {
				firstOpenPage = true
			}
		}
	}
}




