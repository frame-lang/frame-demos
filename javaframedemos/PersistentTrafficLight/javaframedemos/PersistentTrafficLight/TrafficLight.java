package javaframedemos.PersistentTrafficLight;

import java.util.*;

class TrafficLight{

    // creating private properties
    private int __state;
    private TrafficLightCompartment _compartment_;
    private TrafficLightCompartment _nextCompartment_;
    protected TrafficLightMom _manager;
    FrameEvent e = new FrameEvent();
    TrafficLight tLight;
    public TrafficLight(TrafficLightMom manager) {
        this.__state = this._sBegin_(e);
        this._manager = manager;
        this._compartment_ = new TrafficLightCompartment(this.__state);
        this._nextCompartment_ = null;
        
        // Send system start event
        e = new FrameEvent(">", null);
        try{
           this._mux_(e);
        }catch(Exception ex){
            //ex.printStackTrace();
        }   
    }

    // state enum
    private enum TrafficLightState {
        BEGIN(0),
        RED(1),
        GREEN(2),
        YELLOW(3),
        FLASHINGRED(4),
        END(5),
        WORKING(6);

        private final int value;

        private TrafficLightState(int value) {
            this.value = value;
        }

        public int getValue() {
            return value;
        }

    }

    // ===================Interface Block===================//

    public void stop()  throws Exception{
        e = new FrameEvent("stop", null);
        this._mux_(e);
    }

    public void tick()  throws Exception{
        e = new FrameEvent("tick", null);
        this._mux_(e);
    }

    public void systemError()  throws Exception{
        e = new FrameEvent("systemError", null);
        this._mux_(e);
    }

    public void systemRestart()  throws Exception{
        e = new FrameEvent("systemRestart", null);
        this._mux_(e);
    }

    public static TrafficLight loadTrafficLight(TrafficLightMomController manager, TrafficLightController data) {
            data._manager = manager;
        return data;
    }

    public TrafficLight marshal() {
        TrafficLight data = this;
        return data;
    }

    // ===================Multiplexer===================//
    public void _mux_(FrameEvent e)  throws Exception{
        if (this._compartment_.state == TrafficLightState.BEGIN.getValue()) {
            this._sBegin_(e);
        } else if (this._compartment_.state == TrafficLightState.RED.getValue()) {
            this._sRed_(e);
        } else if (this._compartment_.state == TrafficLightState.GREEN.getValue()) {
            this._sGreen_(e);
        } else if (this._compartment_.state == TrafficLightState.YELLOW.getValue()) {
            this._sYellow_(e);
        } else if (this._compartment_.state == TrafficLightState.FLASHINGRED.getValue()) {
            this._sFlashingRed_(e);
        } else if (this._compartment_.state == TrafficLightState.END.getValue()) {
            this._sEnd_(e);
        } else if (this._compartment_.state == TrafficLightState.WORKING.getValue()) {
            this._sWorking_(e);
        }

        if (this._nextCompartment_ != null) {
            TrafficLightCompartment nextCompartment = this._nextCompartment_;
            this._nextCompartment_ = null;
            if (nextCompartment._forwardEvent != null &&
                    nextCompartment._forwardEvent._message == ">") {
                this._mux_(new FrameEvent("<", this._compartment_.ExitArgs));
                this._compartment_ = nextCompartment;
                this._mux_(nextCompartment._forwardEvent);
            } else {
                this._doTransition_(nextCompartment);
                if (nextCompartment._forwardEvent != null) {
                    this._mux_(nextCompartment._forwardEvent);
                }
            }
            nextCompartment._forwardEvent = null;
        }
    }

    // =================== Machine Block ===================//

    private int _sBegin_(FrameEvent e) {
        if (e._message == ">") {
            this.startWorkingTimer_do();
            TrafficLightCompartment compartment = new TrafficLightCompartment(TrafficLightState.RED.getValue());
            return this._transition_(compartment);
        }
        return 0;
    }

    private void _sRed_(FrameEvent e)  throws Exception{
        if (e._message == ">") {
            this.enterRed_do();
            return;
        } else if (e._message == "tick") {
            TrafficLightCompartment compartment = new TrafficLightCompartment(TrafficLightState.GREEN.getValue());
            this._transition_(compartment);
            return;
        }
        this._sWorking_(e);

    }

    private void _sGreen_(FrameEvent e)  throws Exception{
        if (e._message == ">") {
            this.enterGreen_do();
            return;
        } else if (e._message == "tick") {
            TrafficLightCompartment compartment = new TrafficLightCompartment(TrafficLightState.YELLOW.getValue());
            this._transition_(compartment);
            return;
        }
        this._sWorking_(e);

    }

    private void _sYellow_(FrameEvent e)  throws Exception{
        if (e._message == ">") {
            this.enterYellow_do();
            return;
        } else if (e._message == "tick") {
            TrafficLightCompartment compartment = new TrafficLightCompartment(TrafficLightState.RED.getValue());
            this._transition_(compartment);
            return;
        }
        this._sWorking_(e);

    }

    private void _sFlashingRed_(FrameEvent e) {
        if (e._message == ">") {
            this.enterFlashingRed_do();
            this.stopWorkingTimer_do();
            this.startFlashingTimer_do();
            return;
        } else if (e._message == "<") {
            this.exitFlashingRed_do();
            this.stopFlashingTimer_do();
            this.startWorkingTimer_do();
            return;
        } else if (e._message == "tick") {
            this.changeFlashingAnimation_do();
            return;
        } else if (e._message == "systemRestart") {
            TrafficLightCompartment compartment = new TrafficLightCompartment(TrafficLightState.RED.getValue());
            this._transition_(compartment);
            return;
        } else if (e._message == "stop") {
            TrafficLightCompartment compartment = new TrafficLightCompartment(TrafficLightState.END.getValue());
            this._transition_(compartment);
            return;
        }

    }

    private void _sEnd_(FrameEvent e) {
        if (e._message == ">") {
            this.stopWorkingTimer_do();
            return;
        }
    }

    private void _sWorking_(FrameEvent e)  throws Exception{
        if (e._message == "stop") {
            TrafficLightCompartment compartment = new TrafficLightCompartment(TrafficLightState.END.getValue());
            this._doTransition_(compartment);
            return;
        } else if (e._message == "systemError") {
            TrafficLightCompartment compartment = new TrafficLightCompartment(TrafficLightState.FLASHINGRED.getValue());
            this._doTransition_(compartment);
            return;
        }
    }

    // =================== Action Block ==================//
    // Unimplemented Actions

    protected void enterRed_do() throws Exception { throw new UnsupportedOperationException(); }
    protected void enterGreen_do() throws Exception { throw new UnsupportedOperationException(); }
    protected void enterYellow_do() throws Exception { throw new UnsupportedOperationException(); }
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

    // =============== Machinery and Mechanisms ==============//

    private int _transition_(TrafficLightCompartment compartment) {
        this._nextCompartment_ = compartment;
        return this._nextCompartment_.state;
    }

    private void _doTransition_(TrafficLightCompartment nextCompartment)  throws Exception{
        this._mux_(new FrameEvent("<", this._compartment_.ExitArgs));
        this._compartment_ = nextCompartment;
        this._mux_(new FrameEvent(">", this._compartment_.EnterArgs));
    }

}

// =============== Compartment ==============//
class TrafficLightCompartment {
    int state;

    TrafficLightCompartment(int state) {
        this.state = state;
    }

    HashMap<String, Object> StateArgs = new HashMap<String, Object>();
    HashMap<String, Object> StateVars = new HashMap<String, Object>();
    HashMap<String, Object> EnterArgs = new HashMap<String, Object>();
    HashMap<String, Object> ExitArgs = new HashMap<String, Object>();
    FrameEvent _forwardEvent = new FrameEvent();
}

class TrafficLightController extends TrafficLight {
    TrafficLightController(TrafficLightMom manager) {
        super(manager);
    }
    @Override
    protected void enterRed_do() {
        //System.out.println("RED");
        //this._manager.enterRed_do();
        try {
            this._manager.enterRed();
        } catch (Exception e) {
           // e.printStackTrace();
        }
    }

    @Override
    protected void enterGreen_do() {
        //System.out.println("GREEN");
        try {
            this._manager.enterGreen();
        } catch (Exception e) {
           // e.printStackTrace();
        }
    }

    @Override
    protected void enterYellow_do() {
        //System.out.println("YELLOW");
        try {
            this._manager.enterYellow();
        } catch (Exception e) {
           // e.printStackTrace();
        }
    }

    protected void enterFlashingRed_do() {
    }

    protected void exitFlashingRed_do() {
    }

    protected void startWorkingTimer_do() {
    }

    protected void stopWorkingTimer_do() {
    }

    protected void startFlashingTimer_do() {
    }

    protected void stopFlashingTimer_do() {
    }

    protected void changeColor_do(String color) {
    }

    protected void startFlashing_do() {
    }

    protected void stopFlashing_do() {
    }

    protected void changeFlashingAnimation_do() {
    }

    protected void log_do(String msg) {
    }

}