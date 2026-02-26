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
	"github.com/censys/censys-sdk-go/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var (
	testCertIDs = []string{
		"00000002741c89f06524afbbb4720876bc173aca3a6ce344e08584859b9ac34e",
		"000000033b547e13ee216c65b0ff50237f0decef12acb76fce0a96afa9c70d50",
	}
	testHostIDs        = []string{"1.1.1.1", "8.8.8.8"}
	testWebPropertyIDs = []string{"104.236.29.250:443", "78.133.74.135:49152"}
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

// ---------------------------------------------------------------------------
// GlobalData — Certificates
// ---------------------------------------------------------------------------

func (s *SDKTestSuite) TestGlobalData_GetCertificates() {
	res, err := s.client.GlobalData.GetCertificates(s.ctx, operations.V3GlobaldataAssetCertificateListPostRequest{
		AssetCertificateListInputBody: components.AssetCertificateListInputBody{
			CertificateIds: testCertIDs,
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestGlobalData_GetCertificatesRaw() {
	res, err := s.client.GlobalData.GetCertificatesRaw(s.ctx, operations.V3GlobaldataAssetCertificateListRawPostRequest{
		AssetCertificateListInputBody: components.AssetCertificateListInputBody{
			CertificateIds: testCertIDs,
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestGlobalData_GetCertificate() {
	res, err := s.client.GlobalData.GetCertificate(s.ctx, operations.V3GlobaldataAssetCertificateRequest{
		CertificateID: testCertIDs[0],
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestGlobalData_GetCertificateRaw() {
	res, err := s.client.GlobalData.GetCertificateRaw(s.ctx, operations.V3GlobaldataAssetCertificateRawRequest{
		CertificateID: testCertIDs[0],
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
	defer res.ResponseStream.Close()

	buf := new(bytes.Buffer)
	n, err := io.Copy(buf, res.ResponseStream)
	require.NoError(s.T(), err)
	require.Positive(s.T(), n)
}

// ---------------------------------------------------------------------------
// GlobalData — Hosts
// ---------------------------------------------------------------------------

func (s *SDKTestSuite) TestGlobalData_GetHosts() {
	res, err := s.client.GlobalData.GetHosts(s.ctx, operations.V3GlobaldataAssetHostListPostRequest{
		AssetHostListInputBody: components.AssetHostListInputBody{
			HostIds: testHostIDs,
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestGlobalData_GetHost() {
	atTime := mustParseTime("2025-03-20T00:00:00Z")
	res, err := s.client.GlobalData.GetHost(s.ctx, operations.V3GlobaldataAssetHostRequest{
		HostID: "125.13.31.107",
		AtTime: &atTime,
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestGlobalData_GetHostTimeline() {
	startTime := mustParseTime("2025-03-20T00:00:00Z")
	endTime := mustParseTime("2025-03-22T00:00:00Z")
	res, err := s.client.GlobalData.GetHostTimeline(s.ctx, operations.V3GlobaldataAssetHostTimelineRequest{
		HostID:    "125.13.31.107",
		StartTime: startTime,
		EndTime:   endTime,
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

// ---------------------------------------------------------------------------
// GlobalData — Web Properties
// ---------------------------------------------------------------------------

func (s *SDKTestSuite) TestGlobalData_GetWebProperties() {
	res, err := s.client.GlobalData.GetWebProperties(s.ctx, operations.V3GlobaldataAssetWebpropertyListPostRequest{
		AssetWebpropertyListInputBody: components.AssetWebpropertyListInputBody{
			WebpropertyIds: testWebPropertyIDs,
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestGlobalData_GetWebProperty() {
	res, err := s.client.GlobalData.GetWebProperty(s.ctx, operations.V3GlobaldataAssetWebpropertyRequest{
		WebpropertyID: testWebPropertyIDs[0],
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

// ---------------------------------------------------------------------------
// GlobalData — Search
// ---------------------------------------------------------------------------

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

func (s *SDKTestSuite) TestGlobalData_ConvertLegacySearchQueries() {
	res, err := s.client.GlobalData.ConvertLegacySearchQueries(s.ctx, operations.V3GlobaldataSearchConvertRequest{
		SearchConvertQueryInputBody: components.SearchConvertQueryInputBody{
			Queries: []string{"parsed.names: censys.io"},
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

// ---------------------------------------------------------------------------
// GlobalData — Tracked Scans
// ---------------------------------------------------------------------------

func (s *SDKTestSuite) TestGlobalData_TrackedScans() {
	createRes, err := s.client.GlobalData.CreateTrackedScan(s.ctx, operations.V3GlobaldataScansRescanRequest{
		ScansRescanInputBody: components.ScansRescanInputBody{
			Target: components.CreateScansRescanInputBodyTargetOne(components.One{
				ServiceID: components.TargetServiceID{
					IP:                "1.1.1.1",
					Port:              80,
					Protocol:          "HTTP",
					TransportProtocol: components.TargetTransportProtocolTCP,
				},
			}),
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), createRes)

	scanID := createRes.GetResponseEnvelopeTrackedScan().GetResult().GetTrackedScanID()
	require.NotNil(s.T(), scanID)

	getRes, err := s.client.GlobalData.GetTrackedScan(s.ctx, operations.V3GlobaldataScansGetRequest{
		ScanID: *scanID,
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), getRes)
}

// ---------------------------------------------------------------------------
// Collections — Full CRUD + Search
// ---------------------------------------------------------------------------

func (s *SDKTestSuite) TestCollections_List() {
	pageSize := int64(2)
	res, err := s.client.Collections.List(s.ctx, operations.V3CollectionsCrudListRequest{
		PageSize: &pageSize,
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestCollections_CRUD() {
	createRes, err := s.client.Collections.Create(s.ctx, operations.V3CollectionsCrudCreateRequest{
		CrudCreateInputBody: &components.CrudCreateInputBody{
			Name:        "SDK Smoke Test Collection",
			Description: strPtr("Created by Go SDK smoke tests"),
			Query:       "host.services.protocol='SSH' and host.location.country = 'Netherlands' and host.services.port = 9100 and host.autonomous_system.name = 'WORLDSTREAM'",
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), createRes)

	collectionUID := createRes.GetResponseEnvelopeCollection().GetResult().GetID()
	s.T().Cleanup(func() {
		_, _ = s.client.Collections.Delete(s.ctx, operations.V3CollectionsCrudDeleteRequest{
			CollectionUID: collectionUID,
		})
	})

	s.Run("Get", func() {
		res, err := s.client.Collections.Get(s.ctx, operations.V3CollectionsCrudGetRequest{
			CollectionUID: collectionUID,
		})
		require.NoError(s.T(), err)
		require.NotNil(s.T(), res)
	})

	s.Run("Update", func() {
		res, err := s.client.Collections.Update(s.ctx, operations.V3CollectionsCrudUpdateRequest{
			CollectionUID: collectionUID,
			CrudUpdateInputBody: &components.CrudUpdateInputBody{
				Name:        "Updated SDK Smoke Test",
				Description: strPtr("Updated description"),
				Query:       "host.services.protocol='SSH' and host.location.country = 'Netherlands' and host.services.port = 9100 and host.autonomous_system.name = 'WORLDSTREAM'",
			},
		})
		require.NoError(s.T(), err)
		require.NotNil(s.T(), res)

		getRes, err := s.client.Collections.Get(s.ctx, operations.V3CollectionsCrudGetRequest{
			CollectionUID: collectionUID,
		})
		require.NoError(s.T(), err)
		require.Equal(s.T(), "Updated description", getRes.GetResponseEnvelopeCollection().GetResult().GetDescription())
	})

	s.Run("ListEvents", func() {
		res, err := s.client.Collections.ListEvents(s.ctx, operations.V3CollectionsListEventsRequest{
			CollectionUID: collectionUID,
		})
		require.NoError(s.T(), err)
		require.NotNil(s.T(), res)
	})

	s.Run("Aggregate", func() {
		res, err := s.client.Collections.Aggregate(s.ctx, operations.V3CollectionsSearchAggregateRequest{
			CollectionUID: collectionUID,
			SearchAggregateInputBody: components.SearchAggregateInputBody{
				Field:           "host.autonomous_system.name",
				NumberOfBuckets: 10,
				Query:           "host.services.labels.value = 'REMOTE_ACCESS'",
			},
		})
		require.NoError(s.T(), err)
		require.NotNil(s.T(), res)
	})

	s.Run("Search", func() {
		res, err := s.client.Collections.Search(s.ctx, operations.V3CollectionsSearchQueryRequest{
			CollectionUID: collectionUID,
			SearchQueryInputBody: components.SearchQueryInputBody{
				Query: "host.services.labels.value = 'REMOTE_ACCESS'",
			},
		})
		require.NoError(s.T(), err)
		require.NotNil(s.T(), res)
	})

	s.Run("Delete", func() {
		res, err := s.client.Collections.Delete(s.ctx, operations.V3CollectionsCrudDeleteRequest{
			CollectionUID: collectionUID,
		})
		require.NoError(s.T(), err)
		require.NotNil(s.T(), res)

		_, err = s.client.Collections.Get(s.ctx, operations.V3CollectionsCrudGetRequest{
			CollectionUID: collectionUID,
		})
		require.Error(s.T(), err)
	})
}

// ---------------------------------------------------------------------------
// Account Management — Organization
// ---------------------------------------------------------------------------

func (s *SDKTestSuite) TestAccountManagement_GetOrganizationDetails() {
	includeCounts := true
	res, err := s.client.AccountManagement.GetOrganizationDetails(s.ctx, operations.V3AccountmanagementOrgDetailsRequest{
		OrganizationID:      s.orgID,
		IncludeMemberCounts: &includeCounts,
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestAccountManagement_GetOrganizationCredits() {
	res, err := s.client.AccountManagement.GetOrganizationCredits(s.ctx, operations.V3AccountmanagementOrgCreditsRequest{
		OrganizationID: s.orgID,
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestAccountManagement_GetOrganizationCreditUsage() {
	startDate := types.MustNewDateFromString(thirtyDaysAgo())
	res, err := s.client.AccountManagement.GetOrganizationCreditUsage(s.ctx, operations.V3AccountmanagementOrgCreditsUsageRequest{
		OrganizationID: s.orgID,
		StartDate:      startDate,
		Granularity:    operations.GranularityDaily,
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

// ---------------------------------------------------------------------------
// Account Management — Members
// ---------------------------------------------------------------------------

func (s *SDKTestSuite) TestAccountManagement_ListOrganizationMembers() {
	pageSize := 5
	res, err := s.client.AccountManagement.ListOrganizationMembers(s.ctx, operations.V3AccountmanagementListOrgMembersRequest{
		OrganizationID: s.orgID,
		PageSize:       &pageSize,
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestAccountManagement_GetMemberCreditUsage() {
	members, err := s.client.AccountManagement.ListOrganizationMembers(s.ctx, operations.V3AccountmanagementListOrgMembersRequest{
		OrganizationID: s.orgID,
	})
	require.NoError(s.T(), err)

	memberList := members.GetResponseEnvelopeOrganizationMembersList().GetResult().GetMembers()
	require.NotEmpty(s.T(), memberList)

	startDate := types.MustNewDateFromString(thirtyDaysAgo())
	res, err := s.client.AccountManagement.GetMemberCreditUsage(s.ctx, operations.V3AccountmanagementMemberCreditsUsageRequest{
		OrganizationID: s.orgID,
		UserID:         memberList[0].GetUID(),
		StartDate:      startDate,
		Granularity:    operations.QueryParamGranularityDaily,
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

// ---------------------------------------------------------------------------
// Account Management — User (self)
// ---------------------------------------------------------------------------

func (s *SDKTestSuite) TestAccountManagement_GetUserCredits() {
	res, err := s.client.AccountManagement.GetUserCredits(s.ctx)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestAccountManagement_GetUserCreditsUsage() {
	startDate := types.MustNewDateFromString(thirtyDaysAgo())
	res, err := s.client.AccountManagement.GetUserCreditsUsage(s.ctx, operations.V3AccountmanagementUserCreditsUsageRequest{
		StartDate:   startDate,
		Granularity: operations.V3AccountmanagementUserCreditsUsageQueryParamGranularityDaily,
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

// ---------------------------------------------------------------------------
// Threat Hunting
// ---------------------------------------------------------------------------

func (s *SDKTestSuite) TestThreatHunting_ValueCounts() {
	res, err := s.client.ThreatHunting.ValueCounts(s.ctx, operations.V3ThreathuntingValueCountsRequest{
		SearchValueCountsInputBody: components.SearchValueCountsInputBody{
			AndCountConditions: []components.CountCondition{
				{
					FieldValuePairs: []components.FieldValuePair{
						{Field: "host.services.port", Value: "80"},
					},
				},
			},
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestThreatHunting_GetHostObservationsWithCertificate() {
	res, err := s.client.ThreatHunting.GetHostObservationsWithCertificate(s.ctx, operations.V3ThreathuntingGetHostObservationsWithCertificateRequest{
		CertificateID: testCertIDs[0],
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestThreatHunting_ListThreats() {
	res, err := s.client.ThreatHunting.ListThreats(s.ctx, operations.V3ThreathuntingThreatsListRequest{})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *SDKTestSuite) TestThreatHunting_TrackedScans() {
	createRes, err := s.client.ThreatHunting.CreateTrackedScan(s.ctx, operations.V3ThreathuntingScansDiscoveryRequest{
		ScansDiscoveryInputBody: components.ScansDiscoveryInputBody{
			Target: components.CreateScansDiscoveryInputBodyTargetTarget1(components.Target1{
				HostPort: components.HostPort{
					IP:   "1.1.1.1",
					Port: 443,
				},
			}),
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), createRes)

	scanID := createRes.GetResponseEnvelopeTrackedScan().GetResult().GetTrackedScanID()
	require.NotNil(s.T(), scanID)

	getRes, err := s.client.ThreatHunting.GetTrackedScanThreatHunting(s.ctx, operations.V3ThreathuntingScansGetRequest{
		ScanID: *scanID,
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), getRes)
}

func strPtr(s string) *string {
	return &s
}

func mustParseTime(value string) time.Time {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		panic(err)
	}
	return t
}

func thirtyDaysAgo() string {
	return time.Now().AddDate(0, 0, -30).Format("2006-01-02")
}
