import api from "@/api/axiosInstance";
import { useEffect, useState } from "react";

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

type Room = {
  aimTags: AimTag[];
  description: string;
  members: Member[];
  name: string;
  ownerId: string;
  roomId: string;
};

export function useGetRooms() {
  const [rooms, setRooms] = useState<Room[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);

  useEffect(() => {
    const token = localStorage.getItem("code");
    api
      .get("/v1/rooms", {
        headers: {
          Authorization: token ?? "",
        },
      })
      .then((res) => {
        setRooms(res.data.rooms);
        setLoading(false);
      })
      .catch((err) => {
        setError(err as Error);
        setLoading(false);
      });
  }, []);

  return { rooms, loading, error };
}
