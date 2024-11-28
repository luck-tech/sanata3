import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "@/index.css";
import { QueryClientProvider } from "@tanstack/react-query";
import queryClient from "@/lib/queryClient.ts";
import { RouterProvider } from "@tanstack/react-router";
import { router } from "@/routes.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
    </QueryClientProvider>
  </StrictMode>
);
