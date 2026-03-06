<<<<<<< HEAD
<script setup>
const navItems = [
  { name: 'Home', to: '/' },
  { name: 'About', to: '/about' },
  { name: 'Projects', to: '/projects' },
  { name: 'Contact', to: '/contact' },
]
</script>

<template>
  <header class="navbar">
    <div class="navbar__brand">My Porto</div>
    <nav class="navbar__nav">
      <RouterLink
        v-for="item in navItems"
        :key="item.to"
        :to="item.to"
        class="navbar__link"
        active-class="navbar__link--active"
      >
        {{ item.name }}
      </RouterLink>
=======
<template>
  <header class="navbar">
    <nav>
      <RouterLink to="/">Home</RouterLink>
      <RouterLink to="/about">About</RouterLink>
      <RouterLink to="/projects">Projects</RouterLink>
      <RouterLink to="/contact">Contact</RouterLink>
>>>>>>> origin/main
    </nav>
  </header>
</template>

<style scoped>
.navbar {
<<<<<<< HEAD
  position: sticky;
  top: 0;
  z-index: 10;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1rem 1.5rem;
  background: #0f172a;
  color: #fff;
}
.navbar__brand { font-weight: 700; }
.navbar__nav { display: flex; gap: 1rem; }
.navbar__link { color: #cbd5e1; text-decoration: none; }
.navbar__link--active { color: #fff; border-bottom: 2px solid #38bdf8; }
=======
  padding: 16px;
  border-bottom: 1px solid #ddd;
}

nav {
  display: flex;
  gap: 12px;
}
>>>>>>> origin/main
</style>
