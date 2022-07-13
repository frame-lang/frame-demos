

from machine import Countdown


class CountdownController(Countdown):
    def __init__(self,i):
        super().__init__(i)

sm = CountdownController(5)