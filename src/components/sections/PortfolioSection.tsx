"use client";

import { useEffect, useRef, useState } from "react";
import type { LightboxItem } from "../../types/lightbox";
import { Icon } from "../ui";

type PortfolioSectionProps = {
  onScrollTo?: (id: string) => void;
  onOpenLightbox: (item: LightboxItem) => void;
};

interface DivisiCardData {
  id: string;
  number: string;
  title: string;
  kicker: string;
  description: string;
  image: string;
}

const DIVISI_ITEMS: DivisiCardData[] = [
  {
    id: "pemrograman",
    number: "01",
    title: "Pemrograman",
    kicker: "Pengembangan Software & Open Source",
    description:
      "Fokus pada rekayasa perangkat lunak modern, pengembangan web, arsitektur sistem, dan kontribusi aktif di ekosistem open source.",
    image:
      "https://images.unsplash.com/photo-1555066931-4365d14bab8c?q=80&w=1000&auto=format&fit=crop",
  },
  {
    id: "jaringan",
    number: "02",
    title: "Jaringan",
    kicker: "Infrastruktur Server & Sistem Cloud",
    description:
      "Mendalami administrasi sistem Linux, arsitektur jaringan terdistribusi, otomatisasi deployment, serta keandalan infrastruktur cloud.",
    image:
      "https://images.unsplash.com/photo-1544197150-b99a580bb7a8?q=80&w=1000&auto=format&fit=crop",
  },
  {
    id: "data",
    number: "03",
    title: "Data",
    kicker: "Sains Data & Kecerdasan Buatan",
    description:
      "Eksplorasi analitika data, pemodelan statistik, pembelajaran mesin (machine learning), dan penemuan wawasan cerdas dari pola data.",
    image:
      "https://images.unsplash.com/photo-1551288049-bebda4e38f71?q=80&w=1000&auto=format&fit=crop",
  },
  {
    id: "kreatif",
    number: "04",
    title: "Kreatif",
    kicker: "Identitas Visual & Media Interaktif",
    description:
      "Merancang identitas visual, estetika antarmuka (UI/UX), narasi multimedia, dan kreasi konten publikasi komunitas yang berkarakter.",
    image:
      "https://images.unsplash.com/photo-1618005182384-a83a8bd57fbe?q=80&w=1000&auto=format&fit=crop",
  },
];

export function PortfolioSection({ onOpenLightbox }: PortfolioSectionProps) {
  const [headerVisible, setHeaderVisible] = useState(false);
  const [visibleCards, setVisibleCards] = useState<Record<string, boolean>>({});
  const headerRef = useRef<HTMLDivElement | null>(null);
  const cardRefs = useRef<(HTMLDivElement | null)[]>([]);

  useEffect(() => {
    const cardObserver = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            const cardId = entry.target.getAttribute("data-card-id");
            if (cardId) {
              setVisibleCards((prev) => ({ ...prev, [cardId]: true }));
            }
          }
        });
      },
      {
        threshold: 0.25,
        rootMargin: "0px 0px -30% 0px",
      },
    );

    cardRefs.current.forEach((el) => {
      if (el) cardObserver.observe(el);
    });

    const headerObserver = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting) {
          setHeaderVisible(true);
        }
      },
      {
        threshold: 0.2,
        rootMargin: "0px 0px -40px 0px",
      },
    );

    if (headerRef.current) {
      headerObserver.observe(headerRef.current);
    }

    return () => {
      cardObserver.disconnect();
      headerObserver.disconnect();
    };
  }, []);

  return (
    <section id="portfolio" className="portfolio-section">
      <div className="container">
        {/* Section Header */}
        <div ref={headerRef} className="portfolio-header-stack">
          <div className="section-label">Divisi</div>
          <h2
            className={`portfolio-heading divisi-header-heading ${
              headerVisible ? "is-in-view" : ""
            }`}
          >
            Empat pilar eksplorasi teknologi dan inovasi terbuka
          </h2>
        </div>

        {/* Staggered Waterfall Grid */}
        <div className="divisi-waterfall-grid">
          {DIVISI_ITEMS.map((item, index) => {
            const isVisible = !!visibleCards[item.id];

            return (
              <div
                key={item.id}
                ref={(el) => {
                  cardRefs.current[index] = el;
                }}
                data-card-id={item.id}
                className={`divisi-card-wrapper divisi-stagger-${index} ${
                  isVisible ? "is-in-view" : ""
                }`}
              >
                <button
                  type="button"
                  className="divisi-card"
                  onClick={() =>
                    onOpenLightbox({
                      src: item.image,
                      caption: `${item.title} - ${item.kicker}`,
                    })
                  }
                >
                  {/* Image container with 4/5 aspect ratio and shutter reveal */}
                  <div className="divisi-img-frame">
                    <div className="divisi-shutter-inner">
                      <img
                        src={item.image}
                        alt={item.title}
                        className="divisi-img"
                        loading="lazy"
                      />
                    </div>
                  </div>

                  {/* Card Content with Text Blur Rise */}
                  <div className="divisi-content">
                    <div className="divisi-meta">
                      <span className="divisi-kicker">{item.kicker}</span>
                    </div>

                    <h3 className="divisi-card-title">{item.title}</h3>

                    <p className="divisi-card-desc">{item.description}</p>

                    <div className="divisi-card-footer">
                      <span className="divisi-link-text">
                        Eksplorasi Divisi
                      </span>
                      <Icon
                        name="arrow-right-up"
                        size={16}
                        className="divisi-link-icon"
                      />
                    </div>
                  </div>
                </button>
              </div>
            );
          })}
        </div>
      </div>
    </section>
  );
}
