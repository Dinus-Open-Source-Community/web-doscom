"use client";

import type { FormEvent, RefObject } from "react";
import { Icon } from "../ui";

type ContactModalProps = {
  open: boolean;
  submitted: boolean;
  formRef: RefObject<HTMLFormElement | null>;
  onSubmit: (event: FormEvent<HTMLFormElement>) => void;
  onClose: () => void;
};

export function ContactModal({
  open,
  submitted,
  formRef,
  onSubmit,
  onClose,
}: ContactModalProps) {
  const overlayState = open
    ? "opacity-100 pointer-events-auto"
    : "opacity-0 pointer-events-none";

  return (
    <div
      id="contactModal"
      className={`modal-overlay modal-backdrop-translucent ${overlayState}`}
    >
      <button
        type="button"
        className="modal-dismiss"
        onClick={onClose}
        aria-label="Close contact modal"
      />
      <div className="contact-modal-box">
        <button
          type="button"
          onClick={onClose}
          className="contact-modal-close"
          aria-label="Close contact modal"
        >
          <Icon name="close" size={24} />
        </button>
        <div className="contact-modal-header">
          <img src="/logo-doscom.svg" alt="DOSCOM Logo" />
          <span className="contact-modal-badge">Start a Project</span>
        </div>
        <h3 className="contact-modal-title">Say hello to DOSCOM</h3>
        <p className="contact-modal-desc">
          Tell us about your next visionary product or brand identity.
        </p>

        <form
          ref={formRef}
          onSubmit={onSubmit}
          className={`contact-form${submitted ? " hidden" : ""}`}
        >
          <div className="form-group">
            <label htmlFor="contactName">Your Name</label>
            <input
              type="text"
              id="contactName"
              name="name"
              required
              placeholder="Alex Rivera"
            />
          </div>
          <div className="form-group">
            <label htmlFor="contactEmail">Email Address</label>
            <input
              type="email"
              id="contactEmail"
              name="email"
              required
              placeholder="alex@domain.com"
            />
          </div>
          <div className="form-group">
            <label htmlFor="contactMessage">Project Details</label>
            <textarea
              id="contactMessage"
              name="message"
              rows={3}
              required
              placeholder="Briefly describe your goals..."
            />
          </div>
          <button type="submit" className="btn-form-submit">
            Send Message
          </button>
        </form>
        <div
          id="contactSuccess"
          className={`contact-success-state${submitted ? "" : " hidden"}`}
        >
          <Icon name="check" size={40} />
          <h4>Message received!</h4>
          <p>We&apos;ll get back to you within 24 business hours.</p>
        </div>
      </div>
    </div>
  );
}
