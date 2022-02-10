export const throttle = (func, delay, { leading } = {}) => {
  let timerId;
  let firstRun = true;

  return (...args) => {
    if (timerId) {
      return;
    }
    if (leading && firstRun) {
      firstRun = false;
      func(...args);
    }

    timerId = setTimeout(() => {
      func(...args);
      timerId = null;
    }, delay);
  };
};
