# emitted from framec_v0.10.0
# get include files at https://github.com/frame-lang/frame-ancillary-files
from framelang.framelang import FrameEvent

class HistoryStateContext:
    
    def __init__(self):
        
        # Create state stack.
        
        self.__state_stack = []
        
        # Create and intialize start state compartment.
        self.__state = self.__historystatecontext_state_Start
        self.__compartment: 'HistoryStateContextCompartment' = HistoryStateContextCompartment(self.__state)
        self.__next_compartment: 'HistoryStateContextCompartment' = None
        
        # Initialize domain
        
        # Send system start event
        frame_event = FrameEvent(">", None)
        self.__mux(frame_event)
    
    # ===================== Interface Block =================== #
    
    def start(self,):
        e = FrameEvent(">>",None)
        self.__mux(e)
    
    def switchState(self,):
        e = FrameEvent("switchState",None)
        self.__mux(e)
    
    def gotoDeadEnd(self,):
        e = FrameEvent("gotoDeadEnd",None)
        self.__mux(e)
    
    def back(self,):
        e = FrameEvent("back",None)
        self.__mux(e)
    
    # ====================== Multiplexer ==================== #
    
    def __mux(self, e):
        if self.__compartment.state == self.__historystatecontext_state_Start:
            self.__historystatecontext_state_Start(e)
        elif self.__compartment.state == self.__historystatecontext_state_S0:
            self.__historystatecontext_state_S0(e)
        elif self.__compartment.state == self.__historystatecontext_state_S1:
            self.__historystatecontext_state_S1(e)
        elif self.__compartment.state == self.__historystatecontext_state_DeadEnd:
            self.__historystatecontext_state_DeadEnd(e)
        
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
    
    def __historystatecontext_state_Start(self, e):
        if e._message == ">>":
            compartment = HistoryStateContextCompartment(self.__historystatecontext_state_S0)
            compartment.state_vars["enterMsg"] = "Enter S0"
            self.__transition(compartment)
            return
        
    def __historystatecontext_state_S0(self, e):
        if e._message == ">":
            self.print_do((self.__compartment.state_vars["enterMsg"]))
            return
        
        elif e._message == "switchState":
            # Switch\nState
            compartment = HistoryStateContextCompartment(self.__historystatecontext_state_S1)
            self.__transition(compartment)
            return
        
        elif e._message == "gotoDeadEnd":
            self.__state_stack_push(self.__compartment)
            # Goto\nDead End
            compartment = HistoryStateContextCompartment(self.__historystatecontext_state_DeadEnd)
            self.__transition(compartment)
            return
        
    def __historystatecontext_state_S1(self, e):
        if e._message == ">":
            self.print_do("Enter S1")
            return
        
        elif e._message == "switchState":
            # Switch\nState
            compartment = HistoryStateContextCompartment(self.__historystatecontext_state_S0)
            compartment.state_vars["enterMsg"] = "Enter S0"
            self.__transition(compartment)
            return
        
        elif e._message == "gotoDeadEnd":
            self.__state_stack_push(self.__compartment)
            # Goto\nDead End
            compartment = HistoryStateContextCompartment(self.__historystatecontext_state_DeadEnd)
            self.__transition(compartment)
            return
        
    def __historystatecontext_state_DeadEnd(self, e):
        if e._message == ">":
            self.print_do("Enter $DeadEnd")
            return
        
        elif e._message == "back":
            # Go Back
            compartment = self.__state_stack_pop()
            self.__transition(compartment)
            return
        
    
    # ===================== Actions Block =================== #
    
    
    # Unimplemented Actions
    
    def print_do(self,msg: str):
        raise NotImplementedError
    
    
    # =============== Machinery and Mechanisms ============== #
    
    def __transition(self, compartment: 'HistoryStateContextCompartment'):
        self.__next_compartment = compartment
    
    def  __do_transition(self, next_compartment: 'HistoryStateContextCompartment'):
        self.__mux(FrameEvent("<", self.__compartment.exit_args))
        self.__compartment = next_compartment
        self.__mux(FrameEvent(">", self.__compartment.enter_args))
    
    def __state_stack_push(self, compartment: 'HistoryStateContextCompartment'):
        self.__state_stack.append(compartment)
    
    def __state_stack_pop(self):
        return self.__state_stack.pop()
    
    

# ===================== Compartment =================== #

class HistoryStateContextCompartment:

    def __init__(self, state):
        self.state = state
        self.state_args = {}
        self.state_vars = {}
        self.enter_args = {}
        self.exit_args = {}
        self.forward_event = FrameEvent(None, None)
    


# ********************

#class HistoryStateContextController(HistoryStateContext):
	#def __init__(self,):
	    #super().__init__()

    #def print_do(self,msg: str):
        #pass

# ********************


