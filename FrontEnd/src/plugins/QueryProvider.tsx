import { QueryClientProvider } from "@tanstack/react-query";
import type { ComponentType, ReactNode } from "react";
import { getQueryClient } from "../hooks/query-client";

export default function QueryProvider({ children }: { children: ReactNode }) {
  return (
    <QueryClientProvider client={getQueryClient()}>
      {children}
    </QueryClientProvider>
  );
}

export function withQueryProvider<P extends object>(
  Component: ComponentType<P>,
): ComponentType<P> {
  return function WithQueryProvider(props: P) {
    return (
      <QueryProvider>
        <Component {...props} />
      </QueryProvider>
    );
  };
}
