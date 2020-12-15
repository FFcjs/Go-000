package main

import (
	"log"

	helloworld "Go-000/Week04/api/myservice"
	"Go-000/Week04/internal/myservice/service"
	"github.com/go-kratos/kratos/v2"
	httptransport "github.com/go-kratos/kratos/v2/transport/http"

	grpctransport "github.com/go-kratos/kratos/v2/transport/grpc"

	_ "github.com/go-kratos/kratos/v2/encoding/json"
	_ "github.com/go-kratos/kratos/v2/encoding/proto"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Version is the version of the compiled software.
	Version string
	// Branch is current branch name the code is built off.
	Branch string
	// Revision is the short commit hash of source tree.
	Revision string
	// BuildDate is the date when the binary was built.
	BuildDate string
)

func main() {
	log.Printf("service version: %s\n", Version)
	//init resource

	//uc := biz.NewUserUsecase(data.NewUserRepo())
	uc, cleanup, err := InitializeUserUsecase()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()
	// transport server
	httpSrv := httptransport.NewServer(httptransport.WithAddress(":8000"), httptransport.WithErrorHandler(httptransport.DefaultErrorHandler))
	grpcSrv := grpctransport.NewServer(grpctransport.WithAddress(":9000"))

	// register service
	gs := service.NewGreeterService(uc)
	helloworld.RegisterGreeterServer(grpcSrv, gs)
	helloworld.RegisterGreeterHTTPServer(httpSrv, gs)

	// application lifecycle
	app := kratos.New()
	app.Append(kratos.Hook{OnStart: httpSrv.Start, OnStop: httpSrv.Stop})
	app.Append(kratos.Hook{OnStart: grpcSrv.Start, OnStop: grpcSrv.Stop})

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		log.Printf("startup failed: %v\n", err)
	}
}
