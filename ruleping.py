from HABApp import Rule
from HABApp.openhab.items import NumberItem
import time
from datetime import timedelta

class RulePing(Rule):
    def __init__(self):
        super().__init__()
        self.ping_item = NumberItem.get_item("HABApp_LastRulePing")
        self.run.every(timedelta(seconds=1), timedelta(seconds=30), self.check)

    async def check(self):
      self.ping_item.oh_post_update(int(time.time()))

RulePing()
