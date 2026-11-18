"use client";

import { useEffect, useRef } from "react";
import type { LightboxItem } from "../../types/lightbox";
import { Icon } from "../ui";

type LightboxModalProps = {
  item: LightboxItem | null;
  onClose: () => void;
};

export function LightboxModal({ item, onClose }: LightboxModalProps) {
  const dialogRef = useRef<HTMLDivElement>(null);
  const closeRef = useRef<HTMLButtonElement>(null);
  const returnFocusRef = useRef<Element | null>(null);

  useEffect(() => {
    if (!item) return;

    returnFocusRef.current = document.activeElement;
    closeRef.current?.focus();

    return () => {
      const target = returnFocusRef.current;
      if (target instanceof HTMLElement) target.focus();
    };
  }, [item]);

  if (!item) return null;

  const trapFocus = (event: React.KeyboardEvent) => {
    if (event.key !== "Tab") return;

    const nodes = dialogRef.current?.querySelectorAll<HTMLElement>(
      'a[href], button:not([disabled]), [tabindex]:not([tabindex="-1"])',
    );
    if (!nodes || nodes.length === 0) return;

    const first = nodes[0];
    const last = nodes[nodes.length - 1];

    if (event.shiftKey && document.activeElement === first) {
      event.preventDefault();
      last.focus();
    } else if (!event.shiftKey && document.activeElement === last) {
      event.preventDefault();
      first.focus();
    }
  };

  return (
    <div
      ref={dialogRef}
      id="lightboxModal"
      className="modal-overlay modal-backdrop-dark opacity-100 pointer-events-auto"
      role="dialog"
      aria-modal="true"
      aria-label="Project preview"
      onKeyDown={trapFocus}
    >
      <button
        type="button"
        className="modal-dismiss"
        onClick={onClose}
        aria-label="Close lightbox"
        tabIndex={-1}
        aria-hidden="true"
      />
      <button
        ref={closeRef}
        type="button"
        onClick={onClose}
        className="modal-close-btn"
        aria-label="Close lightbox"
      >
        <Icon name="close" size={24} />
      </button>
      <div className="lightbox-content">
        <img id="lightboxImg" src={item.src} alt="Project Preview" />
        <p id="lightboxCaption">{item.caption}</p>
      </div>
    </div>
  );
}
