<template>
  <header class="navbar">
    <nav>
      <RouterLink to="/">Home</RouterLink>
      <RouterLink to="/about">About</RouterLink>
      <RouterLink to="/projects">Projects</RouterLink>
      <RouterLink to="/contact">Contact</RouterLink>
    </nav>
  </header>
</template>

<style scoped>
.navbar {
  padding: 16px;
  border-bottom: 1px solid #ddd;
}

nav {
  display: flex;
  gap: 12px;
}
</style>
