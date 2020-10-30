package domain

type Media struct {
	Code string
	Url  string
}

type SupportThumbnail struct {
	Code         string
	Value  string
	Designation  string
	Type         string
	SupportModel string
	Media        []Media
}

type SupportThumbnailRepository interface {
	GetSupportThumbnails() []SupportThumbnail
}
