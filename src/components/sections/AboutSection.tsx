"use client";

import Image from "next/image";
import { Counter } from "../ui";

const stats = [
  { target: 4, suffix: "", label: "Divisi inti" },
  { target: 11, suffix: "", label: "Bidang keahlian" },
  { target: 100, suffix: "%", label: "Open source" },
] as const;

export function AboutSection() {
  return (
    <section id="about" className="about-section">
      <div className="container about-inner">
        <div className="about-grid">
          <figure className="about-media">
            <div className="about-image-frame">
              <Image
                src="/images/about-community.webp"
                alt="Anggota DOSCOM berfoto bersama dalam kegiatan komunitas open source"
                fill
                sizes="(max-width: 1023px) 90vw, 40vw"
                className="about-image"
              />
            </div>
            <figcaption className="about-caption">
              Foto kegiatan komunitas DOSCOM.
            </figcaption>
          </figure>

          <div className="about-content">
            <p className="about-eyebrow">
              <span className="about-eyebrow-bar" aria-hidden="true" />
              Tentang Kami
            </p>
            <h2 className="about-heading">
              Komunitas Open Source untuk{" "}
              <span>Belajar, Berbagi, dan Berkarya.</span>
            </h2>
            <p className="about-description">
              DOSCOM adalah komunitas open source di Universitas Dian Nuswantoro
              yang berfokus pada pengembangan, edukasi, dan kolaborasi dalam
              dunia teknologi terbuka. Kami percaya bahwa semangat open source
              bukan hanya tentang kode, tetapi juga tentang berbagi pengetahuan,
              membangun bersama, dan tumbuh sebagai satu komunitas.
            </p>

            <div className="about-stats">
              {stats.map((stat) => (
                <article className="about-stat" key={stat.label}>
                  <p className="about-stat-value">
                    <Counter target={stat.target} />
                    {stat.suffix ? (
                      <span className="about-stat-suffix">{stat.suffix}</span>
                    ) : null}
                  </p>
                  <h3 className="about-stat-label">{stat.label}</h3>
                </article>
              ))}
            </div>
          </div>
        </div>

        <blockquote className="about-vision">
          <p className="about-vision-label">Visi DOSCOM</p>
          <p className="about-vision-quote">
            “Memasyarakatkan Open Source dan Meng-Open Source-kan Masyarakat.”
          </p>
          <footer>Dinus Open Source Community</footer>
        </blockquote>
      </div>
    </section>
  );
}
