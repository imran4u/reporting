package reportstore

type Store interface {
	CreateReport(r CreateReportRequest) error
}

type CreateReportRequest struct {
	ReportID string
	UserID   string
	Status   string
	Title    string
}
