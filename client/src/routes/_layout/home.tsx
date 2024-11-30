import RoomCard from "@/components/roomCard";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_layout/home")({
  component: RouteComponent,
});

function RouteComponent() {
  const { search }: { search: string } = Route.useSearch();
  return (
    <div className="max-w-screen-lg mx-auto w-full px-6">
      <h2 className="text-xl py-8 text-primary">
        {search ? "検索結果" : "おすすめ"}
      </h2>
      <div className="grid gap-6">
        <RoomCard />
        <RoomCard />
        <RoomCard />
        <RoomCard />
      </div>
    </div>
  );
}
