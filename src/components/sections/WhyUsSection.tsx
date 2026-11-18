"use client";

import { Counter, Icon } from "../ui";

type WhyUsSectionProps = {
  onOpenContact: () => void;
};

export function WhyUsSection({ onOpenContact }: WhyUsSectionProps) {
  return (
    <section id="why-us" className="container why-us-section">
      <div className="why-us-top-row">
        <div className="why-us-label-desktop">
          <div className="section-label sticky-label">Project Kami</div>
        </div>
        <div className="why-us-top-col">
          <div className="section-label why-us-label-mobile">Project Kami</div>
          <h2 className="why-us-heading">
            We design with purpose, build with{" "}
            <span className="text-muted">passion, and deliver with impact</span>
          </h2>
        </div>
      </div>

      <div className="why-us-content-grid">
        <div className="why-us-portrait-col">
          <div className="why-us-portrait-card">
            <img src="/images/why-portrait.jpg" alt="artistic profile" />
          </div>
        </div>
        <div className="why-us-cards-col">
          <button type="button" className="why-card" onClick={onOpenContact}>
            <div className="why-card-top">
              <div className="why-card-thumbnail bg-dark">
                <img src="/images/why-thumb-1.jpg" alt="skincare bottles" />
              </div>
              <div className="why-card-arrow-btn">
                <Icon name="arrow-right-up" size={16} />
              </div>
            </div>
            <div>
              <h4 className="why-card-title">We give brands a visual voice</h4>
              <p className="why-card-desc">
                Translating core values into compelling visual narratives that
                command attention.
              </p>
            </div>
          </button>

          <button type="button" className="why-card" onClick={onOpenContact}>
            <div className="why-card-top">
              <div className="why-card-thumbnail bg-light">
                <img src="/images/why-thumb-2.jpg" alt="red paint splash" />
              </div>
              <div className="why-card-arrow-btn">
                <Icon name="arrow-right-up" size={16} />
              </div>
            </div>
            <div>
              <h4 className="why-card-title">We design with intention</h4>
              <p className="why-card-desc">
                Every pixel placed has a purpose, ensuring form always follows
                function beautifully.
              </p>
            </div>
          </button>

          <div className="why-mini-card why-mini-satisfaction">
            <div className="counter-value">
              <Counter target={95} />%
            </div>
            <div className="why-mini-label">Customer satisfaction rate</div>
          </div>

          <div className="why-mini-card why-mini-avatars">
            <div className="avatar-stack">
              <img src="/images/avatar-4.jpg" alt="avatar" />
              <img src="/images/avatar-5.jpg" alt="avatar" />
            </div>
            <div className="why-mini-avatars-caption">
              Sustainable creativity, timeless appeal.
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
