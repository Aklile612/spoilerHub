import type {
  StorySection,
  CharacterFate,
  KeyMoment,
  InterpretationSection,
} from "~/types/movie";

/**
 * Parses the structured markdown spoiler text from Gemini into
 * typed sections for the premium detail page.
 */
export function useSpoilerParser(spoilerText: Ref<string | undefined>) {
  /** Extract content between two h2 headings */
  function extractSection(text: string, heading: string): string {
    // Match ## heading (possibly with emoji) and capture until next ## or end
    const escaped = heading.replace(/[.*+?^${}()|[\]\\]/g, "\\$&");
    const regex = new RegExp(
      `##\\s*(?:⚠️\\s*)?${escaped}[\\s\\S]*?\\n([\\s\\S]*?)(?=\\n##\\s|$)`,
      "i"
    );
    const match = text.match(regex);
    return match ? match[1].trim() : "";
  }

  /** Movie overview (non-spoiler) */
  const movieOverview = computed(() => {
    if (!spoilerText.value) return "";
    return extractSection(spoilerText.value, "Movie Overview");
  });

  /** Story breakdown sections */
  const storySections = computed((): StorySection[] => {
    if (!spoilerText.value) return [];
    const sections: StorySection[] = [];

    const sectionDefs = [
      { id: "beginning", title: "The Beginning", icon: "🎬" },
      { id: "turning-point", title: "Major Turning Point", icon: "🔄" },
      { id: "climax", title: "The Climax", icon: "⚡" },
      { id: "ending", title: "Ending Explained", icon: "🎭" },
      { id: "post-credit", title: "Post-Credit Scene", icon: "🎞️" },
    ];

    for (const def of sectionDefs) {
      const content = extractSection(spoilerText.value!, def.title);
      if (content) {
        sections.push({ ...def, content });
      }
    }

    return sections;
  });

  /** Character fates */
  const characterFates = computed((): CharacterFate[] => {
    if (!spoilerText.value) return [];
    const section = extractSection(spoilerText.value, "Character Fates");
    if (!section) return [];

    const fates: CharacterFate[] = [];
    // Match: - **Name** | Actor | STATUS | Summary
    const lines = section.split("\n").filter((l) => l.trim().startsWith("-"));

    for (const line of lines) {
      // Try format: **Name** | Actor | STATUS | summary
      const match = line.match(
        /\*\*(.+?)\*\*\s*\|\s*(.+?)\s*\|\s*(ALIVE|DEAD|UNKNOWN)\s*\|\s*(.+)/i
      );
      if (match) {
        fates.push({
          name: match[1].trim(),
          actor: match[2].trim(),
          status: match[3].trim().toUpperCase() as "ALIVE" | "DEAD" | "UNKNOWN",
          summary: match[4].trim(),
        });
      } else {
        // Fallback: try **Name** — description
        const fallback = line.match(/\*\*(.+?)\*\*\s*[-—]\s*(.+)/);
        if (fallback) {
          fates.push({
            name: fallback[1].trim(),
            actor: "",
            status: "UNKNOWN",
            summary: fallback[2].trim(),
          });
        }
      }
    }

    return fates.slice(0, 6);
  });

  /** Key moments */
  const keyMoments = computed((): KeyMoment[] => {
    if (!spoilerText.value) return [];
    const section = extractSection(spoilerText.value, "Key Moments");
    if (!section) return [];

    const moments: KeyMoment[] = [];
    const lines = section.split("\n").filter((l) => l.trim().startsWith("-"));

    for (const line of lines) {
      // Match: - **Title** — Description
      const match = line.match(/\*\*\[?(.+?)\]?\*\*\s*[-—]\s*(.+)/);
      if (match) {
        moments.push({
          title: match[1].trim(),
          description: match[2].trim(),
        });
      }
    }

    return moments.slice(0, 6);
  });

  /** Interpretation sections */
  const interpretations = computed((): InterpretationSection[] => {
    if (!spoilerText.value) return [];
    const mainSection = extractSection(
      spoilerText.value,
      "What It Really Means"
    );
    if (!mainSection) return [];

    const subSections: InterpretationSection[] = [];
    const subDefs = [
      "Symbolism",
      "Hidden Clues",
      "Fan Theories",
      "Unanswered Questions",
    ];

    for (const sub of subDefs) {
      const escaped = sub.replace(/[.*+?^${}()|[\]\\]/g, "\\$&");
      const regex = new RegExp(
        `###\\s*${escaped}\\s*\\n([\\s\\S]*?)(?=\\n###\\s|$)`,
        "i"
      );
      const match = mainSection.match(regex);
      if (match) {
        subSections.push({
          title: sub,
          content: match[1].trim(),
        });
      }
    }

    return subSections;
  });

  /** Check if spoiler has structured content */
  const hasStructuredContent = computed(() => {
    return storySections.value.length > 0;
  });

  return {
    movieOverview,
    storySections,
    characterFates,
    keyMoments,
    interpretations,
    hasStructuredContent,
  };
}
