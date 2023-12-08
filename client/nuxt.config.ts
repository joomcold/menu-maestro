// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: false,
  devtools: { enabled: true },
  devServer: { port: 5173 },
  modules: ['@nuxt/ui', '@pinia/nuxt', '@nuxt/content', '@vee-validate/nuxt']
})
