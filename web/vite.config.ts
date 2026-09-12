import react, { reactCompilerPreset } from '@vitejs/plugin-react'
import babel from '@rolldown/plugin-babel'
import { defineConfig } from 'vitest/config'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), babel({ presets: [reactCompilerPreset()] })],
  server: {
    proxy: {
      // /todos このパスで始まるリクエストだけ転送
      '/todos': {
        target: 'http://api:8080', // 転送先 Compose 内のAPIコンテナ api:8080
        changeOrigin: true, // 転送先のHostヘッダをapi:8080に変更
      },
    },
  },
  test: {
    environment: 'jsdom',
  },
})
