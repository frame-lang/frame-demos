const FrameEvent = require("./framelang/framelang");

class FizzBuzz {
  
  constructor() {
    this._state = this._FizzBuzzState_Begin_;
    this._compartment = this.fizzBuzzCompartment(this._state);
    this._nextCompartment = null
    this._e = FrameEvent(">");
    this._mux_(this._e);
  }

  

  //===================== Interface Block ===================//
  start() {
    const e = FrameEvent(">>", null);
    this._mux_(e);
  }
  //====================== Multiplexer ====================//
  _mux_(e) {
    switch (this._compartment.state) {
      case this._FizzBuzzState_Begin_:
        this._FizzBuzzState_Begin_(e);
        break;
      case this._FizzBuzzState_Fizz_:
        this._FizzBuzzState_Fizz_(e);
        break;
      case this._FizzBuzzState_Buzz_:
        this._FizzBuzzState_Buzz_(e);
        break;
      case this._FizzBuzzState_Digit_:
        this._FizzBuzzState_Digit_(e);
        break;
      case this._FizzBuzzState_End_:
        this._FizzBuzzState_End_(e);
        break;
    }

    if (this._nextCompartment != null) {
      let nextCompartment = this._nextCompartment;
      this._nextCompartment = null;
      if (
        nextCompartment._forwardEvent != null &&
        nextCompartment._forwardEvent._message == ">"
      ) {
        const e = FrameEvent.call(this)
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
  _FizzBuzzState_Begin_(e) {
    switch (e._message) {
      case ">>":
        let compartment = this.fizzBuzzCompartment(this._FizzBuzzState_Fizz_);
        compartment.EnterArgs["i"] = 1;
        this._transition_(compartment);
        break;
    }
    return;
  }
  _FizzBuzzState_Fizz_(e) {
    switch (e._message) {
      case ">":
        if (this.gt_100(parseInt(e._parameters["i"]))) {
          // i > 100
          let compartment = this.fizzBuzzCompartment(this._FizzBuzzState_End_);
          this._transition_(compartment);
          return;
        }
        if (this.mod3_eq0(parseInt(e._parameters["i"]))) {
          this.print("Fizz");
          // i % 3 == 0
          let compartment = this.fizzBuzzCompartment(this._FizzBuzzState_Buzz_);
          compartment.EnterArgs["i"] = parseInt(e._parameters["i"]);
          compartment.EnterArgs["fizzed"] = true;
          this._transition_(compartment);
        } else {
          // i % 3 != 0
          let compartment = this.fizzBuzzCompartment(this._FizzBuzzState_Buzz_);
          compartment.EnterArgs["i"] = parseInt(e._parameters["i"]);
          compartment.EnterArgs["fizzed"] = false;
          this._transition_(compartment);
        }
        return;
    }
  }
  _FizzBuzzState_Buzz_(e) {
    switch (e._message) {
      case ">":
        if (this.mod5_eq0(parseInt(e._parameters["i"]))) {
          this.print("Buzz");
          // i % 5 == 0
          this._compartment.ExitArgs["output"] = " ";
          let compartment = this.fizzBuzzCompartment(this._FizzBuzzState_Fizz_);
          compartment.EnterArgs["i"] = this.plus_1(
            parseInt(e._parameters["i"])
          );
          this._transition_(compartment);
          return;
        }

        if (Boolean(e._parameters["fizzed"])) {
          // fizzed
          this._compartment.ExitArgs["output"] = " ";
          let compartment = this.fizzBuzzCompartment(this._FizzBuzzState_Fizz_);
          compartment.EnterArgs["i"] = this.plus_1(
            parseInt(e._parameters["i"])
          );
          this._transition_(compartment);
          return;
        }
        // ! mod3 | mod5
        this._compartment.ExitArgs["output"] = " ";
        let compartment = this.fizzBuzzCompartment(this._FizzBuzzState_Digit_);
        compartment.EnterArgs["i"] = parseInt(e._parameters["i"]);
        this._transition_(compartment);
        return;
      case "<":
        this.print(e._parameters["output"].toString());
        return;
    }
  }
  _FizzBuzzState_Digit_(e) {
    switch (e._message) {
      case ">":
        this.print(e._parameters["i"].toString());
        this.print(" ");
        let compartment = this.fizzBuzzCompartment(this._FizzBuzzState_Fizz_);
        compartment.EnterArgs["i"] = this.plus_1(e._parameters["i"]);
        this._transition_(compartment);
        return;
    }
  }
  _FizzBuzzState_End_(e) {
    switch (e._message) {
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
  print(msg) {
    throw new Error("Action not implemented.");
  }
  gt_100(i) {
    throw new Error("Action not implemented.");
  }
  mod3_eq0(i) {
    throw new Error("Action not implemented.");
  }
  mod5_eq0(i) {
    throw new Error("Action not implemented.");
  }
  plus_1(i) {
    throw new Error("Action not implemented.");
  }

  //=============== Compartment ==============//

  fizzBuzzCompartment(state) {
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


class fizzBuzzController extends FizzBuzz{

  constructor(){
    super()
  }

  print(msg) {
    process.stdout.write(msg)
  }
  gt_100(i) {
    return i > 100
  }
  mod3_eq0(i) {
    return i%3 == 0
  }
  mod5_eq0(i) {
    return i%5 == 0
  }
  plus_1(i) {
    return i + 1
  }
  
}


let fizzbuzz = new fizzBuzzController()

fizzbuzz.start()