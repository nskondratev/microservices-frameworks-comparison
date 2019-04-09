const crypto = require('crypto')

module.exports = {
  name: 'hash',

  actions: {
    sha256 (ctx) {
      const hash = crypto.createHash('sha256')
      hash.update(ctx.params.string)
      return { hash: hash.digest('hex') }
    }
  }
}
