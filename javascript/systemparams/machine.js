const FrameEvent = require("./framelang/framelang");

class SystemParams {
  
  _nextCompartment = null
  constructor(message) {
    this._e = FrameEvent.call(this, '>', {"msg" : message});
    this._state = this._SystemParamsState_Begin_
    this._compartment = this.systemParamsCompartment(this._state)
    this._mux_(this._e)
  }

  //====================== Multiplexer ====================//
  _mux_(e) {
    switch (this._compartment.state) {
      case this._SystemParamsState_Begin_:
        this._SystemParamsState_Begin_(e);
    }
    if (this._nextCompartment != null) {
      let nextCompartment = this._nextCompartment;
      this._nextCompartment = null;
      if (
        nextCompartment._forwardEvent != null &&
        nextCompartment._forwardEvent._message == ">"
      ) {
        const e = FrameEvent.call(this);
        e._message = "<";
        e._parameters = this._compartment.ExitArgs;
        e._return = null;
        this._mux_(e);
        this._compartment = nextCompartment;
        this._mux_(nextCompartment._forwardEvent);
      } else {
        this._do_transition_(nextCompartment);
        if (nextCompartment._forwardEvent != null) {
          this._mux_(nextCompartment._forwardEvent);
        }
      }
      nextCompartment._forwardEvent = null;
    }
  }

  //===================== Machine Block ===================//

  _SystemParamsState_Begin_(e) {
    switch (e._message) {
      case ">":
        this.print(e._parameters["msg"].toString());
        return;
    }
  }

  //=============== Machinery and Mechanisms ==============//
  _transition_(compartment) {
    this._nextCompartment = compartment;
  }

  _do_transition_(nextCompartment) {
    const e = FrameEvent.call(this );
    e._message = "<";
    e._parameters = this._compartment.ExitArgs;
    e._return = null;

    this._mux_(e);

    this._compartment = nextCompartment;

    e._message = ">";
    e._parameters = this._compartment.EnterArgs;
    e._return = null;

    this._mux_(e);
  }

  //===================== Actions Block ===================//

  print(msg) {
    throw new Error('Action not implemented.');
  }

  //=============== Compartment ==============//

  systemParamsCompartment(state) {
    let that = {};

    that.state = state;
    that.StateArgs = {};
    that.StateVars = {};
    that.EnterArgs = {};
    that.ExitArgs = {};
    that._forwardEvent = FrameEvent.call(this);

    return that;
  }
}


// ****************************************************************//

class Controller extends SystemParams {
    constructor(message){
      super(message)
    }
  
    print(msg){
      console.log(msg)
    }
  }
  
let a = new Controller("Hello world")