package main

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"
	"time"

	sdk "github.com/censys/censys-sdk-go"
	"github.com/censys/censys-sdk-go/models/components"
	"github.com/censys/censys-sdk-go/models/operations"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type SDKTestSuite struct {
	suite.Suite
	client *sdk.SDK
	ctx    context.Context
	orgID  string
}

func (s *SDKTestSuite) SetupSuite() {
	apiKey := os.Getenv("CENSYS_PAT")
	require.NotEmpty(s.T(), apiKey, "CENSYS_PAT must be set")
	orgID := os.Getenv("CENSYS_ORG_ID")
	require.NotEmpty(s.T(), orgID, "CENSYS_ORG_ID must be set")
	s.orgID = orgID

	s.client = sdk.New(
		sdk.WithSecurity(apiKey),
		sdk.WithOrganizationID(orgID),
	)
	s.ctx = context.Background()
}

func TestSDKSuite(t *testing.T) {
	suite.Run(t, new(SDKTestSuite))
}

func (s *SDKTestSuite) TestGlobalData_GetCertificates() {
	certIDs := []string{
		"00000002741c89f06524afbbb4720876bc173aca3a6ce344e08584859b9ac34e",
		"000000033b547e13ee216c65b0ff50237f0decef12acb76fce0a96afa9c70d50",
	}
	res, err := s.client.GlobalData.GetCertificates(s.ctx, operations.V3GlobaldataAssetCertificateListPostRequest{
		AssetCertificateListInputBody: components.AssetCertificateListInputBody{
			CertificateIds: certIDs,
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestGlobalData_GetCertificatesRaw() {
	certIDs := []string{
		"00000002741c89f06524afbbb4720876bc173aca3a6ce344e08584859b9ac34e",
		"000000033b547e13ee216c65b0ff50237f0decef12acb76fce0a96afa9c70d50",
	}
	res, err := s.client.GlobalData.GetCertificatesRaw(s.ctx, operations.V3GlobaldataAssetCertificateListRawPostRequest{
		AssetCertificateListInputBody: components.AssetCertificateListInputBody{
			CertificateIds: certIDs,
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestGlobalData_GetCertificate() {
	certID := "00000002741c89f06524afbbb4720876bc173aca3a6ce344e08584859b9ac34e"
	res, err := s.client.GlobalData.GetCertificate(s.ctx, operations.V3GlobaldataAssetCertificateRequest{
		CertificateID: certID,
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestGlobalData_GetCertificateRaw() {
	certID := "00000002741c89f06524afbbb4720876bc173aca3a6ce344e08584859b9ac34e"
	res, err := s.client.GlobalData.GetCertificateRaw(s.ctx, operations.V3GlobaldataAssetCertificateRawRequest{
		CertificateID: certID,
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
	defer res.ResponseStream.Close()
	buf := new(bytes.Buffer)
	n, err := io.Copy(buf, res.ResponseStream)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), n)
}

func (s *SDKTestSuite) TestGlobalData_GetHosts() {
	hostIDs := []string{"1.1.1.1", "8.8.8.8"}
	res, err := s.client.GlobalData.GetHosts(s.ctx, operations.V3GlobaldataAssetHostListPostRequest{
		AssetHostListInputBody: components.AssetHostListInputBody{
			HostIds: hostIDs,
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestGlobalData_GetHost() {
	hostID := "125.13.31.107"
	atTime, err := time.Parse(time.RFC3339, "2025-03-20T00:00:00Z")
	require.NoError(s.T(), err)
	res, err := s.client.GlobalData.GetHost(s.ctx, operations.V3GlobaldataAssetHostRequest{
		HostID: hostID,
		AtTime: &atTime,
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestGlobalData_GetHostTimeline() {
	hostID := "125.13.31.107"
	startTime, err := time.Parse(time.RFC3339, "2025-03-20T00:00:00Z")
	require.NoError(s.T(), err)
	endTime, err := time.Parse(time.RFC3339, "2025-03-22T00:00:00Z")
	require.NoError(s.T(), err)
	res, err := s.client.GlobalData.GetHostTimeline(s.ctx, operations.V3GlobaldataAssetHostTimelineRequest{
		HostID:    hostID,
		StartTime: startTime,
		EndTime:   endTime,
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestGlobalData_GetWebProperties() {
	webPropertyIDs := []string{
		"104.236.29.250:443",
		"78.133.74.135:49152",
	}
	res, err := s.client.GlobalData.GetWebProperties(s.ctx, operations.V3GlobaldataAssetWebpropertyListPostRequest{
		AssetWebpropertyListInputBody: components.AssetWebpropertyListInputBody{
			WebpropertyIds: webPropertyIDs,
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestGlobalData_GetWebProperty() {
	webPropertyID := "104.236.29.250:443"
	res, err := s.client.GlobalData.GetWebProperty(s.ctx, operations.V3GlobaldataAssetWebpropertyRequest{
		WebpropertyID: webPropertyID,
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestGlobalData_Aggregate() {
	res, err := s.client.GlobalData.Aggregate(s.ctx, operations.V3GlobaldataSearchAggregateRequest{
		SearchAggregateInputBody: components.SearchAggregateInputBody{
			Field:           "web.endpoints.http.status_reason",
			NumberOfBuckets: 2,
			Query:           "web.port: *",
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestGlobalData_Search() {
	pageSize := int64(3)
	res, err := s.client.GlobalData.Search(s.ctx, operations.V3GlobaldataSearchQueryRequest{
		SearchQueryInputBody: components.SearchQueryInputBody{
			Query:    "web.port: *",
			PageSize: &pageSize,
			Fields:   []string{"web.port"},
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestCollections() {
	res, err := s.client.Collections.Create(s.ctx, operations.V3CollectionsCrudCreateRequest{
		CrudCreateInputBody: &components.CrudCreateInputBody{
			Name:        "Test Collection NL",
			Description: strPtr("Test Collection NL"),
			Query:       "host.services.protocol='SSH' and host.location.country = 'Netherlands' and host.services.port = 9100 and host.autonomous_system.name = 'WORLDSTREAM'",
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)

	collectionUID := res.GetResponseEnvelopeCollection().GetResult().GetID()
	getRes, err := s.client.Collections.Get(s.ctx, operations.V3CollectionsCrudGetRequest{
		CollectionUID: collectionUID,
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), getRes)

	// list events
	listEventsRes, err := s.client.Collections.ListEvents(s.ctx, operations.V3CollectionsListEventsRequest{
		CollectionUID: collectionUID,
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), listEventsRes)

	// search aggregate
	searchAggregateRes, err := s.client.Collections.Aggregate(s.ctx, operations.V3CollectionsSearchAggregateRequest{
		CollectionUID: collectionUID,
		SearchAggregateInputBody: components.SearchAggregateInputBody{
			Field:           "host.autonomous_system.name",
			NumberOfBuckets: 10,
			Query:           "host.services.labels.value = 'REMOTE_ACCESS'",
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), searchAggregateRes)

	// search query
	searchQueryRes, err := s.client.Collections.Search(s.ctx, operations.V3CollectionsSearchQueryRequest{
		CollectionUID: collectionUID,
		SearchQueryInputBody: components.SearchQueryInputBody{
			Query: "host.services.labels.value = 'REMOTE_ACCESS'",
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), searchQueryRes)

	// update it
	updateRes, err := s.client.Collections.Update(s.ctx, operations.V3CollectionsCrudUpdateRequest{
		CollectionUID: collectionUID,
		CrudUpdateInputBody: &components.CrudUpdateInputBody{
			Description: strPtr("New desc"),
			Name:        "New name",
			Query:       "host.services.protocol='SSH' and host.location.country = 'Netherlands' and host.services.port = 9100 and host.autonomous_system.name = 'WORLDSTREAM'",
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), updateRes)

	// make sure it's updated
	getRes, err = s.client.Collections.Get(s.ctx, operations.V3CollectionsCrudGetRequest{
		CollectionUID: collectionUID,
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), getRes)
	require.Equal(s.T(), "New desc", getRes.GetResponseEnvelopeCollection().GetResult().GetDescription())

	// delete it
	deleteRes, err := s.client.Collections.Delete(s.ctx, operations.V3CollectionsCrudDeleteRequest{
		CollectionUID: collectionUID,
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), deleteRes)

	// make sure it's deleted
	getRes, err = s.client.Collections.Get(s.ctx, operations.V3CollectionsCrudGetRequest{
		CollectionUID: collectionUID,
	})
	require.Error(s.T(), err)
	require.Nil(s.T(), getRes)
}

func (s *SDKTestSuite) TestThreatHunting() {
	res, err := s.client.ThreatHunting.ValueCounts(s.ctx, operations.V3ThreathuntingValueCountsRequest{
		SearchValueCountsInputBody: components.SearchValueCountsInputBody{
			AndCountConditions: []components.CountCondition{
				{
					FieldValuePairs: []components.FieldValuePair{
						{
							Field: "host.services.port",
							Value: "80",
						},
					},
				},
			},
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func strPtr(s string) *string {
	return &s
}
