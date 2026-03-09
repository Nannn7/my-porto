/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {
      colors: {
        bg: 'var(--bg)',
        'bg-overlay-1': 'var(--bg-overlay-1)',
        'bg-overlay-2': 'var(--bg-overlay-2)',
        'bg-elevated': 'var(--bg-elevated)',
        surface: 'var(--surface)',
        'surface-soft': 'var(--surface-soft)',
        ink: 'var(--ink)',
        heading: 'var(--heading)',
        muted: 'var(--muted)',
        brand: 'var(--brand)',
        'brand-deep': 'var(--brand-deep)',
        'brand-soft': 'var(--brand-soft)',
        accent: 'var(--accent)',
        'accent-soft': 'var(--accent-soft)',
        border: 'var(--border)',
        link: 'var(--link)',
        focus: 'var(--focus)',
      },
      borderRadius: {
        lgx: 'var(--radius-lg)',
        mdx: 'var(--radius-md)',
      },
      boxShadow: {
        soft: 'var(--shadow)',
      },
      fontFamily: {
        sans: ['Manrope', 'Segoe UI', 'Tahoma', 'sans-serif'],
        display: ['Sora', 'Trebuchet MS', 'sans-serif'],
      },
    },
  },
  plugins: [],
}
