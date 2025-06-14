package challenge

type ChallengeTestcases struct {
	Input          string `json:"input"`
	ExpectedOutput string `json:"expectedOutput"`
}

type ChallengeResult struct {
	Input          string `json:"input"`
	Output         string `json:"output"`
	ExpectedOutput string `json:"expectedOutput"`
	Passed         bool   `json:"passed"`
}

type Challenge struct {
	ID              string               `json:"id"`
	Name            string               `json:"name"`
	Description     string               `json:"description"`
	TestCases       []ChallengeTestcases `json:"testCases"`
	BoilerplateCode string               `json:"boilerplateCode"`
}
