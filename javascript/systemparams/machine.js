
const FrameEvent = require("./framelang/framelang");

class NewSystemParams {
  
  constructor(message){
    this.message = message;
    this._state_ = "SystemParamsState_Begin"
    this.m = new SystemParamsStruct()
    this.m._compartment_= NewSystemParamsCompartment(this._state_)
    this.params = {"msg" : this.message }
    this.e = FrameEvent
    this.e.message = ">"
    this.e.params = this.params;
    this.e.return = null
    this.m._mux_(this.e)
  }

}

class SystemParamsStruct {
  constructor() {
    this._compartment_ = new SystemParamsCompartment();
    this._nextCompartment_ = new SystemParamsCompartment();
  }

  //====================== Multiplexer ====================//
  _mux_(e) {
    switch (this._compartment_.state) {
      case "SystemParamsState_Begin":
        this._SystemParamsState_Begin_(e)
    }
    if (this._nextCompartment_ != null) {
      let nextCompartment = this._nextCompartment_;
      this._nextCompartment_ = null;
      if(nextCompartment._forwardEvent_ != null && nextCompartment._forwardEvent_.message == ">") {
        const e = FrameEvent
        e.message = "<"
        e.params = this._compartment_.ExitArgs
        e.return = null
        this._mux_(e);
        this._compartment_ = nextCompartment;
        this._mux_(nextCompartment._forwardEvent_)
      } else {
        this._do_transition_(nextCompartment)
        if (nextCompartment._forwardEvent_ != null) {
          this._mux_(nextCompartment._forwardEvent_)
        }
      }
      nextCompartment._forwardEvent_ = null;
    }
  }

  //===================== Machine Block ===================//

  _SystemParamsState_Begin_(e) {
    switch (e.message ){
      case ">":
        this.print(e.params['msg'].toString())
        return
    }
  }

  //=============== Machinery and Mechanisms ==============//
  _transition_(compartment) {
    this._nextCompartment_ = compartment;
  }

  _do_transition_(nextCompartment) {
    const e = FrameEvent;
    e.message = "<";
    e.params = this._compartment_.ExitArgs;
    e.return = null;

    this._mux_(e);

    this._compartment_ = nextCompartment;

    e.message = ">";
    e.params = this._compartment_.EnterArgs;
    e.return = null;

    this._mux_(e);
  }

  //===================== Actions Block ===================//

  print(message){
    console.log(message)
  }
}

//=============== Compartment ==============//

class SystemParamsCompartment {
  constructor(state) {
    this.state = state;
  }
  StateArgs = {};
  StateVars = {};
  EnterArgs = {};
  ExitArgs = {};
  _forwardEvent_ = FrameEvent;
}

const NewSystemParamsCompartment = (state) => {
  const c = new SystemParamsCompartment(state);
  c.StateArgs = {}
	c.StateVars = {}
	c.EnterArgs = {}
	c.ExitArgs = {}
  return c
}


/********************/

let a = new NewSystemParams("Hello")


