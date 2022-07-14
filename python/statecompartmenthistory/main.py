
from state_history_basic import HistoryBasic
from state_compartment_history import HistoryStateContext

class HistoryStateContextController(HistoryStateContext):
    def __init__(self,):
        super().__init__()

    def print_do(self,msg):
        print(msg)
class HistoryBasicController(HistoryBasic):
    def __init__(self,):
        super().__init__()

    def print_do(self,msg):
        print(msg)


def run():
    basic = HistoryBasicController()
    basic.start()
    basic.switchState()
    basic.gotoDeadEnd()
    basic.back()
    
    compartment = HistoryStateContextController()
    compartment.start()
    compartment.switchState()
    compartment.gotoDeadEnd()
    compartment.back()


run()