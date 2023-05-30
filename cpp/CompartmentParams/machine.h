// // emitted from framec_v0.10.0
// // get include files at https://github.com/frame-lang/frame-ancillary-files
// #include <iostream>
// #include <unordered_map>
// #include <string>
// #include "frameEvent.h"

// //=============== Compartment ==============//

// class CompartmentParamsCompartment
// {
// public:
//     int state;
    
//     CompartmentParamsCompartment(int state)
//     {
//         this->state = state;
//     }
    
//     std::unordered_map<std::string, std::any> stateArgs;
//     std::unordered_map<std::string, std::any> stateVars;
//     std::unordered_map<std::string, std::any> enterArgs;
//     std::unordered_map<std::string, std::any> exitArgs;
//     FrameEvent *_forwardEvent = nullptr;
// };

// class CompartmentParams
// {
// private:
//     CompartmentParamsCompartment *_compartment_;
//     CompartmentParamsCompartment *_nextCompartment_;
    
    
    
// public:
//     CompartmentParams(int state_param, int enter_param)
//     {
//         // Create and intialize start state compartment.
//         _state_ = static_cast<int>(CompartmentParamsState::S0);
        
//         _compartment_ = new CompartmentParamsCompartment(this->_state_);
//         _nextCompartment_ = nullptr;
//         _compartment_->stateArgs["state_param"] = state_param;
//         _compartment_->stateVars["state_var"] = 100;
//         _compartment_->enterArgs["enter_param"] = enter_param;
        
        
//         // Send system start event
//         FrameEvent frame_event(">", this->_compartment_->enterArgs);
//         _mux_(&frame_event);
    
//     }
    
//     // states enum
// private:
//     enum class CompartmentParamsState
//     {
//         S0 = 0,
//         S1 = 1
//     };
    
//     //====================== Multiplexer ====================//

// private:
//     void _mux_(FrameEvent *e)
//     {
//         if(this->_compartment_->state == static_cast<int>(CompartmentParamsState::S0))
//         {
//             this->_sS0_(e);
//         }
//         else if(this->_compartment_->state == static_cast<int>(CompartmentParamsState::S1))
//         {
//             this->_sS1_(e);
//         }
        
        
//         if(this->_nextCompartment_ != nullptr)
//         {
//             CompartmentParamsCompartment *nextCompartment = this->_nextCompartment_;
//             this->_nextCompartment_ = nullptr;
//             if (nextCompartment->_forwardEvent != nullptr && 
//                 nextCompartment->_forwardEvent->_message == ">")
//             {
//                 this->_mux_(new FrameEvent( "<", this->_compartment_->exitArgs));
//                 this->_compartment_ = nextCompartment;
//                 this->_mux_(nextCompartment->_forwardEvent);
//             }
//             else
//             {
//                 this->_doTransition_(nextCompartment);
//                 if (nextCompartment->_forwardEvent != nullptr)
//                 {
//                     this->_mux_(nextCompartment->_forwardEvent);
//                 }
//             }
//             nextCompartment->_forwardEvent = nullptr;
//         }
//     }
    
        
    
//     //===================== Machine Block ===================//

// private:
    
//     void _sS0_(FrameEvent *e)
//     {
//         if (e->_message == ">") {
//             print_do(std::to_string(any_cast<int>(this->_compartment_->stateArgs["state_param"])) + " " + std::to_string(any_cast<int>(this->_compartment_->stateVars["state_var"])) + " " + std::to_string(any_cast<int>(e->_parameters["enter_param"])));
//             CompartmentParamsCompartment *compartment =  new CompartmentParamsCompartment(static_cast<int>(CompartmentParamsState::S1));
            
//             compartment->_forwardEvent = e;
//             compartment->stateArgs["state_param"] = any_cast<int>(this->_compartment_->stateArgs["state_param"]) + 20;
//             compartment->stateVars["state_var"] = 200;
            
//             this->_transition_(compartment);
//             return;
//         }
//     }
    
//     void _sS1_(FrameEvent *e)
//     {
//         if (e->_message == ">") {
//             print_do(std::to_string(any_cast<int>(this->_compartment_->stateArgs["state_param"])) + " " + std::to_string(any_cast<int>(this->_compartment_->stateVars["state_var"])) + " " + std::to_string(any_cast<int>(e->_parameters["enter_param"])));
//             return;
//         }
//     }

// //===================== Actions Block ===================//
    
    
    
// public:
//     void print_do(std::string s)
//     {
//          std::cout << s << std::endl;
//     }
    
//     // Unimplemented Actions
    
    
    
//     //=============== Machinery and Mechanisms ==============//

// private:
//     int _state_;
    
//     void _transition_(CompartmentParamsCompartment *compartment)
//     {
//         _nextCompartment_ = compartment;
//     }
    
//     void _doTransition_(CompartmentParamsCompartment *nextCompartment)
//     {
//         this->_mux_(new FrameEvent("<", _compartment_->exitArgs));
//         this->_compartment_ = nextCompartment;
//         this->_mux_(new FrameEvent(">", this->_compartment_->enterArgs));
//     }
    
// };


// class CompartmentParamsController : public CompartmentParams
// {
// public:
// 	CompartmentParamsController(int state_param, int enter_param) : CompartmentParams(state_param, enter_param) {}
// };
// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files
// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files
// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files
#include <iostream>
#include <unordered_map>
#include <string>
#include "frameEvent.h"

//=============== Compartment ==============//

class CompartmentParamsCompartment
{
public:
    int state;
    
    CompartmentParamsCompartment(int state)
    {
        this->state = state;
    }
    
    std::unordered_map<std::string, std::any> stateArgs;
    std::unordered_map<std::string, std::any> stateVars;
    std::unordered_map<std::string, std::any> enterArgs;
    std::unordered_map<std::string, std::any> exitArgs;
    FrameEvent *_forwardEvent = nullptr;
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
        _mux_(&frame_event);
    
    }
    
    // states enum
private:
    enum class CompartmentParamsState
    {
        S0 = 0,
        S1 = 1
    };
    
    //====================== Multiplexer ====================//

private:
    void _mux_(FrameEvent *e)
    {
        if(this->_compartment_->state == static_cast<int>(CompartmentParamsState::S0))
        {
            this->_sS0_(e);
        }
        else if(this->_compartment_->state == static_cast<int>(CompartmentParamsState::S1))
        {
            this->_sS1_(e);
        }
        
        
        if(this->_nextCompartment_ != nullptr)
        {
            CompartmentParamsCompartment *nextCompartment = this->_nextCompartment_;
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
            print_do(to_string(any_cast<int>(this->_compartment_->stateArgs["state_param"])) + " " + to_string((any_cast<int>(this->_compartment_->stateVars["state_var"]))) + " " + to_string(any_cast<int>(e->_parameters["enter_param"])));
            CompartmentParamsCompartment *compartment =  new CompartmentParamsCompartment(static_cast<int>(CompartmentParamsState::S1));
            
            compartment->_forwardEvent = e;
            compartment->stateArgs["state_param"] = any_cast<int>(this->_compartment_->stateArgs["state_param"]) + 20;
            compartment->stateVars["state_var"] = 200;
            
            this->_transition_(compartment);
            return;
        }
    }
    
    void _sS1_(FrameEvent *e)
    {
        if (e->_message == ">") {
            print_do(to_string(any_cast<int>(this->_compartment_->stateArgs["state_param"])) + " " + to_string((any_cast<int>(this->_compartment_->stateVars["state_var"]))) + " " + to_string(any_cast<int>(e->_parameters["enter_param"])));
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
        this->_mux_(new FrameEvent("<", _compartment_->exitArgs));
        this->_compartment_ = nextCompartment;
        this->_mux_(new FrameEvent(">", this->_compartment_->enterArgs));
    }
    
};

class CompartmentParamsController : public CompartmentParams
{
public:
	CompartmentParamsController(int state_param, int enter_param) : CompartmentParams(state_param, enter_param) {}
};















