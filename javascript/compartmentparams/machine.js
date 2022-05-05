const FrameEvent = require("../fizzbuzz/framelang/framelang");

class CompartmentParams {
  constructor(state_param, enter_param) {
    this._compartment = this.CompartmentParamsCompartment(
      this.CompartmentParamsState_S0
    );
    this._nextCompartment = null;
    this._compartment.StateArgs["state_param"] = state_param;
    this._compartment.StateVars["state_var"] = 100;
    this._e = FrameEvent(">", { enter_param: enter_param });
    this._mux_(this._e);
  }
  //====================== Multiplexer ====================//
  _mux_(e) {
    switch (this._compartment.state) {
      case this.CompartmentParamsState_S0:
        this.CompartmentParamsState_S0(e);
        break;
      case this.CompartmentParamsState_S1:
        this.CompartmentParamsState_S1(e);
        break;
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
  CompartmentParamsState_S0(e) {
    switch (e._message) {
      case ">":
        this.print(
          parseInt(this._compartment.StateArgs["state_param"]) +
            " " +
            parseInt(this._compartment.StateVars["state_var"]) +
            " " +
            e._parameters["enter_param"]
        );
        let compartment = this.CompartmentParamsCompartment(
          this.CompartmentParamsState_S1
        );
        compartment._forwardEvent = e;
        compartment.StateArgs["state_param"] =
          parseInt(this._compartment.StateArgs["state_param"]) + 20;
        compartment.StateVars["state_var"] = 200;
        this._transition_(compartment);
        return;
    }
  }
  CompartmentParamsState_S1(e) {
    switch (e._message) {
      case ">":
        this.print(
          parseInt(this._compartment.StateArgs["state_param"]) +
            " " +
            parseInt(this._compartment.StateVars["state_var"]) +
            " " +
            e._parameters["enter_param"]
        );
        return;
    }
  }
  //=============== Machinery and Mechanisms ==============//
  _transition_(compartment) {
    this._nextCompartment = compartment;
  }

  _do_transition_(nextCompartment) {
    const e = FrameEvent("<", this._compartment.ExitArgs);
    this._mux_(e);

    this._compartment = nextCompartment;

    e._message = ">";
    e._parameters = this._compartment.EnterArgs;
    e._return = null;

    this._mux_(e);
  }
  //===================== Actions Block ===================//
  // Unimplemented Actions
  print(s) {
    throw new Error("Action is not implemented");
  }
  //=============== Compartment ==============//

  CompartmentParamsCompartment(state) {
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

// ***********************************

class CompartmentParamsController extends CompartmentParams {
  constructor(state_param, enter_param) {
    super(state_param, enter_param);
  }

  print(s) {
    console.log(s);
  }
}

new CompartmentParamsController(1, 2);
