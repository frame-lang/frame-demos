const FrameEvent = require("./framelang/framelang");
// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files

class Countdown {
    
    // creating private properties
    
    #state
    #compartment
    #nextCompartment
    
    
    constructor (i) {
        
        // Create and intialize start state compartment.
        
        this.#state = this.#sS0_;
        this.#compartment = new CountdownCompartment(this.#state);
        this.#nextCompartment = null;
        this.#compartment.StateArgs["i"] = i;
        this.#compartment.StateVars["dec"] = 1;
        
        // Initialize domain
        
        // Send system start event
        const frameEvent = FrameEvent(">", null);
        this.#mux(frameEvent);
    }
    
    //====================== Multiplexer ====================//
    
    #mux(e) {
        switch (this.#compartment.state) {
            case this.#sS0_:
                this.#sS0_(e);
                break;
            case this.#sS1_:
                this.#sS1_(e);
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
    
    #sS0_(e) {
        switch (e._message) {
            case ">":
                {
                (this.#compartment.StateArgs["i"]) = (this.#compartment.StateArgs["i"]) - (this.#compartment.StateVars["dec"]);
                this.print_do((this.#compartment.StateArgs["i"]).toString());
                if ((this.#compartment.StateArgs["i"]) == 0) {
                    let compartment =  new CountdownCompartment(this.#sStop_);
                    
                    
                    this.#transition(compartment);
                    return;
                }
                let compartment =  new CountdownCompartment(this.#sS1_);
                
                compartment.EnterArgs["i"] = this.#compartment.StateArgs["i"];
                
                this.#transition(compartment);
                return;
                }
                
        }
    }
    
    #sS1_(e) {
        switch (e._message) {
            case ">":
                {
                let compartment =  new CountdownCompartment(this.#sS0_);
                
                compartment.StateArgs["i"] = e._parameters["i"];
                compartment.StateVars["dec"] = 1;
                
                this.#transition(compartment);
                return;
                }
                
        }
    }
    
    #sStop_(e) {
        switch (e._message) {
            case ">":
                {
                this.print_do("done");
                return;
                }
                
        }
    }
    
    //===================== Actions Block ===================//
    
    print_do (s) {
        console.log(s)
    }
    
    // Unimplemented Actions
    
    
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

class CountdownCompartment {

    constructor(state) {
        this.state = state
    }
    
    StateArgs = {};
    StateVars = {};
    EnterArgs = {};
    ExitArgs = {};
    _forwardEvent = FrameEvent.call(this)
}




class CountdownController extends Countdown {

	constructor(i) {
	  super(i)
	}
};

module.exports = CountdownController


