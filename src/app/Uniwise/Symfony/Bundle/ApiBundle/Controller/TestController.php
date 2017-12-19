<?php
namespace Uniwise\Symfony\Bundle\ApiBundle\Controller;

use FOS\RestBundle\Controller\Annotations\Get;
use FOS\RestBundle\Controller\Annotations\Route;
use FOS\RestBundle\Controller\FOSRestController;
use Symfony\Component\HttpFoundation\Request;

/**
 * @Route("/test")
 */
class TestController extends FOSRestController {

    /**
     * @param Request $request
     * @Get("/")
     *
     * @return \FOS\RestBundle\View\View
     */
    public function getList(Request $request) {
        return $this->view(["hello" => "world"]);
    }
}