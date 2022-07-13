# emitted from framec_v0.10.0
# get include files at https://github.com/frame-lang/frame-ancillary-files

from framelang.framelang import FrameEvent

class FizzBuzz:
    
    def __init__(self, ):
        
        # Create and intialize start state compartment.
        self.state = self._sBegin_
        self.compartment = FizzBuzzCompartment(self.state)
        self.next_compartment = None
        
        # Initialize domain
        
        # Send system start event
        frame_event = FrameEvent(">", None)
        self.mux(frame_event)
    
    # ===================== Interface Block =================== #
    
    def start(self):
        e = FrameEvent(">>",None)
        self.mux(e)
    
    # ====================== Multiplexer ==================== #
    
    def mux(self, e):
        if self.compartment.state == self._sBegin_:
            self._sBegin_(e)
        elif self.compartment.state == self._sFizz_:
            self._sFizz_(e)
        elif self.compartment.state == self._sBuzz_:
            self._sBuzz_(e)
        elif self.compartment.state == self._sDigit_:
            self._sDigit_(e)
        elif self.compartment.state == self._sEnd_:
            self._sEnd_(e)
        
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
        if e._message == ">>":
            # start
            compartment = FizzBuzzCompartment(self._sFizz_)
            compartment.enter_args["i"] = 1
            self.transition(compartment)
            return
        
    def _sFizz_(self, e):
        if e._message == ">":
            if  self.gt_100_do(e._parameters["i"]):
                # i > 100
                compartment = FizzBuzzCompartment(self._sEnd_)
                self.transition(compartment)
                return
            
            if  self.mod3_eq0_do(e._parameters["i"]):
                self.print_do("Fizz")
                # i % 3 == 0
                compartment = FizzBuzzCompartment(self._sBuzz_)
                compartment.enter_args["i"] = e._parameters["i"]
                compartment.enter_args["fizzed"] = True
                self.transition(compartment)
            else:
                # i % 3 != 0
                compartment = FizzBuzzCompartment(self._sBuzz_)
                compartment.enter_args["i"] = e._parameters["i"]
                compartment.enter_args["fizzed"] = False
                self.transition(compartment)
            
            return
        
    def _sBuzz_(self, e):
        if e._message == ">":
            if  self.mod5_eq0_do(e._parameters["i"]):
                self.print_do("Buzz")
                # i % 5 == 0
                self.compartment.exit_args["output"] = " "
                compartment = FizzBuzzCompartment(self._sFizz_)
                compartment.enter_args["i"] = self.plus_1_do(e._parameters["i"])
                self.transition(compartment)
                return
            
            if  e._parameters["fizzed"]:
                # fizzed
                self.compartment.exit_args["output"] = " "
                compartment = FizzBuzzCompartment(self._sFizz_)
                compartment.enter_args["i"] = self.plus_1_do(e._parameters["i"])
                self.transition(compartment)
                return
            
            # ! mod3 | mod5
            self.compartment.exit_args["output"] = ""
            compartment = FizzBuzzCompartment(self._sDigit_)
            compartment.enter_args["i"] = e._parameters["i"]
            self.transition(compartment)
            return
        
        elif e._message == "<":
            self.print_do(e._parameters["output"])
            return
        
    def _sDigit_(self, e):
        if e._message == ">":
            self.print_do(str(e._parameters["i"]))
            self.print_do(" ")
            # loop
            compartment = FizzBuzzCompartment(self._sFizz_)
            compartment.enter_args["i"] = self.plus_1_do(e._parameters["i"])
            self.transition(compartment)
            return
        
    def _sEnd_(self, e):
        pass
        
    
    # ===================== Actions Block =================== #
    
    
    # Unimplemented Actions
    
    def print_do(self,msg):
        raise NotImplementedError
    
    def gt_100_do(self,i):
        raise NotImplementedError
    
    def mod3_eq0_do(self,i):
        raise NotImplementedError
    
    def mod5_eq0_do(self,i):
        raise NotImplementedError
    
    def plus_1_do(self,i):
        raise NotImplementedError
    
    
    # =============== Machinery and Mechanisms ============== #
    
    def transition(self, compartment):
        self.next_compartment = compartment
    
    def do_transition(self, next_compartment):
        self.mux(FrameEvent("<", self.compartment.exit_args))
        self.compartment = next_compartment
        self.mux(FrameEvent(">", self.compartment.enter_args))
    

# ===================== Compartment =================== #

class FizzBuzzCompartment:

    def __init__(self, state):
        self.state = state
        self.state_args = {}
        self.state_vars = {}
        self.enter_args = {}
        self.exit_args = {}
        self.forward_event = FrameEvent(None, None)
    


# ********************

#class FizzBuzzController(FizzBuzz):
	#def __init__(self,):
	    #super().__init__()

    #def print_do(self,msg):
        #pass

    #def gt_100_do(self,i):
        #pass

    #def mod3_eq0_do(self,i):
        #pass

    #def mod5_eq0_do(self,i):
        #pass

    #def plus_1_do(self,i):
        #pass

# ********************


