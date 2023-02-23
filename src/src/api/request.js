import axios from "axios"
// axios.defaults.withCredentials = true;
let config = {

}
const request = (type, url, params) => {
  config = {
    type,
    url
  }
  return new Promise((resolve, reject) => {
    return axios(config).then(res => {
      resolve(res.data)
      return
    }).catch(err => {
      reject(err.data)
      return
    })
  })
}

export default request 

