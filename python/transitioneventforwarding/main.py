from machine import TransitionEventForwarding

class TransitionEventForwardingController(TransitionEventForwarding):

    def __init__(self,cycles):
        super().__init__(cycles)

    def print_do(self,msg):
        print(msg)

sm = TransitionEventForwardingController(5)
