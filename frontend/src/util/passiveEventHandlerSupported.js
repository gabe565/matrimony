export const passiveEventHandlerSupported = () => {
  let passiveSupported = false;

  try {
    const options = {
      get passive() {
        // This function will be called when the browser
        //   attempts to access the passive property.
        passiveSupported = true;
        return false;
      },
    };

    window.addEventListener("test", null, options);
    window.removeEventListener("test", null, options);
  } catch (err) {
    passiveSupported = false;
  }
  return passiveSupported;
};
