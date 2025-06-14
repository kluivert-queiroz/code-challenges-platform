import type { Challenge } from "@/types";
import { useEffect, useState } from "react";
import { mockedChallenge } from "./useChallenges";

export const useChallenge = ({ challengeId }: { challengeId: string }) => {
  const [challenge, setChallenge] = useState<Challenge | null>(null);

  useEffect(() => {
    setChallenge(mockedChallenge);
  }, [challengeId]);

  return { challenge };
};
