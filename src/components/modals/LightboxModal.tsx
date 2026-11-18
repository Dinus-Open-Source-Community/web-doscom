"use client";

import type { LightboxItem } from "../../types/lightbox";
import { Icon } from "../ui";

type LightboxModalProps = {
  item: LightboxItem | null;
  onClose: () => void;
};

export function LightboxModal({ item, onClose }: LightboxModalProps) {
  const overlayState = item
    ? "opacity-100 pointer-events-auto"
    : "opacity-0 pointer-events-none";

  return (
    <div
      id="lightboxModal"
      className={`modal-overlay modal-backdrop-dark ${overlayState}`}
    >
      <button
        type="button"
        className="modal-dismiss"
        onClick={onClose}
        aria-label="Close lightbox"
      />
      <button
        type="button"
        onClick={onClose}
        className="modal-close-btn"
        aria-label="Close lightbox"
      >
        <Icon name="close" size={24} />
      </button>
      <div className="lightbox-content">
        {item ? (
          <>
            <img id="lightboxImg" src={item.src} alt="Project Preview" />
            <p id="lightboxCaption">{item.caption}</p>
          </>
        ) : null}
      </div>
    </div>
  );
}
