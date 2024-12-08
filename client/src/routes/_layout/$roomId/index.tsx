import api from "@/api/axiosInstance";
import ChatCard from "@/components/chatCard";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import { UserItem } from "@/components/userItem";
import { Room } from "@/types/room";
import { useMutation, useSuspenseQuery } from "@tanstack/react-query";
import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { DoorClosed, Headphones, Send } from "lucide-react";
import { useEffect, useRef, useState } from "react";

export const Route = createFileRoute("/_layout/$roomId/")({
  component: RouteComponent,
  pendingComponent: () => {
    return <div className="px-6 py-5 text-lg font-bold">Loading...</div>;
  },
  errorComponent: () => {
    return <div className="px-6 py-5 text-lg font-bold">Error</div>;
  },
});

function RouteComponent() {
  const [text, setText] = useState("");
  const textareaRef = useRef<HTMLTextAreaElement>(null);
  const { roomId }: { roomId: string } = Route.useParams();
  const navigate = useNavigate();
  // const eventSource = useRef<EventSource>();

  const { data } = useSuspenseQuery({
    queryKey: ["room", roomId],
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
    mutationKey: ["leaveMember"],
    mutationFn: async () => {
      const token = localStorage.getItem("code");
      const res = await api.delete(`/v1/rooms/${roomId}/members`, {
        headers: {
          Authorization: token,
        },
      });
      return res.data;
    },
    onSuccess: () => {
      navigate({ to: "/home" });
    },
    onError: (error) => {
      console.error(error);
    },
  });

  const mutationMessage = useMutation({
    mutationKey: ["message"],
    mutationFn: async () => {
      const token = localStorage.getItem("code");
      const userId = localStorage.getItem("userId");
      const res = await api.post(
        `/v1/rooms/${roomId}/chat`,
        {
          message: text,
          roomID: roomId,
          userId: userId,
        },
        {
          headers: {
            Authorization: token,
          },
        }
      );
      console.log(res.data);
      return res.data;
    },
    onError: (error) => {
      console.error(error);
    },
  });

  useEffect(() => {
    const token = localStorage.getItem("code");
    const apiUrl = import.meta.env.VITE_API_BASE_URL || "http://localhost:8080";
    const evtSource = new EventSource(
      `${apiUrl}/v1/rooms/${roomId}/chat?auth=${token}`
    );

    console.log(evtSource);

    // メッセージを受信したときの処理
    evtSource.onmessage = (event) => {
      console.log(event.data);
    };

    // コンポーネントがアンマウントされたときに接続を閉じる
    return () => {
      evtSource.close();
    };
  }, [roomId]);

  useEffect(() => {
    const textarea = textareaRef.current;
    if (textarea) {
      // デフォルトの高さにリセット
      textarea.style.height = "auto";

      // スクロールの高さを計算
      const scrollHeight = textarea.scrollHeight;

      // 最大3行（約80px）まで拡大
      const maxHeight = 80;

      if (scrollHeight <= maxHeight) {
        textarea.style.height = `${scrollHeight}px`;
      } else {
        textarea.style.height = `${maxHeight}px`;
        textarea.style.overflowY = "auto";
      }
    }
  }, [text]);

  // const joinMeeting = async () => {
  //   // Fetch the meeting and attendee data from your server application
  //   const response = await fetch("/my-server");
  //   const data = await response.json();

  //   // Initalize the `MeetingSessionConfiguration`
  //   const meetingSessionConfiguration = new MeetingSessionConfiguration(
  //     data.Meeting,
  //     data.Attendee
  //   );

  //   // Create a `MeetingSession` using `join()` function with the `MeetingSessionConfiguration`
  //   await meetingManager.join(meetingSessionConfiguration);

  //   // At this point you could let users setup their devices, or by default
  //   // the SDK will select the first device in the list for the kind indicated
  //   // by `deviceLabels` (the default value is DeviceLabels.AudioAndVideo)

  //   // Start the `MeetingSession` to join the meeting
  //   await meetingManager.start();
  // };

  return (
    <div className="px-6 py-5 flex flex-col h-[calc(100vh-64px)]">
      <div className="flex flex-col md:flex-row gap-4 h-full">
        <div className="flex flex-1 flex-col gap-4">
          <div className="flex justify-between items-center w-full">
            <h2 className="text-lg font-bold">{data.name}</h2>
            <Button size={"icon"} variant={"outline"}>
              <Headphones />
            </Button>
          </div>
          <div className="flex-grow overflow-y-auto flex flex-col-reverse px-2">
            <ChatCard />
            <ChatCard />
            <ChatCard />
            <ChatCard />
            <ChatCard />
          </div>
          <div className="flex gap-1 items-end">
            <Textarea
              rows={1}
              ref={textareaRef}
              value={text}
              onChange={(e) => setText(e.target.value)}
              placeholder="メッセージを入力..."
              className="focus-visible:ring-1 focus-visible:ring-offset-0 resize-none min-h-fit max-h-20"
            />
            <Button
              size={"icon"}
              disabled={!text.trim()}
              onClick={() => mutationMessage.mutate()}
            >
              <Send />
            </Button>
          </div>
        </div>
        <div className="w-full md:w-56 border-l pl-4 h-full flex flex-col justify-between">
          <div>
            <div>
              <h3 className="font-semibold pb-2">概要</h3>
              <p className="text-muted-foreground text-sm">
                {data.description}
              </p>
            </div>
            <div className="py-6">
              <h3 className="font-semibold pb-2">メンバー</h3>
              <div className="flex flex-col gap-2">
                {data.members.map((member) => (
                  <UserItem
                    user={{ name: member.name, icon: member.icon }}
                    key={member.id}
                  />
                ))}
              </div>
            </div>
          </div>
          <Button
            variant={"ghost"}
            size={"sm"}
            onClick={() => mutation.mutate()}
          >
            <DoorClosed /> このルームから抜ける
          </Button>
        </div>
      </div>
    </div>
  );
}
