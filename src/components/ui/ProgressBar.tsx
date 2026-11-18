"use client";

import { useEffect, useRef } from "react";

export function ProgressBar({ value }: { value: string }) {
  const ref = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const el = ref.current;
    if (!el) return;

    const observer = new IntersectionObserver(
      ([entry], obs) => {
        if (!entry.isIntersecting) return;
        el.style.width = value;
        obs.unobserve(el);
      },
      { threshold: 0.3 },
    );

    observer.observe(el);
    return () => observer.disconnect();
  }, [value]);

  return <div ref={ref} className="progress-bar" />;
}
