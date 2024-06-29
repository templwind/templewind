/** @type {import('tailwindcss').Config} */
export default {
  content: ["./**/*.{html,js,templ}"],
  theme: {
    extend: {},
  },
  plugins: [
    require("@tailwindcss/typography"), 
    require('@tailwindcss/forms'),
    require('daisyui'),
  ],
}

