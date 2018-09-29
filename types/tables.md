整体思路：
> 此为企业最前端负载均衡
> server [1 to 1]-> location [1 to 1]-> upstream [1 to many]-> nodes

# 数据表

## nginx_servers 表
    id
    listen int
    root string
    server_name string -> golang []domain_string
    index string

## nginx_location
    id
    health_check bool
    server_id int

## upstreams 表
table: nginx_upstreams
    id int
    upstream_name string
    location_id int

## upstream 节点表
table: nginx_node
    id int
    upstream_id int -> nginx_upstreams.id
    server string
    port int
    weight int
    fail_timeout int
    slow_start int
    resolve bool
    backup bool
    down bool