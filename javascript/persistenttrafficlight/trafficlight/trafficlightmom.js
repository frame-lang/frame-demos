// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files

const FrameEvent = require("../framelang/framelang");
const TrafficLightController = require("./trafficlight");


class TrafficLightMom {
    
    // creating private properties
    
    #state
    #compartment
    #nextCompartment
    #trafficLight;
    #data;
    
    constructor () {

        // Create and intialize start state compartment.
        
        this.#state = this.#sNew_;
        this.#compartment = new TrafficLightMomCompartment(this.#state);
        this.#nextCompartment = null;
        
        // Initialize domain
        this.#trafficLight = null;
        this.#data = null;
        
        // Send system start event
        const frameEvent = FrameEvent(">", null);
        this.#mux(frameEvent);
    }
    
    //===================== Interface Block ===================//
    
    stop() {
        let e = FrameEvent("stop",null);
        this.#mux(e);
    }
    
    tick() {
        let e = FrameEvent("tick",null);
        this.#mux(e);
    }
    
    enterRed() {
        let e = FrameEvent("enterRed",null);
        this.#mux(e);
    }
    
    enterGreen() {
        let e = FrameEvent("enterGreen",null);
        this.#mux(e);
    }
    
    enterYellow() {
        let e = FrameEvent("enterYellow",null);
        this.#mux(e);
    }
    
    enterFlashingRed() {
        let e = FrameEvent("enterFlashingRed",null);
        this.#mux(e);
    }
    
    exitFlashingRed() {
        let e = FrameEvent("exitFlashingRed",null);
        this.#mux(e);
    }
    
    startWorkingTimer() {
        let e = FrameEvent("startWorkingTimer",null);
        this.#mux(e);
    }
    
    stopWorkingTimer() {
        let e = FrameEvent("stopWorkingTimer",null);
        this.#mux(e);
    }
    
    startFlashingTimer() {
        let e = FrameEvent("startFlashingTimer",null);
        this.#mux(e);
    }
    
    stopFlashingTimer() {
        let e = FrameEvent("stopFlashingTimer",null);
        this.#mux(e);
    }
    
    changeColor(color) {
        let e = FrameEvent("changeColor",{"color":color});
        this.#mux(e);
    }
    
    startFlashing() {
        let e = FrameEvent("startFlashing",null);
        this.#mux(e);
    }
    
    stopFlashing() {
        let e = FrameEvent("stopFlashing",null);
        this.#mux(e);
    }
    
    changeFlashingAnimation() {
        let e = FrameEvent("changeFlashingAnimation",null);
        this.#mux(e);
    }
    
    systemError() {
        let e = FrameEvent("systemError",null);
        this.#mux(e);
    }
    
    systemRestart() {
        let e = FrameEvent("systemRestart",null);
        this.#mux(e);
    }
    
    log(msg) {
        let e = FrameEvent("log",{"msg":msg});
        this.#mux(e);
    }
    
    //====================== Multiplexer ====================//
    
    #mux(e) {
        switch (this.#compartment.state) {
            case this.#sNew_:
                this.#sNew_(e);
                break;
            case this.#sSaving_:
                this.#sSaving_(e);
                break;
            case this.#sPersisted_:
                this.#sPersisted_(e);
                break;
            case this.#sWorking_:
                this.#sWorking_(e);
                break;
            case this.#sTrafficLightApi_:
                this.#sTrafficLightApi_(e);
                break;
            case this.#sEnd_:
                this.#sEnd_(e);
                break;
        }
        
        if( this.#nextCompartment != null) {
            let nextCompartment = this.#nextCompartment
            this.#nextCompartment = null
            if (nextCompartment._forwardEvent != null && 
               nextCompartment._forwardEvent._message == ">") {
                this.#mux(FrameEvent( "<", this.#compartment.ExitArgs))
                this.#compartment = nextCompartment
                this.#mux(nextCompartment._forwardEvent)
            } else {
                this.#doTransition(nextCompartment)
                if (nextCompartment._forwardEvent != null) {
                    this.#mux(nextCompartment._forwardEvent)
                }
            }
            nextCompartment._forwardEvent = null
        }
    }
    
    //===================== Machine Block ===================//
    
    #sNew_(e) {
        switch (e._message) {
            case ">":
                {
                this.#trafficLight = new TrafficLightController(this);
                // Traffic Light\nStarted
                let compartment =  new TrafficLightMomCompartment(this.#sSaving_);
                
                
                this.#transition(compartment);
                return;
                }
                
        }
        this.#sTrafficLightApi_(e);
        
    }
    
    #sSaving_(e) {
        switch (e._message) {
            case ">":
                {
                this.#data = this.#trafficLight.marshal();
                this.#trafficLight = null;
                // Saved
                let compartment =  new TrafficLightMomCompartment(this.#sPersisted_);
                
                
                this.#transition(compartment);
                return;
                }
                
        }
    }
    
    #sPersisted_(e) {
        switch (e._message) {
            case "tick":
                {
                // Tick
                let compartment =  new TrafficLightMomCompartment(this.#sWorking_);
                
                compartment._forwardEvent = e;
                
                this.#transition(compartment);
                return;
                }
                
            case "systemError":
                {
                // System Error
                let compartment =  new TrafficLightMomCompartment(this.#sWorking_);
                
                compartment._forwardEvent = e;
                
                this.#transition(compartment);
                return;
                }
                
            case "stop":
                {
                // Stop
                let compartment =  new TrafficLightMomCompartment(this.#sEnd_);
                
                
                this.#transition(compartment);
                return;
                }
                
        }
    }
    
    #sWorking_(e) {
        switch (e._message) {
            case ">":
                {
                this.#trafficLight = TrafficLightController.loadTrafficLight(this,this.#data);
                return;
                }
                
            case "tick":
                {
                this.#trafficLight.tick();
                // Done
                let compartment =  new TrafficLightMomCompartment(this.#sSaving_);
                
                
                this.#transition(compartment);
                return;
                }
                
            case "systemError":
                {
                this.#trafficLight.systemError_do();
                // Done
                let compartment =  new TrafficLightMomCompartment(this.#sSaving_);
                
                
                this.#transition(compartment);
                return;
                }
                
        }
        this.#sTrafficLightApi_(e);
        
    }
    
    #sTrafficLightApi_(e) {
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
    
    #sEnd_(e) {
        switch (e._message) {
            case ">":
                {
                this.#trafficLight = TrafficLightController.loadTrafficLight(this,this.#data);
                this.#trafficLight.stop();
                this.#trafficLight = null;
                return;
                }
                
        }
        this.#sTrafficLightApi_(e);
        
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
    
    #transition(compartment) {
        this.#nextCompartment = compartment;
    }
    
    #doTransition(nextCompartment) {
        this.#mux(FrameEvent("<", this.#compartment.ExitArgs));
        this.#compartment = nextCompartment;
        this.#mux(FrameEvent(">", this.#compartment.EnterArgs));
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

module.exports = TrafficLightMomController

