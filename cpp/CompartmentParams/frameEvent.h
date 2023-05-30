#include <string>
#include <unordered_map>
#include <any>
using namespace std;

class FrameEvent {
public:
    std::string _message;
    std::unordered_map<std::string, std::any> _parameters;
    void* _return;

    FrameEvent(std::string message, std::unordered_map<std::string, std::any> parameters) :
        _message(message), _parameters(parameters) {}

    FrameEvent() {}
};

