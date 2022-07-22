# emitted from framec_v0.10.0
# get include files at https://github.com/frame-lang/frame-ancillary-files
import copy
from framelang.framelang import FrameEvent

class TrafficLight:
    
    def __init__(self, manager):
        
        # Create and intialize start state compartment.
        self.__state = self.__trafficlight_state_Begin
        self._manager = manager
        self.__compartment: 'TrafficLightCompartment' = TrafficLightCompartment(self.__state)
        self.__next_compartment: 'TrafficLightCompartment' = None
        
        # Initialize domain
        
        self.flashColor : str = ""
        
        # Send system start event
        frame_event = FrameEvent(">", None)
        self.__mux(frame_event)
    
    # ===================== Interface Block =================== #
    
    def stop(self,):
        e = FrameEvent("stop",None)
        self.__mux(e)
    
    def tick(self,):
        e = FrameEvent("tick",None)
        self.__mux(e)
    
    def systemError(self,):
        e = FrameEvent("systemError",None)
        self.__mux(e)
    
    def systemRestart(self,):
        e = FrameEvent("systemRestart",None)
        self.__mux(e)
    
    
    
    @staticmethod
    def loadTrafficLight(manager, data):
        
        data._manager = manager
        
        return data
        
    
    
    def marshal(self):
        
        data = copy.deepcopy(self)
        return data
        
    
    # ====================== Multiplexer ==================== #
    
    def __mux(self, e):
        if self.__compartment.state == self.__trafficlight_state_Begin:
            self.__trafficlight_state_Begin(e)
        elif self.__compartment.state == self.__trafficlight_state_Red:
            self.__trafficlight_state_Red(e)
        elif self.__compartment.state == self.__trafficlight_state_Green:
            self.__trafficlight_state_Green(e)
        elif self.__compartment.state == self.__trafficlight_state_Yellow:
            self.__trafficlight_state_Yellow(e)
        elif self.__compartment.state == self.__trafficlight_state_FlashingRed:
            self.__trafficlight_state_FlashingRed(e)
        elif self.__compartment.state == self.__trafficlight_state_End:
            self.__trafficlight_state_End(e)
        elif self.__compartment.state == self.__trafficlight_state_Working:
            self.__trafficlight_state_Working(e)
        
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
    
    def __trafficlight_state_Begin(self, e):
        if e._message == ">":
            self.startWorkingTimer_do()
            compartment = TrafficLightCompartment(self.__trafficlight_state_Red)
            self.__transition(compartment)
            return
        
    def __trafficlight_state_Red(self, e):
        if e._message == ">":
            self.enterRed_do()
            return
        
        elif e._message == "tick":
            compartment = TrafficLightCompartment(self.__trafficlight_state_Green)
            self.__transition(compartment)
            return
        
        self.__trafficlight_state_Working(e)
        
    def __trafficlight_state_Green(self, e):
        if e._message == ">":
            self.enterGreen_do()
            return
        
        elif e._message == "tick":
            compartment = TrafficLightCompartment(self.__trafficlight_state_Yellow)
            self.__transition(compartment)
            return
        
        self.__trafficlight_state_Working(e)
        
    def __trafficlight_state_Yellow(self, e):
        if e._message == ">":
            self.enterYellow_do()
            return
        
        elif e._message == "tick":
            compartment = TrafficLightCompartment(self.__trafficlight_state_Red)
            self.__transition(compartment)
            return
        
        self.__trafficlight_state_Working(e)
        
    def __trafficlight_state_FlashingRed(self, e):
        if e._message == ">":
            self.enterFlashingRed_do()
            self.stopWorkingTimer_do()
            self.startFlashingTimer_do()
            return
        
        elif e._message == "<":
            self.exitFlashingRed_do()
            self.stopFlashingTimer_do()
            self.startWorkingTimer_do()
            return
        
        elif e._message == "tick":
            self.changeFlashingAnimation_do()
            return
        
        elif e._message == "systemRestart":
            compartment = TrafficLightCompartment(self.__trafficlight_state_Red)
            self.__transition(compartment)
            return
        
        elif e._message == "stop":
            compartment = TrafficLightCompartment(self.__trafficlight_state_End)
            self.__transition(compartment)
            return
        
    def __trafficlight_state_End(self, e):
        if e._message == ">":
            self.stopWorkingTimer_do()
            return
        
    def __trafficlight_state_Working(self, e):
        if e._message == "stop":
            compartment = TrafficLightCompartment(self.__trafficlight_state_End)
            self.__transition(compartment)
            return
        
        elif e._message == "systemError":
            compartment = TrafficLightCompartment(self.__trafficlight_state_FlashingRed)
            self.__transition(compartment)
            return
        
    
    # ===================== Actions Block =================== #
    
    
    # Unimplemented Actions
    
    def enterRed_do(self):
        raise NotImplementedError
    
    def enterGreen_do(self):
        raise NotImplementedError
    
    def enterYellow_do(self):
        raise NotImplementedError
    
    def enterFlashingRed_do(self):
        raise NotImplementedError
    
    def exitFlashingRed_do(self):
        raise NotImplementedError
    
    def startWorkingTimer_do(self):
        raise NotImplementedError
    
    def stopWorkingTimer_do(self):
        raise NotImplementedError
    
    def startFlashingTimer_do(self):
        raise NotImplementedError
    
    def stopFlashingTimer_do(self):
        raise NotImplementedError
    
    def changeColor_do(self,color: str):
        raise NotImplementedError
    
    def startFlashing_do(self):
        raise NotImplementedError
    
    def stopFlashing_do(self):
        raise NotImplementedError
    
    def changeFlashingAnimation_do(self):
        raise NotImplementedError
    
    def log_do(self,msg: str):
        raise NotImplementedError
    
    
    # =============== Machinery and Mechanisms ============== #
    
    def __transition(self, compartment: 'TrafficLightCompartment'):
        self.__next_compartment = compartment
    
    def  __do_transition(self, next_compartment: 'TrafficLightCompartment'):
        self.__mux(FrameEvent("<", self.__compartment.exit_args))
        self.__compartment = next_compartment
        self.__mux(FrameEvent(">", self.__compartment.enter_args))
    

# ===================== Compartment =================== #

class TrafficLightCompartment:

    def __init__(self, state):
        self.state = state
        self.state_args = {}
        self.state_vars = {}
        self.enter_args = {}
        self.exit_args = {}
        self.forward_event = FrameEvent(None, None)
    


# ********************

class TrafficLightController(TrafficLight):
    def __init__(self, manager):
        super().__init__(manager)

    def enterRed_do(self):
        self._manager.enterRed()
    def enterGreen_do(self):
        self._manager.enterGreen()

    def enterYellow_do(self):
        self._manager.enterYellow()

    def enterFlashingRed_do(self):
        pass

    def exitFlashingRed_do(self):
        pass

    def startWorkingTimer_do(self):
        pass

    def stopWorkingTimer_do(self):
        pass

    def startFlashingTimer_do(self):
        pass

    def stopFlashingTimer_do(self):
        pass

    def changeColor_do(self,color):
        pass

    def startFlashing_do(self):
        pass

    def stopFlashing_do(self):
        pass

    def changeFlashingAnimation_do(self):
        pass

    def log_do(self,msg):
        pass

# ********************


