from machine import FizzBuzz


class FizzBuzzController(FizzBuzz):

    t = ""
    def __init__(self) -> None:
        super().__init__()

    def print_do(self, msg):
        self.t += msg
        print(self.t)

    def gt_100_do(self, i):
        return i > 100

    def mod3_eq0_do(self, i):
        return i % 3 == 0

    def mod5_eq0_do(self, i):
        return i % 5 == 0

    def plus_1_do(self, i):
        return i + 1


sm = FizzBuzzController()
sm.start()