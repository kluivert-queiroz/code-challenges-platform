package runner

import "kluivert-queiroz/api/pkg/submission"

type LanguageSpecifications struct {
	Image         string
	Command       string
	ContainerName string
	Args          []string
	Extension     string
}

var languageSpecifications = map[submission.Language]LanguageSpecifications{
	submission.LanguageNode: {
		Image:         "node:23-alpine",
		Command:       "node",
		ContainerName: "node-code-executor",
		Args:          []string{"--experimental-strip-types", "--no-warnings"},
		Extension:     "ts",
	},
}
