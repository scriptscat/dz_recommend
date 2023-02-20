<?php

if (!defined('IN_DISCUZ')) {
    exit('Access Denied');
}

include "vendor/autoload.php";
require "lib/es.php";


switch ($_GET["operation"]) {
    case "search":
        search();
        break;
    default:
}

function search()
{
    global $_G;
    $keyword = $_GET["keyword"];
    $page = $_GET["page"] ?? 1;
    if (empty($keyword)) {
        showmessage("请输入关键词");
        return;
    }
    if ($page < 1) {
        $page = 1;
    } else if ($page > 100) {
        $page = 100;
    }
    // 使用缓存限流
    $cacheKey = "codfrm_recommend:search:{$_G['uid']}";
    $ok = loadcache($cacheKey);
    if ($ok && $_G['cache'][$cacheKey]['time'] > time() - 5) {
        showmessage("操作过于频繁，请稍后再试");
        return;
    }
    // 通过title和content,es搜索,并获取分词结果
    $resp = Es::getClient()->search([
        "index" => "dz.forum_thread",
        "body" => [
            "query" => [
                "multi_match" => [
                    "query" => $keyword,
                    "fields" => ["title", "content"],
                ],
            ],
        ],
        "size" => 20,
        "from" => ($page - 1) * 20,
    ]);
    // 获取分词结果

    $analyzeResp = Es::getClient()->indices()->analyze([
        "index" => "dz.forum_thread",
        "body" => [
            "text" => $keyword,
            "analyzer" => "ik_max_word",
        ],
    ]);

    $list = $resp['hits']['hits'];
    foreach ($list as $k => &$item) {
        $item = $item['_source'];
        // 查询帖子作者
        $item['author'] = C::t('common_member')->fetch($item['authorid'])["username"];
        // 截取内容
        $item['content'] = cutstr(strip_tags($item['content']), 200);
    }

    // 写缓存
    savecache($cacheKey, [
        "time" => time(),
    ]);

    echo json_encode([
        "code" => 0,
        "data" => $list,
        "total" => $resp['hits']['total']['value'],
        "analyze" => $analyzeResp['tokens'],
    ], JSON_UNESCAPED_UNICODE);
}