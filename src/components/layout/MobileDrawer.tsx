"use client";

import { navigationItems } from "../../data/navigation";
import { Icon } from "../ui";

type MobileDrawerProps = {
  onClose: () => void;
};

export function MobileDrawer({ onClose }: MobileDrawerProps) {
  return (
    <div id="mobileMenu">
      <button
        type="button"
        id="mobileBackdrop"
        onClick={onClose}
        aria-label="Close menu"
      />
      <div id="mobileDrawer">
        <div>
          <div className="drawer-header">
            <div className="drawer-logo">
              <img src="/logo-doscom.svg" alt="DOSCOM Logo" />
              <span>DOSCOM</span>
            </div>
            <button
              type="button"
              onClick={onClose}
              className="drawer-close-btn"
              aria-label="Close menu"
            >
              <Icon name="close" size={20} />
            </button>
          </div>
          <nav className="drawer-nav">
            {navigationItems.map((item) => (
              <a
                key={item.href}
                href={item.href}
                className="nav-link"
                onClick={onClose}
              >
                <span>{item.label}</span>
                <Icon name="arrow-right" size={18} />
              </a>
            ))}
          </nav>
        </div>

        <div className="drawer-footer">
          <p>Siap bergabung dan berkarya bersama DOSCOM?</p>
          <a
            href="mailto:doscom.go@gmail.com?subject=Gabung%20DOSCOM"
            onClick={onClose}
            className="btn btn-pill-white"
          >
            Gabung Sekarang
          </a>
        </div>
      </div>
    </div>
  );
}
