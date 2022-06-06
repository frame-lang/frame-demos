// emitted from framec_v0.10.0

const TrafficLightController = require("./machine");
const FrameEvent = require("../framelang/framelang");
// get include files at https://github.com/frame-lang/frame-ancillary-files

class TrafficLightMom {
    
    constructor () {
        
        // Create and intialize start state compartment.
        
        this._state = this._sNew_;
        this._compartment = new TrafficLightMomCompartment(this._state);
        this._nextCompartment = null;
        
        // Initialize domain
        this.trafficLight = null;
        this.data = null;
        
        // Send system start event
        this._frameEvent = FrameEvent(">", null);
        this._mux_(this._frameEvent);
    }
    
    //===================== Interface Block ===================//
    
    stop() {
        let e = FrameEvent("stop",null);
        this._mux_(e);
    }
    
    tick() {
        let e = FrameEvent("tick",null);
        this._mux_(e);
    }
    
    enterRed() {
        let e = FrameEvent("enterRed",null);
        this._mux_(e);
    }
    
    enterGreen() {
        let e = FrameEvent("enterGreen",null);
        this._mux_(e);
    }
    
    enterYellow() {
        let e = FrameEvent("enterYellow",null);
        this._mux_(e);
    }
    
    enterFlashingRed() {
        let e = FrameEvent("enterFlashingRed",null);
        this._mux_(e);
    }
    
    exitFlashingRed() {
        let e = FrameEvent("exitFlashingRed",null);
        this._mux_(e);
    }
    
    startWorkingTimer() {
        let e = FrameEvent("startWorkingTimer",null);
        this._mux_(e);
    }
    
    stopWorkingTimer() {
        let e = FrameEvent("stopWorkingTimer",null);
        this._mux_(e);
    }
    
    startFlashingTimer() {
        let e = FrameEvent("startFlashingTimer",null);
        this._mux_(e);
    }
    
    stopFlashingTimer() {
        let e = FrameEvent("stopFlashingTimer",null);
        this._mux_(e);
    }
    
    changeColor(color) {
        let e = FrameEvent("changeColor",{"color":color});
        this._mux_(e);
    }
    
    startFlashing() {
        let e = FrameEvent("startFlashing",null);
        this._mux_(e);
    }
    
    stopFlashing() {
        let e = FrameEvent("stopFlashing",null);
        this._mux_(e);
    }
    
    changeFlashingAnimation() {
        let e = FrameEvent("changeFlashingAnimation",null);
        this._mux_(e);
    }
    
    systemError() {
        let e = FrameEvent("systemError",null);
        this._mux_(e);
    }
    
    systemRestart() {
        let e = FrameEvent("systemRestart",null);
        this._mux_(e);
    }
    
    log(msg) {
        let e = FrameEvent("log",{"msg":msg});
        this._mux_(e);
    }
    
    //====================== Multiplexer ====================//
    
    _mux_(e) {
        switch (this._compartment.state) {
            case this._sNew_:
                this._sNew_(e);
                break;
            case this._sSaving_:
                this._sSaving_(e);
                break;
            case this._sPersisted_:
                this._sPersisted_(e);
                break;
            case this._sWorking_:
                this._sWorking_(e);
                break;
            case this._sTrafficLightApi_:
                this._sTrafficLightApi_(e);
                break;
            case this._sEnd_:
                this._sEnd_(e);
                break;
        }
        
        if( this._nextCompartment != null) {
            let nextCompartment = this._nextCompartment
            this._nextCompartment = null
            if (nextCompartment._forwardEvent != null && 
               nextCompartment._forwardEvent._message == ">") {
                this._mux_(FrameEvent( "<", this._compartment.ExitArgs))
                this._compartment = nextCompartment
                this._mux_(nextCompartment._forwardEvent)
            } else {
                this._do_transition_(nextCompartment)
                if (nextCompartment._forwardEvent != null) {
                    this._mux_(nextCompartment._forwardEvent)
                }
            }
            nextCompartment._forwardEvent = null
        }
    }
    
    //===================== Machine Block ===================//
    
    _sNew_(e) {
        switch (e._message) {
            case ">":
                {
                this.trafficLight = new TrafficLightController(this);
                // Traffic Light\nStarted
                let compartment =  new TrafficLightMomCompartment(this._sSaving_);
                
                
                this._transition_(compartment);
                return;
                }
                
        }
        this._sTrafficLightApi_(e);
        
    }
    
    _sSaving_(e) {
        switch (e._message) {
            case ">":
                {
                this.data = this.trafficLight.marshal();
                this.trafficLight = null;
                // Saved
                let compartment =  new TrafficLightMomCompartment(this._sPersisted_);
                
                
                this._transition_(compartment);
                return;
                }
                
        }
    }
    
    _sPersisted_(e) {
        switch (e._message) {
            case "tick":
                {
                // Tick
                let compartment =  new TrafficLightMomCompartment(this._sWorking_);
                
                compartment._forwardEvent = e;
                
                this._transition_(compartment);
                return;
                }
                
            case "systemError":
                {
                // System Error
                let compartment =  new TrafficLightMomCompartment(this._sWorking_);
                
                compartment._forwardEvent = e;
                
                this._transition_(compartment);
                return;
                }
                
            case "stop":
                {
                // Stop
                let compartment =  new TrafficLightMomCompartment(this._sEnd_);
                
                
                this._transition_(compartment);
                return;
                }
                
        }
    }
    
    _sWorking_(e) {
        switch (e._message) {
            case ">":
                {
                this.trafficLight = TrafficLightController.loadTrafficLight(this,this.data);
                return;
                }
                
            case "tick":
                {
                this.trafficLight.tick();
                // Done
                let compartment =  new TrafficLightMomCompartment(this._sSaving_);
                
                
                this._transition_(compartment);
                return;
                }
                
            case "systemError":
                {
                this.trafficLight.systemError_do();
                // Done
                let compartment =  new TrafficLightMomCompartment(this._sSaving_);
                
                
                this._transition_(compartment);
                return;
                }
                
        }
        this._sTrafficLightApi_(e);
        
    }
    
    _sTrafficLightApi_(e) {
        switch (e._message) {
            case "enterRed":
                {
                this.enterRed_do();
                return;
                }
                
            case "enterGreen":
                {
                this.enterGreen_do();
                return;
                }
                
            case "enterYellow":
                {
                this.enterYellow_do();
                return;
                }
                
            case "enterFlashingRed":
                {
                this.enterFlashingRed_do();
                return;
                }
                
            case "exitFlashingRed":
                {
                this.exitFlashingRed_do();
                return;
                }
                
            case "startWorkingTimer":
                {
                this.startWorkingTimer_do();
                return;
                }
                
            case "stopWorkingTimer":
                {
                this.stopWorkingTimer_do();
                return;
                }
                
            case "startFlashingTimer":
                {
                this.startFlashingTimer_do();
                return;
                }
                
            case "stopFlashingTimer":
                {
                this.stopFlashingTimer_do();
                return;
                }
                
            case "changeColor":
                {
                this.changeColor_do((e._parameters["color"]));
                return;
                }
                
            case "startFlashing":
                {
                this.startFlashing_do();
                return;
                }
                
            case "stopFlashing":
                {
                this.stopFlashing_do();
                return;
                }
                
            case "changeFlashingAnimation":
                {
                this.changeFlashingAnimation_do();
                return;
                }
                
            case "systemError":
                {
                this.systemError_do();
                return;
                }
                
            case "systemRestart":
                {
                this.systemRestart_do();
                return;
                }
                
            case "log":
                {
                this.log_do((e._parameters["msg"]));
                return;
                }
                
        }
    }
    
    _sEnd_(e) {
        switch (e._message) {
            case ">":
                {
                this.trafficLight = TrafficLightController.loadTrafficLight(this,this.data);
                this.trafficLight.stop();
                this.trafficLight = null;
                return;
                }
                
        }
        this._sTrafficLightApi_(e);
        
    }
    
    //===================== Actions Block ===================//
    
    // Unimplemented Actions
    
    enterRed_do() { throw new Error('Action not implemented.'); }
    enterGreen_do() { throw new Error('Action not implemented.'); }
    enterYellow_do() { throw new Error('Action not implemented.'); }
    enterFlashingRed_do() { throw new Error('Action not implemented.'); }
    exitFlashingRed_do() { throw new Error('Action not implemented.'); }
    startWorkingTimer_do() { throw new Error('Action not implemented.'); }
    stopWorkingTimer_do() { throw new Error('Action not implemented.'); }
    startFlashingTimer_do() { throw new Error('Action not implemented.'); }
    stopFlashingTimer_do() { throw new Error('Action not implemented.'); }
    changeColor_do(color) { throw new Error('Action not implemented.'); }
    startFlashing_do() { throw new Error('Action not implemented.'); }
    stopFlashing_do() { throw new Error('Action not implemented.'); }
    changeFlashingAnimation_do() { throw new Error('Action not implemented.'); }
    systemError_do() { throw new Error('Action not implemented.'); }
    systemRestart_do() { throw new Error('Action not implemented.'); }
    log_do(msg) { throw new Error('Action not implemented.'); }
    
    //=============== Machinery and Mechanisms ==============//
    
    _transition_(compartment) {
        this._nextCompartment = compartment;
    }
    
    _do_transition_(nextCompartment) {
        this._mux_(FrameEvent("<", this._compartment.ExitArgs));
        this._compartment = nextCompartment;
        this._mux_(FrameEvent(">", this._compartment.EnterArgs));
    }
    
    
    
};


//=============== Compartment ==============//

class TrafficLightMomCompartment {

    constructor(state) {
        this.state = state
    }
    
    StateArgs = {};
    StateVars = {};
    EnterArgs = {};
    ExitArgs = {};
    _forwardEvent = FrameEvent.call(this)
}


// /********************

class TrafficLightMomController extends TrafficLightMom {

	constructor() {
	  super()
	}

	enterRed_do() {
        console.log("Red")
    }
	enterGreen_do() {
        console.log("Green")
    }
	enterYellow_do() {
        console.log("Yellow")
    }
	enterFlashingRed_do() {}
	exitFlashingRed_do() {}
	startWorkingTimer_do() {}
	stopWorkingTimer_do() {}
	startFlashingTimer_do() {}
	stopFlashingTimer_do() {}
	changeColor_do(color) {}
	startFlashing_do() {}
	stopFlashing_do() {}
	changeFlashingAnimation_do() {}
	systemError_do() {}
	systemRestart_do() {}
	log_do(msg) {}

};

// ********************/


module.exports = TrafficLightMomController