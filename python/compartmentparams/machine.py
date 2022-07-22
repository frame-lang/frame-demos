# emitted from framec_v0.10.0
# get include files at https://github.com/frame-lang/frame-ancillary-files
from framelang.framelang import FrameEvent
class CompartmentParams:
    
    def __init__(self, state_param,enter_param):
        
        # Create and intialize start state compartment.
        self.__state = self.__compartmentparams_state_S0
        self.__compartment: 'CompartmentParamsCompartment' = CompartmentParamsCompartment(self.__state)
        self.__next_compartment: 'CompartmentParamsCompartment' = None
        self.__compartment.state_args["state_param"] = state_param
        self.__compartment.state_vars["state_var"] = 100
        self.__compartment.enter_args["enter_param"] = enter_param
        
        # Initialize domain
        
        # Send system start event
        frame_event = FrameEvent(">", self.__compartment.enter_args)
        self.__mux(frame_event)
    
    # ====================== Multiplexer ==================== #
    
    def __mux(self, e):
        if self.__compartment.state == self.__compartmentparams_state_S0:
            self.__compartmentparams_state_S0(e)
        elif self.__compartment.state == self.__compartmentparams_state_S1:
            self.__compartmentparams_state_S1(e)
        
        if self.__next_compartment != None:
            next_compartment = self.__next_compartment
            self.__next_compartment = None
            if(next_compartment.forward_event is not None and 
               next_compartment.forward_event._message == ">"):
                self.__mux(FrameEvent( "<", self.__compartment.exit_args))
                self.__compartment = next_compartment
                self.__mux(next_compartment.forward_event)
            else:
                self.__do_transition(next_compartment)
                if next_compartment.forward_event is not None:
                    self.__mux(next_compartment.forward_event)
            next_compartment.forward_event = None
    
    # ===================== Machine Block =================== #
    
    def __compartmentparams_state_S0(self, e):
        if e._message == ">":
            self.print_do(str((self.__compartment.state_args["state_param"])) + " " + str((self.__compartment.state_vars["state_var"])) + " " + str(e._parameters["enter_param"]))
            compartment = CompartmentParamsCompartment(self.__compartmentparams_state_S1)
            compartment.forward_event = e
            compartment.state_args["state_param"] = self.__compartment.state_args["state_param"] + 20
            compartment.state_vars["state_var"] = 200
            self.__transition(compartment)
            return
        
    def __compartmentparams_state_S1(self, e):
        if e._message == ">":
            self.print_do(str((self.__compartment.state_args["state_param"])) + " " + str((self.__compartment.state_vars["state_var"])) + " " + str(e._parameters["enter_param"]))
            return
        
    
    # ===================== Actions Block =================== #
    
    
    def print_do(self,s: str):
        print(s)
    
    # Unimplemented Actions
    
    
    # =============== Machinery and Mechanisms ============== #
    
    def __transition(self, compartment: 'CompartmentParamsCompartment'):
        self.__next_compartment = compartment
    
    def  __do_transition(self, next_compartment: 'CompartmentParamsCompartment'):
        self.__mux(FrameEvent("<", self.__compartment.exit_args))
        self.__compartment = next_compartment
        self.__mux(FrameEvent(">", self.__compartment.enter_args))
    

# ===================== Compartment =================== #

class CompartmentParamsCompartment:

    def __init__(self, state):
        self.state = state
        self.state_args = {}
        self.state_vars = {}
        self.enter_args = {}
        self.exit_args = {}
        self.forward_event = FrameEvent(None, None)
    


# ********************

#class CompartmentParamsController(CompartmentParams):
	#def __init__(self,state_param,enter_param):
	    #super().__init__(state_param,enter_param)

# ********************


