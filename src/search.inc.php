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
    case "user":
        user();
        break;
    default:
}

function search()
{
    global $_G;
    $keyword = $_GET["keyword"];
    $page = $_GET["page"] ?? 1;
    $uid = $_GET["uid"] ?? 0;
    if (empty($keyword)) {
        echo json_encode([
            "code" => -1,
            "msg" => "请输入关键词"
        ], JSON_UNESCAPED_UNICODE);
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
    if ($ok && $_G['cache'][$cacheKey]['time'] > time() - 1) {
        echo json_encode([
            "code" => -2,
            "msg" => "操作过于频繁，请稍后再试"
        ], JSON_UNESCAPED_UNICODE);
        return;
    }
    // 通过title和content,es搜索,并获取分词结果
    $query = [
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
    ];
    // 如果指定了用户id,则只搜索该用户的帖子
    if ($uid > 0) {
        $query["body"]["query"] = ["bool" => [
            "must" => [
                "multi_match" => [
                    "query" => $keyword,
                    "fields" => ["title", "content"],
                ],
            ],
            "filter" => [
                "term" => [
                    "authorid" => $uid
                ]
            ]
        ]];
    }
    $resp = Es::getClient()->search($query);
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
        "page"=>$page,
    ], JSON_UNESCAPED_UNICODE);
}

// user 用户名前缀搜索用户
function user()
{
    global $_G;
    $username = $_GET["username"];
    if (empty($username)) {
        showmessage("请输入用户名");
        return;
    }

    $list = C::t('common_member')->fetch_all_by_like_username($username, 0, 5);
    if (empty($list)) {
        try {
            // 从归档表中查询
            $list = C::t('common_member_archive')->fetch_all_by_like_username($username, 0, 5);
        } catch (Exception $e) {
            // ignore
        }
    }
    $data = [];
    foreach ($list as $k => &$item) {
        $data[] = [
            "uid" => $item["uid"],
            "username" => $item["username"],
        ];
    }

    echo json_encode([
        "code" => 0,
        "data" => $data,
    ], JSON_UNESCAPED_UNICODE);
}