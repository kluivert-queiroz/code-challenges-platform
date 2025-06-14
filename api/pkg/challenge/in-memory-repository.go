package challenge

import "errors"

var challenges = []Challenge{
	{ID: "1", Name: "Sum Two Numbers", Description: "Sum two numbers", TestCases: []ChallengeTestcases{
		{Input: "1 2", ExpectedOutput: "3"},
		{Input: "2 3", ExpectedOutput: "5"},
		{Input: "3 4", ExpectedOutput: "7"},
		{Input: "4 5", ExpectedOutput: "9"},
		{Input: "5 6", ExpectedOutput: "11"},
		{Input: "6 7", ExpectedOutput: "13"},
		{Input: "7 8", ExpectedOutput: "15"},
		{Input: "8 9", ExpectedOutput: "17"},
		{Input: "9 10", ExpectedOutput: "19"},
	},
		BoilerplateCode: `const readline = require("readline");const rl = readline.createInterface({  input: process.stdin,  output: process.stdout,});let input:string[] = [];rl.on('line', (line:string) => {  input.push(...line.split(' '));});rl.on('close', () => {  console.log(sum(parseInt(input[0]), parseInt(input[1])));});`,
	},
}

type inMemoryChallengeRepository struct {
}

func NewInMemoryChallengeRepository() ChallengeRepository {
	return &inMemoryChallengeRepository{}
}

func (r *inMemoryChallengeRepository) GetChallenge(id string) (Challenge, error) {
	for _, challenge := range challenges {
		if challenge.ID == id {
			return challenge, nil
		}
	}
	return Challenge{}, errors.New("challenge not found")
}

func (r *inMemoryChallengeRepository) GetChallenges() ([]Challenge, error) {
	return challenges, nil
}
