"use client";

import { useEffect, useState } from "react";

const GLYPHS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_#@*&~";

type ScrambleTextProps = {
  text: string;
  trigger?: boolean;
  className?: string;
  speed?: number;
};

export function ScrambleText({
  text,
  trigger = true,
  className,
  speed = 30,
}: ScrambleTextProps) {
  const [displayText, setDisplayText] = useState(text);

  useEffect(() => {
    if (!trigger) {
      setDisplayText(text);
      return;
    }

    let frame = 0;
    const totalFrames = text.length * 3;
    let timer: NodeJS.Timeout | null = null;

    const tick = () => {
      frame++;
      const progress = frame / totalFrames;
      const revealCount = Math.floor(progress * text.length);

      const scrambled = text
        .split("")
        .map((char, index) => {
          if (char === " ") return " ";
          if (index < revealCount) return char;
          const randomChar = GLYPHS[Math.floor(Math.random() * GLYPHS.length)];
          return randomChar;
        })
        .join("");

      setDisplayText(scrambled);

      if (frame < totalFrames) {
        timer = setTimeout(tick, speed);
      } else {
        setDisplayText(text);
      }
    };

    timer = setTimeout(tick, speed);

    return () => {
      if (timer) clearTimeout(timer);
    };
  }, [trigger, text, speed]);

  return <span className={className}>{displayText}</span>;
}
