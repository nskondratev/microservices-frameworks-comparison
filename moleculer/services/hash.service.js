const crypto = require('crypto')

module.exports = {
  name: 'hash',

  actions: {
    sha256: {
      cache: true,
      handler (ctx) {
        const hash = crypto.createHash('sha256')
        hash.update(ctx.params.string)
        return { hash: hash.digest('hex') }
      }
    }
  }
}
