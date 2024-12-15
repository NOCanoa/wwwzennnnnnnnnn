const fs = require('fs');
const path = require('path');

// Ensure the flag is passed as an argument
const flag = process.argv[2];
if (!flag) {
    console.error('Please provide a flag.');
    process.exit(1);
}

// Define the new site URL based on the flag
const newSiteURL = `https://zen-astro-test.b-cdn.net/${flag}`;

// Path to the astro.config.mjs file
const configPath = path.join(__dirname, 'astro.config.mjs');

// Read the existing config file
let config = fs.readFileSync(configPath, 'utf-8');

// Use a regular expression to find and replace the site URL
const updatedConfig = config.replace(/site: '.*?'/, `site: '${newSiteURL}'`);

// Write the updated config back to the file
fs.writeFileSync(configPath, updatedConfig);

console.log(`Updated site URL to ${newSiteURL}`);
