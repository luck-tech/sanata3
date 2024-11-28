import { Outlet } from "@tanstack/react-router";

//一旦
export default function Layout() {
  return (
    <div className="flex flex-col min-h-screen">
      <header className="bg-blue-600 text-white p-4">
        <nav>
          <a href="/" className="mr-4">
            Home
          </a>
          <a href="/about">About</a>
        </nav>
      </header>
      <main className="flex-1 p-4">
        <Outlet /> {/* 子ルートがここにレンダリングされる */}
      </main>
    </div>
  );
}
