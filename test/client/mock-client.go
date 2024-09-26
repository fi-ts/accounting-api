package client

import (
	"testing"

	v1 "github.com/fi-ts/accounting-go/pkg/apis/v1"
	accountingclient "github.com/fi-ts/accounting-go/pkg/client"
	accmocks "github.com/fi-ts/accounting-go/test/mocks/v1"
	"github.com/stretchr/testify/mock"

	healthv1 "google.golang.org/grpc/health/grpc_health_v1"
)

type AccountingMockFns struct {
	Cluster            func(mock *mock.Mock)
	Pod                func(mock *mock.Mock)
	S3                 func(mock *mock.Mock)
	IP                 func(mock *mock.Mock)
	Info               func(mock *mock.Mock)
	NetworkTraffic     func(mock *mock.Mock)
	Volume             func(mock *mock.Mock)
	Postgres           func(mock *mock.Mock)
	ProductOption      func(mock *mock.Mock)
	Machine            func(mock *mock.Mock)
	MachineReservation func(mock *mock.Mock)
}

type AccountingMockClient struct {
	ClusterService            *accmocks.ClusterServiceClient
	PodService                *accmocks.PodServiceClient
	S3Service                 *accmocks.S3ServiceClient
	IPService                 *accmocks.IPServiceClient
	InfoService               *accmocks.InfoServiceClient
	NetworkTrafficService     *accmocks.NetworkTrafficServiceClient
	VolumeService             *accmocks.VolumeServiceClient
	PostgresService           *accmocks.PostgresServiceClient
	ProductOptionService      *accmocks.ProductOptionServiceClient
	MachineService            *accmocks.MachineServiceClient
	MachineReservationService *accmocks.MachineReservationServiceClient
}

func NewAccountingMockClient(mockFns *AccountingMockFns) (*AccountingMockClient, accountingclient.AccountingClient) {
	a := &AccountingMockClient{
		ClusterService:            &accmocks.ClusterServiceClient{},
		PodService:                &accmocks.PodServiceClient{},
		S3Service:                 &accmocks.S3ServiceClient{},
		IPService:                 &accmocks.IPServiceClient{},
		InfoService:               &accmocks.InfoServiceClient{},
		NetworkTrafficService:     &accmocks.NetworkTrafficServiceClient{},
		VolumeService:             &accmocks.VolumeServiceClient{},
		PostgresService:           &accmocks.PostgresServiceClient{},
		ProductOptionService:      &accmocks.ProductOptionServiceClient{},
		MachineService:            &accmocks.MachineServiceClient{},
		MachineReservationService: &accmocks.MachineReservationServiceClient{},
	}

	if mockFns != nil {
		if mockFns.Cluster != nil {
			mockFns.Cluster(&a.ClusterService.Mock)
		}
		if mockFns.Pod != nil {
			mockFns.Pod(&a.PodService.Mock)
		}
		if mockFns.S3 != nil {
			mockFns.S3(&a.S3Service.Mock)
		}
		if mockFns.IP != nil {
			mockFns.IP(&a.IPService.Mock)
		}
		if mockFns.Info != nil {
			mockFns.Info(&a.InfoService.Mock)
		}
		if mockFns.NetworkTraffic != nil {
			mockFns.NetworkTraffic(&a.NetworkTrafficService.Mock)
		}
		if mockFns.Volume != nil {
			mockFns.Volume(&a.VolumeService.Mock)
		}
		if mockFns.Postgres != nil {
			mockFns.Postgres(&a.PostgresService.Mock)
		}
		if mockFns.ProductOption != nil {
			mockFns.ProductOption(&a.ProductOptionService.Mock)
		}
		if mockFns.Machine != nil {
			mockFns.Machine(&a.MachineService.Mock)
		}
		if mockFns.MachineReservation != nil {
			mockFns.MachineReservation(&a.MachineReservationService.Mock)
		}
	}

	return a, a
}

func (c *AccountingMockClient) Close() error {
	return nil
}

func (c *AccountingMockClient) Cluster() v1.ClusterServiceClient {
	return c.ClusterService
}

func (c *AccountingMockClient) Pod() v1.PodServiceClient {
	return c.PodService
}

func (c *AccountingMockClient) S3() v1.S3ServiceClient {
	return c.S3Service
}

func (c *AccountingMockClient) IP() v1.IPServiceClient {
	return c.IPService
}

func (c *AccountingMockClient) Info() v1.InfoServiceClient {
	return c.InfoService
}

func (c *AccountingMockClient) NetworkTraffic() v1.NetworkTrafficServiceClient {
	return c.NetworkTrafficService
}

func (c *AccountingMockClient) Volume() v1.VolumeServiceClient {
	return c.VolumeService
}

func (c *AccountingMockClient) Health() healthv1.HealthClient {
	return nil
}

func (c *AccountingMockClient) Postgres() v1.PostgresServiceClient {
	return c.PostgresService
}

func (c *AccountingMockClient) ProductOption() v1.ProductOptionServiceClient {
	return c.ProductOptionService
}

func (c *AccountingMockClient) Machine() v1.MachineServiceClient {
	return c.MachineService
}

func (c *AccountingMockClient) MachineReservation() v1.MachineReservationServiceClient {
	return c.MachineReservationService
}

func (c *AccountingMockClient) AssertExpectations(t *testing.T) {
	_ = c.ClusterService.AssertExpectations(t)
	_ = c.PodService.AssertExpectations(t)
	_ = c.S3Service.AssertExpectations(t)
	_ = c.IPService.AssertExpectations(t)
	_ = c.InfoService.AssertExpectations(t)
	_ = c.NetworkTrafficService.AssertExpectations(t)
	_ = c.VolumeService.AssertExpectations(t)
	_ = c.PostgresService.AssertExpectations(t)
	_ = c.ProductOptionService.AssertExpectations(t)
	_ = c.MachineService.AssertExpectations(t)
	_ = c.MachineReservationService.AssertExpectations(t)
}
