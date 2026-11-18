"use client";

import { Counter, Icon, ProgressBar } from "../ui";

export function AboutSection() {
  return (
    <section id="about" className="container about-section">
      <div className="about-label-desktop">
        <div className="section-label sticky-label">Tentang Kami</div>
      </div>
      <div className="about-img-col">
        <div className="section-label about-label-mobile">Tentang Kami</div>
        <div className="about-img-card">
          <img src="/images/about.jpg" alt="cinematic lighting" />
        </div>
      </div>
      <div className="about-content-col">
        <h2 className="about-heading">
          The role of user experience in modern design architectures.
        </h2>
        <div className="about-stats-grid">
          <div>
            <div className="counter-value">
              <Counter target={4} />K
            </div>
            <h3 className="about-stat-title">Happy customers</h3>
            <p className="about-stat-desc">
              We believe that the key to our success lies in understanding the
              core needs of our clients.
            </p>
            <div className="avatar-group">
              <div className="avatar-stack">
                <img src="/images/avatar-1.jpg" alt="avatar" />
                <img src="/images/avatar-2.jpg" alt="avatar" />
                <img src="/images/avatar-3.jpg" alt="avatar" />
              </div>
              <span className="avatar-caption">Sustainable creativity.</span>
            </div>
          </div>
          <div className="satisfaction-card">
            <Icon name="like" className="satisfaction-icon" size={24} />
            <div>
              <div className="satisfaction-label">Satisfaction rate</div>
              <div className="counter-value">
                <Counter target={95} />%
              </div>
            </div>
          </div>
        </div>

        <div className="progress-list">
          <div>
            <div className="progress-header">
              <span>Strategy</span>
              <span>75%</span>
            </div>
            <div className="progress-track">
              <ProgressBar value="75%" />
            </div>
          </div>
          <div>
            <div className="progress-header">
              <span>Design</span>
              <span>85%</span>
            </div>
            <div className="progress-track">
              <ProgressBar value="85%" />
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
