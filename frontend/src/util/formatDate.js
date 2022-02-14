export const FullDate = {
  weekday: "long",
  month: "long",
  day: "numeric",
  year: "numeric",
};

export const FullDateTime = {
  ...FullDate,
  hour: "numeric",
  minute: "numeric",
};

export const formatDate = (str, format = FullDateTime) => {
  if (!str) {
    return null;
  }
  const date = new Date(str);
  return date.toLocaleString("en-US", format);
};
