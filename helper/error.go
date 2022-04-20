package helper

func ThrowError(err error) {
	if err != nil {
		panic(err)
	}
}
