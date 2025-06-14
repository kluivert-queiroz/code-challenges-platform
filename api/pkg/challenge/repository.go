package challenge

type ChallengeRepository interface {
	GetChallenge(id string) (Challenge, error)
	GetChallenges() ([]Challenge, error)
}
