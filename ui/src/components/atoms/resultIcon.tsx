import { IconCheck, IconX } from "@tabler/icons-react";

export const ResultIcon = ({ passed }: { passed: boolean }) => {
  return passed ? (
    <IconCheck className="w-4 h-4 text-green-500" />
  ) : (
    <IconX className="w-4 h-4 text-red-500" />
  );
};
