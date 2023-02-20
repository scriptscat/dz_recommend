<?php


use Codfrm\DzMarkdown\ParsedownExt;
use Michelf\MarkdownExtra;

if (!defined('IN_DISCUZ')) {
    exit('Access Denied');
}

include "vendor/autoload.php";
require "lib/es.php";

class plugin_codfrm_recommend
{

}

class plugin_codfrm_recommend_forum extends plugin_codfrm_recommend
{

    // 帖子 viewthread_modaction_output hook
    function viewthread_modaction_output()
    {
        global $_G;
        include_once template('codfrm_recommend:module');
        // 缓存
        $cacheKey = "codfrm_recommend:module:{$_G['tid']}";
        $cache = loadcache($cacheKey);
        if ($cache && $_G['cache'][$cacheKey]['time'] > time() - 3600) {
            return tpl_viewthread_postbottom_recommend($_G['cache'][$cacheKey]['list']);
        }
        $list = [];
        try {
            $resp = Es::getClient()->search([
                "index" => "dz.forum_thread",
                "body" => [
                    "query" => [
                        "match" => [
                            "title" => $_G['forum_thread']['subject']
                        ]
                    ]
                ],
                "size" => 7,
            ]);
            // 排除当前帖子
            $list = $resp['hits']['hits'];
            $list = array_filter($list, function ($item) {
                global $_G;
                return $item['_id'] != $_G['tid'];
            });
            // 只取6个
            if (count($list) > 6) {
                $list = array_slice($list, 0, 6);
            }
        } catch (Exception $e) {
            // 屏蔽错误
        }
        // 写缓存
        savecache($cacheKey, [
            "list" => $list,
            "time" => time(),
        ]);
        return tpl_viewthread_postbottom_recommend($list);
    }

    /**
     * 拦截帖子操作
     * @param $param
     * @return void
     */
    public function post_message($param)
    {
        global $_G;
        try {
            switch ($param['param'][0]) {
                case "post_newthread_succeed":
                    // 添加新记录
                    Es::insertThread(
                        $param['param'][2]['tid'],
                        $param['param'][2]['fid'],
                        $_POST['subject'],
                        $_POST['message'],
                        $_G['uid'],
                        time(),
                    );
                    break;
                case "post_edit_succeed":
                    // 更新记录
                    Es::updateThread(
                        $param['param'][2]['tid'],
                        $param['param'][2]['fid'],
                        $_POST['subject'],
                        $_POST['message'],
                    );
                    break;
            }
        } catch (Exception $e) {
            // 屏蔽错误
        }
    }

}

class mobileplugin_codfrm_recommend_forum extends plugin_codfrm_recommend_forum
{

}