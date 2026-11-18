"use client";

import { navigationItems } from "../../data/navigation";
import { Icon } from "../ui";

type HeroSectionProps = {
  onOpenMenu: () => void;
};

export function HeroSection({ onOpenMenu }: HeroSectionProps) {
  return (
    <section id="home" className="hero-section">
      <div className="hero-bg">
        <img src="/hero.jpg" alt="DOSCOM Community" />
      </div>

      <div className="hero-overlay-dark" />
      <div className="hero-overlay-gradient" />

      <header className="hero-header">
        <a href="#home" className="hero-brand">
          <img src="/logo-doscom.svg" alt="DOSCOM Logo" />
          <span>DOSCOM</span>
        </a>

        <nav className="hero-nav">
          {navigationItems.map((item) => (
            <a key={item.href} href={item.href}>
              {item.label}
            </a>
          ))}
        </nav>

        <div className="hero-actions">
          <a
            href="mailto:doscom.go@gmail.com?subject=Gabung%20DOSCOM"
            className="hero-btn-start"
          >
            Gabung Sekarang
          </a>
          <button
            type="button"
            onClick={onOpenMenu}
            className="hero-menu-btn"
            aria-label="Open menu"
          >
            <Icon name="menu" size={32} />
          </button>
        </div>
      </header>

      <div className="hero-content">
        <div className="hero-title-col">
          <h1 className="hero-title">
            <span>Dinus Open</span>
            <span className="text-gradient">Source Community</span>
          </h1>
        </div>
        <div className="hero-desc-col">
          <a href="#about" className="btn btn-outline-white">
            <span>Jelajahi</span>
            <Icon name="arrow-down" size={16} />
          </a>
        </div>
      </div>
    </section>
  );
}
