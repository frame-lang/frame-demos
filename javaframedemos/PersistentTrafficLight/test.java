// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files
package java.PersistentTrafficLight;
import java.util.*;

import javaframedemos.PersistentTrafficLight.FrameEvent;

class TrafficLightController {
    
    private TrafficLightControllerCompartment _compartment_;
    private TrafficLightControllerCompartment _nextCompartment_;
    
    TrafficLightController() {
        
        // Create and intialize start state compartment.
        this._compartment_ = new TrafficLightControllerCompartment(this._state_);
        this._nextCompartment_ = null;
        
        
        // Send system start event
        FrameEvent frame_event = new FrameEvent(">", null);
        this._mux_(frame_event);
    
    }
    
    // states enum
    private enum TrafficLightControllerState {
        
        BEGIN(0),
        RED(1),
        GREEN(2),
        YELLOW(3),
        FLASHINGRED(4),
        END(5),
        WORKING(6);
        
        public final int value;
        
        private TrafficLightControllerState(int value) {
            this.value=value;
        }
        
        public int getValue() {
            return value;
        }
    }
    
    //====================== Multiplexer ====================//
    
    private void _mux_(FrameEvent e) {
        if(this._compartment_.state == TrafficLightControllerState.BEGIN.getValue()) {
            this._sBegin_(e);
        }else if(this._compartment_.state == TrafficLightControllerState.RED.getValue()) {
            this._sRed_(e);
        }else if(this._compartment_.state == TrafficLightControllerState.GREEN.getValue()) {
            this._sGreen_(e);
        }else if(this._compartment_.state == TrafficLightControllerState.YELLOW.getValue()) {
            this._sYellow_(e);
        }else if(this._compartment_.state == TrafficLightControllerState.FLASHINGRED.getValue()) {
            this._sFlashingRed_(e);
        }else if(this._compartment_.state == TrafficLightControllerState.END.getValue()) {
            this._sEnd_(e);
        }else if(this._compartment_.state == TrafficLightControllerState.WORKING.getValue()) {
            this._sWorking_(e);
        }
        
        
        if(this._nextCompartment_ != null) {
            TrafficLightControllerCompartment nextCompartment = this._nextCompartment_;
            this._nextCompartment_ = null;
            if(nextCompartment._forwardEvent != null && 
             nextCompartment._forwardEvent._message == ">") {
                this._mux_(new FrameEvent( "<", this._compartment_.exitArgs));
                this._compartment_ = nextCompartment;
                this._mux_(nextCompartment._forwardEvent);
            } else {
                this._doTransition_(nextCompartment);
                if(nextCompartment._forwardEvent != null) {
                    this._mux_(nextCompartment._forwardEvent);
                }
            }
            nextCompartment._forwardEvent = null;
        }
    }
    
    //===================== Interface Block ===================//
    
    public void stop() {
        FrameEvent e = new FrameEvent("stop", null);
        this._mux_(e);
    }
    
    public void tick() {
        FrameEvent e = new FrameEvent("tick", null);
        this._mux_(e);
    }
    
    public void systemError() {
        FrameEvent e = new FrameEvent("systemError", null);
        this._mux_(e);
    }
    
    public void systemRestart() {
        FrameEvent e = new FrameEvent("systemRestart", null);
        this._mux_(e);
    }
    
    public void loadTrafficLight(TrafficLightMomController manager,TrafficLightController data) {
        HashMap<String,Object> parameters = new HashMap<String,Object>();
        parameters.put("manager", manager);

        parameters.put("data", data);

        FrameEvent e = new FrameEvent("loadTrafficLight", parameters);
        this._mux_(e);
    }
    
    
    //===================== Machine Block ===================//
    
    private void _sBegin_(FrameEvent e) {
        if(e._message == ">") {
            startWorkingTimer_do();
            TrafficLightControllerCompartment compartment =  new TrafficLightControllerCompartment(TrafficLightControllerState.RED.getValue());
            
            this._transition_(compartment);
            return;
        }
    }
    
    private void _sRed_(FrameEvent e) {
        if(e._message == ">") {
            enterRed_do();
            return;
        }
        else if(e._message == "tick") {
            TrafficLightControllerCompartment compartment =  new TrafficLightControllerCompartment(TrafficLightControllerState.GREEN.getValue());
            
            this._transition_(compartment);
            return;
        }
        _sWorking_(e);
        
    }
    
    private void _sGreen_(FrameEvent e) {
        if(e._message == ">") {
            enterGreen_do();
            return;
        }
        else if(e._message == "tick") {
            TrafficLightControllerCompartment compartment =  new TrafficLightControllerCompartment(TrafficLightControllerState.YELLOW.getValue());
            
            this._transition_(compartment);
            return;
        }
        _sWorking_(e);
        
    }
    
    private void _sYellow_(FrameEvent e) {
        if(e._message == ">") {
            enterYellow_do();
            return;
        }
        else if(e._message == "tick") {
            TrafficLightControllerCompartment compartment =  new TrafficLightControllerCompartment(TrafficLightControllerState.RED.getValue());
            
            this._transition_(compartment);
            return;
        }
        _sWorking_(e);
        
    }
    
    private void _sFlashingRed_(FrameEvent e) {
        if(e._message == ">") {
            enterFlashingRed_do();
            stopWorkingTimer_do();
            startFlashingTimer_do();
            return;
        }
        else if(e._message == "<") {
            exitFlashingRed_do();
            stopFlashingTimer_do();
            startWorkingTimer_do();
            return;
        }
        else if(e._message == "tick") {
            changeFlashingAnimation_do();
            return;
        }
        else if(e._message == "systemRestart") {
            TrafficLightControllerCompartment compartment =  new TrafficLightControllerCompartment(TrafficLightControllerState.RED.getValue());
            
            this._transition_(compartment);
            return;
        }
        else if(e._message == "stop") {
            TrafficLightControllerCompartment compartment =  new TrafficLightControllerCompartment(TrafficLightControllerState.END.getValue());
            
            this._transition_(compartment);
            return;
        }
    }
    
    private void _sEnd_(FrameEvent e) {
        if(e._message == ">") {
            stopWorkingTimer_do();
            return;
        }
    }
    
    private void _sWorking_(FrameEvent e) {
        if(e._message == "stop") {
            TrafficLightControllerCompartment compartment =  new TrafficLightControllerCompartment(TrafficLightControllerState.END.getValue());
            
            this._transition_(compartment);
            return;
        }
        else if(e._message == "systemError") {
            TrafficLightControllerCompartment compartment =  new TrafficLightControllerCompartment(TrafficLightControllerState.FLASHINGRED.getValue());
            
            this._transition_(compartment);
            return;
        }
    }
    
    //===================== Actions Block ===================//
    
    // Unimplemented Actions
    
    protected void enterRed_do() { throw new UnsupportedOperationException(); }
    protected void enterGreen_do() { throw new UnsupportedOperationException(); }
    protected void enterYellow_do() { throw new UnsupportedOperationException(); }
    protected void enterFlashingRed_do() { throw new UnsupportedOperationException(); }
    protected void exitFlashingRed_do() { throw new UnsupportedOperationException(); }
    protected void startWorkingTimer_do() { throw new UnsupportedOperationException(); }
    protected void stopWorkingTimer_do() { throw new UnsupportedOperationException(); }
    protected void startFlashingTimer_do() { throw new UnsupportedOperationException(); }
    protected void stopFlashingTimer_do() { throw new UnsupportedOperationException(); }
    protected void changeColor_do(String color) { throw new UnsupportedOperationException(); }
    protected void startFlashing_do() { throw new UnsupportedOperationException(); }
    protected void stopFlashing_do() { throw new UnsupportedOperationException(); }
    protected void changeFlashingAnimation_do() { throw new UnsupportedOperationException(); }
    protected void log_do(String msg) { throw new UnsupportedOperationException(); }
    
    //===================== Domain Block ===================//
    
    private String flashColor  = "";
    
    //=============== Machinery and Mechanisms ==============//
    
    private int _state_;
    
    private void _transition_(TrafficLightControllerCompartment compartment) {
        this._nextCompartment_ = compartment;
    }
    
    private void _doTransition_(TrafficLightControllerCompartment nextCompartment) {
        this._mux_(new FrameEvent("<", this._compartment_.exitArgs));
        this._compartment_ = nextCompartment;
        this._mux_(new FrameEvent(">", this._compartment_.enterArgs));
    }
    
}

//=============== Compartment ==============//

class TrafficLightControllerCompartment {

    int state;
    
    TrafficLightControllerCompartment(int state) {
        this.state = state;
    }
    
    HashMap<String, Object> stateArgs = new HashMap<String, Object>();
    HashMap<String, Object> stateVars = new HashMap<String, Object>();
    HashMap<String, Object> enterArgs = new HashMap<String, Object>();
    HashMap<String, Object> exitArgs = new HashMap<String, Object>();
    FrameEvent _forwardEvent = new FrameEvent();
}


/********************

class TrafficLightControllerController extends TrafficLightController {
	TrafficLightControllerController() {
	  super();
	}

    protected void enterRed_do() {}

    protected void enterGreen_do() {}

    protected void enterYellow_do() {}

    protected void enterFlashingRed_do() {}

    protected void exitFlashingRed_do() {}

    protected void startWorkingTimer_do() {}

    protected void stopWorkingTimer_do() {}

    protected void startFlashingTimer_do() {}

    protected void stopFlashingTimer_do() {}

    protected void changeColor_do(String color) {}

    protected void startFlashing_do() {}

    protected void stopFlashing_do() {}

    protected void changeFlashingAnimation_do() {}

    protected void log_do(String msg) {}
}

********************/

