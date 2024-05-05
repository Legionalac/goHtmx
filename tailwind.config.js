/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["web/templates/*.templ"],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: ["light", "dark", "black"],
  },
  plugins: [
    require('daisyui'),
  ],
}

