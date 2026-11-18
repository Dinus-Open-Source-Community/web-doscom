import "../styles/globals.css";
import type { Metadata } from "next";
import { Plus_Jakarta_Sans } from "next/font/google";

const jakartaSans = Plus_Jakarta_Sans({
  variable: "--font-jakarta-sans",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Doscom",
  description: "Dinus Open Source Community",
  icons: {
    icon: "/logo-doscom.svg",
  },
};

export default function RootLayout({ children }: LayoutProps<"/">) {
  return (
    <html lang="en" className={jakartaSans.variable}>
      <body>{children}</body>
    </html>
  );
}
