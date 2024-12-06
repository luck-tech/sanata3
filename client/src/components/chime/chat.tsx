import ChatInput from "./chatInput";
import Messages from "./messages";
import { StyledChat, StyledTitle } from "./Styled";

export default function Chat() {
  return (
    <StyledChat className="bg-secondary border-t border-l border-border">
      <StyledTitle className="border-b border-muted-foreground">
        <div className="text-sm">チャット</div>
      </StyledTitle>
      <Messages />
      <ChatInput />
    </StyledChat>
  );
}
