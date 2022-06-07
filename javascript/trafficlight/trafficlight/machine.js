// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files

const FrameEvent = require("../framelang/framelang");

class TrafficLight {
    
    // creating private properties
    
    #state
    #compartment
    #nextCompartment
    #flashColor;
    #mom;
    #ticker;
    
    constructor () {
        // Create and intialize start state compartment.
        
        this.#state = this.#sBegin_;
        this.#compartment = new TrafficLightCompartment(this.#state);
        this.#nextCompartment = null;
        
        // Initialize domain
        this.#flashColor = "";
        this.#mom = null;
        this.#ticker = null;
        
        // Send system start event
        const frameEvent = FrameEvent(">", null);
        this.#mux(frameEvent);
    }
    
    //===================== Interface Block ===================//
    
    start(mom) {
        let e = FrameEvent(">>",{"mom":mom});
        this.#mux(e);
    }
    
    stop() {
        let e = FrameEvent("stop",null);
        this.#mux(e);
    }
    
    tick() {
        let e = FrameEvent("tick",null);
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
    
    //====================== Multiplexer ====================//
    
    #mux(e) {
        switch (this.#compartment.state) {
            case this.#sBegin_:
                this.#sBegin_(e);
                break;
            case this.#sRed_:
                this.#sRed_(e);
                break;
            case this.#sGreen_:
                this.#sGreen_(e);
                break;
            case this.#sYellow_:
                this.#sYellow_(e);
                break;
            case this.#sFlashingRed_:
                this.#sFlashingRed_(e);
                break;
            case this.#sEnd_:
                this.#sEnd_(e);
                break;
            case this.#sWorking_:
                this.#sWorking_(e);
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
    
    #sBegin_(e) {
        switch (e._message) {
            case ">>":
                {
                this.startWorkingTimer_do();
                let compartment =  new TrafficLightCompartment(this.#sRed_);
                
                
                this.#transition(compartment);
                return;
                }
                
        }
    }
    
    #sRed_(e) {
        switch (e._message) {
            case ">":
                {
                this.enterRed_do();
                return;
                }
                
            case "tick":
                {
                let compartment =  new TrafficLightCompartment(this.#sGreen_);
                
                
                this.#transition(compartment);
                return;
                }
                
        }
        this.#sWorking_(e);
        
    }
    
    #sGreen_(e) {
        switch (e._message) {
            case ">":
                {
                this.enterGreen_do();
                return;
                }
                
            case "tick":
                {
                let compartment =  new TrafficLightCompartment(this.#sYellow_);
                
                
                this.#transition(compartment);
                return;
                }
                
        }
        this.#sWorking_(e);
        
    }
    
    #sYellow_(e) {
        switch (e._message) {
            case ">":
                {
                this.enterYellow_do();
                return;
                }
                
            case "tick":
                {
                let compartment =  new TrafficLightCompartment(this.#sRed_);
                
                
                this.#transition(compartment);
                return;
                }
                
        }
        this.#sWorking_(e);
        
    }
    
    #sFlashingRed_(e) {
        switch (e._message) {
            case ">":
                {
                this.enterFlashingRed_do();
                this.stopWorkingTimer_do();
                this.startFlashingTimer_do();
                return;
                }
                
            case "<":
                {
                this.exitFlashingRed_do();
                this.stopFlashingTimer_do();
                this.startWorkingTimer_do();
                return;
                }
                
            case "tick":
                {
                this.changeFlashingAnimation_do();
                return;
                }
                
            case "systemRestart":
                {
                let compartment =  new TrafficLightCompartment(this.#sRed_);
                
                
                this.#transition(compartment);
                return;
                }
                
            case "stop":
                {
                let compartment =  new TrafficLightCompartment(this.#sEnd_);
                
                
                this.#transition(compartment);
                return;
                }
                
        }
    }
    
    #sEnd_(e) {
        switch (e._message) {
            case ">":
                {
                this.stopWorkingTimer_do();
                return;
                }
                
        }
    }
    
    #sWorking_(e) {
        switch (e._message) {
            case "stop":
                {
                let compartment =  new TrafficLightCompartment(this.#sEnd_);
                
                
                this.#transition(compartment);
                return;
                }
                
            case "systemError":
                {
                let compartment =  new TrafficLightCompartment(this.#sFlashingRed_);
                
                
                this.#transition(compartment);
                return;
                }
                
        }
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

class TrafficLightCompartment {

    constructor(state) {
        this.state = state
    }
    
    StateArgs = {};
    StateVars = {};
    EnterArgs = {};
    ExitArgs = {};
    _forwardEvent = FrameEvent.call(this)
}




class TrafficLightController extends TrafficLight {

	constructor() {
	  super()
	}
	enterRed_do() {
        console.log("enterRed_do()")
    }
	enterGreen_do() {
        console.log("enterGreen_do()")
    }
	enterYellow_do() {
        console.log("enterYellow_do()")
    }
	enterFlashingRed_do() {}
	exitFlashingRed_do() {}
	startWorkingTimer_do() {
        this.ticker = setInterval(() => {
            this.tick()
        }, 1000)
    }
	stopWorkingTimer_do() {
        clearInterval(this.ticker)
    }
	startFlashingTimer_do() {}
	stopFlashingTimer_do() {}
	changeColor_do(color) {}
	startFlashing_do() {}
	stopFlashing_do() {}
	changeFlashingAnimation_do() {}
	log_do(msg) {}
};

module.exports = TrafficLightController


