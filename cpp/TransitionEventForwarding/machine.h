// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files

#include <iostream>
#include <unordered_map>
#include <string>
#include <any>
#include "frameEvent.h"
using namespace std;

//=============== Compartment ==============//

class TransitionEventForwardingCompartment
{
public:
    int state;
    
    TransitionEventForwardingCompartment(int state)
    {
        this->state = state;
    }
    
    std::unordered_map<std::string, std::any> stateArgs;
    std::unordered_map<std::string, std::any> stateVars;
    std::unordered_map<std::string, std::any> enterArgs;
    std::unordered_map<std::string, std::any> exitArgs;
    FrameEvent *_forwardEvent = nullptr;
};

class TransitionEventForwarding
{
private:
    TransitionEventForwardingCompartment *_compartment_;
    TransitionEventForwardingCompartment *_nextCompartment_;
    
    
    
public:
    TransitionEventForwarding(int cycles)
    {
        // Create and intialize start state compartment.
        _state_ = static_cast<int>(TransitionEventForwardingState::START);
        
        _compartment_ = new TransitionEventForwardingCompartment(this->_state_);
        _nextCompartment_ = nullptr;
        _compartment_->enterArgs["cycles"] = cycles;
        
        
        // Send system start event
        FrameEvent frame_event(">", this->_compartment_->enterArgs);
        _mux_(&frame_event);
    
    }
    
    // states enum
private:
    enum class TransitionEventForwardingState
    {
        START = 0,
        FORWARDEVENTAGAIN = 1,
        DECREMENT = 2,
        STOP = 3
    };
    
    //====================== Multiplexer ====================//

private:
    void _mux_(FrameEvent *e)
    {
        if(this->_compartment_->state == static_cast<int>(TransitionEventForwardingState::START))
        {
            this->_sStart_(e);
        }
        else if(this->_compartment_->state == static_cast<int>(TransitionEventForwardingState::FORWARDEVENTAGAIN))
        {
            this->_sForwardEventAgain_(e);
        }
        else if(this->_compartment_->state == static_cast<int>(TransitionEventForwardingState::DECREMENT))
        {
            this->_sDecrement_(e);
        }
        else if(this->_compartment_->state == static_cast<int>(TransitionEventForwardingState::STOP))
        {
            this->_sStop_(e);
        }
        
        
        if(this->_nextCompartment_ != nullptr)
        {
            TransitionEventForwardingCompartment *nextCompartment = this->_nextCompartment_;
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
    
    void _sStart_(FrameEvent *e)
    {
        if (e->_message == ">") {
            if (any_cast<int>(e->_parameters["cycles"]) == 0) {
                this->_compartment_->exitArgs["msg"] = string("stopping");
                TransitionEventForwardingCompartment *compartment =  new TransitionEventForwardingCompartment(static_cast<int>(TransitionEventForwardingState::STOP));
                
                compartment->_forwardEvent = e;
                
                this->_transition_(compartment);
                return;
            } else {
                this->_compartment_->exitArgs["msg"] = string("keep going");
                TransitionEventForwardingCompartment *compartment =  new TransitionEventForwardingCompartment(static_cast<int>(TransitionEventForwardingState::FORWARDEVENTAGAIN));
                
                compartment->_forwardEvent = e;
                
                this->_transition_(compartment);
            }
            return;
        }
        else if (e->_message == "<") {
            print_do(any_cast<string>(e->_parameters["msg"]));
            return;
        }
    }
    
    void _sForwardEventAgain_(FrameEvent *e)
    {
        if (e->_message == ">") {
            TransitionEventForwardingCompartment *compartment =  new TransitionEventForwardingCompartment(static_cast<int>(TransitionEventForwardingState::DECREMENT));
            
            compartment->_forwardEvent = e;
            
            this->_transition_(compartment);
            return;
        }
    }
    
    void _sDecrement_(FrameEvent *e)
    {
        if (e->_message == ">") {
            print_do(to_string(any_cast<int>(e->_parameters["cycles"])));
            TransitionEventForwardingCompartment *compartment =  new TransitionEventForwardingCompartment(static_cast<int>(TransitionEventForwardingState::START));
            compartment->enterArgs["cycles"] = {any_cast<int>(e->_parameters["cycles"]) - 1};
            
            this->_transition_(compartment);
            return;
        }
    }
    
    void _sStop_(FrameEvent *e)
    {
        if (e->_message == ">") {
            print_do(to_string(any_cast<int>(e->_parameters["cycles"])));
            print_do("done");
            return;
        }
    }

//===================== Actions Block ===================//
    
    
    
public:
    void print_do(string msg)
    {
        cout << msg << endl;
    }
    
    // Unimplemented Actions
    
    
    
    //=============== Machinery and Mechanisms ==============//

private:
    int _state_;
    
    void _transition_(TransitionEventForwardingCompartment *compartment)
    {
        _nextCompartment_ = compartment;
    }
    
    void _doTransition_(TransitionEventForwardingCompartment *nextCompartment)
    {
        this->_mux_(new FrameEvent("<", _compartment_->exitArgs));
        this->_compartment_ = nextCompartment;
        this->_mux_(new FrameEvent(">", this->_compartment_->enterArgs));
    }
    
};

class TransitionEventForwardingController : public TransitionEventForwarding
{
public:
	TransitionEventForwardingController(int cycles) : TransitionEventForwarding(cycles) {}
};


