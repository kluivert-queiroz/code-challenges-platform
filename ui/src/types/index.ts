export interface ChallengeTestcase {
  input: string;
  expectedOutput: string;
}

export interface Challenge {
  id: string;
  name: string;
  description: string;
  boilerplateCode: string;
  testCases: ChallengeTestcase[];
  defaultCode: string;
}

export interface CodeSubmission {
  language: string;
  code: string;
  challengeId: string;
}

export interface ChallengeResult {
  input: string;
  expectedOutput: string;
  output: string;
  passed: boolean;
}
