import { useState, useRef } from "react";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { ScrollArea } from "@/components/ui/scroll-area";
import { TagsInputProps } from "@/types/form";
import { X } from "lucide-react";
import api from "@/api/axiosInstance";

const TagsInput = ({
  selectedTags,
  setSelectedTags,
  isEditing = true,
}: TagsInputProps) => {
  const [inputValue, setInputValue] = useState<string>("");
  const [availableTags, setAvailableTags] = useState<string[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [dropdownVisible, setDropdownVisible] = useState<boolean>(false);
  const [isComposing, setIsComposing] = useState<boolean>(false); // 日本語入力変換中かどうか
  const inputRef = useRef<HTMLInputElement>(null);

  const fetchTags = async (query: string) => {
    try {
      setLoading(true);
      const token = localStorage.getItem("code");

      const response = await api.get("/v1/skilltags", {
        params: {
          limit: 5,
          tag: query,
        },
        headers: {
          Authorization: token,
        },
      });
      setAvailableTags(response.data.tags || []);
      setDropdownVisible(true); // 結果が返ってきたらドロップダウンを表示
    } catch (error) {
      console.error("Error fetching tags:", error);
      setAvailableTags([]);
      setDropdownVisible(true); // エラー時もドロップダウンを表示（0件扱い）
    } finally {
      setLoading(false);
    }
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setInputValue(value);

    if (value.trim() !== "") {
      fetchTags(value);
    } else {
      setAvailableTags([]);
      setDropdownVisible(false); // 入力が空の場合はドロップダウンを非表示
    }
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (
      e.key === "Enter" &&
      !isComposing && // 日本語入力中ではない場合のみ実行
      dropdownVisible &&
      availableTags.length <= 0 &&
      inputValue.trim() !== ""
    ) {
      addTag(inputValue.trim());
    }
  };

  const handleCompositionStart = () => setIsComposing(true);
  const handleCompositionEnd = () => setIsComposing(false);

  const addTag = (tag: string) => {
    if (!selectedTags.includes(tag)) {
      setSelectedTags([...selectedTags, tag]);
    }
    setInputValue("");
    setAvailableTags([]);
    setDropdownVisible(false); // タグ追加時はドロップダウンを閉じる
    if (inputRef.current) {
      inputRef.current.focus();
    }
  };

  const removeTag = (tag: string) => {
    setSelectedTags(selectedTags.filter((t) => t !== tag));
    if (inputRef.current) {
      inputRef.current.focus();
    }
  };

  return (
    <div className="w-full max-w-screen-sm">
      <div
        className={`flex flex-wrap gap-2 ${
          isEditing ? "border border-input" : ""
        } rounded-md p-2 cursor-text`}
        onClick={() => {
          if (inputRef.current) {
            inputRef.current.focus();
          }
        }}
      >
        {selectedTags.map((tag) => (
          <Badge
            key={tag}
            variant="secondary"
            className={`flex items-center space-x-1 p-[6px_10px] ${
              isEditing
                ? "bg-muted text-black hover:bg-muted"
                : "bg-primary text-white hover:bg-primary"
            }`}
          >
            <span className="text-[14px]">{tag}</span>
            {isEditing && (
              <button
                type="button"
                onClick={(e) => {
                  e.stopPropagation();
                  removeTag(tag);
                }}
                className="ml-1 text-muted-foreground hover:text-foreground"
              >
                <X size={16} />
              </button>
            )}
          </Badge>
        ))}
        {isEditing && (
          <input
            ref={inputRef}
            className="flex-1 outline-none bg-transparent leading-[30px]"
            value={inputValue}
            onChange={handleInputChange}
            onKeyDown={handleKeyDown} // Enterキー押下を検知
            onCompositionStart={handleCompositionStart} // 日本語入力開始
            onCompositionEnd={handleCompositionEnd} // 日本語入力確定
            placeholder="ex: React, ITパスポート"
          />
        )}
      </div>
      {isEditing && dropdownVisible && (
        <div className="absolute z-50 w-full border border-gray-300 rounded-md mt-1 p-2 bg-white shadow-md">
          {loading ? (
            <div className="text-sm text-muted-foreground">読み込み中...</div>
          ) : availableTags.length > 0 ? (
            <ScrollArea className="h-40 w-full">
              <div className="flex flex-col space-y-1">
                {availableTags.map((tag) => (
                  <Button
                    key={tag}
                    variant="ghost"
                    className="justify-start"
                    onClick={() => addTag(tag)}
                    disabled={selectedTags.includes(tag)}
                  >
                    {tag}
                  </Button>
                ))}
              </div>
            </ScrollArea>
          ) : (
            <div className="p-2 text-sm text-muted-foreground">
              <span>
                一致するタグがありません。Enterキーで「{inputValue}
                」を追加します。
              </span>
            </div>
          )}
        </div>
      )}
    </div>
  );
};

export default TagsInput;
