"use client";

import { useEffect, useId, useRef, useState } from "react";
import {
  type PengurusMember,
  pengurus,
  pengurusPeriods,
} from "../../data/pengurus";
import Lanyard from "../ui/Lanyard/Lanyard";

function initials(name: string) {
  return name
    .split(/\s+/)
    .filter(Boolean)
    .slice(0, 2)
    .map((word) => word[0]?.toUpperCase() ?? "")
    .join("");
}

function PengurusCard({
  person,
  onSelect,
}: {
  person: PengurusMember;
  onSelect: (person: PengurusMember) => void;
}) {
  const [imageFailed, setImageFailed] = useState(false);
  const showImage = Boolean(person.image) && !imageFailed;

  return (
    <button
      type="button"
      className="pengurus-card"
      onClick={() => onSelect(person)}
      aria-label={`Lihat detail ${person.name}, ${person.role}`}
    >
      <div className="pengurus-portrait">
        {showImage ? (
          <img
            src={person.image}
            alt={`Foto ${person.name}`}
            width={480}
            height={600}
            loading="lazy"
            onError={() => setImageFailed(true)}
          />
        ) : (
          <span className="pengurus-initials" aria-hidden="true">
            {initials(person.name)}
          </span>
        )}
      </div>
      <div className="pengurus-body">
        <span className="pengurus-name">{person.name}</span>
        <span className="pengurus-role">{person.role}</span>
      </div>
    </button>
  );
}

function PengurusModal({
  person,
  onClose,
}: {
  person: PengurusMember;
  onClose: () => void;
}) {
  const dialogRef = useRef<HTMLDialogElement>(null);
  const history = person.history
    ?.slice()
    .sort(
      (a, b) => Number.parseInt(b.period, 10) - Number.parseInt(a.period, 10),
    );

  useEffect(() => {
    const dialog = dialogRef.current;
    if (dialog && !dialog.open) dialog.showModal();
  }, []);

  return (
    <dialog
      ref={dialogRef}
      className="pengurus-modal"
      data-lenis-prevent
      aria-label={`Detail ${person.name}`}
      onClose={onClose}
      onClick={(event) => {
        if (event.target === event.currentTarget) event.currentTarget.close();
      }}
      onKeyDown={(event) => {
        if (event.key === "Escape") event.currentTarget.close();
      }}
    >
      <button
        type="button"
        className="pengurus-modal-close"
        aria-label="Tutup detail pengurus"
        onClick={() => dialogRef.current?.close()}
      >
        <svg
          width="20"
          height="20"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          strokeWidth="1.75"
          strokeLinecap="round"
          aria-hidden="true"
        >
          <path d="m6 6 12 12M18 6 6 18" />
        </svg>
      </button>
      <div className="pengurus-modal-content">
        <div className="pengurus-lanyard">
          <Lanyard
            frontImage={person.image}
            backImage="/logo-doscom.svg"
            imageFit="cover"
            cardColor="#ffffff"
            orientation="portrait"
            finish="glossy"
            cornerRadius={0.3}
            size={0.6}
            anchor="center"
            strapLength={0.5}
            strapColor="#111111"
            strapWidth={0.65}
            metal="silver"
            gravity={1}
            damping={0.5}
            elasticity={0.5}
            breeze={0.5}
            interactive
            intro
          />
        </div>
        <div className="pengurus-modal-details">
          <h2 className="pengurus-modal-name">{person.name}</h2>
          <p className="pengurus-modal-role">{person.role}</p>
          <dl className="pengurus-modal-meta">
            <div>
              <dt>Divisi</dt>
              <dd>{person.division}</dd>
            </div>
            <div>
              <dt>Periode kepengurusan</dt>
              <dd>{person.period}</dd>
            </div>
          </dl>
          <p className="pengurus-modal-description">{person.description}</p>
          <section
            className="pengurus-modal-section"
            aria-labelledby="pengurus-social-title"
          >
            <h3 id="pengurus-social-title">Media sosial</h3>
            {person.socialMedia?.length ? (
              <ul className="pengurus-social-list">
                {person.socialMedia.map(({ platform, url }) => (
                  <li key={`${platform}-${url}`}>
                    <a href={url} target="_blank" rel="noreferrer">
                      {platform}
                    </a>
                  </li>
                ))}
              </ul>
            ) : (
              <p className="pengurus-modal-empty">
                Belum ada tautan media sosial.
              </p>
            )}
          </section>
          <section
            className="pengurus-modal-section"
            aria-labelledby="pengurus-history-title"
          >
            <h3 id="pengurus-history-title">Riwayat keanggotaan</h3>
            {history?.length ? (
              <ol className="pengurus-history-list">
                {history.map(({ role, period }) => (
                  <li key={`${role}-${period}`}>
                    <span>{role}</span>
                    <span>{period}</span>
                  </li>
                ))}
              </ol>
            ) : (
              <p className="pengurus-modal-empty">Riwayat belum ditambahkan.</p>
            )}
          </section>
        </div>
      </div>
    </dialog>
  );
}

export function PengurusSection() {
  const [period, setPeriod] = useState<string>(pengurusPeriods[0] ?? "");
  const [open, setOpen] = useState(false);
  const [expanded, setExpanded] = useState(false);
  const [selectedPerson, setSelectedPerson] = useState<PengurusMember | null>(
    null,
  );
  const pickerRef = useRef<HTMLDivElement>(null);
  const triggerRef = useRef<HTMLButtonElement>(null);
  const optionRefs = useRef<Array<HTMLButtonElement | null>>([]);
  const labelId = useId();
  const listboxId = useId();

  const selectedIndex = Math.max(pengurusPeriods.indexOf(period), 0);
  const members = pengurus.filter((person) => person.period === period);
  const remainingCount = Math.max(members.length - 8, 0);
  const visibleMembers = expanded ? members : members.slice(0, 8);

  useEffect(() => {
    if (!open) return;
    optionRefs.current[selectedIndex]?.focus();

    const onPointerDown = (event: PointerEvent) => {
      if (!pickerRef.current?.contains(event.target as Node)) setOpen(false);
    };
    document.addEventListener("pointerdown", onPointerDown);
    return () => document.removeEventListener("pointerdown", onPointerDown);
  }, [open, selectedIndex]);

  const choose = (value: string) => {
    setPeriod(value);
    setOpen(false);
    triggerRef.current?.focus();
  };

  const onTriggerKeyDown = (event: React.KeyboardEvent<HTMLButtonElement>) => {
    if (
      event.key === "ArrowDown" ||
      event.key === "Enter" ||
      event.key === " "
    ) {
      event.preventDefault();
      setOpen(true);
    }
    if (event.key === "ArrowUp") {
      event.preventDefault();
      setOpen(true);
    }
  };

  const onOptionKeyDown = (
    event: React.KeyboardEvent<HTMLButtonElement>,
    index: number,
  ) => {
    const last = pengurusPeriods.length - 1;

    if (event.key === "ArrowDown" || event.key === "ArrowUp") {
      event.preventDefault();
      const next =
        event.key === "ArrowDown"
          ? index === last
            ? 0
            : index + 1
          : index === 0
            ? last
            : index - 1;
      optionRefs.current[next]?.focus();
      return;
    }
    if (event.key === "Home" || event.key === "End") {
      event.preventDefault();
      optionRefs.current[event.key === "Home" ? 0 : last]?.focus();
      return;
    }
    if (event.key === "Enter" || event.key === " ") {
      event.preventDefault();
      choose(pengurusPeriods[index]);
      return;
    }
    if (event.key === "Escape") {
      event.preventDefault();
      setOpen(false);
      triggerRef.current?.focus();
      return;
    }
    if (event.key === "Tab") setOpen(false);
  };

  return (
    <section id="pengurus" className="container pengurus-section">
      <div className="pengurus-header">
        <div className="section-label">Pengurus</div>
        <div className="pengurus-period-picker" ref={pickerRef}>
          <span id={labelId} className="pengurus-period-label">
            Periode
          </span>
          <div className="pengurus-select-wrap">
            <button
              ref={triggerRef}
              type="button"
              className="pengurus-select"
              aria-haspopup="listbox"
              aria-expanded={open}
              aria-controls={open ? listboxId : undefined}
              aria-labelledby={`${labelId} ${labelId}-value`}
              onClick={() => setOpen((value) => !value)}
              onKeyDown={onTriggerKeyDown}
            >
              <span id={`${labelId}-value`}>{period}</span>
              <svg
                className="pengurus-select-chevron"
                width="16"
                height="16"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                strokeWidth="2"
                strokeLinecap="round"
                strokeLinejoin="round"
                aria-hidden="true"
              >
                <path d="m6 9 6 6 6-6" />
              </svg>
            </button>
            {open ? (
              <div
                id={listboxId}
                role="listbox"
                aria-labelledby={labelId}
                className="pengurus-select-menu"
                data-lenis-prevent
              >
                {pengurusPeriods.map((option, index) => (
                  <button
                    key={option}
                    ref={(node) => {
                      optionRefs.current[index] = node;
                    }}
                    type="button"
                    role="option"
                    aria-selected={option === period}
                    className="pengurus-option"
                    onClick={() => choose(option)}
                    onKeyDown={(event) => onOptionKeyDown(event, index)}
                  >
                    {option}
                  </button>
                ))}
              </div>
            ) : null}
          </div>
        </div>
      </div>
      <div className="pengurus-list">
        {visibleMembers.map((person) => (
          <PengurusCard
            key={`${person.period}-${person.role}-${person.name}`}
            person={person}
            onSelect={setSelectedPerson}
          />
        ))}
      </div>
      {members.length > 8 ? (
        <div className="pengurus-more-wrap">
          <button
            type="button"
            className="pengurus-more-btn"
            aria-expanded={expanded}
            onClick={() => setExpanded((prev) => !prev)}
          >
            <span>
              {expanded
                ? "Tampilkan Lebih Sedikit"
                : `Lihat Semua Pengurus (${remainingCount})`}
            </span>
            <svg
              className={`pengurus-more-chevron ${expanded ? "is-expanded" : ""}`}
              width="16"
              height="16"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              strokeWidth="2"
              strokeLinecap="round"
              strokeLinejoin="round"
              aria-hidden="true"
            >
              <path d="m6 9 6 6 6-6" />
            </svg>
          </button>
        </div>
      ) : null}
      {selectedPerson ? (
        <PengurusModal
          person={selectedPerson}
          onClose={() => setSelectedPerson(null)}
        />
      ) : null}
    </section>
  );
}
