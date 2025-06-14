import type { Challenge } from "@/types";
import { Button } from "../ui/button";
import { IconPlayerPlayFilled } from "@tabler/icons-react";
import { Link } from "@tanstack/react-router";

export const ChallengeBar = ({ challenge }: { challenge: Challenge }) => {
  return (
    <div className="flex flex-row justify-between items-center bg-slate-800 p-4 rounded-md mb-4 gap-4">
      <div className="flex flex-col">
        <p>{challenge.name}</p>
        <p className="text-sm text-white/50 line-clamp-2">
          {challenge.description}
        </p>
      </div>
      <div>
        <Link to="/challenges/$challengeId" params={{ challengeId: challenge.id }}>
          <Button className="cursor-pointer">
            <IconPlayerPlayFilled /> Start
          </Button>
        </Link>
      </div>
    </div>
  );
};
