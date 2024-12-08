import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Badge } from "./ui/badge";
import { Link } from "@tanstack/react-router";

type AimTag = {
  id: number;
  name: string;
};

type Member = {
  description: string;
  icon: string;
  id: string;
  name: string;
};

type RoomCardProps = {
  roomId: string;
  name: string;
  members: Member[];
  aimTags: AimTag[];
};

export default function RoomCard({
  roomId,
  name,
  members,
  aimTags,
}: RoomCardProps) {
  return (
    <Link to="/$roomId/description" params={{ roomId }}>
      <Card className="hover:bg-sidebar">
        <CardHeader className="flex-row items-center justify-between space-y-0">
          <CardTitle className="text-xl">{name}</CardTitle>
          <CardDescription className="text-base">
            {members.length}人
          </CardDescription>
        </CardHeader>
        <CardContent className="flex flex-wrap gap-2">
          {aimTags.map((tag) => (
            <Badge key={tag.id} className="text-sm">
              {tag.name}
            </Badge>
          ))}
        </CardContent>
      </Card>
    </Link>
  );
}
