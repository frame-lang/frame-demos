const FrameEvent = require("./framelang/framelang");

// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files

class FizzBuzz {
  constructor() {
    // Create and intialize start state compartment.

    this._state = this._sBegin_;
    this._compartment = new FizzBuzzCompartment(this._state);
    this._nextCompartment = null;

    // Initialize domain

    // Send system start event
    this._frameEvent = FrameEvent(">", null);
    this._mux_(this._frameEvent);
  }

  //===================== Interface Block ===================//

  start() {
    let e = FrameEvent(">>", null);
    this._mux_(e);
  }

  //====================== Multiplexer ====================//

  _mux_(e) {
    switch (this._compartment.state) {
      case this._sBegin_:
        this._sBegin_(e);
        break;
      case this._sFizz_:
        this._sFizz_(e);
        break;
      case this._sBuzz_:
        this._sBuzz_(e);
        break;
      case this._sDigit_:
        this._sDigit_(e);
        break;
      case this._sEnd_:
        this._sEnd_(e);
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
      case ">>": {
        // start
        let compartment = new FizzBuzzCompartment(this._sFizz_);

        compartment.EnterArgs["i"] = 1;

        this._transition_(compartment);
        return;
      }
    }
  }

  _sFizz_(e) {
    switch (e._message) {
      case ">": {
        if (this.gt_100_do(e._parameters["i"])) {
          // i > 100
          let compartment = new FizzBuzzCompartment(this._sEnd_);

          this._transition_(compartment);
          return;
        }
        if (this.mod3_eq0_do(e._parameters["i"])) {
          this.print_do("Fizz");
          // i % 3 == 0
          let compartment = new FizzBuzzCompartment(this._sBuzz_);

          compartment.EnterArgs["i"] = e._parameters["i"];
          compartment.EnterArgs["fizzed"] = true;

          this._transition_(compartment);
        } else {
          // i % 3 != 0
          let compartment = new FizzBuzzCompartment(this._sBuzz_);

          compartment.EnterArgs["i"] = e._parameters["i"];
          compartment.EnterArgs["fizzed"] = false;

          this._transition_(compartment);
        }
        return;
      }
    }
  }

  _sBuzz_(e) {
    switch (e._message) {
      case ">": {
        if (this.mod5_eq0_do(e._parameters["i"])) {
          this.print_do("Buzz");
          // i % 5 == 0
          this._compartment.ExitArgs["output"] = " ";
          let compartment = new FizzBuzzCompartment(this._sFizz_);

          compartment.EnterArgs["i"] = this.plus_1_do(e._parameters["i"]);

          this._transition_(compartment);
          return;
        }
        if (e._parameters["fizzed"]) {
          // fizzed
          this._compartment.ExitArgs["output"] = " ";
          let compartment = new FizzBuzzCompartment(this._sFizz_);

          compartment.EnterArgs["i"] = this.plus_1_do(e._parameters["i"]);

          this._transition_(compartment);
          return;
        }
        // ! mod3 | mod5
        this._compartment.ExitArgs["output"] = "";
        let compartment = new FizzBuzzCompartment(this._sDigit_);

        compartment.EnterArgs["i"] = e._parameters["i"];

        this._transition_(compartment);
        return;
      }

      case "<": {
        this.print_do(e._parameters["output"]);
        return;
      }
    }
  }

  _sDigit_(e) {
    switch (e._message) {
      case ">": {
        this.print_do(e._parameters["i"]);
        this.print_do(" ");
        // loop
        let compartment = new FizzBuzzCompartment(this._sFizz_);

        compartment.EnterArgs["i"] = this.plus_1_do(e._parameters["i"]);

        this._transition_(compartment);
        return;
      }
    }
  }

  _sEnd_(e) {
    switch (e._message) {
    }
  }

  //===================== Actions Block ===================//

  // Unimplemented Actions

  print_do(msg) {
    throw new Error("Action not implemented.");
  }
  gt_100_do(i) {
    throw new Error("Action not implemented.");
  }
  mod3_eq0_do(i) {
    throw new Error("Action not implemented.");
  }
  mod5_eq0_do(i) {
    throw new Error("Action not implemented.");
  }
  plus_1_do(i) {
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

let fizzbuzz = new FizzBuzzController();

fizzbuzz.start();
