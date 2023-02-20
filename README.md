# dz_recommend

油猴中文网用文章相似推荐

## 开发配置
启动docker后,在runtime/config/config_global.php文件后加入
```php
$_config['plugindeveloper'] = 2;
```

## es 索引模板

```
PUT _index_template/dz.forum_thread
{
  "template": {
    "mappings": {
      "properties": {
        "tid": {
          "type": "long"
        },
        "fid": {
          "type": "long"
        },
        "title": {
          "type": "text",
          "index": true,
          "analyzer": "ik_max_word",
          "eager_global_ordinals": false,
          "index_phrases": false,
          "norms": true,
          "fielddata": false,
          "store": false,
          "index_options": "positions",
          "search_analyzer": "ik_smart"
        },
        "content": {
          "type": "text",
          "index": true,
          "analyzer": "ik_max_word",
          "eager_global_ordinals": false,
          "index_phrases": false,
          "norms": true,
          "fielddata": false,
          "store": false,
          "index_options": "positions",
          "search_analyzer": "ik_smart"
        }
      }
    }
  },
  "index_patterns": [
    "dz.forum_thread"
  ]
}
```