// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files

#include <iostream>
#include <unordered_map>
#include <any>
#include <string>
#include "frameEvent.h"
using namespace std;

//=============== Compartment ==============//

class SystemParamsCompartment
{
public:
    int state;
    
    SystemParamsCompartment(int state)
    {
        this->state = state;
    }
    
    std::unordered_map<std::string, std::any> stateArgs;
    std::unordered_map<std::string, std::any> stateVars;
    std::unordered_map<std::string, std::any> enterArgs;
    std::unordered_map<std::string, std::any> exitArgs;
    FrameEvent *_forwardEvent = nullptr;
};

class SystemParams
{
private:
    SystemParamsCompartment *_compartment_;
    SystemParamsCompartment *_nextCompartment_;
    
    
    
public:
    SystemParams(string stateMsg, string enterMsg)
    {
        // Create and intialize start state compartment.
        _state_ = static_cast<int>(SystemParamsState::BEGIN);
        
        _compartment_ = new SystemParamsCompartment(this->_state_);
        _nextCompartment_ = nullptr;
        _compartment_->stateArgs["stateMsg"] = stateMsg;
        _compartment_->enterArgs["enterMsg"] = enterMsg;
        
        
        // Send system start event
        FrameEvent frame_event(">", this->_compartment_->enterArgs);
        _mux_(&frame_event);
    
    }
    
    // states enum
private:
    enum class SystemParamsState
    {
        BEGIN = 0
    };
    
    //====================== Multiplexer ====================//

private:
    void _mux_(FrameEvent *e)
    {
        if(this->_compartment_->state == static_cast<int>(SystemParamsState::BEGIN))
        {
            this->_sBegin_(e);
        }
        
        
        if(this->_nextCompartment_ != nullptr)
        {
            SystemParamsCompartment *nextCompartment = this->_nextCompartment_;
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
    
    void _sBegin_(FrameEvent *e)
    {
        if (e->_message == ">") {
            print_do(any_cast<string>(this->_compartment_->stateArgs["stateMsg"]) + " " + any_cast<string>(e->_parameters["enterMsg"]));
            return;
        }
    }

//===================== Actions Block ===================//
    
    
    
public:
    void print_do(string msg)
    {
        cout << msg << " ";
    }
    
    // Unimplemented Actions
    
    
    
    //=============== Machinery and Mechanisms ==============//

private:
    int _state_;
    
    void _transition_(SystemParamsCompartment *compartment)
    {
        _nextCompartment_ = compartment;
    }
    
    void _doTransition_(SystemParamsCompartment *nextCompartment)
    {
        this->_mux_(new FrameEvent("<", _compartment_->exitArgs));
        this->_compartment_ = nextCompartment;
        this->_mux_(new FrameEvent(">", this->_compartment_->enterArgs));
    }
    
};

class SystemParamsController : public SystemParams
{
public:
	SystemParamsController(string stateMsg, string enterMsg) : SystemParams(stateMsg, enterMsg) {}
};





