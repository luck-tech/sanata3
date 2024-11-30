import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_layout/new')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/_layout/new"!</div>
}
