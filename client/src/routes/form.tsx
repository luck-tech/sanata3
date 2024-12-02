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
    <div className="flex flex-col items-center justify-center min-h-screen max-w-lg w-full mx-auto px-6 py-4">
      <h1 className="text-4xl mb-10 font-bold">アプリ名</h1>
      <p className="text-muted-foreground mb-10 text-center">
        入力された内容からおすすめのルームを表示します
      </p>
      <div className="w-full py-4">
        <p className="mb-3">学びたい技術や資格</p>
        <TagsInput
          selectedTags={selectedWantTags}
          setSelectedTags={setSelectedWantTags}
        />
      </div>
      <div className="w-full">
        <div className="w-full py-4">
          <p className="mb-3">使っている技術、取得済みの資格</p>
          <TagsInput
            selectedTags={selectedUsingTags}
            setSelectedTags={setSelectedUsingTags}
          />
        </div>
        <div className="flex justify-end">
          <Button>送信</Button>
        </div>
      </div>
    </div>
  );
}
