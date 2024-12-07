import { Avatar, AvatarFallback, AvatarImage } from "./ui/avatar";

export const UserItem = ({
  user,
}: {
  user?: { name: string; icon: string };
}) => {
  return (
    <div className="flex items-center gap-2">
      <Avatar>
        <AvatarImage
          src={user?.icon ? user.icon : "https://github.com/shadcn.png"}
        />
        <AvatarFallback>{user?.name.slice(0, 2)}</AvatarFallback>
      </Avatar>
      <div className="text-lg break-all">
        {user?.name ? user.name : "username"}
      </div>
    </div>
  );
};
