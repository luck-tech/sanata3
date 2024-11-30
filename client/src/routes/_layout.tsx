import { AppSidebar } from "@/components/app-sidebar";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Separator } from "@/components/ui/separator";
import {
  SidebarInset,
  SidebarProvider,
  SidebarTrigger,
} from "@/components/ui/sidebar";
import { Outlet, createFileRoute } from "@tanstack/react-router";
import { Search } from "lucide-react";

export const Route = createFileRoute("/_layout")({
  component: LayoutComponent,
});

function LayoutComponent() {
  return (
    <SidebarProvider>
      <AppSidebar />
      <main className="w-full">
        <SidebarInset>
          <header className="flex h-16 shrink-0 items-center border-b gap-2 px-4 sticky top-0 bg-white">
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
