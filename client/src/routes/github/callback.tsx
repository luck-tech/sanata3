import api from "@/api/axiosInstance";
import { createFileRoute } from "@tanstack/react-router";
import { useRouter } from "@tanstack/react-router";
import { useEffect, useState } from "react";

export const Route = createFileRoute("/github/callback")({
  component: GitHubCallback,
});

function GitHubCallback() {
  const router = useRouter();
  const searchParams = new URLSearchParams(router.state.location.search);
  const code = searchParams.get("code");

  const [loginStatus, setLoginStatus] = useState(true);

  useEffect(() => {
    const fetchGitHubLogin = async () => {
      if (!code) router.navigate({ to: "/" });
      try {
        const response = await api.post("/login/github", { code });
        localStorage.setItem("access code", response.data.code);

        setLoginStatus(false);
        localStorage.setItem("code", response.data.code);
        localStorage.setItem("userId", response.data.id);
        router.navigate({ to: "/form" });
      } catch (error) {
        console.error("Login Error: ", error);
      }
    };

    fetchGitHubLogin();
  }, [code, router]);

  if (loginStatus) {
    return <p>isLoading...</p>;
  }
  return;
}

export default GitHubCallback;
