<?php


use Codfrm\DzMarkdown\ParsedownExt;
use Michelf\MarkdownExtra;

if (!defined('IN_DISCUZ')) {
    exit('Access Denied');
}

class plugin_codfrm_recommend
{

}

class plugin_codfrm_recommend_forum extends plugin_codfrm_recommend
{

    function viewthread_modaction_output()
    {
        global $_G;
        include_once template('codfrm_recommend:module');
        $resp = file_get_contents("http://127.0.0.1:20141/recommend?tid={$_G['tid']}");
        $resp = json_decode($resp, true);
        $list = $resp['data'] ?? [];
        return tpl_viewthread_postbottom_recommend($list);
    }

}
