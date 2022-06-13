const path = require('path')
require('dotenv').config()
module.exports = {
  env: {
    NEXT_PUBLIC_API_URL: process.env.NEXT_PUBLIC_API_URL,
    NEXT_PUBLIC_FIREBASE_API_KEY: process.env.NEXT_PUBLIC_FIREBASE_API_KEY,
    NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN: process.env.NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN,
    NEXT_PUBLIC_FIREBASE_PROJECT_ID: process.env.NEXT_PUBLIC_FIREBASE_PROJECT_ID,
    NEXT_PUBLIC_FIREBASE_STORAGE_BUCKET: process.env.NEXT_PUBLIC_FIREBASE_STORAGE_BUCKET,
    NEXT_PUBLIC_FIREBASE_MESSAGING_SENDER_ID: process.env.NEXT_PUBLIC_FIREBASE_MESSAGING_SENDER_ID,
    NEXT_PUBLIC_FIREBASE_APP_ID: process.env.NEXT_PUBLIC_FIREBASE_APP_ID,
    NEXT_PUBLIC_FIREBASE_TENANT_ID: process.env.NEXT_PUBLIC_FIREBASE_TENANT_ID
  },
  webpack: config => {
    config.plugins = config.plugins || []

    config.plugins = [...config.plugins]

    config.resolve.alias['src'] = path.join(__dirname, 'src')
    config.resolve.alias['modules'] = path.join(__dirname, '/src/modules')
    config.resolve.alias['pages'] = path.join(__dirname, '/src/pages')
    config.resolve.alias['components'] = path.join(__dirname, '/src/components')
    config.resolve.alias['states'] = path.join(__dirname, '/src/states')
    return config
  },
  distDir: 'dist/src/.next'
}
