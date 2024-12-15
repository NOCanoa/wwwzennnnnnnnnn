// @ts-check
import { defineConfig } from 'astro/config';
import tailwind from '@astrojs/tailwind';

import react from '@astrojs/react';

import { defineConfig } from 'astro/config';

const prId = process.env.PR_ID || 'default';  // Use 'default' if PR_ID is not set


// https://astro.build/config
export default defineConfig({
    integrations: [tailwind(), react()],
    site: `https://prs-zen-browser-site.b-cdn.net/${prId}`,
});