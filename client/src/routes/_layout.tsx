import api from "@/api/axiosInstance";
import { AppSidebar } from "@/components/app-sidebar";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Separator } from "@/components/ui/separator";
import {
  SidebarInset,
  SidebarProvider,
  SidebarTrigger,
} from "@/components/ui/sidebar";
import { User } from "@/types/user";
import { useSuspenseQuery } from "@tanstack/react-query";
import {
  Navigate,
  Outlet,
  createFileRoute,
  useRouter,
} from "@tanstack/react-router";
import { Search } from "lucide-react";

export const Route = createFileRoute("/_layout")({
  component: LayoutComponent,
  errorComponent: () => {
    return <Navigate to="/" />;
  },
});

function LayoutComponent() {
  const token = localStorage.getItem("code");
  const userId = localStorage.getItem("userId");
  const router = useRouter();
  if (!token) router.navigate({ to: "/" });

  const { data: user } = useSuspenseQuery({
    queryKey: ["user"],
    queryFn: async (): Promise<User> => {
      const res = await api.get(`/v1/users/${userId}`, {
        headers: {
          Authorization: token,
        },
      });
      return res.data;
    },
  });

  const { data: rooms } = useSuspenseQuery({
    queryKey: ["rooms"],
    queryFn: async () => {
      const token = localStorage.getItem("code");
      const res = await api.get(`/v1/rooms`, {
        headers: {
          Authorization: token,
        },
      });
      return res.data;
    },
  });

  return (
    <SidebarProvider>
      <AppSidebar user={user} rooms={rooms.rooms} />
      <main className="w-full">
        <SidebarInset className="min-h-screen">
          <header className="flex h-16 shrink-0 items-center border-b gap-2 px-4 sticky top-0 bg-white z-20">
            <SidebarTrigger className="-ml-1" />
            <Separator orientation="vertical" className="mr-2 h-4" />
            <div className="flex justify-center items-center w-full gap-2 pr-7">
              <Input
                className="max-w-lg h-9 focus-visible:ring-1 focus-visible:ring-offset-0"
                placeholder="ルームを検索"
              />
              {/* Enterで検索が良さそう、面倒そうだったらこのままボタンで */}
              <Button variant={"outline"} size={"sm"}>
                <Search />
                検索
              </Button>
            </div>
          </header>
          <Outlet />
        </SidebarInset>
      </main>
    </SidebarProvider>
  );
}
