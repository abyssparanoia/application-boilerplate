import express from 'express'
import next from 'next'

const dev = process.env.NODE_ENV !== 'production'
const PORT = process.env.PORT || 3000

const app = next({ dir: '.', dev })
const handle = app.getRequestHandler()

export const appFactory = async () => {
  await app.prepare()
  const server = express()

  // nextjs routing
  server.get('*', (req, res) => handle(req, res))
  server.post('*', (req, res) => handle(req, res))
  server.put('*', (req, res) => handle(req, res))
  server.patch('*', (req, res) => handle(req, res))
  server.delete('*', (req, res) => handle(req, res))

  return server
}

const run = async () => {
  try {
    const server = await appFactory()
    server.listen(PORT, () => {
      console.log(`> Ready on http://localhost:${PORT}`)
    })
  } catch (err) {
    console.error(err)
  }
}

run()
