<script setup>
import Navbar from './components/Navbar.vue'
import Footer from './components/Footer.vue'
</script>

<template>
  <div class="relative min-h-screen">
    <div
      class="orb-float pointer-events-none fixed left-[-90px] top-[140px] z-0 aspect-square w-[280px] rounded-full bg-brand-soft opacity-35 blur-[60px] motion-reduce:animate-none"
      aria-hidden="true"
    ></div>
    <div
      class="orb-float pointer-events-none fixed bottom-[110px] right-[-70px] z-0 aspect-square w-[280px] rounded-full bg-accent-soft opacity-35 blur-[60px] [animation-delay:0.6s] motion-reduce:animate-none"
      aria-hidden="true"
    ></div>

    <Navbar />

    <main class="relative z-[1] mx-auto mt-[1.4rem] w-[min(1120px,94vw)]">
      <RouterView v-slot="{ Component, route }">
        <Transition name="page" mode="out-in">
          <component :is="Component" :key="route.fullPath" />
        </Transition>
      </RouterView>
    </main>

    <Footer class="reveal reveal--delay-1" />
  </div>
</template>
