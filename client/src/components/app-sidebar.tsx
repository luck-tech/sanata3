import * as React from "react";
import { GalleryVerticalEnd, Home, Plus } from "lucide-react";
import { Link, useMatchRoute } from "@tanstack/react-router";

import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarRail,
} from "@/components/ui/sidebar";
import { NavUser } from "./nav-user";
import { User } from "@/types/user";
import { Room } from "@/types/room";

export function AppSidebar({ user, rooms }: { user: User; rooms: Room[] }) {
  return (
    <Sidebar>
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton size="lg" asChild>
              <span>
                <div className="flex aspect-square size-8 items-center justify-center rounded-lg bg-sidebar-primary text-sidebar-primary-foreground">
                  <GalleryVerticalEnd className="size-4" />
                </div>
                <div className="flex flex-col gap-0.5 leading-none">
                  <span className="font-semibold">Haru World</span>
                </div>
              </span>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
        <SidebarMenu>
          <SidebarLink to="/home">
            <Home /> <span>ホーム</span>
          </SidebarLink>
          <SidebarLink to="/new">
            <Plus /> <span>ルーム作成</span>
          </SidebarLink>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>参加中のルーム</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {rooms === null ? (
                <div>入ってるルームはありません</div>
              ) : (
                rooms.map((room) => (
                  <SidebarLink to={room.roomId} key={room.roomId}>
                    {room.name}
                  </SidebarLink>
                ))
              )}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
      <SidebarFooter>
        <NavUser user={user} />
      </SidebarFooter>
      <SidebarRail />
    </Sidebar>
  );
}

const SidebarLink = ({
  to,
  children,
}: {
  to: string;
  children: React.ReactNode;
}) => {
  const matchRoute = useMatchRoute();
  const params = matchRoute({ to: to });

  return (
    <>
      <SidebarMenuItem>
        <SidebarMenuButton asChild isActive={!!params}>
          <Link to={to}>{children}</Link>
        </SidebarMenuButton>
      </SidebarMenuItem>
    </>
  );
};
