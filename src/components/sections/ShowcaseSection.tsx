"use client";

import type { LightboxItem } from "../../types/lightbox";
import { Icon } from "../ui";

type ShowcaseSectionProps = {
  onOpenLightbox: (item: LightboxItem) => void;
};

export function ShowcaseSection({ onOpenLightbox }: ShowcaseSectionProps) {
  return (
    <section className="dark-showcase-section">
      <div className="container">
        <h2 className="showcase-heading">
          Modern design and development for real-world solutions
        </h2>
        <div className="showcase-grid">
          <button
            type="button"
            className="showcase-card showcase-accent"
            onClick={() =>
              onOpenLightbox({
                src: "/images/showcase-3d.jpg",
                caption: "3D Form & Dynamics",
              })
            }
          >
            <img src="/images/showcase-3d.jpg" alt="3d rendering" />
          </button>

          <button
            type="button"
            className="showcase-card showcase-overlay-card"
            onClick={() =>
              onOpenLightbox({
                src: "/images/showcase-packaging.jpg",
                caption: "Tactile Packaging Design",
              })
            }
          >
            <img src="/images/showcase-packaging.jpg" alt="packaging design" />
            <div className="showcase-overlay">
              <h3 className="showcase-overlay-title">Packaging design</h3>
              <p className="showcase-overlay-desc">
                Tactile aesthetics that elevate product perception.
              </p>
              <div className="showcase-plus-btn">
                <Icon name="add" size={20} />
              </div>
            </div>
          </button>

          <button
            type="button"
            className="showcase-card showcase-dark"
            onClick={() =>
              onOpenLightbox({
                src: "/images/showcase-workspace.jpg",
                caption: "Creative Workspace Culture",
              })
            }
          >
            <img src="/images/showcase-workspace.jpg" alt="person working" />
          </button>
        </div>
      </div>
    </section>
  );
}
