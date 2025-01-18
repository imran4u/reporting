package reporting

import (
	"github.com/imran4u/reporting/reportstore"
	"github.com/imran4u/reporting/uuid"
)

type ReportManager struct {
	uuidGenerator uuid.Generator
	store         reportstore.Store
}

func NewReportManager(gen uuid.Generator, store reportstore.Store) *ReportManager {
	return &ReportManager{
		uuidGenerator: gen,
		store:         store,
	}
}

func (m *ReportManager) CreateReport(request CreateReportRequest) (response CreateReportResponse, err error) {
	reportID := m.uuidGenerator.Generate()

	r := reportstore.CreateReportRequest{
		ReportID: reportID,
		UserID:   request.UserID,
		Status:   reportstore.ReportStatusPending.String(),
		Title:    request.Title,
	}
	err = m.store.CreateReport(r)

	response.ReportID = r.ReportID
	return
}
