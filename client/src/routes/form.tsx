import api from "@/api/axiosInstance";
import TagsInput from "@/components/TagsInput";
import { Button } from "@/components/ui/button";
import { createFileRoute, useRouter } from "@tanstack/react-router";
import { useEffect, useState } from "react";

export const Route = createFileRoute("/form")({
  component: RouteComponent,
});

function RouteComponent() {
  const [selectedWantTags, setSelectedWantTags] = useState<string[]>([]);
  const [selectedUsingTags, setSelectedUsingTags] = useState<string[]>([]);
  const userId = localStorage.getItem("userId");
  const token = localStorage.getItem("code");
  const router = useRouter();
  if (!token) router.navigate({ to: "/" });

  useEffect(() => {
    const fetchUserSkills = async () => {
      try {
        const response = await api.get(`/v1/users/${userId}`, {
          headers: {
            Authorization: token,
          },
        });

        const data = response.data;
        setSelectedUsingTags(
          data.usedSkills.map((skill: { name: string }) => skill.name)
        );
      } catch (error) {
        console.error("Error fetching user skills: ", error);
      }
    };

    fetchUserSkills();
  }, [token]);

  const handleSubmit = async () => {
    try {
      const response = await api.put(
        `/v1/users/${userId}`,
        {
          usedSkills: selectedUsingTags,
          userID: userId,
          wantLearnSkills: selectedWantTags,
        },
        {
          headers: {
            Authorization: token,
          },
        }
      );
      console.log("Response data:", response.data);
      router.navigate({ to: "/home" }); // 成功時に/homeにリダイレクト
    } catch (error) {
      console.error("Error sending data:", error);
    }
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-screen max-w-lg w-full mx-auto px-6 py-4">
      <h1 className="text-4xl mb-10 font-bold">Haru World</h1>
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
          <Button onClick={handleSubmit}>送信</Button>
        </div>
      </div>
    </div>
  );
}
