// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files

const FrameEvent = require("./framelang/framelang");

class HistoryStateContext {
    
    // creating private properties
    
    #state
    #compartment
    #nextCompartment
    #stateStack
    
    constructor () {
        
        // Create state stack.
        
        this.#stateStack = [];
        
        // Create and intialize start state compartment.
        
        this.#state = this.#sStart_;
        this.#compartment = new HistoryStateContextCompartment(this.#state);
        this.#nextCompartment = null;
        
        // Initialize domain
        
        // Send system start event
        const frameEvent = FrameEvent(">", null);
        this.#mux(frameEvent);
    }
    
    //===================== Interface Block ===================//
    
    start() {
        let e = FrameEvent(">>",null);
        this.#mux(e);
    }
    
    switchState() {
        let e = FrameEvent("switchState",null);
        this.#mux(e);
    }
    
    gotoDeadEnd() {
        let e = FrameEvent("gotoDeadEnd",null);
        this.#mux(e);
    }
    
    back() {
        let e = FrameEvent("back",null);
        this.#mux(e);
    }
    
    //====================== Multiplexer ====================//
    
    #mux(e) {
        switch (this.#compartment.state) {
            case this.#sStart_:
                this.#sStart_(e);
                break;
            case this.#sS0_:
                this.#sS0_(e);
                break;
            case this.#sS1_:
                this.#sS1_(e);
                break;
            case this.#sDeadEnd_:
                this.#sDeadEnd_(e);
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
    
    #sStart_(e) {
        switch (e._message) {
            case ">>":
                {
                let compartment =  new HistoryStateContextCompartment(this.#sS0_);
                
                compartment.StateVars["enterMsg"] = "Enter S0";
                
                this.#transition(compartment);
                return;
                }
                
        }
    }
    
    #sS0_(e) {
        switch (e._message) {
            case ">":
                {
                this.print_do((this.#compartment.StateVars["enterMsg"]));
                return;
                }
                
            case "switchState":
                {
                // Switch\nState
                let compartment =  new HistoryStateContextCompartment(this.#sS1_);
                
                
                this.#transition(compartment);
                return;
                }
                
            case "gotoDeadEnd":
                {
                this.#stateStack_push(this.#compartment);
                // Goto\nDead End
                let compartment =  new HistoryStateContextCompartment(this.#sDeadEnd_);
                
                
                this.#transition(compartment);
                return;
                }
                
        }
    }
    
    #sS1_(e) {
        switch (e._message) {
            case ">":
                {
                this.print_do("Enter S1");
                return;
                }
                
            case "switchState":
                {
                // Switch\nState
                let compartment =  new HistoryStateContextCompartment(this.#sS0_);
                
                compartment.StateVars["enterMsg"] = "Enter S0";
                
                this.#transition(compartment);
                return;
                }
                
            case "gotoDeadEnd":
                {
                this.#stateStack_push(this.#compartment);
                // Goto\nDead End
                let compartment =  new HistoryStateContextCompartment(this.#sDeadEnd_);
                
                
                this.#transition(compartment);
                return;
                }
                
        }
    }
    
    #sDeadEnd_(e) {
        switch (e._message) {
            case ">":
                {
                this.print_do("Enter $DeadEnd");
                return;
                }
                
            case "back":
                {
                // Go Back
                let compartment = this.#stateStack_pop()
                this.#transition(compartment)
                return;
                }
                
        }
    }
    
    //===================== Actions Block ===================//
    
    // Unimplemented Actions
    
    print_do(msg) { throw new Error('Action not implemented.'); }
    
    //=============== Machinery and Mechanisms ==============//
    
    #transition(compartment) {
        this.#nextCompartment = compartment;
    }
    
    #doTransition(nextCompartment) {
        this.#mux(FrameEvent("<", this.#compartment.ExitArgs));
        this.#compartment = nextCompartment;
        this.#mux(FrameEvent(">", this.#compartment.EnterArgs));
    }
    
    #stateStack_push(compartment) {
        this.#stateStack.push(compartment);
    }
    
    #stateStack_pop(){
        return this.#stateStack.pop();
    }
    
    
    
};

//=============== Compartment ==============//

class HistoryStateContextCompartment {

    constructor(state) {
        this.state = state
    }
    
    StateArgs = {};
    StateVars = {};
    EnterArgs = {};
    ExitArgs = {};
    _forwardEvent = FrameEvent.call(this)
}



class HistoryStateContextController extends HistoryStateContext {

	constructor() {
	  super()
	}
	print_do(msg) {
        console.log(msg)
    }
};

module.exports = HistoryStateContextController
