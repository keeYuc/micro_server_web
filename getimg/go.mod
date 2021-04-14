module getimg

go 1.16

require (
	github.com/afocus/captcha v0.0.0-20191010092841-4bd1f21c8868
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	golang.org/x/image v0.0.0-20210220032944-ac19c3e999fb // indirect
	google.golang.org/grpc v1.37.0
	keeyu/message v0.0.1
	keeyu/tool v0.0.0-00010101000000-000000000000
)

replace (
	keeyu/message => ../proto
	keeyu/tool => ../tool
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
