import type { Challenge } from "@/types";
import { useState } from "react";

import { useEffect } from "react";
export const mockedChallenge = {
  id: "1",
  name: "Sum Two Numbers",
  description:
    "It's a simple challenge to sum two numbers. You will be given two numbers and you need to return the sum of them.",
  testCases: [
    {
      input: "1 2",
      expectedOutput: "3",
    },
    {
      input: "2 3",
      expectedOutput: "5",
    },
  ],
  boilerplateCode:
    "const readline = require(\"readline\");const rl = readline.createInterface({  input: process.stdin,  output: process.stdout,});let input:string[] = [];rl.on('line', (line:string) => {  input.push(...line.split(' '));});rl.on('close', () => {  console.log(sum(parseInt(input[0]), parseInt(input[1])));});",
  defaultCode: "const sum = (a: number, b: number) => {\n//your code here\n}",
};
export const useChallenges = () => {
  const [challenges, setChallenges] = useState<Challenge[]>([]);

  useEffect(() => {
    const challenges = [mockedChallenge];
    setChallenges(challenges);
  }, []);
  return { challenges };
};
