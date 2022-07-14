# emitted from framec_v0.10.0
# get include files at https://github.com/frame-lang/frame-ancillary-files

from framelang.framelang import FrameEvent

class HistoryBasic:
    
    def __init__(self, ):
        
        # Create state stack.
        
        self.state_stack = []
        
        # Create and intialize start state compartment.
        self.state = self._sStart_
        self.compartment = HistoryBasicCompartment(self.state)
        self.next_compartment = None
        
        # Initialize domain
        
        # Send system start event
        frame_event = FrameEvent(">", None)
        self.mux(frame_event)
    
    # ===================== Interface Block =================== #
    
    def start(self):
        e = FrameEvent(">>",None)
        self.mux(e)
    
    def switchState(self):
        e = FrameEvent("switchState",None)
        self.mux(e)
    
    def gotoDeadEnd(self):
        e = FrameEvent("gotoDeadEnd",None)
        self.mux(e)
    
    def back(self):
        e = FrameEvent("back",None)
        self.mux(e)
    
    # ====================== Multiplexer ==================== #
    
    def mux(self, e):
        if self.compartment.state == self._sStart_:
            self._sStart_(e)
        elif self.compartment.state == self._sS0_:
            self._sS0_(e)
        elif self.compartment.state == self._sS1_:
            self._sS1_(e)
        elif self.compartment.state == self._sDeadEnd_:
            self._sDeadEnd_(e)
        
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
        if e._message == ">>":
            compartment = HistoryBasicCompartment(self._sS0_)
            self.transition(compartment)
            return
        
    def _sS0_(self, e):
        if e._message == ">":
            self.print_do("Enter S0")
            return
        
        elif e._message == "switchState":
            # Switch\nState
            compartment = HistoryBasicCompartment(self._sS1_)
            self.transition(compartment)
            return
        
        elif e._message == "gotoDeadEnd":
            self.state_stack_push(self.compartment)
            # Goto\nDead End
            compartment = HistoryBasicCompartment(self._sDeadEnd_)
            self.transition(compartment)
            return
        
    def _sS1_(self, e):
        if e._message == ">":
            self.print_do("Enter S1")
            return
        
        elif e._message == "switchState":
            # Switch\nState
            compartment = HistoryBasicCompartment(self._sS0_)
            self.transition(compartment)
            return
        
        elif e._message == "gotoDeadEnd":
            self.state_stack_push(self.compartment)
            # Goto\nDead End
            compartment = HistoryBasicCompartment(self._sDeadEnd_)
            self.transition(compartment)
            return
        
    def _sDeadEnd_(self, e):
        if e._message == ">":
            self.print_do("Enter $DeadEnd")
            return
        
        elif e._message == "back":
            # Go Back
            compartment = self.state_stack_pop()
            self.transition(compartment)
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
    
    def state_stack_push(self, compartment):
        self.state_stack.append(compartment)
    
    def state_stack_pop(self):
        return self.state_stack.pop()
    
    

# ===================== Compartment =================== #

class HistoryBasicCompartment:

    def __init__(self, state):
        self.state = state
        self.state_args = {}
        self.state_vars = {}
        self.enter_args = {}
        self.exit_args = {}
        self.forward_event = FrameEvent(None, None)
    


# ********************

#class HistoryBasicController(HistoryBasic):
	#def __init__(self,):
	    #super().__init__()

    #def print_do(self,msg):
        #pass

# ********************


