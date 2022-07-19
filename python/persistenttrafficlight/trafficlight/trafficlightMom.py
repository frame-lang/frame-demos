# emitted from framec_v0.10.0
# get include files at https://github.com/frame-lang/frame-ancillary-files
from trafficlight.trafficlight import TrafficLightController 

from framelang.framelang import FrameEvent


class TrafficLightManager:
    
    def __init__(self):
        
        # Create and intialize start state compartment.
        self.state = self._sNew_
        self.compartment = TrafficLightManagerCompartment(self.state)
        self.next_compartment = None
        
        # Initialize domain
        
        self.trafficLight = None
        self.data = None
        
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
    
    def enterRed(self):
        e = FrameEvent("enterRed",None)
        self.mux(e)
    
    def enterGreen(self):
        e = FrameEvent("enterGreen",None)
        self.mux(e)
    
    def enterYellow(self):
        e = FrameEvent("enterYellow",None)
        self.mux(e)
    
    def enterFlashingRed(self):
        e = FrameEvent("enterFlashingRed",None)
        self.mux(e)
    
    def exitFlashingRed(self):
        e = FrameEvent("exitFlashingRed",None)
        self.mux(e)
    
    def startWorkingTimer(self):
        e = FrameEvent("startWorkingTimer",None)
        self.mux(e)
    
    def stopWorkingTimer(self):
        e = FrameEvent("stopWorkingTimer",None)
        self.mux(e)
    
    def startFlashingTimer(self):
        e = FrameEvent("startFlashingTimer",None)
        self.mux(e)
    
    def stopFlashingTimer(self):
        e = FrameEvent("stopFlashingTimer",None)
        self.mux(e)
    
    def changeColor(self,color):
        parameters = {}
        parameters["color"] = color

        e = FrameEvent("changeColor",parameters)
        self.mux(e)
    
    def startFlashing(self):
        e = FrameEvent("startFlashing",None)
        self.mux(e)
    
    def stopFlashing(self):
        e = FrameEvent("stopFlashing",None)
        self.mux(e)
    
    def changeFlashingAnimation(self):
        e = FrameEvent("changeFlashingAnimation",None)
        self.mux(e)
    
    def systemError(self):
        e = FrameEvent("systemError",None)
        self.mux(e)
    
    def systemRestart(self):
        e = FrameEvent("systemRestart",None)
        self.mux(e)
    
    def log(self,msg):
        parameters = {}
        parameters["msg"] = msg

        e = FrameEvent("log",parameters)
        self.mux(e)
    
    # ====================== Multiplexer ==================== #
    
    def mux(self, e):
        if self.compartment.state == self._sNew_:
            self._sNew_(e)
        elif self.compartment.state == self._sSaving_:
            self._sSaving_(e)
        elif self.compartment.state == self._sPersisted_:
            self._sPersisted_(e)
        elif self.compartment.state == self._sWorking_:
            self._sWorking_(e)
        elif self.compartment.state == self._sTrafficLightApi_:
            self._sTrafficLightApi_(e)
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
    
    def _sNew_(self, e):
        if e._message == ">":
            self.trafficLight = TrafficLightController(self)
            # Traffic Light\nStarted
            compartment = TrafficLightManagerCompartment(self._sSaving_)
            self.transition(compartment)
            return
        
        self._sTrafficLightApi_(e)
        
    def _sSaving_(self, e):
        if e._message == ">":
            self.data = self.trafficLight.marshal()
            self.trafficLight = None
            # Saved
            compartment = TrafficLightManagerCompartment(self._sPersisted_)
            self.transition(compartment)
            return
        
    def _sPersisted_(self, e):
        if e._message == "tick":
            # Tick
            compartment = TrafficLightManagerCompartment(self._sWorking_)
            compartment.forward_event = e
            self.transition(compartment)
            return
        
        elif e._message == "systemError":
            # System Error
            compartment = TrafficLightManagerCompartment(self._sWorking_)
            compartment.forward_event = e
            self.transition(compartment)
            return
        
        elif e._message == "stop":
            # Stop
            compartment = TrafficLightManagerCompartment(self._sEnd_)
            self.transition(compartment)
            return
        
    def _sWorking_(self, e):
        if e._message == ">":
            self.trafficLight = TrafficLightController.loadTrafficLight(self,self.data)
            return
        
        elif e._message == "tick":
            self.trafficLight.tick()
            # Done
            compartment = TrafficLightManagerCompartment(self._sSaving_)
            self.transition(compartment)
            return
        
        elif e._message == "systemError":
            self.trafficLight.systemError_do()
            # Done
            compartment = TrafficLightManagerCompartment(self._sSaving_)
            self.transition(compartment)
            return
        
        self._sTrafficLightApi_(e)
        
    def _sTrafficLightApi_(self, e):
        if e._message == "enterRed":
            self.enterRed_do()
            return
        
        elif e._message == "enterGreen":
            self.enterGreen_do()
            return
        
        elif e._message == "enterYellow":
            self.enterYellow_do()
            return
        
        elif e._message == "enterFlashingRed":
            self.enterFlashingRed_do()
            return
        
        elif e._message == "exitFlashingRed":
            self.exitFlashingRed_do()
            return
        
        elif e._message == "startWorkingTimer":
            self.startWorkingTimer_do()
            return
        
        elif e._message == "stopWorkingTimer":
            self.stopWorkingTimer_do()
            return
        
        elif e._message == "startFlashingTimer":
            self.startFlashingTimer_do()
            return
        
        elif e._message == "stopFlashingTimer":
            self.stopFlashingTimer_do()
            return
        
        elif e._message == "changeColor":
            self.changeColor_do(e._parameters["color"])
            return
        
        elif e._message == "startFlashing":
            self.startFlashing_do()
            return
        
        elif e._message == "stopFlashing":
            self.stopFlashing_do()
            return
        
        elif e._message == "changeFlashingAnimation":
            self.changeFlashingAnimation_do()
            return
        
        elif e._message == "systemError":
            self.systemError_do()
            return
        
        elif e._message == "systemRestart":
            self.systemRestart_do()
            return
        
        elif e._message == "log":
            self.log_do(e._parameters["msg"])
            return
        
    def _sEnd_(self, e):
        if e._message == ">":
            self.trafficLight = TrafficLightController.loadTrafficLight(self,self.data)
            self.trafficLight.stop()
            self.trafficLight = None
            return
        
        self._sTrafficLightApi_(e)
        
    
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
    
    def systemError_do(self):
        raise NotImplementedError
    
    def systemRestart_do(self):
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

class TrafficLightManagerCompartment:

    def __init__(self, state):
        self.state = state
        self.state_args = {}
        self.state_vars = {}
        self.enter_args = {}
        self.exit_args = {}
        self.forward_event = FrameEvent(None, None)
    


# ********************

class TrafficLightManagerController(TrafficLightManager):
    def __init__(self,):
        super().__init__()
    
    def enterRed_do(self):
        print("Red")

    def enterGreen_do(self):
        print("Green")

    def enterYellow_do(self):
        print("Yellow")

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

    def systemError_do(self):
        pass

    def systemRestart_do(self):
        pass

    def log_do(self,msg):
        pass
    
    

# ********************


