import type { ChallengeResult } from "@/types";
import { ResultIcon } from "../atoms/resultIcon";
import { IconX } from "@tabler/icons-react";

export const TestcasesResult = ({
  results,
  onClose,
}: {
  results: ChallengeResult[];
  onClose: () => void;
}) => {
  const passedCount = results.filter((result) => result.passed).length;
  return (
    <div className="flex flex-col relative">
      <div className="flex flex-col gap-4 bg-neutral-900 p-4">
        <div className="flex flex-row justify-between">
          <h2 className="text-2xl font-bold">
            Testcases Result ({passedCount} / {results.length})
          </h2>
          <IconX className="w-6 h-6 cursor-pointer" onClick={onClose} />
        </div>
        {results.map((result) => (
          <div key={result.input} className="flex flex-row gap-4">
            <div className="flex flex-col gap-2">
              <p className="text-sm text-gray-500">Input</p>
              <p className="text-sm text-gray-500">{result.input}</p>
            </div>
            <div className="flex flex-col gap-2">
              <p className="text-sm text-gray-500">Expected Output</p>
              <p className="text-sm text-gray-500">{result.expectedOutput}</p>
            </div>
            <div className="flex flex-col gap-2">
              <p className="text-sm text-gray-500">Output</p>
              <p className="text-sm text-gray-500">{result.output}</p>
            </div>
            <div className="flex flex-col gap-2">
              <p className="text-sm text-gray-500">Passed</p>
              <p className="text-sm text-gray-500 flex justify-center items-center">
                <ResultIcon passed={result.passed} />
              </p>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};
