#include <string>
#include <unordered_map>

class FrameEvent {
public:
    std::string _message;
    std::unordered_map<std::string, int> _parameters;
    void* _return;

    FrameEvent(std::string message, std::unordered_map<std::string, int> parameters) :
        _message(message), _parameters(parameters) {}

    FrameEvent() {}
};

