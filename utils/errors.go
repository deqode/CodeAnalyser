package utils

//HandleErr todo: need to create/implement error handler
func HandleErr(err error) {
	if err != nil {
		Logger(err)
	}
}
