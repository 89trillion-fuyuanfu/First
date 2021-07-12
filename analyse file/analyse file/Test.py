from locust import HttpUser, task


class WebsiteUser(HttpUser):
    min_wait = 3000  # 执行事务之间用户等待时间的下界（单位：毫秒）
    max_wait = 6000  # 执行事务之间用户等待时间的上界（单位：毫秒

    @task(1)
    def getcombatpoints(self):
        self.client.get("/getcombatpoints")

    @task(2)
    def getsolider(self):
        self.client.get("/getsolider")

    @task(3)
    def getrarity(self):
        self.client.get("/getrarity")

    @task(4)
    def getlegalsolider(self):
        self.client.get("/getlegalsolider")

    @task(5)
    def getsoliderjson(self):
        self.client.get("/getsoliderjson")