"use client";

import Lenis from "lenis";
import "lenis/dist/lenis.css";
import { useEffect, useRef, useState } from "react";
import "../styles/landing.css";
import { MobileDrawer, SiteFooter, StickyNav } from "../components/layout";
import { ContactModal, LightboxModal } from "../components/modals";
import {
  AboutSection,
  HeroSection,
  PengurusSection,
  PortfolioSection,
  ServicesSection,
  ShowcaseSection,
  VideoReelSection,
  WhyUsSection,
} from "../components/sections";
import type { LightboxItem } from "../types/lightbox";

export default function Home() {
  const [stickyVisible, setStickyVisible] = useState(false);
  const [mobileOpen, setMobileOpen] = useState(false);
  const [contactOpen, setContactOpen] = useState(false);
  const [submitted, setSubmitted] = useState(false);
  const [lightbox, setLightbox] = useState<LightboxItem | null>(null);

  const formRef = useRef<HTMLFormElement>(null);
  const resetTimer = useRef<ReturnType<typeof setTimeout> | null>(null);
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
      lenis.start();
      lenis.scrollTo(target as HTMLElement, { offset: -20, duration: 1.2 });
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
    const overlayOpen = contactOpen || lightbox !== null;

    body.classList.toggle("menu-open", mobileOpen);
    html.classList.toggle("menu-open", mobileOpen);
    body.classList.toggle("modal-open", overlayOpen);
    html.classList.toggle("modal-open", overlayOpen);

    const lenis = lenisRef.current;
    if (lenis) {
      if (mobileOpen || overlayOpen) {
        lenis.stop();
      } else {
        lenis.start();
      }
    }

    return () => {
      body.classList.remove("menu-open", "modal-open");
      html.classList.remove("menu-open", "modal-open");
    };
  }, [mobileOpen, contactOpen, lightbox]);

  useEffect(() => {
    const onKeyDown = (event: KeyboardEvent) => {
      if (event.key !== "Escape") return;
      setMobileOpen(false);
      setContactOpen(false);
      setSubmitted(false);
      setLightbox(null);
    };
    window.addEventListener("keydown", onKeyDown);
    return () => window.removeEventListener("keydown", onKeyDown);
  }, []);

  useEffect(
    () => () => {
      if (resetTimer.current) clearTimeout(resetTimer.current);
    },
    [],
  );

  const scrollToSection = (id: string) => {
    const target = document.getElementById(id);
    if (!target) return;
    if (lenisRef.current) {
      lenisRef.current.scrollTo(target, { offset: -20, duration: 1.2 });
    } else {
      target.scrollIntoView({ behavior: "smooth", block: "start" });
    }
  };

  const openContactModal = () => {
    setMobileOpen(false);
    setContactOpen(true);
  };

  const closeContactModal = () => {
    setContactOpen(false);
    setSubmitted(false);
  };

  const handleContactSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setSubmitted(true);
    if (resetTimer.current) clearTimeout(resetTimer.current);
    resetTimer.current = setTimeout(() => {
      setContactOpen(false);
      setSubmitted(false);
      formRef.current?.reset();
    }, 2500);
  };

  return (
    <>
      <StickyNav
        visible={stickyVisible}
        onOpenContact={openContactModal}
        onScrollTo={scrollToSection}
        onOpenMenu={() => setMobileOpen(true)}
      />
      <MobileDrawer
        onClose={() => setMobileOpen(false)}
        onOpenContact={openContactModal}
      />
      <HeroSection
        onOpenContact={openContactModal}
        onOpenMenu={() => setMobileOpen(true)}
      />
      <AboutSection />
      <PortfolioSection
        onScrollTo={scrollToSection}
        onOpenLightbox={setLightbox}
      />
      <ServicesSection onOpenContact={openContactModal} />
      <ShowcaseSection onOpenLightbox={setLightbox} />
      <WhyUsSection onOpenContact={openContactModal} />
      <PengurusSection />
      <VideoReelSection onOpenLightbox={setLightbox} />
      <SiteFooter onOpenContact={openContactModal} />
      <LightboxModal item={lightbox} onClose={() => setLightbox(null)} />
      <ContactModal
        open={contactOpen}
        submitted={submitted}
        formRef={formRef}
        onSubmit={handleContactSubmit}
        onClose={closeContactModal}
      />
    </>
  );
}
