module github.com/myk4040okothogodo/tutorial/holders

go 1.18

require (
	github.com/myk4040okothogodo/tutorial/db v0.0.0-00010101000000-000000000000
	github.com/myk4040okothogodo/tutorial/holders/server v0.0.0-00010101000000-000000000000
)

require (
	github.com/arangodb/go-driver v1.3.2 // indirect
	github.com/arangodb/go-velocypack v0.0.0-20200318135517-5af53c29c67e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.46.2 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/myk4040okothogodo/tutorial/db => ../db

replace github.com/myk4040okothogodo/tutorial/holders/server => ./server
