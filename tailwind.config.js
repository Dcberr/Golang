/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
      "./templates/**/*.{html,js,templ,go}",
      "./templates/common/**/*.{html,js,templ,go}",
    ],
    theme: {
      extend: {},
      spacing: {
        '128': "32rem"
      },
      fontFamily: {
        sans: ["Quicksand"],
      },
    },
    plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography")],
  };
  