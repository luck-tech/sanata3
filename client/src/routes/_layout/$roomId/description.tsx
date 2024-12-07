import api from "@/api/axiosInstance";
import { Button } from "@/components/ui/button";
import { UserItem } from "@/components/userItem";
import { Room } from "@/types/room";
import { useMutation, useSuspenseQuery } from "@tanstack/react-query";
import { createFileRoute, Navigate } from "@tanstack/react-router";

export const Route = createFileRoute("/_layout/$roomId/description")({
  component: RouteComponent,
  errorComponent: () => <Navigate to="/" />,
});

function RouteComponent() {
  const { roomId }: { roomId: string } = Route.useParams();

  const { data } = useSuspenseQuery({
    queryKey: ["room-description"],
    queryFn: async (): Promise<Room> => {
      const token = localStorage.getItem("code");
      const res = await api.get(`/v1/rooms/${roomId}`, {
        headers: {
          Authorization: token,
        },
      });
      return res.data;
    },
  });

  const mutation = useMutation({
    mutationFn: async () => {
      const token = localStorage.getItem("code");
      const res = await api.post(`/v1/rooms/${roomId}/members`, {
        headers: {
          Authorization: token,
        },
      });
      return res.data;
    },
    onError: (error) => {
      console.error(error);
    },
  });

  const userId = localStorage.getItem("userId");
  const isIncluded = data.members.some((member) => member.id === userId);
  if (isIncluded) {
    return <Navigate to="/$roomId" params={{ roomId: roomId }} />;
  }

  const owner = data.members.find((member) => member.id === data.ownerId);
  const members = data.members.filter((member) => member.id !== data.ownerId);
  if (!owner) return;

  return (
    <div className="max-w-screen-lg mx-auto w-full px-6 py-12">
      <div className="w-full flex justify-between items-center gap-5 flex-wrap">
        <h2 className="text-2xl text-primary font-bold">{data.name}</h2>
        <Button onClick={() => mutation.mutate()}>このルームに入る</Button>
      </div>
      <p className="py-8">{data.description}</p>

      <div className="py-6 space-y-3">
        <h3 className="font-bold">作成者</h3>
        <UserItem user={{ name: owner.name, icon: owner.icon }} />
      </div>
      <div className="py-6 space-y-3">
        <div className="flex gap-2 items-center">
          <h3 className="font-bold">参加者</h3>
          <div className="w-6 h-6 flex justify-center items-center leading-none text-sm rounded-full bg-primary text-primary-foreground">
            {members.length}
          </div>
        </div>
        <div className="flex flex-wrap gap-8">
          {members.map((member) => (
            <UserItem
              user={{ name: member.name, icon: member.icon }}
              key={member.id}
            />
          ))}
        </div>
      </div>
    </div>
  );
}
