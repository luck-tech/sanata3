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

// This is sample data.
const data = {
  user: {
    name: "shadcn",
    avatar: "/avatars/shadcn.jpg",
  },
  navMain: [
    {
      title: "React",
      url: "/react", // ルームのIdとして受け取りurlにする
    },
    {
      title: "Next.js",
      url: "/nextjs",
    },
    {
      title: "Hono",
      url: "/hono",
    },
    {
      title: "Web開発全般",
      url: "/web",
    },
    {
      title: "ITパスポート",
      url: "/it",
    },
  ],
};

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  return (
    <Sidebar {...props}>
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton size="lg" asChild>
              <span>
                <div className="flex aspect-square size-8 items-center justify-center rounded-lg bg-sidebar-primary text-sidebar-primary-foreground">
                  <GalleryVerticalEnd className="size-4" />
                </div>
                <div className="flex flex-col gap-0.5 leading-none">
                  <span className="font-semibold">サービス名</span>
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
              {data.navMain.map((item, idx) => (
                <SidebarLink to={item.url} key={idx}>
                  {item.title}
                </SidebarLink>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
      <SidebarFooter>
        <NavUser user={data.user} />
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
