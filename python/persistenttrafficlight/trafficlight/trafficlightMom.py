# emitted from framec_v0.10.0
# get include files at https://github.com/frame-lang/frame-ancillary-files
from framelang.framelang import FrameEvent
from trafficlight.trafficlight import TrafficLightController

class TrafficLightManager:
    
    def __init__(self):
        
        # Create and intialize start state compartment.
        self.__state = self.__trafficlightmanager_state_New
        self.__compartment: 'TrafficLightManagerCompartment' = TrafficLightManagerCompartment(self.__state)
        self.__next_compartment: 'TrafficLightManagerCompartment' = None
        
        # Initialize domain
        
        self.trafficLight  = None
        self.data  = None
        
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
    
    def enterRed(self,):
        e = FrameEvent("enterRed",None)
        self.__mux(e)
    
    def enterGreen(self,):
        e = FrameEvent("enterGreen",None)
        self.__mux(e)
    
    def enterYellow(self,):
        e = FrameEvent("enterYellow",None)
        self.__mux(e)
    
    def enterFlashingRed(self,):
        e = FrameEvent("enterFlashingRed",None)
        self.__mux(e)
    
    def exitFlashingRed(self,):
        e = FrameEvent("exitFlashingRed",None)
        self.__mux(e)
    
    def startWorkingTimer(self,):
        e = FrameEvent("startWorkingTimer",None)
        self.__mux(e)
    
    def stopWorkingTimer(self,):
        e = FrameEvent("stopWorkingTimer",None)
        self.__mux(e)
    
    def startFlashingTimer(self,):
        e = FrameEvent("startFlashingTimer",None)
        self.__mux(e)
    
    def stopFlashingTimer(self,):
        e = FrameEvent("stopFlashingTimer",None)
        self.__mux(e)
    
    def changeColor(self,color: str):
        parameters = {}
        parameters["color"] = color

        e = FrameEvent("changeColor",parameters)
        self.__mux(e)
    
    def startFlashing(self,):
        e = FrameEvent("startFlashing",None)
        self.__mux(e)
    
    def stopFlashing(self,):
        e = FrameEvent("stopFlashing",None)
        self.__mux(e)
    
    def changeFlashingAnimation(self,):
        e = FrameEvent("changeFlashingAnimation",None)
        self.__mux(e)
    
    def systemError(self,):
        e = FrameEvent("systemError",None)
        self.__mux(e)
    
    def systemRestart(self,):
        e = FrameEvent("systemRestart",None)
        self.__mux(e)
    
    def log(self,msg: str):
        parameters = {}
        parameters["msg"] = msg

        e = FrameEvent("log",parameters)
        self.__mux(e)
    
    # ====================== Multiplexer ==================== #
    
    def __mux(self, e):
        if self.__compartment.state == self.__trafficlightmanager_state_New:
            self.__trafficlightmanager_state_New(e)
        elif self.__compartment.state == self.__trafficlightmanager_state_Saving:
            self.__trafficlightmanager_state_Saving(e)
        elif self.__compartment.state == self.__trafficlightmanager_state_Persisted:
            self.__trafficlightmanager_state_Persisted(e)
        elif self.__compartment.state == self.__trafficlightmanager_state_Working:
            self.__trafficlightmanager_state_Working(e)
        elif self.__compartment.state == self.__trafficlightmanager_state_TrafficLightApi:
            self.__trafficlightmanager_state_TrafficLightApi(e)
        elif self.__compartment.state == self.__trafficlightmanager_state_End:
            self.__trafficlightmanager_state_End(e)
        
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
    
    def __trafficlightmanager_state_New(self, e):
        if e._message == ">":
            self.trafficLight = TrafficLightController(self)
            # Traffic Light\nStarted
            compartment = TrafficLightManagerCompartment(self.__trafficlightmanager_state_Saving)
            self.__transition(compartment)
            return
        
        self.__trafficlightmanager_state_TrafficLightApi(e)
        
    def __trafficlightmanager_state_Saving(self, e):
        if e._message == ">":
            self.data = self.trafficLight.marshal()
            self.trafficLight = None
            # Saved
            compartment = TrafficLightManagerCompartment(self.__trafficlightmanager_state_Persisted)
            self.__transition(compartment)
            return
        
    def __trafficlightmanager_state_Persisted(self, e):
        if e._message == "tick":
            # Tick
            compartment = TrafficLightManagerCompartment(self.__trafficlightmanager_state_Working)
            compartment.forward_event = e
            self.__transition(compartment)
            return
        
        elif e._message == "systemError":
            # System Error
            compartment = TrafficLightManagerCompartment(self.__trafficlightmanager_state_Working)
            compartment.forward_event = e
            self.__transition(compartment)
            return
        
        elif e._message == "stop":
            # Stop
            compartment = TrafficLightManagerCompartment(self.__trafficlightmanager_state_End)
            self.__transition(compartment)
            return
        
    def __trafficlightmanager_state_Working(self, e):
        if e._message == ">":
            self.trafficLight = TrafficLightController.loadTrafficLight(self,self.data)
            return
        
        elif e._message == "tick":
            self.trafficLight.tick()
            # Done
            compartment = TrafficLightManagerCompartment(self.__trafficlightmanager_state_Saving)
            self.__transition(compartment)
            return
        
        elif e._message == "systemError":
            self.trafficLight.systemError_do()
            # Done
            compartment = TrafficLightManagerCompartment(self.__trafficlightmanager_state_Saving)
            self.__transition(compartment)
            return
        
        self.__trafficlightmanager_state_TrafficLightApi(e)
        
    def __trafficlightmanager_state_TrafficLightApi(self, e):
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
        
    def __trafficlightmanager_state_End(self, e):
        if e._message == ">":
            self.trafficLight = TrafficLightController.loadTrafficLight(self,self.data)
            self.trafficLight.stop()
            self.trafficLight = None
            return
        
        self.__trafficlightmanager_state_TrafficLightApi(e)
        
    
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
    
    def systemError_do(self):
        raise NotImplementedError
    
    def systemRestart_do(self):
        raise NotImplementedError
    
    def log_do(self,msg: str):
        raise NotImplementedError
    
    
    # =============== Machinery and Mechanisms ============== #
    
    def __transition(self, compartment: 'TrafficLightManagerCompartment'):
        self.__next_compartment = compartment
    
    def  __do_transition(self, next_compartment: 'TrafficLightManagerCompartment'):
        self.__mux(FrameEvent("<", self.__compartment.exit_args))
        self.__compartment = next_compartment
        self.__mux(FrameEvent(">", self.__compartment.enter_args))
    

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


