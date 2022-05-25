const FrameEvent = require("../fizzbuzz/framelang/framelang");

class CompartmentParams {
    
  constructor (state_param,enter_param) {
      
      // Create and intialize start state compartment.
      
      this._state = this._sS0_;
      this._compartment = new CompartmentParamsCompartment(this._state);
      this._nextCompartment = null;
      this._compartment.StateArgs["state_param"] = state_param;
      this._compartment.StateVars["state_var"] = 100;
      this._compartment.EnterArgs["enter_param"] = enter_param;
      
      // Initialize domain
      
      // Send system start event
      this._frameEvent = FrameEvent(">", this._compartment.EnterArgs);
      this._mux_(this._frameEvent);
  }
  
  //====================== Multiplexer ====================//
  
  _mux_(e) {
      switch (this._compartment.state) {
          case this._sS0_:
              this._sS0_(e);
              break;
          case this._sS1_:
              this._sS1_(e);
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
  
  _sS0_(e) {
      switch (e._message) {
          case ">":
              {
              this.print_do((this._compartment.StateArgs["state_param"]) + " " + (this._compartment.StateVars["state_var"]) + " " + (e._parameters["enter_param"]));
              let compartment =  new CompartmentParamsCompartment(this._sS1_);
              
              compartment._forwardEvent = e;
              compartment.StateArgs["state_param"] = this._compartment.StateArgs["state_param"] + 20;
              compartment.StateVars["state_var"] = 200;
              
              this._transition_(compartment);
              return;
              }
              
      }
  }
  
  _sS1_(e) {
      switch (e._message) {
          case ">":
              {
              this.print_do((this._compartment.StateArgs["state_param"]) + " " + (this._compartment.StateVars["state_var"]) + " " + (e._parameters["enter_param"]));
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

