import request from "./request";
let YOUZHONG_URL = '/plugin.php'
export default {
  // 查询内容
  searchContent : (keyword, uid = '') => {
    return request('get', `${YOUZHONG_URL}?id=codfrm_recommend:search&operation=search&keyword=${keyword}&uid=${uid}`)
  },
  searchUserName : (userName) => {
    return request('get', `${YOUZHONG_URL}?id=codfrm_recommend:search&operation=user&username=${userName}`)
  }
}