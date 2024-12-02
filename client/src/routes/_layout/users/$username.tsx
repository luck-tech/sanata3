import { Button } from "@/components/ui/button";
import { createFileRoute } from "@tanstack/react-router";
import mockData from "@/mockData.json";
import { useState } from "react";
import TagsInput from "@/components/TagsInput";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";

export const Route = createFileRoute("/_layout/users/$username")({
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
    <div className="flex flex-col justify-center items-center gap-12 min-h-[calc(100vh-64px)] max-w-screen-sm w-full mx-auto px-6 py-8">
      <div className="flex gap-8 items-center w-full">
        <Avatar className="h-24 w-24 bg-secondary">
          <AvatarImage src="/vite.svg" alt="shadcn" className="" />
          <AvatarFallback className="">CN</AvatarFallback>
        </Avatar>
        <h1 className="text-2xl font-bold">ユーザー名</h1>
      </div>
      <div className="flex flex-col gap-10 pb-10">
        <div>
          <div className="flex justify-between items-center mb-2">
            <p>学びたい技術や資格</p>
            <Button
              variant={!isEditingWantTags ? "secondary" : "default"}
              size={"sm"}
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
              size={"sm"}
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
