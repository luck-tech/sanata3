import { useRef } from "react";
import { ChatBubble } from "amazon-chime-sdk-component-library-react";
import { StyledMessages } from "./Styled";

export default function Messages() {
  const scrollRef = useRef<HTMLDivElement>(null);
  const messages = [
    {
      isSelf: false,
      senderName: "aa",
      timestamp: new Date().toLocaleTimeString(),
      message: "メッセージ内容 メッセージ内容",
    },
    {
      isSelf: true,
      senderName: "aa",
      timestamp: new Date().toLocaleTimeString() + 10,
      message: "メッセージ内容 メッセージ内容",
    },
  ];

  // useEffect(() => {
  //   if (messages.length === 0) {
  //     return;
  //   }
  //   if (scrollRef.current) {
  //     scrollRef.current.scrollTop = scrollRef.current.scrollHeight;
  //   }
  // }, [messages.length]);

  const renderMessages = () => {
    return messages.map((message) => (
      <ChatBubble
        variant={message.isSelf ? "outgoing" : "incoming"}
        senderName={message.senderName}
        key={message.timestamp}
        showTail={message.isSelf ? true : false}
      >
        {message.message}
      </ChatBubble>
    ));
  };

  return <StyledMessages ref={scrollRef}>{renderMessages()}</StyledMessages>;
}
