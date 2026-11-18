"use client";

import { useEffect, useState } from "react";
import { navigationItems } from "../../data/navigation";
import { Icon } from "../ui";

export function SiteFooter() {
  const [currentYear, setCurrentYear] = useState<number | null>(null);

  useEffect(() => setCurrentYear(new Date().getFullYear()), []);

  return (
    <footer id="contact" className="site-footer">
      <div className="container">
        <div className="footer-main-grid">
          <div className="footer-brand-col">
            <div className="footer-logo-row">
              <img src="/logo-doscom.svg" alt="DOSCOM Logo" />
              <span>DOSCOM</span>
            </div>
            <p className="footer-tagline">
              Open Your Mind with Open Source.
              <br />
              Tempat untuk belajar tanpa batas, berbagi tanpa ragu, berkarya
              bersama, dan menciptakan dampak nyata melalui semangat kolaborasi
              open source.
            </p>
            <nav className="footer-social-links" aria-label="Social media">
              <a
                href="https://www.linkedin.com/company/doscomedia"
                target="_blank"
                rel="noreferrer"
                aria-label="DOSCOM on LinkedIn"
              >
                <Icon name="linkedin" size={18} />
                <span>LinkedIn</span>
              </a>
              <a
                href="https://www.instagram.com/doscomedia/"
                target="_blank"
                rel="noreferrer"
                aria-label="DOSCOM on Instagram"
              >
                <Icon name="instagram" size={18} />
                <span>Instagram</span>
              </a>
              <a
                href="https://github.com/Dinus-Open-Source-Community"
                target="_blank"
                rel="noreferrer"
                aria-label="DOSCOM on GitHub"
              >
                <Icon name="github" size={18} />
                <span>GitHub</span>
              </a>
            </nav>
            <a className="footer-email" href="mailto:doscom.go@gmail.com">
              <Icon name="mail" size={18} />
              <span>doscom.go@gmail.com</span>
            </a>
          </div>
          <nav className="footer-nav-list" aria-label="Footer navigation">
            {navigationItems.map((item) => (
              <a key={item.href} href={item.href} className="footer-nav-item">
                <span>{item.label}</span>
                <Icon name="arrow-right" size={20} />
              </a>
            ))}
          </nav>
          <div className="footer-contacts-strip">
            <div>
              <div className="contact-strip-title">Location</div>
              <p className="footer-address">
                Gedung D, Universitas Dian Nuswantoro, Jl. Nakula I No.5-11,
                Semarang
              </p>
              <iframe
                className="footer-map"
                src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d247.51506315180333!2d110.40929758056724!3d-6.9808621806776525!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2e708b4c1fb56ab3%3A0xbfd9daa65aacb459!2sJl.%20Sadewa%202%208%2C%20Pendrikan%20Kidul%2C%20Kec.%20Semarang%20Tengah%2C%20Kota%20Semarang%2C%20Jawa%20Tengah%2050131!5e0!3m2!1sid!2sid!4v1794981281954!5m2!1sid!2sid"
                title="DOSCOM location on Google Maps"
                width="600"
                height="450"
                style={{ border: 0 }}
                allowFullScreen
                loading="lazy"
                referrerPolicy="strict-origin-when-cross-origin"
              />
            </div>
          </div>
        </div>

        <div className="footer-bottom-bar">
          <div>
            © 2007 - {currentYear ?? ""} DOSCOM,{" "}
            <a href="https://www.dinus.ac.id/" target="_blank" rel="noreferrer">
              Universitas Dian Nuswantoro
            </a>
            . All rights reserved.
          </div>
          <div className="footer-legal-links">
            <span>Dibuat oleh Dinus Open Source Community dengan ♥️</span>
          </div>
        </div>
      </div>
    </footer>
  );
}
