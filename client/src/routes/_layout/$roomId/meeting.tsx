import Chat from "@/components/chime/chat";
import MeetingRoster from "@/components/chime/meetingRoster";
import { createFileRoute, useNavigate } from "@tanstack/react-router";
import {
  AudioInputControl,
  AudioOutputControl,
  ContentShareControl,
  ControlBar,
  ControlBarButton,
  Flex,
  Modal,
  ModalBody,
  ModalButton,
  ModalButtonGroup,
  ModalHeader,
  Phone,
  UserActivityProvider,
  VideoInputControl,
  VideoTileGrid,
} from "amazon-chime-sdk-component-library-react";
import { useState } from "react";

export const Route = createFileRoute("/_layout/$roomId/meeting")({
  component: RouteComponent,
});

function RouteComponent() {
  const [showModal, setShowModal] = useState(false);
  const toggleModal = (): void => setShowModal(!showModal);
  const navigate = useNavigate();

  const leaveMeeting = async (): Promise<void> => {
    navigate({ to: "/home" });
  };

  return (
    <UserActivityProvider>
      <div className="min-h-[calc(100vh-64px)] w-full flex">
        <VideoTileGrid className="p-5 flex-1" />
        <ControlBar
          layout="undocked-horizontal"
          className="absolute bottom-10 mx-auto right-0 left-0 w-fit"
          showLabels
        >
          <AudioInputControl />
          <VideoInputControl />
          <ContentShareControl />
          <AudioOutputControl />
          <ControlBarButton
            icon={<Phone />}
            onClick={toggleModal}
            label={"Leave"}
          />
          {showModal && (
            <Modal size="md" onClose={toggleModal} rootId="modal-root">
              <ModalHeader title="End Meeting" />
              <ModalBody>
                <p className="py-4">
                  Leave meeting or you can end the meeting for all. The meeting
                  cannot be used once it ends.
                </p>
              </ModalBody>
              <ModalButtonGroup
                primaryButtons={[
                  <ModalButton
                    key="leave-meeting"
                    onClick={leaveMeeting}
                    variant="primary"
                    label="Leave Meeting"
                    closesModal
                  />,
                  <ModalButton
                    key="cancel-meeting-ending"
                    variant="secondary"
                    label="Cancel"
                    closesModal
                  />,
                ]}
              />
            </Modal>
          )}
        </ControlBar>
        <Flex layout="stack">
          <MeetingRoster />
          <Chat />
        </Flex>
      </div>
    </UserActivityProvider>
  );
}
