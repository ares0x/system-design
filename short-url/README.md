# README.md

## 核心逻辑
短链接生成平台的核心逻辑是 /internal/service/service.go 中的那个 interface，其中包含 Encode 和 Decode 两个方法。Encode 方法通过一个唯一 id
将长链接编码为 62 位的短链接，并将该 id 作为唯一id存储在数据库中；Decode 则将短链接反编码为原始 id，通过 id 获取与之映射的长链接。 

## 可以优化的地方
* uuid 的生成方式。在分布式环境下建议通过第三方组件来生成递增的唯一 id，以保证生成的 shortUrl 不重复
* 可以通过 redis 做缓存来加速访问
