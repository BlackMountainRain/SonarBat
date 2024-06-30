# consul
Fork from [https://github.com/go-kratos/kratos/tree/main/contrib/registry/consul](https://github.com/go-kratos/kratos/tree/main/contrib/registry/consul),
and split the service into `HTTP` and `GRPC` services when registering services to the consul.

Because the original registry service as `TCP`, which can't be used in the `consul-template` to generate the `nginx` configuration file.
