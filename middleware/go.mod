module middleware

go 1.15

replace (
	model => ../model
	serializer => ../serializer
)

require (
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.6.3
	model v0.0.0-00010101000000-000000000000
	serializer v0.0.0-00010101000000-000000000000
)
