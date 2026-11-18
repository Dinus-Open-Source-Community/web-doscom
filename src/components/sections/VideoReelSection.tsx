"use client";

import type { LightboxItem } from "../../types/lightbox";
import { Icon } from "../ui";

type VideoReelSectionProps = {
  onOpenLightbox: (item: LightboxItem) => void;
};

export function VideoReelSection({ onOpenLightbox }: VideoReelSectionProps) {
  return (
    <section className="container video-reel-section">
      <button
        type="button"
        className="video-reel-banner"
        onClick={() =>
          onOpenLightbox({
            src: "/images/video-banner.jpg",
            caption: "DOSCOM Showreel",
          })
        }
      >
        <img src="/images/video-banner.jpg" alt="office work" />
        <div className="video-reel-overlay" />

        <div className="video-reel-play-btn">
          <Icon name="play" size={32} />
          <span>Play Reel</span>
        </div>
      </button>
    </section>
  );
}
