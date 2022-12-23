namespace go hello.example

struct HelloReq {
    1: string Name (api.query="name");
}

struct HelloResp {
    1: string RespBody;
}


service HelloService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello", api.handler_path="v1");
    HelloResp HelloMethod2(1: HelloReq request) (api.get="/hello2", api.handler_path="v1");
    HelloResp HelloMethod3(1: HelloReq request) (api.get="/hello3", api.handler_path="v1");
    HelloResp HelloMethod4(1: HelloReq request) (api.get="/hello4", api.handler_path="v2");
    HelloResp HelloMethod5(1: HelloReq request) (api.get="/hello5", api.handler_path="v2");
}
