# emitted from framec_v0.10.0
# get include files at https://github.com/frame-lang/frame-ancillary-files

from framelang.framelang import FrameEvent


class StringTools:
    
    def __init__(self, ):
        
        # Create and intialize start state compartment.
        self.state = self._sRouter_
        self.compartment = StringToolsCompartment(self.state)
        self.next_compartment = None
        
        # Initialize domain
        
        # Send system start event
        frame_event = FrameEvent(">", None)
        self.mux(frame_event)
    
    # ===================== Interface Block =================== #
    
    def reverse(self,str):
        parameters = {}
        parameters["str"] = str

        e = FrameEvent("reverse",parameters)
        self.mux(e)
        return e._return
    
    def makePalindrome(self,str):
        parameters = {}
        parameters["str"] = str

        e = FrameEvent("makePalindrome",parameters)
        self.mux(e)
        return e._return
    
    # ====================== Multiplexer ==================== #
    
    def mux(self, e):
        if self.compartment.state == self._sRouter_:
            self._sRouter_(e)
        elif self.compartment.state == self._sReverse_:
            self._sReverse_(e)
        elif self.compartment.state == self._sMakePalindrome_:
            self._sMakePalindrome_(e)
        
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
    
    def _sRouter_(self, e):
        if e._message == "makePalindrome":
            # make\npalindrome
            compartment = StringToolsCompartment(self._sMakePalindrome_)
            compartment.forward_event = e
            self.transition(compartment)
            return
        
        elif e._message == "reverse":
            # reverse
            compartment = StringToolsCompartment(self._sReverse_)
            compartment.forward_event = e
            self.transition(compartment)
            return
        
    def _sReverse_(self, e):
        if e._message == "reverse":
            e._return = self.reverse_str_do(e._parameters["str"])
            # ready
            compartment = StringToolsCompartment(self._sRouter_)
            self.transition(compartment)
            return
        
    def _sMakePalindrome_(self, e):
        if e._message == "makePalindrome":
            e._return = e._parameters["str"] + self.reverse_str_do(e._parameters["str"])
            # ready
            compartment = StringToolsCompartment(self._sRouter_)
            self.transition(compartment)
            return
        
    
    # ===================== Actions Block =================== #
    
    
    # Unimplemented Actions
    
    def reverse_str_do(self,str):
        raise NotImplementedError
    
    
    # =============== Machinery and Mechanisms ============== #
    
    def transition(self, compartment):
        self.next_compartment = compartment
    
    def do_transition(self, next_compartment):
        self.mux(FrameEvent("<", self.compartment.exit_args))
        self.compartment = next_compartment
        self.mux(FrameEvent(">", self.compartment.enter_args))
    

# ===================== Compartment =================== #

class StringToolsCompartment:

    def __init__(self, state):
        self.state = state
        self.state_args = {}
        self.state_vars = {}
        self.enter_args = {}
        self.exit_args = {}
        self.forward_event = FrameEvent(None, None)
    


# ********************

#class StringToolsController(StringTools):
	#def __init__(self,):
	    #super().__init__()

    #def reverse_str_do(self,str):
        #pass

# ********************


