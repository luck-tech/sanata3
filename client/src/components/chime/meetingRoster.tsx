import { useState, ChangeEvent } from "react";
import {
  Roster,
  RosterHeader,
  RosterGroup,
  useRosterState,
  RosterAttendeeType,
  RosterAttendee,
  RosterCell,
} from "amazon-chime-sdk-component-library-react";

const MeetingRoster = () => {
  const { roster } = useRosterState();
  const [filter, setFilter] = useState("");

  let attendees = Object.values(roster);

  if (filter) {
    attendees = attendees.filter((attendee: RosterAttendeeType) =>
      attendee?.name?.toLowerCase().includes(filter.trim().toLowerCase())
    );
  }

  const handleSearch = (e: ChangeEvent<HTMLInputElement>) => {
    setFilter(e.target.value);
  };

  const attendeeItems = attendees.map((attendee: RosterAttendeeType) => {
    const { chimeAttendeeId } = attendee || {};
    return <RosterAttendee attendeeId={chimeAttendeeId} />;
  });

  return (
    <Roster className="h-full">
      <RosterHeader
        searchValue={filter}
        onSearch={handleSearch}
        title="参加者"
        badge={attendees.length}
      />
      <RosterGroup>
        <RosterCell name="aa" muted={false} videoEnabled={false} />
        {attendeeItems}
      </RosterGroup>
    </Roster>
  );
};

export default MeetingRoster;
