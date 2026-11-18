"use client";

import { useEffect, useState } from "react";
import { Icon } from "../ui";

type SiteFooterProps = {
  onOpenContact: () => void;
};

export function SiteFooter({ onOpenContact }: SiteFooterProps) {
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
              Crafting digital solutions that merge robust functionality with
              refined aesthetic principles.
            </p>
            <button
              type="button"
              onClick={onOpenContact}
              className="btn btn-primary"
            >
              Say hello!
            </button>
            <nav className="footer-social-links" aria-label="Social media">
              <a
                href="https://www.linkedin.com/company/doscomedia/"
                target="_blank"
                rel="noreferrer"
                aria-label="DOSCOM on LinkedIn"
              >
                <Icon name="linkedin" size={20} />
              </a>
              <a
                href="https://www.instagram.com/doscomedia/"
                target="_blank"
                rel="noreferrer"
                aria-label="DOSCOM on Instagram"
              >
                <Icon name="instagram" size={20} />
              </a>
              <a
                href="https://github.com/DOSCOM"
                target="_blank"
                rel="noreferrer"
                aria-label="DOSCOM on GitHub"
              >
                <Icon name="github" size={20} />
              </a>
            </nav>
          </div>
          <div className="footer-links-col">
            <div className="footer-nav-list">
              {[
                { href: "#home", label: "Home" },
                { href: "#portfolio", label: "Portfolio" },
                { href: "#services", label: "Services" },
                { href: "#about", label: "About" },
                { href: "#contact", label: "Contact" },
              ].map((item) => (
                <a key={item.href} href={item.href} className="footer-nav-item">
                  <span>{item.label}</span>
                  <Icon name="arrow-right" size={20} />
                </a>
              ))}
            </div>
          </div>
        </div>

        <div className="footer-contacts-strip">
          <div>
            <div className="contact-strip-title">Inquiries</div>
            <a href="mailto:info@bloomava.com" className="contact-strip-desc">
              info@bloomava.com
            </a>
          </div>
          <div>
            <div className="contact-strip-title">Contact us</div>
            <a href="tel:+15551234567" className="contact-strip-desc">
              (555) 123-4567
            </a>
          </div>
          <div>
            <div className="contact-strip-title">Location</div>
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

        <div className="footer-bottom-bar">
          <div>© 2007 - {currentYear ?? ""} DOSCOM. All rights reserved.</div>
          <div className="footer-legal-links">
            <span>
              Dibuat oleh Dinus Open Source Community dengan ♥️ | Di bawah
              naungan{" "}
              <a
                href="https://www.dinus.ac.id/"
                target="_blank"
                rel="noreferrer"
              >
                Universitas Dian Nuswantoro
              </a>
            </span>
          </div>
        </div>
      </div>
    </footer>
  );
}
