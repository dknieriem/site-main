/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/*.go.tmpl"],
  darkMode: 'selector',
  theme: {
    extend: {
      fontFamily: {
        sans: ["Open Sans", "sans-serif"],
        mono: ["Office Code Pro", "monospace"],
      }
    },
  },
  plugins: [],
}

