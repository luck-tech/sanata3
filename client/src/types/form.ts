export interface TagsInputProps {
  selectedTags: string[];
  setSelectedTags: React.Dispatch<React.SetStateAction<string[]>>;
  isEditing?: boolean;
}
