import TagsInput from "@/components/TagsInput";
import { Button } from "@/components/ui/button";
import { createFileRoute } from "@tanstack/react-router";
import { useState } from "react";

export const Route = createFileRoute("/form")({
  component: RouteComponent,
});

function RouteComponent() {
  const [selectedWantTags, setSelectedWantTags] = useState<string[]>([]);
  const [selectedUsingTags, setSelectedUsingTags] = useState<string[]>([]);

  return (
    <div className="flex flex-col items-center justify-center min-h-screen">
      <h1 className="text-[36px] mb-[51px] font-bold">アプリ名</h1>
      <p className="text-[#7B7B7B] mb-[83px]">
        入力された内容からおすすめのルームを表示します
      </p>
      <div>
        <p className="mb-[16px]">学びたい技術や資格</p>
        <TagsInput
          selectedTags={selectedWantTags}
          setSelectedTags={setSelectedWantTags}
        />
      </div>
      <div>
        <div>
          <p className="m-[51px_0_16px_0]">使っている技術、取得済みの資格</p>
          <TagsInput
            selectedTags={selectedUsingTags}
            setSelectedTags={setSelectedUsingTags}
          />
        </div>
        <div className="mt-[39px] flex justify-end">
          <Button className="w-[60px]">送信</Button>
        </div>
      </div>
    </div>
  );
}
