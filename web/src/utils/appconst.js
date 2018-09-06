const AppName = require('../../package.json').name
const IsDev = process.env.NODE_ENV !== 'production'

export default {
  AppName,
  IsDev
}
