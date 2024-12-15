// @ts-check
import { defineConfig } from 'astro/config';
import tailwind from '@astrojs/tailwind';
import react from '@astrojs/react';

const prId = process.env.PR_ID || 'default'; // Use 'default' if PR_ID is not set
console.log(`PR_ID: ${prId}`);
console.log(`Site URL: https://prs-zen-browser-site.b-cdn.net/${prId}`);


// https://astro.build/config
export default defineConfig({
    site: `https://prs-zen-browser-site.b-cdn.net/${prId}`,
    base: '/${prId}',
    integrations: [tailwind(), react()],
});