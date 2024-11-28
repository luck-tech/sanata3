import {
  createRootRoute,
  createRoute,
  createRouter,
} from "@tanstack/react-router";
import Layout from "./components/layout";

// layout
const rootRoute = createRootRoute({
  component: Layout,
});

// 下記のようにルートを追加していく
const homeRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/",
  component: () => <p>home</p>,
});

const aboutRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/about",
  component: () => <p>about</p>,
});

const routeTree = rootRoute.addChildren([homeRoute, aboutRoute]);

export const router = createRouter({ routeTree });

declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}
