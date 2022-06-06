const FrameEvent = require("../fizzbuzz/framelang/framelang");

// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files

class CompartmentParams {
    
    // creating private properties
    
    #state
    #compartment
    #nextCompartment
    
    
    constructor (state_param,enter_param) {
        
        // Create and intialize start state compartment.
        
        this.#state = this.#sS0_;
        this.#compartment = new CompartmentParamsCompartment(this.#state);
        this.#nextCompartment = null;
        this.#compartment.StateArgs["state_param"] = state_param;
        this.#compartment.StateVars["state_var"] = 100;
        this.#compartment.EnterArgs["enter_param"] = enter_param;
        
        // Initialize domain
        
        // Send system start event
        const frameEvent = FrameEvent(">", this.#compartment.EnterArgs);
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
                this.print_do((this.#compartment.StateArgs["state_param"]) + " " + (this.#compartment.StateVars["state_var"]) + " " + (e._parameters["enter_param"]));
                let compartment =  new CompartmentParamsCompartment(this.#sS1_);
                
                compartment._forwardEvent = e;
                compartment.StateArgs["state_param"] = this.#compartment.StateArgs["state_param"] + 20;
                compartment.StateVars["state_var"] = 200;
                
                this.#transition(compartment);
                return;
                }
                
        }
    }
    
    #sS1_(e) {
        switch (e._message) {
            case ">":
                {
                this.print_do((this.#compartment.StateArgs["state_param"]) + " " + (this.#compartment.StateVars["state_var"]) + " " + (e._parameters["enter_param"]));
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

class CompartmentParamsCompartment {

  constructor(state) {
      this.state = state
  }
  
  StateArgs = {};
  StateVars = {};
  EnterArgs = {};
  ExitArgs = {};
  _forwardEvent = FrameEvent.call(this)
}




class CompartmentParamsController extends CompartmentParams {

constructor(state_param,enter_param) {
  super(state_param,enter_param)
  
}
};


module.exports = CompartmentParamsController;

