import { ChallengeBar } from "../organisms/challengeBar";
import { Card, CardContent, CardHeader, CardTitle } from "../ui/card";
import { useChallenges } from "@/hooks/useChallenges";

export const BrowseChallenges = () => {
  const { challenges } = useChallenges();
  return (
    <div className="flex flex-row min-h-screen justify-center items-center">
      <Card className="w-full max-w-md">
        <CardHeader>
          <CardTitle>Choose a challenge!</CardTitle>
        </CardHeader>
        <CardContent>
          {challenges.map((challenge) => (
            <ChallengeBar key={challenge.id} challenge={challenge} />
          ))}
        </CardContent>
      </Card>
    </div>
  );
};
