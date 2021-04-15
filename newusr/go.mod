module newusr

go 1.16

require (
	github.com/hashicorp/consul/api v1.8.1
	google.golang.org/grpc v1.37.0
	keeyu/message v0.0.0-00010101000000-000000000000
)

replace (
	keeyu/message => ../proto
	keeyu/tool => ../tool
)
