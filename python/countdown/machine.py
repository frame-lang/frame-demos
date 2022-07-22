# emitted from framec_v0.10.0
# get include files at https://github.com/frame-lang/frame-ancillary-files
from framelang.framelang import FrameEvent

class Countdown:
    
    def __init__(self, i):
        
        # Create and intialize start state compartment.
        self.__state = self.__countdown_state_S0
        self.__compartment: 'CountdownCompartment' = CountdownCompartment(self.__state)
        self.__next_compartment: 'CountdownCompartment' = None
        self.__compartment.state_args["i"] = i
        self.__compartment.state_vars["dec"] = 1
        
        # Initialize domain
        
        # Send system start event
        frame_event = FrameEvent(">", None)
        self.__mux(frame_event)
    
    # ====================== Multiplexer ==================== #
    
    def __mux(self, e):
        if self.__compartment.state == self.__countdown_state_S0:
            self.__countdown_state_S0(e)
        elif self.__compartment.state == self.__countdown_state_S1:
            self.__countdown_state_S1(e)
        elif self.__compartment.state == self.__countdown_state_Stop:
            self.__countdown_state_Stop(e)
        
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
    
    def __countdown_state_S0(self, e):
        if e._message == ">":
            (self.__compartment.state_args["i"]) = (self.__compartment.state_args["i"]) - (self.__compartment.state_vars["dec"])
            self.print_do(str((self.__compartment.state_args["i"])))
            if  (self.__compartment.state_args["i"]) == 0:
                compartment = CountdownCompartment(self.__countdown_state_Stop)
                self.__transition(compartment)
                return
            
            compartment = CountdownCompartment(self.__countdown_state_S1)
            compartment.enter_args["i"] = self.__compartment.state_args["i"]
            self.__transition(compartment)
            return
        
    def __countdown_state_S1(self, e):
        if e._message == ">":
            compartment = CountdownCompartment(self.__countdown_state_S0)
            compartment.state_args["i"] = e._parameters["i"]
            compartment.state_vars["dec"] = 1
            self.__transition(compartment)
            return
        
    def __countdown_state_Stop(self, e):
        if e._message == ">":
            self.print_do("done")
            return
        
    
    # ===================== Actions Block =================== #
    
    
    def print_do(self,s: str):
        print(s)
    
    # Unimplemented Actions
    
    
    # =============== Machinery and Mechanisms ============== #
    
    def __transition(self, compartment: 'CountdownCompartment'):
        self.__next_compartment = compartment
    
    def  __do_transition(self, next_compartment: 'CountdownCompartment'):
        self.__mux(FrameEvent("<", self.__compartment.exit_args))
        self.__compartment = next_compartment
        self.__mux(FrameEvent(">", self.__compartment.enter_args))
    

# ===================== Compartment =================== #

class CountdownCompartment:

    def __init__(self, state):
        self.state = state
        self.state_args = {}
        self.state_vars = {}
        self.enter_args = {}
        self.exit_args = {}
        self.forward_event = FrameEvent(None, None)
    


# ********************

#class CountdownController(Countdown):
	#def __init__(self,i):
	    #super().__init__(i)

# ********************


