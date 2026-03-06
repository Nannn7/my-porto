<script setup>
import Navbar from './components/Navbar.vue'
import Footer from './components/Footer.vue'
</script>

<template>
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
}
</style>
