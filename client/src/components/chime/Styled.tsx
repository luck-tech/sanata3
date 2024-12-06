import styled from "styled-components";

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const StyledChat = styled.aside<any>`
  display: grid;
  grid-template-rows: auto 1fr auto;
  grid-template-areas:
    "chat-header"
    "messages"
    "chat-input";
  width: 100%;
  height: 100%;
  padding-bottom: 1rem;
  overflow-y: auto;
`;

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const StyledTitle = styled.div<any>`
  grid-area: chat-header;
  position: relative;
  display: flex;
  align-items: center;
  padding: 1rem 1rem;
  margin-bottom: 0.5rem;
`;

export const StyledChatInputContainer = styled.div`
  grid-area: chat-input;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-top: 0.75rem;

  .ch-input-wrapper {
    width: 90%;

    .ch-input {
      width: 100%;
    }
  }
`;

export const StyledMessages = styled.div`
  grid-area: messages;
  display: flex;
  flex-direction: column;
  height: 100%;
  overflow-y: auto;
  row-gap: 0.5rem;
  padding: 0 0.75rem;
`;
