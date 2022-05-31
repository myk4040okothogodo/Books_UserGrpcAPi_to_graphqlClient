module github.com/myk4040okothogodo/tutorial/holders/server

go 1.18

require (
	github.com/arangodb/go-driver v1.3.2
	github.com/myk4040okothogodo/tutorial/db v0.0.0-00010101000000-000000000000
	github.com/myk4040okothogodo/tutorial/gen/go/proto/holders latest
	google.golang.org/grpc v1.46.2
)

require (
	github.com/arangodb/go-velocypack v0.0.0-20200318135517-5af53c29c67e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.0.0-20220526153639-5463443f8c37 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/myk4040okothogodo/tutorial/gen/go/proto/holders => ../../gen/go/proto/holders

replace github.com/myk4040okothogodo/tutorial/db => ../../db
