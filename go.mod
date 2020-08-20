module datacenter

go 1.14

replace (
	api => ./api
	config => ./config
	model => ./model
	router => ./router
	service => ./service
	serializer => ./serializer
)

require (
	config v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
	model v0.0.0-00010101000000-000000000000
	router v0.0.0-00010101000000-000000000000
	service v0.0.0-00010101000000-000000000000
)
