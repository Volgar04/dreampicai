// tailwind.config.js

module.exports = {
  content: ["./**/*.html", "./**/*.templ", "./**/*.go",],
  safelist: [],
  theme: {},
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["dark"]
  }
}
