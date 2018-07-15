package common

//StartUp operations
func StartUp() {

	initConfig()

	initKeys()

	createDbSession()

	addIndexes()

}
