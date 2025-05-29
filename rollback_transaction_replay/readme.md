## 如何使用软回滚工具

1. 首先要在epoch卡住的阶段把交易信息存到表里
    使用save_transaction里面的工具
    - basePath要改成bfc工具所在的目录
    - 修改common.GetAllCheckPoints(cli, ${epoch}) 中的epoch为当前的epoch
    - 在区块链浏览器获取到最新的checkpoint number sequence 并且修改end的值为此值
    - 运行工具，数据会存在 rollback_data.db中
2. 进行一遍硬回滚
3. 回放交易到链上
    - 使用replay_transaction中的工具
    - 将之前存好的数据copy到这个目录下
    - 同样修改common.GetAllCheckPoints(cli, ${epoch}) 中的epoch为当前的epoch
    - 将end改为之前的checkpoint number sequence
    - 运行工具，检查是否数据库中的rollbackdata表中的交易是否为success
4. 检查链上的交易是否被回放

