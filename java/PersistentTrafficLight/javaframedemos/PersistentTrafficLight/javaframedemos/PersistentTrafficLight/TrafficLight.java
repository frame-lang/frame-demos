package javaframedemos.PersistentTrafficLight;

import java.util.*;


class TrafficLight {
    
    private TrafficLightCompartment _compartment_;
    private TrafficLightCompartment _nextCompartment_;
    
    protected TrafficLightMom _manager;
    
    TrafficLight(TrafficLightMom manager) {
        
        // Create and intialize start state compartment.
        this._manager = manager;
        this._compartment_ = new TrafficLightCompartment(this._state_);
        this._nextCompartment_ = null;
        
        
        // Send system start event
        FrameEvent frame_event = new FrameEvent(">", null);
        this._mux_(frame_event);
    
    }
    
    // states enum
    private enum TrafficLightState {
        
        BEGIN(0),
        RED(1),
        GREEN(2),
        YELLOW(3),
        FLASHINGRED(4),
        END(5),
        WORKING(6);
        
        public final int value;
        
        private TrafficLightState(int value) {
            this.value=value;
        }
        
        public int getValue() {
            return value;
        }
    }
    
    
    public static TrafficLightController loadTrafficLight(TrafficLightMomController manager, TrafficLightController data){
        
        data._manager = manager;
        return data;
        
    }
    
    public TrafficLight marshal(){
        
        TrafficLight data = this;
        return data;
        
    }
    
    //====================== Multiplexer ====================//
    
    private void _mux_(FrameEvent e) {
        if(this._compartment_.state == TrafficLightState.BEGIN.getValue()) {
            this._sBegin_(e);
        }else if(this._compartment_.state == TrafficLightState.RED.getValue()) {
            this._sRed_(e);
        }else if(this._compartment_.state == TrafficLightState.GREEN.getValue()) {
            this._sGreen_(e);
        }else if(this._compartment_.state == TrafficLightState.YELLOW.getValue()) {
            this._sYellow_(e);
        }else if(this._compartment_.state == TrafficLightState.FLASHINGRED.getValue()) {
            this._sFlashingRed_(e);
        }else if(this._compartment_.state == TrafficLightState.END.getValue()) {
            this._sEnd_(e);
        }else if(this._compartment_.state == TrafficLightState.WORKING.getValue()) {
            this._sWorking_(e);
        }
        
        
        if(this._nextCompartment_ != null) {
            TrafficLightCompartment nextCompartment = this._nextCompartment_;
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
    
    
    //===================== Machine Block ===================//
    
    private void _sBegin_(FrameEvent e) {
        if(e._message == ">") {
            startWorkingTimer_do();
            TrafficLightCompartment compartment =  new TrafficLightCompartment(TrafficLightState.RED.getValue());
            
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
            TrafficLightCompartment compartment =  new TrafficLightCompartment(TrafficLightState.GREEN.getValue());
            
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
            TrafficLightCompartment compartment =  new TrafficLightCompartment(TrafficLightState.YELLOW.getValue());
            
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
            TrafficLightCompartment compartment =  new TrafficLightCompartment(TrafficLightState.RED.getValue());
            
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
            TrafficLightCompartment compartment =  new TrafficLightCompartment(TrafficLightState.RED.getValue());
            
            this._transition_(compartment);
            return;
        }
        else if(e._message == "stop") {
            TrafficLightCompartment compartment =  new TrafficLightCompartment(TrafficLightState.END.getValue());
            
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
            TrafficLightCompartment compartment =  new TrafficLightCompartment(TrafficLightState.END.getValue());
            
            this._transition_(compartment);
            return;
        }
        else if(e._message == "systemError") {
            TrafficLightCompartment compartment =  new TrafficLightCompartment(TrafficLightState.FLASHINGRED.getValue());
            
            this._transition_(compartment);
            return;
        }
    }
    
    //===================== Actions Block ===================//
    protected void enterRed_do() {
        try {this._manager.enterRed_do();}catch(Exception e){}
    }
    
    protected void enterGreen_do() {
        try {this._manager.enterGreen_do();}catch(Exception e){}
    }
    
    protected void enterYellow_do() {
        try {this._manager.enterYellow_do();}catch(Exception e){}
    }
    
    
    // Unimplemented Actions
    
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
    
    private int _transition_(TrafficLightCompartment compartment) {
        this._nextCompartment_ = compartment;
        return this._nextCompartment_.state;
    }
    
    private void _doTransition_(TrafficLightCompartment nextCompartment) {
        this._mux_(new FrameEvent("<", this._compartment_.exitArgs));
        this._compartment_ = nextCompartment;
        this._mux_(new FrameEvent(">", this._compartment_.enterArgs));
    }
    
}

//=============== Compartment ==============//

class TrafficLightCompartment {

    int state;
    
    TrafficLightCompartment(int state) {
        this.state = state;
    }
    
    HashMap<String, Object> stateArgs = new HashMap<String, Object>();
    HashMap<String, Object> stateVars = new HashMap<String, Object>();
    HashMap<String, Object> enterArgs = new HashMap<String, Object>();
    HashMap<String, Object> exitArgs = new HashMap<String, Object>();
    FrameEvent _forwardEvent = new FrameEvent();
}


class TrafficLightController extends TrafficLight {

	TrafficLightController(TrafficLightMom manager) {
	  super(manager);
	}
    
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