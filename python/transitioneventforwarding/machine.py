# emitted from framec_v0.10.0
# get include files at https://github.com/frame-lang/frame-ancillary-files

from framelang.framelang import FrameEvent

class TransitionEventForwarding:
    
    def __init__(self, cycles):
        
        # Create and intialize start state compartment.
        self.state = self._sStart_
        self.compartment = TransitionEventForwardingCompartment(self.state)
        self.next_compartment = None
        self.compartment.enter_args["cycles"] = cycles
        
        # Initialize domain
        
        # Send system start event
        frame_event = FrameEvent(">", self.compartment.enter_args)
        self.mux(frame_event)
    
    # ====================== Multiplexer ==================== #
    
    def mux(self, e):
        if self.compartment.state == self._sStart_:
            self._sStart_(e)
        elif self.compartment.state == self._sForwardEventAgain_:
            self._sForwardEventAgain_(e)
        elif self.compartment.state == self._sDecrement_:
            self._sDecrement_(e)
        elif self.compartment.state == self._sStop_:
            self._sStop_(e)
        
        if self.next_compartment != None:
            next_compartment = self.next_compartment
            self.next_compartment = None
            if(next_compartment.forward_event != None and 
               next_compartment.forward_event._message == ">"):
                self.mux(FrameEvent( "<", self.compartment.exit_args))
                self.compartment = next_compartment
                self.mux(next_compartment.forward_event)
            else:
                self.do_transition(next_compartment)
                if next_compartment.forward_event != None:
                    self.mux(next_compartment.forward_event)
            next_compartment.forward_event = None
    
    # ===================== Machine Block =================== #
    
    def _sStart_(self, e):
        if e._message == ">":
            if  e._parameters["cycles"] == 0:
                self.compartment.exit_args["msg"] = "stopping"
                compartment = TransitionEventForwardingCompartment(self._sStop_)
                compartment.forward_event = e
                self.transition(compartment)
                return
            else:
                self.compartment.exit_args["msg"] = "keep going"
                compartment = TransitionEventForwardingCompartment(self._sForwardEventAgain_)
                compartment.forward_event = e
                self.transition(compartment)
            
            return
        
        elif e._message == "<":
            self.print_do(e._parameters["msg"])
            return
        
    def _sForwardEventAgain_(self, e):
        if e._message == ">":
            compartment = TransitionEventForwardingCompartment(self._sDecrement_)
            compartment.forward_event = e
            self.transition(compartment)
            return
        
    def _sDecrement_(self, e):
        if e._message == ">":
            self.print_do(str(e._parameters["cycles"]))
            compartment = TransitionEventForwardingCompartment(self._sStart_)
            compartment.enter_args["cycles"] = (e._parameters["cycles"] - 1)
            self.transition(compartment)
            return
        
    def _sStop_(self, e):
        if e._message == ">":
            self.print_do(str(e._parameters["cycles"]))
            self.print_do("done")
            return
        
    
    # ===================== Actions Block =================== #
    
    
    # Unimplemented Actions
    
    def print_do(self,msg):
        raise NotImplementedError
    
    
    # =============== Machinery and Mechanisms ============== #
    
    def transition(self, compartment):
        self.next_compartment = compartment
    
    def do_transition(self, next_compartment):
        self.mux(FrameEvent("<", self.compartment.exit_args))
        self.compartment = next_compartment
        self.mux(FrameEvent(">", self.compartment.enter_args))
    

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
	#def __init__(self,cycles):
	    #super().__init__(cycles)

    #def print_do(self,msg):
        #pass

# ********************


