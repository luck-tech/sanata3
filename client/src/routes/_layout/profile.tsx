import { Button } from "@/components/ui/button";
import { Avatar, AvatarFallback, AvatarImage } from "@radix-ui/react-avatar";
import { createFileRoute } from "@tanstack/react-router";
import mockData from "@/mockData.json";
import { useState } from "react";
import TagsInput from "@/components/TagsInput";

export const Route = createFileRoute("/_layout/profile")({
  component: RouteComponent,
});

function RouteComponent() {
  const [selectedWantTags, setSelectedWantTags] = useState<string[]>(
    mockData.availableTags
  );
  const [selectedUsingTags, setSelectedUsingTags] = useState<string[]>(
    mockData.availableTags
  );
  const [isEditingWantTags, setIsEditingWantTags] = useState(false);
  const [isEditingUsingTags, setIsEditingUsingTags] = useState(false);

  return (
    <div className="flex flex-col flex-grow justify-center items-center gap-12 w-[579px] m-[0_auto]">
      <div className="flex gap-8 items-center w-full">
        <Avatar className="h-[100px] w-[100px] flex justify-center items-center rounded-lg">
          <AvatarImage
            src="../../../public/vite.svg"
            alt="shadcn"
            className="flex-grow"
          />
          <AvatarFallback className="rounded-lg">CN</AvatarFallback>
        </Avatar>
        <h1 className="text-[24px] font-bold">ユーザー名</h1>
      </div>
      <div className="flex flex-col gap-10">
        <div>
          <div className="flex justify-between items-center mb-2">
            <p>学びたい技術や資格</p>
            <Button
              variant={!isEditingWantTags ? "secondary" : "default"}
              onClick={() => setIsEditingWantTags((prev) => !prev)}
            >
              {isEditingWantTags ? "更新" : "編集"}
            </Button>
          </div>
          <TagsInput
            selectedTags={selectedWantTags}
            setSelectedTags={setSelectedWantTags}
            isEditing={isEditingWantTags}
          />
        </div>
        <div>
          <div className="flex justify-between items-center mb-2">
            <p>使ってる技術、取得済みの資格</p>
            <Button
              variant={!isEditingUsingTags ? "secondary" : "default"}
              onClick={() => setIsEditingUsingTags((prev) => !prev)}
            >
              {isEditingUsingTags ? "更新" : "編集"}
            </Button>
          </div>
          <TagsInput
            selectedTags={selectedUsingTags}
            setSelectedTags={setSelectedUsingTags}
            isEditing={isEditingUsingTags}
          />
        </div>
      </div>
    </div>
  );
}
