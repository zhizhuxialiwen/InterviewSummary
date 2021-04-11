#ifndef FOOINTERFACE_H_
#define FOOINTERFACE_H_
 
#include <string>
 
namespace seamless {
 
class FooInterface {
public:
    virtual std::string getArbitraryString() = 0;
    virtual ~FooInterface() {}
        
};
 
}  // namespace seamless
 
#endif // FOOINTERFACE_H_