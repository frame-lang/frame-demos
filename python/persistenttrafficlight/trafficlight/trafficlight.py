# emitted from framec_v0.10.0
# get include files at https://github.com/frame-lang/frame-ancillary-files
import copy
from framelang.framelang import FrameEvent

class TrafficLight:
    
    def __init__(self, manager):
        
        # Create and intialize start state compartment.
        self.state = self._sBegin_
        self._manager = manager
        self.compartment = TrafficLightCompartment(self.state)
        self.next_compartment = None
        
        # Initialize domain
        
        self.flashColor = ""
        
        # Send system start event
        frame_event = FrameEvent(">", None)
        self.mux(frame_event)
    
    # ===================== Interface Block =================== #
    
    def stop(self):
        e = FrameEvent("stop",None)
        self.mux(e)
    
    def tick(self):
        e = FrameEvent("tick",None)
        self.mux(e)
    
    def systemError(self):
        e = FrameEvent("systemError",None)
        self.mux(e)
    
    def systemRestart(self):
        e = FrameEvent("systemRestart",None)
        self.mux(e)
    
    
    
    @staticmethod
    def loadTrafficLight(manager, data):
        
        data._manager = manager
        
        return data
        
    
    
    def marshal(self):
        
        data = copy.deepcopy(self)
        return data
        
    
    # ====================== Multiplexer ==================== #
    
    def mux(self, e):
        if self.compartment.state == self._sBegin_:
            self._sBegin_(e)
        elif self.compartment.state == self._sRed_:
            self._sRed_(e)
        elif self.compartment.state == self._sGreen_:
            self._sGreen_(e)
        elif self.compartment.state == self._sYellow_:
            self._sYellow_(e)
        elif self.compartment.state == self._sFlashingRed_:
            self._sFlashingRed_(e)
        elif self.compartment.state == self._sEnd_:
            self._sEnd_(e)
        elif self.compartment.state == self._sWorking_:
            self._sWorking_(e)
        
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
            self.startWorkingTimer_do()
            compartment = TrafficLightCompartment(self._sRed_)
            self.transition(compartment)
            return
        
    def _sRed_(self, e):
        if e._message == ">":
            self.enterRed_do()
            return
        
        elif e._message == "tick":
            compartment = TrafficLightCompartment(self._sGreen_)
            self.transition(compartment)
            return
        
        self._sWorking_(e)
        
    def _sGreen_(self, e):
        if e._message == ">":
            self.enterGreen_do()
            return
        
        elif e._message == "tick":
            compartment = TrafficLightCompartment(self._sYellow_)
            self.transition(compartment)
            return
        
        self._sWorking_(e)
        
    def _sYellow_(self, e):
        if e._message == ">":
            self.enterYellow_do()
            return
        
        elif e._message == "tick":
            compartment = TrafficLightCompartment(self._sRed_)
            self.transition(compartment)
            return
        
        self._sWorking_(e)
        
    def _sFlashingRed_(self, e):
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
            compartment = TrafficLightCompartment(self._sRed_)
            self.transition(compartment)
            return
        
        elif e._message == "stop":
            compartment = TrafficLightCompartment(self._sEnd_)
            self.transition(compartment)
            return
        
    def _sEnd_(self, e):
        if e._message == ">":
            self.stopWorkingTimer_do()
            return
        
    def _sWorking_(self, e):
        if e._message == "stop":
            compartment = TrafficLightCompartment(self._sEnd_)
            self.transition(compartment)
            return
        
        elif e._message == "systemError":
            compartment = TrafficLightCompartment(self._sFlashingRed_)
            self.transition(compartment)
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
    
    def changeColor_do(self,color):
        raise NotImplementedError
    
    def startFlashing_do(self):
        raise NotImplementedError
    
    def stopFlashing_do(self):
        raise NotImplementedError
    
    def changeFlashingAnimation_do(self):
        raise NotImplementedError
    
    def log_do(self,msg):
        raise NotImplementedError
    
    
    # =============== Machinery and Mechanisms ============== #
    
    def transition(self, compartment):
        self.next_compartment = compartment
    
    def do_transition(self, next_compartment):
        self.mux(FrameEvent("<", self.compartment.exit_args))
        self.compartment = next_compartment
        self.mux(FrameEvent(">", self.compartment.enter_args))
    

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


