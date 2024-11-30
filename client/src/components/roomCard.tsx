import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Badge } from "./ui/badge";
import { Link } from "@tanstack/react-router";

export default function RoomCard() {
  return (
    <Link to="/$roomId/description" params={{ roomId: "1" }}>
      <Card className="hover:bg-sidebar">
        <CardHeader className="flex-row items-center justify-between space-y-0">
          <CardTitle className="text-xl">ルーム名</CardTitle>
          <CardDescription className="text-base">4人</CardDescription>
        </CardHeader>
        <CardContent className="flex flex-wrap gap-2">
          <Badge className="text-sm">Frontend</Badge>
          <Badge className="text-sm">React</Badge>
          <Badge className="text-sm">TypeScript</Badge>
        </CardContent>
      </Card>
    </Link>
  );
}
