package utils

func PanicIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}
