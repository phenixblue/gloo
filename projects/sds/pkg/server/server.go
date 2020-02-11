package server

import (
	"context"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"net"
	"os"

	"github.com/solo-io/go-utils/hashutils"

	"github.com/solo-io/go-utils/contextutils"

	"google.golang.org/grpc"

	auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	sds "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v2"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/envoyproxy/go-control-plane/pkg/server"
)

const (
	sdsClient = "sds_client"
)

var (
	grpcOptions = []grpc.ServerOption{grpc.MaxConcurrentStreams(10000)}
)

type EnvoyKey struct{}

func (h *EnvoyKey) ID(node *core.Node) string {
	return sdsClient
}

func SetupEnvoySDS() (*grpc.Server, cache.SnapshotCache) {
	grpcServer := grpc.NewServer(grpcOptions...)
	hasher := &EnvoyKey{}
	snapshotCache := cache.NewSnapshotCache(false, hasher, nil)
	svr := server.NewServer(context.Background(), snapshotCache, nil)

	// register services
	sds.RegisterSecretDiscoveryServiceServer(grpcServer, svr)
	return grpcServer, snapshotCache
}

func RunSDSServer(ctx context.Context, grpcServer *grpc.Server, serverAddress string) error {
	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		return err
	}
	contextutils.LoggerFrom(ctx).Info(fmt.Sprintf("sds server listening on %s\n", serverAddress))
	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			contextutils.LoggerFrom(ctx).Error(fmt.Sprintf("Stopping sds server listening on %s\n", serverAddress))
			os.Exit(1)
		}
	}()
	go func() {
		<-ctx.Done()
		contextutils.LoggerFrom(ctx).Info(fmt.Sprintf("stopping sds server on %s\n", serverAddress))
		grpcServer.GracefulStop()
	}()
	return nil
}

func GetSnapshotVersion(sslKeyFile, sslCertFile, sslCaFile string) (string, error) {
	var err error
	key, err := ioutil.ReadFile(sslKeyFile)
	if err != nil {
		return "", err
	}
	cert, err := ioutil.ReadFile(sslCertFile)
	if err != nil {
		return "", err
	}
	ca, err := ioutil.ReadFile(sslCaFile)
	if err != nil {
		return "", err
	}
	hash, err := hashutils.HashAllSafe(fnv.New64(), key, cert, ca)
	return fmt.Sprintf("%d", hash), err
}

func UpdateSDSConfig(ctx context.Context, sslKeyFile, sslCertFile, sslCaFile string, snapshotCache cache.SnapshotCache) error {
	snapshotVersion, err := GetSnapshotVersion(sslKeyFile, sslCertFile, sslCaFile)
	if err != nil {
		return err
	}
	contextutils.LoggerFrom(ctx).Info(fmt.Sprintf("Updating SDS config. Snapshot version is %s", snapshotVersion))

	items := []cache.Resource{
		serverCertSecret(sslCertFile, sslKeyFile),
		validationContextSecret(sslCaFile),
	}
	secretSnapshot := cache.Snapshot{}
	secretSnapshot.Resources[cache.Secret] = cache.NewResources(snapshotVersion, items)
	return snapshotCache.SetSnapshot(sdsClient, secretSnapshot)
}

func serverCertSecret(certFile, keyFile string) cache.Resource {
	return &auth.Secret{
		Name: "server_cert",
		Type: &auth.Secret_TlsCertificate{
			TlsCertificate: &auth.TlsCertificate{
				CertificateChain: &core.DataSource{
					Specifier: &core.DataSource_Filename{
						Filename: certFile,
					},
				},
				PrivateKey: &core.DataSource{
					Specifier: &core.DataSource_Filename{
						Filename: keyFile,
					},
				},
			},
		},
	}
}

func validationContextSecret(caFile string) cache.Resource {
	return &auth.Secret{
		Name: "validation_context",
		Type: &auth.Secret_ValidationContext{
			ValidationContext: &auth.CertificateValidationContext{
				TrustedCa: &core.DataSource{
					Specifier: &core.DataSource_Filename{
						Filename: caFile,
					},
				},
			},
		},
	}
}
