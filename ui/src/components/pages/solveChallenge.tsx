import { useChallenge } from "@/hooks/useChallenge";
import { Link, useParams } from "@tanstack/react-router";
import { CodeEditor } from "../organisms/codeEditor";
import { useSubmitCode } from "@/hooks/useSubmitCode";
import { TestcasesResult } from "../organisms/testcasesResult";
import { useEffect, useState } from "react";
import { IconArrowLeft } from "@tabler/icons-react";

export const SolveChallenge = () => {
  const { challengeId } = useParams({ from: "/challenges/$challengeId" });
  const { challenge } = useChallenge({ challengeId });
  const {
    mutate: submitCode,
    isPending: isSubmitting,
    data: submission,
    isSuccess,
  } = useSubmitCode();
  const [isTestcasesResultOpen, setIsTestcasesResultOpen] = useState(false);
  useEffect(() => {
    if (isSuccess) {
      setIsTestcasesResultOpen(true);
    }
  }, [isSuccess]);
  const handleSubmit = (code: string) => {
    submitCode({
      language: "typescript",
      code,
      challengeId,
    });
  };

  return (
    <div className="flex flex-row">
      <div className="flex flex-col gap-4 col-span-1 p-4">
        <div className="flex flex-row gap-2 items-center">
          <Link to="/">
            <IconArrowLeft className="w-6 h-6" />
          </Link>
          <h1 className="text-2xl font-bold">{challenge?.name}</h1>
        </div>
        <p className="text-sm text-gray-500">{challenge?.description}</p>
      </div>
      <div className="col-span-auto w-full h-svh">
        <div className="flex flex-col h-full">
          <div className={`grow ${isSuccess ? "h-1/2" : "h-full"}`}>
            <CodeEditor
              language="typescript"
              defaultValue={challenge?.defaultCode ?? ""}
              theme="vs-dark"
              onSubmit={handleSubmit}
              isSubmitting={isSubmitting}
            />
          </div>
          <div className="grow-0 max-h-1/6 overflow-y-scroll resize-y">
            {isSuccess && isTestcasesResultOpen && (
              <TestcasesResult
                results={submission?.data}
                onClose={() => setIsTestcasesResultOpen(false)}
              />
            )}
          </div>
        </div>
      </div>
    </div>
  );
};
