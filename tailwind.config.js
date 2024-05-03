/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/*.tmpl"],
  theme: {
    extend: {
      fontFamily: {
        sans: ["Poppins", "sans-serif"],
        mono: ["Office Code Pro", "monospace"],
      }
    },
  },
  plugins: [],
}

