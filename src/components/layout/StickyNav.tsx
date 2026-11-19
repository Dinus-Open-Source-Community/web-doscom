"use client";

import { navigationItems } from "../../data/navigation";
import { Icon } from "../ui";

type StickyNavProps = {
  visible: boolean;
  onScrollTo: (id: string) => void;
  onOpenMenu: () => void;
};

export function StickyNav({ visible, onScrollTo, onOpenMenu }: StickyNavProps) {
  return (
    <nav
      id="stickyNav"
      aria-label="Quick navigation"
      className={visible ? "nav-visible" : undefined}
    >
      <div className="container sticky-nav-inner">
        <a href="#home" className="hero-brand" aria-label="DOSCOM Home">
          <img
            src="/logo-with-text.svg"
            alt="DOSCOM"
            className="sticky-logo-img"
          />
        </a>

        <div className="sticky-nav-links">
          {navigationItems.map((item) => (
            <a key={item.href} href={item.href} className="sticky-nav-link">
              {item.label}
            </a>
          ))}
        </div>

        <div className="sticky-nav-actions">
          <a
            href="mailto:doscom.go@gmail.com?subject=Gabung%20DOSCOM"
            className="btn btn-primary btn-talk"
          >
            <span>Gabung Sekarang</span>
            <Icon name="arrow-right-up" size={12} />
          </a>
          <button
            type="button"
            onClick={() => onScrollTo("home")}
            title="Back to top"
            aria-label="Back to top"
            className="icon-circle-btn"
          >
            <Icon name="arrow-up" size={16} />
          </button>
          <button
            type="button"
            onClick={onOpenMenu}
            className="icon-circle-btn icon-circle-btn-mobile sticky-mobile-trigger"
            aria-label="Open mobile menu"
          >
            <Icon name="menu" size={20} />
          </button>
        </div>
      </div>
    </nav>
  );
}
