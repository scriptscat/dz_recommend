import React, { useState } from 'react'
import EveryArticleItem from '../components/EveryArticleItem'
import { Select, Button, Input, Row, Col,  Space, Modal } from 'antd'
import youZhongImage from '../../src/static/image/youzhong.png'
import './YouZhoneSearch.css'
import Search from '../../src/api/search' 

const { Option } = Select;
function YouZhongSearch() {
  // 搜索输入的input内容
  const [searchInputContent, setSearchInputContent] = useState('')
  // 输入用户名的input显示隐藏
  const [isDisplayInputWithUser, setIsDisplayInputWithUser] = useState(false)
  // 搜索到的文章内容
  const [searchContentData, setSearcContentData] = useState([])
  // 搜索到的高亮部分内容
  const [searchHighlightContents, setSearchHighlightContents] = useState([])
  // 搜索到的内容数量
  const [contentTotal, setContentTotal] = useState(0)

  // 选择搜索的方式
  const selectedMethods = (method) => {
    if (method === 'searchContent') {
      setIsDisplayInputWithUser(false)
    } else if (method === 'searchContentWithUser') {
      setIsDisplayInputWithUser(true)
    }
  }

  // 调用搜索接口
  const search = () => {
    Search.searchContent().then(
      res => {
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

  // 搜索内容改变
  const searchInputChange = (e) => {
    console.log(e.target.value)
    setSearchInputContent(e.target.value)
  }
  return (
    <div style={{marginLeft: '15px'}}>
      <div>
        <div style={{ display: 'flex', marginBottom: '20px' }}>
          <div style={{ marginRight: '10px' }}>
            <img src={youZhongImage} />
          </div>
          <div className='searchBox' style={{ width: '100%'}}>
            <Input.Group compact>
              <Select defaultValue="searchContent" style={{ width: '130px'}} onChange={selectedMethods}>
                <Option value="searchContent">普通搜索</Option>
                <Option value="searchContentWithUser">搜索指定用户</Option>
              </Select>
              <Input style={{ width: '200px', display: `${isDisplayInputWithUser ? '' : 'none'}` }} defaultValue=""  placeholder='请输入用户名'/>
              <Input style={{ width: '500px' }} onChange={searchInputChange}/>
              <Button type="primary" onClick={search}>查询</Button>
            </Input.Group>
          </div>
        </div>
        <div className='searchResult'><h2 style={{ fontSize: '14px'}}>结果: 找到<span>{`" ${searchInputContent} "`}</span>相关内容<span>{contentTotal} 个</span></h2></div>
        {searchContentData.map((divContent, index) => {
          return <EveryArticleItem key={index} divContent={divContent} highlightContents={searchHighlightContents}/>
        })}
      </div>
    </div>
  );
}

export default YouZhongSearch