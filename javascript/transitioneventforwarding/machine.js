const FrameEvent = require("../fizzbuzz/framelang/framelang");

// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files

class TransitionEventForwarding {
    
    // creating private properties
    
    #state
    #compartment
    #nextCompartment
    
    
    constructor (cycles) {
        
        // Create and intialize start state compartment.
        
        this.#state = this.#sStart_;
        this.#compartment = new TransitionEventForwardingCompartment(this.#state);
        this.#nextCompartment = null;
        this.#compartment.EnterArgs["cycles"] = cycles;
        
        // Initialize domain
        
        // Send system start event
        const frameEvent = FrameEvent(">", this.#compartment.EnterArgs);
        this.#mux(frameEvent);
    }
    
    //====================== Multiplexer ====================//
    
    #mux(e) {
        switch (this.#compartment.state) {
            case this.#sStart_:
                this.#sStart_(e);
                break;
            case this.#sForwardEventAgain_:
                this.#sForwardEventAgain_(e);
                break;
            case this.#sDecrement_:
                this.#sDecrement_(e);
                break;
            case this.#sStop_:
                this.#sStop_(e);
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
            case ">":
                {
                if ((e._parameters["cycles"]) == 0) {
                    this.#compartment.ExitArgs["msg"] = "stopping";
                    let compartment =  new TransitionEventForwardingCompartment(this.#sStop_);
                    
                    compartment._forwardEvent = e;
                    
                    this.#transition(compartment);
                    return;
                } else {
                    this.#compartment.ExitArgs["msg"] = "keep going";
                    let compartment =  new TransitionEventForwardingCompartment(this.#sForwardEventAgain_);
                    
                    compartment._forwardEvent = e;
                    
                    this.#transition(compartment);
                }
                return;
                }
                
            case "<":
                {
                this.print_do((e._parameters["msg"]));
                return;
                }
                
        }
    }
    
    #sForwardEventAgain_(e) {
        switch (e._message) {
            case ">":
                {
                let compartment =  new TransitionEventForwardingCompartment(this.#sDecrement_);
                
                compartment._forwardEvent = e;
                
                this.#transition(compartment);
                return;
                }
                
        }
    }
    
    #sDecrement_(e) {
        switch (e._message) {
            case ">":
                {
                this.print_do((e._parameters["cycles"]).toString());
                let compartment =  new TransitionEventForwardingCompartment(this.#sStart_);
                
                compartment.EnterArgs["cycles"] = (e._parameters["cycles"] - 1);
                
                this.#transition(compartment);
                return;
                }
                
        }
    }
    
    #sStop_(e) {
        switch (e._message) {
            case ">":
                {
                this.print_do((e._parameters["cycles"]).toString());
                this.print_do("done");
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
    
    
    
};

//=============== Compartment ==============//

class TransitionEventForwardingCompartment {

    constructor(state) {
        this.state = state
    }
    
    StateArgs = {};
    StateVars = {};
    EnterArgs = {};
    ExitArgs = {};
    _forwardEvent = FrameEvent.call(this)
}




class TransitionEventForwardingController extends TransitionEventForwarding {

	constructor(cycles) {
	  super(cycles)
	}
	print_do(msg) {
        console.log(msg)
    }
};



module.exports = TransitionEventForwardingController
