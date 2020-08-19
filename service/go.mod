module service

go 1.14

replace (
	model => ../model
	serializer => ../serializer
)

require (
	github.com/gin-gonic/gin v1.6.3
	model v0.0.0-00010101000000-000000000000
	serializer v0.0.0-00010101000000-000000000000
)
