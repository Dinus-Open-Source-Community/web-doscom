"use client";

import { Counter, Icon } from "../ui";

type ServicesSectionProps = {
  onOpenContact: () => void;
};

export function ServicesSection({ onOpenContact }: ServicesSectionProps) {
  return (
    <section id="services" className="container services-section">
      <div className="services-top-row">
        <div className="services-label-desktop">
          <div className="section-label sticky-label">Program Kerja</div>
        </div>
        <div className="services-top-content">
          <div className="section-label services-label-mobile">
            Program Kerja
          </div>
          <h2 className="services-heading">
            Your story deserves more than just visuals - we design with meaning
          </h2>
          <div className="services-big-counter">
            <div className="big-number-wrap">
              <div className="big-number">
                <Counter target={15} />
              </div>
              <div className="big-number-plus">+</div>
            </div>
            <div className="big-number-label">Years of experience</div>
          </div>
        </div>
      </div>

      <div className="services-main-grid">
        <div className="services-left-col">
          <h3 className="services-subheading">
            We build brands that break the mold
          </h3>
          <button
            type="button"
            onClick={onOpenContact}
            className="btn btn-primary"
          >
            View Services
          </button>
        </div>
        <div className="services-cards-col">
          <button
            type="button"
            className="service-card"
            onClick={onOpenContact}
          >
            <div>
              <div className="service-card-header">
                <h4 className="service-card-title">Branding &amp; Identity</h4>
                <Icon
                  name="arrow-right-up"
                  className="service-card-icon"
                  size={20}
                />
              </div>
              <p className="service-card-desc">
                Strategic positioning and visual identity systems that resonate.
              </p>
            </div>
            <div className="service-card-cta">Learn more &rarr;</div>
          </button>

          <button
            type="button"
            className="service-card"
            onClick={onOpenContact}
          >
            <div>
              <div className="service-card-header">
                <h4 className="service-card-title">Web &amp; Digital Design</h4>
                <Icon
                  name="arrow-right-up"
                  className="service-card-icon"
                  size={20}
                />
              </div>
              <p className="service-card-desc">
                Immersive digital experiences tailored for conversion.
              </p>
            </div>
            <div className="service-card-cta">Learn more &rarr;</div>
          </button>

          <div className="service-card-inverted">
            <div className="inverted-label">Customer satisfaction rate</div>
            <div className="inverted-counter">
              <Counter target={95} />%
            </div>
            <div className="inverted-glow" />
          </div>
        </div>
      </div>
    </section>
  );
}
