<script setup>
import Navbar from './components/Navbar.vue'
import Footer from './components/Footer.vue'
</script>

<template>
<<<<<<< HEAD
  <div class="app-shell">
    <Navbar />
    <main class="main-content">
      <RouterView />
    </main>
    <Footer />
  </div>
</template>

<style scoped>
.app-shell {
  min-height: 100vh;
  background: #f8fafc;
}
.main-content {
  max-width: 960px;
  margin: 0 auto;
  padding: 1.5rem;
=======
  <Navbar />
  <main class="container">
    <RouterView />
  </main>
  <Footer />
</template>

<style scoped>
.container {
  max-width: 900px;
  margin: 0 auto;
  padding: 16px;
>>>>>>> origin/main
}
</style>
