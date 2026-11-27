"use client";

import Lenis from "lenis";
import "lenis/dist/lenis.css";
import { useEffect, useRef, useState } from "react";
import "../styles/landing.css";
import { MobileDrawer, SiteFooter, StickyNav } from "../components/layout";
import { LightboxModal } from "../components/modals";
import {
  AboutSection,
  HeroSection,
  PengurusSection,
  PortfolioSection,
  ProjectSection,
  ServicesSection,
  ShowcaseSection,
  VideoReelSection,
  WhyUsSection,
} from "../components/sections";
import type { LightboxItem } from "../types/lightbox";

const STICKY_NAV_SCROLL_THRESHOLD = 350;

function getAnchorOffset(target: HTMLElement) {
  const stickyNav = document.getElementById("stickyNav");
  const navHeight = stickyNav?.offsetHeight ?? 0;
  const targetTop = target.getBoundingClientRect().top + window.scrollY;

  if (navHeight > 0 && targetTop > STICKY_NAV_SCROLL_THRESHOLD + navHeight) {
    return -(navHeight + 16);
  }

  return -20;
}

function scrollToAnchor(lenis: Lenis, target: HTMLElement) {
  const scroll = () => {
    lenis.start();
    lenis.resize();
    lenis.scrollTo(target, {
      offset: getAnchorOffset(target),
      duration: 1.2,
    });
  };

  if (document.documentElement.classList.contains("menu-open")) {
    requestAnimationFrame(() => requestAnimationFrame(scroll));
  } else {
    scroll();
  }
}

export default function Home() {
  const [stickyVisible, setStickyVisible] = useState(false);
  const [mobileOpen, setMobileOpen] = useState(false);
  const [lightbox, setLightbox] = useState<LightboxItem | null>(null);
  const [projectModalOpen, setProjectModalOpen] = useState(false);

  const lenisRef = useRef<Lenis | null>(null);

  useEffect(() => {
    const lenis = new Lenis({
      duration: 1.2,
      easing: (t) => Math.min(1, 1.001 - 2 ** (-10 * t)),
      orientation: "vertical",
      gestureOrientation: "vertical",
      smoothWheel: true,
      wheelMultiplier: 1,
      touchMultiplier: 2,
      infinite: false,
    });
    lenisRef.current = lenis;

    const onScroll = ({ scroll }: { scroll: number }) =>
      setStickyVisible(scroll > 350);
    lenis.on("scroll", onScroll);
    onScroll({ scroll: window.scrollY });

    let rafId = requestAnimationFrame(function raf(time) {
      lenis.raf(time);
      rafId = requestAnimationFrame(raf);
    });

    const onAnchorClick = (event: MouseEvent) => {
      const anchor = (event.target as Element | null)?.closest?.(
        'a[href^="#"]',
      );
      const href = anchor?.getAttribute("href");
      if (!href || href === "#") return;
      const target = document.querySelector(href);
      if (!target) return;

      event.preventDefault();
      scrollToAnchor(lenis, target as HTMLElement);
    };
    document.addEventListener("click", onAnchorClick);

    return () => {
      cancelAnimationFrame(rafId);
      document.removeEventListener("click", onAnchorClick);
      lenis.off("scroll", onScroll);
      lenis.destroy();
      lenisRef.current = null;
    };
  }, []);

  useEffect(() => {
    const html = document.documentElement;
    const body = document.body;
    const lightboxOpen = lightbox !== null;
    const scrollLocked = mobileOpen || lightboxOpen || projectModalOpen;

    body.classList.toggle("menu-open", mobileOpen);
    html.classList.toggle("menu-open", mobileOpen);
    // ProjectModal mengelola scroll lock-nya sendiri lewat class
    // project-modal-open, jadi jangan pakai modal-open milik lightbox.
    body.classList.toggle("modal-open", lightboxOpen);
    html.classList.toggle("modal-open", lightboxOpen);

    const lenis = lenisRef.current;
    if (lenis) {
      if (scrollLocked) {
        lenis.stop();
      } else {
        lenis.start();
      }
    }

    return () => {
      body.classList.remove("menu-open", "modal-open");
      html.classList.remove("menu-open", "modal-open");
    };
  }, [mobileOpen, lightbox, projectModalOpen]);

  useEffect(() => {
    const onKeyDown = (event: KeyboardEvent) => {
      if (event.key !== "Escape") return;
      setMobileOpen(false);
      setLightbox(null);
    };
    window.addEventListener("keydown", onKeyDown);
    return () => window.removeEventListener("keydown", onKeyDown);
  }, []);

  const scrollToSection = (id: string) => {
    const target = document.getElementById(id);
    if (!target) return;
    if (lenisRef.current) {
      scrollToAnchor(lenisRef.current, target);
    } else {
      target.scrollIntoView({ behavior: "smooth", block: "start" });
    }
  };

  return (
    <>
      <StickyNav
        visible={stickyVisible}
        onScrollTo={scrollToSection}
        onOpenMenu={() => setMobileOpen(true)}
      />
      <MobileDrawer onClose={() => setMobileOpen(false)} />
      <HeroSection onOpenMenu={() => setMobileOpen(true)} />
      <AboutSection />
      <PortfolioSection
        onScrollTo={scrollToSection}
        onOpenLightbox={setLightbox}
      />
      <ServicesSection />
      <ProjectSection onModalOpenChange={setProjectModalOpen} />
      <ShowcaseSection onOpenLightbox={setLightbox} />
      <WhyUsSection />
      <PengurusSection />
      <VideoReelSection onOpenLightbox={setLightbox} />
      <SiteFooter />
      <LightboxModal item={lightbox} onClose={() => setLightbox(null)} />
    </>
  );
}
