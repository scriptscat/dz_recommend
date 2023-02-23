import request from "./request";

export default {
  // 查询内容
  searchContent : () => {
    return request('get', 'http://localhost/plugin.php?id=codfrm_recommend:search&operation=search&keyword=%E9%98%BF%E6%96%AF%E9%A1%BF')
  }
}