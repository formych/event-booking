https://github.com/formych/event-booking


go get github.com/formych/event-booking
go run main.go


静态文件放在public下面，html,css,js,images

通过localhost:8080加上参数访问接口和html文件

r.Static("/", "public")

u := r.Group("/api/user")
{
    u.POST("/signin", model.SignIn)
    u.POST("/signup", model.SignUp)
    u.POST("/signout", model.SignOut)
}

e := r.Group("/api/event")
e.Use(model.Authentication())
{
    e.POST("/add", model.EventAdd)
    e.DELETE("/delete", model.EventDelete)
}