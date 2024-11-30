import { Button } from "@/components/ui/button";
import { UserItem } from "@/components/userItem";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_layout/$roomId/description")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <div className="max-w-screen-lg mx-auto w-full px-6 py-12">
      <div className="w-full flex justify-between items-center gap-5 flex-wrap">
        <h2 className="text-2xl text-primary font-bold">ルーム名</h2>
        <Button>このルームに入る</Button>
      </div>
      <p className="py-8">
        Lorem ipsum dolor, sit amet consectetur adipisicing elit. Quo, optio
        recusandae ex autem sed, iure itaque voluptatibus dolorum accusamus sunt
        quam quod placeat accusantium minus! Molestias, cum? Aspernatur,
        excepturi nisi?
      </p>

      <div className="py-6 space-y-3">
        <h3 className="font-bold">作成者</h3>
        <UserItem />
      </div>
      <div className="py-6 space-y-3">
        <div className="flex gap-2 items-center">
          <h3 className="font-bold">参加者</h3>
          <div className="w-6 h-6 flex justify-center items-center leading-none text-sm rounded-full bg-primary text-primary-foreground">
            3
          </div>
        </div>
        <div className="flex flex-wrap gap-8">
          <UserItem />
          <UserItem />
          <UserItem />
        </div>
      </div>
    </div>
  );
}
