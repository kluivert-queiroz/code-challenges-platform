package runner

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"kluivert-queiroz/api/pkg/challenge"
	"kluivert-queiroz/api/pkg/docker"
	"kluivert-queiroz/api/pkg/submission"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

type Runner interface {
	Assert(code string, language submission.Language, challenge challenge.Challenge) ([]challenge.ChallengeResult, error)
	RemoveContainers(ctx context.Context) error
}

type runnerService struct {
}

func NewRunnerService() Runner {
	return &runnerService{}
}

func (e *runnerService) RemoveContainers(ctx context.Context) error {
	return docker.RemoveContainers(ctx, []string{
		languageSpecifications[submission.LanguageNode].ContainerName,
	})
}

func (e *runnerService) Assert(code string, language submission.Language, c challenge.Challenge) ([]challenge.ChallengeResult, error) {
	asserts := c.TestCases
	filename := fmt.Sprintf("/tmp/submissions/%s.%s", uuid.New().String(), languageSpecifications[language].Extension)
	boilerplateCode := c.BoilerplateCode

	if err := os.WriteFile(filename, []byte(boilerplateCode+"\n"+code), 0644); err != nil {
		return nil, err
	}

	if err := docker.EnsureContainerIsRunning(
		context.Background(),
		languageSpecifications[language].ContainerName,
		languageSpecifications[language].Image,
	); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	results := make([]challenge.ChallengeResult, len(asserts))
	resultsChannel := make(chan challenge.ChallengeResult, len(asserts))

	for _, assert := range asserts {
		go func(assert challenge.ChallengeTestcases) {
			response, err := assertCode(ctx, language, filename, assert)
			if err != nil {
				resultsChannel <- challenge.ChallengeResult{
					Input:          assert.Input,
					ExpectedOutput: assert.ExpectedOutput,
					Output:         err.Error(),
					Passed:         false,
				}
				return
			}
			resultsChannel <- response
		}(assert)
	}

	for i := range asserts {
		results[i] = <-resultsChannel
	}

	if err := os.Remove(filename); err != nil {
		return nil, err
	}

	return results, nil
}

func assertCode(ctx context.Context, language submission.Language, codePath string, assert challenge.ChallengeTestcases) (challenge.ChallengeResult, error) {
	execCmd := docker.ExecuteCodeCommand(
		ctx,
		languageSpecifications[language].ContainerName,
		languageSpecifications[language].Command,
		languageSpecifications[language].Args,
		codePath,
	)

	var out bytes.Buffer
	var stderr bytes.Buffer
	execCmd.Stdout = &out
	execCmd.Stderr = &stderr
	execCmd.Stdin = bytes.NewBufferString(assert.Input)

	if err := execCmd.Run(); err != nil {
		log.Error(stderr.String())
		return challenge.ChallengeResult{}, errors.New(stderr.String())
	}

	output := out.String()
	if strings.Contains(output, "\n") {
		output = strings.Trim(strings.Join(strings.Split(output, "\n"), " "), " ")
	}

	return challenge.ChallengeResult{
		Input:          assert.Input,
		ExpectedOutput: assert.ExpectedOutput,
		Output:         output,
		Passed:         output == assert.ExpectedOutput,
	}, nil
}
