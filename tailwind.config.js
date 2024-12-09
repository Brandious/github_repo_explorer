/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.templ}", "./**/*.templ"],
  darkMode: 'class', // Enable dark mode
  theme: {
    extend: {
      colors: {
        // Electric blue as primary
        primary: {
          50: '#E6F8FF',
          100: '#B3ECFF',
          200: '#80DFFF',
          300: '#4DD2FF',
          400: '#1AC5FF',
          500: '#00A3FF', // primary
          600: '#0082CC',
          700: '#006199',
          800: '#004166',
          900: '#002033',
        },
        // Purple accent
        secondary: {
          50: '#F5E6FF',
          100: '#E5B3FF',
          200: '#D580FF',
          300: '#C54DFF',
          400: '#B51AFF',
          500: '#A500FF',
          600: '#8400CC',
          700: '#630099',
          800: '#420066',
          900: '#210033',
        },
        // Modern dark theme colors
        dark: {
          // Main backgrounds
          background: '#0F172A',    // Deep blue-black
          surface: '#1E293B',       // Lighter blue-black
          card: '#334155',          // Card background

          // Text colors
          text: {
            primary: '#F1F5F9',     // Almost white
            secondary: '#94A3B8',   // Muted gray
            accent: '#00A3FF',      // Electric blue
          },

          // UI elements
          border: '#2D3748',        // Subtle borders
          hover: '#2C5282',         // Hover state
          active: '#2B6CB0',        // Active state

          // Accent colors
          accent: {
            cyan: '#06B6D4',
            purple: '#A500FF',
            green: '#10B981',
            red: '#EF4444',
          },

          // Gradient colors
          gradient: {
            start: '#00A3FF',
            end: '#A500FF',
          }
        },
      },
      fontFamily: {
        sans: ['Inter', 'system-ui', 'sans-serif'],
        heading: ['Poppins', 'sans-serif'],
        mono: ['JetBrains Mono', 'monospace'], // Good for code
      },
      spacing: {
        '128': '32rem',
        '144': '36rem',
      },
      borderRadius: {
        '4xl': '2rem',
      },
      // Add glass effect utilities
      backdropBlur: {
        xs: '2px',
      },
      // Custom box shadows
      boxShadow: {
        'neon': '0 0 5px theme(colors.primary.500), 0 0 20px theme(colors.primary.500)',
        'neon-lg': '0 0 10px theme(colors.primary.500), 0 0 40px theme(colors.primary.500)',
        'neon-purple': '0 0 5px theme(colors.secondary.500), 0 0 20px theme(colors.secondary.500)',
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
    // Add custom utilities
    function ({ addUtilities }) {
      const newUtilities = {
        '.glass': {
          backgroundColor: 'rgba(255, 255, 255, 0.05)',
          backdropFilter: 'blur(10px)',
          border: '1px solid rgba(255, 255, 255, 0.1)',
        },
        '.glass-dark': {
          backgroundColor: 'rgba(15, 23, 42, 0.75)',
          backdropFilter: 'blur(10px)',
          border: '1px solid rgba(255, 255, 255, 0.1)',
        },
      }
      addUtilities(newUtilities)
    },
  ],
}