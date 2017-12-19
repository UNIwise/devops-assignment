<?php
require_once('../Uniwise/Symfony/autoload.php');

$environments = ["local", "dev", "stage", "prod", "test"];
$server_environ = getenv("SERVER_ENVIRONMENT");

// Exception if environment it not set correctly
if (!$server_environ || !in_array($server_environ, $environments)) {
    throw new Exception("SERVER ENVIRONMENT NOT SET");
}

$debug = true;
$kernel = new AppKernel($server_environ, $debug);
$kernel->boot();

$request = \Symfony\Component\HttpFoundation\Request::createFromGlobals();
$response = $kernel->handle($request);
$response->send();
$kernel->terminate($request, $response);
