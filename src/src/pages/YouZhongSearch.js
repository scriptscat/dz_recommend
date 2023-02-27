import React, { useState, useEffect } from 'react'
import EveryArticleItem from '../components/EveryArticleItem'
import { Select, Button, Input, Row, Col,  Space, Modal, message } from 'antd'
import youZhongImage from '../../src/static/image/youzhong.png'
import './YouZhoneSearch.css'
import Search from '../../src/api/search' 

const { Option } = Select;
function YouZhongSearch() {
  // 搜索输入的input内容
  const [searchInputContent, setSearchInputContent] = useState('')
  // 搜索到的用户名结果
  const [searchInputResult, setSearchInputResult] = useState('')
  // 选择指定的用户
  const [selectedUserName, setSelectedUserName] = useState('')
  // 输入用户名的input显示隐藏
  const [isDisplayInputWithUser, setIsDisplayInputWithUser] = useState(false)
  // 搜索到的文章内容
  const [searchContentData, setSearcContentData] = useState([])
  // 搜索到的高亮部分内容
  const [searchHighlightContents, setSearchHighlightContents] = useState([])
  // 搜索到的内容数量
  const [contentTotal, setContentTotal] = useState(0)
  // 显示搜索结果
  const [displayResultDiv, setDisplayResultDiv] = useState('none')
  // 显示搜索用户结果
  const [userNameListResult, setUserNameListResult] = useState([])

  useEffect(() => {
    // 获取当前 URL
    const url = window.location.href;
    // 解析 URL，获取其中的参数
    const urlParams = new URLSearchParams(url.split("?")[1]);
    // 获取 keyword 参数的值
    const keyword = urlParams.get("keyword");
    console.log(urlParams)
    console.log(keyword)
    if(keyword) {
      setSearchInputContent(keyword)
      setTimeout(() => {
        search()
      }, 0);
    }
  }, [])

  // 选择搜索的方式
  const selectedMethods = (method) => {
    if (method === 'searchContent') {
      setIsDisplayInputWithUser(false)
    } else if (method === 'searchContentWithUser') {
      setIsDisplayInputWithUser(true)
    }
  }

  // 搜索内容改变
  const searchInputChange = (e) => {
    console.log(e.target.value)
    setSearchInputContent(e.target.value)
  }

  // 搜索用户
  const onSearchUser = (userName) => {
    console.log(userName)
    Search.searchUserName(userName).then(
      res => {
        console.log(res)
        if (res.data.length > 0) {
          const userNameListOptions = res.data.map((item, index) => {
            return {
              value: item.uid, 
              label: item.username
            }
          })
          console.log(userNameListOptions)
          setUserNameListResult(userNameListOptions)

        }
      }
    ).catch(
      err => {
        console.log(err)
      }
    )
  }

  // 选择指定用户
  const onChangeUserName = (userName) => {
    console.log(userName)
    setSelectedUserName(userName)
  }

  // 调用搜索接口
  const search = () => {
    if (!searchInputContent) {
      message.warning('请输入查询的内容')
      return
    }
    Search.searchContent(searchInputContent, selectedUserName).then(
      res => {
        setDisplayResultDiv('')
        setSearchInputResult(searchInputContent)
        setSearcContentData(res.data)
        setSearchHighlightContents(res.analyze)
        setContentTotal(res.total)
      }
    ).catch(
      err => {
        console.log(err)
      }
    )
  }

  return (
    <div style={{ marginLeft: '15px' }}>
      <div>
        <div style={{ display: 'flex', marginBottom: '20px' }}>
          <div style={{ marginRight: '10px' }}>
            <img src={youZhongImage} />
          </div>
          <div className="searchBox" style={{ width: '100%' }}>
            <Input.Group compact>
              <Select
                defaultValue="searchContent"
                style={{ width: '130px' }}
                onChange={selectedMethods}
              >
                <Option value="searchContent">普通搜索</Option>
                <Option value="searchContentWithUser">搜索指定用户</Option>
              </Select>
              {/* <Input
                style={{
                  width: '200px',
                  display: `${isDisplayInputWithUser ? '' : 'none'}`,
                }}
                defaultValue=""
                placeholder="请输入用户名"
              /> */}
              <Select
                style={{
                  width: '200px',
                  display: `${isDisplayInputWithUser ? '' : 'none'}`,
                }}
                showSearch
                placeholder="请输入用户名"
                optionFilterProp="children"
                onChange={onChangeUserName}
                onSearch={onSearchUser}
                filterOption={(input, option) =>
                  (option?.label ?? '')
                    .toLowerCase()
                    .includes(input.toLowerCase())
                }
                options={userNameListResult}
              />
              <Input style={{ width: '500px' }} onChange={searchInputChange} onPressEnter={search}/>
              <Button type="primary" onClick={search}>
                查询
              </Button>
            </Input.Group>
          </div>
        </div>
        <div className="searchResult" style={{ display: displayResultDiv }}>
          <h2 style={{ fontSize: '14px' }}>
            结果: 找到<span>{`" ${searchInputResult} "`}</span>相关内容
            <span>{contentTotal} 个</span>
          </h2>
        </div>
        {searchContentData.map((divContent, index) => {
          return (
            <EveryArticleItem
              key={index}
              divContent={divContent}
              highlightContents={searchHighlightContents}
            />
          );
        })}
      </div>
    </div>
  );
}

export default YouZhongSearch