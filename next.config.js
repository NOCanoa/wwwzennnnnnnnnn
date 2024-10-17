const { PHASE_DEVELOPMENT_SERVER } = require("next/constants");

/** @type {import('next').NextConfig} */
const nextConfig = (phase, { defaultConfig }) => {
	const defaultConfigWWW = {
		images: {
			remotePatterns: [
				{
					protocol: "https",
					hostname: "raw.githubusercontent.com",
				},
				{
					protocol: "https",
					hostname: "cdn.jsdelivr.net",
					port: "",
					pathname: "/gh/zen-browser/**",
				},
			],
			domains: ["localhost", "cdn.jsdelivr.net", "raw.githubusercontent.com"], // Allow images from jsDelivr
		},

		compiler: {
			styledComponents: true,
		},
	};
	return {
		...defaultConfigWWW,
		// production only config options here
		output: "export",
	};
};