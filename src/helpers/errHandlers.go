package helpers

//CheckErr checks if err is not nil and panics
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
