import type { ReactNode } from "react";

export type IconName =
  | "arrow-right-up"
  | "arrow-up"
  | "menu"
  | "close"
  | "arrow-right"
  | "arrow-down"
  | "like"
  | "play"
  | "add"
  | "check"
  | "linkedin"
  | "instagram"
  | "github"
  | "mail";

const iconPaths: Record<IconName, ReactNode> = {
  "arrow-right-up": <path d="M7 17 17 7M9 7h8v8" />,
  "arrow-up": <path d="M12 19V5m-6 6 6-6 6 6" />,
  menu: <path d="M4 7h16M4 12h16M4 17h16" />,
  close: (
    <>
      <circle cx="12" cy="12" r="9" />
      <path d="m15 9-6 6m0-6 6 6" />
    </>
  ),
  "arrow-right": <path d="M5 12h14m-6-6 6 6-6 6" />,
  "arrow-down": <path d="M12 5v14m-6-6 6 6 6-6" />,
  like: (
    <path d="M7 10v10H4V10h3Zm0 9h10.3a2 2 0 0 0 2-1.7l.8-5A2 2 0 0 0 18.3 10H14V6a2 2 0 0 0-2-2l-5 6v9Z" />
  ),
  play: (
    <>
      <circle cx="12" cy="12" r="9" />
      <path d="m10 8.5 5 3.5-5 3.5z" />
    </>
  ),
  add: (
    <>
      <circle cx="12" cy="12" r="9" />
      <path d="M12 8v8M8 12h8" />
    </>
  ),
  check: (
    <>
      <circle cx="12" cy="12" r="9" />
      <path d="m8 12 3 3 5-6" />
    </>
  ),
  linkedin: (
    <>
      <rect x="3" y="3" width="18" height="18" rx="2" />
      <path d="M8 10v7m0-10v.01M12 17v-4a3 3 0 0 1 6 0v4m-6-7v7" />
    </>
  ),
  instagram: (
    <>
      <rect x="3" y="3" width="18" height="18" rx="5" />
      <circle cx="12" cy="12" r="4" />
      <path d="M17.5 6.5h.01" />
    </>
  ),
  github: (
    <>
      <path d="M9 19c-4.3 1.4-4.3-2.2-6-2.5M15 21v-3.5a3 3 0 0 0-.8-2.3c2.7-.3 5.5-1.3 5.5-6A4.7 4.7 0 0 0 18.5 6 4.3 4.3 0 0 0 18.4 3S17.4 2.7 15 4.5a12 12 0 0 0-6 0C6.6 2.7 5.6 3 5.6 3A4.3 4.3 0 0 0 5.5 6a4.7 4.7 0 0 0-1.2 3.2c0 4.7 2.8 5.7 5.5 6A3 3 0 0 0 9 17.5V21" />
    </>
  ),
  mail: (
    <>
      <rect x="3" y="5" width="18" height="14" rx="2" />
      <path d="m3 7 9 6 9-6" />
    </>
  ),
};

type IconProps = {
  name: IconName;
  size?: number;
  className?: string;
  strokeWidth?: number;
};

export function Icon({
  name,
  size = 16,
  className,
  strokeWidth = 1.5,
}: IconProps) {
  return (
    <svg
      aria-hidden="true"
      className={className}
      fill="none"
      height={size}
      stroke="currentColor"
      strokeLinecap="round"
      strokeLinejoin="round"
      strokeWidth={strokeWidth}
      viewBox="0 0 24 24"
      width={size}
    >
      {iconPaths[name]}
    </svg>
  );
}
