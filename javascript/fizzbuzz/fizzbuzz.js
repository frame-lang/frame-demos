const FrameEvent = require("./framelang/framelang");

// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files

class FizzBuzz {
    
    // creating private properties
    
    #state
    #compartment
    #nextCompartment
    
    
    constructor () {
        
        // Create and intialize start state compartment.
        
        this.#state = this.#sBegin_;
        this.#compartment = new FizzBuzzCompartment(this.#state);
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
    
    //====================== Multiplexer ====================//
    
    #mux(e) {
        switch (this.#compartment.state) {
            case this.#sBegin_:
                this.#sBegin_(e);
                break;
            case this.#sFizz_:
                this.#sFizz_(e);
                break;
            case this.#sBuzz_:
                this.#sBuzz_(e);
                break;
            case this.#sDigit_:
                this.#sDigit_(e);
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
    
    #sBegin_(e) {
        switch (e._message) {
            case ">>":
                {
                // start
                let compartment =  new FizzBuzzCompartment(this.#sFizz_);
                
                compartment.EnterArgs["i"] = 1;
                
                this.#transition(compartment);
                return;
                }
                
        }
    }
    
    #sFizz_(e) {
        switch (e._message) {
            case ">":
                {
                if (this.gt_100_do((e._parameters["i"]))) {
                    // i > 100
                    let compartment =  new FizzBuzzCompartment(this.#sEnd_);
                    
                    
                    this.#transition(compartment);
                    return;
                }
                if (this.mod3_eq0_do((e._parameters["i"]))) {
                    this.print_do("Fizz");
                    // i % 3 == 0
                    let compartment =  new FizzBuzzCompartment(this.#sBuzz_);
                    
                    compartment.EnterArgs["i"] = e._parameters["i"];
                    compartment.EnterArgs["fizzed"] = true;
                    
                    this.#transition(compartment);
                } else {
                    // i % 3 != 0
                    let compartment =  new FizzBuzzCompartment(this.#sBuzz_);
                    
                    compartment.EnterArgs["i"] = e._parameters["i"];
                    compartment.EnterArgs["fizzed"] = false;
                    
                    this.#transition(compartment);
                }
                return;
                }
                
        }
    }
    
    #sBuzz_(e) {
        switch (e._message) {
            case ">":
                {
                if (this.mod5_eq0_do((e._parameters["i"]))) {
                    this.print_do("Buzz");
                    // i % 5 == 0
                    this.#compartment.ExitArgs["output"] = " ";
                    let compartment =  new FizzBuzzCompartment(this.#sFizz_);
                    
                    compartment.EnterArgs["i"] = this.plus_1_do(e._parameters["i"]);
                    
                    this.#transition(compartment);
                    return;
                }
                if ((e._parameters["fizzed"])) {
                    // fizzed
                    this.#compartment.ExitArgs["output"] = " ";
                    let compartment =  new FizzBuzzCompartment(this.#sFizz_);
                    
                    compartment.EnterArgs["i"] = this.plus_1_do(e._parameters["i"]);
                    
                    this.#transition(compartment);
                    return;
                }
                // ! mod3 | mod5
                this.#compartment.ExitArgs["output"] = "";
                let compartment =  new FizzBuzzCompartment(this.#sDigit_);
                
                compartment.EnterArgs["i"] = e._parameters["i"];
                
                this.#transition(compartment);
                return;
                }
                
            case "<":
                {
                this.print_do((e._parameters["output"]));
                return;
                }
                
        }
    }
    
    #sDigit_(e) {
        switch (e._message) {
            case ">":
                {
                this.print_do((e._parameters["i"]).toString());
                this.print_do(" ");
                // loop
                let compartment =  new FizzBuzzCompartment(this.#sFizz_);
                
                compartment.EnterArgs["i"] = this.plus_1_do(e._parameters["i"]);
                
                this.#transition(compartment);
                return;
                }
                
        }
    }
    
    #sEnd_(e) {
        switch (e._message) {
        }
    }
    
    //===================== Actions Block ===================//
    
    // Unimplemented Actions
    
    print_do(msg) { throw new Error('Action not implemented.'); }
    gt_100_do(i) { throw new Error('Action not implemented.'); }
    mod3_eq0_do(i) { throw new Error('Action not implemented.'); }
    mod5_eq0_do(i) { throw new Error('Action not implemented.'); }
    plus_1_do(i) { throw new Error('Action not implemented.'); }
    
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

class FizzBuzzCompartment {
  constructor(state) {
    this.state = state;
  }

  StateArgs = {};
  StateVars = {};
  EnterArgs = {};
  ExitArgs = {};
  _forwardEvent = FrameEvent.call(this);
}

class FizzBuzzController extends FizzBuzz {
  constructor() {
    super();
    
  }
  print_do(msg) {
    process.stdout.write(msg.toString());
  }
  gt_100_do(i) {
    return i > 100;
  }
  mod3_eq0_do(i) {
    return i % 3 == 0;
  }
  mod5_eq0_do(i) {
    return i % 5 == 0;
  }
  plus_1_do(i) {
    return i + 1;
  }
}

module.exports = FizzBuzzController

