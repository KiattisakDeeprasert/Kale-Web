export default function AuthLayout({
    children,
  }: Readonly<{ children: React.ReactNode }>) {
    return <>{children}</>; // ไม่มี Navbar
  }
  