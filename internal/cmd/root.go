package cmd

// Execute is the starting point of app
func Execute() {
	db, err := mustPrepareDB()
	if err != nil {
		panic(err)
	}
	serveApp(db)
}
