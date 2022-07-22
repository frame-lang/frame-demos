# emitted from framec_v0.10.0
# get include files at https://github.com/frame-lang/frame-ancillary-files
from framelang.framelang import FrameEvent

class SystemParams:
    
    def __init__(self, stateMsg: str,enterMsg: str):
        
        # Create and intialize start state compartment.
        self.__state = self.__systemparams_state_Begin
        self.__compartment: 'SystemParamsCompartment' = SystemParamsCompartment(self.__state)
        self.__next_compartment: 'SystemParamsCompartment' = None
        self.__compartment.state_args["stateMsg"] = stateMsg
        self.__compartment.enter_args["enterMsg"] = enterMsg
        
        # Initialize domain
        
        # Send system start event
        frame_event = FrameEvent(">", self.__compartment.enter_args)
        self.__mux(frame_event)
    
    # ====================== Multiplexer ==================== #
    
    def __mux(self, e):
        if self.__compartment.state == self.__systemparams_state_Begin:
            self.__systemparams_state_Begin(e)
        
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
    
    def __systemparams_state_Begin(self, e):
        if e._message == ">":
            self.print_do((self.__compartment.state_args["stateMsg"]) + " " + e._parameters["enterMsg"])
            return
        
    
    # ===================== Actions Block =================== #
    
    
    def print_do(self,msg: str):
        
        print(msg)
    
    
    # Unimplemented Actions
    
    
    # =============== Machinery and Mechanisms ============== #
    
    def __transition(self, compartment: 'SystemParamsCompartment'):
        self.__next_compartment = compartment
    
    def  __do_transition(self, next_compartment: 'SystemParamsCompartment'):
        self.__mux(FrameEvent("<", self.__compartment.exit_args))
        self.__compartment = next_compartment
        self.__mux(FrameEvent(">", self.__compartment.enter_args))
    

# ===================== Compartment =================== #

class SystemParamsCompartment:

    def __init__(self, state):
        self.state = state
        self.state_args = {}
        self.state_vars = {}
        self.enter_args = {}
        self.exit_args = {}
        self.forward_event = FrameEvent(None, None)
    


# ********************

#class SystemParamsController(SystemParams):
	#def __init__(self,stateMsg: str,enterMsg: str):
	    #super().__init__(stateMsg: str,enterMsg: str)

# ********************


