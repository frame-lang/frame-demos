# emitted from framec_v0.10.0
# get include files at https://github.com/frame-lang/frame-ancillary-files

from framelang.framelang import FrameEvent

class Countdown:
    
    def __init__(self, i):
        
        # Create and intialize start state compartment.
        self.state = self._sS0_
        self.compartment = CountdownCompartment(self.state)
        self.next_compartment = None
        self.compartment.state_args["i"] = i
        self.compartment.state_vars["dec"] = 1
        
        # Initialize domain
        
        # Send system start event
        frame_event = FrameEvent(">", None)
        self.mux(frame_event)
    
    # ====================== Multiplexer ==================== #
    
    def mux(self, e):
        if self.compartment.state == self._sS0_:
            self._sS0_(e)
        elif self.compartment.state == self._sS1_:
            self._sS1_(e)
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
    
    def _sS0_(self, e):
        if e._message == ">":
            (self.compartment.state_args["i"]) = (self.compartment.state_args["i"]) - (self.compartment.state_vars["dec"])
            self.print_do(str((self.compartment.state_args["i"])))
            if  (self.compartment.state_args["i"]) == 0:
                compartment = CountdownCompartment(self._sStop_)
                self.transition(compartment)
                return
            
            compartment = CountdownCompartment(self._sS1_)
            compartment.enter_args["i"] = self.compartment.state_args["i"]
            self.transition(compartment)
            return
        
    def _sS1_(self, e):
        if e._message == ">":
            compartment = CountdownCompartment(self._sS0_)
            compartment.state_args["i"] = e._parameters["i"]
            compartment.state_vars["dec"] = 1
            self.transition(compartment)
            return
        
    def _sStop_(self, e):
        if e._message == ">":
            self.print_do("done")
            return
        
    
    # ===================== Actions Block =================== #
    
    
    def print_do(self,s):
        print(s)
    
    # Unimplemented Actions
    
    
    # =============== Machinery and Mechanisms ============== #
    
    def transition(self, compartment):
        self.next_compartment = compartment
    
    def do_transition(self, next_compartment):
        self.mux(FrameEvent("<", self.compartment.exit_args))
        self.compartment = next_compartment
        self.mux(FrameEvent(">", self.compartment.enter_args))
    

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


