import { useState, useRef } from "react";
import { Badge } from "@/components/ui/badge";
import {
  Popover,
  PopoverTrigger,
  PopoverContent,
} from "@/components/ui/popover";
import { Button } from "@/components/ui/button";
import { ScrollArea } from "@/components/ui/scroll-area";
import { TagsInputProps } from "@/types/form";
import mockData from "@/mockData.json";

const TagsInput = ({ selectedTags, setSelectedTags }: TagsInputProps) => {
  const [isOpen, setIsOpen] = useState<boolean>(false);
  const [inputValue, setInputValue] = useState<string>("");
  const inputRef = useRef<HTMLInputElement>(null);
  const [availableTags, setAvailableTags] = useState<string[]>(
    mockData.availableTags
  );

  const addTag = (tag: string) => {
    if (!selectedTags.includes(tag)) {
      setSelectedTags([...selectedTags, tag]);
    }
    setInputValue("");
    // 入力フィールドにフォーカスを戻す
    if (inputRef.current) {
      inputRef.current.focus();
    }
  };

  const removeTag = (tag: string) => {
    setSelectedTags(selectedTags.filter((t) => t !== tag));
    // 入力フィールドにフォーカスを戻す
    if (inputRef.current) {
      inputRef.current.focus();
    }
  };

  // 入力値に基づいてタグをフィルタリング
  const filteredTags = availableTags.filter((tag) =>
    tag.toLowerCase().includes(inputValue.toLowerCase())
  );

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      e.preventDefault();
      if (inputValue.trim() !== "") {
        // タグを `selectedTags` に追加
        addTag(inputValue.trim());

        // タグを `availableTags` に追加
        if (!availableTags.includes(inputValue.trim())) {
          setAvailableTags((prev) => [...prev, inputValue.trim()]);
        }

        // 入力値をクリア
        setInputValue("");
      }
    }
  };

  return (
    <div className="w-[579px]">
      <Popover open={isOpen} onOpenChange={setIsOpen}>
        <PopoverTrigger asChild>
          <div
            className="flex flex-wrap gap-2 border border-input rounded-md p-2 cursor-text"
            onClick={() => {
              setIsOpen(true);
              if (inputRef.current) {
                inputRef.current.focus();
              }
            }}
          >
            {selectedTags.map((tag) => (
              <Badge
                key={tag}
                variant="secondary"
                className="flex items-center space-x-1 p-[6px_10px]"
              >
                <span className="text-[14px]">{tag}</span>
                <button
                  type="button"
                  onClick={(e) => {
                    e.stopPropagation();
                    removeTag(tag);
                  }}
                  className="ml-1 text-muted-foreground hover:text-foreground "
                >
                  &times;
                </button>
              </Badge>
            ))}
            <input
              ref={inputRef}
              className="flex-1 outline-none bg-transparent"
              value={inputValue}
              onChange={(e) => setInputValue(e.target.value)}
              onKeyDown={handleKeyDown}
              placeholder="ex: React, ITパスポート"
            />
          </div>
        </PopoverTrigger>
        <PopoverContent
          className="w-[579px] p-2"
          // フォーカスが移動しないように設定
          onOpenAutoFocus={(event) => event.preventDefault()}
          onCloseAutoFocus={(event) => event.preventDefault()}
        >
          <ScrollArea className="h-40">
            {filteredTags.length > 0 ? (
              <div className="flex flex-col space-y-1">
                {filteredTags.map((tag) => (
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
            ) : (
              <div className="p-2 text-sm text-muted-foreground">
                <span>
                  一致するタグがありません。Enterキーで「{inputValue}
                  」を追加します。
                </span>
              </div>
            )}
          </ScrollArea>
        </PopoverContent>
      </Popover>
    </div>
  );
};

export default TagsInput;
