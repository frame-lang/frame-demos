
// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files
#include <iostream>
#include <unordered_map>
#include <string>
#include "frameEvent.h"

//=============== Compartment ==============//

class FizzBuzzCompartment
{
public:
    int state;
    
    FizzBuzzCompartment(int state)
    {
        this->state = state;
    }
    
    std::unordered_map<std::string, std::any> stateArgs;
    std::unordered_map<std::string, std::any> stateVars;
    std::unordered_map<std::string, std::any> enterArgs;
    std::unordered_map<std::string, std::any> exitArgs;
    FrameEvent *_forwardEvent = nullptr;
};

class FizzBuzz
{
private:
    FizzBuzzCompartment *_compartment_;
    FizzBuzzCompartment *_nextCompartment_;
    
    
    
public:
    FizzBuzz()
    {
        // Create and intialize start state compartment.
        _state_ = static_cast<int>(FizzBuzzState::BEGIN);
        
        _compartment_ = new FizzBuzzCompartment(this->_state_);
        _nextCompartment_ = nullptr;
        
        
        // Send system start event
        FrameEvent frame_event(">", std::unordered_map<std::string, std::any>());
        _mux_(&frame_event);
    
    }
    
    // states enum
private:
    enum class FizzBuzzState
    {
        BEGIN = 0,
        FIZZ = 1,
        BUZZ = 2,
        DIGIT = 3,
        END = 4
    };
    
    //====================== Multiplexer ====================//

private:
    void _mux_(FrameEvent *e)
    {
        if(this->_compartment_->state == static_cast<int>(FizzBuzzState::BEGIN))
        {
            this->_sBegin_(e);
        }
        else if(this->_compartment_->state == static_cast<int>(FizzBuzzState::FIZZ))
        {
            this->_sFizz_(e);
        }
        else if(this->_compartment_->state == static_cast<int>(FizzBuzzState::BUZZ))
        {
            this->_sBuzz_(e);
        }
        else if(this->_compartment_->state == static_cast<int>(FizzBuzzState::DIGIT))
        {
            this->_sDigit_(e);
        }
        else if(this->_compartment_->state == static_cast<int>(FizzBuzzState::END))
        {
            this->_sEnd_(e);
        }
        
        
        if(this->_nextCompartment_ != nullptr)
        {
            FizzBuzzCompartment *nextCompartment = this->_nextCompartment_;
            this->_nextCompartment_ = nullptr;
            if (nextCompartment->_forwardEvent != nullptr && 
                nextCompartment->_forwardEvent->_message == ">")
            {
                this->_mux_(new FrameEvent( "<", this->_compartment_->exitArgs));
                this->_compartment_ = nextCompartment;
                this->_mux_(nextCompartment->_forwardEvent);
            }
            else
            {
                this->_doTransition_(nextCompartment);
                if (nextCompartment->_forwardEvent != nullptr)
                {
                    this->_mux_(nextCompartment->_forwardEvent);
                }
            }
            nextCompartment->_forwardEvent = nullptr;
        }
    }
    
        
        //===================== Interface Block ===================//
    public:
        void start() {
            FrameEvent e(">>", unordered_map<string, any>());
            this->_mux_(&e);
        }
        
    
    //===================== Machine Block ===================//

private:
    
    void _sBegin_(FrameEvent *e)
    {
        if (e->_message == ">>") {
            // start
            FizzBuzzCompartment *compartment =  new FizzBuzzCompartment(static_cast<int>(FizzBuzzState::FIZZ));
            compartment->enterArgs["i"] = 1;
            
            this->_transition_(compartment);
            return;
        }
    }
    
    void _sFizz_(FrameEvent *e)
    {
        if (e->_message == ">") {
            if (gt_100_do(any_cast<int>(e->_parameters["i"]))) {
                // i > 100
                FizzBuzzCompartment *compartment =  new FizzBuzzCompartment(static_cast<int>(FizzBuzzState::END));
                
                this->_transition_(compartment);
                return;
            }
            if (mod3_eq0_do(any_cast<int>(e->_parameters["i"]))) {
                print_do("Fizz");
                // i % 3 == 0
                FizzBuzzCompartment *compartment =  new FizzBuzzCompartment(static_cast<int>(FizzBuzzState::BUZZ));
                compartment->enterArgs["i"] = any_cast<int>(e->_parameters["i"]);
                compartment->enterArgs["fizzed"] = true;
                
                this->_transition_(compartment);
            } else {
                // i % 3 != 0
                FizzBuzzCompartment *compartment =  new FizzBuzzCompartment(static_cast<int>(FizzBuzzState::BUZZ));
                compartment->enterArgs["i"] = any_cast<int>(e->_parameters["i"]);
                compartment->enterArgs["fizzed"] = false;
                
                this->_transition_(compartment);
            }
            return;
        }
    }
    
    void _sBuzz_(FrameEvent *e)
    {
        if (e->_message == ">") {
            if (mod5_eq0_do(any_cast<int>(e->_parameters["i"]))) {
                print_do("Buzz");
                // i % 5 == 0
                this->_compartment_->exitArgs["output"] = string(" ");
                FizzBuzzCompartment *compartment =  new FizzBuzzCompartment(static_cast<int>(FizzBuzzState::FIZZ));
                compartment->enterArgs["i"] = plus_1_do(any_cast<int>(e->_parameters["i"]));
                
                this->_transition_(compartment);
                return;
            }
            if (any_cast<bool>(e->_parameters["fizzed"])) {
                // fizzed
                this->_compartment_->exitArgs["output"] = string(" ");
                FizzBuzzCompartment *compartment =  new FizzBuzzCompartment(static_cast<int>(FizzBuzzState::FIZZ));
                compartment->enterArgs["i"] = plus_1_do(any_cast<int>(e->_parameters["i"]));
                
                this->_transition_(compartment);
                return;
            }
            // ! mod3 | mod5
            this->_compartment_->exitArgs["output"] = string("");
            FizzBuzzCompartment *compartment =  new FizzBuzzCompartment(static_cast<int>(FizzBuzzState::DIGIT));
            compartment->enterArgs["i"] = any_cast<int>(e->_parameters["i"]);
            
            this->_transition_(compartment);
            return;
        }
        else if (e->_message == "<") {
            print_do(any_cast<string>(e->_parameters["output"]));
            return;
        }
    }
    
    void _sDigit_(FrameEvent *e)
    {
        if (e->_message == ">") {
            print_do(to_string(any_cast<int>(e->_parameters["i"])));
            print_do(" ");
            // loop
            FizzBuzzCompartment *compartment =  new FizzBuzzCompartment(static_cast<int>(FizzBuzzState::FIZZ));
            compartment->enterArgs["i"] = plus_1_do(any_cast<int>(e->_parameters["i"]));
            
            this->_transition_(compartment);
            return;
        }
    }
    
    void _sEnd_(FrameEvent *e)
    {
    }

//===================== Actions Block ===================//
    
    
    
public:
    void print_do(string msg)
    {
        cout << msg << " ";
    }
    
    
public:
    bool gt_100_do(int i)
    {
        return i > 100;
    }
    
    
public:
    bool mod3_eq0_do(int i)
    {
        return i % 3 == 0;
    }
    
    
public:
    bool mod5_eq0_do(int i)
    {
        return i % 5 == 0;
    }
    
    
public:
    int plus_1_do(int i)
    {
        return i + 1;
    }
    
    // Unimplemented Actions
    
    
    
    //=============== Machinery and Mechanisms ==============//

private:
    int _state_;
    
    void _transition_(FizzBuzzCompartment *compartment)
    {
        _nextCompartment_ = compartment;
    }
    
    void _doTransition_(FizzBuzzCompartment *nextCompartment)
    {
        this->_mux_(new FrameEvent("<", _compartment_->exitArgs));
        this->_compartment_ = nextCompartment;
        this->_mux_(new FrameEvent(">", this->_compartment_->enterArgs));
    }
    
};

class FizzBuzzController : public FizzBuzz
{
public:
	FizzBuzzController() : FizzBuzz() {}
};
















