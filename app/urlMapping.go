package app

func mapUrls() {
	v1 := router.Group("api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/")
		}
	}
}
