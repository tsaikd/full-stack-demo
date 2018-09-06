import BackendBase from './BackendBase'
import BackendError from './BackendError'

class BackendAPI extends BackendBase {
  ErrInvalidParam = new BackendError('ErrInvalidParam')

  // Get coin history
  GetCoinHistory ({ symbol }, config) {
    const params = Object.assign({
      'symbol': symbol
    }, (config || {}).params)
    return this.do(Object.assign({
      'method': 'get',
      'url': '/get/coinHistory'
    }, config, { params }))
  }

  // Get version info
  Version (config) {
    const params = Object.assign({
    }, (config || {}).params)
    return this.do(Object.assign({
      'method': 'get',
      'url': '/version'
    }, config, { params }))
  }
}

export default BackendAPI
