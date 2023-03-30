#include <iostream>
#include <unordered_map>
#include <string>
#include "frameEvent.h"

class CompartmentParamsCompartment
{
public:
    int state;

    CompartmentParamsCompartment(int state)
    {
        this->state = state;
    }

    std::unordered_map<std::string, int> stateArgs;
    std::unordered_map<std::string, int> stateVars;
    std::unordered_map<std::string, int> enterArgs;
    std::unordered_map<std::string, int> exitArgs;
    FrameEvent *_forwardEvent;
};

class CompartmentParams
{
private:
    CompartmentParamsCompartment *_compartment_;
    CompartmentParamsCompartment *_nextCompartment_;

public:
    CompartmentParams(int state_param, int enter_param)
    {
        // Create and intialize start state compartment.
        _state_ = static_cast<int>(CompartmentParamsState::S0);
        _compartment_ = new CompartmentParamsCompartment(this->_state_);
        _nextCompartment_ = nullptr;
        _compartment_->stateArgs["state_param"] = state_param;
        _compartment_->stateVars["state_var"] = 100;
        _compartment_->enterArgs["enter_param"] = enter_param;

        // Send system start event
        FrameEvent frame_event(">", this->_compartment_->enterArgs);
        _mux_(frame_event);
    }

private:
    enum class CompartmentParamsState
    {
        S0 = 0,
        S1 = 1
    };

    //====================== Multiplexer ====================//
private:
    void _mux_(FrameEvent e)
    {
        if (this->_compartment_->state == static_cast<int>(CompartmentParamsState::S0))
        {
            this->_sS0_(e);
        }
        else if (this->_compartment_->state == static_cast<int>(CompartmentParamsState::S1))
        {
            this->_sS1_(e);
        }

        if (this->_nextCompartment_ != nullptr)
        {
            CompartmentParamsCompartment *nextCompartment = this->_nextCompartment_;
            this->_nextCompartment_ = nullptr;
            if (nextCompartment->_forwardEvent != nullptr &&
                nextCompartment->_forwardEvent->_message == ">")
            {
                this->_mux_(FrameEvent("<", this->_compartment_->exitArgs));
                this->_compartment_ = nextCompartment;
                this->_mux_(*nextCompartment->_forwardEvent);
            }
            else
            {
                this->_doTransition_(nextCompartment);
                if (nextCompartment->_forwardEvent != nullptr)
                {
                    this->_mux_(*nextCompartment->_forwardEvent);
                }
            }
            nextCompartment->_forwardEvent = nullptr;
        }
    }

    //===================== Machine Block ===================//
private:
    void _sS0_(FrameEvent e)
    {
        if (e._message == ">")
        {
            print_do(std::to_string((int)this->_compartment_->stateArgs["state_param"]) + " " +
                     std::to_string((int)this->_compartment_->stateVars["state_var"]) + " " +
                     std::to_string((int)e._parameters["enter_param"]));
            CompartmentParamsCompartment *compartment =
                new CompartmentParamsCompartment(this->_state_ + 1);

            compartment->_forwardEvent = new FrameEvent(e);
            compartment->stateArgs["state_param"] = (int)this->_compartment_->stateArgs["state_param"] + 20;
            compartment->stateVars["state_var"] = 200;

            this->_transition_(compartment);
            return;
        }
    }

private:
    void _sS1_(FrameEvent e)
    {
        if (e._message == ">")
        {
            print_do(std::to_string((int)this->_compartment_->stateArgs["state_param"]) + " " +
                     std::to_string((int)this->_compartment_->stateVars["state_var"]) + " " +
                     std::to_string((int)e._parameters["enter_param"]));
            return;
        }
    }

    //===================== Actions Block ===================//

public:
    void print_do(std::string s)
    {
        std::cout << s << std::endl;
    }

    // Unimplemented Actions

    //=============== Machinery and Mechanisms ==============//

private:
    int _state_;

    void _transition_(CompartmentParamsCompartment *compartment)
    {
        _nextCompartment_ = compartment;
    }

    void _doTransition_(CompartmentParamsCompartment *nextCompartment)
    {
        this->_mux_(FrameEvent("<", _compartment_->exitArgs));
        this->_compartment_ = nextCompartment;
        this->_mux_(FrameEvent(">", _compartment_->enterArgs));
        // _nextCompartment_ = nullptr;
    }
};
//=============== Compartment ==============//

class CompartmentParamsController : public CompartmentParams
{
public:
    CompartmentParamsController(int state_param, int enter_param) : CompartmentParams(state_param, enter_param) {}
};
