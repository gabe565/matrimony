export const toSlug = (str, separator = "-") =>
  str
    .trim()
    .toLowerCase()
    .replace(/[^a-z0-9 -]/g, "")
    .replace(/\s+/g, separator)
    .replace(/-+/g, separator);
