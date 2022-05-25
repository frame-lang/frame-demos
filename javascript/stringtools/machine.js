// emitted from framec_v0.10.0

const FrameEvent = require("./framelang/framelang");

// get include files at https://github.com/frame-lang/frame-ancillary-files

class StringTools {
    
    constructor () {
        
        // Create and intialize start state compartment.
        
        this._state = this._sRouter_;
        this._compartment = new StringToolsCompartment(this._state);
        this._nextCompartment = null;
        
        // Initialize domain
        
        // Send system start event
        this._frameEvent = FrameEvent(">", null);
        this._mux_(this._frameEvent);
    }
    
    //===================== Interface Block ===================//
    
    reverse(str) {
        let e = FrameEvent("reverse",{"str":str});
        this._mux_(e);
        return e._return;
    }
    
    makePalindrome(str) {
        let e = FrameEvent("makePalindrome",{"str":str});
        this._mux_(e);
        return e._return;
    }
    
    //====================== Multiplexer ====================//
    
    _mux_(e) {
        switch (this._compartment.state) {
            case this._sRouter_:
                this._sRouter_(e);
                break;
            case this._sReverse_:
                this._sReverse_(e);
                break;
            case this._sMakePalindrome_:
                this._sMakePalindrome_(e);
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
    
    _sRouter_(e) {
        switch (e._message) {
            case "makePalindrome":
                {
                // make\npalindrome
                let compartment =  new StringToolsCompartment(this._sMakePalindrome_);
                
                compartment._forwardEvent = e;
                
                this._transition_(compartment);
                return;
                }
                
            case "reverse":
                {
                // reverse
                let compartment =  new StringToolsCompartment(this._sReverse_);
                
                compartment._forwardEvent = e;
                
                this._transition_(compartment);
                return;
                }
                
        }
    }
    
    _sReverse_(e) {
        switch (e._message) {
            case "reverse":
                {
                e._return = this.reverse_str_do((e._parameters["str"]));
                // ready
                let compartment =  new StringToolsCompartment(this._sRouter_);
                
                
                this._transition_(compartment);
                return;
                }
                
        }
    }
    
    _sMakePalindrome_(e) {
        switch (e._message) {
            case "makePalindrome":
                {
                e._return = (e._parameters["str"]) + this.reverse_str_do((e._parameters["str"]));
                // ready
                let compartment =  new StringToolsCompartment(this._sRouter_);
                
                
                this._transition_(compartment);
                return;
                }
                
        }
    }
    
    //===================== Actions Block ===================//
    
    // Unimplemented Actions
    
    reverse_str_do(str) { throw new Error('Action not implemented.'); }
    
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

class StringToolsCompartment {

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

class StringToolsController extends StringTools {

	constructor() {
	  super()
	}
	reverse_str_do(str) {
		let result = ""
		for (let char of str) {
			result = char + result
		}
		return result
	}
};

module.exports = StringToolsController
// ********************/
