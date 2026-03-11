<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import Navbar from './components/Navbar.vue'
import Footer from './components/Footer.vue'

const route = useRoute()
const isAdminLayout = computed(() => route.meta.layout === 'admin')
</script>

<template>
  <div class="relative min-h-screen">
    <template v-if="!isAdminLayout">
      <div
        class="orb-float pointer-events-none fixed left-[-90px] top-[140px] z-0 aspect-square w-[280px] rounded-full bg-brand-soft opacity-35 blur-[60px] motion-reduce:animate-none"
        aria-hidden="true"
      ></div>
      <div
        class="orb-float pointer-events-none fixed bottom-[110px] right-[-70px] z-0 aspect-square w-[280px] rounded-full bg-accent-soft opacity-35 blur-[60px] [animation-delay:0.6s] motion-reduce:animate-none"
        aria-hidden="true"
      ></div>

      <Navbar />
    </template>

    <main :class="isAdminLayout ? 'relative z-[1] w-full px-4 py-6 md:px-8 lg:px-12' : 'relative z-[1] mt-[1.4rem] w-full px-4 md:px-8 lg:px-12'">
      <RouterView />
    </main>

    <Footer v-if="!isAdminLayout" class="reveal reveal--delay-1" />
  </div>
</template>
