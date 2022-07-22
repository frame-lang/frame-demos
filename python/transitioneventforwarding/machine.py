# emitted from framec_v0.10.0
# get include files at https://github.com/frame-lang/frame-ancillary-files
from framelang.framelang import FrameEvent

class TransitionEventForwarding:
    
    def __init__(self, cycles: int):
        
        # Create and intialize start state compartment.
        self.__state = self.__transitioneventforwarding_state_Start
        self.__compartment: 'TransitionEventForwardingCompartment' = TransitionEventForwardingCompartment(self.__state)
        self.__next_compartment: 'TransitionEventForwardingCompartment' = None
        self.__compartment.enter_args["cycles"] = cycles
        
        # Initialize domain
        
        # Send system start event
        frame_event = FrameEvent(">", self.__compartment.enter_args)
        self.__mux(frame_event)
    
    # ====================== Multiplexer ==================== #
    
    def __mux(self, e):
        if self.__compartment.state == self.__transitioneventforwarding_state_Start:
            self.__transitioneventforwarding_state_Start(e)
        elif self.__compartment.state == self.__transitioneventforwarding_state_ForwardEventAgain:
            self.__transitioneventforwarding_state_ForwardEventAgain(e)
        elif self.__compartment.state == self.__transitioneventforwarding_state_Decrement:
            self.__transitioneventforwarding_state_Decrement(e)
        elif self.__compartment.state == self.__transitioneventforwarding_state_Stop:
            self.__transitioneventforwarding_state_Stop(e)
        
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
    
    def __transitioneventforwarding_state_Start(self, e):
        if e._message == ">":
            if  e._parameters["cycles"] == 0:
                self.__compartment.exit_args["msg"] = "stopping"
                compartment = TransitionEventForwardingCompartment(self.__transitioneventforwarding_state_Stop)
                compartment.forward_event = e
                self.__transition(compartment)
                return
            else:
                self.__compartment.exit_args["msg"] = "keep going"
                compartment = TransitionEventForwardingCompartment(self.__transitioneventforwarding_state_ForwardEventAgain)
                compartment.forward_event = e
                self.__transition(compartment)
            
            return
        
        elif e._message == "<":
            self.print_do(e._parameters["msg"])
            return
        
    def __transitioneventforwarding_state_ForwardEventAgain(self, e):
        if e._message == ">":
            compartment = TransitionEventForwardingCompartment(self.__transitioneventforwarding_state_Decrement)
            compartment.forward_event = e
            self.__transition(compartment)
            return
        
    def __transitioneventforwarding_state_Decrement(self, e):
        if e._message == ">":
            self.print_do(str(e._parameters["cycles"]))
            compartment = TransitionEventForwardingCompartment(self.__transitioneventforwarding_state_Start)
            compartment.enter_args["cycles"] = (e._parameters["cycles"] - 1)
            self.__transition(compartment)
            return
        
    def __transitioneventforwarding_state_Stop(self, e):
        if e._message == ">":
            self.print_do(str(e._parameters["cycles"]))
            self.print_do("done")
            return
        
    
    # ===================== Actions Block =================== #
    
    
    # Unimplemented Actions
    
    def print_do(self,msg: str):
        raise NotImplementedError
    
    
    # =============== Machinery and Mechanisms ============== #
    
    def __transition(self, compartment: 'TransitionEventForwardingCompartment'):
        self.__next_compartment = compartment
    
    def  __do_transition(self, next_compartment: 'TransitionEventForwardingCompartment'):
        self.__mux(FrameEvent("<", self.__compartment.exit_args))
        self.__compartment = next_compartment
        self.__mux(FrameEvent(">", self.__compartment.enter_args))
    

# ===================== Compartment =================== #

class TransitionEventForwardingCompartment:

    def __init__(self, state):
        self.state = state
        self.state_args = {}
        self.state_vars = {}
        self.enter_args = {}
        self.exit_args = {}
        self.forward_event = FrameEvent(None, None)
    


# ********************

#class TransitionEventForwardingController(TransitionEventForwarding):
	#def __init__(self,cycles: int):
	    #super().__init__(cycles: int)

    #def print_do(self,msg: str):
        #pass

# ********************


