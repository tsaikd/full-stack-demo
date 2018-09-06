/*
custom error handler with backend:
  Backend.API_FUNCTION_CALL({ PARAMETERS }, { disableErrorHandler: true })
    .catch((err) => { ... })

example:
  Backend.Version({ disableErrorHandler: true })
    .catch(Backend.util.ErrorHandler)
 */

import axios from 'axios'
import appconst from '@/utils/appconst.js'

export default class BackendBase {
  util = new BackendUtil()
  errorHandler = this.util.ErrorHandler

  do (config) {
    config.url = '/api/1' + (config.url || '')

    return axios.request(config)
      .then((res) => {
        return res.data
      })
      .catch((err) => {
        if (!config.disableErrorHandler) {
          return this.errorHandler(err, config)
        }
        throw err
      })
  }

  getAPIURL (resourcePath = '', params = {}) {
    return '/api/1' + resourcePath + '?' + Object.keys(params).map(function (key) {
      return [key, params[key]].map(encodeURIComponent).join('=')
    }).join('&')
  }
}

class BackendUtil {
  ErrorHandler (err, config) {
    if (err && err.response && err.response.data) {
      if (err.response.data.error) {
        console.error(err.response.data.error, err)
        throw err
      }
      console.error(err.response.data, err)
      throw err
    }
    console.error(err)
    throw err
  }
}
