"use client";

import { usePathname } from "next/navigation";

export default function ClientWrapper({ children }: { children: React.ReactNode }) {
  const pathname = usePathname();

  const noLayoutPages = ["/auth/login", "/auth/register"];

  if (noLayoutPages.includes(pathname)) {
    return <>{children}</>; 
  }

  return <>{children}</>;
}
