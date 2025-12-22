/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./internal/ui/**/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/typography'),
  ],
};
