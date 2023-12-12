package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"time"

	v1 "github.com/fi-ts/accounting-go/pkg/apis/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	healthv1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
)

type AccountingClient interface {
	Cluster() v1.ClusterServiceClient
	Pod() v1.PodServiceClient
	S3() v1.S3ServiceClient
	Volume() v1.VolumeServiceClient
	Info() v1.InfoServiceClient
	IP() v1.IPServiceClient
	NetworkTraffic() v1.NetworkTrafficServiceClient
	Health() healthv1.HealthClient
	Postgres() v1.PostgresServiceClient
	Machine() v1.MachineServiceClient
	ProductOption() v1.ProductOptionServiceClient
	Close() error
}

type client struct {
	conn *grpc.ClientConn
}

func NewClient(ctx context.Context, hostname string, port int, certFile string, keyFile string, caFile string, additionalOpts ...grpc.DialOption) (AccountingClient, error) {
	address := fmt.Sprintf("%s:%d", hostname, port)

	certPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, fmt.Errorf("failed to load system credentials: %w", err)
	}

	if caFile != "" {
		ca, err := os.ReadFile(caFile)
		if err != nil {
			return nil, fmt.Errorf("could not read ca certificate: %w", err)
		}

		ok := certPool.AppendCertsFromPEM(ca)
		if !ok {
			return nil, fmt.Errorf("failed to append ca certs: %s", caFile)
		}
	}

	clientCertificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("could not load client key pair: %w", err)
	}

	creds := credentials.NewTLS(&tls.Config{
		ServerName:   hostname,
		Certificates: []tls.Certificate{clientCertificate},
		RootCAs:      certPool,
		MinVersion:   tls.VersionTLS12,
	})

	kacp := keepalive.ClientParameters{
		Time:                20 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             3 * time.Second,  // wait 1 second for ping ack before considering the connection dead
		PermitWithoutStream: true,             // send pings even without active streams
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(kacp),
	}

	additionalOpts = append(opts, additionalOpts...)

	conn, err := grpc.DialContext(ctx, address, additionalOpts...)
	if err != nil {
		return nil, err
	}

	return client{
		conn: conn,
	}, nil
}

// Close the underlying connection
func (c client) Close() error {
	return c.conn.Close()
}

func (c client) Cluster() v1.ClusterServiceClient {
	return v1.NewClusterServiceClient(c.conn)
}

func (c client) Pod() v1.PodServiceClient {
	return v1.NewPodServiceClient(c.conn)
}

func (c client) S3() v1.S3ServiceClient {
	return v1.NewS3ServiceClient(c.conn)
}

func (c client) IP() v1.IPServiceClient {
	return v1.NewIPServiceClient(c.conn)
}

func (c client) Info() v1.InfoServiceClient {
	return v1.NewInfoServiceClient(c.conn)
}

func (c client) NetworkTraffic() v1.NetworkTrafficServiceClient {
	return v1.NewNetworkTrafficServiceClient(c.conn)
}

func (c client) Volume() v1.VolumeServiceClient {
	return v1.NewVolumeServiceClient(c.conn)
}

// Health is the root accessor for health related functions
func (c client) Health() healthv1.HealthClient {
	return healthv1.NewHealthClient(c.conn)
}

func (c client) Postgres() v1.PostgresServiceClient {
	return v1.NewPostgresServiceClient(c.conn)
}

func (c client) Machine() v1.MachineServiceClient {
	return v1.NewMachineServiceClient(c.conn)
}

func (c client) ProductOption() v1.ProductOptionServiceClient {
	return v1.NewProductOptionServiceClient(c.conn)
}
