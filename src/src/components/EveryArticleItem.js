import React from 'react'
import './EveryArticleItem.css'
import DOMPurify from "dompurify";
import parse from "html-react-parser";

function EveryArticleItem(props) {
  const { divContent, highlightContents } = props 
  console.log(divContent, highlightContents)
  const { title, content, author } = divContent
  console.log(props)

  // 对标题和内容展示高亮内容
  const showHighLightsAboutTitleOrContent = (content, highlightContents) => {
    console.log(content, highlightContents)
    highlightContents.map((highlightContent, index) => {
      const _reg = new RegExp(highlightContent['token'], 'g')
      console.log(_reg.exec(content))
      content = content.replace(_reg, `<span class='hightLightContent'>${highlightContent['token']}</span>`)
    })
    console.log(content)
    return DOMPurify.sanitize(content)
  }


  return (
    <div style={{ marginBottom: '15px'}}>
      <div style={{ fontSize: '16px', marginBottom: '3px'}}>{parse(showHighLightsAboutTitleOrContent(title, highlightContents))}</div>
      <div className="replayAndSearchCounts">0个回复 - 56次查看</div>
      <div style={{ fontSize: '12px', marginBottom: '3px'}}>{parse(showHighLightsAboutTitleOrContent(content, highlightContents))}</div>
      <div className="publishTimeAndAuthorAndTheme">
        <span className='publishTime'>2023-2-9 00:36</span> - <span className="author">{author}</span> -{' '}
        <span className="theme">网页开发</span>
      </div>
    </div>
  );
}

export default EveryArticleItem