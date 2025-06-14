package submission

type Language string

const (
	LanguageNode Language = "typescript"
)

type Submission struct {
	Code        string   `json:"code" validate:"required"`
	Language    Language `json:"language" validate:"required,oneof=typescript"`
	ChallengeID string   `json:"challengeId" validate:"required"`
}
