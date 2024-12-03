import { Avatar, AvatarFallback, AvatarImage } from "./ui/avatar";

export const UserItem = () => {
  return (
    <div className="flex items-center gap-2">
      <Avatar>
        <AvatarImage src="https://github.com/shadcn.png" />
        <AvatarFallback>CN</AvatarFallback>
      </Avatar>
      <div className="text-lg break-all">username</div>
    </div>
  );
};
