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

        $list = $resp['data'] ?? [];
        return tpl_viewthread_postbottom_recommend($list);
    }

}
