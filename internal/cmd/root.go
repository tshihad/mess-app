package cmd

// Execute is the starting point of app
func Execute() {
	c, err := mustReadConfig()
	if err != nil {
		panic(err)
	}
	if err := mustValidateConfig(c); err != nil {
		panic(err)
	}
	db, err := mustPrepareDB(c)
	if err != nil {
		panic(err)
	}
	serveApp(db, c)
}
