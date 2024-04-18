/* eslint-env node */
/** @type {import('tailwindcss').Config} */

module.exports = {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {
      colors: {
        primary: {
          50: '#E8EAF2',
          100: '#D5D8E7',
          200: '#AAB2CF',
          300: '#808BB7',
          400: '#566499',
          500: '#3F4970',
          600: '#323958',
          700: '#272D45',
          800: '#1A1E2E',
          900: '#0D0F17',
        },
        accent: {
          50: '#FFFAFA',
          100: '#FFF5F5',
          200: '#FFE5E6',
          300: '#FFDBDC',
          400: '#FFCCCE',
          500: '#FFB9BB',
          600: '#FF9EA1',
          700: '#FF757A',
          800: '#FF333A',
          900: '#C70007',
        },
      },
      container: {
        screens: {
          xl: '1236px',
        },
      },
      fontFamily: {
        script: ['"Dancing Script"', 'cursive'],
      },
    },
  },
  plugins: [require('@tailwindcss/typography'), require('@tailwindcss/forms')],
}
