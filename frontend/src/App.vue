<script setup>
import Navbar from './components/Navbar.vue'
import Footer from './components/Footer.vue'
</script>

<template>
  <div class="app-shell">
    <div class="orb orb--left" aria-hidden="true"></div>
    <div class="orb orb--right" aria-hidden="true"></div>

    <Navbar />

    <main class="main-content">
      <RouterView v-slot="{ Component, route }">
        <Transition name="page" mode="out-in">
          <component :is="Component" :key="route.fullPath" />
        </Transition>
      </RouterView>
    </main>

    <Footer class="reveal reveal--delay-1" />
  </div>
</template>

<style scoped>
.app-shell {
  min-height: 100vh;
  position: relative;
}

.main-content {
  width: min(1120px, 94vw);
  margin: 1.4rem auto 0;
  position: relative;
  z-index: 1;
}

.orb {
  position: fixed;
  width: 280px;
  aspect-ratio: 1;
  border-radius: 999px;
  filter: blur(60px);
  opacity: 0.35;
  z-index: 0;
  pointer-events: none;
  animation: orb-float 8s ease-in-out infinite alternate;
}

.orb--left {
  left: -90px;
  top: 140px;
  background: var(--brand-soft);
}

.orb--right {
  right: -70px;
  bottom: 110px;
  background: var(--accent-soft);
  animation-delay: 0.6s;
}

@keyframes orb-float {
  from {
    transform: translateY(0);
  }

  to {
    transform: translateY(-16px);
  }
}

@media (prefers-reduced-motion: reduce) {
  .orb {
    animation: none;
  }
}
</style>
