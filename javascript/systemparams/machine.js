const FrameEvent = require("./framelang/framelang");

// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files

class SystemParams {
    
  constructor (stateMsg,enterMsg) {
      
      // Create and intialize start state compartment.
      
      this._state = this._sBegin_;
      this._compartment = new SystemParamsCompartment(this._state);
      this._nextCompartment = null;
      this._compartment.StateArgs["stateMsg"] = stateMsg;
      this._compartment.EnterArgs["enterMsg"] = enterMsg;
      
      // Initialize domain
      
      // Send system start event
      this._frameEvent = FrameEvent(">", this._compartment.EnterArgs);
      this._mux_(this._frameEvent);
  }
  
  //====================== Multiplexer ====================//
  
  _mux_(e) {
      switch (this._compartment.state) {
          case this._sBegin_:
              this._sBegin_(e);
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
  
  _sBegin_(e) {
      switch (e._message) {
          case ">":
              {
              this.print_do((this._compartment.StateArgs["stateMsg"]) + " " + (e._parameters["enterMsg"]));
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




let a = new SystemParamsController("Hello world", "Deepak");
