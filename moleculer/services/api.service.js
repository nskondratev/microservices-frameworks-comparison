'use strict'

const fastify = require('fastify')

module.exports = {
  name: 'api',

  app: null,

  created () {
    this.app = fastify({ logger: false })

    this.app.get('/hash/sha256', request => this.broker.call('hash.sha256', { string: request.query.string }))
  },

  async started () {
    await this.app.listen(this.settings.port)
  },

  async stopped () {
    await this.app.close()
  },

  settings: {
    port: process.env.PORT || 3000
  }
}
