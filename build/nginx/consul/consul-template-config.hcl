template {
    source = "/etc/consul-template/nginx.ctmpl"
    destination = "/etc/nginx/conf.d/default.conf"
    command = "touch /nginx_signal/reload"
}