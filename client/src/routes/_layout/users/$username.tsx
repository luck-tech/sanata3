import { Button } from "@/components/ui/button";
import { createFileRoute } from "@tanstack/react-router";
import { useEffect, useState } from "react";
import TagsInput from "@/components/TagsInput";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import api from "@/api/axiosInstance";

export const Route = createFileRoute("/_layout/users/$username")({
  component: RouteComponent,
});

function RouteComponent() {
  const [selectedWantTags, setSelectedWantTags] = useState<string[]>([]);
  const [selectedUsingTags, setSelectedUsingTags] = useState<string[]>([]);
  const [isEditingWantTags, setIsEditingWantTags] = useState(false);
  const [isEditingUsingTags, setIsEditingUsingTags] = useState(false);
  const [userName, setUserName] = useState("");
  const [avatarIcon, setAvatarIcon] = useState("");

  useEffect(() => {
    const userId = localStorage.getItem("userId");
    const token = localStorage.getItem("code");

    const fetchUserData = async () => {
      try {
        const response = await api.get(`/v1/users/${userId}`, {
          headers: {
            Authorization: token,
          },
        });
        const data = response.data;

        setUserName(data.name);
        setAvatarIcon(data.icon);
        setSelectedUsingTags(
          data.usedSkills.map((skill: { name: string }) => skill.name)
        );
        setSelectedWantTags(
          data.wantLeanSkills.map((skill: { name: string }) => skill.name)
        );
      } catch (error) {
        console.error("Failed to fetch user data:", error);
      }
    };

    fetchUserData();
  }, []);

  return (
    <div className="flex flex-col justify-center items-center gap-12 min-h-[calc(100vh-64px)] max-w-screen-sm w-full mx-auto px-6 py-8">
      <div className="flex gap-8 items-center w-full">
        <Avatar className="h-24 w-24 bg-secondary">
          <AvatarImage src={avatarIcon} alt="User Avatar" className="" />
          <AvatarFallback className="">CN</AvatarFallback>
        </Avatar>
        <h1 className="text-2xl font-bold">{userName}</h1>
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
