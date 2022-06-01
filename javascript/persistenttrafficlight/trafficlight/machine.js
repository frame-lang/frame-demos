// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files
const FrameEvent = require("../framelang/framelang");
class TrafficLight {
  constructor(manager) {
    // Create and intialize start state compartment.

    this._state = this._sBegin_;
    this._manager = manager;
    this._compartment = new TrafficLightCompartment(this._state);
    this._nextCompartment = null;

    // Initialize domain
    this.flashColor = "";

    // Send system start event
    this._frameEvent = FrameEvent(">", null);
    this._mux_(this._frameEvent);
  }

  //===================== Interface Block ===================//

  stop() {
    let e = FrameEvent("stop", null);
    this._mux_(e);
  }

  tick() {
    let e = FrameEvent("tick", null);
    this._mux_(e);
  }

  systemError() {
    let e = FrameEvent("systemError", null);
    this._mux_(e);
  }

  systemRestart() {
    let e = FrameEvent("systemRestart", null);
    this._mux_(e);
  }

  static loadTrafficLight(manager, data) {
    data._manager = manager;
    data._compartment = data._compartment;

    data.flashColor = data.flashColor;

    return data;
  }


  marshal() {
    let data = this;
    return data;
  }

  //====================== Multiplexer ====================//

  _mux_(e) {
    switch (this._compartment.state) {
      case this._sBegin_:
        this._sBegin_(e);
        break;
      case this._sRed_:
        this._sRed_(e);
        break;
      case this._sGreen_:
        this._sGreen_(e);
        break;
      case this._sYellow_:
        this._sYellow_(e);
        break;
      case this._sFlashingRed_:
        this._sFlashingRed_(e);
        break;
      case this._sEnd_:
        this._sEnd_(e);
        break;
      case this._sWorking_:
        this._sWorking_(e);
        break;
    }

    if (this._nextCompartment != null) {
      let nextCompartment = this._nextCompartment;
      this._nextCompartment = null;
      if (
        nextCompartment._forwardEvent != null &&
        nextCompartment._forwardEvent._message == ">"
      ) {
        this._mux_(FrameEvent("<", this._compartment.ExitArgs));
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

  _sBegin_(e) {
    switch (e._message) {
      case ">": {
        this.startWorkingTimer_do();
        let compartment = new TrafficLightCompartment(this._sRed_);

        this._transition_(compartment);
        return;
      }
    }
  }

  _sRed_(e) {
    switch (e._message) {
      case ">": {
        this.enterRed_do();
        return;
      }

      case "tick": {
        let compartment = new TrafficLightCompartment(this._sGreen_);

        this._transition_(compartment);
        return;
      }
    }
    this._sWorking_(e);
  }

  _sGreen_(e) {
    switch (e._message) {
      case ">": {
        this.enterGreen_do();
        return;
      }

      case "tick": {
        let compartment = new TrafficLightCompartment(this._sYellow_);

        this._transition_(compartment);
        return;
      }
    }
    this._sWorking_(e);
  }

  _sYellow_(e) {
    switch (e._message) {
      case ">": {
        this.enterYellow_do();
        return;
      }

      case "tick": {
        let compartment = new TrafficLightCompartment(this._sRed_);

        this._transition_(compartment);
        return;
      }
    }
    this._sWorking_(e);
  }

  _sFlashingRed_(e) {
    switch (e._message) {
      case ">": {
        this.enterFlashingRed_do();
        this.stopWorkingTimer_do();
        this.startFlashingTimer_do();
        return;
      }

      case "<": {
        this.exitFlashingRed_do();
        this.stopFlashingTimer_do();
        this.startWorkingTimer_do();
        return;
      }

      case "tick": {
        this.changeFlashingAnimation_do();
        return;
      }

      case "systemRestart": {
        let compartment = new TrafficLightCompartment(this._sRed_);

        this._transition_(compartment);
        return;
      }

      case "stop": {
        let compartment = new TrafficLightCompartment(this._sEnd_);

        this._transition_(compartment);
        return;
      }
    }
  }

  _sEnd_(e) {
    switch (e._message) {
      case ">": {
        this.stopWorkingTimer_do();
        return;
      }
    }
  }

  _sWorking_(e) {
    switch (e._message) {
      case "stop": {
        let compartment = new TrafficLightCompartment(this._sEnd_);

        this._transition_(compartment);
        return;
      }

      case "systemError": {
        let compartment = new TrafficLightCompartment(this._sFlashingRed_);

        this._transition_(compartment);
        return;
      }
    }
  }

  //===================== Actions Block ===================//

  // Unimplemented Actions

  enterRed_do() {
    throw new Error("Action not implemented.");
  }
  enterGreen_do() {
    throw new Error("Action not implemented.");
  }
  enterYellow_do() {
    throw new Error("Action not implemented.");
  }
  enterFlashingRed_do() {
    throw new Error("Action not implemented.");
  }
  exitFlashingRed_do() {
    throw new Error("Action not implemented.");
  }
  startWorkingTimer_do() {
    throw new Error("Action not implemented.");
  }
  stopWorkingTimer_do() {
    throw new Error("Action not implemented.");
  }
  startFlashingTimer_do() {
    throw new Error("Action not implemented.");
  }
  stopFlashingTimer_do() {
    throw new Error("Action not implemented.");
  }
  changeColor_do(color) {
    throw new Error("Action not implemented.");
  }
  startFlashing_do() {
    throw new Error("Action not implemented.");
  }
  stopFlashing_do() {
    throw new Error("Action not implemented.");
  }
  changeFlashingAnimation_do() {
    throw new Error("Action not implemented.");
  }
  log_do(msg) {
    throw new Error("Action not implemented.");
  }

  //=============== Machinery and Mechanisms ==============//

  _transition_(compartment) {
    this._nextCompartment = compartment;
  }

  _do_transition_(nextCompartment) {
    this._mux_(FrameEvent("<", this._compartment.ExitArgs));
    this._compartment = nextCompartment;
    this._mux_(FrameEvent(">", this._compartment.EnterArgs));
  }
}

//=============== Compartment ==============//

class TrafficLightCompartment {
  constructor(state) {
    this.state = state;
  }

  StateArgs = {};
  StateVars = {};
  EnterArgs = {};
  ExitArgs = {};
  _forwardEvent = FrameEvent.call(this);
}

// /********************

class TrafficLightController extends TrafficLight {
  constructor(manager) {
    super(manager);
  }
  enterRed_do() {
    this._manager.enterRed();
  }
  enterGreen_do() {
    this._manager.enterGreen();
  }
  enterYellow_do() {
    this._manager.enterYellow();
  }
  enterFlashingRed_do() {}
  exitFlashingRed_do() {}
  startWorkingTimer_do() {}
  stopWorkingTimer_do() {}
  startFlashingTimer_do() {}
  stopFlashingTimer_do() {}
  changeColor_do(color) {}
  startFlashing_do() {}
  stopFlashing_do() {}
  changeFlashingAnimation_do() {}
  log_do(msg) {}
}

module.exports = TrafficLightController;
// ********************/
