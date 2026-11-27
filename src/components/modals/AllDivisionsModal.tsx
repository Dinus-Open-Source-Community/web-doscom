"use client";

import { useEffect, useRef, useState } from "react";
import { Icon } from "../ui";

export interface DivisionItem {
  id: string;
  category: string;
  title: string;
  kicker: string;
  description: string;
  image: string;
}

export const DIVISIONS_TOP_4: DivisionItem[] = [
  {
    id: "bph",
    category: "Kepemimpinan & Tata Kelola",
    title: "Badan Pengurus Harian",
    kicker: "Pimpinan & Tata Kelola Organisasi",
    description:
      "Mengoordinasikan arah strategis, pengelolaan administrasi, keuangan, dan kebijakan operasional Dinus Open Source Community secara menyeluruh.",
    image:
      "https://images.unsplash.com/photo-1522071820081-009f0129c71c?q=80&w=1000&auto=format&fit=crop",
  },
  {
    id: "tech-lead",
    category: "Kepemimpinan & Tata Kelola",
    title: "Tech Lead",
    kicker: "Arsitektur Teknis & Standarisasi",
    description:
      "Menentukan arah arsitektur perangkat lunak, standar kode open source, riset teknologi mutakhir, serta pembinaan kualitas teknis tim pengembang.",
    image:
      "https://images.unsplash.com/photo-1531403009284-440f080d1e12?q=80&w=1000&auto=format&fit=crop",
  },
  {
    id: "project-management",
    category: "Kepemimpinan & Tata Kelola",
    title: "Project Manager",
    kicker: "Manajemen Proyek & Timeline",
    description:
      "Mengelola siklus pengembangan produk komunitas, alokasi tim kerja, koordinasi milestone, dan ketercapaian target rilis program kerja.",
    image:
      "https://images.unsplash.com/photo-1552664730-d307ca884978?q=80&w=1000&auto=format&fit=crop",
  },
  {
    id: "humas",
    category: "Relasi & Komunikasi",
    title: "Hubungan Masyarakat",
    kicker: "Komunikasi Eksternal & Kemitraan",
    description:
      "Membangun jejaring kolaborasi dengan komunitas teknologi luar, media partner, instansi akademik, dan perwakilan resmi DOSCOM.",
    image:
      "https://images.unsplash.com/photo-1511632765486-a01980e01a18?q=80&w=1000&auto=format&fit=crop",
  },
];

export const DIVISIONS_BOTTOM_4: DivisionItem[] = [
  {
    id: "pemrograman",
    category: "Pilar Eksplorasi Teknis",
    title: "Pemrograman",
    kicker: "Pengembangan Software & Open Source",
    description:
      "Fokus pada rekayasa perangkat lunak modern, pengembangan web, arsitektur sistem backend, dan kontribusi aktif di ekosistem open source.",
    image:
      "https://images.unsplash.com/photo-1555066931-4365d14bab8c?q=80&w=1000&auto=format&fit=crop",
  },
  {
    id: "jaringan",
    category: "Pilar Eksplorasi Teknis",
    title: "Jaringan",
    kicker: "Infrastruktur Server & Sistem Cloud",
    description:
      "Mendalami administrasi sistem Linux, arsitektur jaringan terdistribusi, otomatisasi deployment server, serta keandalan infrastruktur cloud.",
    image:
      "https://images.unsplash.com/photo-1544197150-b99a580bb7a8?q=80&w=1000&auto=format&fit=crop",
  },
  {
    id: "data",
    category: "Pilar Eksplorasi Teknis",
    title: "Data",
    kicker: "Sains Data & Kecerdasan Buatan",
    description:
      "Eksplorasi analitika data, pemodelan statistik, pembelajaran mesin (machine learning), dan penemuan wawasan cerdas dari pola data.",
    image:
      "https://images.unsplash.com/photo-1551288049-bebda4e38f71?q=80&w=1000&auto=format&fit=crop",
  },
  {
    id: "kreatif",
    category: "Pilar Eksplorasi Teknis",
    title: "Kreatif",
    kicker: "Identitas Visual & Media Interaktif",
    description:
      "Merancang identitas visual, estetika antarmuka (UI/UX), narasi multimedia, dan kreasi konten publikasi komunitas yang berkarakter.",
    image:
      "https://images.unsplash.com/photo-1618005182384-a83a8bd57fbe?q=80&w=1000&auto=format&fit=crop",
  },
];

type AllDivisionsModalProps = {
  isOpen: boolean;
  origin: { x: number; y: number } | null;
  onClose: () => void;
};

export function AllDivisionsModal({
  isOpen,
  origin,
  onClose,
}: AllDivisionsModalProps) {
  const [mounted, setMounted] = useState(false);
  const [circleActive, setCircleActive] = useState(false);
  const [contentActive, setContentActive] = useState(false);
  const [expandedDone, setExpandedDone] = useState(false);
  const isClosingRef = useRef(false);

  useEffect(() => {
    if (isOpen) {
      isClosingRef.current = false;
      setMounted(true);
      window.dispatchEvent(new CustomEvent("lenis:stop"));

      const tCircle = setTimeout(() => {
        setCircleActive(true);
      }, 30);
      const tContent = setTimeout(() => {
        setContentActive(true);
      }, 450);
      const tDone = setTimeout(() => {
        setExpandedDone(true);
      }, 1450);

      return () => {
        clearTimeout(tCircle);
        clearTimeout(tContent);
        clearTimeout(tDone);
      };
    }
  }, [isOpen]);

  const handleClose = () => {
    if (isClosingRef.current) return;
    isClosingRef.current = true;

    setExpandedDone(false);
    setContentActive(false);

    const tCircle = setTimeout(() => {
      setCircleActive(false);
    }, 280);

    const tUnmount = setTimeout(() => {
      setMounted(false);
      window.dispatchEvent(new CustomEvent("lenis:start"));
      onClose();
    }, 1500);

    return () => {
      clearTimeout(tCircle);
      clearTimeout(tUnmount);
    };
  };

  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if (e.key === "Escape" && isOpen && mounted) {
        handleClose();
      }
    };
    window.addEventListener("keydown", handleKeyDown);
    return () => window.removeEventListener("keydown", handleKeyDown);
  }, [isOpen, mounted]);

  if (!mounted) return null;

  const posX =
    origin?.x ?? (typeof window !== "undefined" ? window.innerWidth / 2 : 0);
  const posY =
    origin?.y ?? (typeof window !== "undefined" ? window.innerHeight / 2 : 0);

  const clipPathStyle = expandedDone
    ? "none"
    : circleActive
      ? `circle(150vmax at ${posX}px ${posY}px)`
      : `circle(0px at ${posX}px ${posY}px)`;

  return (
    <div
      className={`all-divisi-modal-overlay ${circleActive ? "is-circle-expanded" : ""} ${
        expandedDone ? "is-settled" : ""
      }`}
      data-lenis-prevent
      style={{
        clipPath: clipPathStyle,
      }}
    >
      <div
        className={`all-divisi-modal-inner ${
          contentActive
            ? expandedDone
              ? "is-content-visible is-settled"
              : "is-content-visible"
            : "is-content-closing"
        }`}
      >
        {/* Top Header Bar */}
        <header className="all-divisi-header">
          <div className="all-divisi-header-left">
            <div className="all-divisi-label">
              Seluruh Divisi &middot; DOSCOM
            </div>
            <h2 className="all-divisi-heading">
              Struktur Kepengurusan &amp; Pilar Komunitas
            </h2>
          </div>
          <button
            type="button"
            onClick={handleClose}
            className="all-divisi-close-btn"
            aria-label="Tutup Tampilan Divisi"
          >
            <span>Tutup</span>
            <span className="all-divisi-close-circle">&times;</span>
          </button>
        </header>

        {/* 3x3 Symmetrical Grid with Centerpiece Logo Card */}
        <div className="all-divisi-grid">
          {/* Row 1 & Row 2 Left (Top 4 Divisions) */}
          {DIVISIONS_TOP_4.map((item) => (
            <div key={item.id} className="all-divisi-card">
              <div className="all-divisi-card-img-wrap">
                <img
                  src={item.image}
                  alt={item.title}
                  className="all-divisi-card-img"
                  loading="lazy"
                />
              </div>
              <div className="all-divisi-card-info">
                <span className="all-divisi-card-kicker">{item.kicker}</span>
                <h3 className="all-divisi-card-title">{item.title}</h3>
                <p className="all-divisi-card-desc">{item.description}</p>
              </div>
            </div>
          ))}

          {/* Centerpiece: DOSCOM White Logo Card (Slot 5 in 3x3 Grid) */}
          <div className="all-divisi-card all-divisi-centerpiece">
            <div className="centerpiece-card-inner">
              <div className="centerpiece-brand-wrap">
                <img
                  src="/logo-with-text-white.svg"
                  alt="DOSCOM Logo"
                  className="centerpiece-logo-img"
                />
              </div>
              <div className="centerpiece-text-wrap">
                <span className="centerpiece-est">
                  EST. 2008 &middot; SEMARANG
                </span>
                <p className="centerpiece-desc">
                  Dinus Open Source Community adalah wadah eksplorasi,
                  kolaborasi riset, dan pengembangan perangkat lunak sumber
                  terbuka berbasis komunitas mahasiswa.
                </p>
              </div>
              <div className="centerpiece-badge-row">
                <span className="centerpiece-tag">#OpenSource</span>
                <span className="centerpiece-tag">#DOSCOM</span>
                <span className="centerpiece-tag">#Community</span>
              </div>
            </div>
          </div>

          {/* Row 2 Right & Row 3 (Bottom 4 Divisions) */}
          {DIVISIONS_BOTTOM_4.map((item) => (
            <div key={item.id} className="all-divisi-card">
              <div className="all-divisi-card-img-wrap">
                <img
                  src={item.image}
                  alt={item.title}
                  className="all-divisi-card-img"
                  loading="lazy"
                />
              </div>
              <div className="all-divisi-card-info">
                <span className="all-divisi-card-kicker">{item.kicker}</span>
                <h3 className="all-divisi-card-title">{item.title}</h3>
                <p className="all-divisi-card-desc">{item.description}</p>
              </div>
            </div>
          ))}
        </div>

        {/* Bottom Bar */}
        <div className="all-divisi-footer">
          <button
            type="button"
            onClick={handleClose}
            className="btn btn-primary"
          >
            <span>Kembali ke Beranda</span>
            <Icon name="arrow-right" size={16} />
          </button>
        </div>
      </div>
    </div>
  );
}
