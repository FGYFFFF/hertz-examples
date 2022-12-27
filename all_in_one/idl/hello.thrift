namespace go hello.example

struct HelloReq {
    1: string Name (api.query="name");
}

struct HelloResp {
    1: string RespBody;
}


service HelloService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello");
    HelloResp HelloMethod2(1: HelloReq request) (api.get="/hello2");
    HelloResp HelloMethod3(1: HelloReq request) (api.get="/hello3");
    HelloResp HelloMethod4(1: HelloReq request) (api.get="/hello4");
    HelloResp HelloMethod5(1: HelloReq request) (api.get="/hello5", api.handler_path="v5");
    HelloResp HelloMethod6(1: HelloReq request) (api.get="/hello6", api.handler_path="v6");
    HelloResp HelloMethod7(1: HelloReq request) (api.get="/hello7", api.handler_path="v7");
    HelloResp HelloMethod8(1: HelloReq request) (api.get="/hello8", api.handler_path="v8");
    HelloResp HelloMethod9(1: HelloReq request) (api.get="/hello9", api.handler_path="v9");
    HelloResp HelloMethod10(1: HelloReq request) (api.get="/hello10", api.handler_path="v10");
    HelloResp HelloMethod11(1: HelloReq request) (api.get="/hello11", api.handler_path="v11");
    HelloResp HelloMethod12(1: HelloReq request) (api.get="/hello12");
}
