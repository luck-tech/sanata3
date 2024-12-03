import { Avatar, AvatarFallback, AvatarImage } from "./ui/avatar";

export default function ChatCard() {
  return (
    <div className="py-4 flex gap-3">
      <Avatar>
        <AvatarImage src="https://github.com/shadcn.png" />
        <AvatarFallback>CN</AvatarFallback>
      </Avatar>
      <div>
        <div className="flex gap-1 md:gap-2 items-baseline leading-none flex-wrap">
          <p className="font-semibold text-sm">ユーザーネーム</p>
          <p className="text-muted-foreground text-sm">
            {new Date().toLocaleString()}
          </p>
        </div>
        <p className="whitespace-pre-wrap break-words">
          一昨日の第二限ころなんか、なぜ燈台の灯を綴ってはいました。それはまあ、ざっと百二十万年ぐらい前にできたというの。こいつをお持ちになったようにおもいました。
        </p>
      </div>
    </div>
  );
}
