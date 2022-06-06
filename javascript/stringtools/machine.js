// emitted from framec_v0.10.0

const FrameEvent = require("./framelang/framelang");

// get include files at https://github.com/frame-lang/frame-ancillary-files

class StringTools {
    
    // creating private properties
    
    #state
    #compartment
    #nextCompartment
    
    
    constructor () {
        
        // Create and intialize start state compartment.
        
        this.#state = this.#sRouter_;
        this.#compartment = new StringToolsCompartment(this.#state);
        this.#nextCompartment = null;
        
        // Initialize domain
        
        // Send system start event
        const frameEvent = FrameEvent(">", null);
        this.#mux(frameEvent);
    }
    
    //===================== Interface Block ===================//
    
    reverse(str) {
        let e = FrameEvent("reverse",{"str":str});
        this.#mux(e);
        return e._return;
    }
    
    makePalindrome(str) {
        let e = FrameEvent("makePalindrome",{"str":str});
        this.#mux(e);
        return e._return;
    }
    
    //====================== Multiplexer ====================//
    
    #mux(e) {
        switch (this.#compartment.state) {
            case this.#sRouter_:
                this.#sRouter_(e);
                break;
            case this.#sReverse_:
                this.#sReverse_(e);
                break;
            case this.#sMakePalindrome_:
                this.#sMakePalindrome_(e);
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
    
    #sRouter_(e) {
        switch (e._message) {
            case "makePalindrome":
                {
                // make\npalindrome
                let compartment =  new StringToolsCompartment(this.#sMakePalindrome_);
                
                compartment._forwardEvent = e;
                
                this.#transition(compartment);
                return;
                }
                
            case "reverse":
                {
                // reverse
                let compartment =  new StringToolsCompartment(this.#sReverse_);
                
                compartment._forwardEvent = e;
                
                this.#transition(compartment);
                return;
                }
                
        }
    }
    
    #sReverse_(e) {
        switch (e._message) {
            case "reverse":
                {
                e._return = this.reverse_str_do((e._parameters["str"]));
                // ready
                let compartment =  new StringToolsCompartment(this.#sRouter_);
                
                
                this.#transition(compartment);
                return;
                }
                
        }
    }
    
    #sMakePalindrome_(e) {
        switch (e._message) {
            case "makePalindrome":
                {
                e._return = (e._parameters["str"]) + this.reverse_str_do((e._parameters["str"]));
                // ready
                let compartment =  new StringToolsCompartment(this.#sRouter_);
                
                
                this.#transition(compartment);
                return;
                }
                
        }
    }
    
    //===================== Actions Block ===================//
    
    // Unimplemented Actions
    
    reverse_str_do(str) { throw new Error('Action not implemented.'); }
    
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
