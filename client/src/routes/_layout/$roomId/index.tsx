import ChatCard from "@/components/chatCard";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import { UserItem } from "@/components/userItem";
import { createFileRoute } from "@tanstack/react-router";
import { Headphones, Send } from "lucide-react";
import { useEffect, useRef, useState } from "react";

export const Route = createFileRoute("/_layout/$roomId/")({
  component: RouteComponent,
});

function RouteComponent() {
  const [text, setText] = useState("");
  const textareaRef = useRef<HTMLTextAreaElement>(null);
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
  return (
    <div className="px-6 py-5 flex flex-col h-[calc(100vh-64px)]">
      <div className="flex flex-col md:flex-row gap-4 h-full">
        <div className="flex flex-1 flex-col gap-4">
          <div className="flex justify-between items-center w-full">
            <h2 className="text-lg font-bold">ルーム名</h2>
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
            <Button size={"icon"} disabled={!text.trim()}>
              <Send />
            </Button>
          </div>
        </div>
        <div className="w-full md:w-56 border-l pl-4">
          <div>
            <h3 className="font-semibold pb-2">概要</h3>
            <p className="text-muted-foreground text-sm">
              ほんとうにそのいるかのかたちのおかしいことは、二人のうしろで聞こえました。鳥捕りは、だまって見ていました。青年はなんとも言えずさびしい気がして、そっちに祈ってくれました。
            </p>
          </div>
          <div className="py-6">
            <h3 className="font-semibold pb-2">メンバー</h3>
            <div className="flex flex-col gap-2">
              <UserItem />
              <UserItem />
              <UserItem />
              <UserItem />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
