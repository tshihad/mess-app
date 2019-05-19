package cmd

// Execute is the starting point of app
func Execute() {
	logger, err := mustPrepareLogger()
	if err != nil {
		panic(err)
	}
	db, err := mustPrepareDB()
	if err != nil {
		logger.Panic(err)
	}
	redis, err := mustPrepareRedis()
	if err != nil {
		logger.Panic(err)
	}
	serveApp(db, redis, logger)
}
