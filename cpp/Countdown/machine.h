// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files
#include <iostream>
#include <unordered_map>
#include <any>
#include <string>
#include "frameEvent.h"
using namespace std;

//=============== Compartment ==============//

class CountdownCompartment
{
public:
    int state;
    
    CountdownCompartment(int state)
    {
        this->state = state;
    }
    
    std::unordered_map<std::string, std::any> stateArgs;
    std::unordered_map<std::string, std::any> stateVars;
    std::unordered_map<std::string, std::any> enterArgs;
    std::unordered_map<std::string, std::any> exitArgs;
    FrameEvent *_forwardEvent = nullptr;
};

class Countdown
{
private:
    CountdownCompartment *_compartment_;
    CountdownCompartment *_nextCompartment_;
    
    
    
public:
    Countdown(int i)
    {
        // Create and intialize start state compartment.
        _state_ = static_cast<int>(CountdownState::S0);
        
        _compartment_ = new CountdownCompartment(this->_state_);
        _nextCompartment_ = nullptr;
        _compartment_->stateArgs["i"] = i;
        _compartment_->stateVars["dec"] = 1;
        
        
        // Send system start event
        FrameEvent frame_event(">", std::unordered_map<std::string, std::any>());
        _mux_(&frame_event);
    
    }
    
    // states enum
private:
    enum class CountdownState
    {
        S0 = 0,
        S1 = 1,
        STOP = 2
    };
    
    //====================== Multiplexer ====================//

private:
    void _mux_(FrameEvent *e)
    {
        if(this->_compartment_->state == static_cast<int>(CountdownState::S0))
        {
            this->_sS0_(e);
        }
        else if(this->_compartment_->state == static_cast<int>(CountdownState::S1))
        {
            this->_sS1_(e);
        }
        else if(this->_compartment_->state == static_cast<int>(CountdownState::STOP))
        {
            this->_sStop_(e);
        }
        
        
        if(this->_nextCompartment_ != nullptr)
        {
            CountdownCompartment *nextCompartment = this->_nextCompartment_;
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
    
        
    
    //===================== Machine Block ===================//

private:
    
    void _sS0_(FrameEvent *e)
    {
        if (e->_message == ">") {
            this->_compartment_->stateArgs["i"] = any_cast<int>(this->_compartment_->stateArgs["i"]) - (any_cast<int>(this->_compartment_->stateVars["dec"]));
            print_do(to_string(any_cast<int>(this->_compartment_->stateArgs["i"])));
            if (any_cast<int>(this->_compartment_->stateArgs["i"]) == 0) {
                CountdownCompartment *compartment =  new CountdownCompartment(static_cast<int>(CountdownState::STOP));
                
                this->_transition_(compartment);
                return;
            }
            CountdownCompartment *compartment =  new CountdownCompartment(static_cast<int>(CountdownState::S1));
            compartment->enterArgs["i"] = any_cast<int>(this->_compartment_->stateArgs["i"]);
            
            this->_transition_(compartment);
            return;
        }
    }
    
    void _sS1_(FrameEvent *e)
    {
        if (e->_message == ">") {
            CountdownCompartment *compartment =  new CountdownCompartment(static_cast<int>(CountdownState::S0));
            compartment->stateArgs["i"] = any_cast<int>(e->_parameters["i"]);
            compartment->stateVars["dec"] = 1;
            
            this->_transition_(compartment);
            return;
        }
    }
    
    void _sStop_(FrameEvent *e)
    {
        if (e->_message == ">") {
            print_do("done");
            return;
        }
    }

//===================== Actions Block ===================//
    
    
    
public:
    void print_do(string s)
    {
        std::cout << s << std::endl;
    }
    
    // Unimplemented Actions
    
    
    
    //=============== Machinery and Mechanisms ==============//

private:
    int _state_;
    
    void _transition_(CountdownCompartment *compartment)
    {
        _nextCompartment_ = compartment;
    }
    
    void _doTransition_(CountdownCompartment *nextCompartment)
    {
        this->_mux_(new FrameEvent("<", _compartment_->exitArgs));
        this->_compartment_ = nextCompartment;
        this->_mux_(new FrameEvent(">", this->_compartment_->enterArgs));
    }
    
};

class CountdownController : public Countdown
{
public:
	CountdownController(int i) : Countdown(i) {}
};



















