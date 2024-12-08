import RoomCard from "@/components/roomCard";
import { createFileRoute } from "@tanstack/react-router";
import { useGetRooms } from "@/hooks/useGetRooms";

export const Route = createFileRoute("/_layout/home")({
  component: RouteComponent,
});

function RouteComponent() {
  const { search }: { search: string } = Route.useSearch();
  const { rooms, loading, error } = useGetRooms();

  if (loading) return <div>Loading...</div>;
  if (error) return <div>エラーが発生しました: {error.message}</div>;

  return (
    <div className="max-w-screen-lg mx-auto w-full px-6">
      <h2 className="text-xl py-8 text-primary">
        {search ? "検索結果" : "おすすめ"}
      </h2>
      <div className="grid gap-6">
        {rooms.map((room) => (
          <RoomCard
            key={room.roomId}
            roomId={room.roomId}
            name={room.name}
            members={room.members}
            aimTags={room.aimTags}
          />
        ))}
      </div>
    </div>
  );
}
