from machine import SystemParams


class SystemParamsController(SystemParams):

    def __init__(self, stateMsg, enterMsg):
        super().__init__(stateMsg, enterMsg)
    

SystemParamsController("Hello", "World")