from machine import CompartmentParams


class CompartmentParamsController(CompartmentParams):
    def __init__(self,state_param,enter_param):
        super().__init__(state_param,enter_param)

    def print_do(self,s):
        print(s)

sm = CompartmentParamsController(1,2)
