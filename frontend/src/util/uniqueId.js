export const uniqueId = (n = 6) =>
  [...Array(n)].map(() => Math.random().toString(36)[2]).join("");
