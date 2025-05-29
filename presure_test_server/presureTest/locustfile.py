# https://www.jianshu.com/p/5b1381f17444
import random

from locust import HttpUser, task, between
import utils



# 相当于模拟一个用户
class WebUser(HttpUser):
    # 服务器的地址
    host = 'http://localhost:13344'

    threshold = 't32'
    index = 0

    # 下一个任务执行之前等待的时间，用于模式用户的思考时间
    wait_time = between(1, 3)  # 这里使用随机 3，5 秒钟
    user = utils.random_user()
    print(user)

    # 定义一个测试任务，类似于一个取样器

    @task(0)
    def start_sign(self):
        data = "userId=c@c.com&msg=68656c6c6f&t="+self.threshold
        self.client.get('/signing?'+data)

    @task(100)
    def start_paybfc(self):
        accountid = utils.random_account(self.index)
        self.client.get('/ping' )
        self.index = self.index + 1


    def on_start(self):
        # 测试之前执行的操作
        print('用户登录')
        #auth = {'username': 'root', 'password': 'root1234'}
        #self.client.post('/login/', json=auth)

    def on_stop(self):
        # 测试结束执行的操作
        print('用户退出')
        #self.client.delete('/logout/')





