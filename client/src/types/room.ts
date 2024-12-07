export type Room = {
  aimTags: [
    {
      id: number;
      name: string;
    },
  ];
  description: string;
  members: [
    {
      description: string;
      icon: string;
      id: string;
      name: string;
    },
  ];
  name: string;
  ownerId: string;
  roomId: string;
};
