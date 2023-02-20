<?php

if (!defined('IN_DISCUZ')) {
    exit('Access Denied');
}

include "vendor/autoload.php";
require "lib/es.php";


switch ($_GET["operation"]) {
    case "refresh":
        refresh();
        break;
    default:
        include_once template('codfrm_recommend:manager');
        echo tpl_manager();
}

function refresh()
{
    set_time_limit(0);
    ignore_user_abort(true);
    // SSE fetch
    header('Content-Type: text/event-stream');
    header('Cache-Control: no-cache');
    header('Connection: keep-alive');
    try {
        // 删除之前的索引
        $resp = Es::getClient()->indices()->delete(["index" => "dz.forum_thread"]);
        var_dump($resp);
    } catch (Exception $e) {
        var_dump($e->getMessage());
    }
    $startPos = 0;
    while (true) {
        // 从mysql中查询dz帖子
        $list = DB::fetch_all("SELECT * FROM %t LIMIT %d, %d", ["forum_thread", $startPos, 10]);
        if (empty($list)) {
            break;
        }
        foreach ($list as $data) {
            // 将帖子数据写入es
            // 查询帖子内容
            $post = C::t("forum_post")->fetch_all_by_tid("tid:" . $data['tid'], $data['tid'], true, '', 0, 1);
            $post = array_pop($post);
            Es::insertThread(
                $data['tid'], $data['fid'], $data['subject'], $post['message'],
                $data['authorid'], $data['dateline']
            );
            echo "data: {$data['tid']} {$data['subject']}";
            echo PHP_EOL;
            ob_flush();
            flush();
        }
        $startPos += 10;
    }
    echo "data: done";
}