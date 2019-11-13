<?php

namespace App\Controller;

use Jphp\Application\WebApplication;
use Jphp\Controller\HttpController;
use Jphp\Http\Request;
use Jphp\Http\Response;
use Jphp\Template\Template;

/**
 * Class TestController
 * @package App\Controller
 */
class TestController extends HttpController {

	/**
	 * @methods=["GET","POST"];
	 * @param Request $request
	 * @return Response
	 */
	public function indexAction(Request $request) {

        $app = WebApplication::getApp();
        $view_path = $app->getViewPath();
        $test_tpl = $view_path . DIRECTORY_SEPARATOR . "index.html";
//        $content = file_get_contents($test_tpl);
        $data = array("a" => 1, "b" => 2);
        $template = new Template();
        $content = $template->parse($test_tpl, $data);

		return new Response($content);
	}

}
