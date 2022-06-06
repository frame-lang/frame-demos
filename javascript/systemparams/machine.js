const FrameEvent = require("./framelang/framelang");

// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files

class SystemParams {
    
    // creating private properties
    
    #state
    #compartment
    #nextCompartment
    
    
    constructor (stateMsg,enterMsg) {
        
        // Create and intialize start state compartment.
        
        this.#state = this.#sBegin_;
        this.#compartment = new SystemParamsCompartment(this.#state);
        this.#nextCompartment = null;
        this.#compartment.StateArgs["stateMsg"] = stateMsg;
        this.#compartment.EnterArgs["enterMsg"] = enterMsg;
        
        // Initialize domain
        
        // Send system start event
        const frameEvent = FrameEvent(">", this.#compartment.EnterArgs);
        this.#mux(frameEvent);
    }
    
    //====================== Multiplexer ====================//
    
    #mux(e) {
        switch (this.#compartment.state) {
            case this.#sBegin_:
                this.#sBegin_(e);
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
            case ">":
                {
                this.print_do((this.#compartment.StateArgs["stateMsg"]) + " " + (e._parameters["enterMsg"]));
                return;
                }
                
        }
    }
    
    //===================== Actions Block ===================//
    
    print_do (msg) {
        
        console.log(msg)
    
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

class SystemParamsCompartment {

  constructor(state) {
      this.state = state
  }
  
  StateArgs = {};
  StateVars = {};
  EnterArgs = {};
  ExitArgs = {};
  _forwardEvent = FrameEvent.call(this)
}




class SystemParamsController extends SystemParams {

constructor(stateMsg,enterMsg) {
  super(stateMsg,enterMsg)
}
};


module.exports = SystemParamsController


