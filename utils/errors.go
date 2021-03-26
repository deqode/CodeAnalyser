package utils

func HandleErr(err error) {
	if err != nil {
		Logger(err)
	}
}
