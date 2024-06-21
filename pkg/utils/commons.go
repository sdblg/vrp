package utils

// panic if any error encounterred
func PanicIfErr(err error, ignoreErrors ...error) {
	if err == nil {
		return
	}

	for _, e := range ignoreErrors {
		if err == e {
			return
		}
	}

	panic(err)
}
