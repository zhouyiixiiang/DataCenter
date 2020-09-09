module datacenter

go 1.14

replace (
	api => ./api
	config => ./config
	middleware => ./middleware
	model => ./model
	router => ./router
	serializer => ./serializer
	service => ./service
)

require (
	config v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
	middleware v0.0.0-00010101000000-000000000000 // indirect
	model v0.0.0-00010101000000-000000000000
	router v0.0.0-00010101000000-000000000000
	service v0.0.0-00010101000000-000000000000
)
