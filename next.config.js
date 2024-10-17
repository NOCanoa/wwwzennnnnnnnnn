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
    experimental: {
      serverActions: {
        // edit: updated to new key. Was previously `allowedForwardedHosts`
        allowedOrigins: ["localhost:3000", "get-zen.vercel.app"],
      },
    },
    compiler: {
      styledComponents: true,
    },
  };
};

module.exports = {
  ...nextConfig,
  output: "static",
  distDir: "out",
  
};