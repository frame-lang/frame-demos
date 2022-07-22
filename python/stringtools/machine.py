# emitted from framec_v0.10.0
# get include files at https://github.com/frame-lang/frame-ancillary-files
from framelang.framelang import FrameEvent

class StringTools:
    
    def __init__(self):
        
        # Create and intialize start state compartment.
        self.__state = self.__stringtools_state_Router
        self.__compartment: 'StringToolsCompartment' = StringToolsCompartment(self.__state)
        self.__next_compartment: 'StringToolsCompartment' = None
        
        # Initialize domain
        
        # Send system start event
        frame_event = FrameEvent(">", None)
        self.__mux(frame_event)
    
    # ===================== Interface Block =================== #
    
    def reverse(self,str: str):
        parameters = {}
        parameters["str"] = str

        e = FrameEvent("reverse",parameters)
        self.__mux(e)
        return e._return
    
    def makePalindrome(self,str: str):
        parameters = {}
        parameters["str"] = str

        e = FrameEvent("makePalindrome",parameters)
        self.__mux(e)
        return e._return
    
    # ====================== Multiplexer ==================== #
    
    def __mux(self, e):
        if self.__compartment.state == self.__stringtools_state_Router:
            self.__stringtools_state_Router(e)
        elif self.__compartment.state == self.__stringtools_state_Reverse:
            self.__stringtools_state_Reverse(e)
        elif self.__compartment.state == self.__stringtools_state_MakePalindrome:
            self.__stringtools_state_MakePalindrome(e)
        
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
    
    def __stringtools_state_Router(self, e):
        if e._message == "makePalindrome":
            # make\npalindrome
            compartment = StringToolsCompartment(self.__stringtools_state_MakePalindrome)
            compartment.forward_event = e
            self.__transition(compartment)
            return
        
        elif e._message == "reverse":
            # reverse
            compartment = StringToolsCompartment(self.__stringtools_state_Reverse)
            compartment.forward_event = e
            self.__transition(compartment)
            return
        
    def __stringtools_state_Reverse(self, e):
        if e._message == "reverse":
            e._return = self.reverse_str_do(e._parameters["str"])
            # ready
            compartment = StringToolsCompartment(self.__stringtools_state_Router)
            self.__transition(compartment)
            return
        
    def __stringtools_state_MakePalindrome(self, e):
        if e._message == "makePalindrome":
            e._return = e._parameters["str"] + self.reverse_str_do(e._parameters["str"])
            # ready
            compartment = StringToolsCompartment(self.__stringtools_state_Router)
            self.__transition(compartment)
            return
        
    
    # ===================== Actions Block =================== #
    
    
    # Unimplemented Actions
    
    def reverse_str_do(self,str: str):
        raise NotImplementedError
    
    
    # =============== Machinery and Mechanisms ============== #
    
    def __transition(self, compartment: 'StringToolsCompartment'):
        self.__next_compartment = compartment
    
    def  __do_transition(self, next_compartment: 'StringToolsCompartment'):
        self.__mux(FrameEvent("<", self.__compartment.exit_args))
        self.__compartment = next_compartment
        self.__mux(FrameEvent(">", self.__compartment.enter_args))
    

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

    #def reverse_str_do(self,str: str):
        #pass

# ********************


