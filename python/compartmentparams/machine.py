# emitted from framec_v0.10.0
# get include files at https://github.com/frame-lang/frame-ancillary-files

from framelang.framelang import FrameEvent

class CompartmentParams:
    
    def __init__(self, state_param,enter_param):
        
        # Create and intialize start state compartment.
        self.state = self._sS0_
        self.compartment = CompartmentParamsCompartment(self.state)
        self.next_compartment = None
        self.compartment.state_args["state_param"] = state_param
        self.compartment.state_vars["state_var"] = 100
        self.compartment.enter_args["enter_param"] = enter_param
        
        # Initialize domain
        
        # Send system start event
        frame_event = FrameEvent(">", self.compartment.enter_args)
        self.mux(frame_event)
    
    # ====================== Multiplexer ==================== #
    
    def mux(self, e):
        if self.compartment.state == self._sS0_:
            self._sS0_(e)
        elif self.compartment.state == self._sS1_:
            self._sS1_(e)
        
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
            self.print_do(str((self.compartment.state_args["state_param"])) + " " + str((self.compartment.state_vars["state_var"])) + " " + str(e._parameters["enter_param"]))
            
            compartment = CompartmentParamsCompartment(self._sS1_)
            compartment.forward_event = e
            compartment.state_args["state_param"] = self.compartment.state_args["state_param"] + 20
            compartment.state_vars["state_var"] = 200
            
            self.transition(compartment)
            return
        
    def _sS1_(self, e):
        if e._message == ">":
            self.print_do(str((self.compartment.state_args["state_param"])) + " " + str((self.compartment.state_vars["state_var"])) + " " + str(e._parameters["enter_param"]))
            return
        
    
    # ===================== Actions Block =================== #
    
    def print_do(self,s):
        raise NotImplementedError
    
    
    # =============== Machinery and Mechanisms ============== #
    
    def transition(self, compartment):
        self.next_compartment = compartment
    
    def do_transition(self, next_compartment):
        self.mux(FrameEvent("<", self.compartment.exit_args))
        self.compartment = next_compartment
        self.mux(FrameEvent(">", self.compartment.enter_args))
    

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

    #def print_do(self,s):
        #pass

# ********************


