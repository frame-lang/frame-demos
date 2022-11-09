// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files
package java.PersistentTrafficLight;
import java.util.*;

import javaframedemos.PersistentTrafficLight.FrameEvent;


class TrafficLightMomController {
    
    private TrafficLightMomControllerCompartment _compartment_;
    private TrafficLightMomControllerCompartment _nextCompartment_;
    
    TrafficLightMomController() {
        
        // Create and intialize start state compartment.
        this._compartment_ = new TrafficLightMomControllerCompartment(this._state_);
        this._nextCompartment_ = null;
        
        
        // Send system start event
        FrameEvent frame_event = new FrameEvent(">", null);
        this._mux_(frame_event);
    
    }
    
    // states enum
    private enum TrafficLightMomControllerState {
        
        NEW(0),
        SAVING(1),
        PERSISTED(2),
        WORKING(3),
        TRAFFICLIGHTAPI(4),
        END(5);
        
        public final int value;
        
        private TrafficLightMomControllerState(int value) {
            this.value=value;
        }
        
        public int getValue() {
            return value;
        }
    }
    
    //====================== Multiplexer ====================//
    
    private void _mux_(FrameEvent e) {
        if(this._compartment_.state == TrafficLightMomControllerState.NEW.getValue()) {
            this._sNew_(e);
        }else if(this._compartment_.state == TrafficLightMomControllerState.SAVING.getValue()) {
            this._sSaving_(e);
        }else if(this._compartment_.state == TrafficLightMomControllerState.PERSISTED.getValue()) {
            this._sPersisted_(e);
        }else if(this._compartment_.state == TrafficLightMomControllerState.WORKING.getValue()) {
            this._sWorking_(e);
        }else if(this._compartment_.state == TrafficLightMomControllerState.TRAFFICLIGHTAPI.getValue()) {
            this._sTrafficLightApi_(e);
        }else if(this._compartment_.state == TrafficLightMomControllerState.END.getValue()) {
            this._sEnd_(e);
        }
        
        
        if(this._nextCompartment_ != null) {
            TrafficLightMomControllerCompartment nextCompartment = this._nextCompartment_;
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
    
    public void enterRed() {
        FrameEvent e = new FrameEvent("enterRed", null);
        this._mux_(e);
    }
    
    public void enterGreen() {
        FrameEvent e = new FrameEvent("enterGreen", null);
        this._mux_(e);
    }
    
    public void enterYellow() {
        FrameEvent e = new FrameEvent("enterYellow", null);
        this._mux_(e);
    }
    
    public void enterFlashingRed() {
        FrameEvent e = new FrameEvent("enterFlashingRed", null);
        this._mux_(e);
    }
    
    public void exitFlashingRed() {
        FrameEvent e = new FrameEvent("exitFlashingRed", null);
        this._mux_(e);
    }
    
    public void startWorkingTimer() {
        FrameEvent e = new FrameEvent("startWorkingTimer", null);
        this._mux_(e);
    }
    
    public void stopWorkingTimer() {
        FrameEvent e = new FrameEvent("stopWorkingTimer", null);
        this._mux_(e);
    }
    
    public void startFlashingTimer() {
        FrameEvent e = new FrameEvent("startFlashingTimer", null);
        this._mux_(e);
    }
    
    public void stopFlashingTimer() {
        FrameEvent e = new FrameEvent("stopFlashingTimer", null);
        this._mux_(e);
    }
    
    public void changeColor(String color) {
        HashMap<String,Object> parameters = new HashMap<String,Object>();
        parameters.put("color", color);

        FrameEvent e = new FrameEvent("changeColor", parameters);
        this._mux_(e);
    }
    
    public void startFlashing() {
        FrameEvent e = new FrameEvent("startFlashing", null);
        this._mux_(e);
    }
    
    public void stopFlashing() {
        FrameEvent e = new FrameEvent("stopFlashing", null);
        this._mux_(e);
    }
    
    public void changeFlashingAnimation() {
        FrameEvent e = new FrameEvent("changeFlashingAnimation", null);
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
    
    public void log(String msg) {
        HashMap<String,Object> parameters = new HashMap<String,Object>();
        parameters.put("msg", msg);

        FrameEvent e = new FrameEvent("log", parameters);
        this._mux_(e);
    }
    
    
    //===================== Machine Block ===================//
    
    private void _sNew_(FrameEvent e) {
        if(e._message == ">") {
            this.trafficLight = TrafficLightController(this);
            // Traffic Light\nStarted
            TrafficLightMomControllerCompartment compartment =  new TrafficLightMomControllerCompartment(TrafficLightMomControllerState.SAVING.getValue());
            
            this._transition_(compartment);
            return;
        }
        _sTrafficLightApi_(e);
        
    }
    
    private void _sSaving_(FrameEvent e) {
        if(e._message == ">") {
            this.data = this.trafficLight.marshal();
            this.trafficLight = null;
            // Saved
            TrafficLightMomControllerCompartment compartment =  new TrafficLightMomControllerCompartment(TrafficLightMomControllerState.PERSISTED.getValue());
            
            this._transition_(compartment);
            return;
        }
    }
    

    private void _sPersisted_(FrameEvent e) {
        if(e._message == "tick") {
            // Tick
            TrafficLightMomControllerCompartment compartment =  new TrafficLightMomControllerCompartment(TrafficLightMomControllerState.WORKING.getValue());
            
            compartment._forwardEvent = e;
            
            this._transition_(compartment);
            return;
        }
        else if(e._message == "systemError") {
            // System Error
            TrafficLightMomControllerCompartment compartment =  new TrafficLightMomControllerCompartment(TrafficLightMomControllerState.WORKING.getValue());
            
            compartment._forwardEvent = e;
            
            this._transition_(compartment);
            return;
        }
        else if(e._message == "stop") {
            // Stop
            TrafficLightMomControllerCompartment compartment =  new TrafficLightMomControllerCompartment(TrafficLightMomControllerState.END.getValue());
            
            this._transition_(compartment);
            return;
        }
    }
    
    private void _sWorking_(FrameEvent e) {
        if(e._message == ">") {
            this.trafficLight = TrafficLightController.loadTrafficLight(this,this.data);
            return;
        }
        else if(e._message == "tick") {
            this.trafficLight.tick();
            // Done
            TrafficLightMomControllerCompartment compartment =  new TrafficLightMomControllerCompartment(TrafficLightMomControllerState.SAVING.getValue());
            
            this._transition_(compartment);
            return;
        }
        else if(e._message == "systemError") {
            this.trafficLight.systemError_do();
            // Done
            TrafficLightMomControllerCompartment compartment =  new TrafficLightMomControllerCompartment(TrafficLightMomControllerState.SAVING.getValue());
            
            this._transition_(compartment);
            return;
        }
        _sTrafficLightApi_(e);
        
    }
    
    private void _sTrafficLightApi_(FrameEvent e) {
        if(e._message == "enterRed") {
            enterRed_do();
            return;
        }
        else if(e._message == "enterGreen") {
            enterGreen_do();
            return;
        }
        else if(e._message == "enterYellow") {
            enterYellow_do();
            return;
        }
        else if(e._message == "enterFlashingRed") {
            enterFlashingRed_do();
            return;
        }
        else if(e._message == "exitFlashingRed") {
            exitFlashingRed_do();
            return;
        }
        else if(e._message == "startWorkingTimer") {
            startWorkingTimer_do();
            return;
        }
        else if(e._message == "stopWorkingTimer") {
            stopWorkingTimer_do();
            return;
        }
        else if(e._message == "startFlashingTimer") {
            startFlashingTimer_do();
            return;
        }
        else if(e._message == "stopFlashingTimer") {
            stopFlashingTimer_do();
            return;
        }
        else if(e._message == "changeColor") {
            changeColor_do(String.valueOf(((String) e._parameters.get("color"))));
            return;
        }
        else if(e._message == "startFlashing") {
            startFlashing_do();
            return;
        }
        else if(e._message == "stopFlashing") {
            stopFlashing_do();
            return;
        }
        else if(e._message == "changeFlashingAnimation") {
            changeFlashingAnimation_do();
            return;
        }
        else if(e._message == "systemError") {
            systemError_do();
            return;
        }
        else if(e._message == "systemRestart") {
            systemRestart_do();
            return;
        }
        else if(e._message == "log") {
            log_do(String.valueOf(((String) e._parameters.get("msg"))));
            return;
        }
    }
    
    private void _sEnd_(FrameEvent e) {
        if(e._message == ">") {
            this.trafficLight = TrafficLightController.loadTrafficLight(this,this.data);
            this.trafficLight.stop();
            this.trafficLight = null;
            return;
        }
        _sTrafficLightApi_(e);
        
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
    protected void systemError_do() { throw new UnsupportedOperationException(); }
    protected void systemRestart_do() { throw new UnsupportedOperationException(); }
    protected void log_do(String msg) { throw new UnsupportedOperationException(); }
    
    //===================== Domain Block ===================//
    
    private TrafficLightMomController trafficLight  = null;
    private String data  = null;
    
    //=============== Machinery and Mechanisms ==============//
    
    private int _state_;
    
    private void _transition_(TrafficLightMomControllerCompartment compartment) {
        this._nextCompartment_ = compartment;
    }
    
    private void _doTransition_(TrafficLightMomControllerCompartment nextCompartment) {
        this._mux_(new FrameEvent("<", this._compartment_.exitArgs));
        this._compartment_ = nextCompartment;
        this._mux_(new FrameEvent(">", this._compartment_.enterArgs));
    }
    
}

//=============== Compartment ==============//

class TrafficLightMomControllerCompartment {

    int state;
    
    TrafficLightMomControllerCompartment(int state) {
        this.state = state;
    }
    
    HashMap<String, Object> stateArgs = new HashMap<String, Object>();
    HashMap<String, Object> stateVars = new HashMap<String, Object>();
    HashMap<String, Object> enterArgs = new HashMap<String, Object>();
    HashMap<String, Object> exitArgs = new HashMap<String, Object>();
    FrameEvent _forwardEvent = new FrameEvent();
}


/********************

class TrafficLightMomControllerController extends TrafficLightMomController {
	TrafficLightMomControllerController() {
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

    protected void systemError_do() {}

    protected void systemRestart_do() {}

    protected void log_do(String msg) {}
}

********************/

