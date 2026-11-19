"use client";

import type { LightboxItem } from "../../types/lightbox";
import { Icon } from "../ui";

type PortfolioSectionProps = {
  onScrollTo: (id: string) => void;
  onOpenLightbox: (item: LightboxItem) => void;
};

export function PortfolioSection({
  onScrollTo,
  onOpenLightbox,
}: PortfolioSectionProps) {
  return (
    <section id="portfolio" className="portfolio-section">
      <div className="container">
        <div className="portfolio-header-grid">
          <div className="portfolio-label-col">
            <div className="section-label">Divisi</div>
          </div>
          <div className="portfolio-title-col">
            <h2 className="portfolio-heading">
              Explore our latest branding and design projects
            </h2>
            <button
              type="button"
              onClick={() => onScrollTo("services")}
              className="btn btn-primary"
            >
              <span>View All</span>
              <Icon name="arrow-right" size={16} />
            </button>
          </div>
        </div>

        <div className="portfolio-grid">
          <button
            type="button"
            className="bento-item bento-big"
            onClick={() =>
              onOpenLightbox({
                src: "/images/portfolio-pixelcraft.jpg",
                caption: "Pixelcraft Studio - Abstract 3D Architecture",
              })
            }
          >
            <div className="bento-img-wrap">
              <img
                src="/images/portfolio-pixelcraft.jpg"
                alt="abstract golden 3d shape"
              />
            </div>
            <div className="bento-info">
              <span>Pixelcraft Studio</span>
              <span className="bento-tag">3D Visual &amp; Architecture</span>
            </div>
            <div className="bento-hover-badge">
              <div className="badge-view">View</div>
            </div>
          </button>

          <button
            type="button"
            className="bento-item bento-medium"
            onClick={() =>
              onOpenLightbox({
                src: "/images/portfolio-novaworks.jpg",
                caption: "Novaworks - Dark Textured Organic Identity",
              })
            }
          >
            <div className="bento-img-wrap">
              <img
                src="/images/portfolio-novaworks.jpg"
                alt="dark textured organic shape"
              />
              <div className="bento-hover-badge">
                <div className="badge-view">View</div>
              </div>
            </div>
            <div className="bento-info">
              <span>Novaworks</span>
              <span className="bento-tag">Brand Narrative</span>
            </div>
          </button>

          <button
            type="button"
            className="bento-item bento-medium"
            onClick={() =>
              onOpenLightbox({
                src: "/images/portfolio-visionary.jpg",
                caption: "Visionary Threads - Minimalist Product Packaging",
              })
            }
          >
            <div className="bento-img-wrap bento-item-bg-neutral">
              <img src="/images/portfolio-visionary.jpg" alt="white bottle" />
              <div className="bento-hover-badge">
                <div className="badge-view">View</div>
              </div>
            </div>
            <div className="bento-info">
              <span>Visionary Threads</span>
              <span className="bento-tag">Product Experience</span>
            </div>
          </button>
        </div>
      </div>
    </section>
  );
}
