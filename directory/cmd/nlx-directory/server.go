package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"

	"github.com/VNG-Realisatie/nlx/directory/directoryapi"
)

// newGRPCSplitterHandlerFunc returns an http.Handler that delegates gRPC connections to grpcServer
// and all other connections to otherHandler.
func newGRPCSplitterHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// gRPC connection detected when HTTP protocol is version 2 and content-type is application/grpc
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

// runServer is a blocking function which sets up the grpc and http/json server and runs them on a single address/port.
func runServer(log *zap.Logger, address string, caCertPool *x509.CertPool, certKeyPair tls.Certificate, directoryService directoryapi.DirectoryServer) {

	// setup logrus connection for global grpc logging
	grpc_zap.ReplaceGrpcLogger(log)

	recoveryOptions := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(func(p interface{}) error {
			log.Warn("recovered from a panic in a grpc request handler", zap.ByteString("stack", debug.Stack()))
			return grpc.Errorf(codes.Internal, "%s", p)
		}),
	}

	// prepare grpc server options
	opts := []grpc.ServerOption{
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(log),
			grpc_recovery.StreamServerInterceptor(recoveryOptions...),
		),
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(log),
			grpc_recovery.UnaryServerInterceptor(recoveryOptions...),
		),
	}

	// start grpc server and attach director service
	grpcServer := grpc.NewServer(opts...)
	directoryapi.RegisterDirectoryServer(grpcServer, directoryService)

	// setup client credentials for grpc gateway
	gatewayDialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(
			credentials.NewTLS(&tls.Config{
				Certificates:       []tls.Certificate{certKeyPair}, // using the grpc server's own cert to connect to it, perhaps find a way for the http/json gatewy to bypass tls locally
				RootCAs:            caCertPool,
				InsecureSkipVerify: true, // This is a local connection; hostname won't match
			}),
		),
	}

	// root http serve mux
	httpRouter := http.NewServeMux()
	httpRouter.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
		io.Copy(w, strings.NewReader(directoryapi.SwaggerJSONDirectory))
	})

	// setup grpc gateway and attach to main mux
	gatewayMux := runtime.NewServeMux()
	err := directoryapi.RegisterDirectoryHandlerFromEndpoint(context.Background(), gatewayMux, address, gatewayDialOptions)
	if err != nil {
		fmt.Printf("serve: %v\n", err)
		return
	}
	httpRouter.Handle("/", gatewayMux)

	// start a simple tcp listener
	tcpListener, err := net.Listen("tcp", address)
	if err != nil {
		log.With(zap.Error(err)).Fatal(fmt.Sprintf("could not listen on %s", address))
	}

	// wrap the tcp listener with tls
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{certKeyPair},
		NextProtos:   []string{"h2"},
		ClientCAs:    caCertPool,
		ClientAuth:   tls.VerifyClientCertIfGiven,
		// TODO: ClientAUth must actually be tls.RequireAndVerifyClientCert, but we can only do that when the registration endpoint (client-tls mandatory)
		// has been seperated from the inspection endpoint (client-tls optional). And perhaps the gateway should run seperately all together.
	}
	tlsListener := tls.NewListener(tcpListener, tlsConfig)

	// let server handle connections on the tls Listener
	srv := &http.Server{
		Addr:    address,
		Handler: newGRPCSplitterHandlerFunc(grpcServer, httpRouter),
	}
	err = srv.Serve(tlsListener)
	if err != nil {
		log.With(zap.Error(err)).Fatal("ListenAndServe failed")
	}

	return
}