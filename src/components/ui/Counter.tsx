"use client";

import { useEffect, useRef } from "react";

export function Counter({ target }: { target: number }) {
  const ref = useRef<HTMLSpanElement>(null);

  useEffect(() => {
    const el = ref.current;
    if (!el) return;

    let animated = false;
    let frameId: number | undefined;

    const observer = new IntersectionObserver(
      ([entry], obs) => {
        if (animated || !entry.isIntersecting) return;
        animated = true;
        obs.unobserve(el);

        const start = performance.now();
        const step = (now: number) => {
          const progress = Math.min((now - start) / 1500, 1);
          const eased = progress === 1 ? 1 : 1 - 2 ** (-10 * progress);
          el.textContent = String(Math.floor(eased * target));
          if (progress < 1) {
            frameId = requestAnimationFrame(step);
          } else {
            el.textContent = String(target);
          }
        };
        frameId = requestAnimationFrame(step);
      },
      { threshold: 0.2 },
    );

    observer.observe(el);
    return () => {
      observer.disconnect();
      if (frameId !== undefined) cancelAnimationFrame(frameId);
    };
  }, [target]);

  return <span ref={ref}>{target}</span>;
}
