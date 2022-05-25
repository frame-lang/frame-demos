const FrameEvent = require("../fizzbuzz/framelang/framelang");

// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files

class TransitionEventForwarding {
    
    constructor (cycles) {
        
        // Create and intialize start state compartment.
        
        this._state = this._sStart_;
        this._compartment = new TransitionEventForwardingCompartment(this._state);
        this._nextCompartment = null;
        this._compartment.EnterArgs["cycles"] = cycles;
        
        // Initialize domain
        
        // Send system start event
        this._frameEvent = FrameEvent(">", this._compartment.EnterArgs);
        this._mux_(this._frameEvent);
    }
    
    //====================== Multiplexer ====================//
    
    _mux_(e) {
        switch (this._compartment.state) {
            case this._sStart_:
                this._sStart_(e);
                break;
            case this._sForwardEventAgain_:
                this._sForwardEventAgain_(e);
                break;
            case this._sDecrement_:
                this._sDecrement_(e);
                break;
            case this._sStop_:
                this._sStop_(e);
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
    
    _sStart_(e) {
        switch (e._message) {
            case ">":
                {
                if ((e._parameters["cycles"]) == 0) {
                    this._compartment.ExitArgs["msg"] = "stopping";
                    let compartment =  new TransitionEventForwardingCompartment(this._sStop_);
                    
                    compartment._forwardEvent = e;
                    
                    this._transition_(compartment);
                    return;
                } else {
                    this._compartment.ExitArgs["msg"] = "keep going";
                    let compartment =  new TransitionEventForwardingCompartment(this._sForwardEventAgain_);
                    
                    compartment._forwardEvent = e;
                    
                    this._transition_(compartment);
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
    
    _sForwardEventAgain_(e) {
        switch (e._message) {
            case ">":
                {
                let compartment =  new TransitionEventForwardingCompartment(this._sDecrement_);
                
                compartment._forwardEvent = e;
                
                this._transition_(compartment);
                return;
                }
                
        }
    }
    
    _sDecrement_(e) {
        switch (e._message) {
            case ">":
                {
                this.print_do((e._parameters["cycles"]).toString());
                let compartment =  new TransitionEventForwardingCompartment(this._sStart_);
                
                compartment.EnterArgs["cycles"] = (e._parameters["cycles"] - 1);
                
                this._transition_(compartment);
                return;
                }
                
        }
    }
    
    _sStop_(e) {
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


// /********************

class TransitionEventForwardingController extends TransitionEventForwarding {

	constructor(cycles) {
	  super(cycles)
	}
	print_do(msg) {
        console.log(msg)
    }
};

// ********************/



new TransitionEventForwardingController(5)