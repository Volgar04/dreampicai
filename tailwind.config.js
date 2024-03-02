// tailwind.config.js

module.exports = {
  content: ["./view/**/*.templ", "./**/*.templ"],
  safelist: [],
  theme: {},
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["retro"]
  }
}
