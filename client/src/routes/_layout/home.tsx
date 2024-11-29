import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_layout/home")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <div className="max-w-screen-lg mx-auto w-full p-8">
      <h2 className="text-2xl font-bold">おすすめ</h2>
    </div>
  );
}
