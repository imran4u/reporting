package reporting

import (
	"errors"
	"fmt"
	"testing"

	"github.com/imran4u/reporting/reportstore"
	mockReportStore "github.com/imran4u/reporting/reportstore/mock"
	mockuuid "github.com/imran4u/reporting/uuid/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type ReportManagerSuit struct {
	suite.Suite
	*require.Assertions

	ctrl              *gomock.Controller
	mockReportStore   *mockReportStore.MockStore
	mockUUIDGenerator *mockuuid.MockGenerator

	manager *ReportManager
}

func TestReportManagerSuit(t *testing.T) {
	suite.Run(t, new(ReportManagerSuit))
}

// Run before each and every test
func (s *ReportManagerSuit) SetupTest() {
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())

	s.mockReportStore = mockReportStore.NewMockStore(s.ctrl)
	s.mockUUIDGenerator = mockuuid.NewMockGenerator(s.ctrl)

	s.manager = NewReportManager(s.mockUUIDGenerator, s.mockReportStore)
}

// Run after each test
func (s *ReportManagerSuit) TearDownTest(t *testing.T) {
	s.ctrl.Finish()
}

func (s *ReportManagerSuit) TestCreateReport() {
	reportId := "reportId"
	userId := "testId"
	title := "titile"

	s.mockUUIDGenerator.EXPECT().Generate().Return(reportId).Times(1)

	s.mockReportStore.EXPECT().CreateReport(gomock.Eq(reportstore.CreateReportRequest{
		ReportID: reportId,
		UserID:   userId,
		Status:   reportstore.ReportStatusPending.String(),
		Title:    title,
	})).Return(nil).Times(1)

	actualResponse, err := s.manager.CreateReport(CreateReportRequest{
		UserID: userId,
		Title:  title,
	})
	s.NoError(err)

	expectedResponse := CreateReportResponse{
		ReportID: reportId,
	}
	s.Equal(actualResponse, expectedResponse)

}

func (s *ReportManagerSuit) TestCreateReportError() {
	s.mockUUIDGenerator.EXPECT().Generate().Return("reportId").Times(1)
	createError := errors.New("my custom error")
	s.mockReportStore.EXPECT().CreateReport(gomock.Any()).Return(createError).Times(1)

	_, err := s.manager.CreateReport(CreateReportRequest{})
	fmt.Println(createError, ": ", err)
	s.Equal(createError, err)
}
