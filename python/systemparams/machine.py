# emitted from framec_v0.10.0
# get include files at https://github.com/frame-lang/frame-ancillary-files
from framelang.framelang import FrameEvent
class SystemParams:
    
    def __init__(self, stateMsg,enterMsg):
        
        # Create and intialize start state compartment.
        self.state = self._sBegin_
        self.compartment = SystemParamsCompartment(self.state)
        self.next_compartment = None
        self.compartment.state_args["stateMsg"] = stateMsg
        self.compartment.enter_args["enterMsg"] = enterMsg
        
        # Initialize domain
        
        # Send system start event
        frame_event = FrameEvent(">", self.compartment.enter_args)
        self.mux(frame_event)
    
    # ====================== Multiplexer ==================== #
    
    def mux(self, e):
        if self.compartment.state == self._sBegin_:
            self._sBegin_(e)
        
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
    
    def _sBegin_(self, e):
        if e._message == ">":
            self.print_do((self.compartment.state_args["stateMsg"]) + " " + e._parameters["enterMsg"])
            return
        
    
    # ===================== Actions Block =================== #
    
    
    def print_do(self,msg):
        
        print(msg)
    
    
    # Unimplemented Actions
    
    
    # =============== Machinery and Mechanisms ============== #
    
    def transition(self, compartment):
        self.next_compartment = compartment
    
    def do_transition(self, next_compartment):
        self.mux(FrameEvent("<", self.compartment.exit_args))
        self.compartment = next_compartment
        self.mux(FrameEvent(">", self.compartment.enter_args))
    

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
	#def __init__(self,stateMsg,enterMsg):
	    #super().__init__(stateMsg,enterMsg)

# ********************


