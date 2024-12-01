package users

func CheckAge(age int) bool {
	if age >= 18 {
		return true
	}

	return false
}

func CheckAge2(age int) bool {
	return age >= 18
}
