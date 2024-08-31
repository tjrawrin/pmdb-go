import tailwindcss from 'tailwindcss'
import { defineConfig } from 'vite'

export default defineConfig({
	css: {
		postcss: {
			plugins: [tailwindcss()],
		},
	},
	publicDir: 'web/public',
	build: {
		manifest: true,
		outDir: 'web/dist',
		rollupOptions: {
			input: ['web/assets/script.js', 'web/assets/style.css'],
			output: {
				entryFileNames: 'assets/[name]-[hash].js',
				chunkFileNames: 'assets/[name]-[hash].js',
				assetFileNames: 'assets/[name]-[hash].[ext]',
			},
		},
	},
})
