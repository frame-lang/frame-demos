# emitted from framec_v0.10.0
# get include files at https://github.com/frame-lang/frame-ancillary-files
from framelang.framelang import FrameEvent

class FizzBuzz:
    
    def __init__(self):
        
        # Create and intialize start state compartment.
        self.__state = self.__fizzbuzz_state_Begin
        self.__compartment: 'FizzBuzzCompartment' = FizzBuzzCompartment(self.__state)
        self.__next_compartment: 'FizzBuzzCompartment' = None
        
        # Initialize domain
        
        # Send system start event
        frame_event = FrameEvent(">", None)
        self.__mux(frame_event)
    
    # ===================== Interface Block =================== #
    
    def start(self,):
        e = FrameEvent(">>",None)
        self.__mux(e)
    
    # ====================== Multiplexer ==================== #
    
    def __mux(self, e):
        if self.__compartment.state == self.__fizzbuzz_state_Begin:
            self.__fizzbuzz_state_Begin(e)
        elif self.__compartment.state == self.__fizzbuzz_state_Fizz:
            self.__fizzbuzz_state_Fizz(e)
        elif self.__compartment.state == self.__fizzbuzz_state_Buzz:
            self.__fizzbuzz_state_Buzz(e)
        elif self.__compartment.state == self.__fizzbuzz_state_Digit:
            self.__fizzbuzz_state_Digit(e)
        elif self.__compartment.state == self.__fizzbuzz_state_End:
            self.__fizzbuzz_state_End(e)
        
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
    
    def __fizzbuzz_state_Begin(self, e):
        if e._message == ">>":
            # start
            compartment = FizzBuzzCompartment(self.__fizzbuzz_state_Fizz)
            compartment.enter_args["i"] = 1
            self.__transition(compartment)
            return
        
    def __fizzbuzz_state_Fizz(self, e):
        if e._message == ">":
            if  self.gt_100_do(e._parameters["i"]):
                # i > 100
                compartment = FizzBuzzCompartment(self.__fizzbuzz_state_End)
                self.__transition(compartment)
                return
            
            if  self.mod3_eq0_do(e._parameters["i"]):
                self.print_do("Fizz")
                # i % 3 == 0
                compartment = FizzBuzzCompartment(self.__fizzbuzz_state_Buzz)
                compartment.enter_args["i"] = e._parameters["i"]
                compartment.enter_args["fizzed"] = True
                self.__transition(compartment)
            else:
                # i % 3 != 0
                compartment = FizzBuzzCompartment(self.__fizzbuzz_state_Buzz)
                compartment.enter_args["i"] = e._parameters["i"]
                compartment.enter_args["fizzed"] = False
                self.__transition(compartment)
            
            return
        
    def __fizzbuzz_state_Buzz(self, e):
        if e._message == ">":
            if  self.mod5_eq0_do(e._parameters["i"]):
                self.print_do("Buzz")
                # i % 5 == 0
                self.__compartment.exit_args["output"] = " "
                compartment = FizzBuzzCompartment(self.__fizzbuzz_state_Fizz)
                compartment.enter_args["i"] = self.plus_1_do(e._parameters["i"])
                self.__transition(compartment)
                return
            
            if  e._parameters["fizzed"]:
                # fizzed
                self.__compartment.exit_args["output"] = " "
                compartment = FizzBuzzCompartment(self.__fizzbuzz_state_Fizz)
                compartment.enter_args["i"] = self.plus_1_do(e._parameters["i"])
                self.__transition(compartment)
                return
            
            # ! mod3 | mod5
            self.__compartment.exit_args["output"] = ""
            compartment = FizzBuzzCompartment(self.__fizzbuzz_state_Digit)
            compartment.enter_args["i"] = e._parameters["i"]
            self.__transition(compartment)
            return
        
        elif e._message == "<":
            self.print_do(e._parameters["output"])
            return
        
    def __fizzbuzz_state_Digit(self, e):
        if e._message == ">":
            self.print_do(str(e._parameters["i"]))
            self.print_do(" ")
            # loop
            compartment = FizzBuzzCompartment(self.__fizzbuzz_state_Fizz)
            compartment.enter_args["i"] = self.plus_1_do(e._parameters["i"])
            self.__transition(compartment)
            return
        
    def __fizzbuzz_state_End(self, e):
        pass
        
    
    # ===================== Actions Block =================== #
    
    
    # Unimplemented Actions
    
    def print_do(self,msg: str):
        raise NotImplementedError
    
    def gt_100_do(self,i: int):
        raise NotImplementedError
    
    def mod3_eq0_do(self,i: int):
        raise NotImplementedError
    
    def mod5_eq0_do(self,i: int):
        raise NotImplementedError
    
    def plus_1_do(self,i: int):
        raise NotImplementedError
    
    
    # =============== Machinery and Mechanisms ============== #
    
    def __transition(self, compartment: 'FizzBuzzCompartment'):
        self.__next_compartment = compartment
    
    def  __do_transition(self, next_compartment: 'FizzBuzzCompartment'):
        self.__mux(FrameEvent("<", self.__compartment.exit_args))
        self.__compartment = next_compartment
        self.__mux(FrameEvent(">", self.__compartment.enter_args))
    

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

    #def print_do(self,msg: str):
        #pass

    #def gt_100_do(self,i: int):
        #pass

    #def mod3_eq0_do(self,i: int):
        #pass

    #def mod5_eq0_do(self,i: int):
        #pass

    #def plus_1_do(self,i: int):
        #pass

# ********************


