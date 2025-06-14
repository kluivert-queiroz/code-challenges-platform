import { api } from "@/lib/axios";
import type { ChallengeResult, CodeSubmission } from "@/types";
import { useMutation } from "@tanstack/react-query";

export const useSubmitCode = () => {
  return useMutation({
    mutationFn: (codeSubmission: CodeSubmission) => {
      return api.post<ChallengeResult[]>("/submissions", codeSubmission);
    },
  });
};
