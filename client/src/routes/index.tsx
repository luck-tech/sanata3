import { Button } from "@/components/ui/button";
import { createFileRoute } from "@tanstack/react-router";
import { Github } from "lucide-react";

export const Route = createFileRoute("/")({
  component: Home,
});

function Home() {
  const handleLogin = () => {
    const clientId = import.meta.env.VITE_GITHUB_CLIENT_ID;
    const authorizeUrl = `https://github.com/login/oauth/authorize?client_id=${clientId}`;

    window.location.href = authorizeUrl;
  };

  return (
    <div className="container mx-auto flex flex-col gap-10 justify-center items-center min-h-screen">
      <h1 className="text-5xl font-bold">サービス名</h1>
      <Button
        className="text-lg [&_svg]:size-6"
        size={"lg"}
        onClick={handleLogin}
      >
        <Github />
        GitHubでログイン
      </Button>
    </div>
  );
}

export default Home;
