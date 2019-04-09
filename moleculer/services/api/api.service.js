const fastify = require('fastify')

module.exports = {
  name: 'api',

  app: null,

  // dependencies: ['hash'],

  created () {
    this.logger.info('Created hook in api service')

    this.app = fastify({ logger: false })

    this.app.get('/hash/sha256', async (request, reply) => {
      const { hash } = await this.broker.call('hash.sha256', { string: request.query.string })
      return { hash }
    })
  },

  async started () {
    this.logger.info('Started hook in api service')
    try {
      await this.app.listen(8081)
      this.logger.info(`Fastify app start listening on ${this.app.server.address().port} port`)
    } catch (err) {
      this.logger.error(err)
    }
  },

  async stopped () {
    return this.app.close()
  }
}
