function getErrFac (res) {
  if (!res) {
    return res
  }
  if (res.errfac) {
    return res.errfac
  }
  if (res.response) {
    return getErrFac(res.response)
  }
  if (res.data) {
    return getErrFac(res.data)
  }
  return null
}

export default class BackendError {
  constructor (err) {
    this.err = err
  }

  match (res) {
    const errfac = getErrFac(res)
    return errfac && (this.err in errfac)
  }
}
