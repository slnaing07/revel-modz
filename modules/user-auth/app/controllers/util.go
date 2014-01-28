package controllers

func checkPass(err error) {
	if err != nil {
		println("Error: ", err)
	}
}

func checkFail(err error) {
	if err != nil {
		panic(err)
	}
}
